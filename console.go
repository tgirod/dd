package main

import (
//	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/mattn/go-shellwords"
)

const MAX_RESULTS int = 10
const COUNTDOWN time.Duration = time.Minute

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID int

	// racine de l'arbre des commandes
	Branch

	// informations sur la connexion au réseau
	*Session

	// interface neurale directe
	DNI bool

	// liste des dernières commandes évaluées
	Results []Result

	// trace en cours
	Alert bool
}

type Session struct {
	Server                  // serveur auquel la session seréfère
	User                    // compte utilisateur actif dans ce serveur
	Identity                // identité active dans la session
	Countdown time.Duration // temps restant avant déconnexion
	Mem       []MemoryZone  // zones mémoires disponibles pour une évasion
	Parent    *Session      // session précédente
}

type MemoryZone struct {
	Address string
	Used    bool
}

func (z MemoryZone) Value() any {
	return z.Address
}

func (z MemoryZone) Desc() string {
	if z.Used {
		return "occupée"
	} else {
		return "disponible"
	}
}

func (c *Console) WithSession(server Server, user User, identity Identity, reset bool) {
	countdown := COUNTDOWN
	if c.Alert && !reset {
		countdown = 0
	}

	parent := c.Session
	if reset {
		parent = nil
		c.Alert = false
	}

	c.Session = &Session{
		Server:    server,
		User:      user,
		Identity:  identity,
		Countdown: countdown,
		Mem:       InitMem(),
		Parent:    parent,
	}
}

// TimeLeft retourne le temps restant avant la finalisation de la trace
func (c *Console) TimeLeft() time.Duration {
	var left time.Duration
	sess := c.Session
	for sess != nil {
		left += sess.Countdown
		sess = sess.Parent
	}
	return left
}

// StartAlert démarre l'alerte si elle ne l'est pas déjà
func (c *Console) StartAlert() {
	if !c.Alert {
		c.Alert = true
	}
}

// Security décrémente le countdown dans la session
func (s *Session) Security() bool {
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

// FIXME codes hexa pour le chargement des commandes spéciales ?
var Hack = map[string]Cmd{
	"jack":  jack,
	"evade": evade,
	"door":  door,
	"imp":   imp,
	"scan":  scan,
}

var baseCmds = Branch{
	name: "root",
	cmds: []Cmd{
		back,
		yes,
		connect,
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

var MonCmds = Branch{
	name: "root",
	cmds: []Cmd{
		sudo_id,
		sudo_msg,
		sudo_yes,
		sudo_forum,
		sudo_reg,
		back,
		yes,
		connect,
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

func NewConsole(monitoring bool) *Console {
	app.Log("nouvelle console")
	if monitoring {
		return &Console{
			Branch:  MonCmds,
			Session: &Session{},
			DNI:     false,
			Results: []Result{},
			Alert:   false,
		}
	}
	return &Console{
		Branch:  baseCmds,
		Session: &Session{},
		DNI:     false,
		Results: []Result{},
		Alert:   false,
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

func InitMem() []MemoryZone {
	mz := make([]MemoryZone, 10)
	for i := 0; i < len(mz); i++ {
		addr := fmt.Sprintf("%04x", uint16(rand.Uint32()))
		mz[i] = MemoryZone{Address: addr}
	}
	return mz
}

func (c *Console) Disconnect() {
	c.Session = &Session{}
	c.Identity = Identity{}
	c.Alert = false
	c.DNI = false

	// remove Hackers Cmds from Branch, they should be at the end
	// plus propre que de raccourcir la liste et marche aussi pour
	// Monitor.
	id := len(c.Branch.cmds) - 1
	clean := false
	for id >= 0 && !clean {
		_, ok := Hack[c.Branch.cmds[id].name]
		if ok { // it is a Hack Cmd
			// remove last element
			c.Branch.cmds = c.Branch.cmds[:id]
			id = id - 1
		} else {
			clean = true
		}
	}

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
	o.Output = strings.ReplaceAll(o.Output, "\t", "    ")
	c.Results = append(c.Results, o)
	if len(c.Results) > MAX_RESULTS {
		c.Results = c.Results[len(c.Results)-MAX_RESULTS : len(c.Results)]
	}
}

func (c *Console) Delay() time.Duration {
	security := 1 << c.Session.Server.Security
	speed := time.Second / time.Duration(security)

	if c.DNI {
		return speed * DNISpeed
	} else {
		return speed
	}
}

// FIXME a virer
var shutdown = []string{
	dd.Address,
	frozdd.Address,
	maravdd.Address,
}

func (c *Console) Connect(address string, identity Identity, force bool, reset bool) error {
	// FIXME a virer
	/*
	ok := false
	for _, s := range shutdown {
		if s == address {
			ok = true
			break
		}
	}
	if !ok {
		return errors.New(`POUR DES RAISONS DE SÉCURITÉ, CETTE LIAISON RÉSEAU A ÉTÉ COUPÉE
POUR TOUTE DEMANDE DE RÉTABLISSEMENT,
MERCI DE CONTACTER LE SERVICE CLIENT DE LEGBA VOODOOCOM `)
	}
	*/
	/// fin

	server, err := FindServer(address)
	if err != nil {
		return err
	}

	// compte associé à l'identité active
	user, err := server.FindUser(identity.Login)

	// serveur privé, pas de compte, pas de connexion forcée
	if server.Private && err != nil && !force {
		return errInvalidUser
	}

	c.WithSession(server, user, identity, reset)

	if c.User.Backdoor {
		c.RemoveUser(user)
		RemoveIdentity(c.Identity)
	}

	return nil
}

func (c *Console) IsConnected() bool {
	return c.Session.Server.Address != ""
}
