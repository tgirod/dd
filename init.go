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
	amathison     = Identity{"amathison", "hai3ja", "Alan Mathison", true}            // alan
	mmathison     = Identity{"mmathison", "mie6oo", "Mélody Mathison", true}          // mel
	mbellamy      = Identity{"mbellamy", "ahng7e", "Margherita Bellamy", true}        // rita
	sbronner      = Identity{"sbronner", "rahk0u", "Sebastian Bronner", true}         // styx
	cbellamy      = Identity{"cbellamy", "xoh7sh", "Camélia Bellamy", true}           // kapo
	jvillanova    = Identity{"jvillanova", "ay9aef", "Julius Villanova", true}        // scalpel
	ecanto        = Identity{"ecanto", "ti3eim", "Eddy Canto", true}                  // greko
	ejohannesen   = Identity{"ejohannesen", "obo4ie", "Edwin Johannesen", true}       // jesus
	jbranson      = Identity{"jbranson", "aich8g", "Jonathan Branson", true}          // escobar
	jmfright      = Identity{"jmfright", "uruw5g", "John Mac Fright", true}           // cageot
	skmihalec     = Identity{"skmihalec", "paeh3l", "Sylvia Kemija Mihalec", true}    // la fouine
	emartin       = Identity{"emartin", "thooy1", "Eva Martin", true}                 // eva
	mdubian       = Identity{"mdubian", "iup1ie", "Michael Dubian", true}             // fat mike
	cmihalec      = Identity{"cmihalec", "uequ8u", "Carlotta Mihalec", true}          // kennedy
	sjohannesen   = Identity{"sjohannesen", "aiphu4", "Sabrina Johannesen", true}     // savage girl
	rmichu        = Identity{"rmichu", "ool7ch", "Raoul Michu", true}                 // raoul cool
	rglass        = Identity{"rglass", "ahzae2", "Rupert Glass", true}                // green glass
	sglass        = Identity{"sglass", "si6aeb", "Stefie Glass", true}                // stefie
	djohannesen   = Identity{"djohannesen", "loh1ie", "Daisy Johannesen", true}       // chilly daisy
	dbonenfant    = Identity{"dbonenfant", "de4oiv", "Désiré Bonenfant", true}        // frère ping
	hproskychev   = Identity{"hproskychev", "ooj4an", "Harald Proskychev", true}      // papa proxy
	njasinski     = Identity{"njasinski", "eveth3", "Nikole Jasinski", true}          // nikki
	sjasinski     = Identity{"sjasinski", "ie7uo2", "Stefan Jasinski", true}          // sasquatch
	ffceline      = Identity{"ffceline", "boh6ay", "Franz-Ferdinand Celine", true}    // celine
	cmills        = Identity{"cmills", "thue1d", "Camélia Mills", true}               // cramille
	lseptembre    = Identity{"lseptembre", "cul1ol", "Lilas Septembre", true}         // tiger doll
	edubian       = Identity{"edubian", "rooch7", "Eloïse Dubian", true}              // sister morphine
	zabasolo      = Identity{"zabasolo", "aipho0", "Zilmir Abasolo", true}            // zilmir
	ebranson      = Identity{"ebranson", "rae2ie", "Elisabeth Branson", true}         // betty b
	jkievain      = Identity{"jkievain", "nie3oo", "Jordan Kievain", true}            // abraham
	fmanson       = Identity{"fmanson", "tiuf0y", "Frédéric Manson", true}            // crunch
	rkievain      = Identity{"rkievain", "aso2qu", "Rodolph Kievain", true}           // one kick
	pdoberty      = Identity{"pdoberty", "aivei1", "Pete Doberty", true}              // jacob
	rwhite        = Identity{"rwhite", "ies2su", "Richard White", true}               // oggy
	ajolivet      = Identity{"ajolivet", "quai1a", "Adrien Jolivet", true}            // cyrano
	mklebert      = Identity{"mklebert", "eis6ku", "Mickael Klebert", true}           // iron mike
	jvazzanna     = Identity{"jvazzanna", "ueth4k", "Joseph Vazzanna", true}          // small joe
	jbatista      = Identity{"jbatista", "yah6ae", "Johaquim Batista", true}          // joe-rez
	gsuleymanoglu = Identity{"gsuleymanoglu", "zo1daa", "Georges Suleymanoglu", true} // georges

	// PNJs
	afrieman = Identity{"afrieman", "far3ik", "Anton Frieman", true} // PNJ fan blackwave
	// TODO quelques employé•e•s de la kramps
	// TODO quelques employé•e•s de legba voodoocom
)

