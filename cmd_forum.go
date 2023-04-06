package main

// Une Command mais aussi un BubbleTea.Model qui profite de la fenêtre Modale
// du Client

// FIXME if ENTER on a Post with 0 Answer, enter directly reading ?????
// FIXME someimes, shorter list does not clear previous entries

/*
   ForumCmd::Run(*Client, args)
     + ForumCmd.Forum, with Forum.User
     + ForumCmd.list, from Items = GenFromTopic(Forum), list.Title
     + state = TopicState

   TopicState
     -- ENTER --> Topic ? Forum.EnterTopicIndex, list.Items = GenFromTopic(Forum)
                  Post ? list.Items = GenFromPost(m.Forum, i.Index, i.Answers))
                    + state = ThreadMode

     -- QUIT ---> Forum.LeaveTopic
                    + list.Items = GenFromTopic(m.Forum)

     -- +TOPIC -> [EnterWriteMsg](needTopic:true) => WritingMode

     -- +POST --> [EnterWriteMsg](needTopic:false, needTitle) => WritingMode

   ThreadMode
     -- ENTER --> m.reader.SetPost( Forum, index of Answers, Index )
                    + state = ReadingMode
     -- QUIT ---> list.Items = GenFromTopic(m.Forum)
                    + state = Topicmode

   ReadingMode
     -- QUIT ---> [QuitReadingMsg] => ThreadMode
     -- +/- ----> nothing changes in Forum
     -- ANSWER -> [EnterWritingMsg](title = r.forum.GetTitleFromTopic())
                    => writer.SetAnsweredPost(Forum, title)
                    => WritingMode

   WritingMode
     -- CANCEL -> [QuitWritingMsg]
                  => UpdatePostInThread => ThreadMode
                  => UpdateTopic => TopicMode
     -- SEND ---> Forum.AddPost( ... )
                    + [QuitWritingMsg]
                       => UpdatePostInThread => ThreadMode
                       => UpdateTopic => TopicMode
   --------------------------------
   func GenFromTopic(forum) {
     + Item for each Topic + 0 answers
                     Post + Answers = list of finfo index
   }

   func GenFromPost(forum, idPost, idList) {
     + Item for each Post And forum.TopicList[idA] for each idA in idList
   }

 **/

import (
	//"errors"
	"fmt"
	"io"
	"io/fs"

	//"strconv"
	"strings"
	"time"

	hhelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type ForumMsg struct{}

var forum = Cmd{
	Name:      "forum",
	ShortHelp: "consulter le forum d'un serveur",
	Connected: true,
	Parse: func(args []string) any {
		return ForumMsg{}
	},
}

// *****************************************************************************
// ********************************************************* ForumCmd as Command
// *****************************************************************************
type StateForum int

const (
	TopicMode StateForum = iota
	ThreadMode
	ReadingMode
	WritingMode
)

type QuitReadingMsg struct{}
type EnterWritingMsg struct {
	needTopic bool
	title     string
}
type QuitWritingMsg struct{}

// Cmd separated from CmdModel to be able to use ShortHelp in Cmd and
// also as a BubbleTea.help
type ForumCmd struct{}
type ForumCmdModel struct {
	list   list.Model  // list Topics and Posts
	help   hhelp.Model // help while in List
	keys   *listKeyMap // additionnal Keys for list
	width  int
	height int
	reader ReaderModel
	writer WriterModel

	Forum
	prevState StateForum
	state     StateForum
}

// Implement Command = ParseName, ShortHelp, LongHelp, Run
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
		return Eval{
			Cmd:   "forum",
			Error: errNotConnected,
		}
	}

	fo, err := c.Server.GetForum()
	//c.Console.Forum = forum
	if err != nil {
		return Eval{
			Cmd:   "forum " + strings.Join(args, " "),
			Error: errForumUnreachable,
		}
	}
	fo.User = c.Login
	// finfo in TopicList are already ordered T > Post, dates
	Items := GenFromTopic(fo)

	const defaultWidth = 50

	listKeys := newListKeyMap()
	l := list.New(Items, ForumDelegate{}, 5, 5) // DEL defaultWidth, listHeight)
	// TODO Change Key => remove ForceQuit
	l.DisableQuitKeybindings()

	// FIXME/TODO update title
	l.Title = "Le FORUM du dd.local"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.Help.Styles.FullDesc = helpTextStyle
	l.Help.Styles.ShortDesc = helpTextStyle
	l.Help.Styles.FullKey = helpTextStyle
	l.Help.Styles.ShortKey = helpTextStyle

	h := hhelp.New()
	h.Styles.FullDesc = helpTextStyle
	h.Styles.ShortDesc = helpTextStyle
	h.Styles.FullKey = helpTextStyle
	h.Styles.ShortKey = helpTextStyle

	fcm := ForumCmdModel{
		list:      l,
		help:      h,
		keys:      listKeys,
		reader:    NewReader(),
		writer:    NewWriter(),
		Forum:     fo,
		prevState: TopicMode,
		state:     TopicMode,
	}

	return OpenModalMsg(&fcm)
}

