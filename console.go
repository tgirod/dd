package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID int

	// racine de l'arbre des commandes
	Node

	// identité active sur la console
	Identity string

	// identifiant dans le serveur actuel
	Login string

	// admin dans le serveur actuel ?
	Admin bool

	// l'alerte est-elle activée ?
	Alert bool

	// durée avant la déconnexion forcée
	Countdown time.Duration

	// zones mémoires disponibles pour une évasion
	Mem map[string]bool

	// interface neurale directe
	DNI bool

	// Pile de serveurs visités lors de cette connexion
	History Stack

	// serveur auquel la console est connectée
	*Server

	// état interne du jeu
	*Game
}

var Hack = map[string]Command{
	"yyqz": Jack{},
	"zfcq": Evade{},
}

var baseCmds = []Command{
	Back{},
	Connect{},
	Data,
	Help{},
	Identify{},
	Index{},
	Link{},
	Load{},
	Plug{},
	Quit{},
	Registry,
	Pop{},
}

func NewConsole(game *Game) *Console {
	return &Console{
		Node: Node{
			Sub: baseCmds,
		},
		Game: game,
	}
}

func (c *Console) connect(address string) error {
	var err error
	var server *Server

	// récupérer le serveur
	if server, err = c.FindServer(address); err != nil {
		return fmt.Errorf("%s : %w", address, err)
	}

	// vérifier que l'utilisateur a le droit de se connecter
	if c.Admin, err = server.CheckAccount(c.Login); err != nil {
		return fmt.Errorf("%s : %w", c.Login, err)
	}

	// enregistrer le nouveau serveur
	c.Server = server
	c.InitMem()
	return nil
}

func (c *Console) Connect(address string) error {
	if err := c.connect(address); err != nil {
		return err
	}

	c.History.Clear()
	c.History.Push(Target{address, ""})
	fmt.Println(c.History)
	return nil
}

func (c *Console) Link(id int) error {
	if id < 0 || id >= len(c.Server.Targets) {
		return errInvalidArgument
	}

	target := c.Server.Targets[id]
	if err := c.connect(target.Address); err != nil {
		return err
	}

	c.History.Push(target)
	fmt.Println(c.History)
	return nil
}

func (c *Console) Back() error {
	if !c.IsConnected() {
		return errNotConnected
	}

	if len(c.History) == 1 {
		return errInvalidCommand
	}

	// enlever le serveur actuel
	c.History.Pop()

	prevTarget, _ := c.History.Peek()

	fmt.Println(c.History)
	return c.connect(prevTarget.Address)
}

func (c *Console) IsConnected() bool {
	return c.Server != nil
}

func (c *Console) InitMem() {
	c.Mem = make(map[string]bool)
	for i := 0; i < 5; i++ {
		addr := fmt.Sprintf("%08x", rand.Uint32())
		c.Mem[addr] = true
	}
}

func (c *Console) Disconnect() error {
	if !c.IsConnected() {
		return errNotConnected
	}

	c.Server = nil
	c.Login = ""
	c.Admin = false
	c.Alert = false
	c.History.Clear()
	c.Node.Sub = baseCmds
	return nil
}

func (c *Console) Load(code string) error {
	command, ok := Hack[code]
	if !ok {
		return fmt.Errorf("%s : %w", code, errInvalidArgument)
	}

	c.Node.Sub = append(c.Node.Sub, command)
	return nil
}

func (c *Console) Plug() error {
	if c.IsConnected() {
		return errConnected
	}

	c.DNI = true
	return nil
}

func (c *Console) Jack(id int) error {
	if !c.IsConnected() {
		return errNotConnected
	}

	if id < 0 || id >= len(c.Server.Targets) {
		return errInvalidArgument
	}

	target := c.Server.Targets[id]
	server, err := c.Game.FindServer(target.Address)
	if err != nil {
		return err
	}

	c.Server = server
	c.Login = "illegal"
	c.Admin = false
	c.InitMem()
	c.History.Push(target)
	return nil
}

func (c *Console) StartSecurity() {
	if !c.Alert {
		c.Alert = true
		c.Countdown = c.Server.Scan
	} else if c.Server.Scan < c.Countdown {
		c.Countdown = c.Server.Scan
	}
}

func (c *Console) DataSearch(keyword string) ([]Entry, error) {
	var search []Entry

	if !c.IsConnected() {
		return search, errNotConnected
	}

	if len([]rune(keyword)) < 3 {
		return search, fmt.Errorf("%s : %w", keyword, errKeywordTooShort)
	}

	search = c.Server.DataSearch(keyword, c.Login)
	return search, nil
}

func (c *Console) DataView(id string) (Entry, error) {
	var err error
	var entry Entry

	if !c.IsConnected() {
		return entry, errNotConnected
	}

	entry, err = c.Server.FindEntry(id, c.Login)
	if err != nil {
		return entry, err
	}

	return entry, nil
}

func (c *Console) Evade(zone string) error {
	if !c.IsConnected() {
		return errNotConnected
	}

	available, exist := c.Mem[zone]
	if !exist {
		return fmt.Errorf("%s : %w", zone, errMemNotFound)
	}

	if !available {
		return fmt.Errorf("%s : %w", zone, errMemUnavailable)
	}

	c.Mem[zone] = false
	c.Countdown = c.Server.Scan

	return nil
}

func (c *Console) RegistrySearch(name string) ([]Register, error) {
	var search []Register

	if !c.IsConnected() {
		return search, errNotConnected
	}

	search = c.Server.RegistrySearch(name)
	return search, nil
}

func (c *Console) RegistryEdit(name string) error {
	if !c.IsConnected() {
		return errNotConnected
	}

	// sauver l'état du jeu
	c.Game.Serialize()

	return c.Server.RegistryEdit(name)
}
