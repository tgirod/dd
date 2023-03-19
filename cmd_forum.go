package main

// Une Command mais aussi un BubbleTea.Model qui profite de la fenêtre Modale
// du Client

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ForumCmd permet d'entrer dans un Forum (s'il existe)

// var ForumCmd = Node{
// 	Name: "forum",
// 	Help: "interagir avec un forum",
// 	Sub: []Command{
// 		ForumEnter{},
// 		ForumRead{},
// 		ForumLeave{},
// 		ForumShow{},
// 		ForumWritePost{},
// 		ForumAnswerPost{},
// 		ForumAddTopic{},
// 	},
// }

// *****************************************************************************
// ********************************************************* ForumCmd as Command
// *****************************************************************************
type ForumCmd struct {
	list list.Model
	Forum
	state         StateForum
	width, height int
}

func (f ForumCmd) ParseName() string {
	return "forum"
}

func (f ForumCmd) ShortHelp() string {
	return "entre dans un Forum (s'il existe)"
}

func (f ForumCmd) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(f.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  forum\n")
	return b.String()
}

func (f ForumCmd) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "forum",
			Error: errNotConnected}
	}

	fo, err := c.Server.GetForum()
	//c.Console.Forum = forum
	if err != nil {
		return ResultMsg{
			Cmd:   "forum " + strings.Join(args, " "),
			Error: errForumUnreachable}
	}
	// finfo in TopicList are already ordered T > Post, dates
	Items := GenFromTopic(fo)

	const defaultWidth = 50

	l := list.New(Items, ForumDelegate{}, defaultWidth, listHeight)
	// TODO Change Key => remove ForceQuit
	l.DisableQuitKeybindings()

	l.Title = "Le FORUM du dd.local"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return OpenModalMsg(&ForumCmd{
		list:  l,
		Forum: fo,
		state: TopicMode,
		c.width, c.height})
	// return ShowForumInternal(c, fmt.Sprintf("Forum : you are authorized to enter %s\n",
	// 	c.Server.Address), 0)
}

// ***************************************************************************
// ************************************************ ForumCmd as BubleTea.Model
// ***************************************************************************
func (m ForumCmd) Init() tea.Cmd {
	return nil
}

func (m ForumCmd) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			switch m.state {
			case TopicMode:
				i, ok := m.list.SelectedItem().(ItemForum)
				if ok {
					if i.IsTopic {
						m.Forum.EnterTopicIndex(i.Index)
						m.list.SetItems(GenFromTopic(m.Forum))
					} else {
						m.list.SetItems(GenFromPost(m.Forum, i.Index, i.Answers))
						m.list.ResetSelected()
						m.state = ThreadMode
					}
				}
				return m, nil
			case ThreadMode:
				i, ok := m.list.SelectedItem().(ItemForum)
				if ok {
					fmt.Printf("SELECT %s\n", i.Title)
				}
				return m, nil
			}
			return m, nil
		}
		if msg.String() == "backspace" {
			switch m.state {
			case TopicMode:
				m.Forum.LeaveTopic()
				m.list.SetItems(GenFromTopic(m.Forum))
				return m, nil
			case ThreadMode:
				m.list.SetItems(GenFromTopic(m.Forum))
				m.state = TopicMode
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
func (m ForumCmd) View() string {
	return "\n" + m.list.View()
}

// ***************************************************************************
// ************************** helpers for list;model: ForumItem, ForumDelegate
// ***************************************************************************

const listHeight = 35
const maxEntryDisplay = 20

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)

	errTopicExists   = errors.New("TOPIC existe déjà")
	errEntryNotFound = errors.New("entrée introuvable")
)

type ForumItem struct {
	Title string
	// Reference to the current Forum.TopicList
	Index    int
	IsTopic  bool
	IsThread bool
	Answers  []int
}

// ***************************************************************** ForumItem
// An Item with Title, etc for a list entry
func (i ForumItem) FilterValue() string {
	return i.Title
}
func (i ForumItem) AddAnswer(ida int) ForumItem {
	return ForumItem{i.Title, i.Index, i.IsTopic, i.IsThread,
		append(i.Answers, ida)}
}

