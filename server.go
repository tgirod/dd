package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"

	"gopkg.in/yaml.v3"
)

// Server représente un serveur sur le Net
type Server struct {
	// Addresse du serveur sur le réseau
	Address string `storm:"id"`

	// ce serveur accepte-t-il des connexions anonymes ?
	Private bool

	// informations affichées lors de la connexion
	Description string

	// niveau de sécurité : plus le chiffre est élevé, plus rapide est la trace
	Security int
}

func (s Server) HasResource() q.Matcher {
	return q.Eq("Server", s.Address)
}

// User représente un compte utilisateur sur un serveur
type User struct {
	ID       int    `storm:"id,increment"`
	Login    string `storm:"index"`
	Server   string `storm:"index"` // le serveur concerné
	Backdoor bool
	Groups   Groups
}

func (u User) Value() any {
	return u.Login
}

func (u User) Desc() string {
	return strings.Join(u.Groups, " ")
}

type Groups []string

func (u User) HasAccess() q.Matcher {
	return q.NewFieldMatcher("Group", u.Groups)
}

// Match permet de vérifier si une donnée est accessible depuis un compte
func (gs Groups) MatchField(v any) (bool, error) {
	group, ok := v.(string)
	if !ok {
		return false, storm.ErrBadType
	}

	// aucun groupe == public
	if group == "" {
		return true, nil
	}

	for _, g := range gs {
		if group == g {
			return true, nil
		}
	}
	return false, nil
}

func (s Server) Users() []User {
	users, err := Find[User](s.HasResource())
	if err != nil {
		panic(err)
	}
	return users
}

// FindUser cherche un compte utilisateur correspondant au login
func (s Server) FindUser(login string) (User, error) {
	return First[User](
		s.HasResource(),
		q.Eq("Login", login),
	)
}

func (s Server) RemoveUser(user User) error {
	return Delete(user)
}

type Link struct {
	ID     int    `storm:"id,increment"`
	Server string `storm:"index"`
	Group  string `storm:"index"`

	// adresse du serveur de destination
	Address string

	// description du lien
	Description string
}

func (l Link) Value() any {
	return l.ID
}

func (l Link) Desc() string {
	return l.Description
}

func (s Server) Links(u User) []Link {
	links, err := Find[Link](
		s.HasResource(),
		u.HasAccess(),
	)
	if err != nil {
		panic(err)
	}
	return links
}

func (s Server) Link(id int, u User) (Link, error) {
	return First[Link](
		s.HasResource(),
		u.HasAccess(),
		q.Eq("ID", id),
	)
}

// Register représente registre mémoire qui peut être modifié pour contrôler quelque chose
type Register struct {
	Server      string `storm:"index"`
	Group       string `storm:"index"`
	ID          int    `storm:"id,increment"`
	Description string
	State       RegisterState   // état actuel
	Options     []RegisterState // valeurs possible
}

type RegisterState string

func (s RegisterState) Value() any {
	return s
}

func (s RegisterState) Desc() string {
	return string(s)
}

func (r Register) Value() any {
	return r.ID
}

func (r Register) Desc() string {
	return fmt.Sprintf("%s : %s", r.Description, r.State)
}

func (s Server) Registers(u User) []Register {
	registers, err := Find[Register](
		s.HasResource(),
		u.HasAccess(),
	)
	if err != nil {
		panic(err)
	}
	return registers
}

func (s Server) Register(id int, u User) (Register, error) {
	return First[Register](
		s.HasResource(),
		u.HasAccess(),
		q.Eq("ID", id),
	)
}

// CreateBackdoor créé une backdoor dans le serveur
func (s Server) CreateBackdoor(identity Identity) (User, error) {
	acc := User{
		Login:    identity.Login,
		Server:   s.Address,
		Backdoor: true,
	}
	return Save(acc)
}

type Post struct {
	Server  string `storm:"index"`
	Group   string `storm:"index"`
	ID      int    `storm:"id,increment"`
	Parent  int    `storm:"index"`
	Date    time.Time
	Author  string
	Subject string
	Content string
}

func (p Post) Value() any {
	return p.ID
}

func (p Post) Desc() string {
	return fmt.Sprintf("%s\t%s\t%s\t", p.Date.Format(time.DateTime), p.Author, p.Subject)
}

func (p Post) Dump() {
	fmt.Printf("--- Dump Post:")
	fmt.Printf("\n Server: [%s]", p.Server)
	fmt.Printf("\n Group: [%s]", p.Group)
	fmt.Printf("\n ID: [%d]", p.ID)
	fmt.Printf("\n Parent: [%d]", p.Parent)
	fmt.Printf("\n Date: [%s]", p.Date)
	fmt.Printf("\n Author: [%s]", p.Author)
	fmt.Printf("\n Subject: [%s]", p.Subject)
	fmt.Printf("\n Content: [%v]", p.Content)
}

