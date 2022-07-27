package main

import (
	"strings"
	"time"
)

const (
	SEC1 = time.Minute * 10
	SEC2 = time.Minute * 5
	SEC3 = time.Minute * 2
	SEC4 = time.Minute * 1
	SEC5 = time.Second * 30
)

type ID struct {
	Login    string
	Password string
	Name     string
}

func (i ID) Keywords() []string {
	return strings.Fields(i.Name)
}

var (
	alan           = ID{"amathison", "GGCGGTAGCCCCTCTCGAGC", "Alan Mathison"}
	mel            = ID{"mmathison", "GGCCAAAGCTCCTTCGGAGC", "Mélody Mathison"}
	rocky          = ID{"jdoe7624", "CCGCGCAGAATCATAGCTGT", "John Doe 7624"} // pas d'ID
	rita           = ID{"mbellamy", "CAAAGTTCTAGGCATAGGGA", "Margherita Bellamy"}
	styx           = ID{"sbronner", "TTAGCTCGATATCCTAACCC", "Sebastian Bronner"}
	kapo           = ID{"cbellamy", "GAACTGCTTTAGTTGACGGA", "Camélia Bellamy"}
	scalpel        = ID{"jvillanova", "TGAAAGAGACATGATGCCTT", "Julius Villanova"}
	greko          = ID{"ecanto", "TCTGAGGTTTATTGATTTCG", "Eddy Canto"}
	jesus          = ID{"ejohannesen", "TTCGGGATTACTGCGTGCTG", "Edwin Johannesen"}
	escobar        = ID{"jbranson", "GGAGGACACCCCAAACGCAT", "Jonathan Branson"}
	cageot         = ID{"", "GCCCTTGTCATGTACTTAGT", ""} // TODO
	lafouine       = ID{"skmihalec", "CTGTCACCCAATCTACAGCG", "Sylvia Kemija Mihalec"}
	eva            = ID{"", "CTGTTGTAGTGACATGTTTC", "Eva"} // TODO
	fatmike        = ID{"mdubian", "AACCTTGGGCACGGTCGGTA", "Michael Dubian"}
	kennedy        = ID{"", "CCCGCGGGCAAAGCTGACAG", ""} // TODO
	savagegirl     = ID{"jdoe", "GGGTCTATAGGTCAAACGGT", "Jane Doe 2645"}
	raoulcool      = ID{"rmichu", "GTCACAAGGTTGTTTAATGG", "Raoul Michu"}
	greenglass     = ID{"rglass", "ATGCCTACCTCCAATGATTA", "Rupert Glass"}
	chillydaisy    = ID{"djohannesen", "CGGGAGACACGTTCAGTCTT", "Daisy Johannesen"}
	frereping      = ID{"dbonenfant", "GCATGGCCGAATTCCTCATT", "Désiré Bonenfant"}
	papaproxy      = ID{"hproskychev", "CGATTTGTATTGGATACGGA", "Harald Proskychev"}
	nikki          = ID{"njasinski", "ACGAACCTAGAGCCGCACGC", "Nikole Jasinski"}
	celine         = ID{"ffceline", "CGCTCCCATTTCATGTCAGC", "Franz-Ferdinand Celine"}
	cramille       = ID{"cmills", "TTTGGGAGAAGCTTATGCAC", "Camélia Mills"}
	tigerdoll      = ID{"lseptembre", "ATATGTTGAGCGTAAAGGCG", "Lilas Septembre"}
	sistermorphine = ID{"edubian", "CCATCCGGCGGACCTTATGC", "Eloïse Dubian"}
	zilmir         = ID{"zabasolo", "GACGGGATACCTACTCTCGA", "Zilmir Abasolo"}
	bettyb         = ID{"ebranson", "ATTCCGACTCAGGGTACCGG", "Elisabeth Branson"}
	abraham        = ID{"jkievain", "TGGCGTCTCTAATTCTTGCC", "Jordan Kievain"}
	crunch         = ID{"", "TTCAAGCTGAATATGAAAGG", ""}
	onekick        = ID{"rkievain", "GTCAAATCTGAGACTCTTGC", "Rodolph Kievain"}
	jacob          = ID{"", "TGAAAGAGACAGTATGCCGT", "Pete"}
	cyrano         = ID{"ajolivet", "TTCGACTGAATGTTTGATGT", "Adrien Jolivet"}
	smalljoe       = ID{"jvazzanna", "TATCGACGCACGGGACTTGG", "Joseph Vazzanna"}
	ironmike       = ID{"mklebert", "CGAGAAATGACAGAGTTGTA", "Mickael Klebert"}
	paula          = ID{"pjolivet", "GGGTGATCTGTTGCCCCCTG", "Paula Jolivet"}
	ringo          = ID{"rjolivet", "AACTGACGGATTCGATCATG", "Ringo Jolivet"}
	georges        = ID{"gchang", "GTTTGCACGGAACATGCAAC", "Georges Chang"}
	jeanne         = ID{"jkolinsky", "GACCCGTATTTCGCTGATTG", "Jeanne Kolinsky"}
	oggy           = ID{"rwhite", "TCAGCTTCTAACGTTCGGGA", "Richard White"}
)

