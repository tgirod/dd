package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	//"github.com/lithammer/fuzzysearch/fuzzy"
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	Network    []Server
	Identities []Identity
}

type Identity struct {
	Login    string
	Password string
	Name     string
	Yes      int
	Messages []Message
}

type Message struct {
	Recipient string // expéditeur
	Sender    string // destinataire
	Subject   string // titre du message
	Content   string // contenu du message
	Unread    bool   // pas encore lu
}

func (g *Game) MessageSend(m Message) error {
	// trouver le destinataire
	recipient, err := g.FindIdentity(m.Recipient)
	if err != nil {
		return err
	}

	recipient.Messages = append(recipient.Messages, m)
	return nil
}

func (g *Game) Pay(from, to string, amount int) error {
	var src, dst *Identity
	var err error

	if src, err = g.FindIdentity(from); err != nil {
		return errIdentityNotFound
	}

	if dst, err = g.FindIdentity(to); err != nil {
		return errIdentityNotFound
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

func (g *Game) CreateRandomIdentity() Identity {
	login := randomString()
	password := randomString()
	id := Identity{
		Login:    login,
		Password: password,
		Name:     "",
		Yes:      0,
	}
	g.Identities = append(g.Identities, id)
	return id
}

func (g *Game) RemoveIdentity(login string) {

}

func (g *Game) CheckIdentity(login, password string) (*Identity, error) {
	for i, id := range g.Identities {
		if id.Login == login && id.Password == password {
			return &g.Identities[i], nil
		}
	}
	return nil, errInvalidIdentity
}

func (g Game) FindIdentity(login string) (*Identity, error) {
	for i, identity := range g.Identities {
		if identity.Login == login {
			return &g.Identities[i], nil
		}
	}

	return nil, fmt.Errorf("%s : %w", login, errIdentityNotFound)
}

func (g Game) FindServer(address string) (*Server, error) {
	for i, server := range g.Network {
		if server.Address == address {
			return &g.Network[i], nil
		}
	}
	return nil, fmt.Errorf("%s : %w", address, errServerNotFound)
}

func (g Game) Serialize() {
	ret, err := json.MarshalIndent(g, "", " ")
	if err != nil {
		fmt.Println(err)
	} else {

		//What you get is a byte array, which needs to be converted into a string
		//fmt.Println(string(ret))

		// Write byte array to file
		_ = ioutil.WriteFile("network.json", ret, 0644)

	}
}

func (g Game) UnSerialize(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot open JSON file")
	}
	err = json.Unmarshal(content, &g)
	if err != nil {
		fmt.Println("Can't deserislize", content)
	}
}
