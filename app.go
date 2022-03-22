package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

type App struct {
	s       *ssh.Server
	network *Network // la représentation du réseau
}

// NewApp créé un nouvel objet application
func NewApp() *App {
	var err error
	a := new(App)

	a.s, err = wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(a.handler),
			lm.Middleware(),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}
	return a
}

// Start démarre le serveur, en attente de connexions
func (a *App) Start() {
	var err error
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

// handler prend en charge la connexion entrante et créé les objets nécessaires
func (a *App) handler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// si le terminal qui tente de se connecter est invalide
	pty, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}

	// l'objet console est le modèle utilisé par bubbletea
	model := NewConsole(pty)
	model.network = a.network

	return model, []tea.ProgramOption{}
}