// serveur local du dirty district
var dd = Server{
	Address: "dd.local",
	Credentials: []Cred{
		{"public", "public", 1},
		{"jesus", "roxor", 5},
	},
	Description: ddDesc,
	Targets: []Target{
		{d22.Address, "serveur public du District 22", 1, "public", "public"},
	},
	Registers: []Register{
		{"cafe", false, "machine à café", 1},
		{"sono", true, "et je coupe le son ...", 3},
	},
	Entries: []Entry{
		{"bluemars", []string{"boisson"}, 1, "jesus", "Blue Mars - le cocktail parfait", "la recette"},
		{"mojito", []string{"boisson"}, 1, "", "Mojito - le cocktail classique", "Menthe, glace pilée, citron vert et plein de rhum"},
	},
	Scan: SEC1,
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

// serveur public du district 22
var d22 = Server{
	Address: "d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: dd22Desc,
	Targets: []Target{
		{legba.Address, "Legba Voodoocom", 1, "public", "public"},
		{kramps.Address, "Kramps Security", 1, "public", "public"},
		{corp.Address, "Central Services", 1, "public", "public"},
		{abus.Address, "Association des Banques Unifiées Suisses", 1, "public", "public"},
		{greendata.Address, "Green Data, solution environnementale", 1, "public", "public"},
	},
	Scan: SEC2,
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

// serveur public de la kramps
var kramps = Server{
	Address: "kramps.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Targets: []Target{
		{kramps_priv.Address, "Serveur réservé au personnel", 3, "personnel", "kramps1234"},
	},
	Description: kpubDesc,
	Scan:        SEC2,
}

var kpubDesc = `
      ___           ___           ___           ___           ___         ___     
     /__/|         /  /\         /  /\         /__/\         /  /\       /  /\    
    |  |:|        /  /::\       /  /::\       |  |::\       /  /::\     /  /:/_   
    |  |:|       /  /:/\:\     /  /:/\:\      |  |:|:\     /  /:/\:\   /  /:/ /\  
  __|  |:|      /  /:/~/:/    /  /:/~/::\   __|__|:|\:\   /  /:/~/:/  /  /:/ /::\ 
 /__/\_|:|____ /__/:/ /:/___ /__/:/ /:/\:\ /__/::::| \:\ /__/:/ /:/  /__/:/ /:/\:\
 \  \:\/:::::/ \  \:\/:::::/ \  \:\/:/__\/ \  \:\~~\__\/ \  \:\/:/   \  \:\/:/~/:/
  \  \::/~~~~   \  \::/~~~~   \  \::/       \  \:\        \  \::/     \  \::/ /:/ 
   \  \:\        \  \:\        \  \:\        \  \:\        \  \:\      \__\/ /:/  
    \  \:\        \  \:\        \  \:\        \  \:\        \  \:\       /__/:/   
     \__\/         \__\/         \__\/         \__\/         \__\/       \__\/    
                                                                 
                                                                +-+-+-+-+-+-+-+-+
                                                                |S|e|c|u|r|i|t|y|
   _|_  | | _   |\/| _  _  _| _    _ |    _  (~ ∧  _  _|_       +-+-+-+-+-+-+-+-+
    !   |_|| |  |  |(_)| |(_|(/_  |_)||_|_\  _)|_||    !        | | | | | | | | |
                                  |                             | | | | | | | | |
                                                                | | | | | | | | |
                                                                | | | | | | | | |
                                                                | | | | | | | | |
                    Plus de 4.000.000.000 ¥€$ de capital        | | | | | | | | |
`

// serveur privé de la kramps
var kramps_priv = Server{
	Address: "priv.kramps.d22.eu",
	Credentials: []Cred{
		{"personnel", "kramps1234", 3}, // accès depuis le serveur public
		{"akremmer", "sexgod22", 3},    // backdoor, vol de compte utilisateur
	},
	Targets: []Target{
		{kramps_inmates.Address, "Gestion des prisonniers", 3, "personnel", "kramps1234"},
		{kramps_sec.Address, "Sécurité des installations", 5, "admin", "lkjqsod"},
	},
	Scan:        SEC3,
	Description: kperDesc,

	Entries: []Entry{
		{"G-F5", []string{"Sig.", "Raffaellino", "Bombieri"}, 1, "", "Sig. Raffaellino Bombieri", "Sig. Raffaellino Bombieri - 25/5/1963 - 1084dd5a-25ac-43c1-90f4-1bd3d7769680"},
		{"G-H292", []string{"Humberto", "Plana-Chacón"}, 1, "", "Humberto Plana-Chacón", "Humberto Plana-Chacón - 2/5/1999 - 2e285752-3a57-4a6c-b4cc-556d3a9c873d"},
		{"G-A37", []string{"Jonathan", "Swift"}, 1, "", "Jonathan Swift", "Jonathan Swift  - 11/4/1984 - 383de3ff-56eb-4745-b472-e046ff8e552e"},
		{"G-F254", []string{"Éric", "Huet"}, 1, "", "Éric Huet", "Éric Huet - 14/7/1968 - 4ec26008-158e-433d-9fd0-7937a5ea4e13"},
		{"G-J74", []string{"Angelina", "Deledda"}, 1, "", "Angelina Deledda", "Angelina Deledda - 21/12/1965 - d764f740-fca3-436e-b36f-8f4c63e923f1"},
		{"G-C50", []string{"Romana", "Tutino"}, 1, "", "Romana Tutino", "Romana Tutino - 1/5/1986 - 199c14a3-ff9b-4dbe-9e27-1083c890d5d3"},
		{"G-L35", []string{"Azahar", "Marcos", "Piñol"}, 1, "", "Azahar Marcos Piñol", "Azahar Marcos Piñol - 5/12/1990 - f68c0dd9-0a85-41f4-8a23-91dc30efd7b1"},
		{"G-C3", []string{"Harvey", "Zimmermann"}, 1, "", "Harvey Zimmermann", "Harvey Zimmermann  - 15/11/1961 - a37dd901-6526-4913-8900-daf9af5f8fab"},
		{"G-F196", []string{"Mary", "Flynn"}, 1, "", "Mary Flynn", "Mary Flynn - 5/9/1983 - 18518deb-c136-4045-8f58-32ed11621e25"},
		{"G-N19", []string{"Ing.", "Constantin", "Briemer"}, 1, "", "Ing. Constantin Briemer", "Ing. Constantin Briemer - 30/10/2000 - ed3f45c9-4396-4675-b2e3-545ce188bbe5"},
		{"G-C279", []string{"Dott.", "Cecilia", "Passalacqua"}, 1, "", "Dott. Cecilia Passalacqua", "Dott. Cecilia Passalacqua - 7/2/1969 - e3eca928-532f-4281-ab6d-c3c8539bfff4"},
		{"G-S20", []string{"Alfredo", "Vendetti"}, 1, "", "Alfredo Vendetti", "Alfredo Vendetti - 15/5/1958 - 5f945848-4682-4282-b8d7-65aebb52be5c"},
	}, // TODO liste associant prisonnier / matricule / numéro de cellule
	Registers: []Register{
		{"G-F5_AC1_10", false, "", 1},
		{"G-F5_AC2_10", false, "", 1},
		{"G-F5_AC3_10", false, "", 1},
		{"G-F5_DZ_10", false, "", 1},
		{"G-F5_RR_10", true, "", 1},
		{"G-F5_CE_10", false, "", 1},
		{"G-F5_AC1_14", false, "", 1},
		{"G-F5_AC2_14", false, "", 1},
		{"G-F5_AC3_14", true, "", 1},
		{"G-F5_DZ_14", false, "", 1},
		{"G-F5_RR_14", false, "", 1},
		{"G-F5_CE_14", false, "", 1},
		{"G-F5_AC1_16", false, "", 1},
		{"G-F5_AC2_16", false, "", 1},
		{"G-F5_AC3_16", false, "", 1},
		{"G-F5_DZ_16", false, "", 1},
		{"G-F5_RR_16", false, "", 1},
		{"G-F5_CE_16", true, "", 1},
		{"G-F5_AC1_20", false, "", 1},
		{"G-F5_AC2_20", false, "", 1},
		{"G-F5_AC3_20", true, "", 1},
		{"G-F5_DZ_20", false, "", 1},
		{"G-F5_RR_20", false, "", 1},
		{"G-F5_CE_20", false, "", 1},
		{"G-H292_AC1_10", false, "", 1},
		{"G-H292_AC2_10", false, "", 1},
		{"G-H292_AC3_10", false, "", 1},
		{"G-H292_DZ_10", false, "", 1},
		{"G-H292_RR_10", true, "", 1},
		{"G-H292_CE_10", false, "", 1},
		{"G-H292_AC1_14", false, "", 1},
		{"G-H292_AC2_14", false, "", 1},
		{"G-H292_AC3_14", false, "", 1},
		{"G-H292_DZ_14", false, "", 1},
		{"G-H292_RR_14", true, "", 1},
		{"G-H292_CE_14", false, "", 1},
		{"G-H292_AC1_16", false, "", 1},
		{"G-H292_AC2_16", false, "", 1},
		{"G-H292_AC3_16", false, "", 1},
		{"G-H292_DZ_16", true, "", 1},
		{"G-H292_RR_16", false, "", 1},
		{"G-H292_CE_16", false, "", 1},
		{"G-H292_AC1_20", false, "", 1},
		{"G-H292_AC2_20", false, "", 1},
		{"G-H292_AC3_20", false, "", 1},
		{"G-H292_DZ_20", false, "", 1},
		{"G-H292_RR_20", true, "", 1},
		{"G-H292_CE_20", false, "", 1},
		{"G-A37_AC1_10", false, "", 1},
		{"G-A37_AC2_10", false, "", 1},
		{"G-A37_AC3_10", true, "", 1},
		{"G-A37_DZ_10", false, "", 1},
		{"G-A37_RR_10", false, "", 1},
		{"G-A37_CE_10", false, "", 1},
		{"G-A37_AC1_14", false, "", 1},
		{"G-A37_AC2_14", true, "", 1},
		{"G-A37_AC3_14", false, "", 1},
		{"G-A37_DZ_14", false, "", 1},
		{"G-A37_RR_14", false, "", 1},
		{"G-A37_CE_14", false, "", 1},
		{"G-A37_AC1_16", false, "", 1},
		{"G-A37_AC2_16", false, "", 1},
		{"G-A37_AC3_16", false, "", 1},
		{"G-A37_DZ_16", false, "", 1},
		{"G-A37_RR_16", true, "", 1},
		{"G-A37_CE_16", false, "", 1},
		{"G-A37_AC1_20", false, "", 1},
		{"G-A37_AC2_20", true, "", 1},
		{"G-A37_AC3_20", false, "", 1},
		{"G-A37_DZ_20", false, "", 1},
		{"G-A37_RR_20", false, "", 1},
		{"G-A37_CE_20", false, "", 1},
		{"G-F254_AC1_10", false, "", 1},
		{"G-F254_AC2_10", false, "", 1},
		{"G-F254_AC3_10", false, "", 1},
		{"G-F254_DZ_10", false, "", 1},
		{"G-F254_RR_10", false, "", 1},
		{"G-F254_CE_10", true, "", 1},
		{"G-F254_AC1_14", true, "", 1},
		{"G-F254_AC2_14", false, "", 1},
		{"G-F254_AC3_14", false, "", 1},
		{"G-F254_DZ_14", false, "", 1},
		{"G-F254_RR_14", false, "", 1},
		{"G-F254_CE_14", false, "", 1},
		{"G-F254_AC1_16", false, "", 1},
		{"G-F254_AC2_16", false, "", 1},
		{"G-F254_AC3_16", false, "", 1},
		{"G-F254_DZ_16", false, "", 1},
		{"G-F254_RR_16", false, "", 1},
		{"G-F254_CE_16", true, "", 1},
		{"G-F254_AC1_20", true, "", 1},
		{"G-F254_AC2_20", false, "", 1},
		{"G-F254_AC3_20", false, "", 1},
		{"G-F254_DZ_20", false, "", 1},
		{"G-F254_RR_20", false, "", 1},
		{"G-F254_CE_20", false, "", 1},
		{"G-J74_AC1_10", false, "", 1},
		{"G-J74_AC2_10", true, "", 1},
		{"G-J74_AC3_10", false, "", 1},
		{"G-J74_DZ_10", false, "", 1},
		{"G-J74_RR_10", false, "", 1},
		{"G-J74_CE_10", false, "", 1},
		{"G-J74_AC1_14", false, "", 1},
		{"G-J74_AC2_14", false, "", 1},
		{"G-J74_AC3_14", false, "", 1},
		{"G-J74_DZ_14", false, "", 1},
		{"G-J74_RR_14", true, "", 1},
		{"G-J74_CE_14", false, "", 1},
		{"G-J74_AC1_16", false, "", 1},
		{"G-J74_AC2_16", false, "", 1},
		{"G-J74_AC3_16", false, "", 1},
		{"G-J74_DZ_16", false, "", 1},
		{"G-J74_RR_16", true, "", 1},
		{"G-J74_CE_16", false, "", 1},
		{"G-J74_AC1_20", false, "", 1},
		{"G-J74_AC2_20", false, "", 1},
		{"G-J74_AC3_20", false, "", 1},
		{"G-J74_DZ_20", false, "", 1},
		{"G-J74_RR_20", true, "", 1},
		{"G-J74_CE_20", false, "", 1},
		{"G-C50_AC1_10", false, "", 1},
		{"G-C50_AC2_10", false, "", 1},
		{"G-C50_AC3_10", false, "", 1},
		{"G-C50_DZ_10", true, "", 1},
		{"G-C50_RR_10", false, "", 1},
		{"G-C50_CE_10", false, "", 1},
		{"G-C50_AC1_14", false, "", 1},
		{"G-C50_AC2_14", false, "", 1},
		{"G-C50_AC3_14", false, "", 1},
		{"G-C50_DZ_14", true, "", 1},
		{"G-C50_RR_14", false, "", 1},
		{"G-C50_CE_14", false, "", 1},
		{"G-C50_AC1_16", false, "", 1},
		{"G-C50_AC2_16", false, "", 1},
		{"G-C50_AC3_16", false, "", 1},
		{"G-C50_DZ_16", true, "", 1},
		{"G-C50_RR_16", false, "", 1},
		{"G-C50_CE_16", false, "", 1},
		{"G-C50_AC1_20", true, "", 1},
		{"G-C50_AC2_20", false, "", 1},
		{"G-C50_AC3_20", false, "", 1},
		{"G-C50_DZ_20", false, "", 1},
		{"G-C50_RR_20", false, "", 1},
		{"G-C50_CE_20", false, "", 1},
		{"G-L35_AC1_10", false, "", 1},
		{"G-L35_AC2_10", false, "", 1},
		{"G-L35_AC3_10", false, "", 1},
		{"G-L35_DZ_10", true, "", 1},
		{"G-L35_RR_10", false, "", 1},
		{"G-L35_CE_10", false, "", 1},
		{"G-L35_AC1_14", false, "", 1},
		{"G-L35_AC2_14", false, "", 1},
		{"G-L35_AC3_14", true, "", 1},
		{"G-L35_DZ_14", false, "", 1},
		{"G-L35_RR_14", false, "", 1},
		{"G-L35_CE_14", false, "", 1},
		{"G-L35_AC1_16", false, "", 1},
		{"G-L35_AC2_16", false, "", 1},
		{"G-L35_AC3_16", false, "", 1},
		{"G-L35_DZ_16", true, "", 1},
		{"G-L35_RR_16", false, "", 1},
		{"G-L35_CE_16", false, "", 1},
		{"G-L35_AC1_20", true, "", 1},
		{"G-L35_AC2_20", false, "", 1},
		{"G-L35_AC3_20", false, "", 1},
		{"G-L35_DZ_20", false, "", 1},
		{"G-L35_RR_20", false, "", 1},
		{"G-L35_CE_20", false, "", 1},
		{"G-C3_AC1_10", true, "", 1},
		{"G-C3_AC2_10", false, "", 1},
		{"G-C3_AC3_10", false, "", 1},
		{"G-C3_DZ_10", false, "", 1},
		{"G-C3_RR_10", false, "", 1},
		{"G-C3_CE_10", false, "", 1},
		{"G-C3_AC1_14", false, "", 1},
		{"G-C3_AC2_14", false, "", 1},
		{"G-C3_AC3_14", false, "", 1},
		{"G-C3_DZ_14", true, "", 1},
		{"G-C3_RR_14", false, "", 1},
		{"G-C3_CE_14", false, "", 1},
		{"G-C3_AC1_16", false, "", 1},
		{"G-C3_AC2_16", false, "", 1},
		{"G-C3_AC3_16", false, "", 1},
		{"G-C3_DZ_16", false, "", 1},
		{"G-C3_RR_16", true, "", 1},
		{"G-C3_CE_16", false, "", 1},
		{"G-C3_AC1_20", false, "", 1},
		{"G-C3_AC2_20", false, "", 1},
		{"G-C3_AC3_20", true, "", 1},
		{"G-C3_DZ_20", false, "", 1},
		{"G-C3_RR_20", false, "", 1},
		{"G-C3_CE_20", false, "", 1},
		{"G-F196_AC1_10", false, "", 1},
		{"G-F196_AC2_10", false, "", 1},
		{"G-F196_AC3_10", false, "", 1},
		{"G-F196_DZ_10", true, "", 1},
		{"G-F196_RR_10", false, "", 1},
		{"G-F196_CE_10", false, "", 1},
		{"G-F196_AC1_14", false, "", 1},
		{"G-F196_AC2_14", false, "", 1},
		{"G-F196_AC3_14", false, "", 1},
		{"G-F196_DZ_14", false, "", 1},
		{"G-F196_RR_14", true, "", 1},
		{"G-F196_CE_14", false, "", 1},
		{"G-F196_AC1_16", false, "", 1},
		{"G-F196_AC2_16", false, "", 1},
		{"G-F196_AC3_16", false, "", 1},
		{"G-F196_DZ_16", false, "", 1},
		{"G-F196_RR_16", true, "", 1},
		{"G-F196_CE_16", false, "", 1},
		{"G-F196_AC1_20", false, "", 1},
		{"G-F196_AC2_20", false, "", 1},
		{"G-F196_AC3_20", false, "", 1},
		{"G-F196_DZ_20", false, "", 1},
		{"G-F196_RR_20", false, "", 1},
		{"G-F196_CE_20", true, "", 1},
		{"G-N19_AC1_10", false, "", 1},
		{"G-N19_AC2_10", false, "", 1},
		{"G-N19_AC3_10", false, "", 1},
		{"G-N19_DZ_10", false, "", 1},
		{"G-N19_RR_10", true, "", 1},
		{"G-N19_CE_10", false, "", 1},
		{"G-N19_AC1_14", false, "", 1},
		{"G-N19_AC2_14", false, "", 1},
		{"G-N19_AC3_14", false, "", 1},
		{"G-N19_DZ_14", true, "", 1},
		{"G-N19_RR_14", false, "", 1},
		{"G-N19_CE_14", false, "", 1},
		{"G-N19_AC1_16", false, "", 1},
		{"G-N19_AC2_16", false, "", 1},
		{"G-N19_AC3_16", false, "", 1},
		{"G-N19_DZ_16", false, "", 1},
		{"G-N19_RR_16", true, "", 1},
		{"G-N19_CE_16", false, "", 1},
		{"G-N19_AC1_20", true, "", 1},
		{"G-N19_AC2_20", false, "", 1},
		{"G-N19_AC3_20", false, "", 1},
		{"G-N19_DZ_20", false, "", 1},
		{"G-N19_RR_20", false, "", 1},
		{"G-N19_CE_20", false, "", 1},
		{"G-C279_AC1_10", false, "", 1},
		{"G-C279_AC2_10", false, "", 1},
		{"G-C279_AC3_10", false, "", 1},
		{"G-C279_DZ_10", true, "", 1},
		{"G-C279_RR_10", false, "", 1},
		{"G-C279_CE_10", false, "", 1},
		{"G-C279_AC1_14", true, "", 1},
		{"G-C279_AC2_14", false, "", 1},
		{"G-C279_AC3_14", false, "", 1},
		{"G-C279_DZ_14", false, "", 1},
		{"G-C279_RR_14", false, "", 1},
		{"G-C279_CE_14", false, "", 1},
		{"G-C279_AC1_16", false, "", 1},
		{"G-C279_AC2_16", false, "", 1},
		{"G-C279_AC3_16", false, "", 1},
		{"G-C279_DZ_16", true, "", 1},
		{"G-C279_RR_16", false, "", 1},
		{"G-C279_CE_16", false, "", 1},
		{"G-C279_AC1_20", true, "", 1},
		{"G-C279_AC2_20", false, "", 1},
		{"G-C279_AC3_20", false, "", 1},
		{"G-C279_DZ_20", false, "", 1},
		{"G-C279_RR_20", false, "", 1},
		{"G-C279_CE_20", false, "", 1},
		{"G-S20_AC1_10", false, "", 1},
		{"G-S20_AC2_10", false, "", 1},
		{"G-S20_AC3_10", false, "", 1},
		{"G-S20_DZ_10", false, "", 1},
		{"G-S20_RR_10", false, "", 1},
		{"G-S20_CE_10", true, "", 1},
		{"G-S20_AC1_14", true, "", 1},
		{"G-S20_AC2_14", false, "", 1},
		{"G-S20_AC3_14", false, "", 1},
		{"G-S20_DZ_14", false, "", 1},
		{"G-S20_RR_14", false, "", 1},
		{"G-S20_CE_14", false, "", 1},
		{"G-S20_AC1_16", false, "", 1},
		{"G-S20_AC2_16", false, "", 1},
		{"G-S20_AC3_16", false, "", 1},
		{"G-S20_DZ_16", true, "", 1},
		{"G-S20_RR_16", false, "", 1},
		{"G-S20_CE_16", false, "", 1},
		{"G-S20_AC1_20", true, "", 1},
		{"G-S20_AC2_20", false, "", 1},
		{"G-S20_AC3_20", false, "", 1},
		{"G-S20_DZ_20", false, "", 1},
		{"G-S20_RR_20", false, "", 1},
		{"G-S20_CE_20", false, "", 1},
	}, // TODO emploi du temps des prisonniers (extérieur / cellule)
}

var kperDesc = `
      ___           ___           ___           ___           ___         ___     
     /__/|         /  /\         /  /\         /__/\         /  /\       /  /\    
    |  |:|        /  /::\       /  /::\       |  |::\       /  /::\     /  /:/_   
    |  |:|       /  /:/\:\     /  /:/\:\      |  |:|:\     /  /:/\:\   /  /:/ /\  
  __|  |:|      /  /:/~/:/    /  /:/~/::\   __|__|:|\:\   /  /:/~/:/  /  /:/ /::\ 
 /__/\_|:|____ /__/:/ /:/___ /__/:/ /:/\:\ /__/::::| \:\ /__/:/ /:/  /__/:/ /:/\:\
 \  \:\/:::::/ \  \:\/:::::/ \  \:\/:/__\/ \  \:\~~\__\/ \  \:\/:/   \  \:\/:/~/:/
  \  \::/~~~~   \  \::/~~~~   \  \::/       \  \:\        \  \::/     \  \::/ /:/ 
   \  \:\        \  \:\        \  \:\        \  \:\        \  \:\      \__\/ /:/  
    \  \:\        \  \:\        \  \:\        \  \:\        \  \:\       /__/:/   
     \__\/         \__\/         \__\/         \__\/         \__\/       \__\/    
                                                                 
  [-> Serveur du personnel <-]++ toutes les transaction sont loguées (SecLvl 4)

  20/07: Rappel ! Les EdT sont calculés chaque nuit pour le lendemain.
         Aucune réclamation n'est recevable.

  18/07: StuKa est la 3° organisation du personnel (spécialisée Tech niv 1-3) à 
         déposer le bilan cette année.

  02/07: Rappel ! Les dossiers de promotions pour SecAgent, tout niveau, sont à
         renvoyer avant le 31/07. Tarif habituel, voir agence comptable.
`

var kramps_inmates = Server{
	Address:     "inmates.kramps.d22.eu",
	Credentials: []Cred{},
	Scan:        SEC3,
	Description: kinmatesDesc,
	Entries: []Entry{
		{"PR-289", []string{"Josué", "Cobos"}, 1, "", "Josué Cobos", "Josué Cobos - 1/3/1989 - 1e80def0-56e2-4334-a5e9-3900c68e924a"},
		{"PX-149", []string{"Iker", "Diaz-Figuerola"}, 1, "", "Iker Diaz-Figuerola", "Iker Diaz-Figuerola - 19/1/1960 - ecf774b2-d824-4e39-83a6-71a3ea504d48"},
		{"PO-235", []string{"Shannon", "Ward"}, 1, "", "Shannon Ward", "Shannon Ward - 19/9/1965 - 5b5e9a8f-7059-41fa-abc1-834fad65dcae"},
		{"PS-161", []string{"Dipl.-Ing.", "Rebecca", "Briemer"}, 1, "", "Dipl.-Ing. Rebecca Briemer", "Dipl.-Ing. Rebecca Briemer - 26/3/1992 - a62cc240-954f-4731-a2d9-fcf1428a6deb"},
		{"PB-32", []string{"Rembrandt", "Gatto-Togliatti"}, 1, "", "Rembrandt Gatto-Togliatti", "Rembrandt Gatto-Togliatti - 2/3/1955 - 5b073a69-cf82-41f4-a69b-766aa8dabb71"},
		{"PR-58", []string{"Gustavo", "Soto", "Blanca"}, 1, "", "Gustavo Soto Blanca", "Gustavo Soto Blanca - 19/11/1965 - 9a5947a2-b306-4927-987f-41d0171fd21b"},
		{"PK-224", []string{"Sig.", "Sabatino", "Zecchini"}, 1, "", "Sig. Sabatino Zecchini", "Sig. Sabatino Zecchini - 29/9/1980 - ca15d1f5-d2b5-471e-b7ed-bf374a65fe9f"},
		{"PF-47", []string{"Stefan", "Jasinski"}, 1, "", "Stefan Jasinski", "Stefan Jasinski - 15/3/1988 - c54dd982-6526-cd42-8900-cbe27f5f8fab"},
		{"PC-261", []string{"Eusebia", "Pozo-Botella"}, 1, "", "Eusebia Pozo-Botella", "Eusebia Pozo-Botella - 21/7/1967 - 715c2159-72ea-4738-a401-b5a58506e694"},
		{"PN-94", []string{"Nieves", "Vázquez", "Franco"}, 1, "", "Nieves Vázquez Franco", "Nieves Vázquez Franco - 30/9/1961 - 8def9643-e69b-45ba-81fe-41ffbfddb4ef"},
		{"PR-1", []string{"Grégoire", "Menard"}, 1, "", "Grégoire Menard", "Grégoire Menard - 24/3/1955 - 44a26bdf-0210-4607-abd7-aa9cadbbdf1c"},
		{"PO-114", []string{"Sébastien", "Albert"}, 1, "", "Sébastien Albert", "Sébastien Albert - 3/3/2002 - 868b22aa-370e-4d7e-ae0f-26c22eb0ec1c"},
		{"PX-280", []string{"Ing.", "Eveline", "Buchholz"}, 1, "", "Ing. Eveline Buchholz", "Ing. Eveline Buchholz - 9/2/1979 - 39de7fda-a12d-4da7-ac74-4d8802095f69"},
		{"PS-197", []string{"Bruce", "Hoffman"}, 1, "", "Bruce Hoffman", "Bruce Hoffman - 28/4/1978 - a1e241a4-6b8d-4e2e-bb77-3f93242fa942"},
		{"PV-144", []string{"Stefania", "Cagnin"}, 1, "", "Stefania Cagnin", "Stefania Cagnin - 22/8/1994 - 1e6d16c6-2f9c-4d07-804a-1b44d042cc1d"},
		{"PR-25", []string{"Adelgunde", "Krebs", "B.Eng."}, 1, "", "Adelgunde Krebs B.Eng.", "Adelgunde Krebs B.Eng. - 30/10/1974 - fe435f6a-84ae-41c5-80ae-e4239a270b00"},
		{"PN-173", []string{"Sig.", "Cesare", "Mazzi"}, 1, "", "Sig. Cesare Mazzi", "Sig. Cesare Mazzi - 24/6/1956 - ad980090-d686-4637-a204-90a0b956a68b"},
		{"PD-256", []string{"Univ.Prof.", "Wolf-Rüdiger", "Cichorius"}, 1, "", "Univ.Prof. Wolf-Rüdiger Cichorius", "Univ.Prof. Wolf-Rüdiger Cichorius - 23/6/1950 - fe2944cf-bfa1-404f-8d72-4bfcf1129376"},
		{"PA-186", []string{"Isidora", "del", "Olivera"}, 1, "", "Isidora del Olivera", "Isidora del Olivera - 30/9/1952 - e713a1d1-fbdc-4048-a12e-6f1f4961537a"},
		{"PK-222", []string{"Jolanda", "Guinizzelli-Panicucci"}, 1, "", "Jolanda Guinizzelli-Panicucci", "Jolanda Guinizzelli-Panicucci - 21/4/1953 - 9363dd30-7e40-4d91-8ea3-db67ea58c45b"},
		{"PV-28", []string{"Swen", "Ullrich"}, 1, "", "Swen Ullrich", "Swen Ullrich - 23/12/1990 - 6f2287cf-036e-4372-bc3e-ad5d408283d2"},
		{"PZ-176", []string{"Galo", "Santamaria", "Carro"}, 1, "", "Galo Santamaria Carro", "Galo Santamaria Carro - 8/6/1959 - 2b19228c-3c84-4031-a3b6-ff4be465e1b5"},
		{"PS-58", []string{"Philippe", "Guillaume", "de", "Pinto"}, 1, "", "Philippe Guillaume de Pinto", "Philippe Guillaume de Pinto - 23/9/2000 - 167b7a42-4fd6-4d12-aaf1-4f32c3704705"},
		{"PH-270", []string{"André", "de", "la", "Ribeiro"}, 1, "", "André de la Ribeiro", "André de la Ribeiro - 14/2/2002 - 4000c990-9824-49e3-b53e-30d4b9c5bf6c"},
		{"PT-19", []string{"Saverio", "Trotta-Pontecorvo"}, 1, "", "Saverio Trotta-Pontecorvo", "Saverio Trotta-Pontecorvo - 16/12/1953 - b40ca3bc-6f1b-480e-81b7-f6b744ab6d32"},
		{"PI-135", []string{"Erin", "Kelly"}, 1, "", "Erin Kelly", "Erin Kelly - 10/8/1978 - 13496acb-9112-4b34-9890-413f9f6d3343"},
		{"PP-250", []string{"Renata", "Tomasetti-Camicione"}, 1, "", "Renata Tomasetti-Camicione", "Renata Tomasetti-Camicione - 25/11/1953 - 4c0f783a-59fd-46bf-be87-e97b97db9460"},
		{"PE-15", []string{"Isidro", "Cabezas", "Dalmau"}, 1, "", "Isidro Cabezas Dalmau", "Isidro Cabezas Dalmau - 7/12/2000 - 8ec8657f-23e0-4114-b096-1b78ba7e71bb"},
		{"PN-44", []string{"Gregory", "Thomas"}, 1, "", "Gregory Thomas", "Gregory Thomas - 2/5/1969 - fd62a7e7-6a87-4ca8-a629-b03e1cbe3ecf"},
		{"PL-187", []string{"Rocío", "Carbajo"}, 1, "", "Rocío Carbajo", "Rocío Carbajo - 3/3/1983 - f23845bc-2952-4898-a165-99dd7dcfcb87"},
		{"PV-240", []string{"Amanda", "Pinamonte"}, 1, "", "Amanda Pinamonte", "Amanda Pinamonte - 14/9/2000 - 1bd20f1a-9eea-47a3-a0b8-c46e39c762fe"},
		{"PO-270", []string{"Adelia", "de", "Marquez"}, 1, "", "Adelia de Marquez", "Adelia de Marquez - 10/8/1961 - ed4529cc-118e-41b1-b8d5-d8b8f5e824b1"},
		{"PM-58", []string{"Jochem", "Scheibe-Köhler"}, 1, "", "Jochem Scheibe-Köhler", "Jochem Scheibe-Köhler - 24/11/1959 - fa641ac5-2b4b-42b4-9ec7-0fe1a6241221"},
		{"PE-66", []string{"Fabrizia", "Agostini"}, 1, "", "Fabrizia Agostini", "Fabrizia Agostini - 7/2/1981 - b8fa7024-2342-49f1-a94c-aa6eb7d5be6b"},
		{"PW-168", []string{"Jennifer", "Martin"}, 1, "", "Jennifer Martin", "Jennifer Martin - 19/1/1959 - 0471c610-34c5-49cf-aaa1-a9e06584eb88"},
		{"PC-263", []string{"Gülsen", "Trapp"}, 1, "", "Gülsen Trapp", "Gülsen Trapp - 9/10/1968 - 7554027a-4841-42d6-9ee3-222f16518ef8"},
		{"PX-217", []string{"Elizabeth", "Wilson"}, 1, "", "Elizabeth Wilson", "Elizabeth Wilson - 18/11/1987 - 1f2275db-890f-465e-8a66-9ba5e59358c5"},
		{"PI-40", []string{"Che", "Rozas", "Montesinos"}, 1, "", "Che Rozas Montesinos", "Che Rozas Montesinos - 12/12/2000 - 5ac440eb-c586-4051-9abd-736ee40c9ada"},
		{"PK-4", []string{"Gilles", "Pereira"}, 1, "", "Gilles Pereira", "Gilles Pereira - 26/5/2002 - c378d251-d729-4a88-9187-70570b2ddb89"},
		{"PC-52", []string{"Marcial", "Rodrigo"}, 1, "", "Marcial Rodrigo", "Marcial Rodrigo - 19/11/1953 - 3f5fe026-7da2-40b1-a799-159a0c2322d5"},
		{"PA-21", []string{"Pedro", "Ramirez"}, 1, "", "Pedro Ramirez", "Pedro Ramirez - 8/7/1991 - 383ca3ff-58eb-4745-efff-e046ff8e552e"},
		{"PL-25", []string{"Valentine", "Clerc"}, 1, "", "Valentine Clerc", "Valentine Clerc - 30/5/1970 - 0d4dfc82-e8f7-47e2-9c48-3726d2f8c064"},
		{"PG-11", []string{"Miguel", "Williams"}, 1, "", "Miguel Williams", "Miguel Williams - 1/2/1970 - 937859bc-9d1d-4b8d-98ae-08e541df9159"},
		{"PN-144", []string{"Vera", "Pont", "Avilés"}, 1, "", "Vera Pont Avilés", "Vera Pont Avilés - 12/7/1987 - 03c9b13c-f632-4fb6-b0cb-8aefb17e7c74"},
		{"PH-148", []string{"Robert", "Ellis"}, 1, "", "Robert Ellis", "Robert Ellis - 5/3/1971 - 007862ec-20c8-48bf-aeef-43c4ba76b8c6"},
		{"PR-8", []string{"David", "Stewart"}, 1, "", "David Stewart", "David Stewart - 1/2/1964 - acc87482-6c76-4212-a9e1-ab5060d4a6be"},
		{"PV-35", []string{"Marino", "Torre", "Canet"}, 1, "", "Marino Torre Canet", "Marino Torre Canet - 28/6/1956 - d1648b3a-51a5-4239-9da9-edf3c3e8539e"},
		{"PH-169", []string{"Jeanne", "Pages", "de", "la", "Grondin"}, 1, "", "Jeanne Pages de la Grondin", "Jeanne Pages de la Grondin - 29/6/1994 - 25494162-e01b-4b15-b7de-990648ca8d76"},
		{"PO-276", []string{"Sophia", "Nunez"}, 1, "", "Sophia Nunez", "Sophia Nunez - 5/8/1991 - 0f84aada-680f-4450-b21b-07ae9f16f047"},
		{"PO-297", []string{"Gabriela", "Águila-Viana"}, 1, "", "Gabriela Águila-Viana", "Gabriela Águila-Viana - 26/10/1957 - ec29fc1b-c977-4fce-bfba-b07e441aa5e9"},
		{"PJ-1", []string{"Rosa", "María", "Noguera", "Sastre"}, 1, "", "Rosa María Noguera Sastre", "Rosa María Noguera Sastre - 25/12/1956 - 7284cbfb-39a8-4511-a4f9-0f52dc834fe1"},
		{"PW-12", []string{"Lisa", "Steele"}, 1, "", "Lisa Steele", "Lisa Steele - 7/12/1952 - a7580a18-00ed-42ee-90d6-192bed734e86"},
	},
	Registers: []Register{
		{"PR-289_AC1_10", false, "", 1},
		{"PR-289_AC2_10", false, "", 1},
		{"PR-289_AC3_10", false, "", 1},
		{"PR-289_DZ_10", false, "", 1},
		{"PR-289_RR_10", false, "", 1},
		{"PR-289_CE_10", true, "", 1},
		{"PR-289_AC1_14", false, "", 1},
		{"PR-289_AC2_14", false, "", 1},
		{"PR-289_AC3_14", false, "", 1},
		{"PR-289_DZ_14", false, "", 1},
		{"PR-289_RR_14", false, "", 1},
		{"PR-289_CE_14", true, "", 1},
		{"PR-289_AC1_16", false, "", 1},
		{"PR-289_AC2_16", true, "", 1},
		{"PR-289_AC3_16", false, "", 1},
		{"PR-289_DZ_16", false, "", 1},
		{"PR-289_RR_16", false, "", 1},
		{"PR-289_CE_16", false, "", 1},
		{"PR-289_AC1_20", false, "", 1},
		{"PR-289_AC2_20", false, "", 1},
		{"PR-289_AC3_20", false, "", 1},
		{"PR-289_DZ_20", false, "", 1},
		{"PR-289_RR_20", true, "", 1},
		{"PR-289_CE_20", false, "", 1},
		{"PX-149_AC1_10", false, "", 1},
		{"PX-149_AC2_10", false, "", 1},
		{"PX-149_AC3_10", true, "", 1},
		{"PX-149_DZ_10", false, "", 1},
		{"PX-149_RR_10", false, "", 1},
		{"PX-149_CE_10", false, "", 1},
		{"PX-149_AC1_14", false, "", 1},
		{"PX-149_AC2_14", false, "", 1},
		{"PX-149_AC3_14", false, "", 1},
		{"PX-149_DZ_14", false, "", 1},
		{"PX-149_RR_14", false, "", 1},
		{"PX-149_CE_14", true, "", 1},
		{"PX-149_AC1_16", true, "", 1},
		{"PX-149_AC2_16", false, "", 1},
		{"PX-149_AC3_16", false, "", 1},
		{"PX-149_DZ_16", false, "", 1},
		{"PX-149_RR_16", false, "", 1},
		{"PX-149_CE_16", false, "", 1},
		{"PX-149_AC1_20", false, "", 1},
		{"PX-149_AC2_20", false, "", 1},
		{"PX-149_AC3_20", false, "", 1},
		{"PX-149_DZ_20", false, "", 1},
		{"PX-149_RR_20", true, "", 1},
		{"PX-149_CE_20", false, "", 1},
		{"PO-235_AC1_10", false, "", 1},
		{"PO-235_AC2_10", true, "", 1},
		{"PO-235_AC3_10", false, "", 1},
		{"PO-235_DZ_10", false, "", 1},
		{"PO-235_RR_10", false, "", 1},
		{"PO-235_CE_10", false, "", 1},
		{"PO-235_AC1_14", false, "", 1},
		{"PO-235_AC2_14", false, "", 1},
		{"PO-235_AC3_14", false, "", 1},
		{"PO-235_DZ_14", false, "", 1},
		{"PO-235_RR_14", false, "", 1},
		{"PO-235_CE_14", true, "", 1},
		{"PO-235_AC1_16", true, "", 1},
		{"PO-235_AC2_16", false, "", 1},
		{"PO-235_AC3_16", false, "", 1},
		{"PO-235_DZ_16", false, "", 1},
		{"PO-235_RR_16", false, "", 1},
		{"PO-235_CE_16", false, "", 1},
		{"PO-235_AC1_20", false, "", 1},
		{"PO-235_AC2_20", false, "", 1},
		{"PO-235_AC3_20", true, "", 1},
		{"PO-235_DZ_20", false, "", 1},
		{"PO-235_RR_20", false, "", 1},
		{"PO-235_CE_20", false, "", 1},
		{"PS-161_AC1_10", false, "", 1},
		{"PS-161_AC2_10", false, "", 1},
		{"PS-161_AC3_10", false, "", 1},
		{"PS-161_DZ_10", false, "", 1},
		{"PS-161_RR_10", true, "", 1},
		{"PS-161_CE_10", false, "", 1},
		{"PS-161_AC1_14", false, "", 1},
		{"PS-161_AC2_14", false, "", 1},
		{"PS-161_AC3_14", false, "", 1},
		{"PS-161_DZ_14", true, "", 1},
		{"PS-161_RR_14", false, "", 1},
		{"PS-161_CE_14", false, "", 1},
		{"PS-161_AC1_16", true, "", 1},
		{"PS-161_AC2_16", false, "", 1},
		{"PS-161_AC3_16", false, "", 1},
		{"PS-161_DZ_16", false, "", 1},
		{"PS-161_RR_16", false, "", 1},
		{"PS-161_CE_16", false, "", 1},
		{"PS-161_AC1_20", false, "", 1},
		{"PS-161_AC2_20", false, "", 1},
		{"PS-161_AC3_20", false, "", 1},
		{"PS-161_DZ_20", false, "", 1},
		{"PS-161_RR_20", true, "", 1},
		{"PS-161_CE_20", false, "", 1},
		{"PB-32_AC1_10", false, "", 1},
		{"PB-32_AC2_10", true, "", 1},
		{"PB-32_AC3_10", false, "", 1},
		{"PB-32_DZ_10", false, "", 1},
		{"PB-32_RR_10", false, "", 1},
		{"PB-32_CE_10", false, "", 1},
		{"PB-32_AC1_14", false, "", 1},
		{"PB-32_AC2_14", false, "", 1},
		{"PB-32_AC3_14", true, "", 1},
		{"PB-32_DZ_14", false, "", 1},
		{"PB-32_RR_14", false, "", 1},
		{"PB-32_CE_14", false, "", 1},
		{"PB-32_AC1_16", false, "", 1},
		{"PB-32_AC2_16", false, "", 1},
		{"PB-32_AC3_16", false, "", 1},
		{"PB-32_DZ_16", true, "", 1},
		{"PB-32_RR_16", false, "", 1},
		{"PB-32_CE_16", false, "", 1},
		{"PB-32_AC1_20", false, "", 1},
		{"PB-32_AC2_20", false, "", 1},
		{"PB-32_AC3_20", true, "", 1},
		{"PB-32_DZ_20", false, "", 1},
		{"PB-32_RR_20", false, "", 1},
		{"PB-32_CE_20", false, "", 1},
		{"PR-58_AC1_10", false, "", 1},
		{"PR-58_AC2_10", false, "", 1},
		{"PR-58_AC3_10", true, "", 1},
		{"PR-58_DZ_10", false, "", 1},
		{"PR-58_RR_10", false, "", 1},
		{"PR-58_CE_10", false, "", 1},
		{"PR-58_AC1_14", false, "", 1},
		{"PR-58_AC2_14", false, "", 1},
		{"PR-58_AC3_14", false, "", 1},
		{"PR-58_DZ_14", true, "", 1},
		{"PR-58_RR_14", false, "", 1},
		{"PR-58_CE_14", false, "", 1},
		{"PR-58_AC1_16", false, "", 1},
		{"PR-58_AC2_16", false, "", 1},
		{"PR-58_AC3_16", true, "", 1},
		{"PR-58_DZ_16", false, "", 1},
		{"PR-58_RR_16", false, "", 1},
		{"PR-58_CE_16", false, "", 1},
		{"PR-58_AC1_20", false, "", 1},
		{"PR-58_AC2_20", true, "", 1},
		{"PR-58_AC3_20", false, "", 1},
		{"PR-58_DZ_20", false, "", 1},
		{"PR-58_RR_20", false, "", 1},
		{"PR-58_CE_20", false, "", 1},
		{"PK-224_AC1_10", false, "", 1},
		{"PK-224_AC2_10", false, "", 1},
		{"PK-224_AC3_10", false, "", 1},
		{"PK-224_DZ_10", false, "", 1},
		{"PK-224_RR_10", false, "", 1},
		{"PK-224_CE_10", true, "", 1},
		{"PK-224_AC1_14", false, "", 1},
		{"PK-224_AC2_14", false, "", 1},
		{"PK-224_AC3_14", false, "", 1},
		{"PK-224_DZ_14", false, "", 1},
		{"PK-224_RR_14", true, "", 1},
		{"PK-224_CE_14", false, "", 1},
		{"PK-224_AC1_16", false, "", 1},
		{"PK-224_AC2_16", true, "", 1},
		{"PK-224_AC3_16", false, "", 1},
		{"PK-224_DZ_16", false, "", 1},
		{"PK-224_RR_16", false, "", 1},
		{"PK-224_CE_16", false, "", 1},
		{"PK-224_AC1_20", false, "", 1},
		{"PK-224_AC2_20", false, "", 1},
		{"PK-224_AC3_20", true, "", 1},
		{"PK-224_DZ_20", false, "", 1},
		{"PK-224_RR_20", false, "", 1},
		{"PK-224_CE_20", false, "", 1},
		{"PF-47_AC1_10", false, "", 1},
		{"PF-47_AC2_10", false, "", 1},
		{"PF-47_AC3_10", false, "", 1},
		{"PF-47_DZ_10", true, "", 1},
		{"PF-47_RR_10", false, "", 1},
		{"PF-47_CE_10", false, "", 1},
		{"PF-47_AC1_14", false, "", 1},
		{"PF-47_AC2_14", false, "", 1},
		{"PF-47_AC3_14", false, "", 1},
		{"PF-47_DZ_14", false, "", 1},
		{"PF-47_RR_14", false, "", 1},
		{"PF-47_CE_14", true, "", 1},
		{"PF-47_AC1_16", false, "", 1},
		{"PF-47_AC2_16", true, "", 1},
		{"PF-47_AC3_16", false, "", 1},
		{"PF-47_DZ_16", false, "", 1},
		{"PF-47_RR_16", false, "", 1},
		{"PF-47_CE_16", false, "", 1},
		{"PF-47_AC1_20", false, "", 1},
		{"PF-47_AC2_20", false, "", 1},
		{"PF-47_AC3_20", false, "", 1},
		{"PF-47_DZ_20", false, "", 1},
		{"PF-47_RR_20", false, "", 1},
		{"PF-47_CE_20", true, "", 1},
		{"PC-261_AC1_10", true, "", 1},
		{"PC-261_AC2_10", false, "", 1},
		{"PC-261_AC3_10", false, "", 1},
		{"PC-261_DZ_10", false, "", 1},
		{"PC-261_RR_10", false, "", 1},
		{"PC-261_CE_10", false, "", 1},
		{"PC-261_AC1_14", false, "", 1},
		{"PC-261_AC2_14", false, "", 1},
		{"PC-261_AC3_14", false, "", 1},
		{"PC-261_DZ_14", true, "", 1},
		{"PC-261_RR_14", false, "", 1},
		{"PC-261_CE_14", false, "", 1},
		{"PC-261_AC1_16", false, "", 1},
		{"PC-261_AC2_16", false, "", 1},
		{"PC-261_AC3_16", true, "", 1},
		{"PC-261_DZ_16", false, "", 1},
		{"PC-261_RR_16", false, "", 1},
		{"PC-261_CE_16", false, "", 1},
		{"PC-261_AC1_20", false, "", 1},
		{"PC-261_AC2_20", false, "", 1},
		{"PC-261_AC3_20", false, "", 1},
		{"PC-261_DZ_20", true, "", 1},
		{"PC-261_RR_20", false, "", 1},
		{"PC-261_CE_20", false, "", 1},
		{"PN-94_AC1_10", false, "", 1},
		{"PN-94_AC2_10", true, "", 1},
		{"PN-94_AC3_10", false, "", 1},
		{"PN-94_DZ_10", false, "", 1},
		{"PN-94_RR_10", false, "", 1},
		{"PN-94_CE_10", false, "", 1},
		{"PN-94_AC1_14", false, "", 1},
		{"PN-94_AC2_14", false, "", 1},
		{"PN-94_AC3_14", false, "", 1},
		{"PN-94_DZ_14", true, "", 1},
		{"PN-94_RR_14", false, "", 1},
		{"PN-94_CE_14", false, "", 1},
		{"PN-94_AC1_16", false, "", 1},
		{"PN-94_AC2_16", false, "", 1},
		{"PN-94_AC3_16", false, "", 1},
		{"PN-94_DZ_16", false, "", 1},
		{"PN-94_RR_16", true, "", 1},
		{"PN-94_CE_16", false, "", 1},
		{"PN-94_AC1_20", false, "", 1},
		{"PN-94_AC2_20", true, "", 1},
		{"PN-94_AC3_20", false, "", 1},
		{"PN-94_DZ_20", false, "", 1},
		{"PN-94_RR_20", false, "", 1},
		{"PN-94_CE_20", false, "", 1},
		{"PR-1_AC1_10", false, "", 1},
		{"PR-1_AC2_10", true, "", 1},
		{"PR-1_AC3_10", false, "", 1},
		{"PR-1_DZ_10", false, "", 1},
		{"PR-1_RR_10", false, "", 1},
		{"PR-1_CE_10", false, "", 1},
		{"PR-1_AC1_14", false, "", 1},
		{"PR-1_AC2_14", false, "", 1},
		{"PR-1_AC3_14", true, "", 1},
		{"PR-1_DZ_14", false, "", 1},
		{"PR-1_RR_14", false, "", 1},
		{"PR-1_CE_14", false, "", 1},
		{"PR-1_AC1_16", false, "", 1},
		{"PR-1_AC2_16", false, "", 1},
		{"PR-1_AC3_16", false, "", 1},
		{"PR-1_DZ_16", false, "", 1},
		{"PR-1_RR_16", true, "", 1},
		{"PR-1_CE_16", false, "", 1},
		{"PR-1_AC1_20", false, "", 1},
		{"PR-1_AC2_20", false, "", 1},
		{"PR-1_AC3_20", true, "", 1},
		{"PR-1_DZ_20", false, "", 1},
		{"PR-1_RR_20", false, "", 1},
		{"PR-1_CE_20", false, "", 1},
		{"PO-114_AC1_10", false, "", 1},
		{"PO-114_AC2_10", true, "", 1},
		{"PO-114_AC3_10", false, "", 1},
		{"PO-114_DZ_10", false, "", 1},
		{"PO-114_RR_10", false, "", 1},
		{"PO-114_CE_10", false, "", 1},
		{"PO-114_AC1_14", false, "", 1},
		{"PO-114_AC2_14", false, "", 1},
		{"PO-114_AC3_14", true, "", 1},
		{"PO-114_DZ_14", false, "", 1},
		{"PO-114_RR_14", false, "", 1},
		{"PO-114_CE_14", false, "", 1},
		{"PO-114_AC1_16", false, "", 1},
		{"PO-114_AC2_16", false, "", 1},
		{"PO-114_AC3_16", false, "", 1},
		{"PO-114_DZ_16", false, "", 1},
		{"PO-114_RR_16", false, "", 1},
		{"PO-114_CE_16", true, "", 1},
		{"PO-114_AC1_20", false, "", 1},
		{"PO-114_AC2_20", false, "", 1},
		{"PO-114_AC3_20", false, "", 1},
		{"PO-114_DZ_20", false, "", 1},
		{"PO-114_RR_20", false, "", 1},
		{"PO-114_CE_20", true, "", 1},
		{"PX-280_AC1_10", false, "", 1},
		{"PX-280_AC2_10", false, "", 1},
		{"PX-280_AC3_10", false, "", 1},
		{"PX-280_DZ_10", false, "", 1},
		{"PX-280_RR_10", false, "", 1},
		{"PX-280_CE_10", true, "", 1},
		{"PX-280_AC1_14", false, "", 1},
		{"PX-280_AC2_14", false, "", 1},
		{"PX-280_AC3_14", false, "", 1},
		{"PX-280_DZ_14", false, "", 1},
		{"PX-280_RR_14", false, "", 1},
		{"PX-280_CE_14", true, "", 1},
		{"PX-280_AC1_16", false, "", 1},
		{"PX-280_AC2_16", false, "", 1},
		{"PX-280_AC3_16", false, "", 1},
		{"PX-280_DZ_16", false, "", 1},
		{"PX-280_RR_16", true, "", 1},
		{"PX-280_CE_16", false, "", 1},
		{"PX-280_AC1_20", false, "", 1},
		{"PX-280_AC2_20", false, "", 1},
		{"PX-280_AC3_20", false, "", 1},
		{"PX-280_DZ_20", false, "", 1},
		{"PX-280_RR_20", false, "", 1},
		{"PX-280_CE_20", true, "", 1},
		{"PS-197_AC1_10", false, "", 1},
		{"PS-197_AC2_10", false, "", 1},
		{"PS-197_AC3_10", false, "", 1},
		{"PS-197_DZ_10", false, "", 1},
		{"PS-197_RR_10", false, "", 1},
		{"PS-197_CE_10", true, "", 1},
		{"PS-197_AC1_14", true, "", 1},
		{"PS-197_AC2_14", false, "", 1},
		{"PS-197_AC3_14", false, "", 1},
		{"PS-197_DZ_14", false, "", 1},
		{"PS-197_RR_14", false, "", 1},
		{"PS-197_CE_14", false, "", 1},
		{"PS-197_AC1_16", false, "", 1},
		{"PS-197_AC2_16", false, "", 1},
		{"PS-197_AC3_16", true, "", 1},
		{"PS-197_DZ_16", false, "", 1},
		{"PS-197_RR_16", false, "", 1},
		{"PS-197_CE_16", false, "", 1},
		{"PS-197_AC1_20", false, "", 1},
		{"PS-197_AC2_20", false, "", 1},
		{"PS-197_AC3_20", false, "", 1},
		{"PS-197_DZ_20", true, "", 1},
		{"PS-197_RR_20", false, "", 1},
		{"PS-197_CE_20", false, "", 1},
		{"PV-144_AC1_10", false, "", 1},
		{"PV-144_AC2_10", false, "", 1},
		{"PV-144_AC3_10", true, "", 1},
		{"PV-144_DZ_10", false, "", 1},
		{"PV-144_RR_10", false, "", 1},
		{"PV-144_CE_10", false, "", 1},
		{"PV-144_AC1_14", true, "", 1},
		{"PV-144_AC2_14", false, "", 1},
		{"PV-144_AC3_14", false, "", 1},
		{"PV-144_DZ_14", false, "", 1},
		{"PV-144_RR_14", false, "", 1},
		{"PV-144_CE_14", false, "", 1},
		{"PV-144_AC1_16", false, "", 1},
		{"PV-144_AC2_16", false, "", 1},
		{"PV-144_AC3_16", false, "", 1},
		{"PV-144_DZ_16", false, "", 1},
		{"PV-144_RR_16", true, "", 1},
		{"PV-144_CE_16", false, "", 1},
		{"PV-144_AC1_20", false, "", 1},
		{"PV-144_AC2_20", false, "", 1},
		{"PV-144_AC3_20", false, "", 1},
		{"PV-144_DZ_20", false, "", 1},
		{"PV-144_RR_20", true, "", 1},
		{"PV-144_CE_20", false, "", 1},
		{"PR-25_AC1_10", true, "", 1},
		{"PR-25_AC2_10", false, "", 1},
		{"PR-25_AC3_10", false, "", 1},
		{"PR-25_DZ_10", false, "", 1},
		{"PR-25_RR_10", false, "", 1},
		{"PR-25_CE_10", false, "", 1},
		{"PR-25_AC1_14", true, "", 1},
		{"PR-25_AC2_14", false, "", 1},
		{"PR-25_AC3_14", false, "", 1},
		{"PR-25_DZ_14", false, "", 1},
		{"PR-25_RR_14", false, "", 1},
		{"PR-25_CE_14", false, "", 1},
		{"PR-25_AC1_16", false, "", 1},
		{"PR-25_AC2_16", false, "", 1},
		{"PR-25_AC3_16", false, "", 1},
		{"PR-25_DZ_16", false, "", 1},
		{"PR-25_RR_16", false, "", 1},
		{"PR-25_CE_16", true, "", 1},
		{"PR-25_AC1_20", false, "", 1},
		{"PR-25_AC2_20", false, "", 1},
		{"PR-25_AC3_20", false, "", 1},
		{"PR-25_DZ_20", true, "", 1},
		{"PR-25_RR_20", false, "", 1},
		{"PR-25_CE_20", false, "", 1},
		{"PN-173_AC1_10", false, "", 1},
		{"PN-173_AC2_10", false, "", 1},
		{"PN-173_AC3_10", false, "", 1},
		{"PN-173_DZ_10", true, "", 1},
		{"PN-173_RR_10", false, "", 1},
		{"PN-173_CE_10", false, "", 1},
		{"PN-173_AC1_14", true, "", 1},
		{"PN-173_AC2_14", false, "", 1},
		{"PN-173_AC3_14", false, "", 1},
		{"PN-173_DZ_14", false, "", 1},
		{"PN-173_RR_14", false, "", 1},
		{"PN-173_CE_14", false, "", 1},
		{"PN-173_AC1_16", false, "", 1},
		{"PN-173_AC2_16", true, "", 1},
		{"PN-173_AC3_16", false, "", 1},
		{"PN-173_DZ_16", false, "", 1},
		{"PN-173_RR_16", false, "", 1},
		{"PN-173_CE_16", false, "", 1},
		{"PN-173_AC1_20", false, "", 1},
		{"PN-173_AC2_20", false, "", 1},
		{"PN-173_AC3_20", false, "", 1},
		{"PN-173_DZ_20", true, "", 1},
		{"PN-173_RR_20", false, "", 1},
		{"PN-173_CE_20", false, "", 1},
		{"PD-256_AC1_10", true, "", 1},
		{"PD-256_AC2_10", false, "", 1},
		{"PD-256_AC3_10", false, "", 1},
		{"PD-256_DZ_10", false, "", 1},
		{"PD-256_RR_10", false, "", 1},
		{"PD-256_CE_10", false, "", 1},
		{"PD-256_AC1_14", true, "", 1},
		{"PD-256_AC2_14", false, "", 1},
		{"PD-256_AC3_14", false, "", 1},
		{"PD-256_DZ_14", false, "", 1},
		{"PD-256_RR_14", false, "", 1},
		{"PD-256_CE_14", false, "", 1},
		{"PD-256_AC1_16", false, "", 1},
		{"PD-256_AC2_16", false, "", 1},
		{"PD-256_AC3_16", true, "", 1},
		{"PD-256_DZ_16", false, "", 1},
		{"PD-256_RR_16", false, "", 1},
		{"PD-256_CE_16", false, "", 1},
		{"PD-256_AC1_20", false, "", 1},
		{"PD-256_AC2_20", false, "", 1},
		{"PD-256_AC3_20", false, "", 1},
		{"PD-256_DZ_20", true, "", 1},
		{"PD-256_RR_20", false, "", 1},
		{"PD-256_CE_20", false, "", 1},
		{"PA-186_AC1_10", true, "", 1},
		{"PA-186_AC2_10", false, "", 1},
		{"PA-186_AC3_10", false, "", 1},
		{"PA-186_DZ_10", false, "", 1},
		{"PA-186_RR_10", false, "", 1},
		{"PA-186_CE_10", false, "", 1},
		{"PA-186_AC1_14", false, "", 1},
		{"PA-186_AC2_14", false, "", 1},
		{"PA-186_AC3_14", true, "", 1},
		{"PA-186_DZ_14", false, "", 1},
		{"PA-186_RR_14", false, "", 1},
		{"PA-186_CE_14", false, "", 1},
		{"PA-186_AC1_16", true, "", 1},
		{"PA-186_AC2_16", false, "", 1},
		{"PA-186_AC3_16", false, "", 1},
		{"PA-186_DZ_16", false, "", 1},
		{"PA-186_RR_16", false, "", 1},
		{"PA-186_CE_16", false, "", 1},
		{"PA-186_AC1_20", false, "", 1},
		{"PA-186_AC2_20", false, "", 1},
		{"PA-186_AC3_20", false, "", 1},
		{"PA-186_DZ_20", true, "", 1},
		{"PA-186_RR_20", false, "", 1},
		{"PA-186_CE_20", false, "", 1},
		{"PK-222_AC1_10", false, "", 1},
		{"PK-222_AC2_10", false, "", 1},
		{"PK-222_AC3_10", false, "", 1},
		{"PK-222_DZ_10", false, "", 1},
		{"PK-222_RR_10", false, "", 1},
		{"PK-222_CE_10", true, "", 1},
		{"PK-222_AC1_14", true, "", 1},
		{"PK-222_AC2_14", false, "", 1},
		{"PK-222_AC3_14", false, "", 1},
		{"PK-222_DZ_14", false, "", 1},
		{"PK-222_RR_14", false, "", 1},
		{"PK-222_CE_14", false, "", 1},
		{"PK-222_AC1_16", false, "", 1},
		{"PK-222_AC2_16", true, "", 1},
		{"PK-222_AC3_16", false, "", 1},
		{"PK-222_DZ_16", false, "", 1},
		{"PK-222_RR_16", false, "", 1},
		{"PK-222_CE_16", false, "", 1},
		{"PK-222_AC1_20", false, "", 1},
		{"PK-222_AC2_20", true, "", 1},
		{"PK-222_AC3_20", false, "", 1},
		{"PK-222_DZ_20", false, "", 1},
		{"PK-222_RR_20", false, "", 1},
		{"PK-222_CE_20", false, "", 1},
		{"PV-28_AC1_10", false, "", 1},
		{"PV-28_AC2_10", false, "", 1},
		{"PV-28_AC3_10", false, "", 1},
		{"PV-28_DZ_10", false, "", 1},
		{"PV-28_RR_10", false, "", 1},
		{"PV-28_CE_10", true, "", 1},
		{"PV-28_AC1_14", false, "", 1},
		{"PV-28_AC2_14", false, "", 1},
		{"PV-28_AC3_14", true, "", 1},
		{"PV-28_DZ_14", false, "", 1},
		{"PV-28_RR_14", false, "", 1},
		{"PV-28_CE_14", false, "", 1},
		{"PV-28_AC1_16", false, "", 1},
		{"PV-28_AC2_16", false, "", 1},
		{"PV-28_AC3_16", true, "", 1},
		{"PV-28_DZ_16", false, "", 1},
		{"PV-28_RR_16", false, "", 1},
		{"PV-28_CE_16", false, "", 1},
		{"PV-28_AC1_20", false, "", 1},
		{"PV-28_AC2_20", false, "", 1},
		{"PV-28_AC3_20", true, "", 1},
		{"PV-28_DZ_20", false, "", 1},
		{"PV-28_RR_20", false, "", 1},
		{"PV-28_CE_20", false, "", 1},
		{"PZ-176_AC1_10", false, "", 1},
		{"PZ-176_AC2_10", false, "", 1},
		{"PZ-176_AC3_10", false, "", 1},
		{"PZ-176_DZ_10", true, "", 1},
		{"PZ-176_RR_10", false, "", 1},
		{"PZ-176_CE_10", false, "", 1},
		{"PZ-176_AC1_14", false, "", 1},
		{"PZ-176_AC2_14", false, "", 1},
		{"PZ-176_AC3_14", false, "", 1},
		{"PZ-176_DZ_14", false, "", 1},
		{"PZ-176_RR_14", false, "", 1},
		{"PZ-176_CE_14", true, "", 1},
		{"PZ-176_AC1_16", false, "", 1},
		{"PZ-176_AC2_16", false, "", 1},
		{"PZ-176_AC3_16", false, "", 1},
		{"PZ-176_DZ_16", true, "", 1},
		{"PZ-176_RR_16", false, "", 1},
		{"PZ-176_CE_16", false, "", 1},
		{"PZ-176_AC1_20", false, "", 1},
		{"PZ-176_AC2_20", false, "", 1},
		{"PZ-176_AC3_20", false, "", 1},
		{"PZ-176_DZ_20", true, "", 1},
		{"PZ-176_RR_20", false, "", 1},
		{"PZ-176_CE_20", false, "", 1},
		{"PS-58_AC1_10", false, "", 1},
		{"PS-58_AC2_10", false, "", 1},
		{"PS-58_AC3_10", false, "", 1},
		{"PS-58_DZ_10", false, "", 1},
		{"PS-58_RR_10", true, "", 1},
		{"PS-58_CE_10", false, "", 1},
		{"PS-58_AC1_14", false, "", 1},
		{"PS-58_AC2_14", false, "", 1},
		{"PS-58_AC3_14", false, "", 1},
		{"PS-58_DZ_14", false, "", 1},
		{"PS-58_RR_14", false, "", 1},
		{"PS-58_CE_14", true, "", 1},
		{"PS-58_AC1_16", false, "", 1},
		{"PS-58_AC2_16", false, "", 1},
		{"PS-58_AC3_16", false, "", 1},
		{"PS-58_DZ_16", false, "", 1},
		{"PS-58_RR_16", false, "", 1},
		{"PS-58_CE_16", true, "", 1},
		{"PS-58_AC1_20", false, "", 1},
		{"PS-58_AC2_20", false, "", 1},
		{"PS-58_AC3_20", true, "", 1},
		{"PS-58_DZ_20", false, "", 1},
		{"PS-58_RR_20", false, "", 1},
		{"PS-58_CE_20", false, "", 1},
		{"PH-270_AC1_10", false, "", 1},
		{"PH-270_AC2_10", false, "", 1},
		{"PH-270_AC3_10", true, "", 1},
		{"PH-270_DZ_10", false, "", 1},
		{"PH-270_RR_10", false, "", 1},
		{"PH-270_CE_10", false, "", 1},
		{"PH-270_AC1_14", false, "", 1},
		{"PH-270_AC2_14", true, "", 1},
		{"PH-270_AC3_14", false, "", 1},
		{"PH-270_DZ_14", false, "", 1},
		{"PH-270_RR_14", false, "", 1},
		{"PH-270_CE_14", false, "", 1},
		{"PH-270_AC1_16", false, "", 1},
		{"PH-270_AC2_16", false, "", 1},
		{"PH-270_AC3_16", false, "", 1},
		{"PH-270_DZ_16", true, "", 1},
		{"PH-270_RR_16", false, "", 1},
		{"PH-270_CE_16", false, "", 1},
		{"PH-270_AC1_20", false, "", 1},
		{"PH-270_AC2_20", false, "", 1},
		{"PH-270_AC3_20", false, "", 1},
		{"PH-270_DZ_20", false, "", 1},
		{"PH-270_RR_20", false, "", 1},
		{"PH-270_CE_20", true, "", 1},
		{"PT-19_AC1_10", false, "", 1},
		{"PT-19_AC2_10", false, "", 1},
		{"PT-19_AC3_10", false, "", 1},
		{"PT-19_DZ_10", false, "", 1},
		{"PT-19_RR_10", true, "", 1},
		{"PT-19_CE_10", false, "", 1},
		{"PT-19_AC1_14", false, "", 1},
		{"PT-19_AC2_14", true, "", 1},
		{"PT-19_AC3_14", false, "", 1},
		{"PT-19_DZ_14", false, "", 1},
		{"PT-19_RR_14", false, "", 1},
		{"PT-19_CE_14", false, "", 1},
		{"PT-19_AC1_16", false, "", 1},
		{"PT-19_AC2_16", false, "", 1},
		{"PT-19_AC3_16", false, "", 1},
		{"PT-19_DZ_16", true, "", 1},
		{"PT-19_RR_16", false, "", 1},
		{"PT-19_CE_16", false, "", 1},
		{"PT-19_AC1_20", false, "", 1},
		{"PT-19_AC2_20", false, "", 1},
		{"PT-19_AC3_20", false, "", 1},
		{"PT-19_DZ_20", false, "", 1},
		{"PT-19_RR_20", false, "", 1},
		{"PT-19_CE_20", true, "", 1},
		{"PI-135_AC1_10", false, "", 1},
		{"PI-135_AC2_10", false, "", 1},
		{"PI-135_AC3_10", false, "", 1},
		{"PI-135_DZ_10", true, "", 1},
		{"PI-135_RR_10", false, "", 1},
		{"PI-135_CE_10", false, "", 1},
		{"PI-135_AC1_14", true, "", 1},
		{"PI-135_AC2_14", false, "", 1},
		{"PI-135_AC3_14", false, "", 1},
		{"PI-135_DZ_14", false, "", 1},
		{"PI-135_RR_14", false, "", 1},
		{"PI-135_CE_14", false, "", 1},
		{"PI-135_AC1_16", false, "", 1},
		{"PI-135_AC2_16", false, "", 1},
		{"PI-135_AC3_16", false, "", 1},
		{"PI-135_DZ_16", false, "", 1},
		{"PI-135_RR_16", true, "", 1},
		{"PI-135_CE_16", false, "", 1},
		{"PI-135_AC1_20", false, "", 1},
		{"PI-135_AC2_20", false, "", 1},
		{"PI-135_AC3_20", false, "", 1},
		{"PI-135_DZ_20", false, "", 1},
		{"PI-135_RR_20", true, "", 1},
		{"PI-135_CE_20", false, "", 1},
		{"PP-250_AC1_10", false, "", 1},
		{"PP-250_AC2_10", false, "", 1},
		{"PP-250_AC3_10", false, "", 1},
		{"PP-250_DZ_10", false, "", 1},
		{"PP-250_RR_10", true, "", 1},
		{"PP-250_CE_10", false, "", 1},
		{"PP-250_AC1_14", false, "", 1},
		{"PP-250_AC2_14", false, "", 1},
		{"PP-250_AC3_14", false, "", 1},
		{"PP-250_DZ_14", true, "", 1},
		{"PP-250_RR_14", false, "", 1},
		{"PP-250_CE_14", false, "", 1},
		{"PP-250_AC1_16", false, "", 1},
		{"PP-250_AC2_16", true, "", 1},
		{"PP-250_AC3_16", false, "", 1},
		{"PP-250_DZ_16", false, "", 1},
		{"PP-250_RR_16", false, "", 1},
		{"PP-250_CE_16", false, "", 1},
		{"PP-250_AC1_20", false, "", 1},
		{"PP-250_AC2_20", true, "", 1},
		{"PP-250_AC3_20", false, "", 1},
		{"PP-250_DZ_20", false, "", 1},
		{"PP-250_RR_20", false, "", 1},
		{"PP-250_CE_20", false, "", 1},
		{"PE-15_AC1_10", false, "", 1},
		{"PE-15_AC2_10", true, "", 1},
		{"PE-15_AC3_10", false, "", 1},
		{"PE-15_DZ_10", false, "", 1},
		{"PE-15_RR_10", false, "", 1},
		{"PE-15_CE_10", false, "", 1},
		{"PE-15_AC1_14", false, "", 1},
		{"PE-15_AC2_14", false, "", 1},
		{"PE-15_AC3_14", false, "", 1},
		{"PE-15_DZ_14", true, "", 1},
		{"PE-15_RR_14", false, "", 1},
		{"PE-15_CE_14", false, "", 1},
		{"PE-15_AC1_16", false, "", 1},
		{"PE-15_AC2_16", false, "", 1},
		{"PE-15_AC3_16", false, "", 1},
		{"PE-15_DZ_16", false, "", 1},
		{"PE-15_RR_16", true, "", 1},
		{"PE-15_CE_16", false, "", 1},
		{"PE-15_AC1_20", false, "", 1},
		{"PE-15_AC2_20", false, "", 1},
		{"PE-15_AC3_20", false, "", 1},
		{"PE-15_DZ_20", false, "", 1},
		{"PE-15_RR_20", true, "", 1},
		{"PE-15_CE_20", false, "", 1},
		{"PN-44_AC1_10", false, "", 1},
		{"PN-44_AC2_10", false, "", 1},
		{"PN-44_AC3_10", true, "", 1},
		{"PN-44_DZ_10", false, "", 1},
		{"PN-44_RR_10", false, "", 1},
		{"PN-44_CE_10", false, "", 1},
		{"PN-44_AC1_14", false, "", 1},
		{"PN-44_AC2_14", false, "", 1},
		{"PN-44_AC3_14", false, "", 1},
		{"PN-44_DZ_14", false, "", 1},
		{"PN-44_RR_14", false, "", 1},
		{"PN-44_CE_14", true, "", 1},
		{"PN-44_AC1_16", false, "", 1},
		{"PN-44_AC2_16", false, "", 1},
		{"PN-44_AC3_16", false, "", 1},
		{"PN-44_DZ_16", false, "", 1},
		{"PN-44_RR_16", true, "", 1},
		{"PN-44_CE_16", false, "", 1},
		{"PN-44_AC1_20", true, "", 1},
		{"PN-44_AC2_20", false, "", 1},
		{"PN-44_AC3_20", false, "", 1},
		{"PN-44_DZ_20", false, "", 1},
		{"PN-44_RR_20", false, "", 1},
		{"PN-44_CE_20", false, "", 1},
		{"PL-187_AC1_10", false, "", 1},
		{"PL-187_AC2_10", false, "", 1},
		{"PL-187_AC3_10", false, "", 1},
		{"PL-187_DZ_10", true, "", 1},
		{"PL-187_RR_10", false, "", 1},
		{"PL-187_CE_10", false, "", 1},
		{"PL-187_AC1_14", false, "", 1},
		{"PL-187_AC2_14", false, "", 1},
		{"PL-187_AC3_14", false, "", 1},
		{"PL-187_DZ_14", false, "", 1},
		{"PL-187_RR_14", true, "", 1},
		{"PL-187_CE_14", false, "", 1},
		{"PL-187_AC1_16", true, "", 1},
		{"PL-187_AC2_16", false, "", 1},
		{"PL-187_AC3_16", false, "", 1},
		{"PL-187_DZ_16", false, "", 1},
		{"PL-187_RR_16", false, "", 1},
		{"PL-187_CE_16", false, "", 1},
		{"PL-187_AC1_20", false, "", 1},
		{"PL-187_AC2_20", false, "", 1},
		{"PL-187_AC3_20", true, "", 1},
		{"PL-187_DZ_20", false, "", 1},
		{"PL-187_RR_20", false, "", 1},
		{"PL-187_CE_20", false, "", 1},
		{"PV-240_AC1_10", false, "", 1},
		{"PV-240_AC2_10", true, "", 1},
		{"PV-240_AC3_10", false, "", 1},
		{"PV-240_DZ_10", false, "", 1},
		{"PV-240_RR_10", false, "", 1},
		{"PV-240_CE_10", false, "", 1},
		{"PV-240_AC1_14", false, "", 1},
		{"PV-240_AC2_14", false, "", 1},
		{"PV-240_AC3_14", false, "", 1},
		{"PV-240_DZ_14", true, "", 1},
		{"PV-240_RR_14", false, "", 1},
		{"PV-240_CE_14", false, "", 1},
		{"PV-240_AC1_16", false, "", 1},
		{"PV-240_AC2_16", false, "", 1},
		{"PV-240_AC3_16", false, "", 1},
		{"PV-240_DZ_16", false, "", 1},
		{"PV-240_RR_16", true, "", 1},
		{"PV-240_CE_16", false, "", 1},
		{"PV-240_AC1_20", false, "", 1},
		{"PV-240_AC2_20", false, "", 1},
		{"PV-240_AC3_20", false, "", 1},
		{"PV-240_DZ_20", false, "", 1},
		{"PV-240_RR_20", false, "", 1},
		{"PV-240_CE_20", true, "", 1},
		{"PO-270_AC1_10", false, "", 1},
		{"PO-270_AC2_10", true, "", 1},
		{"PO-270_AC3_10", false, "", 1},
		{"PO-270_DZ_10", false, "", 1},
		{"PO-270_RR_10", false, "", 1},
		{"PO-270_CE_10", false, "", 1},
		{"PO-270_AC1_14", false, "", 1},
		{"PO-270_AC2_14", false, "", 1},
		{"PO-270_AC3_14", false, "", 1},
		{"PO-270_DZ_14", false, "", 1},
		{"PO-270_RR_14", false, "", 1},
		{"PO-270_CE_14", true, "", 1},
		{"PO-270_AC1_16", false, "", 1},
		{"PO-270_AC2_16", false, "", 1},
		{"PO-270_AC3_16", false, "", 1},
		{"PO-270_DZ_16", false, "", 1},
		{"PO-270_RR_16", false, "", 1},
		{"PO-270_CE_16", true, "", 1},
		{"PO-270_AC1_20", false, "", 1},
		{"PO-270_AC2_20", false, "", 1},
		{"PO-270_AC3_20", false, "", 1},
		{"PO-270_DZ_20", false, "", 1},
		{"PO-270_RR_20", false, "", 1},
		{"PO-270_CE_20", true, "", 1},
		{"PM-58_AC1_10", false, "", 1},
		{"PM-58_AC2_10", false, "", 1},
		{"PM-58_AC3_10", false, "", 1},
		{"PM-58_DZ_10", false, "", 1},
		{"PM-58_RR_10", false, "", 1},
		{"PM-58_CE_10", true, "", 1},
		{"PM-58_AC1_14", false, "", 1},
		{"PM-58_AC2_14", false, "", 1},
		{"PM-58_AC3_14", false, "", 1},
		{"PM-58_DZ_14", true, "", 1},
		{"PM-58_RR_14", false, "", 1},
		{"PM-58_CE_14", false, "", 1},
		{"PM-58_AC1_16", false, "", 1},
		{"PM-58_AC2_16", false, "", 1},
		{"PM-58_AC3_16", false, "", 1},
		{"PM-58_DZ_16", false, "", 1},
		{"PM-58_RR_16", false, "", 1},
		{"PM-58_CE_16", true, "", 1},
		{"PM-58_AC1_20", true, "", 1},
		{"PM-58_AC2_20", false, "", 1},
		{"PM-58_AC3_20", false, "", 1},
		{"PM-58_DZ_20", false, "", 1},
		{"PM-58_RR_20", false, "", 1},
		{"PM-58_CE_20", false, "", 1},
		{"PE-66_AC1_10", false, "", 1},
		{"PE-66_AC2_10", false, "", 1},
		{"PE-66_AC3_10", true, "", 1},
		{"PE-66_DZ_10", false, "", 1},
		{"PE-66_RR_10", false, "", 1},
		{"PE-66_CE_10", false, "", 1},
		{"PE-66_AC1_14", true, "", 1},
		{"PE-66_AC2_14", false, "", 1},
		{"PE-66_AC3_14", false, "", 1},
		{"PE-66_DZ_14", false, "", 1},
		{"PE-66_RR_14", false, "", 1},
		{"PE-66_CE_14", false, "", 1},
		{"PE-66_AC1_16", false, "", 1},
		{"PE-66_AC2_16", true, "", 1},
		{"PE-66_AC3_16", false, "", 1},
		{"PE-66_DZ_16", false, "", 1},
		{"PE-66_RR_16", false, "", 1},
		{"PE-66_CE_16", false, "", 1},
		{"PE-66_AC1_20", false, "", 1},
		{"PE-66_AC2_20", true, "", 1},
		{"PE-66_AC3_20", false, "", 1},
		{"PE-66_DZ_20", false, "", 1},
		{"PE-66_RR_20", false, "", 1},
		{"PE-66_CE_20", false, "", 1},
		{"PW-168_AC1_10", false, "", 1},
		{"PW-168_AC2_10", false, "", 1},
		{"PW-168_AC3_10", true, "", 1},
		{"PW-168_DZ_10", false, "", 1},
		{"PW-168_RR_10", false, "", 1},
		{"PW-168_CE_10", false, "", 1},
		{"PW-168_AC1_14", false, "", 1},
		{"PW-168_AC2_14", false, "", 1},
		{"PW-168_AC3_14", true, "", 1},
		{"PW-168_DZ_14", false, "", 1},
		{"PW-168_RR_14", false, "", 1},
		{"PW-168_CE_14", false, "", 1},
		{"PW-168_AC1_16", false, "", 1},
		{"PW-168_AC2_16", false, "", 1},
		{"PW-168_AC3_16", false, "", 1},
		{"PW-168_DZ_16", false, "", 1},
		{"PW-168_RR_16", false, "", 1},
		{"PW-168_CE_16", true, "", 1},
		{"PW-168_AC1_20", false, "", 1},
		{"PW-168_AC2_20", true, "", 1},
		{"PW-168_AC3_20", false, "", 1},
		{"PW-168_DZ_20", false, "", 1},
		{"PW-168_RR_20", false, "", 1},
		{"PW-168_CE_20", false, "", 1},
		{"PC-263_AC1_10", true, "", 1},
		{"PC-263_AC2_10", false, "", 1},
		{"PC-263_AC3_10", false, "", 1},
		{"PC-263_DZ_10", false, "", 1},
		{"PC-263_RR_10", false, "", 1},
		{"PC-263_CE_10", false, "", 1},
		{"PC-263_AC1_14", true, "", 1},
		{"PC-263_AC2_14", false, "", 1},
		{"PC-263_AC3_14", false, "", 1},
		{"PC-263_DZ_14", false, "", 1},
		{"PC-263_RR_14", false, "", 1},
		{"PC-263_CE_14", false, "", 1},
		{"PC-263_AC1_16", true, "", 1},
		{"PC-263_AC2_16", false, "", 1},
		{"PC-263_AC3_16", false, "", 1},
		{"PC-263_DZ_16", false, "", 1},
		{"PC-263_RR_16", false, "", 1},
		{"PC-263_CE_16", false, "", 1},
		{"PC-263_AC1_20", false, "", 1},
		{"PC-263_AC2_20", true, "", 1},
		{"PC-263_AC3_20", false, "", 1},
		{"PC-263_DZ_20", false, "", 1},
		{"PC-263_RR_20", false, "", 1},
		{"PC-263_CE_20", false, "", 1},
		{"PX-217_AC1_10", true, "", 1},
		{"PX-217_AC2_10", false, "", 1},
		{"PX-217_AC3_10", false, "", 1},
		{"PX-217_DZ_10", false, "", 1},
		{"PX-217_RR_10", false, "", 1},
		{"PX-217_CE_10", false, "", 1},
		{"PX-217_AC1_14", true, "", 1},
		{"PX-217_AC2_14", false, "", 1},
		{"PX-217_AC3_14", false, "", 1},
		{"PX-217_DZ_14", false, "", 1},
		{"PX-217_RR_14", false, "", 1},
		{"PX-217_CE_14", false, "", 1},
		{"PX-217_AC1_16", false, "", 1},
		{"PX-217_AC2_16", false, "", 1},
		{"PX-217_AC3_16", false, "", 1},
		{"PX-217_DZ_16", false, "", 1},
		{"PX-217_RR_16", false, "", 1},
		{"PX-217_CE_16", true, "", 1},
		{"PX-217_AC1_20", false, "", 1},
		{"PX-217_AC2_20", false, "", 1},
		{"PX-217_AC3_20", false, "", 1},
		{"PX-217_DZ_20", false, "", 1},
		{"PX-217_RR_20", false, "", 1},
		{"PX-217_CE_20", true, "", 1},
		{"PI-40_AC1_10", false, "", 1},
		{"PI-40_AC2_10", false, "", 1},
		{"PI-40_AC3_10", false, "", 1},
		{"PI-40_DZ_10", false, "", 1},
		{"PI-40_RR_10", true, "", 1},
		{"PI-40_CE_10", false, "", 1},
		{"PI-40_AC1_14", false, "", 1},
		{"PI-40_AC2_14", false, "", 1},
		{"PI-40_AC3_14", false, "", 1},
		{"PI-40_DZ_14", false, "", 1},
		{"PI-40_RR_14", false, "", 1},
		{"PI-40_CE_14", true, "", 1},
		{"PI-40_AC1_16", true, "", 1},
		{"PI-40_AC2_16", false, "", 1},
		{"PI-40_AC3_16", false, "", 1},
		{"PI-40_DZ_16", false, "", 1},
		{"PI-40_RR_16", false, "", 1},
		{"PI-40_CE_16", false, "", 1},
		{"PI-40_AC1_20", false, "", 1},
		{"PI-40_AC2_20", false, "", 1},
		{"PI-40_AC3_20", false, "", 1},
		{"PI-40_DZ_20", false, "", 1},
		{"PI-40_RR_20", true, "", 1},
		{"PI-40_CE_20", false, "", 1},
		{"PK-4_AC1_10", false, "", 1},
		{"PK-4_AC2_10", false, "", 1},
		{"PK-4_AC3_10", false, "", 1},
		{"PK-4_DZ_10", false, "", 1},
		{"PK-4_RR_10", true, "", 1},
		{"PK-4_CE_10", false, "", 1},
		{"PK-4_AC1_14", false, "", 1},
		{"PK-4_AC2_14", true, "", 1},
		{"PK-4_AC3_14", false, "", 1},
		{"PK-4_DZ_14", false, "", 1},
		{"PK-4_RR_14", false, "", 1},
		{"PK-4_CE_14", false, "", 1},
		{"PK-4_AC1_16", true, "", 1},
		{"PK-4_AC2_16", false, "", 1},
		{"PK-4_AC3_16", false, "", 1},
		{"PK-4_DZ_16", false, "", 1},
		{"PK-4_RR_16", false, "", 1},
		{"PK-4_CE_16", false, "", 1},
		{"PK-4_AC1_20", false, "", 1},
		{"PK-4_AC2_20", false, "", 1},
		{"PK-4_AC3_20", false, "", 1},
		{"PK-4_DZ_20", false, "", 1},
		{"PK-4_RR_20", false, "", 1},
		{"PK-4_CE_20", true, "", 1},
		{"PC-52_AC1_10", false, "", 1},
		{"PC-52_AC2_10", false, "", 1},
		{"PC-52_AC3_10", true, "", 1},
		{"PC-52_DZ_10", false, "", 1},
		{"PC-52_RR_10", false, "", 1},
		{"PC-52_CE_10", false, "", 1},
		{"PC-52_AC1_14", false, "", 1},
		{"PC-52_AC2_14", false, "", 1},
		{"PC-52_AC3_14", false, "", 1},
		{"PC-52_DZ_14", true, "", 1},
		{"PC-52_RR_14", false, "", 1},
		{"PC-52_CE_14", false, "", 1},
		{"PC-52_AC1_16", false, "", 1},
		{"PC-52_AC2_16", true, "", 1},
		{"PC-52_AC3_16", false, "", 1},
		{"PC-52_DZ_16", false, "", 1},
		{"PC-52_RR_16", false, "", 1},
		{"PC-52_CE_16", false, "", 1},
		{"PC-52_AC1_20", true, "", 1},
		{"PC-52_AC2_20", false, "", 1},
		{"PC-52_AC3_20", false, "", 1},
		{"PC-52_DZ_20", false, "", 1},
		{"PC-52_RR_20", false, "", 1},
		{"PC-52_CE_20", false, "", 1},
		{"PA-21_AC1_10", false, "", 1},
		{"PA-21_AC2_10", false, "", 1},
		{"PA-21_AC3_10", true, "", 1},
		{"PA-21_DZ_10", false, "", 1},
		{"PA-21_RR_10", false, "", 1},
		{"PA-21_CE_10", false, "", 1},
		{"PA-21_AC1_14", false, "", 1},
		{"PA-21_AC2_14", false, "", 1},
		{"PA-21_AC3_14", false, "", 1},
		{"PA-21_DZ_14", false, "", 1},
		{"PA-21_RR_14", true, "", 1},
		{"PA-21_CE_14", false, "", 1},
		{"PA-21_AC1_16", false, "", 1},
		{"PA-21_AC2_16", false, "", 1},
		{"PA-21_AC3_16", false, "", 1},
		{"PA-21_DZ_16", true, "", 1},
		{"PA-21_RR_16", false, "", 1},
		{"PA-21_CE_16", false, "", 1},
		{"PA-21_AC1_20", false, "", 1},
		{"PA-21_AC2_20", false, "", 1},
		{"PA-21_AC3_20", false, "", 1},
		{"PA-21_DZ_20", false, "", 1},
		{"PA-21_RR_20", false, "", 1},
		{"PA-21_CE_20", true, "", 1},
		{"PL-25_AC1_10", false, "", 1},
		{"PL-25_AC2_10", false, "", 1},
		{"PL-25_AC3_10", false, "", 1},
		{"PL-25_DZ_10", false, "", 1},
		{"PL-25_RR_10", true, "", 1},
		{"PL-25_CE_10", false, "", 1},
		{"PL-25_AC1_14", false, "", 1},
		{"PL-25_AC2_14", false, "", 1},
		{"PL-25_AC3_14", true, "", 1},
		{"PL-25_DZ_14", false, "", 1},
		{"PL-25_RR_14", false, "", 1},
		{"PL-25_CE_14", false, "", 1},
		{"PL-25_AC1_16", false, "", 1},
		{"PL-25_AC2_16", false, "", 1},
		{"PL-25_AC3_16", true, "", 1},
		{"PL-25_DZ_16", false, "", 1},
		{"PL-25_RR_16", false, "", 1},
		{"PL-25_CE_16", false, "", 1},
		{"PL-25_AC1_20", false, "", 1},
		{"PL-25_AC2_20", false, "", 1},
		{"PL-25_AC3_20", true, "", 1},
		{"PL-25_DZ_20", false, "", 1},
		{"PL-25_RR_20", false, "", 1},
		{"PL-25_CE_20", false, "", 1},
		{"PG-11_AC1_10", false, "", 1},
		{"PG-11_AC2_10", false, "", 1},
		{"PG-11_AC3_10", false, "", 1},
		{"PG-11_DZ_10", false, "", 1},
		{"PG-11_RR_10", true, "", 1},
		{"PG-11_CE_10", false, "", 1},
		{"PG-11_AC1_14", true, "", 1},
		{"PG-11_AC2_14", false, "", 1},
		{"PG-11_AC3_14", false, "", 1},
		{"PG-11_DZ_14", false, "", 1},
		{"PG-11_RR_14", false, "", 1},
		{"PG-11_CE_14", false, "", 1},
		{"PG-11_AC1_16", false, "", 1},
		{"PG-11_AC2_16", false, "", 1},
		{"PG-11_AC3_16", false, "", 1},
		{"PG-11_DZ_16", false, "", 1},
		{"PG-11_RR_16", true, "", 1},
		{"PG-11_CE_16", false, "", 1},
		{"PG-11_AC1_20", false, "", 1},
		{"PG-11_AC2_20", false, "", 1},
		{"PG-11_AC3_20", false, "", 1},
		{"PG-11_DZ_20", false, "", 1},
		{"PG-11_RR_20", false, "", 1},
		{"PG-11_CE_20", true, "", 1},
		{"PN-144_AC1_10", false, "", 1},
		{"PN-144_AC2_10", false, "", 1},
		{"PN-144_AC3_10", false, "", 1},
		{"PN-144_DZ_10", true, "", 1},
		{"PN-144_RR_10", false, "", 1},
		{"PN-144_CE_10", false, "", 1},
		{"PN-144_AC1_14", false, "", 1},
		{"PN-144_AC2_14", false, "", 1},
		{"PN-144_AC3_14", true, "", 1},
		{"PN-144_DZ_14", false, "", 1},
		{"PN-144_RR_14", false, "", 1},
		{"PN-144_CE_14", false, "", 1},
		{"PN-144_AC1_16", false, "", 1},
		{"PN-144_AC2_16", true, "", 1},
		{"PN-144_AC3_16", false, "", 1},
		{"PN-144_DZ_16", false, "", 1},
		{"PN-144_RR_16", false, "", 1},
		{"PN-144_CE_16", false, "", 1},
		{"PN-144_AC1_20", false, "", 1},
		{"PN-144_AC2_20", true, "", 1},
		{"PN-144_AC3_20", false, "", 1},
		{"PN-144_DZ_20", false, "", 1},
		{"PN-144_RR_20", false, "", 1},
		{"PN-144_CE_20", false, "", 1},
		{"PH-148_AC1_10", false, "", 1},
		{"PH-148_AC2_10", false, "", 1},
		{"PH-148_AC3_10", false, "", 1},
		{"PH-148_DZ_10", true, "", 1},
		{"PH-148_RR_10", false, "", 1},
		{"PH-148_CE_10", false, "", 1},
		{"PH-148_AC1_14", false, "", 1},
		{"PH-148_AC2_14", true, "", 1},
		{"PH-148_AC3_14", false, "", 1},
		{"PH-148_DZ_14", false, "", 1},
		{"PH-148_RR_14", false, "", 1},
		{"PH-148_CE_14", false, "", 1},
		{"PH-148_AC1_16", false, "", 1},
		{"PH-148_AC2_16", false, "", 1},
		{"PH-148_AC3_16", true, "", 1},
		{"PH-148_DZ_16", false, "", 1},
		{"PH-148_RR_16", false, "", 1},
		{"PH-148_CE_16", false, "", 1},
		{"PH-148_AC1_20", true, "", 1},
		{"PH-148_AC2_20", false, "", 1},
		{"PH-148_AC3_20", false, "", 1},
		{"PH-148_DZ_20", false, "", 1},
		{"PH-148_RR_20", false, "", 1},
		{"PH-148_CE_20", false, "", 1},
		{"PR-8_AC1_10", false, "", 1},
		{"PR-8_AC2_10", false, "", 1},
		{"PR-8_AC3_10", true, "", 1},
		{"PR-8_DZ_10", false, "", 1},
		{"PR-8_RR_10", false, "", 1},
		{"PR-8_CE_10", false, "", 1},
		{"PR-8_AC1_14", false, "", 1},
		{"PR-8_AC2_14", false, "", 1},
		{"PR-8_AC3_14", false, "", 1},
		{"PR-8_DZ_14", false, "", 1},
		{"PR-8_RR_14", false, "", 1},
		{"PR-8_CE_14", true, "", 1},
		{"PR-8_AC1_16", false, "", 1},
		{"PR-8_AC2_16", false, "", 1},
		{"PR-8_AC3_16", false, "", 1},
		{"PR-8_DZ_16", false, "", 1},
		{"PR-8_RR_16", true, "", 1},
		{"PR-8_CE_16", false, "", 1},
		{"PR-8_AC1_20", false, "", 1},
		{"PR-8_AC2_20", true, "", 1},
		{"PR-8_AC3_20", false, "", 1},
		{"PR-8_DZ_20", false, "", 1},
		{"PR-8_RR_20", false, "", 1},
		{"PR-8_CE_20", false, "", 1},
		{"PV-35_AC1_10", false, "", 1},
		{"PV-35_AC2_10", false, "", 1},
		{"PV-35_AC3_10", false, "", 1},
		{"PV-35_DZ_10", false, "", 1},
		{"PV-35_RR_10", false, "", 1},
		{"PV-35_CE_10", true, "", 1},
		{"PV-35_AC1_14", false, "", 1},
		{"PV-35_AC2_14", false, "", 1},
		{"PV-35_AC3_14", true, "", 1},
		{"PV-35_DZ_14", false, "", 1},
		{"PV-35_RR_14", false, "", 1},
		{"PV-35_CE_14", false, "", 1},
		{"PV-35_AC1_16", false, "", 1},
		{"PV-35_AC2_16", false, "", 1},
		{"PV-35_AC3_16", false, "", 1},
		{"PV-35_DZ_16", false, "", 1},
		{"PV-35_RR_16", false, "", 1},
		{"PV-35_CE_16", true, "", 1},
		{"PV-35_AC1_20", true, "", 1},
		{"PV-35_AC2_20", false, "", 1},
		{"PV-35_AC3_20", false, "", 1},
		{"PV-35_DZ_20", false, "", 1},
		{"PV-35_RR_20", false, "", 1},
		{"PV-35_CE_20", false, "", 1},
		{"PH-169_AC1_10", false, "", 1},
		{"PH-169_AC2_10", false, "", 1},
		{"PH-169_AC3_10", false, "", 1},
		{"PH-169_DZ_10", true, "", 1},
		{"PH-169_RR_10", false, "", 1},
		{"PH-169_CE_10", false, "", 1},
		{"PH-169_AC1_14", false, "", 1},
		{"PH-169_AC2_14", false, "", 1},
		{"PH-169_AC3_14", false, "", 1},
		{"PH-169_DZ_14", false, "", 1},
		{"PH-169_RR_14", false, "", 1},
		{"PH-169_CE_14", true, "", 1},
		{"PH-169_AC1_16", false, "", 1},
		{"PH-169_AC2_16", true, "", 1},
		{"PH-169_AC3_16", false, "", 1},
		{"PH-169_DZ_16", false, "", 1},
		{"PH-169_RR_16", false, "", 1},
		{"PH-169_CE_16", false, "", 1},
		{"PH-169_AC1_20", true, "", 1},
		{"PH-169_AC2_20", false, "", 1},
		{"PH-169_AC3_20", false, "", 1},
		{"PH-169_DZ_20", false, "", 1},
		{"PH-169_RR_20", false, "", 1},
		{"PH-169_CE_20", false, "", 1},
		{"PO-276_AC1_10", false, "", 1},
		{"PO-276_AC2_10", false, "", 1},
		{"PO-276_AC3_10", false, "", 1},
		{"PO-276_DZ_10", false, "", 1},
		{"PO-276_RR_10", true, "", 1},
		{"PO-276_CE_10", false, "", 1},
		{"PO-276_AC1_14", false, "", 1},
		{"PO-276_AC2_14", false, "", 1},
		{"PO-276_AC3_14", false, "", 1},
		{"PO-276_DZ_14", false, "", 1},
		{"PO-276_RR_14", false, "", 1},
		{"PO-276_CE_14", true, "", 1},
		{"PO-276_AC1_16", false, "", 1},
		{"PO-276_AC2_16", true, "", 1},
		{"PO-276_AC3_16", false, "", 1},
		{"PO-276_DZ_16", false, "", 1},
		{"PO-276_RR_16", false, "", 1},
		{"PO-276_CE_16", false, "", 1},
		{"PO-276_AC1_20", false, "", 1},
		{"PO-276_AC2_20", false, "", 1},
		{"PO-276_AC3_20", true, "", 1},
		{"PO-276_DZ_20", false, "", 1},
		{"PO-276_RR_20", false, "", 1},
		{"PO-276_CE_20", false, "", 1},
		{"PO-297_AC1_10", false, "", 1},
		{"PO-297_AC2_10", true, "", 1},
		{"PO-297_AC3_10", false, "", 1},
		{"PO-297_DZ_10", false, "", 1},
		{"PO-297_RR_10", false, "", 1},
		{"PO-297_CE_10", false, "", 1},
		{"PO-297_AC1_14", true, "", 1},
		{"PO-297_AC2_14", false, "", 1},
		{"PO-297_AC3_14", false, "", 1},
		{"PO-297_DZ_14", false, "", 1},
		{"PO-297_RR_14", false, "", 1},
		{"PO-297_CE_14", false, "", 1},
		{"PO-297_AC1_16", false, "", 1},
		{"PO-297_AC2_16", false, "", 1},
		{"PO-297_AC3_16", false, "", 1},
		{"PO-297_DZ_16", false, "", 1},
		{"PO-297_RR_16", false, "", 1},
		{"PO-297_CE_16", true, "", 1},
		{"PO-297_AC1_20", false, "", 1},
		{"PO-297_AC2_20", false, "", 1},
		{"PO-297_AC3_20", false, "", 1},
		{"PO-297_DZ_20", false, "", 1},
		{"PO-297_RR_20", true, "", 1},
		{"PO-297_CE_20", false, "", 1},
		{"PJ-1_AC1_10", true, "", 1},
		{"PJ-1_AC2_10", false, "", 1},
		{"PJ-1_AC3_10", false, "", 1},
		{"PJ-1_DZ_10", false, "", 1},
		{"PJ-1_RR_10", false, "", 1},
		{"PJ-1_CE_10", false, "", 1},
		{"PJ-1_AC1_14", false, "", 1},
		{"PJ-1_AC2_14", false, "", 1},
		{"PJ-1_AC3_14", false, "", 1},
		{"PJ-1_DZ_14", true, "", 1},
		{"PJ-1_RR_14", false, "", 1},
		{"PJ-1_CE_14", false, "", 1},
		{"PJ-1_AC1_16", false, "", 1},
		{"PJ-1_AC2_16", true, "", 1},
		{"PJ-1_AC3_16", false, "", 1},
		{"PJ-1_DZ_16", false, "", 1},
		{"PJ-1_RR_16", false, "", 1},
		{"PJ-1_CE_16", false, "", 1},
		{"PJ-1_AC1_20", false, "", 1},
		{"PJ-1_AC2_20", false, "", 1},
		{"PJ-1_AC3_20", true, "", 1},
		{"PJ-1_DZ_20", false, "", 1},
		{"PJ-1_RR_20", false, "", 1},
		{"PJ-1_CE_20", false, "", 1},
		{"PW-12_AC1_10", false, "", 1},
		{"PW-12_AC2_10", false, "", 1},
		{"PW-12_AC3_10", false, "", 1},
		{"PW-12_DZ_10", false, "", 1},
		{"PW-12_RR_10", false, "", 1},
		{"PW-12_CE_10", true, "", 1},
		{"PW-12_AC1_14", true, "", 1},
		{"PW-12_AC2_14", false, "", 1},
		{"PW-12_AC3_14", false, "", 1},
		{"PW-12_DZ_14", false, "", 1},
		{"PW-12_RR_14", false, "", 1},
		{"PW-12_CE_14", false, "", 1},
		{"PW-12_AC1_16", false, "", 1},
		{"PW-12_AC2_16", false, "", 1},
		{"PW-12_AC3_16", false, "", 1},
		{"PW-12_DZ_16", true, "", 1},
		{"PW-12_RR_16", false, "", 1},
		{"PW-12_CE_16", false, "", 1},
		{"PW-12_AC1_20", false, "", 1},
		{"PW-12_AC2_20", false, "", 1},
		{"PW-12_AC3_20", false, "", 1},
		{"PW-12_DZ_20", false, "", 1},
		{"PW-12_RR_20", false, "", 1},
		{"PW-12_CE_20", true, "", 1},
	},
}

var kinmatesDesc = `
      ___           ___           ___           ___           ___         ___     
     /__/|         /  /\         /  /\         /__/\         /  /\       /  /\    
    |  |:|        /  /::\       /  /::\       |  |::\       /  /::\     /  /:/_   
    |  |:|       /  /:/\:\     /  /:/\:\      |  |:|:\     /  /:/\:\   /  /:/ /\  
  __|  |:|      /  /:/~/:/    /  /:/~/::\   __|__|:|\:\   /  /:/~/:/  /  /:/ /::\ 
 /__/\_|:|____ /__/:/ /:/___ /__/:/ /:/\:\ /__/::::| \:\ /__/:/ /:/  /__/:/ /:/\:\
 \  \:\/:::::/ \  \:\/:::::/ \  \:\/:/__\/ \  \:\~~\__\/ \  \:\/:/   \  \:\/:/~/:/
  \  \::/~~~~   \  \::/~~~~   \  \::/       \  \:\        \  \::/     \  \::/ /:/ 
   \  \:\        \  \:\        \  \:\        \  \:\        \  \:\      \__\/ /:/  
    \  \:\        \  \:\        \  \:\        \  \:\        \  \:\       /__/:/   
     \__\/         \__\/         \__\/         \__\/         \__\/       \__\/    
                                                                 
  [-> Serveur Prisonniers <-]    ++ toutes les transaction sont loguées (SecLvl 4)
`

// serveur de sécurité de la kramps
var kramps_sec = Server{
	Address: "sec.kramps.d22.eu",
	Credentials: []Cred{
		{"admin", "lkjqsod", 5},
	},
	Scan:        SEC4,
	Description: ksecDesc,
	Registers: []Register{
		// caméras dans les ateliers (AC1/2/3)
		{"CAM-AC1-01", true, "", 5},
		{"CAM-AC1-02", true, "", 5},
		{"CAM-AC1-03", true, "", 5},
		{"CAM-AC2-01", true, "", 5},
		{"CAM-AC2-02", true, "", 5},
		{"CAM-AC2-03", true, "", 5},
		{"CAM-AC3-01", true, "", 5},
		{"CAM-AC3-02", true, "", 5},
		{"CAM-AC3-03", true, "", 5},
		// caméras dans la zone de repos
		{"CAM-RR-01", true, "", 5},
		{"CAM-RR-02", true, "", 5},
		{"CAM-RR-03", true, "", 5},
		{"CAM-RR-04", true, "", 5},
		{"CAM-RR-05", true, "", 5},
		// caméras dans les couloirs des cellules
		{"CAM-CE-01", true, "", 5},
		{"CAM-CE-02", true, "", 5},
		{"CAM-CE-03", true, "", 5},
		{"CAM-CE-04", true, "", 5},
		{"CAM-CE-05", true, "", 5},
		// verrouillage des portes des ateliers
		{"VER-AC1", true, "", 5},
		{"VER-AC2", true, "", 5},
		{"VER-AC3", true, "", 5},
		// verrouillage  des portes de la zone de repos
		{"VER-RR-01", true, "", 5},
		{"VER-RR-02", true, "", 5},
		{"VER-RR-03", true, "", 5},
		// verrouillage des portes des couloirs des cellules
		{"VER-CE-01", true, "", 5},
		{"VER-CE-02", true, "", 5},
		{"VER-CE-03", true, "", 5},
		{"VER-CE-04", true, "", 5},
		{"VER-CE-05", true, "", 5},
		// caméra de la pièce ou il y a le coffre
		{"CAM-DIV-01", true, "", 5},
		// alarmes des divers secteurs
		{"ALM-AC1", false, "", 5},
		{"ALM-AC2", false, "", 5},
		{"ALM-AC3", false, "", 5},
		{"ALM-RR", false, "", 5},
		{"ALM-CE", false, "", 5},
	},
}

var ksecDesc = `
      ___           ___           ___           ___           ___         ___     
     /__/|         /  /\         /  /\         /__/\         /  /\       /  /\    
    |  |:|        /  /::\       /  /::\       |  |::\       /  /::\     /  /:/_   
    |  |:|       /  /:/\:\     /  /:/\:\      |  |:|:\     /  /:/\:\   /  /:/ /\  
  __|  |:|      /  /:/~/:/    /  /:/~/::\   __|__|:|\:\   /  /:/~/:/  /  /:/ /::\ 
 /__/\_|:|____ /__/:/ /:/___ /__/:/ /:/\:\ /__/::::| \:\ /__/:/ /:/  /__/:/ /:/\:\
 \  \:\/:::::/ \  \:\/:::::/ \  \:\/:/__\/ \  \:\~~\__\/ \  \:\/:/   \  \:\/:/~/:/
  \  \::/~~~~   \  \::/~~~~   \  \::/       \  \:\        \  \::/     \  \::/ /:/ 
   \  \:\        \  \:\        \  \:\        \  \:\        \  \:\      \__\/ /:/  
    \  \:\        \  \:\        \  \:\        \  \:\        \  \:\       /__/:/   
     \__\/         \__\/         \__\/         \__\/         \__\/       \__\/    
                                                                 
  [-> Serveur Sécurité <-]        ++ toutes les transaction sont loguées (SecLvl 4)
  
  Vous avez suivi la formation <sécurité> obligatoire, mais nous vous rappelons les
  règles essentielles :
       #1) Respectez la vie privée des autres.
       #2) Réfléchissez avant de taper.
       #3) De grands pouvoirs impliquent de grandes responsabilités.
`

// serveur des services corporatistes D22
var corp = Server{
	Address: "corp.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: cd22Desc,
	Targets: []Target{
		{justice.Address, "services judiciaires", 1, "public", "public"},
	},
}

var cd22Desc = `

   ((ervices  ((orporatistes
   ''                       
    _                       
   [|)istrict  22           

   Ce service **public** vous est proposé **gratuitement** par la Cour Corporatiste.

   Ce service est livré en l'état, et la Cours Corporatiste décline toute responsabilité
   en ce qui concerne les données présentes et l'usage qui en est fait.

   Ce site existe gràce à la généreuse participation de Weyland-Yutani Corp,
   Tyrel Corp, Tessier-Ashpool SA, Disney Dassault, Arasaka, Renraku, Ubik,
   Legba Voodoocom, Avalon, Association des Banques Unifiées Suisses (ABUS).
`

// serveur judiciaire
var justice = Server{
	Address: "justice.corp.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: cd22justDesc,
	Entries: []Entry{
		{mel.Login, mel.Keywords(), 1, "", "Mélody MATHISON", "Disparue - Incident 16485-4346B, Nexkemia Petrochemicals, 07/07/2000"},
		{rocky.Login, rocky.Keywords(), 1, "", "John DOE 7624", "- D22/ag#867533654: agression à main armée (victime Sony HAARTZ)"},
		{rita.Login, rita.Keywords(), 1, "", "Margherita BELLAMY", "- néant"},
		{styx.Login, styx.Keywords(), 1, "", "Sébastian BRONNER", "- néant"},
		{kapo.Login, kapo.Keywords(), 1, "", "Camélia BELLAMY", "- néant"},
		{scalpel.Login, scalpel.Keywords(), 1, "", "Julius VILLANOVA", "***** Personne recherchée, mandat inter-district PJ/676/ER/65534 *****\n- D22/cm#5674243: complicité de meurtre"},
		{greko.Login, greko.Keywords(), 1, "", "Eddy CANTO", "- néant"},
		{jesus.Login, jesus.Keywords(), 1, "", "Edwin JOHANNESEN", "- néant"},
		{escobar.Login, escobar.Keywords(), 1, "", "Jonathan BRANSON", "- néant TODO"},
		{cageot.Login, cageot.Keywords(), 1, "", "John MacFRIGHT", "***** Personne recherchée, mandat inter-district PF/0865/EP/55463 *****\n- D21/rc#12785234452 rupture contrat\n\n\n$$$SPECIAL$$$ contacter cont4yes@kitsu.Keywordseu, ¥€$ en rapport.Keywords"},
		{lafouine.Login, lafouine.Keywords(), 1, "", "Sylvia Kemija MIHALEC", "- néant"},
		{eva.Login, eva.Keywords(), 1, "", "Pamela TODO", "***** Personne recherchée, mandat inter-district PF/1437/PM/02 *****\n- D21/rc#6542867 rupture contrat"},
		{fatmike.Login, fatmike.Keywords(), 1, "", "Michael DUBIAN", "- D22/vm#23842834: vol à l'étalage\n- D22/vm#54327653: vol recette épicerie nuit\n- D22/vm#543299873: vol simple\n- D22/vm#547699823: vol graviscooter\n- D22/vm#753296671: vol à l'étalage"},
		{kennedy.Login, kennedy.Keywords(), 1, "", "Carlotta MIHALEC", "***** Personne recherchée, mandat inter-district PF/0865/EP/55463 *****\n- D22/vd#765428736: vol données confidentielles "},
		{savagegirl.Login, savagegirl.Keywords(), 1, "", savagegirl.Name, "- néant"},
		{raoulcool.Login, raoulcool.Keywords(), 1, "", "Raoul MICHU", "- néant"},
		{greenglass.Login, greenglass.Keywords(), 1, "", "Rupert GLASS", "- néant"},
		{chillydaisy.Login, chillydaisy.Keywords(), 1, "", "Daisy JOHANNESEN", "***** Personne recherchée, mandat inter-district PF/0415/EG/55323 *****\n- D22/me#1275436253: double meurtre, arme à feu\n"},
		{frereping.Login, frereping.Keywords(), 1, "", "Désiré BONENFANT", "- néant"},
		{papaproxy.Login, papaproxy.Keywords(), 1, "", "Harald PROSKYCHEV", "***** Personne recherchée, mandat inter-district PF/2964/EP/98254 *****\n- D22/vd#89875357678: vol données avec copyright"},
		{nikki.Login, nikki.Keywords(), 1, "", "Nicole JASINSKI", "***** Personne recherchée, mandat inter-district PF/7253/EP/90271 *****\n- D22/vd#1100298735: vol données sous brevet"},
		{celine.Login, celine.Keywords(), 1, "", "Franz-Ferdinand CÉLINE", "***** Personne recherchée, mandat inter-district PF/1001/EP/98682 *****\n- D22/pi#9867356873: piratage informatique\n- D22/am#18763725: association malfaiteurs"},
		{cramille.Login, cramille.Keywords(), 1, "", "Camelia MILLS", "- néant"},
		{tigerdoll.Login, tigerdoll.Keywords(), 1, "", "Lilas SEPTEMBRE", "- néant"},
		{sistermorphine.Login, sistermorphine.Keywords(), 1, "", "Eloïse DUBIAN", "- D22/ou#7578538765: outrage et rébellion, EuroPol\n- D22/va#325363552: vandalisme\n- D22/td#89765363: tapage diurne répété\n- D22/tn#101002543: tapage nocturne"},
		{zilmir.Login, zilmir.Keywords(), 1, "", "Zilmir ABASOLO", "- néant"},
		{bettyb.Login, bettyb.Keywords(), 1, "", "Elisabeth BRANSON", "- néant"},
		{abraham.Login, abraham.Keywords(), 1, "", abraham.Name, "- néant"},
		{crunch.Login, crunch.Keywords(), 1, "", "TODO", "- néant"},
		{onekick.Login, onekick.Keywords(), 1, "", "Rodolphe KIÉVAIN", "- néant\n>>> automated procedure: contact@kramps.Keywordseu | #line>2"},
		{jacob.Login, jacob.Keywords(), 1, "", "Pete TODO", "- néant"},
		{cyrano.Login, cyrano.Keywords(), 1, "", "Adrien JOLIVET", "- néant"},
		{smalljoe.Login, smalljoe.Keywords(), 1, "", "Joseph VAZZANNA", "- néant"},
		{ironmike.Login, ironmike.Keywords(), 1, "", "Mickael KLEBERT", "- néant"},
		{paula.Login, paula.Keywords(), 1, "", "Paula JOLIVET", "- néant"},
		{ringo.Login, ringo.Keywords(), 1, "", "Ringo JOLIVET", "- néant"},
		{georges.Login, georges.Keywords(), 1, "", "Georges CHANG", "- néant"},
		{jeanne.Login, jeanne.Keywords(), 1, "", "Jeanne KOLINSKY", "- néant"},
		{oggy.Login, oggy.Keywords(), 1, "", "Richard WHITE", "- néant"},
	},
}

var cd22justDesc = `
   ((ervices  ((orporatistes
   ''                       
    _                       
   [|)istrict  22   === Département JUDICIAIRE ===

   Ce service **public** vous est proposé **gratuitement** par la Cours Corporatiste.

   Ce service est livré en l'état, et la Cours Corporatiste décline toute responsabilité
   en ce qui concerne les données présentes et l'usage qui en est fait.

   Ce site existe gràce à la généreuse participation de Weyland-Yutani Corp,
   Tyrel Corp, Tessier-Ashpool SA, Disney Dassault, Arasaka, Renraku, Ubik,
   Legba Voodoocom, Avalon, Association des Banques Unifiées Suisses (ABUS).
`

// serveur bancaire du D22
var abus = Server{
	Address: "abus.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
		{alan.Login, alan.Password, 1},
		{mel.Login, mel.Password, 1},
		{rita.Login, rita.Password, 1},
	},
	Description: cd22bankDesc,
	Entries: []Entry{
		{
			ID:         alan.Login,
			Keywords:   []string{"propriété"},
			Restricted: 1,
			Owner:      alan.Login,
			Title:      "Titre immobilier DZ8-7687",
			Content: `DZ8-7687 : immeuble-local commercial, District 22.

2000 : ce titre a fait l'objet d'une offre publique d'achat suite à l'incident Nexkemia. M Mathison n'a pas donné suite.
2019 : M Mathison n'a eu aucune activité enregistrée par nos services depuis 2000. Castle Corp souhaite se porter acquéreur du bien. Conformément au cadre légal concernant les personnes portées disparues, M Mathison ou son héritier a jusqu'au 2020-08-01 pour se manifester. Passée cette date, le titre sera remis en vente, au profit de Castle Corp.`,
		},
	},
}

