package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var ForumCmd = Node{
	Name: "forum",
	Help: "interagir avec un forum",
	Sub: []Command{
		ForumEnter{},
		ForumRead{},
		ForumLeave{},
		ForumShow{},
		ForumWritePost{},
		ForumAddTopic{},
	},
}
// By default, in

// *****************************************************************************
// ****************************************************************** ForumEnter
// *****************************************************************************
type ForumEnter struct{}

func (f ForumEnter) ParseName() string {
	return "enter"
}

func (f ForumEnter) ShortHelp() string {
	return "entre dans un Forum (s'il existe)"
}

func (f ForumEnter) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum enter \n")
	return b.String()
}

func (f ForumEnter) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum enter",
			Error: errNotConnected}
	}

	forum, err := c.Server.GetForum()
	c.Console.Forum = forum
	if err != nil {
		return ResultMsg{
			Cmd:   "forum enter " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : you are authorized to enter %s\n",
		c.Server.Address), 0)
}

// *****************************************************************************
// ******************************************************************* ForumRead
// *****************************************************************************
type ForumRead struct{}

func (f ForumRead) ParseName() string {
	return "read"
}

func (f ForumRead) ShortHelp() string {
	return "lit un topic ou un post d'un forum"
}

func (f ForumRead) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum read <ID>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  ID -- l'identifiant du topic/post à lire\n")
	return b.String()
}

func (f ForumRead) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum read " + strings.Join(args, " "),
			Error: errNotConnected}
	}

	if len(args) < 1 {
		return ResultMsg{
			Cmd:   "forum read " + strings.Join(args, " "),
			Error:  errMissingArgument,
			Output: f.LongHelp(),
		}
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return ResultMsg{
			Cmd:   "forum read " + strings.Join(args, " "),
			Error: err,
		}
	}

	err = c.Console.Forum.EnterTopicIndex(id)

	if err != nil {
		return ResultMsg{
			Cmd:   "forum read " + strings.Join(args, " "),
			Error: err,
		}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum read " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : read %s\n",
		c.Server.Address+c.Forum.Topic), 0)
}

// *****************************************************************************
// ****************************************************************** ForumLeave
// *****************************************************************************
type ForumLeave struct{}

func (f ForumLeave) ParseName() string {
	return "leave"
}

func (f ForumLeave) ShortHelp() string {
	return "quitte un Post ou un Topic"
}

func (f ForumLeave) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum leave\n")
	return b.String()
}

func (f ForumLeave) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum leave" + strings.Join(args, " "),
			Error: errNotConnected}
	}

	err := c.Console.Forum.LeaveTopic()

	if err != nil {
		return ResultMsg{
			Cmd:   "forum leave" + strings.Join(args, " "),
			Error: err,
		}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum show " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : leave: vous êtes maintenant dans %s\n",
		c.Server.Address+c.Forum.Topic), 0)
}

// *****************************************************************************
// ******************************************************************* ForumShow
// *****************************************************************************
type ForumShow struct{}

func (f ForumShow) ParseName() string {
	return "show"
}

func (f ForumShow) ShortHelp() string {
	return "montre le Topic/Post courant du forum"
}

func (f ForumShow) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum show [INDEX]\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  [INDEX] -- l'index (optionnel) du premier élément à montrer\n")

	return b.String()
}

