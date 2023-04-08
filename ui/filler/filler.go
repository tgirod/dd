package filler

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type LoginFiller interface {
	SetLogin(string) LoginFiller
}

type PasswordFiller interface {
	SetPassword(string) PasswordFiller
}

type SubjectFiller interface {
	SetSubject(string) SubjectFiller
}

type ContentFiller interface {
	SetContent(string) ContentFiller
}

type FilledMsg struct{}

type Model struct {
	Title  string  // titre de la fenêtre modale
	Msg    tea.Msg // message à remplir
	Fields []field // champs de saisie
	Focus  int     // champ sélectionné
}

type keymap struct {
	Next     key.Binding
	Validate key.Binding
	Cancel   key.Binding
}

var keys = keymap{
	Next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "champ suivant"),
	),
	Validate: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "valide la commande"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annule la commande"),
	),
}

type field interface {
	Focus() tea.Cmd
	Blur()
	Value() string
	Update(tea.Msg) (field, tea.Cmd)
	View() string
}

// LoginField permet de saisir un login
type LoginField struct {
	*textinput.Model
}

func (f LoginField) Update(msg tea.Msg) (field, tea.Cmd) {
	model, cmd := f.Model.Update(msg)
	f.Model = &model
	return f, cmd
}

func newLoginField() LoginField {
	ti := textinput.New()
	ti.Placeholder = "login"
	return LoginField{&ti}
}

// PasswordField permet de saisir un mot de passe
type PasswordField struct {
	*textinput.Model
}

func (f PasswordField) Update(msg tea.Msg) (field, tea.Cmd) {
	model, cmd := f.Model.Update(msg)
	f.Model = &model
	return f, cmd
}

func newPasswordField() PasswordField {
	ti := textinput.New()
	ti.Placeholder = "password"
	ti.EchoMode = textinput.EchoPassword
	return PasswordField{&ti}
}

// SubjectField permet de saisir le sujet d'un message
type SubjectField struct {
	*textinput.Model
}

func (f SubjectField) Update(msg tea.Msg) (field, tea.Cmd) {
	model, cmd := f.Model.Update(msg)
	f.Model = &model
	return f, cmd
}

func newSubjectField() SubjectField {
	ti := textinput.New()
	ti.Placeholder = "subject"
	return SubjectField{&ti}
}

type ContentField struct {
	*textarea.Model
}

func (f ContentField) Update(msg tea.Msg) (field, tea.Cmd) {
	model, cmd := f.Model.Update(msg)
	f.Model = &model
	return f, cmd
}

func newContentField() ContentField {
	ta := textarea.New()
	ta.Placeholder = "content"
	return ContentField{&ta}
}

func New(title string, msg tea.Msg) *Model {
	m := Model{
		Title: title,
		Msg:   msg,
	}

	if _, ok := msg.(LoginFiller); ok {
		m.Fields = append(m.Fields, newLoginField())
	}

	if _, ok := msg.(PasswordFiller); ok {
		m.Fields = append(m.Fields, newPasswordField())
	}

	if _, ok := msg.(SubjectFiller); ok {
		m.Fields = append(m.Fields, newSubjectField())
	}

	if _, ok := msg.(ContentFiller); ok {
		m.Fields = append(m.Fields, newContentField())
	}

	return &m
}

func (m *Model) next() tea.Cmd {
	var cmds []tea.Cmd
	m.Focus = (m.Focus + 1) % len(m.Fields)

	for i := range m.Fields {
		if i == m.Focus {
			cmds = append(cmds, m.Fields[i].Focus())
		} else {
			m.Fields[i].Blur()
		}
	}

	return tea.Batch(cmds...)
}

func (m *Model) Init() tea.Cmd {
	if len(m.Fields) == 0 {
		return nil // ne devrait pas arriver
	}
	return m.Fields[0].Focus()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Next):
			cmds = append(cmds, m.next())

		case key.Matches(msg, keys.Validate):
			// passer au champ suivant si on n'est pas sur le dernier
			if m.Focus != len(m.Fields)-1 {
				cmds = append(cmds, m.next())
				break
			}

			// récupérer les valeurs des champs
			for _, field := range m.Fields {
				switch field := field.(type) {
				case LoginField:
					if msg, ok := m.Msg.(LoginFiller); ok {
						m.Msg = msg.SetLogin(field.Value())
					}
				case PasswordField:
					if msg, ok := m.Msg.(PasswordFiller); ok {
						m.Msg = msg.SetPassword(field.Value())
					}
				case SubjectField:
					if msg, ok := m.Msg.(SubjectFiller); ok {
						m.Msg = msg.SetSubject(field.Value())
					}
				case ContentField:
					if msg, ok := m.Msg.(ContentFiller); ok {
						m.Msg = msg.SetContent(field.Value())
					}
				}
			}

			// retourner le message complété pour qu'il soit traité
			// et un message de cloture
			cmds = append(cmds,
				func() tea.Msg { return m.Msg },
				func() tea.Msg { return FilledMsg{} },
			)

		case key.Matches(msg, keys.Cancel):
			// ne pas retourner de message rempli simplement fermer la fenêtre modale
			cmds = append(cmds,
				func() tea.Msg { return FilledMsg{} },
			)

		default:
			// transférer aux champs de saisie
			for i := range m.Fields {
				m.Fields[i], cmd = m.Fields[i].Update(msg)
				cmds = append(cmds, cmd)
			}
		}

	default:
		// transférer aux champs de saisie
		for i := range m.Fields {
			m.Fields[i], cmd = m.Fields[i].Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	fields := []string{m.Title}
	for i := range m.Fields {
		fields = append(fields, m.Fields[i].View())
	}

	return lg.JoinVertical(lg.Left,
		fields...,
	)
}