// ***************************************************************************
// ******************************************* ForumCmdModel as BubleTea.Model
// ***************************************************************************
type listKeyMap struct {
	enterTopic key.Binding // also enterPost
	exitTopic  key.Binding // also exitPost
	addTopic   key.Binding
	addPost    key.Binding
	answerPost key.Binding
	quitForum  key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		enterTopic: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "enter Topic/Post"),
		),
		exitTopic: key.NewBinding(
			key.WithKeys("backspace"),
			key.WithHelp("backspace", "exit Topic/Post"),
		),
		addTopic: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "add Topic"),
		),
		addPost: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "add Post"),
		),
		answerPost: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "answer Post"),
		),
		quitForum: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("Esc/q", "exit Forum"),
		),
	}
}
func (m ForumCmdModel) AdditionalShortHelp() []key.Binding {
	return []key.Binding{}
}
func (m ForumCmdModel) ShortHelp() []key.Binding {
	kb := []key.Binding{
		m.list.KeyMap.CursorUp,
		m.list.KeyMap.CursorDown,
		m.list.KeyMap.ShowFullHelp,
	}
	return kb
}
func (m ForumCmdModel) AdditionalFullHelp() []key.Binding {
	fmt.Printf("__AddFullHelp state=%v\n", m.state)
	if m.state == TopicMode {
		return []key.Binding{
			m.keys.enterTopic,
			m.keys.exitTopic,
			m.keys.addTopic,
			m.keys.addPost,
			m.keys.quitForum,
		}
	} else {
		return []key.Binding{
			m.keys.enterTopic,
			m.keys.exitTopic,
			m.keys.answerPost,
			m.keys.quitForum,
		}
	}
}
func (m ForumCmdModel) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{{
		m.list.KeyMap.CursorUp,
		m.list.KeyMap.CursorDown,
		m.list.KeyMap.NextPage,
		m.list.KeyMap.PrevPage,
		m.list.KeyMap.GoToStart,
		m.list.KeyMap.GoToEnd,
	}}

	return append(kb,
		m.AdditionalFullHelp(),
		[]key.Binding{
			m.list.KeyMap.CloseFullHelp,
		})
}

// Implement bublletea = Init, Update, View
func (m ForumCmdModel) Init() tea.Cmd {
	return nil
}