func (f ForumShow) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum show " + strings.Join(args, " "),
			Error: errNotConnected}
	}

	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum show " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}

	index :=-1
	if len(args) >= 1 {
		argInt, err := strconv.Atoi(args[0])

		if err != nil {
			return ResultMsg{
				Cmd:   "forum show " + strings.Join(args, " "),
				Error: err,
			}
		}
		index = argInt
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : show %s\n",
		c.Server.Address+c.Forum.Topic),
		index)
}
// *************************************************************** ForumMoveStart
// Une commande qui n'est PAS INTERACTIVE
type ForumMoveStart struct{}
func (f ForumMoveStart) ParseName() string {
	return "moveStart"
}
func (f ForumMoveStart) ShortHelp() string {
	return "met l'index de Show à la première entrée"
}
func (f ForumMoveStart) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum moveStart\n")
	return b.String()
}
func (f ForumMoveStart) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum moveStart " + strings.Join(args, " "),
			Error: errNotConnected}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum moveStart " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : show %s\n",
		c.Server.Address+c.Forum.Topic),
		0)
}
// ***************************************************************** ForumMoveEnd
// Une commande qui n'est PAS INTERACTIVE
type ForumMoveEnd struct{}
func (f ForumMoveEnd) ParseName() string {
	return "moveStart"
}
func (f ForumMoveEnd) ShortHelp() string {
	return "met l'index de Show à la première entrée"
}
func (f ForumMoveEnd) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum moveStart\n")
	return b.String()
}
func (f ForumMoveEnd) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum moveStart " + strings.Join(args, " "),
			Error: errNotConnected}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum moveStart " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	index := len(c.Forum.TopicList) - maxEntryDisplay
	if index < 0 {
		index = 0
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : show %s\n",
		c.Server.Address+c.Forum.Topic),
		index)
}
// ***************************************************************** ForumNext
// Une commande qui n'est PAS INTERACTIVE
type ForumNext struct{}
func (f ForumNext) ParseName() string {
	return "next"
}
func (f ForumNext) ShortHelp() string {
	return "met l'index de Show à la première entrée"
}
func (f ForumNext) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum next\n")
	return b.String()
}
func (f ForumNext) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum next " + strings.Join(args, " "),
			Error: errNotConnected}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum next " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	index := c.Forum.IndexShow + maxEntryDisplay
	if index > len(c.Forum.TopicList) - maxEntryDisplay {
		 index = len(c.Forum.TopicList) - maxEntryDisplay
	}
	if index < 0 {
		index = 0
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : show %s\n",
		c.Server.Address+c.Forum.Topic),
		index)
}

// ***************************************************************** ForumPrev
// Une commande qui n'est PAS INTERACTIVE
type ForumPrev struct{}
func (f ForumPrev) ParseName() string {
	return "prev"
}
func (f ForumPrev) ShortHelp() string {
	return "met l'index de Show à la première entrée"
}
func (f ForumPrev) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum prev\n")
	return b.String()
}
func (f ForumPrev) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum prev " + strings.Join(args, " "),
			Error: errNotConnected}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum prev " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	index := c.Forum.IndexShow - maxEntryDisplay
	if index < 0 {
		index = 0
	}
	return ShowForumInternal(c, fmt.Sprintf("Forum : show %s\n",
		c.Server.Address+c.Forum.Topic),
		index)
}

// *********************************************************** ShowForumInternal
func ShowForumInternal(c *Client, heading string, index int ) ReadMsg {
	// construire la réponse à afficher
	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "%s\n", heading)
	fmt.Fprint(tw, "    => CtrlS = start, CtrlP = prev, CtrlN = next, CtrlE = end\n")

	for _, t := range c.Console.Forum.Display(index) {
		fmt.Fprintf(tw, "%s\n", t)
	}
	tw.Flush()

	return ReadMsg{
		Body: b.String(),
		Callbacks: []CmdMapping{
			{tea.KeyCtrlS, ForumMoveStart{}},
			{tea.KeyCtrlP, ForumPrev{}},
			{tea.KeyCtrlN, ForumNext{}},
			{tea.KeyCtrlE, ForumMoveEnd{}}},
	}
}

// *****************************************************************************
// ************************************************************** ForumWritePost
// *****************************************************************************
type ForumWritePost struct{}

func (f ForumWritePost) ParseName() string {
	return "post"
}

func (f ForumWritePost) ShortHelp() string {
	return "write an new Post to the Forum"
}

func (f ForumWritePost) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum post <TITLE>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  TITLE -- le TITRE du nouveau POST à créer\n")
	return b.String()
}

