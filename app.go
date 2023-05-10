package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	// lm "github.com/charmbracelet/wish/logging"
	//"encoding/json"
)

var (
	errInternalError      = errors.New("erreur interne")
	errServerNotFound     = errors.New("serveur introuvable")
	errServiceNotFound    = errors.New("serveur introuvable")
	errIdentityNotFound   = errors.New("identité introuvable")
	errInvalidCredentials = errors.New("identifiant ou mot de passe invalide")
	errInvalidIdentity    = errors.New("identifiant ou mot de passe invalide")
	errInvalidAccount     = errors.New("compte utilisateur invalide")
	errNotConnected       = errors.New("la console n'est pas connectée")
	errConnected          = errors.New("la console est connectée")
	errLowPrivilege       = errors.New("niveau de privilège insuffisant")
	errEntryNotFound      = errors.New("entrée introuvable")
	errMemNotFound        = errors.New("zone mémoire introuvable")
	errMemUnavailable     = errors.New("zone mémoire indisponible")
	errEmptyHistory       = errors.New("historique de navigation vide")
	errKeywordTooShort    = errors.New("mot clef trop court")
	errRegisterNotFound   = errors.New("registre introuvable")
	errNotIdentified      = errors.New("aucune identité active")
	errLowCredit          = errors.New("crédit insuffisant")
	errNegativeAmount     = errors.New("montant négatif")
	errMessageNotFound    = errors.New("message introuvable")
	errForumUnreachable   = errors.New("Forum inatteignable")
	errTopicExists        = errors.New("TOPIC existe déjà")
	errPostNotFound       = errors.New("Post introuvable")
)

type App struct {
	s     *ssh.Server
	admin *tea.Program
}

// NewApp créé un nouvel objet application
func NewApp(init bool) *App {
	if init {
		Init()
	}

	var err error
	a := new(App)

	if a.s, err = wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(a.Handler),
			// lm.Midleware(),
		),
	); err != nil {
		log.Fatal(err)
	}

	a.admin = tea.NewProgram(NewAdmin())

	return a
}

// Start démarre le serveur, en attente de connexions
func (a *App) Start() {

	// done := make(chan os.Signal, 1)
	// signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("lancement du serveur SSH sur %s:%d", host, port)
	go func() {
		if err := a.s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("lancement de l'interface admin")
	if _, err := a.admin.Run(); err != nil {
		log.Fatalln(err)
	}

	log.Println("fermeture du serveur SSH")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := a.s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) Log(msg string) {
	a.admin.Send(LogMsg(msg))
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
	)

	return client, []tea.ProgramOption{tea.WithAltScreen()}
}
