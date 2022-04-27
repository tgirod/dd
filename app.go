package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asdine/storm/v3"
	gc "github.com/asdine/storm/v3/codec/gob"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

type App struct {
	s *ssh.Server
	Game
}

// NewApp créé un nouvel objet application
func NewApp() *App {
	var err error
	a := new(App)

	if a.s, err = wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(a.Handler),
			lm.Middleware(),
		),
	); err != nil {
		log.Fatal(err)
	}

	return a
}

// Start démarre le serveur, en attente de connexions
func (a *App) Start() {
	var err error

	// ouverture de la BDD
	gob.Register(Node{})
	gob.Register(Connect{})
	gob.Register(Help{})
	gob.Register(Index{})
	gob.Register(Quit{})
	gob.Register(LinkList{})
	gob.Register(LinkConnect{})
	gob.Register(DataSearch{})
	gob.Register(Jack{})
	gob.Register(Rise{})
	db, err := storm.Open("game.db", storm.Codec(gc.Codec))
	if err != nil {
		log.Fatal(err)
	}
	a.Game = Game{db}
	if err := a.Game.Init(); err != nil {
		// FIXME a invoquer uniquement avec un argument "init"
		log.Fatal(err)
	}
	defer a.Game.Close()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err = a.s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := a.s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

// Handler prend en charge la connexion entrante et créé les objets nécessaires
func (a *App) Handler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// si le terminal qui tente de se connecter est invalide
	pty, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}

	// création de l'interface utilisateur
	client := NewClient(
		pty.Window.Width,
		pty.Window.Height,
		a.Game,
	)

	return client, []tea.ProgramOption{tea.WithAltScreen()}
}