var cd22bankDesc = `
   ((ervices  ((orporatistes
   ''                       
    _                       
   [|)istrict  22   === Département BANCAIRE ===        

   Ce service **public** est assuré par
                                          _______     ______                 _______ 
                       Association       (  ___  )   (  ___ \   |\     /|   (  ____ \
                                         | (   ) |   | (   ) )  | )   ( |   | (    \/
                       des Banques       | (___) |   | (__/ /   | |   | |   | (_____ 
                                         |  ___  |   |  __ (    | |   | |   (_____  )
                          Unifiées       | (   ) |   | (  \ \   | |   | |         ) |
                                         | )   ( | _ | )___) )_ | (___) | _ /\____) |
                           Suisses       |/     \|(¥)|/ \___/(€)(_______)($)\_______)
`

// serveur public de Legba Voodoocom
var legba = Server{
	Address: "legba.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: lbDesc,
	Targets: []Target{
		{legba_satcom.Address, "division sat-com", 5, "admin", "satcom9876"},
		{legba_archive.Address, "archives", 3, "personnel", "archive6543"},
	},
}
var lbDesc = `
                 ......                 
           .',,,,,,,,,,,,,,,.           
        .;;,'.            .',;;'        
      ':;.   ...  ,odooodkxxdc;;:,      
    .::..;:oOKXXOOXWWWWNXNWN0xd:';c'     Entrez dans une nouvelle réalité
   .c,  :XMWWWWWKkKK0KKdck0Okdol. .c,   
  .c'   'dxkKNWWO,...;kklxdOWWWXl  .c,       Soyez l'acteur du changement
  :;        .:ollc,cxOOooxkXMWO;.   'l.      
 'c.         'dc.';ox0O,lWWWWKl''.   :;            Matérialisez vos rêves
 ;c         .;l:'..;coOxdKNXKc,kXd.  ;: 
 ;c         .,:od:.   'kKOdl:'cdc.   ;:              Obtenez l'impossible
 'c.           .';c;',l0Oolc:c,.     :; 
  :;              .xXN0c',;,'.      'c.                    Et plus encore
  .c,            'xKKd.            .c,  
   .c,     .',;:lKMKxo:;,'..      .c,                   
    .::. 'xKNWWMMMWMMMMWWNXk:   .;c. .____                ___.              
      .:;;,,:clloooooollc:;'. .;:,   |    |    ____   ____\_ |__ _____      
        .,;;,.            .',;;'     |    |  _/ __ \ / ___\| __ \\__  \     
           .',,,,,,,,,,,',,'.        |    |__\  ___// /_/  > \_\ \/ __ \_   
                 .......             |_______ \___  >___  /|___  (____  /   
          ____   ___                         \/   \/_____/     \/     \/ 
          \   \ /   /___   ____   __| _/____   ____   ____  ____   _____  
           \   Y   /  _ \ /  _ \ / __ |/  _ \ /  _ \_/ ___\/  _ \ /     \ 
            \     (  <_> |  <_> ) /_/ (  <_> |  <_> )  \__(  <_> )  Y Y  \
             \___/ \____/ \____/\____ |\____/ \____/ \___  >____/|__|_|  /
                                     \/                  \/            \/
`