func (m ForumCmdModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// WindowSizeMsg -> pass also to Reader (if active), aggregate
	// KeyMsg -> pass either to Reader or deal here
	// other -> pass either to Reader or deal here

	fmt.Print("__ForumCmdModel::Update start msg=", msg, " state=", m.state, "\n")

	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Print("__forum Size=", msg, "\n")
		m.width = msg.Width - 4
		m.height = msg.Height - 6
		m.list.SetWidth(msg.Width - 4)
		//m.list.SetSize(msg.Width-4, msg.Height-6)

		m.reader, _ = m.reader.Update(msg)
		m.writer, _ = m.writer.Update(msg)
		// FIXME padding+margin+border instead of "magic" number
		return m, nil

	case QuitReadingMsg:
		fmt.Print("__forum::Update QuitReading=", msg, "\n")
		m.state = ThreadMode
		m.prevState = ReadingMode
		return m, nil

	case EnterWritingMsg:
		fmt.Print("__Forum::Update case EnterWritingMsg=", msg, "\n")
		m.prevState = m.state
		m.state = WritingMode
		if msg.title != "" {
			m.writer.SetAnsweredPost(&(m.Forum), msg.title)
		} else {
			m.writer.SetNewPost(&(m.Forum), msg.needTopic)
		}
		return m, nil

	case QuitWritingMsg:
		fmt.Print("__forum::Update QuitWriting=", msg, "\n")
		if m.prevState == ReadingMode {
			fmt.Print("  prev was ReadingMode\n")
			m.state = ThreadMode
			previousIndex := m.Forum.IndexPost
			fmt.Printf("  and prevIndex=%d\n", previousIndex)
			items, newIndex := UpdatePostThread(&(m.Forum), previousIndex)
			fmt.Printf("  in Forum, len(TopicList)=%d\n", len(m.Forum.TopicList))
			m.list.SetItems(items)
			if len(items) > 0 {
				m.list.Select(newIndex)
			}
		} else if m.prevState == TopicMode {
			fmt.Print("  prev was TopicMode\n")

			m.state = TopicMode
			previousIndex := m.list.Index()
			fmt.Printf("  and prevIndex=%d\n", previousIndex)
			items, newIndex := UpdateTopic(&(m.Forum), previousIndex)
			fmt.Printf("  in Forum, len(TopicList)=%d\n", len(m.Forum.TopicList))
			m.list.SetItems(items)
			if len(items) > 0 {
				m.list.Select(newIndex)
			}
		} else {
			fmt.Print("  prev was NOT Reading, NOT TOPIC\n")

			m.state = m.prevState
		}
		m.prevState = WritingMode
		return m, nil

	case tea.KeyMsg:
		// FIXME remove ??
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		switch m.state {
		case WritingMode:
			m.writer, cmd = m.writer.Update(msg)
			cmds = append(cmds, cmd)

		case ReadingMode:
			m.reader, cmd = m.reader.Update(msg)
			cmds = append(cmds, cmd)

		default:
			m, cmd = m.HandleKeys(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

// Handle Keys for TopicMode and ThreadMode
func (m ForumCmdModel) HandleKeys(msg tea.KeyMsg) (ForumCmdModel, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.quitForum):
		return m, func() tea.Msg { return CloseModalMsg{} }

	case key.Matches(msg, m.keys.exitTopic):
		switch m.state {
		case TopicMode:
			m.Forum.LeaveTopic()
			m.list.SetItems(GenFromTopic(m.Forum))
		case ThreadMode:
			m.list.SetItems(GenFromTopic(m.Forum))
			m.state = TopicMode
			m.prevState = ThreadMode
		}
		return m, nil

	case key.Matches(msg, m.keys.enterTopic):
		//DEBUG fmt.Print("__ForumCmdModel::Update EnterTopic\n")
		switch m.state {
		case TopicMode:
			i, ok := m.list.SelectedItem().(ForumItem)
			if ok {
				if i.IsTopic {
					m.Forum.EnterTopicIndex(i.Index)
					m.list.SetItems(GenFromTopic(m.Forum))
				} else {
					m.list.SetItems(GenFromPost(m.Forum, i.Index, i.Answers))
					m.list.ResetSelected()
					m.state = ThreadMode
					m.prevState = TopicMode
				}
			}
			return m, nil

		case ThreadMode:
			i, ok := m.list.SelectedItem().(ForumItem)
			if ok {
				//DEBUG fmt.Printf("SELECT %s\n", i.Title)

				// need indexes of all post in thread
				var indexes []int
				for _, item := range m.list.Items() {
					foIt, _ := item.(ForumItem)
					indexes = append(indexes, foIt.Index)
				}

				fmt.Printf("SELECT %d/%d\n", m.list.Index(), len(i.Answers))
				m.reader.SetPosts(&(m.Forum),
					indexes,
					m.list.Index(),
				)
				m.state = ReadingMode
				m.prevState = ThreadMode
			}
			return m, nil
		}
	case key.Matches(msg, m.keys.addTopic):
		if m.state == TopicMode {
			return m, func() tea.Msg {
				return EnterWritingMsg{
					needTopic: true,
				}
			}
		}

	case key.Matches(msg, m.keys.addPost):
		if m.state == TopicMode {
			return m, func() tea.Msg {
				return EnterWritingMsg{
					needTopic: false,
				}
			}
		}
	case key.Matches(msg, m.list.KeyMap.ShowFullHelp):
		fallthrough
	case key.Matches(msg, m.list.KeyMap.CloseFullHelp):
		m.help.ShowAll = !m.help.ShowAll
		return m, nil

	}
	// by default -> list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ForumCmdModel) View() string {
	// DEBUG
	fmt.Printf("__ForumCmdModelView size=%d, %d\n", m.list.Width(), m.list.Height())
	var subView string
	if m.state == ReadingMode {
		subView = m.reader.View()
	} else if m.state == WritingMode {
		subView = m.writer.View()
	} else {
		availHeight := m.height - 4
		helpStr := m.help.View(m)
		availHeight -= lg.Height(helpStr)
		m.list.SetHeight(availHeight)
		listStr := m.list.View()

		subView = lg.JoinVertical(lg.Left,
			listStr, helpStr)
	}
	return lg.Place(m.width, m.height, lg.Top, lg.Left,
		lg.JoinVertical(lg.Left,
			m.StatusView(),
			subView),
		lg.WithWhitespaceChars(" "))
}

func (m ForumCmdModel) StatusView() string {
	forumName := forumNameStyle.Render("FORUM " + m.Forum.Address)
	topicStr := forumTopicStyle.Render(
		fmt.Sprintf("\n Topic <%s>", m.Forum.Topic))
	if m.state == ThreadMode {
		i, ok := m.list.SelectedItem().(ForumItem)
		if ok {
			topicStr += forumTopicStyle.Render(fmt.Sprintf("\n Thread: %s",
				i.Title))
		}
	}
	render := forumStatusStyle.Render(forumName + topicStr)
	fmt.Printf("__Status view width=%d\n", lg.Width(render))

	return render
}

// ***************************************************************************
// ******************************************************************** Reader
// ***************************************************************************
type readerKeyMap struct {
	up               key.Binding
	down             key.Binding
	nextPostInThread key.Binding
	prevPostInThread key.Binding
	answerPost       key.Binding
	quitReading      key.Binding
}

func newReaderKeyMap(v viewport.Model) *readerKeyMap {
	return &readerKeyMap{
		up:   v.KeyMap.Up,
		down: v.KeyMap.Down,
		nextPostInThread: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "next"),
		),
		prevPostInThread: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "prev"),
		),
		answerPost: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "answer"),
		),
		quitReading: key.NewBinding(
			key.WithKeys("q", "backspace"),
			key.WithHelp("q/backspace", "quit"),
		),
	}
}