// identités virtuelles fournies par Jésus et le FLR
var (
	hope           = Identity{"hope", "011011011011", "Hope", false}
	mel            = Identity{"mel", "hope4ever", "Mel", false}
	rocky          = Identity{"rocky", "pourquoi", "Rocky", false}
	rita           = Identity{"rita", "1wantM0re", "Rita", false}
	styx           = Identity{"styx", "plastobeton", "Styx", false}
	kapo           = Identity{"kapo", "touspour1", "Kapo", false}
	scalpel        = Identity{"scalpel", "m3dic!!", "Scalpel", false}
	greko          = Identity{"greko", "FuckY00", "Greko", false}
	jesus          = Identity{"jesus", "ZtqCtdlb42", "Jesus", false}
	escobar        = Identity{"escobar", "M0n3y++", "Escobar", false}
	cageot         = Identity{"cageot", "fr33dom", "Cageot", false}
	lafouine       = Identity{"lafouine", "cplvfh3h3", "La Fouine", false}
	eva            = Identity{"eva", "n3verAgain", "Eva", false}
	fatmike        = Identity{"fatmike", "tamereenshort", "Fat", false}
	kennedy        = Identity{"kennedy", "jensaisrien", "Kennedy", false}
	savagegirl     = Identity{"savagegirl", "zeStyle!", "Savage", false}
	raoulcool      = Identity{"raoulcool", "barb3rKing", "Raoul Cool", false}
	greenglass     = Identity{"greenglass", "il0veM0m", "Green Glass", false}
	chillydaisy    = Identity{"chillydaisy", "rb0nesQueen", "Chilly Daisy", false}
	ping           = Identity{"ping", "n0tp0ng!!", "Frère Ping", false}
	papaproxy      = Identity{"papaproxy", "4ragnar!", "Papa Proxy", false}
	nikki          = Identity{"nikki", "3141593", "Nikki", false}
	celine         = Identity{"celine", "f0rg3tme", "Céline", false}
	cramille       = Identity{"cramille", "onsenbalance", "Cramille", false}
	tigerdoll      = Identity{"tigerdoll", "karateGirl", "Tiger Doll", false}
	sistermorphine = Identity{"sistermorphine", "Icanfly", "Sister Morphine", false}
	zilmir         = Identity{"zilmir", "cl3v3r", "Zilmir", false}
	bettyb         = Identity{"bettyb", "ZeK0nsol", "Betty B", false}
	abraham        = Identity{"abraham", "killerSolo", "Abraham", false}
	crunch         = Identity{"crunch", "umdpcfpnp3o", "Crunch", false}
	onekick        = Identity{"onekick", "faitchier", "One Kick", false}
	jacob          = Identity{"jacob", "el01se", "Jacob", false}
	oggy           = Identity{"oggy", "y0pasC0ul", "Oggy", false}
	cyrano         = Identity{"cyrano", "rbNbOne", "Cyrano", false}
	ironmike       = Identity{"ironmike", "deadlymike", "Iron Mike", false}
	smallbob       = Identity{"smallbob", "smallbob", "Small Bob", false}
	joerez         = Identity{"joerez", "mfuck3r", "Joe-Rez", false}
	jeanne         = Identity{"jeanne", "j", "Jeanne", false}
	paula          = Identity{"paula", "mdpH@rd", "Paula", false}
	georges        = Identity{"georges", "devine!", "Georges", false}
	ringo          = Identity{"ringo", "l@cherien!", "Ringo", false}
)

type InfoPlayer struct {
	Perso  string
	IdCorp *Identity
	IdVirt Identity
	Wanted bool
	Known  bool
}

var allPlayers = []InfoPlayer{
	{"Hope", nil, hope, false, false},
	{"Mel", &mmathison, mel, false, true},
	{"Rocky", nil, rocky, false, false},
	{"Rita", &mbellamy, rita, false, true},
	{"Styx", &sbronner, styx, false, true},
	{"Kapo", &cbellamy, kapo, false, true},
	{"Scalpel", &jvillanova, scalpel, false, true},
	{"Greko", &ecanto, greko, false, true},
	{"jesus", &ejohannesen, jesus, false, true},
	{"Escobar", &jbranson, escobar, false, true},
	{"Cageot", &jmfright, cageot, false, true},
	{"La Fouine", &skmihalec, lafouine, false, true},
	{"Eva", &emartin, eva, false, true},
	{"Fat Mike", &mdubian, fatmike, false, true},
	{"Kenndy", &cmihalec, kennedy, false, true},
	{"Savage Girl", &sjohannesen, savagegirl, false, true},
	{"Raoul Cool", &rmichu, raoulcool, false, true},
	{"Green Glass", &rglass, greenglass, false, true},
	{"Chilly Daisy", &djohannesen, chillydaisy, false, true},
	{"Frère Ping", &dbonenfant, ping, false, true},
	{"Papa Proxy", &hproskychev, papaproxy, false, true},
	{"Nikki", &njasinski, nikki, false, true},
	{"Céline", &ffceline, celine, false, true},
	{"Cramille", &cmills, cramille, false, true},
	{"Tiger Doll", &lseptembre, tigerdoll, false, true},
	{"Sister Morphine", &edubian, sistermorphine, false, true},
	{"Zilmir", &zabasolo, zilmir, false, true},
	{"Betty B", &ebranson, bettyb, false, true},
	{"Abraham", &jkievain, abraham, false, true},
	{"Crunch", &fmanson, crunch, false, true},
	{"One Kick", &rkievain, onekick, false, true},
	{"Jacob", &pdoberty, jacob, false, true},
	{"Oggy", &rwhite, oggy, false, true},
	{"Iron Mike", &mklebert, ironmike, false, true},
	{"Joe-Rez", &jbatista, joerez, false, true},
	{"Cyrano", &ajolivet, cyrano, false, true},
	{"Small Bob", &jvazzanna, smallbob, false, true},
	{"Jeanne", nil, jeanne, false, true},
	{"Ringo", nil, ringo, false, true},
	{"Georges", &gsuleymanoglu, georges, false, false},
	{"Paula", nil, paula, false, true},
}

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