func (f ForumWritePost) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum post" + strings.Join(args, " "),
			Error: errNotConnected}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum post" + strings.Join(args, " "),
			Error: errForumUnreachable}
	}

	if len(args) < 1 {
		return ResultMsg{
			Cmd:   "forum post" + strings.Join(args, " "),
			Error:  errMissingArgument,
			Output: f.LongHelp(),
		}
	}

	title := strings.Join( args, " " )
	c.Forum.CurrentTitle = title
	// if c.enterWriteMod == false {
	// 	c.enterWriteMod = true
	// 	c.textarea.Focus()
	// }
	// construire la réponse à afficher
	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "\nForum : in Topic <%s>\n  Write new post <%s>\n",
		c.Server.Address+c.Forum.Topic,
		title,
	)

	tw.Flush()

	fmt.Printf("ForumWritePost: %s", b.String())
	return WriteMsg{
		Heading: b.String(),
		OkCmd: ForumPostMsg{},
	}
}
// *****************************************************************************
// **************************************************************** ForumPostMsg
// *****************************************************************************
// Une commande qui n'est PAS INTERACTIVE
type ForumPostMsg struct{}

func (f ForumPostMsg) ParseName() string {
	return "send"
}

func (f ForumPostMsg) ShortHelp() string {
	return "envoie le Post qui vient d'être écrit sur le Forum"
}

func (f ForumPostMsg) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum send\n")
	return b.String()
}

func (f ForumPostMsg) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum send" + strings.Join(args, " "),
			Error: errNotConnected}
	}

	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum send" + strings.Join(args, " "),
			Error: errForumUnreachable}
	}

	// Prepare info for posting message
	now := time.Now()
	timeStr := now.Format("150405")        // HHMMSS
	// TODO adjust pour le GN
	dateStr := "20"+now.Format("0102")     // 20+MMDD
	err := c.Forum.AddPost( dateStr,       // date
		timeStr,	                       // time
		c.Forum.CurrentTitle,              // title
		c.Console.Login,                   // user
		strings.Join( c.msgWrite, "\n"),   // body
	)

	if err != nil {
		return ResultMsg{
			Cmd:   "forum send" + c.Forum.CurrentTitle,
			Error: err,
		}
	}
	// construire la réponse à afficher
	b := strings.Builder{}
	tw := tw(&b)
		fmt.Fprintf(tw, "Forum : send %s\n",
			c.Forum.CurrentTitle )
	tw.Flush()

	return ResultMsg{
		Cmd:   "forum send" + strings.Join(args, " "),
		Output: b.String(),
	}
}

// *****************************************************************************
// *************************************************************** ForumAddTopic
// *****************************************************************************
type ForumAddTopic struct{}

func (f ForumAddTopic) ParseName() string {
	return "topic"
}

func (f ForumAddTopic) ShortHelp() string {
	return "add a new SUBTOPIC to the forum"
}

func (f ForumAddTopic) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum topic <TITLE>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  TITLE -- le TITRE du nouveau SUBTOPIC à créer\n")
	return b.String()
}

func (f ForumAddTopic) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum topic " + strings.Join(args, " "),
			Error: errNotConnected}
	}
	if c.Forum.Address == "" {
		return ResultMsg{
			Cmd:   "forum topic " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}

	if len(args) < 1 {
		return ResultMsg{
			Cmd:   "forum topic " + strings.Join(args, " "),
			Error:  errMissingArgument,
			Output: f.LongHelp(),
		}
	}

	title := strings.Join( args, " " )

	err := c.Forum.AddTopic(title)
	if err != nil {
		return ResultMsg{
			Cmd:   "forum topic " + c.Forum.CurrentTitle,
			Error: err,
		}
	}

	// construire la réponse à afficher
	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "Forum : create new Topic <%s> in <%s>\n",
		title,
		c.Server.Address+c.Forum.Topic,
	)
	tw.Flush()

	return ResultMsg{
		Cmd:   "forum topic" + strings.Join(args, " "),
		Output: b.String(),
	}
}
