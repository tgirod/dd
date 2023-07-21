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
// TODO remplacer les mdp
var (
	amathison   = Identity{"amathison", "GGCGGTAGCCCCTCTCGAGC", "Alan Mathison"}
	mmathison   = Identity{"mmathison", "GGCCAAAGCTCCTTCGGAGC", "Mélody Mathison"}
	jdoe7624    = Identity{"jdoe7624", "CCGCGCAGAATCATAGCTGT", "John Doe 7624"}
	mbellamy    = Identity{"mbellamy", "CAAAGTTCTAGGCATAGGGA", "Margherita Bellamy"}
	sbronner    = Identity{"sbronner", "TTAGCTCGATATCCTAACCC", "Sebastian Bronner"}
	cbellamy    = Identity{"cbellamy", "GAACTGCTTTAGTTGACGGA", "Camélia Bellamy"}
	jvillanova  = Identity{"jvillanova", "TGAAAGAGACATGATGCCTT", "Julius Villanova"}
	ecanto      = Identity{"ecanto", "TCTGAGGTTTATTGATTTCG", "Eddy Canto"}
	ejohannesen = Identity{"ejohannesen", "TTCGGGATTACTGCGTGCTG", "Edwin Johannesen"}
	jbranson    = Identity{"jbranson", "GGAGGACACCCCAAACGCAT", "Jonathan Branson"}
	jmfright    = Identity{"jmfright", "GCCCTTGTCATGTACTTAGT", "John Mac Fright"}
	skmihalec   = Identity{"skmihalec", "CTGTCACCCAATCTACAGCG", "Sylvia Kemija Mihalec"}
	emartin     = Identity{"emartin", "CTGTTGTAGTGACATGTTTC", "Eva Martin"}
	mdubian     = Identity{"mdubian", "AACCTTGGGCACGGTCGGTA", "Michael Dubian"}
	cmihalec    = Identity{"cmihalec", "CCCGCGGGCAAAGCTGACAG", "Carlotta Mihalec"}
	jdoe        = Identity{"jdoe", "GGGTCTATAGGTCAAACGGT", "Jane Doe 2645"}
	rmichu      = Identity{"rmichu", "GTCACAAGGTTGTTTAATGG", "Raoul Michu"}
	rglass      = Identity{"rglass", "ATGCCTACCTCCAATGATTA", "Rupert Glass"}
	sglass      = Identity{"sglass", "ATCGCTACGTCCATAGACTA", "Steffie Glass"}
	djohannesen = Identity{"djohannesen", "CGGGAGACACGTTCAGTCTT", "Daisy Johannesen"}
	dbonenfant  = Identity{"dbonenfant", "GCATGGCCGAATTCCTCATT", "Désiré Bonenfant"}
	hproskychev = Identity{"hproskychev", "CGATTTGTATTGGATACGGA", "Harald Proskychev"}
	njasinski   = Identity{"njasinski", "ACGAACCTAGAGCCGCACGC", "Nikole Jasinski"}
	sjasinski   = Identity{"sjasinski", "ACGAGTAGAGATGTACACGC", "Stefan Jasinski"}
	ffceline    = Identity{"ffceline", "CGCTCCCATTTCATGTCAGC", "Franz-Ferdinand Celine"}
	cmills      = Identity{"cmills", "TTTGGGAGAAGCTTATGCAC", "Camélia Mills"}
	lseptembre  = Identity{"lseptembre", "ATATGTTGAGCGTAAAGGCG", "Lilas Septembre"}
	edubian     = Identity{"edubian", "CCATCCGGCGGACCTTATGC", "Eloïse Dubian"}
	zabasolo    = Identity{"zabasolo", "GACGGGATACCTACTCTCGA", "Zilmir Abasolo"}
	ebranson    = Identity{"ebranson", "ATTCCGACTCAGGGTACCGG", "Elisabeth Branson"}
	jkievain    = Identity{"jkievain", "TGGCGTCTCTAATTCTTGCC", "Jordan Kievain"}
	fmanson     = Identity{"fmanson", "TTCAAGCTGAATATGAAAGG", "Frédéric Manson"}
	rkievain    = Identity{"rkievain", "GTCAAATCTGAGACTCTTGC", "Rodolph Kievain"}
	pdoberty    = Identity{"pdoberty", "TGAAAGAGACAGTATGCCGT", "Pete Doberty"}
	ajolivet    = Identity{"ajolivet", "TTCGACTGAATGTTTGATGT", "Adrien Jolivet"}
	jvazzanna   = Identity{"jvazzanna", "TATCGACGCACGGGACTTGG", "Joseph Vazzanna"}
	mklebert    = Identity{"mklebert", "CGAGAAATGACAGAGTTGTA", "Mickael Klebert"}
	pjolivet    = Identity{"pjolivet", "GGGTGATCTGTTGCCCCCTG", "Paula Jolivet"}
	rjolivet    = Identity{"rjolivet", "AACTGACGGATTCGATCATG", "Ringo Jolivet"}
	gchang      = Identity{"gchang", "GTTTGCACGGAACATGCAAC", "Georges Chang"}
	jkolinsky   = Identity{"jkolinsky", "GACCCGTATTTCGCTGATTG", "Jeanne Kolinsky"}
	rwhite      = Identity{"rwhite", "TCAGCTTCTAACGTTCGGGA", "Richard White"}
	afrieman    = Identity{"afrieman", "ACGTTGCAAACCTGGTACGT", "Anton Frieman"}
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
