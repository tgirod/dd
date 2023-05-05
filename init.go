package main

import (
	"log"
	"time"
)

const (
	SEC1 = time.Minute * 5
	SEC2 = time.Minute * 3
	SEC3 = time.Minute * 2
	SEC4 = time.Minute * 1
	SEC5 = time.Second * 30
)

func InitNetwork(
	identities []Identity,
) {
	for _, i := range identities {
		if _, err := Save(i); err != nil {
			log.Fatalf("%v : %v\n", i, err)
		}
	}
}

func InitServer(
	s Server,
	accounts []Account,
	links []Link,
	entries []Entry,
	registers []Register,
	posts []Post,
) {
	addr := s.Address
	if addr == "" {
		panic("le serveur n'a pas d'adresse")
	}

	if _, err := Save(s); err != nil {
		log.Fatal(err)
	}

	for _, a := range accounts {
		a.Server = addr
		if _, err := Save(a); err != nil {
			log.Fatalf("%v : %v\n", a, err)
		}
	}
	for _, l := range links {
		l.Server = addr
		if _, err := Save(l); err != nil {
			log.Fatalf("%v : %v\n", l, err)
		}
	}
	for _, e := range entries {
		e.Server = addr
		if _, err := Save(e); err != nil {
			log.Fatalf("%v : %v\n", e, err)
		}
	}
	for _, r := range registers {
		r.Server = addr
		if _, err := Save(r); err != nil {
			log.Fatalf("%v : %v\n", r, err)
		}
	}
	for _, p := range posts {
		p.Server = addr
		if _, err := Save(p); err != nil {
			log.Fatalf("%v : %v\n", p, err)
		}
	}
}

func Reset() {
	db.Drop(Identity{})
	db.Drop(Message{})
	db.Drop(Server{})
	db.Drop(Account{})
	db.Drop(Link{})
	db.Drop(Entry{})
	db.Drop(Register{})
	db.Drop(Post{})
}

var ddDesc = `
 ____  _      _           ____  _     _        _      _
|  _ \(_)_ __| |_ _   _  |  _ \(_)___| |_ _ __(_) ___| |_
| | | | | '__| __| | | | | | | | / __| __| '__| |/ __| __|
| |_| | | |  | |_| |_| | | |_| | \__ \ |_| |  | | (__| |_
|____/|_|_|   \__|\__, | |____/|_|___/\__|_|  |_|\___|\__|
                  |___/
Bienvenue sur le serveur communautaire du Dirty District.

Ce serveur est connecté au Net par le biais d'un accès illégal. Merci de ne pas
faire n'importe quoi.

Tape "index" pour avoir la liste des services fournis par le serveur. Si tu as
besoin d'aide, demande à ton nerd préféré.`

var dd = Server{
	Address:     "dd.local",
	Public:      true,
	Description: ddDesc,
	Scan:        SEC1,
}

var dd22Desc = `
                    _____                            _                       
                   |  ___|                          | |                      
                   | |__ _   _ _ __ ___  _ __   ___ | | ___                  
                   |  __| | | | '__/ _ \| '_ \ / _ \| |/ _ \                 
                   | |__| |_| | | | (_) | |_) | (_) | |  __/                 
                   \____/\__,_|_|  \___/| .__/ \___/|_|\___|                 
                                        | |                                  
                                        |_|                                  
               ______ _     _        _      _     _____  _____               
               |  _  (_)   | |      (_)    | |   / __  \/ __  \              
  ______ ___   | | | |_ ___| |_ _ __ _  ___| |_   ' / /' ' / /'   ___ ______ 
 |______/ _ \  | | | | / __| __| '__| |/ __| __|   / /    / /    / _ \______|
       | (_) | | |/ /| \__ \ |_| |  | | (__| |_  ./ /___./ /___ | (_) |      
        \___/  |___/ |_|___/\__|_|  |_|\___|\__| \_____/\_____/  \___/       
                                                                             

           Bienvenue sur le serveur public du District 22 d'Europole.
           Un noeud du plus grand fournisseur d'accès de Méga-Europe.`

var d22 = Server{
	Address:     "d22.eu",
	Public:      true,
	Description: dd22Desc,
	Scan:        SEC2,
}

func Init() {
	Reset()

	InitNetwork(
		[]Identity{
			{
				Login:    "jesus",
				Password: "roxor",
				Name:     "Jesus",
				Yes:      100,
			},
			{
				Login:    "crunch",
				Password: "hack",
				Name:     "Crunch",
				Yes:      0,
			},
		},
	)

	InitServer(dd,
		[]Account{
			{Login: "jesus", Admin: true},
		},
		[]Link{
			{Address: d22.Address, Desc: "serveur public du District 22"},
		},
		[]Entry{
			{ID: "bluemars", Keywords: []string{"boisson"}, Title: "blue mars", Content: "cocktail"},
		},
		[]Register{
			{Description: "machine à café", State: "on", Options: []string{"on", "off", "overdrive"}},
		},
		[]Post{},
	)

	InitServer(d22,
		[]Account{},
		[]Link{},
		[]Entry{},
		[]Register{},
		[]Post{},
	)
}