type YAMLAnswer struct {
	Date    string
	Author  string
	Content string
}
type YAMLPost struct {
	Server  string
	Group   string
	Date    string
	Author  string
	Subject string
	Content string
	Answers []YAMLAnswer
}
type YAMLMessage struct {
	Date    string
	From    string
	To      string
	Subject string
	Content string
}
type YAMLTransaction struct {
	From    string
	To      string
	Yes     int
	Comment string
}
type YAMLRegister struct {
	Server      string
	Group       string
	Description string
	State       string
	Options     []string
}

// *****************************************************************************
// Load/Save Registry **********************************************************
// *****************************************************************************
func LoadRegistries(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var yamlRegs []YAMLRegister
	err = yaml.Unmarshal(buf, &yamlRegs)
	if err != nil {
		fmt.Printf("FATAL cannot unmarshal: %v\n", err)
		panic(err)
	}
	fmt.Printf("__READ Registres\n")
	fmt.Printf("  read %d regs\n", len(yamlRegs))

	for _, yamlr := range yamlRegs {
		fmt.Printf("__Registry in [%s], desc: %s => <%s>\n", yamlr.Server, yamlr.Description, yamlr.State)

		var options []RegisterState
		for _, rs := range yamlr.Options {
			options = append(options, RegisterState(rs))
		}
		reg := Register{
			Server:      yamlr.Server,
			Group:       yamlr.Group,
			Description: yamlr.Description,
			State:       RegisterState(yamlr.State),
			Options:     options,
		}
		//fmt.Printf("REG %v\n", reg)
		_, err = Save(reg)
		if err != nil {
			fmt.Printf("FATAL cannot save reg: %v\n", err)
		}
		fmt.Printf(" reg saved\n")
	}
}

// *****************************************************************************
// Load/Save Transactions ******************************************************
// *****************************************************************************
func SerializeTransactions() {

	transactions, err := Find[Transaction]()
	if err != nil {
		panic(err)
	}

	for _, t := range transactions {
		fmt.Printf("--- Dump Transaction:")
		fmt.Printf("\n ID: [%d]", t.ID)
		fmt.Printf("\n From: [%s]", t.From)
		fmt.Printf("\n To: [%s]", t.To)
		fmt.Printf("\n YES: [%d]", t.Yes)
		fmt.Printf("\n Comment: [%v]", t.Comment)
	}

	// all transactiouns
	d, err := yaml.Marshal(transactions)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- all transactions:\n%s\n\n", d)
}

func LoadTransactions(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var yamlTransactions []YAMLTransaction
	err = yaml.Unmarshal(buf, &yamlTransactions)
	if err != nil {
		fmt.Printf("FATAL cannot unmarshal: %v\n", err)
		panic(err)
	}
	fmt.Printf("__READ Transactions\n")
	fmt.Printf("  read %d trans\n", len(yamlTransactions))

	for _, yamlt := range yamlTransactions {
		fmt.Printf("__Trans from [%s] to %s: <%d\n", yamlt.From, yamlt.To, yamlt.Yes)

		trans := Transaction{
			From:    yamlt.From,
			To:      yamlt.To,
			Yes:     yamlt.Yes,
			Comment: yamlt.Comment,
		}
		_, err = Save(trans)
		if err != nil {
			fmt.Printf("FATAL cannot save trans: %v\n", err)
		}
		fmt.Printf(" trans saved\n")
	}
}

// *****************************************************************************
// Load/Save Messages **********************************************************
// *****************************************************************************
func SerializeMessages() {

	msges, err := Find[Message]()
	if err != nil {
		panic(err)
	}

	for _, m := range msges {
		fmt.Printf("--- Dump Message:")
		fmt.Printf("\n ID: [%d]", m.ID)
		fmt.Printf("\n From: [%s]", m.From)
		fmt.Printf("\n To: [%s]", m.To)
		fmt.Printf("\n Date: [%s]", m.Date)
		fmt.Printf("\n Subject: [%s]", m.Subject)
		fmt.Printf("\n Content: [%v]", m.Content)
	}

	// all messages
	d, err := yaml.Marshal(msges)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- all messages:\n%s\n\n", d)
}
func LoadMessages(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var yamlMessages []YAMLMessage
	err = yaml.Unmarshal(buf, &yamlMessages)
	if err != nil {
		fmt.Printf("FATAL cannot unmarshal: %v\n", err)
		panic(err)
	}
	fmt.Printf("__READ Messages\n")
	fmt.Printf("  read %d msg\n", len(yamlMessages))

	for _, yamlm := range yamlMessages {
		fmt.Printf("__Msg from [%s] to %s: <ùs>from [%s]\n", yamlm.From, yamlm.To, yamlm.Subject)
		date, err := time.Parse("2006-01-02T15:04:05", yamlm.Date)
		if err != nil {
			fmt.Printf("FATAL cannot parse time: %v\n", err)
			panic(err)
		}
		fmt.Printf("Date %v\n", date)

		msg := Message{
			Date:    date,
			From:    yamlm.From,
			To:      yamlm.To,
			Subject: yamlm.Subject,
			Content: yamlm.Content,
			Opened:  false,
		}
		_, err = Save(msg)
		if err != nil {
			fmt.Printf("FATAL cannot save msg: %v\n", err)
		}
		fmt.Printf(" msg saved\n")
	}
}