// serveur privé de la communication satellite
var legba_satcom = Server{
	Address: "satcom.legba.d22.eu",
	Credentials: []Cred{
		{"admin", "satcom9876", 5},
	},
	Description: satDesc,
}
var satDesc = `
SATCOM, une division externalisée de
 ____                ___  
|    |    ____   ____\_ |__ _____   
|    |  _/ __ \ / ___\| __ \\__  \  
|    |__\  ___// /_/  > \_\ \/ __ \_
|_______ \___  >___  /|___  (____  /
 ____   \/_  \/_____/     \/___  \/ 
 \   \ /   /___   ____   __| _/____   ____   ____  ____   _____           
  \   Y   /  _ \ /  _ \ / __ |/  _ \ /  _ \_/ ___\/  _ \ /     \          
   \     (  <_> |  <_> ) /_/ (  <_> |  <_> )  \__(  <_> )  Y Y  \         
    \___/ \____/ \____/\____ |\____/ \____/ \___  >____/|__|_|  /         
                            \/                  \/            \/          

[Accès Restreint]         >>>>>>> entrez vos identifiants <<<<<<<
`

// serveur archive de Silicon Spirit
var legba_archive = Server{
	Address: "archive.legba.d22.eu",
	Credentials: []Cred{
		{"personnel", "archive6543", 3},
	},
	Description: arcDesc,
}
var arcDesc = `
*********************************************************************************
Legba Voodoocom ne peut être tenu responsable de l'usage et des données stockées.
**** WARNING **** : ce service n'est plus maintenu.
*********************************************************************************
━━━╮╭╮╱╱╱╱╱╱╱╱╱╱╱╭━━━╮╱╱╱╱╱╱╭╮    
┃╭━╮┃┃┃╱╱╱╱╱╱╱╱╱╱╱┃╭━╮┃╱╱╱╱╱╭╯╰╮     Division: R&D, Unité 2772
┃╰━━┳┫┃╭┳━━┳━━┳━╮╱┃╰━━┳━━┳┳━╋╮╭╯               Projets spéciaux
╰━━╮┣┫┃┣┫╭━┫╭╮┃╭╮╮╰━━╮┃╭╮┣┫╭╋┫┃                (dir: A.M)
┃╰━╯┃┃╰┫┃╰━┫╰╯┃┃┃┃┃╰━╯┃╰╯┃┃┃┃┃╰╮
╰━━━┻┻━┻┻━━┻━━┻╯╰╯╰━━━┫╭━┻┻╯╰┻━╯     
╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃
╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╰╯
*********************************************************************************
**** WARNING **** : ce service n'est plus maintenu.
Legba Voodoocom ne peut être tenu responsable de l'usage et des données stockées.
*********************************************************************************
[Beware MalvolentKiIA, hack@45EBG56#EACD M@dJ0k3r;3/4/206]
`

