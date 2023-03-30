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
	Cmd

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

	// liste des dernières commandes évaluées
	Evals []Eval

	// serveur auquel la console est connectée
	*Server

	// état interne du jeu
	*Game
}

type Eval struct {
	Cmd    string
	Error  error
	Output string
}

var Hack = map[string]Cmd{
	"yyqz": jack,
	"zfcq": evade,
}

var baseCmds = Cmd{
	SubCmds: []Cmd{
		back,
		connect,
		data,
		help,
		identify,
		index,
		link,
		load,
		plug,
		quit,
		registry,
		pop,
	},
}

func NewConsole(game *Game) *Console {
	return &Console{
		Cmd:  baseCmds,
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
	eval := Eval{
		Cmd: fmt.Sprintf("connect %s", address),
	}

	if err := c.connect(address); err != nil {
		eval.Error = err
		c.AppendOutput(eval)
		return
	}

	c.History.Clear()
	c.History.Push(Target{address, ""})

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", c.Server.Address)
	fmt.Fprintf(&b, "%s\n", c.Server.Description)
	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) LinkList() {
	eval := Eval{
		Cmd: "link",
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ID\tDESCRIPTION\t\n")
	for i, t := range c.Server.Targets {
		fmt.Fprintf(tw, "%d\t%s\t\n", i, t.Description)
	}
	tw.Flush()

	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) Link(id int) {
	eval := Eval{
		Cmd: fmt.Sprintf("link %d", id),
	}

	if id < 0 || id >= len(c.Server.Targets) {
		eval.Error = errInvalidArgument
		c.AppendOutput(eval)
		return
	}

	target := c.Server.Targets[id]
	if err := c.connect(target.Address); err != nil {
		eval.Error = err
		c.AppendOutput(eval)
		return
	}

	c.History.Push(target)
	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", c.Server.Address)
	fmt.Fprintf(&b, "%s\n", c.Server.Description)
	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) Back() {
	e := Eval{
		Cmd: "back",
	}

	if !c.IsConnected() {
		e.Error = errNotConnected
		c.AppendOutput(e)
		return
	}

	if len(c.History) == 1 {
		e.Error = errInvalidCommand
		c.AppendOutput(e)
		return
	}

	// enlever le serveur actuel
	c.History.Pop()

	prevTarget, _ := c.History.Peek()

	if err := c.connect(prevTarget.Address); err != nil {
		e.Error = err
		c.AppendOutput(e)
		return
	}

	e.Output = fmt.Sprintf("connexion établie à l'adresse %s\n\n", c.Server.Address)
	c.AppendOutput(e)
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
	eval := Eval{
		Cmd: "quit",
	}
	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	c.Server = nil
	c.Login = ""
	c.Admin = false
	c.Alert = false
	c.History.Clear()
	c.Cmd = baseCmds

	eval.Output = "déconnexion effectuée"
	c.AppendOutput(eval)
}

func (c *Console) Disconnect() {
	c.Server = nil
	c.Login = ""
	c.Admin = false
	c.Alert = false
	c.History.Clear()
	c.Cmd = baseCmds

	// affichage par défaut
	eval := Eval{
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

	c.AppendOutput(eval)
}

func (c *Console) Load(code string) {
	eval := Eval{
		Cmd: fmt.Sprintf("load %s", code),
	}

	command, ok := Hack[code]
	if !ok {
		eval.Error = fmt.Errorf("%s : %w", code, errInvalidArgument)
		c.AppendOutput(eval)
		return
	}

	c.Cmd.SubCmds = append(c.Cmd.SubCmds, command)
	eval.Output = fmt.Sprintf("%s : commande chargée", command.Name)
	c.AppendOutput(eval)
}

func (c *Console) Plug() {
	eval := Eval{
		Cmd: "plug",
	}

	if c.IsConnected() {
		eval.Error = errConnected
		c.AppendOutput(eval)
		return
	}

	c.DNI = true
	eval.Output = "interface neuronale directe activée"
	c.AppendOutput(eval)
}

func (c *Console) Jack(id int) {
	eval := Eval{
		Cmd: fmt.Sprintf("jack %d", id),
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	if id < 0 || id >= len(c.Server.Targets) {
		eval.Error = errInvalidArgument
		c.AppendOutput(eval)
		return
	}

	target := c.Server.Targets[id]
	server, err := c.Game.FindServer(target.Address)
	if err != nil {
		eval.Error = err
		c.AppendOutput(eval)
		return
	}

	c.Server = server
	c.Login = "illegal"
	c.Admin = false
	c.InitMem()
	c.History.Push(target)

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", c.Server.Address)
	fmt.Fprintf(&b, "%s\n", c.Server.Description)
	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) StartAlert() {
	if !c.Alert {
		c.Alert = true
		c.Countdown = c.Server.Scan
	} else if c.Server.Scan < c.Countdown {
		c.Countdown = c.Server.Scan
	}
}

func (c *Console) DataSearch(keyword string) {
	eval := Eval{
		Cmd: fmt.Sprintf("data search %s", keyword),
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	if len([]rune(keyword)) < 3 {
		eval.Error = fmt.Errorf("%s : %w", keyword, errKeywordTooShort)
		c.AppendOutput(eval)
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

	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) DataView(id string) {
	eval := Eval{
		Cmd: fmt.Sprintf("data view %s", id),
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	entry, err := c.Server.FindEntry(id, c.Login)
	if err != nil {
		eval.Error = err
		c.AppendOutput(eval)
		return
	}

	// construire la réponse à afficher
	b := strings.Builder{}
	fmt.Fprintf(&b, "TITLE: %s\n", entry.Title)
	fmt.Fprintf(&b, "KEYWORDS: %s\n", strings.Join(entry.Keywords, " "))
	fmt.Fprintf(&b, "-------------------------------------\n")
	fmt.Fprintf(&b, entry.Content)

	eval.Output = b.String()
	c.AppendOutput(eval)

}

func (c *Console) Help(args []string) {
	eval := Eval{
		Cmd: fmt.Sprintf("help %s", strings.Join(args, " ")),
	}
	eval.Output = c.Cmd.Help(args)
	c.AppendOutput(eval)
}

func (c *Console) EvadeList() {
	eval := Eval{
		Cmd: "evade",
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ZONE\tDISPONIBILITE\t\n")
	for addr, available := range c.Mem {
		if !available {
			fmt.Fprintf(tw, "%s\t%s\t\n", addr, "INDISPONIBLE")
		} else {
			fmt.Fprintf(tw, "%s\t%s\t\n", addr, "OK")
		}
	}
	tw.Flush()

	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) Evade(zone string) {
	eval := Eval{
		Cmd: fmt.Sprintf("evade %s", zone),
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	available, ok := c.Mem[zone]
	if !ok {
		eval.Error = fmt.Errorf("%s : %w", zone, errMemNotFound)
		c.AppendOutput(eval)
		return
	}

	if !available {
		eval.Error = fmt.Errorf("%s : %w", zone, errMemUnavailable)
		c.AppendOutput(eval)
		return
	}

	c.Mem[zone] = false
	c.Countdown = c.Server.Scan
	eval.Output = fmt.Sprintf("session relocalisée dans la zone mémoire %s", zone)
	c.AppendOutput(eval)
}

func (c *Console) RegistrySearch(name string) {
	eval := Eval{
		Cmd: fmt.Sprintf("registry search %s", name),
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	search := c.Server.RegistrySearch(name)

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "NAME\tSTATE\tDESCRIPTION\t\n")
	for _, r := range search {
		fmt.Fprintf(tw, "%s\t%t\t%s\t\n", r.Name, r.State, r.Description)
	}
	tw.Flush()

	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) RegistryEdit(name string) {
	eval := Eval{
		Cmd: fmt.Sprintf("registry edit %s", name),
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	state, err := c.Server.RegistryEdit(name)

	if err != nil {
		eval.Error = err
		c.AppendOutput(eval)
		return
	}

	eval.Output = fmt.Sprintf("nouvel état du registre %s : %v\n", name, state)
	c.AppendOutput(eval)
}

func (c *Console) Identify(login, password string) {
	eval := Eval{
		Cmd: fmt.Sprintf("identify %s %s", login, password),
	}

	if err := c.CheckIdentity(login, password); err != nil {
		eval.Error = err
		c.AppendOutput(eval)
		return
	}

	c.Login = login

	// si on est connecté à un serveur, on tente d'accéder au compte utilisateur
	if c.Server != nil {
		if admin, err := c.CheckAccount(login); err == nil {
			c.Admin = admin
		}
	}

	eval.Output = fmt.Sprintf("Identité établie. Bienvenue, %s\n", login)
	c.AppendOutput(eval)
}

func (c *Console) Index() {
	eval := Eval{
		Cmd: "index",
	}

	if !c.IsConnected() {
		eval.Error = errNotConnected
		c.AppendOutput(eval)
		return
	}

	b := strings.Builder{}

	s := c.Server
	b.WriteString(s.Description)
	b.WriteString("\n")
	fmt.Fprintf(&b, "LIENS     : %d\n", len(s.Targets))
	fmt.Fprintf(&b, "DONNEES   : %d\n", len(s.Entries))
	fmt.Fprintf(&b, "REGISTRES : %d\n", len(s.Registers))

	eval.Output = b.String()
	c.AppendOutput(eval)
}

func (c *Console) AppendOutput(o Eval) {
	c.Evals = append(c.Evals, o)
	if len(c.Evals) > MAX_LEN_OUTPUT {
		c.Evals = c.Evals[len(c.Evals)-MAX_LEN_OUTPUT : len(c.Evals)]
	}
}
