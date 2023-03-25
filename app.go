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
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
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
	errInvalidIdentity    = errors.New("identifiant ou mot de passe invalide")
	errNotConnected       = errors.New("la console n'est pas connectée")
	errConnected          = errors.New("la console est connectée")
	errLowPrivilege       = errors.New("niveau de privilège insuffisant")
	errEntryNotFound      = errors.New("entrée introuvable")
	errMemNotFound        = errors.New("zone mémoire introuvable")
	errMemUnavailable     = errors.New("zone mémoire indisponible")
	errEmptyHistory       = errors.New("historique de navigation vide")
	errKeywordTooShort    = errors.New("mot clef trop court")
	errRegisterNotFound   = errors.New("registre introuvable")
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
func (a *App) Start(filename string) {

	// UnSerialize from given file
	if filename != "" {
		a.Game.UnSerialize(filename)
		log.Printf("Reading Game Stat from %s", filename)
	}

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
