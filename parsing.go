package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

// WithContext retourne un nouveau contexte ajoutant une étape de parsing
func (c Context) WithContext(node Node, key string, value any) Context {
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

	return fmt.Sprintf("%s %s", c.parent.Prompt(), c.value)
}

// Result créé un objet Result de base par défaut
func (c Context) Result(err error, output string) Result {
	return Result{
		Prompt: c.Prompt(),
		Error:  err,
		Output: output,
	}
}

func (c Context) Resume(args []string) any {
	return c.node.Resume(c, args)
}

func (c Context) Cancel() *Context {
	return c.parent
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

func (b Branch) String() string {
	return b.name
}

func (b Branch) Help() string {
	s := strings.Builder{}
	for _, c := range b.cmds {
		fmt.Fprintf(&s, "%s : %s\n", c.name, c.help)
	}
	return s.String()
}

func (b Branch) Parse(ctx Context, args []string) any {
	// aucun argument à parser, l'exécution s'arrête là
	if len(args) == 0 {
		return ctx.Result(
			fmt.Errorf("%s : %w", b.name, errMissingCommand),
			b.Help(),
		)
	}

	// un argument à parser, on sélectionne le prochain noeud
	for _, cmd := range b.cmds {
		if strings.HasPrefix(cmd.name, args[0]) {
			// une commande correspond, on enregistre dans le contexte et on continue
			ctx = ctx.WithContext(b, b.name, cmd)
			return cmd.next.Parse(ctx, args[1:])
		}
	}

	// aucune commande ne correspond
	return ctx.Result(
		fmt.Errorf("%s : %w", b.name, errInvalidCommand),
		b.Help(),
	)
}

func (b Branch) Resume(ctx Context, args []string) any {
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
	return r(ctx)
}

func (r Run) Resume(ctx Context, args []string) any {
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
	return s.help
}

func (s String) Parse(ctx Context, args []string) any {
	if len(args) == 0 {
		return ctx.Result(
			fmt.Errorf("%s : %w", s.name, errMissingArgument),
			s.Help(),
		)
	}

	ctx = ctx.WithContext(s, s.name, args[0])
	return s.next.Parse(ctx, args[1:])
}

func (s String) Resume(ctx Context, args []string) any {
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
	return n.help
}

func (n Number) Parse(ctx Context, args []string) any {
	// on vérifie si le noeud est déjà parsé (resume)
	if _, ok := ctx.Value(n.name).(int); ok {
		return n.next.Parse(ctx, args)
	}

	if len(args) == 0 {
		return ctx.Result(
			fmt.Errorf("%s : %w", n.name, errMissingArgument),
			n.Help(),
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
	return n.next.Parse(ctx, args)
}

type Select struct {
	name    string
	help    string
	options func(ctx Context) []Option
	next    Node
}

type Option struct {
	value any
	desc  string
}

func (s Select) String() string {
	return s.name
}

func (s Select) Help() string {
	return s.help
}

func (s Select) List(ctx Context) string {
	b := strings.Builder{}
	for _, o := range s.options(ctx) {
		fmt.Fprintf(&b, "%v : %s\n", o.value, o.desc)
	}
	return b.String()
}

func (s Select) Parse(ctx Context, args []string) any {
	if len(args) == 0 {
		// afficher la liste des choix possibles

		return ctx.Result(
			fmt.Errorf("%s : %w", s.name, errMissingArgument),
			s.List(ctx),
		)
	}

	// vérifier que la valeur est valide
	for _, o := range s.options(ctx) {
		if fmt.Sprintf("%v", o.value) == args[0] {
			// la valeur est valide, continuer le parsing
			ctx = ctx.WithContext(s, s.name, o.value)
			return s.next.Parse(ctx, args[1:])
		}
	}

	// la valeur saisie est invalide
	return ctx.Result(
		fmt.Errorf("%s : %w", s.name, errInvalidArgument),
		s.List(ctx),
	)
}

func (s Select) Resume(ctx Context, args []string) any {
	return s.next.Parse(ctx, args)
}