// serveur le bon district
var lbd = Server{
	Address: "lebondistrict.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
}

// green data
var greendata = Server{
	Address: "greendata.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: greenDesc,
}

var greenDesc = `
   
                                   
                             %      ____                        __  __     
                         %%%%%%    /\  _ \                     /\ \/\ \    
                %%%%%%%%%%%%%%%%   \ \ \L\_\  _ __    __     __\ \  \\ \   
            %%%%%%%%%%%%%% %%%%%    \ \ \L_L /\  __\/ __ \ / __ \ \ ,   \  
          %%%%%%%%%%%%% *%%%%%%%     \ \ \/, \ \ \//\  __//\  __/\ \ \ \ \ 
         %%%%%%%%%   %%%%%%%%%%       \ \____/\ \_\\ \____\ \____\\ \_\ \_\
         %%%%   %%%%%%%%%%%%%%         \/___/  \/_/ \/____/\/____/ \/_/\/_/
            .%%%%%%%%%%%%%%%       
          %%%%%%%%%%%%%%%               ____              __    ______     
        %%%                            /\  _ \           /\ \__/\  _  \    
        %%                             \ \ \/\ \     __  \ \ ,_\ \ \L\ \   
                                        \ \ \ \ \  / __ \ \ \ \/\ \  __ \  
       Analyses / Diagnostics            \ \ \_\ \/\ \L\ \_\ \ \_\ \ \/\ \ 
                                          \ \____/\ \__/ \_\\ \__\\ \_\ \_\
       Certifié ISO-56-52-100              \/___/  \/__/\/_/ \/__/ \/_/\/_/


`