// ReaderModel is made of a viewport, a KeyMap, help
type ReaderModel struct {
	viewport viewport.Model
	help     hhelp.Model // help while reading Posts
	keys     *readerKeyMap
	forum    *Forum
	answers  []int // index of all Post in Thread to read
	index    int   // index of the current Post read

	title string // fullTitle of Post currently read
	body  string // content of Post being read

	// FIXME style for title
	// FIXME style for body
	// FIXME style for help
}

// New ReaderModel
func NewReader() ReaderModel {
	view := viewport.New(1, 1) // temporary width and height
	k := newReaderKeyMap(view)
	h := hhelp.New()
	h.Styles.FullDesc = helpTextStyle
	h.Styles.ShortDesc = helpTextStyle
	h.Styles.FullKey = helpTextStyle
	h.Styles.ShortKey = helpTextStyle

	return ReaderModel{
		viewport: view,
		keys:     k,
		help:     h,
	}
}

// Implement help Interface
func (r ReaderModel) ShortHelp() []key.Binding {
	return []key.Binding{
		r.keys.up,
		r.keys.down,
		r.keys.nextPostInThread,
		r.keys.prevPostInThread,
		r.keys.answerPost,
		r.keys.quitReading,
	}
}
func (r ReaderModel) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{
		r.ShortHelp(),
	}
	return kb
}

