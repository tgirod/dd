package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	//"github.com/lithammer/fuzzysearch/fuzzy"
)

// Network contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Network struct {
	Servers    []Server
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
	Opened    bool   // pas encore lu
}

func (n *Network) MessageSend(m Message) error {
	// trouver le destinataire
	recipient, err := n.FindIdentity(m.Recipient)
	if err != nil {
		return err
	}

	recipient.Messages = append(recipient.Messages, m)
	return nil
}

func (n *Network) Pay(from, to string, amount int) error {
	var src, dst *Identity
	var err error

	if src, err = n.FindIdentity(from); err != nil {
		return errIdentityNotFound
	}

	if dst, err = n.FindIdentity(to); err != nil {
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

func (n *Network) CreateRandomIdentity() Identity {
	login := randomString()
	password := randomString()
	id := Identity{
		Login:    login,
		Password: password,
		Name:     "",
		Yes:      0,
	}
	n.Identities = append(n.Identities, id)
	return id
}

func (n *Network) RemoveIdentity(login string) {

}

func (n *Network) CheckIdentity(login, password string) (*Identity, error) {
	for i, id := range n.Identities {
		if id.Login == login && id.Password == password {
			return &n.Identities[i], nil
		}
	}
	return nil, errInvalidIdentity
}

func (n *Network) FindIdentity(login string) (*Identity, error) {
	for i, identity := range n.Identities {
		if identity.Login == login {
			return &n.Identities[i], nil
		}
	}

	return nil, fmt.Errorf("%s : %w", login, errIdentityNotFound)
}

func (n *Network) FindServer(address string) (*Server, error) {
	for i, server := range n.Servers {
		if server.Address == address {
			return &n.Servers[i], nil
		}
	}
	return nil, fmt.Errorf("%s : %w", address, errServerNotFound)
}

func (n *Network) Serialize() {
	ret, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		fmt.Println(err)
	} else {

		//What you get is a byte array, which needs to be converted into a string
		//fmt.Println(string(ret))

		// Write byte array to file
		_ = ioutil.WriteFile("network.json", ret, 0644)

	}
}

func (n *Network) UnSerialize(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot open JSON file")
	}
	err = json.Unmarshal(content, n)
	if err != nil {
		fmt.Println("Can't deserislize", content)
	}
}