var invertedLeaf = `
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@.@@@@@
@@@@@@@@@@@@@@@@@@@@......@@@@
@@@@@@@@@@@................@@@
@@@@@@@..............@.....@@@
@@@@@.............@&.......@@@
@@@@.........@@@..........@@@@
@@@@....@@@..............@@@@@
@@@@@@@@...............@@@@@@@
@@@@@...............@@@@@@@@@@
@@@...@@@@@@@@@@@@@@@@@@@@@@@@
@@@..@@@@@@@@@@@@@@@@@@@@@@@@@
`

// serveur privé de Crunch
var leet = Server{
	Address: "l33t.darknet",
	Credentials: []Cred{
		{"crunch", "hacktheplanet", 5},
	},
	Description: cruDesc,
}

var cruDesc = `

                                                                      
         _/_/_/  _/_/_/    _/    _/  _/      _/    _/_/_/  _/    _/   
      _/        _/    _/  _/    _/  _/_/    _/  _/        _/    _/    
     _/        _/_/_/    _/    _/  _/  _/  _/  _/        _/_/_/_/     
    _/        _/    _/  _/    _/  _/    _/_/  _/        _/    _/      
     _/_/_/  _/    _/    _/_/    _/      _/    _/_/_/  _/    _/       
                                                                      
    is NOT watching you... No need for that.... :o) 


`

