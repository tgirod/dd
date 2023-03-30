package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	src.Yes = src.Yes - amount
	dst.Yes = dst.Yes + amount

	return nil
}

func (g *Game) CheckIdentity(login, password string) error {
	for _, i := range g.Identities {
		if i.Login == login && i.Password == password {
			return nil
		}
	}
	return errInvalidIdentity
}

func (g Game) FindIdentity(login string) (*Identity, error) {
	for i, identity := range g.Identities {
		if identity.Login == login {
			return &g.Identities[i], nil
		}
	}

	return nil, errIdentityNotFound
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