func DefaultForumItem() ForumItem {
	return ForumItem{"", 0, false, false, nil}
}
func NewTopicForumItem(title string, idx int) ForumItem {
	return ForumItem{title, idx, true, false, nil}
}
func NewPostForumItem(title string, idx int) ForumItem {
	return ForumItem{title, idx, false, false, make([]int, 0, 0)}
}
func NewThreadForumItem(title string, idx int) ForumItem {
	return ForumItem{title, idx, false, true, nil}
}

// Find ItemForum with postTtitle
func IncItemForumWithPost(listIF []list.Item, listInfo []fs.FileInfo,
	postname string, idAnswer int) (int, error) {
	for id, itfo := range listIF {
		// an ItemForum ?
		if i, ok := itfo.(ItemForum); ok {
			if !i.IsTopic {
				_, oriTitle, _, _ := GetElements(listInfo[i.Index].Name())
				//fmt.Printf("Check with -%s-\n", oriTitle)
				if oriTitle == postname {
					//fmt.Printf("Found %d\n", id)
					listIF[id] = i.AddAnswer(idAnswer)
					return len(i.Answers) + 1, nil
				}
			}
		}
	}
	return 0, errEntryNotFound
}

func GenFromTopic(fo Forum) []list.Item {
	// finfo in TopicList are already ordered T > Post, dates
	Items := make([]list.Item, 0, len(fo.TopicList))
	for id, v := range fo.TopicList {
		if v.IsDir() {
			item := NewTopicItemForum(v.Name(), id)
			Items = append(Items, item)
		} else {
			_, title, _, _ := GetElements(v.Name())
			// a "real" Post
			if !strings.HasPrefix(title, "Re: ") {
				//fmt.Printf("Adding %s\n", title)
				item := NewPostItemForum(DecodePostTitle(v.Name()), id)
				Items = append(Items, item)
			} else {
				orig_name := title[4:len(title)]
				//fmt.Printf("Answer to -%s-\n", orig_name)
				// find item with this title
				IncItemForumWithPost(Items,
					fo.TopicList,
					orig_name, id)
			}
		}
	}
	return Items
}

func GenFromPost(fo Forum, idPost int, idList []int) []list.Item {
	// Develop Answers, already in order
	Items := make([]list.Item, 0, len(idList)+1)
	// First Post
	item := NewThreadItemForum(DecodePostTitle(fo.TopicList[idPost].Name()), idPost)
	Items = append(Items, item)
	// Then all Answers
	for _, idA := range idList {
		item := NewThreadItemForum(DecodePostTitle(fo.TopicList[idA].Name()), idA)
		Items = append(Items, item)
	}
	return Items
}

// ************************************************************* ForumDelegate
// For list.Model, delegate the rendering an updating of a ForumItem

type ForumDelegate struct{}

// Height returns the delegate's preferred height.
// 1 for Topic, 2 for Post (with Answers
// TODO can be dynamic ??
func (d ForumDelegate) Height() int {
	// 2 for the moment
	return 2
}

// Spacing returns the delegate's spacing.
func (d ForumDelegate) Spacing() int {
	return 0
}

// Update is the update loop for items. All messages in the list's update
// loop will pass through here except when the user is setting a filter.
// Use this method to perform item-level updates appropriate to this
// delegate.
func (d ForumDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

// Render prints an item.
func (d ForumDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		title, desc string
	)

	if i, ok := item.(ForumItem); ok {
		if i.IsTopic {
			title = "T> " + i.Title
			desc = ""
		} else {
			title = "P> " + i.Title
			desc = fmt.Sprintf("  %3d réponse(s)", len(i.Answers))
		}
	} else {
		return
	}

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	if m.Width() <= 0 {
		// short-circuit
		return
	}

	// // Prevent text from exceeding list width
	// textwidth := uint(m.width - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight())
	// title = truncate.StringWithTail(title, textwidth, ellipsis)
	// desc = truncate.StringWithTail(desc, textwidth, ellipsis)

	str := fmt.Sprintf("%s\n%s", title, desc)
	fmt.Fprint(w, fn(str))
	return
}
