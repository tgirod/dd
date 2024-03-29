package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	errInvalidCommand  = errors.New("commande invalide")
	errMissingCommand  = errors.New("commande manquante")
	errMissingArgument = errors.New("argument manquant")
	errInvalidArgument = errors.New("argument invalide")
)

// Context une étape du parsing. A chaque nouvelle étape, un nouveau contexte est créé, encapsulant le précédent.
type Context struct {
	parent *Context
	key    string
	value  any
	node   Node
}

func (c Context) Help() string {
	if c.node == nil {
		return ""
	}

	switch node := c.node.(type) {
	case Cmd:
		return node.Help()
	case String:
		return fmt.Sprintf("%s : %v", node.help, c.value)
	case Number:
		return fmt.Sprintf("%s : %v", node.help, c.value)
	case Select:
		return fmt.Sprintf("%s : %v", node.help, c.value)
	case Hidden:
		return node.Help()
	case Text:
		return node.Help()
	case LongText:
		return node.Help()
	default:
		return "NO HELP"
	}
}

// WithContext retourne un nouveau contexte ajoutant une étape de parsing
func (c Context) WithContext(node Node, key string, value any) Context {
	// BUG CONNECT
	// fmt.Printf("LOG WithContext key %s, new key %s arg %v\n", c.key, key, value)
	return Context{
		parent: &c,
		key:    key,
		value:  value,
		node:   node,
	}
}

// Value accède à une valeur stockée dans le contexte
func (c Context) Value(key any) any {
	if c.key == key {
		return c.value
	}

	if c.parent != nil {
		return c.parent.Value(key)
	}

	return nil
}

func (c Context) Console() *Console {
	return c.Value("console").(*Console)
}

func (c Context) Print() string {
	b := strings.Builder{}
	ctx := &c
	for ctx != nil {
		fmt.Fprintf(&b, "%s : %v\n", ctx.key, ctx.value)
		ctx = ctx.parent
	}
	return b.String()
}

// Prompt reconstruit la commande entrée à l'origine
func (c Context) Prompt() string {
	if c.key == "console" {
		return ""
	}

	if _, ok := c.node.(Hidden); ok {
		// need to convert to string to get length
		val := fmt.Sprintf("%v", c.value)
		val = strings.Repeat("*", len(val))
		return fmt.Sprintf("%v %s", c.parent.Prompt(), val)
	}
	return fmt.Sprintf("%v %v", c.parent.Prompt(), c.value)
}

// Result créé un objet Result de base par défaut
func (c Context) Result(err error, output string) Result {
	return Result{
		Prompt: c.Prompt(),
		Error:  err,
		Output: output,
	}
}

func (c Context) Output(output string) Result {
	return Result{
		Prompt: c.Prompt(),
		Output: output,
	}
}

func (c Context) Error(err error) Result {
	return Result{
		Prompt: c.Prompt(),
		Error:  err,
	}
}

func (c Context) Resume(args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG ctx.Resume key %s args %v\n", c.key, args)
	// fmt.Printf("LOG c.node is %v\n", c.Help())
	return c.node.Resume(c, args)
	//return c.node.Parse(c, args)
}

// Node est un noeud dans l'arbre de parsing
type Node interface {
	String() string                        // nom du noeud
	Help() string                          // aide courte sur une ligne
	Parse(ctx Context, args []string) any  // parsing
	Resume(ctx Context, args []string) any // reprendre le parsing
}

// Branch permet à l'utilisateur de choisir parmi plusieurs commandes en fonction de leur nom
type Branch struct {
	name string
	cmds []Cmd
}

type Cmd struct {
	name       string
	help       string
	connected  bool
	identified bool
	next       Node
}

func (c Cmd) String() string {
	return c.name
}

func (c Cmd) Help() string {
	return fmt.Sprintf("%s : %s", c.name, c.help)
}

func (c Cmd) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Cmd.Parse name %s key %s arg %v\n", c.name, ctx.key, args)
	return c.next.Parse(ctx, args)
}

func (c Cmd) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Cmd.Resume name %s key %s args %v\n", c.name, ctx.key, args)
	return c.next.Parse(ctx, args)
}

func (b Branch) String() string {
	return b.name
}

