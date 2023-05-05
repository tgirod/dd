package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/mattn/go-shellwords"
)

const MAX_RESULTS int = 10

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID int

	// racine de l'arbre des commandes
	Branch

	// informations sur la connexion au réseau
	*Session

	// identité active
	*Identity

	// l'alerte est-elle activée ?
	Alert bool

	// durée avant la déconnexion forcée
	Countdown time.Duration

	// interface neurale directe
	DNI bool

	// liste des dernières commandes évaluées
	Results []Result

	// état interne du jeu
	*Network
}

type Session struct {
	*Server                  // serveur auquel la session se réfère
	*Account                 // compte utilisateur actif dans ce serveur
	Mem      map[string]bool // zones mémoires disponibles pour une évasion
	Parent   *Session        // session précédente
}

func (s Session) Path() string {
	var path []string
	sess := &s
	for sess != nil {
		path = append([]string{sess.Server.Address}, path...)
		sess = sess.Parent
	}
	return strings.Join(path, "/")
}

type Result struct {
	Prompt string
	Error  error
	Output string
}

var Hack = map[string]Cmd{
	"jack":  jack,
	"evade": evade,
	"door":  door,
}

var baseCmds = Branch{
	name: "root",
	cmds: []Cmd{
		back,
		yes,
		connect,
		data,
		identify,
		index,
		link,
		load,
		plug,
		quit,
		registry,
		message,
		forum,
	},
}

// var baseCmds = Cmd{
// 	Name: "root",
// 	SubCmds: []Cmd{
// 		back,
// 		yes,
// 		connect,
// 		data,
// 		help,
// 		identify,
// 		index,
// 		link,
// 		load,
// 		plug,
// 		quit,
// 		registry,
// 		message,
// 		forum,
// 	},
// }

func NewConsole(net *Network) *Console {
	return &Console{
		Branch:  baseCmds,
		Network: net,
	}
}

func (c *Console) Parse(prompt string) any {
	args, err := shellwords.Parse(prompt)
	if err != nil {
		return Result{
			Prompt: prompt,
			Error:  err,
			Output: "",
		}
	}

	ctx := Context{
		parent: nil,
		key:    "console",
		value:  c,
		node:   nil,
	}
	return c.Branch.Parse(ctx, args)
}

func (c *Console) connect(address string) error {
	var err error
	var server *Server
	var account *Account

	// récupérer le serveur
	if server, err = c.FindServer(address); err != nil {
		return fmt.Errorf("%s : %w", address, err)
	}

	// vérifier que l'utilisateur a le droit de se connecter
	login := ""
	if c.Identity != nil {
		login = c.Identity.Login
	}
	if account, err = server.CheckAccount(login); err != nil {
		return err
	}
	c.Account = account

	// enregistrer le nouveau serveur
	c.Server = server

	if c.Account != nil && c.Account.Backdoor {
		c.Server.RemoveAccount(c.Account.Login)
		c.Network.RemoveIdentity(c.Account.Login)
	}
	c.InitMem()
	return nil
}

func (s *Session) InitMem() {
	s.Mem = make(map[string]bool)
	for i := 0; i < 5; i++ {
		addr := fmt.Sprintf("%08x", rand.Uint32())
		s.Mem[addr] = true
	}
}

func (c *Console) Disconnect() {
	c.Server = nil
	c.Identity = nil
	c.Admin = false
	c.Alert = false
	c.Session = nil
	// BUG
	// c.Branch = baseCmds

	// affichage par défaut
	eval := Result{
		Output: "coupure de la connexion au réseau.",
	}

	if c.DNI {
		eval.Output = `
			     DUMPSHOCK !!!!
                     _____
                    /     \
                   | () () |
                    \  ^  /
                     |||||
                     |||||

			PERDS UN POINT DE VIE

coupure de la connexion au réseau.`
	}

	c.AddResult(eval)
}

func (c *Console) StartAlert() {
	if !c.Alert {
		c.Alert = true
		c.Countdown = c.Server.Scan
	} else if c.Server.Scan < c.Countdown {
		c.Countdown = c.Server.Scan
	}
}

func (c *Console) Identify(login, password string) error {
	identity, err := c.CheckIdentity(login, password)
	if err != nil {
		return err
	}
	c.Identity = identity

	// si on est connecté à un serveur, on tente d'accéder au compte utilisateur
	if c.Server != nil {
		if account, err := c.CheckAccount(login); err == nil {
			c.Account = account
		}
	}

	return nil
}

func (c *Console) AddResult(o Result) {
	c.Results = append(c.Results, o)
	if len(c.Results) > MAX_RESULTS {
		c.Results = c.Results[len(c.Results)-MAX_RESULTS : len(c.Results)]
	}
}

func (c *Console) Delay() time.Duration {
	if c.DNI {
		return time.Second * DNISpeed
	} else {
		return time.Second
	}
}

func (c *Console) Connect(address string, force bool) error {
	server, err := c.FindServer(address)
	if err != nil {
		return err
	}

	// login de l'identité courante (si elle existe)
	login := ""
	if c.Identity != nil {
		login = c.Identity.Login
	}

	account := server.FindAccount(login)

	if !force && !server.Public && account == nil {
		return fmt.Errorf("%s : %w", login, errInvalidAccount)
	}

	c.Session = &Session{
		Server:  server,
		Account: account,
		Parent:  c.Session,
	}

	if c.Account != nil && c.Account.Backdoor {
		c.RemoveAccount(login)
		c.RemoveIdentity(login)
	}

	c.InitMem()

	return nil
}
