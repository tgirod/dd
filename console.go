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

	// interface neurale directe
	DNI bool

	// liste des dernières commandes évaluées
	Results []Result
}

type Session struct {
	Server                    // serveur auquel la session se réfère
	User                      // compte utilisateur actif dans ce serveur
	Identity                  // identité active dans la session
	Alert     bool            // l'alerte est-elle active ?
	Countdown time.Duration   // temps restant avant déconnexion
	Mem       map[string]bool // zones mémoires disponibles pour une évasion
	Parent    *Session        // session précédente
}

func (s Session) WithSession(server Server, user User, identity Identity) Session {
	countdown := server.Security
	if s.Alert {
		countdown = 0
	}
	sess := Session{
		Server:    server,
		User:      user,
		Identity:  identity,
		Alert:     s.Alert,
		Countdown: countdown,
		Mem:       InitMem(),
		Parent:    &s,
	}
	return sess
}

// Trace retourne le temps restant avant que la trace soit terminée
func (s Session) Trace() time.Duration {
	if s.Parent == nil {
		return s.Countdown
	}
	return s.Parent.Trace() + s.Countdown
}

func (s *Session) StartAlert() {
	s.Alert = true
}

func (s *Session) Security() bool {
	// décrémenter le temps restant dans cette session
	s.Alert = true
	if s.Countdown > 0 {
		s.Countdown -= time.Second
		return false
	}

	// décrémenter le temps dans les sessions antérieures
	if s.Parent != nil {
		return s.Parent.Security()
	}

	// on est au bout, la trace est complétée
	return true
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
		group,
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

func InitMem() map[string]bool {
	mem := make(map[string]bool)
	for i := 0; i < 5; i++ {
		addr := fmt.Sprintf("%08x", rand.Uint32())
		mem[addr] = true
	}
	return mem
}

func (c *Console) Disconnect() {
	c.Session = Session{}
	c.Identity = Identity{}
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

func (c *Console) Security() bool {
	if c.Alert {
		if c.Session.Security() {
			// trace complétée, on déconnecte
			c.Disconnect()
			return true
		}
	}
	return false
}

func (c *Console) Identify(login, password string) error {
	identity, err := CheckIdentity(login, password)
	if err != nil {
		return err
	}
	c.Identity = identity

	// si on est connecté à un serveur, on tente d'accéder au compte utilisateur
	if c.IsConnected() {
		user, err := c.FindUser(identity.Login)
		if err == nil {
			c.User = user
		} else {
			c.User = User{}
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
	user, err := server.FindUser(identity.Login)

	if server.Private && err != nil {
		if force {
			c.Session = c.Session.WithSession(server, User{}, identity)
			return nil
		}

		return errInvalidUser
	}

	c.Session = c.Session.WithSession(server, user, identity)

	if c.User.Backdoor {
		c.RemoveUser(user)
		RemoveIdentity(c.Identity)
	}

	return nil
}

func (c *Console) IsConnected() bool {
	return c.Session.Server.Address != ""
}