func (r *ReaderModel) SetPosts(forumRef *Forum, answers []int, indexThread int) {
	fmt.Printf("__SetPosts len=%d\n", len(answers))
	r.forum = forumRef
	r.answers = answers
	r.ReadPost(indexThread)
}
func (r *ReaderModel) ReadPost(index int) {
	r.index = index
	fileIndex := r.answers[index]
	// get Title
	r.title = DecodePostTitle(r.forum.TopicList[fileIndex].Name())
	// FIXME check error ?
	body, _ := r.forum.GetPost(fileIndex)
	r.body = body

	//title string, body string) {
	fmt.Print("__ReaderModel::ReadPost\n", body, "\n---------------\n")
	r.viewport.SetContent(bodyStyle.Render(body))
}

// Implement bubbleTea Interface
func (r ReaderModel) Init() tea.Cmd {
	return nil
}

func (r ReaderModel) Update(msg tea.Msg) (ReaderModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// also set viewport
		//m.viewport = viewport.New(msg.Width-8, msg.Height-8)
		r.viewport.Width = msg.Width - 4
		r.viewport.Height = msg.Height - 10
		// FIXME
		bodyStyle.Width(msg.Width - 4)
		return r, nil

	case tea.KeyMsg:
		fmt.Print("__ReaderModel::Update key= ", msg, "\n")

		switch {
		case key.Matches(msg, r.keys.quitReading):
			return r, func() tea.Msg { return QuitReadingMsg{} }

		case key.Matches(msg, r.keys.nextPostInThread):
			r.ReadPost(min(r.index+1, len(r.answers)-1))

		case key.Matches(msg, r.keys.prevPostInThread):
			r.ReadPost(max(0, r.index-1))

		case key.Matches(msg, r.keys.answerPost):
			// FIXME deal with error ?
			title, _ := r.forum.GetTitleFromTopic()
			if !strings.HasPrefix(title, "Re: ") {
				title = "Re: " + title
			}
			return r, func() tea.Msg {
				return EnterWritingMsg{
					title:     title,
					needTopic: false,
				}
			}
		}
		// todo
	}

	var cmdV tea.Cmd
	r.viewport, cmdV = r.viewport.Update(msg)
	cmds = append(cmds, cmdV)

	var cmdH tea.Cmd
	r.help, cmdH = r.help.Update(msg)
	cmds = append(cmds, cmdH)

	return r, tea.Batch(cmds...)
}

