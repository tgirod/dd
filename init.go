package main

import (
	"log"
)

const (
	SEC1 int = iota
	SEC2
	SEC3
	SEC4
	SEC5
)

func InitNetwork(
	identities []Identity,
	transactions []Transaction,
) {
	log.Println("identités")
	for _, i := range identities {
		log.Println("\t", i.Login)
		if _, err := Save(i); err != nil {
			log.Fatalf("%v : %v\n", i, err)
		}
	}

	log.Println("transactions")
	for _, t := range transactions {
		log.Println("\t", t.From, t.To, t.Yes)
		if _, err := Save(t); err != nil {
			log.Fatalf("%v : %v\n", t, err)
		}
	}
}

func InitServer(
	s Server,
	users []User,
	links []Link,
	registers []Register,
	posts []Post,
) {
	addr := s.Address
	if addr == "" {
		panic("le serveur n'a pas d'adresse")
	}

	log.Println("server", s.Address)
	if _, err := Save(s); err != nil {
		log.Fatal(err)
	}

	log.Println("users")
	for _, a := range users {
		log.Println("\t", a.Login)
		a.Server = addr
		if _, err := Save(a); err != nil {
			log.Fatalf("%v : %v\n", a, err)
		}
	}
	log.Println("links")
	for _, l := range links {
		log.Println("\t", l.Address)
		l.Server = addr
		if _, err := Save(l); err != nil {
			log.Fatalf("%v : %v\n", l, err)
		}
	}
	log.Println("registers")
	for _, r := range registers {
		log.Println("\t", r.Description)
		r.Server = addr
		if _, err := Save(r); err != nil {
			log.Fatalf("%v : %v\n", r, err)
		}
	}
	log.Println("posts")
	for _, p := range posts {
		log.Println("\t", p.Subject)
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
	db.Drop(User{})
	db.Drop(Link{})
	db.Drop(Register{})
	db.Drop(Post{})
	db.Drop(Transaction{})
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
	Description: ddDesc,
	Security:    SEC1,
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
	Description: dd22Desc,
	Security:    SEC3,
}

// identités corpo recopiées depuis l'ancienne version
var (
	amathison   = Identity{"amathison", "hai3ja", "Alan Mathison"}
	mmathison   = Identity{"mmathison", "mie6oo", "Mélody Mathison"}
	jdoe7624    = Identity{"jdoe7624", "ayuu2e", "John Doe 7624"}
	mbellamy    = Identity{"mbellamy", "ahng7e", "Margherita Bellamy"}
	sbronner    = Identity{"sbronner", "rahk0u", "Sebastian Bronner"}
	cbellamy    = Identity{"cbellamy", "xoh7sh", "Camélia Bellamy"}
	jvillanova  = Identity{"jvillanova", "ay9aef", "Julius Villanova"}
	ecanto      = Identity{"ecanto", "ti3eim", "Eddy Canto"}
	ejohannesen = Identity{"ejohannesen", "obo4ie", "Edwin Johannesen"}
	jbranson    = Identity{"jbranson", "aich8g", "Jonathan Branson"}
	jmfright    = Identity{"jmfright", "uruw5g", "John Mac Fright"}
	skmihalec   = Identity{"skmihalec", "paeh3l", "Sylvia Kemija Mihalec"}
	emartin     = Identity{"emartin", "thooy1", "Eva Martin"}
	mdubian     = Identity{"mdubian", "iup1ie", "Michael Dubian"}
	cmihalec    = Identity{"cmihalec", "uequ8u", "Carlotta Mihalec"}
	jdoe        = Identity{"jdoe", "aiphu4", "Jane Doe 2645"}
	rmichu      = Identity{"rmichu", "ool7ch", "Raoul Michu"}
	rglass      = Identity{"rglass", "ahzae2", "Rupert Glass"}
	sglass      = Identity{"sglass", "si6aeb", "Steffie Glass"}
	djohannesen = Identity{"djohannesen", "loh1ie", "Daisy Johannesen"}
	dbonenfant  = Identity{"dbonenfant", "de4oiv", "Désiré Bonenfant"}
	hproskychev = Identity{"hproskychev", "ooj4an", "Harald Proskychev"}
	njasinski   = Identity{"njasinski", "eveth3", "Nikole Jasinski"}
	sjasinski   = Identity{"sjasinski", "ie7uo2", "Stefan Jasinski"}
	ffceline    = Identity{"ffceline", "boh6ay", "Franz-Ferdinand Celine"}
	cmills      = Identity{"cmills", "thue1d", "Camélia Mills"}
	lseptembre  = Identity{"lseptembre", "cul1ol", "Lilas Septembre"}
	edubian     = Identity{"edubian", "rooch7", "Eloïse Dubian"}
	zabasolo    = Identity{"zabasolo", "aipho0", "Zilmir Abasolo"}
	ebranson    = Identity{"ebranson", "rae2ie", "Elisabeth Branson"}
	jkievain    = Identity{"jkievain", "nie3oo", "Jordan Kievain"}
	fmanson     = Identity{"fmanson", "tiuf0y", "Frédéric Manson"}
	rkievain    = Identity{"rkievain", "aso2qu", "Rodolph Kievain"}
	pdoberty    = Identity{"pdoberty", "aivei1", "Pete Doberty"}
	ajolivet    = Identity{"ajolivet", "quai1a", "Adrien Jolivet"}
	jvazzanna   = Identity{"jvazzanna", "ueth4k", "Joseph Vazzanna"}
	mklebert    = Identity{"mklebert", "eis6ku", "Mickael Klebert"}
	pjolivet    = Identity{"pjolivet", "foh3ub", "Paula Jolivet"}
	rjolivet    = Identity{"rjolivet", "ee9gee", "Ringo Jolivet"}
	gchang      = Identity{"gchang", "zo1daa", "Georges Chang"}
	jkolinsky   = Identity{"jkolinsky", "oyu8zi", "Jeanne Kolinsky"}
	rwhite      = Identity{"rwhite", "ies2su", "Richard White"}
	afrieman    = Identity{"afrieman", "far3ik", "Anton Frieman"}
)

func Init() {
	log.Println("réinitialisation de la base de données")

	Reset()

	InitNetwork(
		[]Identity{
			{
				Login:    "jesus",
				Password: "roxor",
				Name:     "Jesus",
			},
			{
				Login:    "crunch",
				Password: "hack",
				Name:     "Crunch",
			},
			{
				Login:    "admin",
				Password: "beurk",
				Name:     "Admin",
			},
		},
		[]Transaction{
			{
				To:      "jesus",
				Yes:     100,
				Comment: "report du solde",
			},
			{
				To:      "admin",
				Yes:     1337,
				Comment: "report du solde",
			},
		},
	)

	InitServer(dd,
		[]User{
			{
				Login:    "jesus",
				Server:   "",
				Backdoor: false,
				Groups:   []string{"admin", "h4ck3r"},
			},
			{
				Login:    "crunch",
				Server:   "",
				Backdoor: false,
				Groups:   []string{"h4ck3r"},
			},
		},
		[]Link{
			{Address: d22.Address, Description: "serveur public du District 22"},
		},
		[]Register{
			{Description: "machine à café", State: "on", Options: []RegisterState{"on", "off", "overdrive"}},
		},
		[]Post{},
	)

	InitServer(d22,
		[]User{
			{
				Login:    "jesus",
				Backdoor: false,
			},
		},
		[]Link{},
		[]Register{},
		[]Post{},
	)
}