// serveur privé de Céline
var lair = Server{
	Address: "celine.darknet",
	Credentials: []Cred{
		{"celine", "waytoocool", 5},
	},
	Description: celDesc,
}
var celDesc = `
  ******               *******                    **    
  **////**   **    **  /**////**                  /**    
 **    //  ************/**    /**  ******   ******/**  **
/**       ///**////**/ /**    /** //////** //**//*/** ** 
/**         /**   /**  /**    /**  *******  /** / /****  
//**    ** ************/**    **  **////**  /**   /**/** 
 //****** ///**////**/ /*******  //********/***   /**//**
  //////    //    //   ///////    //////// ///    //  // 

...... Dernier avertissement ............................
`

// serveur mémoriel de Hope
var hope = Server{
	Address: "hope.local",
	Credentials: []Cred{
		{"hope", "tearsintherain", 5},
	},
	Description: hopDesc,
}

// Project "Hope"
// Dépot mémoriel
// - fenêtre temporelle glissante v12.5
// - compression McVaugh-Korba-Yang
// - contenu infix normalisé
// - (c) A.M
var hopDesc = `
01010000 01110010 01101111 01101010 01100101 01100011 01110100  00100010 01001000 01101111 01110000 01100101 00100010                                                                                                                                                                               
01000100 11101001 01110000 01101111 01110100  01101101 11101001 01101101 01101111 01110010 01101001 01100101 01101100                                                                                                                                                                               
00101101  01100110 01100101 01101110 11101010 01110100 01110010 01100101  01110100 01100101 01101101 01110000 01101111 01110010 01100101 01101100 01101100 01100101  01100111 01101100 01101001 01110011 01110011 01100001 01101110 01110100 01100101  01110110 00110001 00110010 00101110 00110101 
00101101  01100011 01101111 01101101 01110000 01110010 01100101 01110011 01110011 01101001 01101111 01101110  01001101 01100011 01010110 01100001 01110101 01100111 01101000 00101101 01001011 01101111 01110010 01100010 01100001 00101101 01011001 01100001 01101110 01100111                     
00101101  01100011 01101111 01101110 01110100 01100101 01101110 01110101  01101001 01101110 01100110 01101001 01111000  01101110 01101111 01110010 01101101 01100001 01101100 01101001 01110011 11101001                                                                                            
00101101  00101000 01100011 00101001  01000001 00101110 01001101                                                                                                                                                                                                                                    
`

var game = &Game{
	Network: []Server{
		dd,
		d22,
		kramps,
		kramps_priv,
		kramps_sec,
		corp,
		justice,
		abus,
		legba,
		legba_satcom,
		legba_archive,
		lbd,
		greendata,
		leet,
		lair,
		hope,
	},
}
