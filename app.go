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
	errInvalidUser        = errors.New("compte utilisateur invalide")
	errNotConnected       = errors.New("cette commande nécessite d'être connecté")
	errConnected          = errors.New("la console est connectée")
	errLowPrivilege       = errors.New("niveau de privilège insuffisant")
	errEntryNotFound      = errors.New("entrée introuvable")
	errMemNotFound        = errors.New("zone mémoire introuvable")
	errMemUnavailable     = errors.New("zone mémoire indisponible")
	errEmptyHistory       = errors.New("historique de navigation vide")
	errKeywordTooShort    = errors.New("mot clef trop court")
	errRegisterNotFound   = errors.New("registre introuvable")
	errNotIdentified      = errors.New("cette commande nécessite une identité active")
	errLowCredit          = errors.New("crédit insuffisant")
	errNegativeAmount     = errors.New("montant négatif")
	errMessageNotFound    = errors.New("message introuvable")
	errForumUnreachable   = errors.New("Forum inatteignable")
	errTopicExists        = errors.New("TOPIC existe déjà")
	errPostNotFound       = errors.New("Post introuvable")
)

type App struct {
	sPlayer     *ssh.Server
	startTime   time.Time
	sMonitor    *ssh.Server
	sessions    map[ssh.Session]*Console
	admin *tea.Program
}

// NewApp créé un nouvel objet application
func NewApp(init bool) *App {
	if init {
		Init()
	}

	var err error
	a := new(App)
	// MON
	a.startTime = time.Now()
	a.sessions = make(map[ssh.Session]*Console)

	// SSH server for the players
	if a.sPlayer, err = wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, portPlayer)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(a.HandlerPlayer),
			MiddlewareMonitor(a.sessions),
			// lm.Midleware(),
		),
	); err != nil {
		log.Fatal(err)
	}

	// SSH server for monitoring
	if a.sMonitor, err = wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, portMonitor)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(a.HandlerMonitor),
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
	log.Printf("lancement du serveur SSH sur %s:%d", host, portPlayer)
	go func() {
		if err := a.sPlayer.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
	// Lancement du serveur SSH de monitoring
	go func() {
		if err := a.sMonitor.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	// DEL list posts
	// SerializePosts("dd.local")
	// LoadPosts("init_post.yaml")

	log.Println("lancement de l'interface admin")
	if _, err := a.admin.Run(); err != nil {
		log.Fatalln(err)
	}

	log.Println("fermeture du serveur SSH")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := a.sPlayer.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func (a *App) Log(v any) {
	a.admin.Send(LogMsg(fmt.Sprintf("%+v", v)))
}

// MON afficher les sessions ouvertes
func (a App) PrintSessions() {
	for _, e := range a.sessions {
		fmt.Println("Session with Console ID ", e.ID)
	}
}

// Handler prend en charge la connexion entrante et créé les objets nécessaires
// MON Ajoute une session
func (a *App) HandlerPlayer(s ssh.Session) (tea.Model, []tea.ProgramOption) {
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
		false,
	)

	// MON ajout session
	a.sessions[s] = client.Console
	a.PrintSessions()

	return client, []tea.ProgramOption{tea.WithAltScreen()}
}

// MON On a besoin d'un Middleware pour deleter en cas de déconnexion
func MiddlewareMonitor(conn map[ssh.Session]*Console) wish.Middleware {
	return func(sh ssh.Handler) ssh.Handler {
		return func(s ssh.Session) {
			fmt.Printf("Begin of adventure for %s (%d conn)\n", s.User(), len(conn))
			sh(s)
			fmt.Printf("End of adventure for %s (%d conn)\n", s.User(), len(conn))
			delete(conn, s)
		}
	}
}

// MON Handler prend en charge la connexion Monitor entrante
// et créé les objets nécessaires
func (a *App) HandlerMonitor(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// si le terminal qui tente de se connecter est invalide
	pty, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}

	// création de l'interface utilisateur
	monitor := NewMonitor(
		a.startTime,
		pty.Window.Width,
		pty.Window.Height,
		//a.Game,
		a.sessions,
	)
	// superuser
	//monitor.Console.MakeMatrix()

	return monitor, []tea.ProgramOption{tea.WithAltScreen()}
}
