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
