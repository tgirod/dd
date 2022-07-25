package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	//"github.com/lithammer/fuzzysearch/fuzzy"
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	Network []Server
}

func (g Game) FindServer(address string) (*Server, error) {
	for _, server := range g.Network {
		if server.Address == address {
			return &server, nil
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
		_ = ioutil.WriteFile( "network.json", ret, 0644 );
		
	}
}

func (g Game) UnSerialize() {
	content, err := ioutil.ReadFile("network.json")
	if err != nil {
		fmt.Println("Cannot open JSON file")
	}
	err = json.Unmarshal(content, &g)
	if err != nil {
        fmt.Println("Can't deserislize", content)
    }
}
	