func (b Branch) Help() string {
	s := strings.Builder{}
	tw := tw(&s)
	fmt.Fprintln(tw, underline.Render("liste des commandes disponibles"))
	for _, c := range b.cmds {
		fmt.Fprintf(tw, "%s\t%s\t\n", c.name, c.help)
	}
	tw.Flush()
	return s.String()
}

func (b Branch) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Branch.Parse name %s key %s arg %v\n", b.name, ctx.key, args)
	// aucun argument à parser, l'exécution s'arrête là
	if len(args) == 0 {
		return ctx.Result(
			fmt.Errorf("%s : %w", b.name, errMissingCommand),
			ctx.Help()+"\n\n"+b.Help(),
		)
	}

	// un argument à parser, on sélectionne le prochain noeud
	for _, cmd := range b.cmds {
		if strings.HasPrefix(cmd.name, args[0]) {
			// HACK vérifier l'identité et la connectivité
			if cmd.connected || cmd.identified {
				console := ctx.Console()
				if cmd.connected && !console.IsConnected() {
					return ctx.Error(errNotConnected)
				}

				if cmd.identified && console.Identity.Login == "" {
					return ctx.Error(errNotIdentified)
				}
			}

			// on enregistre la commande correspondante dans le contexte et on poursuit
			ctx = ctx.WithContext(cmd, "", cmd.name)

			return cmd.next.Parse(ctx, args[1:])
		}
	}
	// aucune commande ne correspond
	return ctx.Result(
		fmt.Errorf("%s : %w", b.name, errInvalidCommand),
		ctx.Help()+"\n\n"+b.Help(),
	)
}

func (b Branch) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Branch.Resume name %s key %s args %v\n", b.name, ctx.key, args)
	// on vérifie si le noeud est déjà parsé (resume)
	if cmd, ok := ctx.Value(b.name).(Cmd); ok {
		return cmd.next.Parse(ctx, args)
	}

	// FIXME ca ne devrait jamais arriver
	panic(errors.New("reprise d'exécution impossible"))
}

// Run est le noeud terminal d'une commande, là où se trouve le code qui altère l'état du jeu
type Run func(ctx Context) any

func (r Run) String() string {
	return ""
}

func (r Run) Help() string {
	return ""
}

func (r Run) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Run.Parse key %s arg %v\n", ctx.key, args)
	return r(ctx)
}

func (r Run) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Run.Resume key %s args %v\n", ctx.key, args)
	return nil
}

// String est un argument chaine de caractère libre
type String struct {
	name string
	help string
	next Node
}

func (s String) String() string {
	return s.name
}

func (s String) Help() string {
	return fmt.Sprintf("%s : %s", s.name, s.help)
}

func (s String) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG String.Parse name %s key %s arg %v\n", s.name, ctx.key, args)
	if len(args) == 0 {
		return ctx.Result(
			fmt.Errorf("%s : %w", s.name, errMissingArgument),
			ctx.Help()+"\n\n"+s.Help(),
		)
	}

	ctx = ctx.WithContext(s, s.name, args[0])
	return s.next.Parse(ctx, args[1:])
}

func (s String) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG String.Resume name %s key %s arg %v\n", s.name, ctx.key, args)
	return s.next.Parse(ctx, args)
}

// Number est un argument chaine de caractère libre
type Number struct {
	name string
	help string
	next Node
}

func (n Number) String() string {
	return n.name
}

func (n Number) Help() string {
	return fmt.Sprintf("%s : %s", n.name, n.help)
}

func (n Number) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Number.Parse name %s key %s arg %v\n", n.name, ctx.key, args)
	// on vérifie si le noeud est déjà parsé (resume)
	if _, ok := ctx.Value(n.name).(int); ok {
		return n.next.Parse(ctx, args)
	}

	if len(args) == 0 {
		return ctx.Result(
			fmt.Errorf("%s : %w", n.name, errMissingArgument),
			ctx.Help()+"\n\n"+n.Help(),
		)
	}

	num, err := strconv.Atoi(args[0])
	if err != nil {
		return ctx.Result(
			fmt.Errorf("%s : %w", n.name, errInvalidArgument),
			n.Help(),
		)
	}

	ctx = ctx.WithContext(n, n.name, num)
	return n.next.Parse(ctx, args[1:])
}

func (n Number) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Number.Resume name %s key %s arg %v\n", n.name, ctx.key, args)
	return n.next.Parse(ctx, args)
}

type Select struct {
	name    string
	help    string
	header  string
	options func(ctx Context) ([]Option, error)
	next    Node
}

