package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const MAX_LEN_OUTPUT int = 10

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

	Output []Output

	// serveur auquel la console est connectée
	*Server

	// état interne du jeu
	*Game
}

type Output struct {
	Cmd     string
	Error   error
	Content string
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

func (c *Console) Connect(address string) {
	output := Output{
		Cmd: fmt.Sprintf("connect %s", address),
	}

	if err := c.connect(address); err != nil {
		output.Error = err
		c.AppendOutput(output)
		return
	}

	c.History.Clear()
	c.History.Push(Target{address, ""})

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", c.Server.Address)
	fmt.Fprintf(&b, "%s\n", c.Server.Description)
	output.Content = b.String()
	c.AppendOutput(output)
}

func (c *Console) LinkList() {
	output := Output{
		Cmd: "link",
	}

	if !c.IsConnected() {
		output.Error = errNotConnected
		c.AppendOutput(output)
		return
	}

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ID\tDESCRIPTION\t\n")
	for i, t := range c.Server.Targets {
		fmt.Fprintf(tw, "%d\t%s\t\n", i, t.Description)
	}
	tw.Flush()

	output.Content = b.String()
	c.AppendOutput(output)
}

func (c *Console) Link(id int) {
	output := Output{
		Cmd: fmt.Sprintf("link %d", id),
	}

	if id < 0 || id >= len(c.Server.Targets) {
		output.Error = errInvalidArgument
		c.AppendOutput(output)
		return
	}

	target := c.Server.Targets[id]
	if err := c.connect(target.Address); err != nil {
		output.Error = err
		c.AppendOutput(output)
		return
	}

	c.History.Push(target)
	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", c.Server.Address)
	fmt.Fprintf(&b, "%s\n", c.Server.Description)
	output.Content = b.String()
	c.AppendOutput(output)
}

func (c *Console) Back() {
	output := Output{
		Cmd: "back",
	}

	if !c.IsConnected() {
		output.Error = errNotConnected
		c.AppendOutput(output)
		return
	}

	if len(c.History) == 1 {
		output.Error = errInvalidCommand
		c.AppendOutput(output)
		return
	}

	// enlever le serveur actuel
	c.History.Pop()

	prevTarget, _ := c.History.Peek()

	if err := c.connect(prevTarget.Address); err != nil {
		output.Error = err
		c.AppendOutput(output)
		return
	}

	output.Content = fmt.Sprintf("connexion établie à l'adresse %s\n\n", c.Server.Address)
	c.AppendOutput(output)
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

func (c *Console) Quit() {
	output := Output{
		Cmd: "quit",
	}
	if !c.IsConnected() {
		output.Error = errNotConnected
		c.AppendOutput(output)
		return
	}

	c.Server = nil
	c.Login = ""
	c.Admin = false
	c.Alert = false
	c.History.Clear()
	c.Node.Sub = baseCmds

	output.Content = "déconnexion effectuée"
	c.AppendOutput(output)
}

func (c *Console) Load(code string) {
	output := Output{
		Cmd: fmt.Sprintf("load %s", code),
	}

	command, ok := Hack[code]
	if !ok {
		output.Error = fmt.Errorf("%s : %w", code, errInvalidArgument)
		c.AppendOutput(output)
		return
	}

	c.Node.Sub = append(c.Node.Sub, command)
	output.Content = fmt.Sprintf("%s : commande chargée", command.ParseName())
	c.AppendOutput(output)
}

func (c *Console) Plug() {
	output := Output{
		Cmd: "plug",
	}

	if c.IsConnected() {
		output.Error = errConnected
		c.AppendOutput(output)
		return
	}

	c.DNI = true
	output.Content = "interface neuronale directe activée"
	c.AppendOutput(output)
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

func (c *Console) DataSearch(keyword string) {
	var output = Output{
		Cmd: fmt.Sprintf("data search %s", keyword),
	}

	if !c.IsConnected() {
		output.Error = errNotConnected
		c.AppendOutput(output)
		return
	}

	if len([]rune(keyword)) < 3 {
		output.Error = fmt.Errorf("%s : %w", keyword, errKeywordTooShort)
		c.AppendOutput(output)
		return
	}

	// construire la réponse à afficher
	entries := c.Server.DataSearch(keyword, c.Login)
	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ID\tKEYWORDS\tTITLE\t\n")
	for _, e := range entries {
		title := e.Title
		fmt.Fprintf(tw, "%s\t%s\t%s\t\n",
			e.ID,
			strings.Join(e.Keywords, " "),
			title,
		)
	}
	tw.Flush()

	output.Content = b.String()
	c.AppendOutput(output)
}

func (c *Console) DataView(id string) {
	output := Output{
		Cmd: fmt.Sprintf("data view %s", id),
	}

	if !c.IsConnected() {
		output.Error = errNotConnected
		c.AppendOutput(output)
		return
	}

	entry, err := c.Server.FindEntry(id, c.Login)
	if err != nil {
		output.Error = err
		c.AppendOutput(output)
		return
	}

	// construire la réponse à afficher
	b := strings.Builder{}
	fmt.Fprintf(&b, "TITLE: %s\n", entry.Title)
	fmt.Fprintf(&b, "KEYWORDS: %s\n", strings.Join(entry.Keywords, " "))
	fmt.Fprintf(&b, "-------------------------------------\n")
	fmt.Fprintf(&b, entry.Content)

	output.Content = b.String()
	c.AppendOutput(output)

}

func (c *Console) Help(args []string) {
	output := Output{
		Cmd: fmt.Sprintf("help %s", strings.Join(args, " ")),
	}

	if len(args) == 0 {
		b := strings.Builder{}
		b.WriteString("COMMANDES DISPONIBLES\n\n")
		tw := tw(&b)
		fmt.Fprintf(tw, "NOM\tDESCRIPTION\t\n")
		for _, s := range c.Node.Sub {
			fmt.Fprintf(tw, "%s\t%s\t\n", s.ParseName(), s.ShortHelp())
		}
		tw.Flush()
		b.WriteString("\nPour plus d'aide, tapez 'help <COMMAND>'\n")

		output.Content = b.String()
		c.AppendOutput(output)
		return
	}

	// FIXME match récursif pour afficher l'aide d'une sous-commande
	match := c.Node.Match(args[0])
	if len(match) == 0 {
		output.Error = fmt.Errorf("%s : %w", args[0], errInvalidCommand)
		c.AppendOutput(output)
		return
	}

	output.Content = match[0].LongHelp()
	c.AppendOutput(output)
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

func (c *Console) Identify(login, password string) {
	output := Output{
		Cmd: fmt.Sprintf("identify %s %s", login, password),
	}

	if err := c.CheckIdentity(login, password); err != nil {
		output.Error = err
		c.AppendOutput(output)
		return
	}

	c.Login = login

	// si on est connecté à un serveur, on tente d'accéder au compte utilisateur
	if c.Server != nil {
		if admin, err := c.CheckAccount(login); err == nil {
			c.Admin = admin
		}
	}

	output.Content = fmt.Sprintf("Identité établie. Bienvenue, %s\n", login)
	c.AppendOutput(output)
}

func (c *Console) Index() {
	output := Output{
		Cmd: "index",
	}

	if !c.IsConnected() {
		output.Error = errNotConnected
		c.AppendOutput(output)
		return
	}

	b := strings.Builder{}

	s := c.Server
	b.WriteString(s.Description)
	b.WriteString("\n")
	fmt.Fprintf(&b, "LIENS     : %d\n", len(s.Targets))
	fmt.Fprintf(&b, "DONNEES   : %d\n", len(s.Entries))
	fmt.Fprintf(&b, "REGISTRES : %d\n", len(s.Registers))

	output.Content = b.String()
	c.AppendOutput(output)
}

func (c *Console) AppendOutput(o Output) {
	c.Output = append(c.Output, o)
	if len(c.Output) > MAX_LEN_OUTPUT {
		c.Output = c.Output[len(c.Output)-MAX_LEN_OUTPUT : len(c.Output)]
	}
}
