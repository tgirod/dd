package main

import (
	"context"
	"errors"
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

	//"encoding/json"
)

var (
	errInternalError      = errors.New("erreur interne")
	errServerNotFound     = errors.New("serveur introuvable")
	errServiceNotFound    = errors.New("serveur introuvable")
	errInvalidCommand     = errors.New("commande invalide")
	errMissingCommand     = errors.New("commande manquante")
	errMissingArgument    = errors.New("argument manquant")
	errInvalidArgument    = errors.New("argument invalide")
	errInvalidCredentials = errors.New("identifiant ou mot de passe invalide")
	errNotConnected       = errors.New("la console n'est pas connectée")
	errConnected          = errors.New("la console est connectée")
	errLowPrivilege       = errors.New("niveau de privilège insuffisant")
	errEntryNotFound      = errors.New("entrée introuvable")
	errMemNotFound        = errors.New("zone mémoire introuvable")
	errMemUnavailable     = errors.New("zone mémoire indisponible")
	errEmptyHistory       = errors.New("historique de navigation vide")
)

type App struct {
	s *ssh.Server
	*Game
}

// NewApp créé un nouvel objet application
func NewApp() *App {
	var err error
	a := new(App)
	a.Game = game

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

	// bytes, err := json.Marshal(a.Game)
    // if err != nil {
    //     fmt.Println("Can't serislize", a.Game)
    // }
	// fmt.Println(string(bytes))
	// fmt.Println(a.Game.Network[0].Address)
	// a.Game.Network[0].Address = "dd.l"
	// a.Game.Serialize()
	a.Game.UnSerialize()
	// Deserialize Game
	fmt.Println(a.Game.Network[0].Address)
	// var ng Game;
	// err = json.Unmarshal(bytes, &ng)
	// if err != nil {
    //     fmt.Println("Can't deserislize", bytes)
    // }
	
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err := a.s.ListenAndServe(); err != nil {
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

// TODO piste pour que le client ait accès à l'objet tea.Program
func (a *App) ProgramHandler(s ssh.Session) *tea.Program {
	return nil
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
