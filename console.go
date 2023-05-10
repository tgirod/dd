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
	Session

	// l'alerte est-elle activée ?
	Alert bool

	// durée avant la déconnexion forcée
	Countdown time.Duration

	// interface neurale directe
	DNI bool

	// liste des dernières commandes évaluées
	Results []Result
}

type Session struct {
	Server                   // serveur auquel la session se réfère
	Account                  // compte utilisateur actif dans ce serveur
	Identity                 // identité active dans la session
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

func (s Session) WithSession(server Server, account Account, identity Identity) Session {
	sess := Session{
		Server:   server,
		Account:  account,
		Identity: identity,
		Parent:   &s,
	}
	sess.InitMem()
	return sess
}

type Result struct {
	Prompt string
	Error  error
	Output string
}

func (r Result) String() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "> %s\n", r.Prompt)
	if r.Error != nil {
		fmt.Fprintln(&b, r.Error.Error())
	}
	if r.Output != "" {
		fmt.Fprintln(&b, r.Output)
	}
	return b.String()
}

var Hack = map[string]Cmd{
	"jack":  jack,
	"evade": evade,
	"door":  door,
	"imp":   imp,
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

func NewConsole() *Console {
	app.Log("nouvelle console")
	return &Console{
		Branch: baseCmds,
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

func (s *Session) InitMem() {
	s.Mem = make(map[string]bool)
	for i := 0; i < 5; i++ {
		addr := fmt.Sprintf("%08x", rand.Uint32())
		s.Mem[addr] = true
	}
}

func (c *Console) Disconnect() {
	c.Session = Session{}
	c.Identity = Identity{}
	c.Alert = false
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

func (c *Console) TickAlert() {
	// décrémenter d'une seconde
	c.Countdown -= time.Second

	if c.Countdown <= 0 {
		c.Disconnect()
	}
}

func (c *Console) Identify(login, password string) error {
	identity, err := CheckIdentity(login, password)
	if err != nil {
		return err
	}
	c.Identity = identity

	// si on est connecté à un serveur, on tente d'accéder au compte utilisateur
	if c.IsConnected() {
		account, err := c.FindAccount(identity.Login)
		if err == nil {
			c.Account = account
		} else {
			c.Account = Account{}
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

func (c *Console) Connect(address string, identity Identity, force bool) error {
	server, err := FindServer(address)
	if err != nil {
		return err
	}

	// compte associé à l'identité active
	account, err := server.FindAccount(identity.Login)

	if server.Private && err != nil {
		if force {
			c.Session = c.Session.WithSession(server, Account{}, identity)
			return nil
		}

		return errInvalidAccount
	}

	c.Session = c.Session.WithSession(server, account, identity)

	if c.Account.Backdoor {
		c.RemoveAccount(account)
		RemoveIdentity(c.Identity)
	}

	return nil
}

func (c *Console) IsConnected() bool {
	return c.Session.Server.Address != ""
}