func (r ReaderModel) View() string {
	// Header
	threadStr := fmt.Sprintf("%d/%d", r.index+1, len(r.answers))
	spaces := strings.Repeat(" ", max(0, r.viewport.Width-lg.Width(threadStr)-lg.Width(r.title)))
	titleStr := titleStyle.Render(lg.JoinHorizontal(lg.Right,
		r.title,
		spaces,
		threadStr))

	// Content
	//fmt.Print("__ReaderModel::View\n", r.body, "\n---------------------\n")

	// Footer
	info := fmt.Sprintf("%3.f%%", r.viewport.ScrollPercent()*100)
	line := strings.Repeat("─", max(0, r.viewport.Width-lg.Width(info)))
	footer := lg.JoinHorizontal(lg.Right, line, info)

	return lg.JoinVertical(lg.Left,
		titleStr,
		r.viewport.View(),
		footer,
		r.help.View(r),
	)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ***************************************************************************
// ******************************************************************** Writer
// ***************************************************************************
type writerKeyMap struct {
	cancel   key.Binding
	send     key.Binding
	validate key.Binding
}

func newWriterKeyMap() *writerKeyMap {
	return &writerKeyMap{
		cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("Esc", "Cancel"),
		),
		send: key.NewBinding(
			key.WithKeys("ctrl+p"),
			key.WithHelp("Ctrl-P", "Send Post"),
		),
		validate: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("Enter", "Validate Title"),
		),
	}
}

// writerModel is build around textarea
type WriterModel struct {
	textarea  textarea.Model
	textinput textinput.Model
	help      hhelp.Model
	keys      *writerKeyMap
	forum     *Forum

	needTopic bool // true if we only need a new Topic
	needTitle bool
	title     string
	body      []string
}

func NewWriter() WriterModel {
	ta := textarea.New()
	ta.Placeholder = "Ecriver votre Post pour le forum..."
	ta.Prompt = "| "
	ta.KeyMap.InsertNewline.SetEnabled(true) // allow Enter to put multiline

	ti := textinput.New()
	ti.Placeholder = "Votre titre..."

	k := newWriterKeyMap()
	h := hhelp.New()
	h.Styles.FullDesc = helpTextStyle
	h.Styles.ShortDesc = helpTextStyle
	h.Styles.FullKey = helpTextStyle
	h.Styles.ShortKey = helpTextStyle

	return WriterModel{
		textarea:  ta,
		textinput: ti,
		keys:      k,
		help:      h,
		needTopic: false,
		needTitle: true,
	}
}

// Implement help Interface
func (w WriterModel) ShortHelp() []key.Binding {
	if w.needTitle {
		return []key.Binding{
			w.keys.cancel,
			w.keys.validate,
		}
	} else {
		return []key.Binding{
			w.keys.cancel,
			w.keys.send,
		}
	}
}
func (w WriterModel) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{
		w.ShortHelp(),
	}
	return kb
}

func (w *WriterModel) SetAnsweredPost(forumRef *Forum, title string) {
	w.textarea.Reset()
	w.forum = forumRef
	w.needTitle = false
	w.title = title
	w.textarea.Focus()
}

func (w *WriterModel) SetNewPost(forumRef *Forum, needTopic bool) {
	w.textarea.Reset()
	w.forum = forumRef
	w.needTopic = needTopic
	w.needTitle = true
	w.textinput.Focus()
}

// Implement bubbleTea Interface
func (w WriterModel) Init() tea.Cmd {
	return nil
}

