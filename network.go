package main

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
)

type Identity struct {
	Login    string `storm:"id"`
	Password string
	Name     string
}

type Transaction struct {
	ID      int    `storm:"id,increment"`
	From    string `storm:"index"`
	To      string `storm:"index"`
	Yes     int
	Comment string
}

type Message struct {
	ID      int       `storm:"id,increment"`
	From    string    `storm:"index"` // destinataire
	To      string    `storm:"index"` // expéditeur
	Date    time.Time `storm:"index"` // date de transmission
	Subject string    // titre du message
	Content string    // contenu du message
	Opened  bool      `storm:"index"` // pas encore lu
}

func (i Identity) Messages() []Message {
	query := Query(
		q.Or(
			q.Eq("From", i.Login),
			q.Eq("To", i.Login),
		),
	).OrderBy("Date").Reverse()
	var messages []Message
	err := query.Find(&messages)
	if err != nil && err != storm.ErrNotFound {
		panic(err)
	}
	return messages
}

func (i Identity) Message(id int) (Message, error) {
	return First[Message](
		q.Eq("ID", id),
		q.Or(
			q.Eq("From", i.Login),
			q.Eq("To", i.Login),
		),
	)
}

func (i Identity) Send(to, subject, content string) (Message, error) {
	// trouver le destinataire
	_, err := FindIdentity(to)
	if err != nil {
		return Message{}, err
	}

	return Save(Message{
		From:    i.Login,
		To:      to,
		Date:    time.Now(),
		Subject: subject,
		Content: content,
		Opened:  false,
	})
}

func (i Identity) Transactions() ([]Transaction, error) {
	return Find[Transaction](
		q.Or(
			q.Eq("From", i.Login),
			q.Eq("To", i.Login),
		),
	)
}

func (i Identity) Balance() (int, error) {
	transactions, err := i.Transactions()
	if err != nil {
		return 0, err
	}

	bal := 0
	for _, t := range transactions {
		if t.From == i.Login {
			bal -= t.Yes
		}
		if t.To == i.Login {
			bal += t.Yes
		}
	}

	return bal, nil
}

func Pay(from, to string, amount int) error {
	var src Identity
	var err error

	if src, err = FindIdentity(from); err != nil {
		return err
	}

	if _, err = FindIdentity(to); err != nil {
		return err
	}

	if amount < 0 {
		return errNegativeAmount
	}

	bal, err := src.Balance()
	if err != nil {
		return err
	}

	if bal < amount {
		return errLowCredit
	}

	tx := Transaction{
		From:    from,
		To:      to,
		Yes:     amount,
		Comment: "",
	}

	if tx, err = Save(tx); err != nil {
		return err
	}

	return nil
}

func randomString() string {
	data := make([]byte, 3)
	rand.Read(data)
	return base64.RawStdEncoding.EncodeToString(data)
}

func CreateRandomIdentity() (Identity, error) {
	login := randomString()
	password := randomString()
	id := Identity{
		Login:    login,
		Password: password,
		Name:     "",
	}

	return Save(id)
}

func RemoveIdentity(identity Identity) error {
	return Delete(identity)
}

func CheckIdentity(login, password string) (Identity, error) {
	identity, err := FindIdentity(login)
	if err != nil {
		return identity, err
	}

	if identity.Password != password {
		return identity, errInvalidIdentity
	}

	return identity, nil
}

func Identities() ([]Identity, error) {
	return Find[Identity]()
}

func FindIdentity(login string) (Identity, error) {
	return One[Identity]("Login", login)
}

func FindServer(address string) (Server, error) {
	return One[Server]("Address", address)
}
