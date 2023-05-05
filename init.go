package main

import (
	"time"
)

const (
	SEC1 = time.Minute * 5
	SEC2 = time.Minute * 3
	SEC3 = time.Minute * 2
	SEC4 = time.Minute * 1
	SEC5 = time.Second * 30
)

var dd = Server{
	Address:     "dd.local",
	Public:      true,
	Description: ddDesc,
	Scan:        SEC1,
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
besoin d'aide, demande à ton nerd préféré.
`

var d22 = Server{
	Address:     "d22.eu",
	Public:      true,
	Description: dd22Desc,
	Scan:        SEC2,
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
           Un noeud du plus grand fournisseur d'accès de Méga-Europe. 
`

func Init() {
	db.Drop(Identity{})
	db.Drop(Message{})
	db.Drop(Server{})
	db.Drop(Account{})
	db.Drop(Link{})
	db.Drop(Entry{})
	db.Drop(Register{})
	db.Drop(Post{})

	Save(dd)
	Save(Account{Login: "jesus", Server: dd.Address, Admin: true, Backdoor: false})
	Save(Link{Server: dd.Address, Address: d22.Address, Desc: "serveur public du District 22"})
	Save(Register{Server: dd.Address, Description: "machine à café", State: "on", Options: []string{"on", "off", "overdrive"}})
	Save(Entry{
		Server:   dd.Address,
		ID:       "bluemars",
		Keywords: []string{"boisson"},
		Title:    "blue mars",
		Content:  "cocktail",
	})
}