func (w WriterModel) Update(msg tea.Msg) (WriterModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// also set textarea
		w.textarea.SetWidth(msg.Width - 4)
		w.textarea.SetHeight(msg.Height - 10)
		//w.textinput.Width = msg.Width - 4
		return w, nil

	case tea.KeyMsg:
		fmt.Print("__WriterModel::Update key= ", msg, "\n")

		switch {
		case key.Matches(msg, w.keys.cancel):
			return w, func() tea.Msg { return QuitWritingMsg{} }

		case key.Matches(msg, w.keys.send):
			if !w.needTitle {
				w.body = append(w.body, w.textarea.Value())
				w.textarea.Reset()

				fmt.Printf("__Writer::Update send %s %d lines from %s\n", w.title, len(w.body), w.forum.User)

				// Prepare new Post
				title := w.title
				now := time.Now()
				timeStr := now.Format("150405") // HHMMSS
				// FIXME/TODO adjust pour le GN
				dateStr := "20" + now.Format("0102") // 20+MMDD
				// FIXME/TODO User "public" can NOT POST
				err := w.forum.AddPost(dateStr, // date
					timeStr, // time
					title,   // title
					// FIXMEfor Debug
					//w.forum.User,               // user
					"DEBUG",
					strings.Join(w.body, "\n"), // body
				)
				// FIXME deal with err
				if err != nil {
					fmt.Printf("__Writer::Update err=%v\n", err)
				}
				// FIXME/NEXT must also update Thread (new Post)
				return w, func() tea.Msg { return QuitWritingMsg{} }
			}

		case key.Matches(msg, w.keys.validate):
			if w.needTopic {
				w.title = w.textinput.Value()
				err := w.forum.AddTopic(w.title)
				// FIXME deal with err
				if err != nil {
					fmt.Printf("__Writer::Update err=%v\n", err)
				}
				return w, func() tea.Msg { return QuitWritingMsg{} }

			}
			if w.needTitle {
				w.title = w.textinput.Value()
				w.needTitle = false
				w.textarea.Focus()
				return w, nil
			}
		}
	}

	var cmd tea.Cmd
	if w.needTitle {
		w.textinput, cmd = w.textinput.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		w.textarea, cmd = w.textarea.Update(msg)
		cmds = append(cmds, cmd)
	}

	return w, tea.Batch(cmds...)
}

func (w WriterModel) View() string {
	if w.needTopic {
		return lg.JoinVertical(lg.Left,
			// FIXME set Width
			"Indiquez le titre de votre Topic",
			w.textinput.View(),
			w.help.View(w),
		)
	} else if w.needTitle {
		return lg.JoinVertical(lg.Left,
			// FIXME set Width
			"Indiquez le titre de votre Post",
			w.textinput.View(),
			w.help.View(w),
		)
	} else {

		// Header
		titleStr := lg.JoinHorizontal(lg.Right,
			// FIXME set msg
			"Answering: ",
			titleStyle.Render(w.title))

		return lg.JoinVertical(lg.Left,
			titleStr,
			w.textarea.View(),
			w.help.View(w),
		)
	}
}

// ***************************************************************************
// ************************** helpers for list;model: ForumItem, ForumDelegate
// ***************************************************************************

const listHeight = 35
const maxEntryDisplay = 20

