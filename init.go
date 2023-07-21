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
	amathison   = Identity{"amathison", "hai3ja", "Alan Mathison"}         // alan
	mmathison   = Identity{"mmathison", "mie6oo", "Mélody Mathison"}       // mel
	mbellamy    = Identity{"mbellamy", "ahng7e", "Margherita Bellamy"}     // rita
	sbronner    = Identity{"sbronner", "rahk0u", "Sebastian Bronner"}      // styx
	cbellamy    = Identity{"cbellamy", "xoh7sh", "Camélia Bellamy"}        // kapo
	jvillanova  = Identity{"jvillanova", "ay9aef", "Julius Villanova"}     // scalpel
	ecanto      = Identity{"ecanto", "ti3eim", "Eddy Canto"}               // greko
	ejohannesen = Identity{"ejohannesen", "obo4ie", "Edwin Johannesen"}    // jesus
	jbranson    = Identity{"jbranson", "aich8g", "Jonathan Branson"}       // escobar
	jmfright    = Identity{"jmfright", "uruw5g", "John Mac Fright"}        // cageot
	skmihalec   = Identity{"skmihalec", "paeh3l", "Sylvia Kemija Mihalec"} // la fouine
	emartin     = Identity{"emartin", "thooy1", "Eva Martin"}              // eva
	mdubian     = Identity{"mdubian", "iup1ie", "Michael Dubian"}          // fat mike
	cmihalec    = Identity{"cmihalec", "uequ8u", "Carlotta Mihalec"}       // kennedy
	sjohannesen = Identity{"sjohannesen", "aiphu4", "Sabrina Johannesen"}  // savage girl
	rmichu      = Identity{"rmichu", "ool7ch", "Raoul Michu"}              // raoul cool
	rglass      = Identity{"rglass", "ahzae2", "Rupert Glass"}             // green glass
	sglass      = Identity{"sglass", "si6aeb", "Stefie Glass"}             // stefie
	djohannesen = Identity{"djohannesen", "loh1ie", "Daisy Johannesen"}    // chilly daisy
	dbonenfant  = Identity{"dbonenfant", "de4oiv", "Désiré Bonenfant"}     // frère ping
	hproskychev = Identity{"hproskychev", "ooj4an", "Harald Proskychev"}   // papa proxy
	njasinski   = Identity{"njasinski", "eveth3", "Nikole Jasinski"}       // nikki
	sjasinski   = Identity{"sjasinski", "ie7uo2", "Stefan Jasinski"}       // sasquatch
	ffceline    = Identity{"ffceline", "boh6ay", "Franz-Ferdinand Celine"} // celine
	cmills      = Identity{"cmills", "thue1d", "Camélia Mills"}            // cramille
	lseptembre  = Identity{"lseptembre", "cul1ol", "Lilas Septembre"}      // tiger doll
	edubian     = Identity{"edubian", "rooch7", "Eloïse Dubian"}           // sister morphine
	zabasolo    = Identity{"zabasolo", "aipho0", "Zilmir Abasolo"}         // zilmir
	ebranson    = Identity{"ebranson", "rae2ie", "Elisabeth Branson"}      // betty b
	jkievain    = Identity{"jkievain", "nie3oo", "Jordan Kievain"}         // abraham
	fmanson     = Identity{"fmanson", "tiuf0y", "Frédéric Manson"}         // crunch
	rkievain    = Identity{"rkievain", "aso2qu", "Rodolph Kievain"}        // one kick
	pdoberty    = Identity{"pdoberty", "aivei1", "Pete Doberty"}           // jacob
	rwhite      = Identity{"rwhite", "ies2su", "Richard White"}            // oggy
	ajolivet    = Identity{"ajolivet", "quai1a", "Adrien Jolivet"}         // cyrano
	mklebert    = Identity{"mklebert", "eis6ku", "Mickael Klebert"}        // iron mike
	jvazzanna   = Identity{"jvazzanna", "ueth4k", "Joseph Vazzanna"}       // small joe
	jbatista    = Identity{"jbatista", "yah6ae", "Johaquim Batista"}       // joe-rez
	gchang      = Identity{"gchang", "zo1daa", "Georges Chang"}            // georges

	// PNJs
	afrieman = Identity{"afrieman", "far3ik", "Anton Frieman"} // PNJ fan blackwave
	// TODO quelques employé•e•s de la kramps
	// TODO quelques employé•e•s de legba voodoocom
)

// identités virtuelles fournies par Jésus et le FLR
var (
	hope           = Identity{"hope", "goo1uh", "Hope"}
	mel            = Identity{"mel", "quoo1d", "Mel"}
	rocky          = Identity{"rocky", "zuc9eo", "Rocky"}
	rita           = Identity{"rita", "caeb5p", "Rita"}
	styx           = Identity{"styx", "ooz3zu", "Styx"}
	kapo           = Identity{"kapo", "ree4iu", "Kapo"}
	scalpel        = Identity{"scalpel", "uaqu9u", "Scalpel"}
	greko          = Identity{"greko", "boo2ae", "Greko"}
	jesus          = Identity{"jesus", "zodu9s", "Jesus"}
	escobar        = Identity{"escobar", "chach6", "Escobar"}
	cageot         = Identity{"cageot", "quai4a", "Cageot"}
	lafouine       = Identity{"lafouine", "ad0eir", "La Fouine"}
	eva            = Identity{"eva", "heib2i", "Eva"}
	fatmike        = Identity{"fatmike", "ogh2pe", "Fat"}
	kennedy        = Identity{"kennedy", "ri1yak", "Kennedy"}
	savagegirl     = Identity{"savagegirl", "ija5ph", "Savage"}
	raoulcool      = Identity{"raoulcool", "ciex5p", "Raoul Cool"}
	greenglass     = Identity{"greenglass", "ahc0si", "Green Glass"}
	chillydaisy    = Identity{"chillydaisy", "aiv5du", "Chilly Daisy"}
	ping           = Identity{"ping", "neip3i", "Frère Ping"}
	papaproxy      = Identity{"papaproxy", "pu5och", "Papa Proxy"}
	nikki          = Identity{"nikki", "eizoh3", "Nikki"}
	celine         = Identity{"celine", "phie9h", "Céline"}
	cramille       = Identity{"cramille", "ohs6fu", "Cramille"}
	tigerdoll      = Identity{"tigerdoll", "xai8ei", "Tiger Doll"}
	sistermorphine = Identity{"sistermorphine", "agoh9o", "Sister Morphine"}
	zilmir         = Identity{"zilmir", "jo7eiw", "Zilmir"}
	bettyb         = Identity{"bettyb", "aic5er", "Betty B"}
	abraham        = Identity{"abraham", "quook4", "Abraham"}
	crunch         = Identity{"crunch", "dae8at", "Crunch"}
	onekick        = Identity{"onekick", "faefi1", "One Kick"}
	jacob          = Identity{"jacob", "eet2sa", "Jacob"}
	oggy           = Identity{"oggy", "jee9ae", "Oggy"}
	cyrano         = Identity{"cyrano", "bor6oh", "Cyrano"}
	ironmike       = Identity{"ironmike", "eij9mo", "Iron Mike"}
	smalljoe       = Identity{"smalljoe", "ue9jae", "Small Joe"}
	joerez         = Identity{"joerez", "je6rie", "Joe-Rez"}
	jeanne         = Identity{"jeanne", "ech4ee", "Jeanne"}
	paula          = Identity{"paula", "johy2m", "Paula"}
	georges        = Identity{"georges", "ob5jah", "Georges"}
	ringo          = Identity{"ringo", "aic1ar", "Ringo"}
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