// *****************************************************************************
// Load/Save Forum *************************************************************
// *****************************************************************************

// TEST : serialize all Posts to YAML
func SerializePosts(addr string) {
	s, err := FindServer(addr)
	if err != nil {
		panic(err)
	}

	posts, err := Find[Post](
		s.HasResource(),
	)
	if err != nil {
		panic(err)
	}

	// all posts
	d, err := yaml.Marshal(posts)
	if err != nil {
		panic(err)
	}
	nowStr := time.Now().Format("2006-01-02T15:04:05")
	fmt.Printf("--- Date %s\n", nowStr)
	fmt.Printf("--- all posts:\n%s\n\n", d)
}

// TEST Load new post from YAML file
func LoadPosts(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// FIXME TODO make sure we can add post in the server (Forum allowed)
	var yamlPosts []YAMLPost
	err = yaml.Unmarshal(buf, &yamlPosts)
	if err != nil {
		fmt.Printf("FATAL cannot unmarshal: %v\n", err)
		panic(err)
	}
	fmt.Printf("__READ Posts\n")
	fmt.Printf("  read %d posts\n", len(yamlPosts))

	for _, yamlp := range yamlPosts {
		fmt.Printf("__Post to %s:<%s> from [%s]\n", yamlp.Server, yamlp.Subject, yamlp.Author)
		date, err := time.Parse("2006-01-02T15:04:05", yamlp.Date)
		if err != nil {
			fmt.Printf("FATAL cannot parse time: %v\n", err)
			panic(err)
		}
		fmt.Printf("Date %v\n", date)

		post := Post{
			Server:  yamlp.Server,
			Group:   yamlp.Group,
			Date:    date,
			Author:  yamlp.Author,
			Subject: yamlp.Subject,
			Content: yamlp.Content,
		}
		post, err = Save(post)
		if err != nil {
			fmt.Printf("FATAL cannot save post: %v\n", err)
		}
		fmt.Printf(" post saved\n")

		fmt.Printf("Post has %d answers\n", len(yamlp.Answers))
		// Answers
		for _, yamla := range yamlp.Answers {
			fmt.Printf("__Answer from [%s]\n", yamla.Author)
			adate, err := time.Parse("2006-01-02T15:04:05", yamla.Date)
			if err != nil {
				fmt.Printf("FATAL cannot parse time: %v\n", err)
				panic(err)
			}
			fmt.Printf("Date %v\n", adate)

			answer := Post{
				Server:  post.Server,
				Group:   post.Group,
				Parent:  post.ID,
				Date:    adate,
				Author:  yamla.Author,
				Subject: fmt.Sprintf("Re: %s", post.Subject),
				Content: yamla.Content,
			}
			answer, err = Save(answer)
			if err != nil {
				fmt.Printf("FATAL cannot save answer: %v\n", err)
			}
			fmt.Printf(" answer saved\n")
		}
	}
}

func (s Server) Post(id int) (Post, error) {
	return First[Post](
		s.HasResource(),
		q.Eq("ID", id),
	)
}

// Topics liste les posts qui n'ont pas de parent
func (s Server) Topics(u User) []Post {
	posts, err := Find[Post](
		s.HasResource(),
		q.Eq("Parent", 0),
		u.HasAccess(),
	)
	if err != nil {
		panic(err)
	}
	return posts
}

// Replies retourne la liste des réponses à un post
func (s Server) Replies(parent int) []Post {
	posts, err := Find[Post](
		s.HasResource(),
		q.Eq("Parent", parent),
	)
	if err != nil {
		panic(err)
	}

	return posts
}

func concat[T any](slices ...[]T) []T {
	var res []T
	for _, s := range slices {
		res = append(res, s...)
	}
	return res
}

type Thread struct {
	Post
	Replies []Post
}

func (s Server) Thread(p Post) (Thread, error) {

	replies, err := Find[Post](
		s.HasResource(),
		q.Eq("Parent", p.ID),
	)

	thread := Thread{Post: p}
	if err != nil {
		return thread, err
	}

	thread.Replies = append(thread.Replies, replies...)
	return thread, nil
}
