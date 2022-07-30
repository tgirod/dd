package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	SEC1 = time.Minute * 5
	SEC2 = time.Minute * 3
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
	rocky          = ID{"jdoe7624", "CCGCGCAGAATCATAGCTGT", "John Doe 7624"}
	rita           = ID{"mbellamy", "CAAAGTTCTAGGCATAGGGA", "Margherita Bellamy"}
	styx           = ID{"sbronner", "TTAGCTCGATATCCTAACCC", "Sebastian Bronner"}
	kapo           = ID{"cbellamy", "GAACTGCTTTAGTTGACGGA", "Camélia Bellamy"}
	scalpel        = ID{"jvillanova", "TGAAAGAGACATGATGCCTT", "Julius Villanova"}
	greko          = ID{"ecanto", "TCTGAGGTTTATTGATTTCG", "Eddy Canto"}
	jesus          = ID{"ejohannesen", "TTCGGGATTACTGCGTGCTG", "Edwin Johannesen"}
	escobar        = ID{"jbranson", "GGAGGACACCCCAAACGCAT", "Jonathan Branson"}
	cageot         = ID{"jmfright", "GCCCTTGTCATGTACTTAGT", "John Mac Fright"}
	lafouine       = ID{"skmihalec", "CTGTCACCCAATCTACAGCG", "Sylvia Kemija Mihalec"}
	eva            = ID{"emartin", "CTGTTGTAGTGACATGTTTC", "Eva Martin"}
	fatmike        = ID{"mdubian", "AACCTTGGGCACGGTCGGTA", "Michael Dubian"}
	kennedy        = ID{"cmihalec", "CCCGCGGGCAAAGCTGACAG", "Carlotta Mihalec"}
	savagegirl     = ID{"jdoe", "GGGTCTATAGGTCAAACGGT", "Jane Doe 2645"}
	raoulcool      = ID{"rmichu", "GTCACAAGGTTGTTTAATGG", "Raoul Michu"}
	greenglass     = ID{"rglass", "ATGCCTACCTCCAATGATTA", "Rupert Glass"}
	steffie        = ID{"sglass", "ATCGCTACGTCCATAGACTA", "Steffie Glass"}
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
	crunch         = ID{"fmanson", "TTCAAGCTGAATATGAAAGG", "Frédéric Manson"}
	onekick        = ID{"rkievain", "GTCAAATCTGAGACTCTTGC", "Rodolph Kievain"}
	jacob          = ID{"pdoberty", "TGAAAGAGACAGTATGCCGT", "Pete Doberty"}
	cyrano         = ID{"ajolivet", "TTCGACTGAATGTTTGATGT", "Adrien Jolivet"}
	smalljoe       = ID{"jvazzanna", "TATCGACGCACGGGACTTGG", "Joseph Vazzanna"}
	ironmike       = ID{"mklebert", "CGAGAAATGACAGAGTTGTA", "Mickael Klebert"}
	paula          = ID{"pjolivet", "GGGTGATCTGTTGCCCCCTG", "Paula Jolivet"}
	ringo          = ID{"rjolivet", "AACTGACGGATTCGATCATG", "Ringo Jolivet"}
	georges        = ID{"gchang", "GTTTGCACGGAACATGCAAC", "Georges Chang"}
	jeanne         = ID{"jkolinsky", "GACCCGTATTTCGCTGATTG", "Jeanne Kolinsky"}
	oggy           = ID{"rwhite", "TCAGCTTCTAACGTTCGGGA", "Richard White"}
	anton          = ID{"afrieman", "ACGTTGCAAACCTGGTACGT", "Anton Frieman"}
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
		{"personnel", "123kramps!", 3},
	},
	Targets: []Target{
		{kramps_pers.Address, "Serveur réservé au personnel", 3, "personnel", "123kramps!"},
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
var kramps_pers = Server{
	Address: "priv.kramps.d22.eu",
	Credentials: []Cred{
		{"personnel", "123kramps!", 1}, // accès depuis le serveur public
		{"akremmer", "sexgod22", 3},    // backdoor, vol de compte utilisateur
	},
	Targets: []Target{
		{kramps_inmates.Address, "Gestion des prisonniers", 3, "personnel", "123kramps!"},
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
	},
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
	},
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
	Address: "inmates.kramps.d22.eu",
	Credentials: []Cred{
		{"personnel", "123kramps!", 1}, // accès depuis le serveur public
	},
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
		{mel.Login, mel.Keywords(), 1, "", mel.Name, "Disparue - Incident 16485-4346B, Nexkemia Petrochemicals, 07/07/2000"},
		{rocky.Login, rocky.Keywords(), 1, "", rocky.Name, "- D22/ag#867533654: agression à main armée (victime Sony HAARTZ)"},
		{rita.Login, rita.Keywords(), 1, "", rita.Name, "- néant"},
		{styx.Login, styx.Keywords(), 1, "", styx.Name, "- néant"},
		{kapo.Login, kapo.Keywords(), 1, "", kapo.Name, "- néant"},
		{scalpel.Login, scalpel.Keywords(), 1, "", scalpel.Name, "***** Personne recherchée, mandat inter-district PJ/676/ER/65534 *****\n- D22/cm#5674243: complicité de meurtre"},
		{greko.Login, greko.Keywords(), 1, "", greko.Name, "- néant"},
		{jesus.Login, jesus.Keywords(), 1, "", jesus.Name, "- néant"},
		{escobar.Login, escobar.Keywords(), 1, "", escobar.Name, "- néant"},
		{cageot.Login, cageot.Keywords(), 1, "", cageot.Name, "***** Personne recherchée, mandat inter-district PF/0865/EP/55463 *****\n- D21/rc#12785234452 rupture de contrat\n\n\n$$$SPECIAL$$$ contacter cont4yes@kitsu.d22.eu, ¥€$ en rapport"},
		{lafouine.Login, lafouine.Keywords(), 1, "", lafouine.Name, "- néant"},
		{eva.Login, eva.Keywords(), 1, "", eva.Name, "***** Personne recherchée, mandat inter-district PF/1437/PM/02 *****\n- D21/rc#6542867 rupture contrat"},
		{fatmike.Login, fatmike.Keywords(), 1, "", fatmike.Name, "- D22/vm#23842834: vol à l'étalage\n- D22/vm#54327653: vol recette épicerie nuit\n- D22/vm#543299873: vol simple\n- D22/vm#547699823: vol graviscooter\n- D22/vm#753296671: vol à l'étalage"},
		{kennedy.Login, kennedy.Keywords(), 1, "", kennedy.Name, "***** Personne recherchée, mandat inter-district PF/0865/EP/55463 *****\n- D22/vd#765428736: vol données confidentielles "},
		{savagegirl.Login, savagegirl.Keywords(), 1, "", savagegirl.Name, "- néant"},
		{raoulcool.Login, raoulcool.Keywords(), 1, "", raoulcool.Name, "- néant"},
		{greenglass.Login, greenglass.Keywords(), 1, "", greenglass.Name, "- néant"},
		{chillydaisy.Login, chillydaisy.Keywords(), 1, "", chillydaisy.Name, "***** Personne recherchée, mandat inter-district PF/0415/EG/55323 *****\n- D22/me#1275436253: double meurtre, arme à feu\n"},
		{frereping.Login, frereping.Keywords(), 1, "", frereping.Name, "- néant"},
		{papaproxy.Login, papaproxy.Keywords(), 1, "", papaproxy.Name, "***** Personne recherchée, mandat inter-district PF/2964/EP/98254 *****\n- D22/vd#89875357678: vol données avec copyright"},
		{nikki.Login, nikki.Keywords(), 1, "", nikki.Name, "***** Personne recherchée, mandat inter-district PF/7253/EP/90271 *****\n- D22/vd#1100298735: vol données sous brevet"},
		{celine.Login, celine.Keywords(), 1, "", celine.Name, "***** Personne recherchée, mandat inter-district PF/1001/EP/98682 *****\n- D22/pi#9867356873: piratage informatique\n- D22/am#18763725: association malfaiteurs"},
		{cramille.Login, cramille.Keywords(), 1, "", cramille.Name, "- néant"},
		{tigerdoll.Login, tigerdoll.Keywords(), 1, "", tigerdoll.Name, "- néant"},
		{sistermorphine.Login, sistermorphine.Keywords(), 1, "", sistermorphine.Name, "- D22/ou#7578538765: outrage et rébellion, EuroPol\n- D22/va#325363552: vandalisme\n- D22/td#89765363: tapage diurne répété\n- D22/tn#101002543: tapage nocturne"},
		{zilmir.Login, zilmir.Keywords(), 1, "", zilmir.Name, "- néant"},
		{bettyb.Login, bettyb.Keywords(), 1, "", bettyb.Name, "- néant"},
		{abraham.Login, abraham.Keywords(), 1, "", abraham.Name, "- néant"},
		{crunch.Login, crunch.Keywords(), 1, "", crunch.Name, "- néant"},
		{onekick.Login, onekick.Keywords(), 1, "", onekick.Name, "- néant\n>>> automated procedure: contact@kramps.d22.eu | #line>2"},
		{jacob.Login, jacob.Keywords(), 1, "", jacob.Name, "- néant"},
		{cyrano.Login, cyrano.Keywords(), 1, "", cyrano.Name, "- néant"},
		{smalljoe.Login, smalljoe.Keywords(), 1, "", smalljoe.Name, "- néant"},
		{ironmike.Login, ironmike.Keywords(), 1, "", ironmike.Name, "- néant"},
		{paula.Login, paula.Keywords(), 1, "", paula.Name, "- néant"},
		{ringo.Login, ringo.Keywords(), 1, "", ringo.Name, "- néant"},
		{georges.Login, georges.Keywords(), 1, "", georges.Name, "- néant"},
		{jeanne.Login, jeanne.Keywords(), 1, "", jeanne.Name, "- néant"},
		{oggy.Login, oggy.Keywords(), 1, "", oggy.Name, "- néant"},
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
			ID:         fmt.Sprintf("%s_KI8FEI", alan.Login),
			Keywords:   []string{"titre"},
			Restricted: 1,
			Owner:      alan.Login,
			Title:      "Titre immobilier DZ8-7687",
			Content: `DZ8-7687 : immeuble-local commercial, District 22.

2000 : ce titre a fait l'objet d'une offre publique d'achat suite à l'incident Nexkemia. M Mathison n'a pas donné suite.
2019 : M Mathison n'a eu aucune activité enregistrée par nos services depuis 2000. Castle Corp souhaite se porter acquéreur du bien. Conformément au cadre légal concernant les personnes portées disparues, M Mathison ou son héritier a jusqu'au 2020-08-01 pour se manifester. Passée cette date, le titre sera remis en vente, au profit de Castle Corp.`,
		},
		{
			ID:         fmt.Sprintf("%s_UGAEH9", rita.Login),
			Keywords:   []string{"compte"},
			Restricted: 1,
			Owner:      rita.Login,
			Title:      "Compte personnel de Margherita Bellamy",
			Content:    `Solde courant : 10.000Y€S`,
		},
		{
			ID:         fmt.Sprintf("%s_EASHO6", rita.Login),
			Keywords:   []string{"compte"},
			Restricted: 1,
			Owner:      rita.Login,
			Title:      "Ligne budgétaire pour Margherita Bellamy, Headshot Records",
			Content:    `Solde courant : 10.000Y€S`,
		},
		{
			ID:         "EENG3M",
			Keywords:   []string{"compte"},
			Restricted: 1,
			Owner:      kennedy.Login,
			Title:      "Compte anonyme",
			Content:    `Solde courant : 50.000Y€S \\ Date de création : 2007-05-21`,
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

var (
	legbaPersonnel = Cred{"personnel", "paparezo", 3}
	legbaAdmin     = Cred{"admin", "foh5wuoh", 5}
	lbDesc         = `
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
	satDesc = `
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
	arcDesc = `

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
)

// serveur public de Legba Voodoocom
var legba = Server{
	Address: "legba.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
		legbaPersonnel,
	},
	Description: lbDesc,
	Targets: []Target{
		{legba_satcom.Address, "division sat-com", 3, legbaPersonnel.Login, legbaPersonnel.Password},
		{legba_archive.Address, "archives", 3, legbaPersonnel.Login, legbaPersonnel.Password},
	},
}

// serveur privé de la communication satellite
var legba_satcom = Server{
	Address: "satcom.legba.d22.eu",
	Credentials: []Cred{
		legbaPersonnel,
		legbaAdmin,
	},
	Description: satDesc,
	Scan:        SEC4,
	Entries: []Entry{
		{"GEO-EU-D01", []string{"GEO", "EU", "D01"}, 3, "", "Europole D01", `azimut:222.473862
altitude:57.545902
Europole District D01`},
		{"GEO-EU-D02", []string{"GEO", "EU", "D02"}, 3, "", "Europole D02", `azimut:239.324897
altitude:40.278407
Europole District D02`},
		{"GEO-EU-D03", []string{"GEO", "EU", "D03"}, 3, "", "Europole D03", `azimut:109.366561
altitude:76.071807
Europole District D03`},
		{"GEO-EU-D04", []string{"GEO", "EU", "D04"}, 3, "", "Europole D04", `azimut:329.297929
altitude:27.991250
Europole District D04`},
		{"GEO-EU-D05", []string{"GEO", "EU", "D05"}, 3, "", "Europole D05", `azimut:196.971308
altitude:89.923900
Europole District D05`},
		{"GEO-EU-D06", []string{"GEO", "EU", "D06"}, 3, "", "Europole D06", `azimut:128.213744
altitude:72.986507
Europole District D06`},
		{"GEO-EU-D07", []string{"GEO", "EU", "D07"}, 3, "", "Europole D07", `azimut:188.000197
altitude:25.475341
Europole District D07`},
		{"GEO-EU-D08", []string{"GEO", "EU", "D08"}, 3, "", "Europole D08", `azimut:18.387379
altitude:39.189764
Europole District D08`},
		{"GEO-EU-D09", []string{"GEO", "EU", "D09"}, 3, "", "Europole D09", `azimut:25.997548
altitude:2.442340
Europole District D09`},
		{"GEO-EU-D10", []string{"GEO", "EU", "D10"}, 3, "", "Europole D10", `azimut:63.563148
altitude:15.524259
Europole District D10`},
		{"GEO-EU-D11", []string{"GEO", "EU", "D11"}, 3, "", "Europole D11", `azimut:45.742992
altitude:66.014371
Europole District D11`},
		{"GEO-EU-D12", []string{"GEO", "EU", "D12"}, 3, "", "Europole D12", `azimut:347.864768
altitude:52.184109
Europole District D12`},
		{"GEO-EU-D13", []string{"GEO", "EU", "D13"}, 3, "", "Europole D13", `azimut:243.115459
altitude:82.627421
Europole District D13`},
		{"GEO-EU-D14", []string{"GEO", "EU", "D14"}, 3, "", "Europole D14", `azimut:116.483503
altitude:18.400539
Europole District D14`},
		{"GEO-EU-D15", []string{"GEO", "EU", "D15"}, 3, "", "Europole D15", `azimut:340.027907
altitude:44.906145
Europole District D15`},
		{"GEO-EU-D16", []string{"GEO", "EU", "D16"}, 3, "", "Europole D16", `azimut:156.679333
altitude:59.225136
Europole District D16`},
		{"GEO-EU-D17", []string{"GEO", "EU", "D17"}, 3, "", "Europole D17", `azimut:98.859250
altitude:70.193535
Europole District D17`},
		{"GEO-EU-D18", []string{"GEO", "EU", "D18"}, 3, "", "Europole D18", `azimut:249.598879
altitude:35.274047
Europole District D18`},
		{"GEO-EU-D19", []string{"GEO", "EU", "D19"}, 3, "", "Europole D19", `azimut:39.792230
altitude:84.093000
Europole District D19`},
		{"GEO-EU-D20", []string{"GEO", "EU", "D20"}, 3, "", "Europole D20", `azimut:181.817280
altitude:44.512595
Europole District D20`},
		{"GEO-EU-D21", []string{"GEO", "EU", "D21"}, 3, "", "Europole D21", `azimut:150.167960
altitude:85.991215
Europole District D21`},
		{"GEO-EU-D22", []string{"GEO", "EU", "D22"}, 3, "", "Europole D22", `azimut:239.977281
altitude:75.689278
Europole District D22`},
		{"GEO-EU-D23", []string{"GEO", "EU", "D23"}, 3, "", "Europole D23", `azimut:351.246429
altitude:34.655470
Europole District D23`},
		{"GEO-EU-D24", []string{"GEO", "EU", "D24"}, 3, "", "Europole D24", `azimut:160.687062
altitude:65.748652
Europole District D24`},
		{"GEO-EU-D25", []string{"GEO", "EU", "D25"}, 3, "", "Europole D25", `azimut:199.657318
altitude:16.504889
Europole District D25`},
		{"GEO-EU-D26", []string{"GEO", "EU", "D26"}, 3, "", "Europole D26", `azimut:113.082227
altitude:70.553254
Europole District D26`},
		{"GEO-EU-D27", []string{"GEO", "EU", "D27"}, 3, "", "Europole D27", `azimut:149.928442
altitude:38.723569
Europole District D27`},
		{"GEO-EU-D28", []string{"GEO", "EU", "D28"}, 3, "", "Europole D28", `azimut:195.343852
altitude:66.718099
Europole District D28`},
		{"GEO-EU-D29", []string{"GEO", "EU", "D29"}, 3, "", "Europole D29", `azimut:68.053002
altitude:59.244626
Europole District D29`},
		{"GEO-EU-D30", []string{"GEO", "EU", "D30"}, 3, "", "Europole D30", `azimut:109.773570
altitude:75.527002
Europole District D30`},
		{"GEO-AM-D01", []string{"GEO", "AM", "D01"}, 3, "", "Amerique D01", `azimut:73.597028
altitude:49.878709
Amerique District D01`},
		{"GEO-AM-D02", []string{"GEO", "AM", "D02"}, 3, "", "Amerique D02", `azimut:123.021633
altitude:17.279446
Amerique District D02`},
		{"GEO-AM-D03", []string{"GEO", "AM", "D03"}, 3, "", "Amerique D03", `azimut:343.543004
altitude:40.210107
Amerique District D03`},
		{"GEO-AM-D04", []string{"GEO", "AM", "D04"}, 3, "", "Amerique D04", `azimut:118.704682
altitude:30.886185
Amerique District D04`},
		{"GEO-AM-D05", []string{"GEO", "AM", "D05"}, 3, "", "Amerique D05", `azimut:332.719647
altitude:66.663091
Amerique District D05`},
		{"GEO-AM-D06", []string{"GEO", "AM", "D06"}, 3, "", "Amerique D06", `azimut:170.711919
altitude:38.124117
Amerique District D06`},
		{"GEO-AM-D07", []string{"GEO", "AM", "D07"}, 3, "", "Amerique D07", `azimut:95.659856
altitude:19.927787
Amerique District D07`},
		{"GEO-AM-D08", []string{"GEO", "AM", "D08"}, 3, "", "Amerique D08", `azimut:112.172492
altitude:19.548745
Amerique District D08`},
		{"GEO-AM-D09", []string{"GEO", "AM", "D09"}, 3, "", "Amerique D09", `azimut:45.613917
altitude:51.208722
Amerique District D09`},
		{"GEO-AM-D10", []string{"GEO", "AM", "D10"}, 3, "", "Amerique D10", `azimut:330.125659
altitude:73.166312
Amerique District D10`},
		{"GEO-AM-D11", []string{"GEO", "AM", "D11"}, 3, "", "Amerique D11", `azimut:87.738024
altitude:46.632757
Amerique District D11`},
		{"GEO-AM-D12", []string{"GEO", "AM", "D12"}, 3, "", "Amerique D12", `azimut:155.679631
altitude:89.617381
Amerique District D12`},
		{"GEO-AM-D13", []string{"GEO", "AM", "D13"}, 3, "", "Amerique D13", `azimut:175.463825
altitude:13.228532
Amerique District D13`},
		{"GEO-AM-D14", []string{"GEO", "AM", "D14"}, 3, "", "Amerique D14", `azimut:182.310405
altitude:12.549442
Amerique District D14`},
		{"GEO-AM-D15", []string{"GEO", "AM", "D15"}, 3, "", "Amerique D15", `azimut:79.390452
altitude:5.071440
Amerique District D15`},
		{"GEO-AM-D16", []string{"GEO", "AM", "D16"}, 3, "", "Amerique D16", `azimut:110.318744
altitude:33.526340
Amerique District D16`},
		{"GEO-AM-D17", []string{"GEO", "AM", "D17"}, 3, "", "Amerique D17", `azimut:298.195798
altitude:84.808764
Amerique District D17`},
		{"GEO-AM-D18", []string{"GEO", "AM", "D18"}, 3, "", "Amerique D18", `azimut:2.873332
altitude:51.596388
Amerique District D18`},
		{"GEO-AM-D19", []string{"GEO", "AM", "D19"}, 3, "", "Amerique D19", `azimut:296.794890
altitude:71.677123
Amerique District D19`},
		{"GEO-AM-D20", []string{"GEO", "AM", "D20"}, 3, "", "Amerique D20", `azimut:231.170081
altitude:30.647222
Amerique District D20`},
		{"GEO-AM-D21", []string{"GEO", "AM", "D21"}, 3, "", "Amerique D21", `azimut:89.500920
altitude:30.522516
Amerique District D21`},
		{"GEO-AM-D22", []string{"GEO", "AM", "D22"}, 3, "", "Amerique D22", `azimut:323.958919
altitude:30.437744
Amerique District D22`},
		{"GEO-AM-D23", []string{"GEO", "AM", "D23"}, 3, "", "Amerique D23", `azimut:253.869255
altitude:30.920316
Amerique District D23`},
		{"GEO-AM-D24", []string{"GEO", "AM", "D24"}, 3, "", "Amerique D24", `azimut:186.124318
altitude:62.824878
Amerique District D24`},
		{"GEO-AM-D25", []string{"GEO", "AM", "D25"}, 3, "", "Amerique D25", `azimut:341.876998
altitude:86.569408
Amerique District D25`},
		{"GEO-AM-D26", []string{"GEO", "AM", "D26"}, 3, "", "Amerique D26", `azimut:345.379005
altitude:56.438396
Amerique District D26`},
		{"GEO-AM-D27", []string{"GEO", "AM", "D27"}, 3, "", "Amerique D27", `azimut:286.611533
altitude:62.474308
Amerique District D27`},
		{"GEO-AM-D28", []string{"GEO", "AM", "D28"}, 3, "", "Amerique D28", `azimut:153.424171
altitude:2.944232
Amerique District D28`},
		{"GEO-AM-D29", []string{"GEO", "AM", "D29"}, 3, "", "Amerique D29", `azimut:300.060128
altitude:70.495054
Amerique District D29`},
		{"GEO-AM-D30", []string{"GEO", "AM", "D30"}, 3, "", "Amerique D30", `azimut:213.636218
altitude:8.239724
Amerique District D30`},
		{"GEO-AS-D01", []string{"GEO", "AS", "D01"}, 3, "", "Asie D01", `azimut:342.051745
altitude:42.218323
Asie District D01`},
		{"GEO-AS-D02", []string{"GEO", "AS", "D02"}, 3, "", "Asie D02", `azimut:250.259304
altitude:71.493491
Asie District D02`},
		{"GEO-AS-D03", []string{"GEO", "AS", "D03"}, 3, "", "Asie D03", `azimut:95.251740
altitude:48.210691
Asie District D03`},
		{"GEO-AS-D04", []string{"GEO", "AS", "D04"}, 3, "", "Asie D04", `azimut:183.859707
altitude:76.592457
Asie District D04`},
		{"GEO-AS-D05", []string{"GEO", "AS", "D05"}, 3, "", "Asie D05", `azimut:333.600815
altitude:61.490801
Asie District D05`},
		{"GEO-AS-D06", []string{"GEO", "AS", "D06"}, 3, "", "Asie D06", `azimut:33.457971
altitude:64.460645
Asie District D06`},
		{"GEO-AS-D07", []string{"GEO", "AS", "D07"}, 3, "", "Asie D07", `azimut:205.287459
altitude:21.342705
Asie District D07`},
		{"GEO-AS-D08", []string{"GEO", "AS", "D08"}, 3, "", "Asie D08", `azimut:290.712095
altitude:83.393417
Asie District D08`},
		{"GEO-AS-D09", []string{"GEO", "AS", "D09"}, 3, "", "Asie D09", `azimut:129.804248
altitude:56.810036
Asie District D09`},
		{"GEO-AS-D10", []string{"GEO", "AS", "D10"}, 3, "", "Asie D10", `azimut:181.933236
altitude:82.832447
Asie District D10`},
		{"GEO-AS-D11", []string{"GEO", "AS", "D11"}, 3, "", "Asie D11", `azimut:1.056958
altitude:46.733304
Asie District D11`},
		{"GEO-AS-D12", []string{"GEO", "AS", "D12"}, 3, "", "Asie D12", `azimut:130.245444
altitude:38.701721
Asie District D12`},
		{"GEO-AS-D13", []string{"GEO", "AS", "D13"}, 3, "", "Asie D13", `azimut:11.754200
altitude:29.779863
Asie District D13`},
		{"GEO-AS-D14", []string{"GEO", "AS", "D14"}, 3, "", "Asie D14", `azimut:63.656253
altitude:50.797300
Asie District D14`},
		{"GEO-AS-D15", []string{"GEO", "AS", "D15"}, 3, "", "Asie D15", `azimut:240.014087
altitude:64.780596
Asie District D15`},
		{"GEO-AS-D16", []string{"GEO", "AS", "D16"}, 3, "", "Asie D16", `azimut:185.696595
altitude:44.346859
Asie District D16`},
		{"GEO-AS-D17", []string{"GEO", "AS", "D17"}, 3, "", "Asie D17", `azimut:204.779937
altitude:13.827560
Asie District D17`},
		{"GEO-AS-D18", []string{"GEO", "AS", "D18"}, 3, "", "Asie D18", `azimut:92.243999
altitude:9.715407
Asie District D18`},
		{"GEO-AS-D19", []string{"GEO", "AS", "D19"}, 3, "", "Asie D19", `azimut:199.837823
altitude:58.591530
Asie District D19`},
		{"GEO-AS-D20", []string{"GEO", "AS", "D20"}, 3, "", "Asie D20", `azimut:53.849202
altitude:36.692392
Asie District D20`},
		{"GEO-AS-D21", []string{"GEO", "AS", "D21"}, 3, "", "Asie D21", `azimut:142.217905
altitude:76.522163
Asie District D21`},
		{"GEO-AS-D22", []string{"GEO", "AS", "D22"}, 3, "", "Asie D22", `azimut:75.695987
altitude:62.240956
Asie District D22`},
		{"GEO-AS-D23", []string{"GEO", "AS", "D23"}, 3, "", "Asie D23", `azimut:285.624536
altitude:57.488524
Asie District D23`},
		{"GEO-AS-D24", []string{"GEO", "AS", "D24"}, 3, "", "Asie D24", `azimut:274.254944
altitude:45.720608
Asie District D24`},
		{"GEO-AS-D25", []string{"GEO", "AS", "D25"}, 3, "", "Asie D25", `azimut:357.594847
altitude:36.241697
Asie District D25`},
		{"GEO-AS-D26", []string{"GEO", "AS", "D26"}, 3, "", "Asie D26", `azimut:87.273324
altitude:84.761897
Asie District D26`},
		{"GEO-AS-D27", []string{"GEO", "AS", "D27"}, 3, "", "Asie D27", `azimut:192.987743
altitude:39.505095
Asie District D27`},
		{"GEO-AS-D28", []string{"GEO", "AS", "D28"}, 3, "", "Asie D28", `azimut:330.694254
altitude:67.430962
Asie District D28`},
		{"GEO-AS-D29", []string{"GEO", "AS", "D29"}, 3, "", "Asie D29", `azimut:127.146828
altitude:19.723466
Asie District D29`},
		{"GEO-AS-D30", []string{"GEO", "AS", "D30"}, 3, "", "Asie D30", `azimut:74.864500
altitude:75.241127
Asie District D30`},
		{"GEO-AU-D01", []string{"GEO", "AU", "D01"}, 3, "", "Australie D01", `azimut:292.083689
altitude:88.194980
Australie District D01`},
		{"GEO-AU-D02", []string{"GEO", "AU", "D02"}, 3, "", "Australie D02", `azimut:82.831966
altitude:21.731413
Australie District D02`},
		{"GEO-AU-D03", []string{"GEO", "AU", "D03"}, 3, "", "Australie D03", `azimut:154.231428
altitude:1.419308
Australie District D03`},
		{"GEO-AU-D04", []string{"GEO", "AU", "D04"}, 3, "", "Australie D04", `azimut:24.530170
altitude:89.837985
Australie District D04`},
		{"GEO-AU-D05", []string{"GEO", "AU", "D05"}, 3, "", "Australie D05", `azimut:57.894963
altitude:45.076034
Australie District D05`},
		{"GEO-AU-D06", []string{"GEO", "AU", "D06"}, 3, "", "Australie D06", `azimut:342.363100
altitude:21.464879
Australie District D06`},
		{"GEO-AU-D07", []string{"GEO", "AU", "D07"}, 3, "", "Australie D07", `azimut:78.005193
altitude:5.087446
Australie District D07`},
		{"GEO-AU-D08", []string{"GEO", "AU", "D08"}, 3, "", "Australie D08", `azimut:225.866362
altitude:67.567458
Australie District D08`},
		{"GEO-AU-D09", []string{"GEO", "AU", "D09"}, 3, "", "Australie D09", `azimut:167.444574
altitude:61.857512
Australie District D09`},
		{"GEO-AU-D10", []string{"GEO", "AU", "D10"}, 3, "", "Australie D10", `azimut:255.402402
altitude:36.556867
Australie District D10`},
		{"GEO-AU-D11", []string{"GEO", "AU", "D11"}, 3, "", "Australie D11", `azimut:140.726251
altitude:11.197472
Australie District D11`},
		{"GEO-AU-D12", []string{"GEO", "AU", "D12"}, 3, "", "Australie D12", `azimut:161.236059
altitude:36.436500
Australie District D12`},
		{"GEO-AU-D13", []string{"GEO", "AU", "D13"}, 3, "", "Australie D13", `azimut:17.379168
altitude:33.483234
Australie District D13`},
		{"GEO-AU-D14", []string{"GEO", "AU", "D14"}, 3, "", "Australie D14", `azimut:51.559107
altitude:18.235380
Australie District D14`},
		{"GEO-AU-D15", []string{"GEO", "AU", "D15"}, 3, "", "Australie D15", `azimut:108.356927
altitude:73.259122
Australie District D15`},
		{"GEO-AU-D16", []string{"GEO", "AU", "D16"}, 3, "", "Australie D16", `azimut:165.263461
altitude:1.147638
Australie District D16`},
		{"GEO-AU-D17", []string{"GEO", "AU", "D17"}, 3, "", "Australie D17", `azimut:316.051053
altitude:70.737982
Australie District D17`},
		{"GEO-AU-D18", []string{"GEO", "AU", "D18"}, 3, "", "Australie D18", `azimut:76.158201
altitude:64.184684
Australie District D18`},
		{"GEO-AU-D19", []string{"GEO", "AU", "D19"}, 3, "", "Australie D19", `azimut:110.508745
altitude:49.884585
Australie District D19`},
		{"GEO-AU-D20", []string{"GEO", "AU", "D20"}, 3, "", "Australie D20", `azimut:255.635603
altitude:77.332306
Australie District D20`},
		{"GEO-AU-D21", []string{"GEO", "AU", "D21"}, 3, "", "Australie D21", `azimut:277.905640
altitude:72.924568
Australie District D21`},
		{"GEO-AU-D22", []string{"GEO", "AU", "D22"}, 3, "", "Australie D22", `azimut:239.945763
altitude:45.988657
Australie District D22`},
		{"GEO-AU-D23", []string{"GEO", "AU", "D23"}, 3, "", "Australie D23", `azimut:44.763104
altitude:87.511096
Australie District D23`},
		{"GEO-AU-D24", []string{"GEO", "AU", "D24"}, 3, "", "Australie D24", `azimut:288.379136
altitude:78.406113
Australie District D24`},
		{"GEO-AU-D25", []string{"GEO", "AU", "D25"}, 3, "", "Australie D25", `azimut:130.566795
altitude:26.739230
Australie District D25`},
		{"GEO-AU-D26", []string{"GEO", "AU", "D26"}, 3, "", "Australie D26", `azimut:122.319196
altitude:86.855428
Australie District D26`},
		{"GEO-AU-D27", []string{"GEO", "AU", "D27"}, 3, "", "Australie D27", `azimut:259.975290
altitude:21.280427
Australie District D27`},
		{"GEO-AU-D28", []string{"GEO", "AU", "D28"}, 3, "", "Australie D28", `azimut:96.620961
altitude:41.302436
Australie District D28`},
		{"GEO-AU-D29", []string{"GEO", "AU", "D29"}, 3, "", "Australie D29", `azimut:146.349130
altitude:81.275968
Australie District D29`},
		{"GEO-AU-D30", []string{"GEO", "AU", "D30"}, 3, "", "Australie D30", `azimut:75.093528
altitude:80.022612
Australie District D30`},
		{"LEO-SATCOM", []string{"LEO", "SATCOM"}, 3, "", "Constellation SATCOM orbite LEO", "6a8e2a76-b0b7-42e4-aec4-9af7d0b1339e"},
		{"LEO-STARLINK", []string{"LEO", "STARLINK"}, 3, "", "Constellation STARLINK orbite LEO", "39a4c6ac-3710-4aca-b802-8ab7bce8b6fa"},
		{"LEO-VIASAT", []string{"LEO", "VIASAT"}, 3, "", "Constellation VIASAT orbite LEO", "843521bc-30f7-4d73-b68c-a1ea707da880"},
		{"LEO-IRIDIUM", []string{"LEO", "IRIDIUM"}, 3, "", "Constellation IRIDIUM orbite LEO", "10e77c5c-ee5e-4324-9547-f2856ea3e3ac"},
		{"MEO-SATCOM", []string{"MEO", "SATCOM"}, 3, "", "Constellation SATCOM orbite MEO", "fbc01335-8224-43e0-a175-298e43832f96"},
		{"MEO-STARLINK", []string{"MEO", "STARLINK"}, 3, "", "Constellation STARLINK orbite MEO", "704cf5b2-1ccd-42ae-a93a-054fa65f7950"},
		{"MEO-VIASAT", []string{"MEO", "VIASAT"}, 3, "", "Constellation VIASAT orbite MEO", "7ba9b084-442f-445e-97ce-b3936c368079"},
		{"MEO-IRIDIUM", []string{"MEO", "IRIDIUM"}, 3, "", "Constellation IRIDIUM orbite MEO", "c65c0fe2-0ada-4555-b641-439239426488"},
	},
	Registers: []Register{
		{"GEO-EU-D01-600MHZ", false, "", 5},
		{"GEO-EU-D01-800MHZ", true, "", 5},
		{"GEO-EU-D01-1200MHZ", true, "", 5},
		{"GEO-EU-D02-600MHZ", false, "", 5},
		{"GEO-EU-D02-800MHZ", false, "", 5},
		{"GEO-EU-D02-1200MHZ", true, "", 5},
		{"GEO-EU-D03-600MHZ", false, "", 5},
		{"GEO-EU-D03-800MHZ", false, "", 5},
		{"GEO-EU-D03-1200MHZ", false, "", 5},
		{"GEO-EU-D04-600MHZ", false, "", 5},
		{"GEO-EU-D04-800MHZ", false, "", 5},
		{"GEO-EU-D04-1200MHZ", false, "", 5},
		{"GEO-EU-D05-600MHZ", false, "", 5},
		{"GEO-EU-D05-800MHZ", false, "", 5},
		{"GEO-EU-D05-1200MHZ", false, "", 5},
		{"GEO-EU-D06-600MHZ", false, "", 5},
		{"GEO-EU-D06-800MHZ", true, "", 5},
		{"GEO-EU-D06-1200MHZ", false, "", 5},
		{"GEO-EU-D07-600MHZ", false, "", 5},
		{"GEO-EU-D07-800MHZ", true, "", 5},
		{"GEO-EU-D07-1200MHZ", false, "", 5},
		{"GEO-EU-D08-600MHZ", false, "", 5},
		{"GEO-EU-D08-800MHZ", true, "", 5},
		{"GEO-EU-D08-1200MHZ", false, "", 5},
		{"GEO-EU-D09-600MHZ", false, "", 5},
		{"GEO-EU-D09-800MHZ", true, "", 5},
		{"GEO-EU-D09-1200MHZ", false, "", 5},
		{"GEO-EU-D10-600MHZ", false, "", 5},
		{"GEO-EU-D10-800MHZ", true, "", 5},
		{"GEO-EU-D10-1200MHZ", true, "", 5},
		{"GEO-EU-D11-600MHZ", false, "", 5},
		{"GEO-EU-D11-800MHZ", true, "", 5},
		{"GEO-EU-D11-1200MHZ", false, "", 5},
		{"GEO-EU-D12-600MHZ", false, "", 5},
		{"GEO-EU-D12-800MHZ", true, "", 5},
		{"GEO-EU-D12-1200MHZ", false, "", 5},
		{"GEO-EU-D13-600MHZ", false, "", 5},
		{"GEO-EU-D13-800MHZ", true, "", 5},
		{"GEO-EU-D13-1200MHZ", true, "", 5},
		{"GEO-EU-D14-600MHZ", false, "", 5},
		{"GEO-EU-D14-800MHZ", false, "", 5},
		{"GEO-EU-D14-1200MHZ", true, "", 5},
		{"GEO-EU-D15-600MHZ", false, "", 5},
		{"GEO-EU-D15-800MHZ", true, "", 5},
		{"GEO-EU-D15-1200MHZ", false, "", 5},
		{"GEO-EU-D16-600MHZ", false, "", 5},
		{"GEO-EU-D16-800MHZ", true, "", 5},
		{"GEO-EU-D16-1200MHZ", true, "", 5},
		{"GEO-EU-D17-600MHZ", false, "", 5},
		{"GEO-EU-D17-800MHZ", false, "", 5},
		{"GEO-EU-D17-1200MHZ", true, "", 5},
		{"GEO-EU-D18-600MHZ", false, "", 5},
		{"GEO-EU-D18-800MHZ", false, "", 5},
		{"GEO-EU-D18-1200MHZ", true, "", 5},
		{"GEO-EU-D19-600MHZ", false, "", 5},
		{"GEO-EU-D19-800MHZ", true, "", 5},
		{"GEO-EU-D19-1200MHZ", false, "", 5},
		{"GEO-EU-D20-600MHZ", false, "", 5},
		{"GEO-EU-D20-800MHZ", false, "", 5},
		{"GEO-EU-D20-1200MHZ", false, "", 5},
		{"GEO-EU-D21-600MHZ", false, "", 5},
		{"GEO-EU-D21-800MHZ", false, "", 5},
		{"GEO-EU-D21-1200MHZ", true, "", 5},
		{"GEO-EU-D22-600MHZ", false, "", 5},
		{"GEO-EU-D22-800MHZ", false, "", 5},
		{"GEO-EU-D22-1200MHZ", false, "", 5},
		{"GEO-EU-D23-600MHZ", false, "", 5},
		{"GEO-EU-D23-800MHZ", true, "", 5},
		{"GEO-EU-D23-1200MHZ", false, "", 5},
		{"GEO-EU-D24-600MHZ", false, "", 5},
		{"GEO-EU-D24-800MHZ", true, "", 5},
		{"GEO-EU-D24-1200MHZ", true, "", 5},
		{"GEO-EU-D25-600MHZ", false, "", 5},
		{"GEO-EU-D25-800MHZ", false, "", 5},
		{"GEO-EU-D25-1200MHZ", true, "", 5},
		{"GEO-EU-D26-600MHZ", false, "", 5},
		{"GEO-EU-D26-800MHZ", false, "", 5},
		{"GEO-EU-D26-1200MHZ", true, "", 5},
		{"GEO-EU-D27-600MHZ", false, "", 5},
		{"GEO-EU-D27-800MHZ", true, "", 5},
		{"GEO-EU-D27-1200MHZ", false, "", 5},
		{"GEO-EU-D28-600MHZ", false, "", 5},
		{"GEO-EU-D28-800MHZ", true, "", 5},
		{"GEO-EU-D28-1200MHZ", false, "", 5},
		{"GEO-EU-D29-600MHZ", false, "", 5},
		{"GEO-EU-D29-800MHZ", true, "", 5},
		{"GEO-EU-D29-1200MHZ", true, "", 5},
		{"GEO-EU-D30-600MHZ", false, "", 5},
		{"GEO-EU-D30-800MHZ", false, "", 5},
		{"GEO-EU-D30-1200MHZ", true, "", 5},
		{"GEO-AM-D01-600MHZ", false, "", 5},
		{"GEO-AM-D01-800MHZ", true, "", 5},
		{"GEO-AM-D01-1200MHZ", true, "", 5},
		{"GEO-AM-D02-600MHZ", false, "", 5},
		{"GEO-AM-D02-800MHZ", true, "", 5},
		{"GEO-AM-D02-1200MHZ", false, "", 5},
		{"GEO-AM-D03-600MHZ", false, "", 5},
		{"GEO-AM-D03-800MHZ", true, "", 5},
		{"GEO-AM-D03-1200MHZ", false, "", 5},
		{"GEO-AM-D04-600MHZ", false, "", 5},
		{"GEO-AM-D04-800MHZ", true, "", 5},
		{"GEO-AM-D04-1200MHZ", true, "", 5},
		{"GEO-AM-D05-600MHZ", false, "", 5},
		{"GEO-AM-D05-800MHZ", true, "", 5},
		{"GEO-AM-D05-1200MHZ", false, "", 5},
		{"GEO-AM-D06-600MHZ", false, "", 5},
		{"GEO-AM-D06-800MHZ", false, "", 5},
		{"GEO-AM-D06-1200MHZ", false, "", 5},
		{"GEO-AM-D07-600MHZ", false, "", 5},
		{"GEO-AM-D07-800MHZ", false, "", 5},
		{"GEO-AM-D07-1200MHZ", true, "", 5},
		{"GEO-AM-D08-600MHZ", false, "", 5},
		{"GEO-AM-D08-800MHZ", true, "", 5},
		{"GEO-AM-D08-1200MHZ", false, "", 5},
		{"GEO-AM-D09-600MHZ", false, "", 5},
		{"GEO-AM-D09-800MHZ", false, "", 5},
		{"GEO-AM-D09-1200MHZ", false, "", 5},
		{"GEO-AM-D10-600MHZ", false, "", 5},
		{"GEO-AM-D10-800MHZ", false, "", 5},
		{"GEO-AM-D10-1200MHZ", false, "", 5},
		{"GEO-AM-D11-600MHZ", false, "", 5},
		{"GEO-AM-D11-800MHZ", false, "", 5},
		{"GEO-AM-D11-1200MHZ", false, "", 5},
		{"GEO-AM-D12-600MHZ", false, "", 5},
		{"GEO-AM-D12-800MHZ", false, "", 5},
		{"GEO-AM-D12-1200MHZ", false, "", 5},
		{"GEO-AM-D13-600MHZ", false, "", 5},
		{"GEO-AM-D13-800MHZ", true, "", 5},
		{"GEO-AM-D13-1200MHZ", true, "", 5},
		{"GEO-AM-D14-600MHZ", false, "", 5},
		{"GEO-AM-D14-800MHZ", false, "", 5},
		{"GEO-AM-D14-1200MHZ", false, "", 5},
		{"GEO-AM-D15-600MHZ", false, "", 5},
		{"GEO-AM-D15-800MHZ", true, "", 5},
		{"GEO-AM-D15-1200MHZ", false, "", 5},
		{"GEO-AM-D16-600MHZ", false, "", 5},
		{"GEO-AM-D16-800MHZ", false, "", 5},
		{"GEO-AM-D16-1200MHZ", false, "", 5},
		{"GEO-AM-D17-600MHZ", false, "", 5},
		{"GEO-AM-D17-800MHZ", true, "", 5},
		{"GEO-AM-D17-1200MHZ", true, "", 5},
		{"GEO-AM-D18-600MHZ", false, "", 5},
		{"GEO-AM-D18-800MHZ", false, "", 5},
		{"GEO-AM-D18-1200MHZ", false, "", 5},
		{"GEO-AM-D19-600MHZ", false, "", 5},
		{"GEO-AM-D19-800MHZ", true, "", 5},
		{"GEO-AM-D19-1200MHZ", true, "", 5},
		{"GEO-AM-D20-600MHZ", false, "", 5},
		{"GEO-AM-D20-800MHZ", true, "", 5},
		{"GEO-AM-D20-1200MHZ", true, "", 5},
		{"GEO-AM-D21-600MHZ", false, "", 5},
		{"GEO-AM-D21-800MHZ", true, "", 5},
		{"GEO-AM-D21-1200MHZ", false, "", 5},
		{"GEO-AM-D22-600MHZ", false, "", 5},
		{"GEO-AM-D22-800MHZ", true, "", 5},
		{"GEO-AM-D22-1200MHZ", true, "", 5},
		{"GEO-AM-D23-600MHZ", false, "", 5},
		{"GEO-AM-D23-800MHZ", false, "", 5},
		{"GEO-AM-D23-1200MHZ", false, "", 5},
		{"GEO-AM-D24-600MHZ", false, "", 5},
		{"GEO-AM-D24-800MHZ", true, "", 5},
		{"GEO-AM-D24-1200MHZ", true, "", 5},
		{"GEO-AM-D25-600MHZ", false, "", 5},
		{"GEO-AM-D25-800MHZ", false, "", 5},
		{"GEO-AM-D25-1200MHZ", true, "", 5},
		{"GEO-AM-D26-600MHZ", false, "", 5},
		{"GEO-AM-D26-800MHZ", true, "", 5},
		{"GEO-AM-D26-1200MHZ", false, "", 5},
		{"GEO-AM-D27-600MHZ", false, "", 5},
		{"GEO-AM-D27-800MHZ", true, "", 5},
		{"GEO-AM-D27-1200MHZ", false, "", 5},
		{"GEO-AM-D28-600MHZ", false, "", 5},
		{"GEO-AM-D28-800MHZ", false, "", 5},
		{"GEO-AM-D28-1200MHZ", false, "", 5},
		{"GEO-AM-D29-600MHZ", false, "", 5},
		{"GEO-AM-D29-800MHZ", false, "", 5},
		{"GEO-AM-D29-1200MHZ", true, "", 5},
		{"GEO-AM-D30-600MHZ", false, "", 5},
		{"GEO-AM-D30-800MHZ", true, "", 5},
		{"GEO-AM-D30-1200MHZ", false, "", 5},
		{"GEO-AS-D01-600MHZ", false, "", 5},
		{"GEO-AS-D01-800MHZ", false, "", 5},
		{"GEO-AS-D01-1200MHZ", true, "", 5},
		{"GEO-AS-D02-600MHZ", false, "", 5},
		{"GEO-AS-D02-800MHZ", true, "", 5},
		{"GEO-AS-D02-1200MHZ", true, "", 5},
		{"GEO-AS-D03-600MHZ", false, "", 5},
		{"GEO-AS-D03-800MHZ", false, "", 5},
		{"GEO-AS-D03-1200MHZ", false, "", 5},
		{"GEO-AS-D04-600MHZ", false, "", 5},
		{"GEO-AS-D04-800MHZ", true, "", 5},
		{"GEO-AS-D04-1200MHZ", false, "", 5},
		{"GEO-AS-D05-600MHZ", false, "", 5},
		{"GEO-AS-D05-800MHZ", false, "", 5},
		{"GEO-AS-D05-1200MHZ", true, "", 5},
		{"GEO-AS-D06-600MHZ", false, "", 5},
		{"GEO-AS-D06-800MHZ", true, "", 5},
		{"GEO-AS-D06-1200MHZ", false, "", 5},
		{"GEO-AS-D07-600MHZ", false, "", 5},
		{"GEO-AS-D07-800MHZ", true, "", 5},
		{"GEO-AS-D07-1200MHZ", true, "", 5},
		{"GEO-AS-D08-600MHZ", false, "", 5},
		{"GEO-AS-D08-800MHZ", true, "", 5},
		{"GEO-AS-D08-1200MHZ", false, "", 5},
		{"GEO-AS-D09-600MHZ", false, "", 5},
		{"GEO-AS-D09-800MHZ", false, "", 5},
		{"GEO-AS-D09-1200MHZ", false, "", 5},
		{"GEO-AS-D10-600MHZ", false, "", 5},
		{"GEO-AS-D10-800MHZ", false, "", 5},
		{"GEO-AS-D10-1200MHZ", true, "", 5},
		{"GEO-AS-D11-600MHZ", false, "", 5},
		{"GEO-AS-D11-800MHZ", true, "", 5},
		{"GEO-AS-D11-1200MHZ", true, "", 5},
		{"GEO-AS-D12-600MHZ", false, "", 5},
		{"GEO-AS-D12-800MHZ", false, "", 5},
		{"GEO-AS-D12-1200MHZ", true, "", 5},
		{"GEO-AS-D13-600MHZ", false, "", 5},
		{"GEO-AS-D13-800MHZ", true, "", 5},
		{"GEO-AS-D13-1200MHZ", true, "", 5},
		{"GEO-AS-D14-600MHZ", false, "", 5},
		{"GEO-AS-D14-800MHZ", true, "", 5},
		{"GEO-AS-D14-1200MHZ", false, "", 5},
		{"GEO-AS-D15-600MHZ", false, "", 5},
		{"GEO-AS-D15-800MHZ", false, "", 5},
		{"GEO-AS-D15-1200MHZ", false, "", 5},
		{"GEO-AS-D16-600MHZ", false, "", 5},
		{"GEO-AS-D16-800MHZ", false, "", 5},
		{"GEO-AS-D16-1200MHZ", false, "", 5},
		{"GEO-AS-D17-600MHZ", false, "", 5},
		{"GEO-AS-D17-800MHZ", false, "", 5},
		{"GEO-AS-D17-1200MHZ", true, "", 5},
		{"GEO-AS-D18-600MHZ", false, "", 5},
		{"GEO-AS-D18-800MHZ", false, "", 5},
		{"GEO-AS-D18-1200MHZ", true, "", 5},
		{"GEO-AS-D19-600MHZ", false, "", 5},
		{"GEO-AS-D19-800MHZ", true, "", 5},
		{"GEO-AS-D19-1200MHZ", true, "", 5},
		{"GEO-AS-D20-600MHZ", false, "", 5},
		{"GEO-AS-D20-800MHZ", true, "", 5},
		{"GEO-AS-D20-1200MHZ", true, "", 5},
		{"GEO-AS-D21-600MHZ", false, "", 5},
		{"GEO-AS-D21-800MHZ", true, "", 5},
		{"GEO-AS-D21-1200MHZ", false, "", 5},
		{"GEO-AS-D22-600MHZ", false, "", 5},
		{"GEO-AS-D22-800MHZ", false, "", 5},
		{"GEO-AS-D22-1200MHZ", false, "", 5},
		{"GEO-AS-D23-600MHZ", false, "", 5},
		{"GEO-AS-D23-800MHZ", true, "", 5},
		{"GEO-AS-D23-1200MHZ", true, "", 5},
		{"GEO-AS-D24-600MHZ", false, "", 5},
		{"GEO-AS-D24-800MHZ", false, "", 5},
		{"GEO-AS-D24-1200MHZ", false, "", 5},
		{"GEO-AS-D25-600MHZ", false, "", 5},
		{"GEO-AS-D25-800MHZ", false, "", 5},
		{"GEO-AS-D25-1200MHZ", true, "", 5},
		{"GEO-AS-D26-600MHZ", false, "", 5},
		{"GEO-AS-D26-800MHZ", true, "", 5},
		{"GEO-AS-D26-1200MHZ", false, "", 5},
		{"GEO-AS-D27-600MHZ", false, "", 5},
		{"GEO-AS-D27-800MHZ", false, "", 5},
		{"GEO-AS-D27-1200MHZ", false, "", 5},
		{"GEO-AS-D28-600MHZ", false, "", 5},
		{"GEO-AS-D28-800MHZ", true, "", 5},
		{"GEO-AS-D28-1200MHZ", true, "", 5},
		{"GEO-AS-D29-600MHZ", false, "", 5},
		{"GEO-AS-D29-800MHZ", true, "", 5},
		{"GEO-AS-D29-1200MHZ", true, "", 5},
		{"GEO-AS-D30-600MHZ", false, "", 5},
		{"GEO-AS-D30-800MHZ", false, "", 5},
		{"GEO-AS-D30-1200MHZ", false, "", 5},
		{"GEO-AU-D01-600MHZ", false, "", 5},
		{"GEO-AU-D01-800MHZ", true, "", 5},
		{"GEO-AU-D01-1200MHZ", false, "", 5},
		{"GEO-AU-D02-600MHZ", false, "", 5},
		{"GEO-AU-D02-800MHZ", false, "", 5},
		{"GEO-AU-D02-1200MHZ", false, "", 5},
		{"GEO-AU-D03-600MHZ", false, "", 5},
		{"GEO-AU-D03-800MHZ", true, "", 5},
		{"GEO-AU-D03-1200MHZ", false, "", 5},
		{"GEO-AU-D04-600MHZ", false, "", 5},
		{"GEO-AU-D04-800MHZ", true, "", 5},
		{"GEO-AU-D04-1200MHZ", true, "", 5},
		{"GEO-AU-D05-600MHZ", false, "", 5},
		{"GEO-AU-D05-800MHZ", false, "", 5},
		{"GEO-AU-D05-1200MHZ", true, "", 5},
		{"GEO-AU-D06-600MHZ", false, "", 5},
		{"GEO-AU-D06-800MHZ", true, "", 5},
		{"GEO-AU-D06-1200MHZ", false, "", 5},
		{"GEO-AU-D07-600MHZ", false, "", 5},
		{"GEO-AU-D07-800MHZ", true, "", 5},
		{"GEO-AU-D07-1200MHZ", false, "", 5},
		{"GEO-AU-D08-600MHZ", false, "", 5},
		{"GEO-AU-D08-800MHZ", true, "", 5},
		{"GEO-AU-D08-1200MHZ", false, "", 5},
		{"GEO-AU-D09-600MHZ", false, "", 5},
		{"GEO-AU-D09-800MHZ", false, "", 5},
		{"GEO-AU-D09-1200MHZ", true, "", 5},
		{"GEO-AU-D10-600MHZ", false, "", 5},
		{"GEO-AU-D10-800MHZ", false, "", 5},
		{"GEO-AU-D10-1200MHZ", false, "", 5},
		{"GEO-AU-D11-600MHZ", false, "", 5},
		{"GEO-AU-D11-800MHZ", false, "", 5},
		{"GEO-AU-D11-1200MHZ", false, "", 5},
		{"GEO-AU-D12-600MHZ", false, "", 5},
		{"GEO-AU-D12-800MHZ", true, "", 5},
		{"GEO-AU-D12-1200MHZ", false, "", 5},
		{"GEO-AU-D13-600MHZ", false, "", 5},
		{"GEO-AU-D13-800MHZ", false, "", 5},
		{"GEO-AU-D13-1200MHZ", true, "", 5},
		{"GEO-AU-D14-600MHZ", false, "", 5},
		{"GEO-AU-D14-800MHZ", true, "", 5},
		{"GEO-AU-D14-1200MHZ", false, "", 5},
		{"GEO-AU-D15-600MHZ", false, "", 5},
		{"GEO-AU-D15-800MHZ", false, "", 5},
		{"GEO-AU-D15-1200MHZ", false, "", 5},
		{"GEO-AU-D16-600MHZ", false, "", 5},
		{"GEO-AU-D16-800MHZ", false, "", 5},
		{"GEO-AU-D16-1200MHZ", true, "", 5},
		{"GEO-AU-D17-600MHZ", false, "", 5},
		{"GEO-AU-D17-800MHZ", false, "", 5},
		{"GEO-AU-D17-1200MHZ", true, "", 5},
		{"GEO-AU-D18-600MHZ", false, "", 5},
		{"GEO-AU-D18-800MHZ", false, "", 5},
		{"GEO-AU-D18-1200MHZ", false, "", 5},
		{"GEO-AU-D19-600MHZ", false, "", 5},
		{"GEO-AU-D19-800MHZ", true, "", 5},
		{"GEO-AU-D19-1200MHZ", false, "", 5},
		{"GEO-AU-D20-600MHZ", false, "", 5},
		{"GEO-AU-D20-800MHZ", true, "", 5},
		{"GEO-AU-D20-1200MHZ", true, "", 5},
		{"GEO-AU-D21-600MHZ", false, "", 5},
		{"GEO-AU-D21-800MHZ", true, "", 5},
		{"GEO-AU-D21-1200MHZ", false, "", 5},
		{"GEO-AU-D22-600MHZ", false, "", 5},
		{"GEO-AU-D22-800MHZ", false, "", 5},
		{"GEO-AU-D22-1200MHZ", false, "", 5},
		{"GEO-AU-D23-600MHZ", false, "", 5},
		{"GEO-AU-D23-800MHZ", false, "", 5},
		{"GEO-AU-D23-1200MHZ", true, "", 5},
		{"GEO-AU-D24-600MHZ", false, "", 5},
		{"GEO-AU-D24-800MHZ", false, "", 5},
		{"GEO-AU-D24-1200MHZ", true, "", 5},
		{"GEO-AU-D25-600MHZ", false, "", 5},
		{"GEO-AU-D25-800MHZ", true, "", 5},
		{"GEO-AU-D25-1200MHZ", false, "", 5},
		{"GEO-AU-D26-600MHZ", false, "", 5},
		{"GEO-AU-D26-800MHZ", true, "", 5},
		{"GEO-AU-D26-1200MHZ", true, "", 5},
		{"GEO-AU-D27-600MHZ", false, "", 5},
		{"GEO-AU-D27-800MHZ", false, "", 5},
		{"GEO-AU-D27-1200MHZ", true, "", 5},
		{"GEO-AU-D28-600MHZ", false, "", 5},
		{"GEO-AU-D28-800MHZ", false, "", 5},
		{"GEO-AU-D28-1200MHZ", true, "", 5},
		{"GEO-AU-D29-600MHZ", false, "", 5},
		{"GEO-AU-D29-800MHZ", true, "", 5},
		{"GEO-AU-D29-1200MHZ", true, "", 5},
		{"GEO-AU-D30-600MHZ", false, "", 5},
		{"GEO-AU-D30-800MHZ", true, "", 5},
		{"GEO-AU-D30-1200MHZ", false, "", 5},
		{"LEO-SATCOM-600MHZ", false, "", 5},
		{"LEO-SATCOM-800MHZ", false, "", 5},
		{"LEO-SATCOM-1200MHZ", true, "", 5},
		{"LEO-STARLINK-600MHZ", false, "", 5},
		{"LEO-STARLINK-800MHZ", true, "", 5},
		{"LEO-STARLINK-1200MHZ", false, "", 5},
		{"LEO-VIASAT-600MHZ", false, "", 5},
		{"LEO-VIASAT-800MHZ", true, "", 5},
		{"LEO-VIASAT-1200MHZ", false, "", 5},
		{"LEO-IRIDIUM-600MHZ", false, "", 5},
		{"LEO-IRIDIUM-800MHZ", true, "", 5},
		{"LEO-IRIDIUM-1200MHZ", true, "", 5},
		{"MEO-SATCOM-600MHZ", false, "", 5},
		{"MEO-SATCOM-800MHZ", true, "", 5},
		{"MEO-SATCOM-1200MHZ", false, "", 5},
		{"MEO-STARLINK-600MHZ", false, "", 5},
		{"MEO-STARLINK-800MHZ", true, "", 5},
		{"MEO-STARLINK-1200MHZ", true, "", 5},
		{"MEO-VIASAT-600MHZ", false, "", 5},
		{"MEO-VIASAT-800MHZ", false, "", 5},
		{"MEO-VIASAT-1200MHZ", true, "", 5},
		{"MEO-IRIDIUM-600MHZ", false, "", 5},
		{"MEO-IRIDIUM-800MHZ", false, "", 5},
		{"MEO-IRIDIUM-1200MHZ", true, "", 5},
	},
}

// serveur archive de Silicon Spirit
var legba_archive = Server{
	Address: "archive.legba.d22.eu",
	Credentials: []Cred{
		legbaPersonnel,
		legbaAdmin,
	},
	Description: arcDesc,
	Entries: []Entry{
		{"art27", []string{"Mandrake", "Yuong", "Herswing", "Mathison"}, 1, "", "Transcendance sous co-routines", `Titre: Transcendance sous co-routines.

Auteurs: D. Yuong, E. Herswing et A. Mathison

Résumé: Des expériences préliminaires sur la Transcendance Numérique (TN)
permettent la formulation d'une nouvelle conjoncture. Lors des phases terminales
d'auto-corections des caractéristiques de personnalité, il paraît primordial de
passer par des co-routines évolutionnaires de Spinksy-Yuong avec des timeout
ré-évalués par adaptation prodonde. Nos premiers résultats montrent en effet
que sans l'interlaçage de ces co-routines, un effondrement neuro-dépressif
acquiert une vraissemblance voisinant les 92.34 pourcents. Notrons que nos
expériences ont été réalisées in vitro, sur le cluster de calcul SaberSapience
v 3.0.23.`},
		{"art95", []string{"Mandrake", "Mathison"}, 1, "", "Transcendance et effondrement neuro-psychologique", `Titre: Transcendance et effondrement neuro-psychologique.

Auteur: A. Mathison

Résumé: Nous présentons une review de la littérature récentes sur la 
«dégénérescence accélérée» issues de la théorie formulée par le laboratoire
@HigherMind(sous mécénat de Legba Voodoocom). Ces travaux, concommitants à
ceux réalisées par notre propre équipe de SiliconSpirit, ont l'avantage
d'avoir défrichés des pistes non-viables. Il ressort en effet de notre
analyse que la complexité des algorithmes de compressions des aires
thalamiques n'est pas un élément essentiel sur la voie de l'élévation de la
conscience numérique. Cela renforce notre hypothèse selon laquelle l'alignement
mémoriel de classe IV est une contrainte incontournable malgré le surcoût
computationnel certain.`},

		{"art164", []string{"Mandrake", "Levain", "Revertin", "Galakievicz"}, 1, "", "Deux encodages mémoriels pour les aires thalamiques intérmédiaires", `Titre: Deux encodages mémoriels pour les aires thalamiques intérmédiaires.

Auteurs: Y. Levain, A.P. Revertin, et J.O. Galakievicz.

Résumé: Nous présentons deux algorithmes en Oo( n2+log(g) ), probablement
epsilon-corrects, pour l'encodage numérique in vivo des activités mémorielles
et sous-conscentes des aires thalammiques humaines du cerveau humain. Le premier
s'appuie sur la librairie sous licence de Gantrell-HypeX, (récemment acquise par
Legba Voodoocom). Le second, dont l'espérence en fiabilité est légèrement
inférieure, sauf pour certains sous-types d'aires thalamiques, est entièrement
nouveau, et libre de droit de propriété.`},
		{"art652", []string{"Mandrake", "Kuipers", "Trebinsky"}, 1, "", "Un bootstrap efficace de l'ontologie phénoménologique", `Titre: Un bootstrap efficace de l'ontologie phénoménologique.

Auteurs: J. Kuipers et A. Trebinsky.

Résumé: L'un des écueils dans la Transcendance Numérique Forte est en passe
d'être levé. Nos premiers résultats expérimentaux, en simulation, montrent en
effet que notre méthode de bootstrap de l'ontologie phénoménologique primaire
permet une croissance quasi-exponentielle de la concordance proto-symbolique.
Il en découle logiquement une voie ouverte et prometteuse pour des
extraction-compressions réussie de la conscience humaine. Soulignons que cette
méta-accroissement peut se réaliser sans alignemnt mémoriel bas-niveaux, un
processus dont la compléxité doublement-exponentielle rend l'application pour
le moins délicate.`},
		{"art841", []string{"Mandrake", "Saint-Janvier", "Yu", "Wu", "Wellit-Ashley", "Modina"}, 1, "", "Une voie nouvelle pour la Conscience Numérique : la trans-fusion", `Titre: Une voie nouvelle pour la Conscience Numérique : la trans-fusion.

Auteurs: L. Saint-Janvier, D.L. Yu, C. Wu, O. Wellit-Ashley et L. Modina.

Résumé: A rebours des recherches "mainstream" dans le domaine de la Transcendance
Numérique (TN), nous présentons un cadre formel pour une voie originale et sobre
vers une Conscience Numérique Forte (au sens de Lashley). Notre idée, formulée
dans la logique épistémique modale de second ordre, s'appuie sur la fusion
multi-modale d'une IA de type A et d'une proto-numérisation des couches
superficielles et intermédiaire d'une personnalité humain handi-adaptée. Nous
avons passé notre proposition dans tous les vérificateurs symboliques de la base
de Lashley et tous donnent une probabilité de réussite dépassant les 60%.`},
		{"art9641", []string{"Mandrake", "San-Jorgeu", "Kuipers", "Trebinsky"}, 1, "", "Du déficit phénoménologique inéluctable des IA de classe A", `Titre: Du déficit phénoménologique inéluctable des IA de classe A.

Auteurs: R.D. San-Jorgeu, J. Kuipers et A. Trebinsky.

Résumé: Nos travaux sur le déficit phénoménologique des IA de classe A montrent
la faiblesse des approches trans-humain et handi-humaines. Nous argumentons
notamment sur les dangers, éthiques et moraux, des travaux mêlant handi-adaption
et ancrage artificielle pour l'émergence de conscience. Au delà de considérations
éthiques, nous mettons en exergue une analyse réfutatoire, par la méthode des
quanta-qbits transitionnels, qui démontre l'imfaisabilité des travaux de
[Saint-Janvier et. al] sur cette trans-fusion.`},
		{"art4251", []string{"Mandrake", "Pernu", "Itchinson", "Kanakuna"}, 1, "", "Symposium Inter-Coprporatiste sur la Transcendance Numérique", `Titre: Symposium Inter-Coprporatiste sur la Transcendance Numérique.

Editeurs: A. Pernu, C.H. Itchinson et P. Kanakuna.

Résumé: Ces deux journées de dialogues sur les dernières avancées en matie de
Transcendance Numériques furent profitables et porteuses de nouvelles voies de
recherche. Nous avons eu l'honneur d'animer ces journées qui permis des
propositions originales et pertinentes. Citons notamment les bootstrap
phénoménologiques de [Juipers et. al.], la trans-fusion IA/handihumanité
[Saint-Janvier et al.] et l'alignement thalamique de [A. Mathison].
Ce symposiuma eu lieu sous le patronnage de Legba Voodoocom et de SiliconSpirit.`},
		{"note99", []string{"Mandrake", "Kuipers"}, 1, "", "Fermeture du Projet Mandrake", `Titre: Fermeture du Projet Mandrake

Auteur: J. Kuipers.

Résumé: Après une analyse des travaux du Projet «Mandrake», j'ai décidé de le
clore et de transferer son budget au Projet «Phénomos». Le Dr. Alan Mathison
est relevé de son poste de directeur et ses accès révoqués. A toutes fin utiles,
j'ai demandé l'archivages du Projet «Mandrake».`},
	},
}

// fichier personnel : alan, harald, ragnar

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
	Entries: []Entry{
		{"propal01", []string{"deadzone", "achat", "dirty"}, 2, "", "Option d'achat", "Castle Corp: 1000 YES/m2"},
		{"propal02", []string{"deadzone", "achat", "dirty"}, 2, "", "Option d'achat", "Leisure United: 980 YES/m2"},
		{"propal03", []string{"deadzone", "achat", "dirty"}, 2, "", "Option d'achat", "Kramps: 500 YES/m2"},
	},
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
	Address: "leet.darknet",
	Credentials: []Cred{
		{"crunch", "hacktheplanet", 5},
		{"celine", "waytoocool", 4},
		{"nikki", "bohw4k", 4},
		{"greenglass", "brianglass", 3},
	},
	Description: cruDesc,
	Scan:        SEC3,
	Entries: []Entry{
		// ID keywords restricted owner title content
		{"vlope20", []string{"flr", "pr0n"}, 3, "", "vanessalope", `login: green pass: nait5zee`},
		{"bitecoin19", []string{"flr", "pr0n"}, 3, "", "lebitecoin", `login: green pass: ohphe0cu`},
		{"qtf20", []string{"flr", "pr0n"}, 3, "", "QueTesFans", `login: green pass: aesahm0l`},
		{"pndr20", []string{"flr", "pr0n"}, 3, "", "Pinederer", `login: green pass: ohdaf9uo`},
		{"jm20", []string{"flr", "pr0n"}, 3, "", "Jockey & Micheline", `login: green pass: eig0thob`},
		{"005672bR1An ", []string{"maman", "grocery"}, 3, "greenglass", "grocetag458", `tiger power, cheetos`},

		{"005673bR1An", []string{"maman", "pr0n"}, 3, "greenglass", "pr0n", `fistfukdenaines2mains.mov`},

		{"005674bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag006", `2006-2006 - The Solsburry Four - NeoLondon - LD-Aurora*Cobalt*Grim*Slice* - Validated contracts 17 - Estimated total number of contracts: 18 - Contract unfulfilled on 2006-06-07 Reason: Aurora Terminated On Duty | estimated error: 8.6%`},

		{"005675bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag459", `cotton tiges, beignets, chouquettes, semoule, cheetos`},

		{"005676bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag007", `2009-2014 - I Volteggiatori -NuevaRoma - LD-Phase*Pins*ShallowWater*Bull* - Validated contracts 7 - Estimated total number of contracts: 8 - Contract unfulfilled on 2009-09-14 Reason: Phase*ShallowWater Terminated On Duty | estimated error: 0.4%`},

		{"005677bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag460", `steaks hachés, poulet frit, dentifrice`},

		{"005678bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag461", `cheetos`},

		{"005679bR1An", []string{"maman", "pr0n", "movie"}, 3, "greenglass", "pr0n", `Mummypees.mov`},

		{"005680bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag008", `2012-2012 - Kanibale - VecchiaFirenze - Ego*Stitches*ZenHook*BlankSheet - Validated contracts 2 - Estimated total number of contracts: 4 - Contract unfulfilled on 2012-06-18 Reason: Blanksheet*Ego*Zenhook Terminated On Duty | estimated error: 16.005%`},

		{"005681bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag009", `2013-2015 - Credo zabójcy - NowaWarszawa - LD-Design*DoubleFeature*ShinyBone*Arch*DeadZone - Validated contracts 19 - Estimated total number of contracts: 19 - Contract unfulfilled 0 `},

		{"005682bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag462", `lingettes, saucisson, mozzarella, cheetos`},

		{"005683bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag463", `tomates, sparadra, bacon, oeufs, chocolat dur, cheetos`},

		{"005684bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag010", `2015-2018 - Liberty-Damen - NordenBerlin - LD-FlukeAnt*LibertyBell*Pins*Design - Validated contracts 31 - Estimated total number of contracts: 32 - Contract unfulfilled on 2018-05-16 Reason: Design*FlukeAnt Terminated On Duty | estimated error: 23.8%`},

		{"005685bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag464", `chaussettes, ceinture, poivron, poires, mousse, lait, steak de petits pois, escalope de carottes, cheetos (les nouveaux au goût vanille)`},

		{"005686bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag011", `2015-2019 - Wonton Soup - OldParis - LD-Cobalt*ImpPulse*Ink*SamouraiShowdown*SliceNDice - Validated contracts 28 - Estimated total number of contracts: 29 - Contract unfulfilled on 2019-01-24 Reason: Ink*SamouraiShodown Terminated On Duty | estimated error: 2.93%`},

		{"005687bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag011", `2018-2019 - King Kong Five - Europole - LD-AlcoLine*Bull*Mirror*PoisonClock - Validated contracts 5 - Estimated total number of contracts: 6 - Contract unfulfilled on 2019-12-01 Reason: Alcoline Missing On Duty | estimated error: 44.5%`},

		{"005688bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag011", `2018-2019 - RoadRunners - Europole -  LD-Mamasita*Pins*ZenHook*BrashBeast - Validated contracts 14 - Estimated total number of contracts: 14 - Contract unfulfilled 0`},

		{"005689bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag465", `Tampons, lingettes, mascara noire, mascara bleu, lait démaquillant, cheetos`},

		{"005690bR1An", []string{"maman", "pr0n"}, 3, "greenglass", "pr0n", `mummycums.mp3`},

		{"005691bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag011", `2018-2019 - BlastingRogues - Europole -  LD-Design*Silverpath*Stitches - Validated contracts 8 - Estimated total number of contracts: 9 - Contract unfulfilled on 2019-10-15 Reason: Cancelled contract | estimated error: 0.09%`},

		{"005692bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag466", `pain, crème pour le cul (hydratante, pas celle qui pique), yaourts hypoallergéniques, bananes roses, gin, whisky`},

		{"005693bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag467", `cheetos`},

		{"005694bR1An", []string{"maman", "music"}, 3, "greenglass", "musicbadass", `blackwavefullalbum.7zip`},

		{"005695bR1An", []string{"maman", "pr0n"}, 3, "greenglass", "pr0n", `fuckedbyherdogs.mov`},

		{"005696bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag468", `crème solaire, nougats, gants de vaisselle, frozeloops, petits sachets plastique, lait, cheetos`},

		{"005697bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag011", `2019-12-31+ -Trigger Legacy - Europole - LD-Cobalt*Design*Pins*Stitches*Bull* - Validated contracts 4 - Estimated total number of contracts: 5 - Contract unfulfilled 0 - Current Contracts: Contract A = Castle Corp. [Disney-Dassault] / estimated ¥€$ = 50.000 end 2020-07-31 | Contract B = Ubink Inc. [MetaSoft] / estimated ¥€$ = 25.000 / end 2020-08-10 |`},

		{"005698bR1An", []string{"maman", "contract"}, 3, "greenglass", "runteamtag011", `07/20/2020 x: [...] nb entrées 0 | total data: 10 | data hypothèses = estimation erreur: 0.002% au 07/20/2020`},

		{"005699bR1An", []string{"maman", "grocery"}, 3, "greenglass", "grocetag469", `cola, cheetos`},

		{"005700bR1An", []string{"maman", "pr0n"}, 3, "greenglass", "pr0n", `MummyshitsonDaddy.mov`},
	},
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
	Entries: []Entry{
		// id, keys, priv, owner, title, content
		{"hack874", []string{"bitlord", "privé", "assurance", "exit", "strategy"}, 5, "", "b>i>tlord", `
nom: Zhao Yung Wa

- serveur d'hébergement payant de films, séries, animés (voir attached file: list des ayant-droit lésés)
`},
		{"noob02", []string{"green", "glass", "privé", "assurance", "exit", "strategy"}, 5, "", "Green Glass", `
nom: Ruppert-Green Glass

- revente d'id (login/passd) sur réseau sociaux, pr0n, public corporate net 

`},
		{"hack66", []string{"dark", "sephiroth", "privé", "assurance", "exit", "strategy"}, 5, "", "Dark Sephiroth66", `
nom: Alexandre Pinchard

- hack Zelda «Final Chapter» (voir attached log files 276t5443) 
`},

		{"hack32", []string{"nikki", "privé", "assurance", "exit", "strategy"}, 5, "", "Nikki", `
nom: Nikole Jasinski

- hack Legba Voodoo Phone (voir attached log file 653BE32)
- denial of service, Disney Dassault, 02/15 (voir attached log file dass_54ab65)
- run sur Techmint, D22, 07/19, (voir attached log files ru6296EB, ru6296ec)

`},
		{"hack61", []string{"crunch", "privé", "assurance", "exit", "strategy"}, 5, "", "Crunch", `
nom: Frédéric Manson

- hack Legba Voodoo Phone (voir attached log file 653BE32)
- admin du l33t.darknet
- concepteur des software pirates leech3.5A, defile_v2.0 (see source files)
- run sur Techmint, D22, 07/19, (voir attached log files ru6296EB, ru6296ec)

`},
	},
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

// serveur de Hope
var hope = Server{
	Address: "hope.local",
	Credentials: []Cred{
		{"hope", "tearsintherain", 5},
		{"mel", "xxx", 3},
	},
	Description: hopDesc,
	Scan:        SEC5,
	Entries: []Entry{
		// journal personnel d'Alan
		{"alan00-07-07", []string{"alan", "journal", "2000"}, 3, "",
			"Fin du monde", `
Aujourd’hui est arrivé dans la Division R&D John-Mickael Fusion. C’est le
cadre Legba Voodoocom en charge de l’évaluation des actifs du département.
C’est un homme qui paraît sensé et prompt à percevoir les champs de recherche
prometteurs non seulement en termes de retombées financières immédiates,
mais aussi en termes de retombées médiatiques positives pour l’entreprise. Le
Terminal et Hope occupent toutes mes pensées mais je dois faire bonne figure
pour le Projet.`,
		},
		{"alan00-07-08", []string{"alan", "journal", "2000"}, 3, "",
			"Soulagement", `
J’ai invité John-Michael à une visite du Projet Mandrake, afin de lui présenter
nos travaux, et leurs avancées récentes suite à l’épisode Ragnar. Il est
convaincu de l’utilité de notre division et je vais pouvoir rassurer mon
équipe sur leur devenir si le rachat de Silicon Spirit est validé par la Cour
Corporatiste. Nous n’avons pas besoin d’un autre incident Proskychev.`,
		},
		{"alan00-07-09", []string{"alan", "journal", "2000"}, 3, "",
			"Quelle bande d'incompétents", `
En pensant à Hope ce soir, j’ai finalement mis le doigt sur ce qui me
tracassait dans l’article de Jordan Kuipers & Anthon trebinsky sur le boostrap
efficace de l’ontologie phénoménologique. C’est un tissu d’approximations et
d’erreurs dignes d’un première année. Je compte bien écrire une réfutation
détaillée et argumentée basée sur mes propres travaux, une fois la période de
rachat terminée. Ces derniers n’ont décidément rien compris.`,
		},
		{"alan00-07-10", []string{"alan", "journal", "2000"}, 3, "",
			"Legba est un serpent", `
Quel fils de chien galeux ! Quelle ordure syphilitique ! Ca ne se passera
pas comme ça, je vais leur coller mes avocats au cul !!! Me bloquer mon badge
d’accès, et faire mettre mes effets personnels au poste de garde d’entrée !
Comme un vulgaire voleur. Et avec ça, aucun accès à mes anciens postes de
travail, mes docs, sauvegardes, notes… `,
		},
		{"alan00-07-15", []string{"alan", "journal", "2000"}, 3, "",
			"Désespoir", `
L’avocat est plutôt clair, tout s’est fait dans la légalité, et je n’ai aucune
chance de récupérer les données et/ou le matériel appartenant à Silicon Spirit,
transmis à présent à Legba Voodoocom. Monde de merde ...`,
		},
		{"alan00-07-17", []string{"alan", "journal", "2000"}, 3, "",
			"Désabusé", `
... Hope me ressert un verre, et je contemple ma plus brillante réussite. Le
monde n’en saura jamais rien, j’espère qu’elle ne se laissera jamais pervertir
par notre profond besoin de surpasser nos contemporains, quel qu’en soit
le coût. Et dire que ces abrutis de Legba ont mis Kuipers à la tête de leur
version de MON Projet Mandrake !`,
		},
	},
}

// Project "Hope"
// Dépot mémoriel
// - fenêtre temporelle glissante v12.5
// - compression McVaugh-Korba-Yang
// - contenu infix normalisé
// - (c) A.M
var hopDesc = `
01010000 01110010 01101111 01101010 01100101 01100011
01110100  00100010 01001000 01101111 01110000 01100101 00100010
01000100 11101001 01110000 01101111 01110100  01101101
11101001 01101101 01101111 01110010 01101001 01100101 01101100
00101101  01100110 01100101 01101110 11101010 01110100 01110010
01100101  01110100 01100101 01101101 01110000 01101111 01110010 01100101
01101100 01101100 01100101  01100111 01101100 01101001 01110011 01110011
01100001 01101110 01110100 01100101  01110110 00110001 00110010 00101110
00110101  00101101  01100011 01101111 01101101 01110000 01110010
01100101 01110011 01110011 01101001 01101111 01101110  01001101 01100011
01010110 01100001 01110101 01100111 01101000 00101101 01001011 01101111
01110010 01100010 01100001 00101101 01011001 01100001 01101110 01100111
00101101  01100011 01101111 01101110 01110100 01100101 01101110 01110101
01101001 01101110 01100110 01101001 01111000  01101110 01101111 01110010
01101101 01100001 01101100 01101001 01110011 11101001
00101101  00101000 01100011 00101001  01000001 00101110 01001101                                                                                                                                                                                                                                    
`

var game = &Game{
	Network: []Server{
		dd,
		d22,
		kramps,
		kramps_pers,
		kramps_sec,
		kramps_inmates,
		corp,
		justice,
		abus,
		legba,
		legba_satcom,
		legba_archive,
		// lbd,
		greendata,
		leet,
		lair,
		hope,
	},
}