var (
	titleStyle        = lg.NewStyle().MarginLeft(2).Background(lg.Color("2"))
	itemStyle         = lg.NewStyle().PaddingLeft(4)
	bodyStyle         = lg.NewStyle().PaddingLeft(4)
	selectedItemStyle = lg.NewStyle().PaddingLeft(2).Foreground(lg.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	helpTextStyle     = lg.NewStyle().Foreground(lg.Color("142"))
	quitTextStyle     = lg.NewStyle().Margin(1, 0, 2, 4)

	forumStatusStyle = lg.NewStyle()
	forumNameStyle   = lg.NewStyle().Background(lg.Color("2"))
	forumTopicStyle  = lg.NewStyle().Foreground(lg.Color("2"))

	// errTopicExists   = errors.New("TOPIC existe déjà")
	// errEntryNotFound = errors.New("entrée introuvable")
)

type ForumItem struct {
	// Title is "Post title [Author, date time]
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

// Find ForumItem with postTtitle
func IncForumItemWithPost(listIF []list.Item, listInfo []fs.FileInfo,
	postname string, idAnswer int) (int, error) {
	for id, itfo := range listIF {
		// an ForumItem ?
		if i, ok := itfo.(ForumItem); ok {
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
			item := NewTopicForumItem(v.Name(), id)
			Items = append(Items, item)
		} else {
			_, title, _, _ := GetElements(v.Name())
			// a "real" Post
			if !strings.HasPrefix(title, "Re: ") {
				//fmt.Printf("Adding %s\n", title)
				item := NewPostForumItem(DecodePostTitle(v.Name()), id)
				Items = append(Items, item)
			} else {
				orig_name := title[4:]
				//fmt.Printf("Answer to -%s-\n", orig_name)
				// find item with this title
				IncForumItemWithPost(Items,
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
	item := NewThreadForumItem(DecodePostTitle(fo.TopicList[idPost].Name()), idPost)
	Items = append(Items, item)
	// Then all Answers
	for _, idA := range idList {
		item := NewThreadForumItem(DecodePostTitle(fo.TopicList[idA].Name()), idA)
		Items = append(Items, item)
	}
	return Items
}

// Was reading Post/Answer idPost, forum may have changed, reread
func UpdatePostThread(fo *Forum, idPost int) ([]list.Item, int) {
	// Store fullname of current Post/Answers
	fullname := fo.Post
	_, fullTitle, _, _ := GetElements(fullname)
	// oriTitle is the title of the "root" Post which is answered
	oriTitle := fullTitle[0:]
	if strings.HasPrefix(oriTitle, "Re: ") {
		oriTitle = oriTitle[4:]
	}
	fmt.Printf("__UpdatePostThread reading=%s, title=%s\n", fullname, oriTitle)
	// Reread Forum current directory
	fmt.Printf("  GetFiles(%s)\n", fo.Topic)
	fo.GetFiles(fo.Topic)
	// Get index of the Post in the new Forum.Topiclist
	newIdPost := -1
	for idFile, finfo := range fo.TopicList {
		if finfo.Name() == fullname {
			newIdPost = idFile
		}
	}
	fmt.Printf("  newIdPost=%d\n", newIdPost)

	// fill Items of this Topic
	items := GenFromTopic(*fo)
	fmt.Printf("  #items=%d, #TopicList=%d\n", len(items), len(fo.TopicList))

	// Look for Item with newIdPost in its Answers
	newThreadIndex := -1
	newIndexInThread := -1
	var answerIndexes = []int{}
lookForAnswer:
	for _, it := range items {
		if i, ok := it.(ForumItem); ok {
			if !i.IsTopic {
				fmt.Printf("  Elements of %s\n", i.Title)
				for idInA, idA := range i.Answers {
					if idA == newIdPost {
						newThreadIndex = i.Index
						answerIndexes = i.Answers
						newIndexInThread = idInA
						break lookForAnswer
					}
				}
			}
		}
	}
	fmt.Printf("  Post found idPost=%d, idAnswer=%d, nb_answers=%d\n", newIdPost, newThreadIndex, len(answerIndexes))
	fmt.Printf("  Answers=%v\n", answerIndexes)

	// fill items for this Thread
	items = GenFromPost(*fo, newThreadIndex, answerIndexes)

	return items, newIndexInThread
}

// Was reading Topic, using Forum.TopicList[prevIdTopic], forum may have changed, reread
func UpdateTopic(fo *Forum, prevIdTopic int) ([]list.Item, int) {
	// name of Post/Topic that was read
	updateIndex := len(fo.TopicList) > 0
	fullname := ""
	if updateIndex {
		fullname = fo.TopicList[prevIdTopic].Name()
	}

	// Reread Forum current directory
	fmt.Printf("  GetFiles(%s)\n", fo.Topic)
	fo.GetFiles(fo.Topic)

	// Get index of the Topic/Post in the new Forum.Topiclist
	newIdTopic := -1
	if updateIndex {
		for idFile, finfo := range fo.TopicList {
			if finfo.Name() == fullname {
				newIdTopic = idFile
			}
		}
	}
	fmt.Printf("  newIdPost=%d\n", newIdTopic)

	// fill Items of this Topic
	items := GenFromTopic(*fo)
	fmt.Printf("  #items=%d, #TopicList=%d\n", len(items), len(fo.TopicList))

	return items, newIdTopic
}

// ************************************************************* ForumDelegate
// For list.Model, delegate the rendering an updating of a ForumItem
// FIXME Maybe ? enterPost/exitPost as keys for Delegate ??

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
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + s[0])
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