// Option représente un choix possible dans une sélection
type Option interface {
	Value() any
	Desc() string
}

func (s Select) String() string {
	return s.name
}

func (s Select) Help() string {
	return fmt.Sprintf("%s : %s", s.name, s.help)
}

var underline = lipgloss.NewStyle().Underline(true)

func (s Select) List(options []Option) string {
	// construire le header
	b := strings.Builder{}
	header := s.header
	if header == "" {
		header = s.help
	}
	fmt.Fprintln(&b, underline.Render(header))

	tw := tw(&b)
	// afficher les options
	for _, o := range options {
		fmt.Fprintf(tw, "%v\t%s\t\n", o.Value(), o.Desc())
	}
	tw.Flush()
	return b.String()
}

func (s Select) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Select.Parse name %s key %s arg %v\n", s.name, ctx.key, args)
	// récupérer la liste des options possibles
	options, err := s.options(ctx)
	if err != nil {
		return ctx.Error(err)
	}

	if len(args) == 0 {
		// afficher la liste des choix possibles
		return ctx.Output(ctx.Help() + "\n\n" + s.List(options))
	}

	// vérifier la validité de l'option choisie
	for _, o := range options {
		if fmt.Sprintf("%v", o.Value()) == args[0] {
			// la valeur est valide, continuer le parsing
			ctx = ctx.WithContext(s, s.name, o.Value())
			return s.next.Parse(ctx, args[1:])
		}
	}

	// la valeur saisie est invalide
	return ctx.Error(
		fmt.Errorf("%s : %w", s.name, errInvalidArgument),
	)
}

func (s Select) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Select.Resume name %s key %s arg %v\n", s.name, ctx.key, args)
	return s.next.Parse(ctx, args)
}

type Hidden struct {
	name string
	help string
	next Node
}

func (h Hidden) String() string {
	return h.name
}

func (h Hidden) Help() string {
	return fmt.Sprintf("%s : %s", h.name, h.help)
}

func (h Hidden) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Hidden.Parse name %s key %s arg %v\n", h.name, ctx.key, args)
	if len(args) == 0 {
		// ouvrir une fenêtre modale
		modal := NewShort(ctx, h, true)
		return OpenModalMsg(modal)
	}

	ctx = ctx.WithContext(h, h.name, args[0])
	return h.next.Parse(ctx, args[1:])
}

func (h Hidden) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Hidden.Resume name %s key %s arg %v\n", h.name, ctx.key, args)
	return h.next.Parse(ctx, args)
}

type Text struct {
	name string
	help string
	next Node
}

func (t Text) String() string {
	return t.name
}

func (t Text) Help() string {
	return fmt.Sprintf("%s : %s", t.name, t.help)
}

func (t Text) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Text.Parse name %s key %s arg %v\n", t.name, ctx.key, args)
	if len(args) == 0 {
		// ouvrir une fenêtre modale
		modal := NewShort(ctx, t, false)
		return OpenModalMsg(modal)
	}

	// BUG CONNECT
	// fmt.Printf("LOG Text.Parse after modal name %s key %s arg %v\n", t.name, ctx.key, args)
	ctx = ctx.WithContext(t, t.name, args[0])
	return t.next.Parse(ctx, args[1:])
}

func (t Text) Resume(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG Text.Resume name %s key %s arg %v\n", t.name, ctx.key, args)
	return t.next.Parse(ctx, args)
}

type LongText struct {
	name string
	help string
	next Node
}

func (t LongText) String() string {
	return t.name
}

func (t LongText) Help() string {
	return fmt.Sprintf("%s : %s", t.name, t.help)
}

func (t LongText) Parse(ctx Context, args []string) any {
	// BUG CONNECT
	// fmt.Printf("LOG LongText.Parse name %s key %s arg %v\n", t.name, ctx.key, args)
	if len(args) == 0 {
		// ouvrir une fenêtre modale
		modal := NewLong(ctx, t)
		return OpenModalMsg(modal)
	}

	ctx = ctx.WithContext(t, t.name, args[0])
	return t.next.Parse(ctx, args[1:])
}

func (t LongText) Resume(ctx Context, args []string) any {
	return t.next.Parse(ctx, args)
}

func ToOptions[T Option](values []T) []Option {
	opts := make([]Option, len(values))
	for i, v := range values {
		opts[i] = v
	}
	return opts
}
