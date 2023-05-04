package main

import (
	"encoding/base64"
	"math/rand"
)

type Identity struct {
	Login    string `storm:"id"`
	Password string
	Name     string
	Yes      int
	Messages []Message
}

type Message struct {
	ID        int    `storm:"id,increment"`
	Recipient string `storm:"index"` // expéditeur
	Sender    string // destinataire
	Subject   string // titre du message
	Content   string // contenu du message
	Opened    bool   `storm:"index"` // pas encore lu
}

func MessageSend(m Message) (Message, error) {
	// trouver le destinataire
	_, err := FindIdentity(m.Recipient)
	if err != nil {
		return m, err
	}

	return Save(m)
}

func Pay(from, to string, amount int) error {
	var src, dst Identity
	var err error

	if src, err = FindIdentity(from); err != nil {
		return err
	}

	if dst, err = FindIdentity(to); err != nil {
		return err
	}

	if src.Yes < amount {
		return errLowCredit
	}

	if amount < 0 {
		return errNegativeAmount
	}

	src.Yes = src.Yes - amount
	dst.Yes = dst.Yes + amount

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
		Yes:      0,
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

func FindIdentity(login string) (Identity, error) {
	return One[Identity]("Login", login)
}

func FindServer(address string) (Server, error) {
	return One[Server]("Address", address)
}
