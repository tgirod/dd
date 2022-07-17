package main

const (
	SEC1 = 0.05
	SEC2 = 0.1
	SEC3 = 0.2
	SEC4 = 0.4
	SEC5 = 0.8
)

const (
	mel            = "GGCCAAAGCTCCTTCGGAGC"
	rocky          = "CCGCGCAGAATCATAGCTGT"
	rita           = "CAAAGTTCTAGGCATAGGGA"
	styx           = "TTAGCTCGATATCCTAACCC"
	kapo           = "GAACTGCTTTAGTTGACGGA"
	scalpel        = "TGAAAGAGACATGATGCCTT"
	greko          = "TCTGAGGTTTATTGATTTCG"
	jesus          = "TTCGGGATTACTGCGTGCTG"
	escobar        = "GGAGGACACCCCAAACGCAT"
	cageot         = "GCCCTTGTCATGTACTTAGT"
	lafouine       = "CTGTCACCCAATCTACAGCG"
	eva            = "CTGTTGTAGTGACATGTTTC"
	fatmike        = "AACCTTGGGCACGGTCGGTA"
	kennedy        = "CCCGCGGGCAAAGCTGACAG"
	savagegirl     = "GGGTCTATAGGTCAAACGGT"
	raoulcool      = "GTCACAAGGTTGTTTAATGG"
	greenglass     = "ATGCCTACCTCCAATGATTA"
	chillydaisy    = "CGGGAGACACGTTCAGTCTT"
	frereping      = "GCATGGCCGAATTCCTCATT"
	papaproxy      = "CGATTTGTATTGGATACGGA"
	nikki          = "ACGAACCTAGAGCCGCACGC"
	celine         = "CGCTCCCATTTCATGTCAGC"
	cramille       = "TTTGGGAGAAGCTTATGCAC"
	tigerdoll      = "ATATGTTGAGCGTAAAGGCG"
	sistermorphine = "CCATCCGGCGGACCTTATGC"
	zilmir         = "GACGGGATACCTACTCTCGA"
	bettyb         = "ATTCCGACTCAGGGTACCGG"
	abraham        = "TGGCGTCTCTAATTCTTGCC"
	crunch         = "TTCAAGCTGAATATGAAAGG"
	onekick        = "GTCAAATCTGAGACTCTTGC"
	jacob          = "TGAAAGAGACAGTATGCCGT"
	gang1          = "TTCGACTGAATGTTTGATGT"
	gang2          = "TATCGACGCACGGGACTTGG"
	gang3          = "CGAGAAATGACAGAGTTGTA"
	paula          = "GGGTGATCTGTTGCCCCCTG"
	ringo          = "AACTGACGGATTCGATCATG"
	georges        = "GTTTGCACGGAACATGCAAC"
	jeanne         = "GACCCGTATTTCGCTGATTG"
	oggy           = "TCAGCTTCTAACGTTCGGGA"
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
	},
	Detection: SEC1,
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
	Description: `Bienvenue sur le serveur public du District 22 d'Europole.`,
	Targets: []Target{
		{legba.Address, "Legba Voodoocom", 1, "public", "public"},
		{kramps.Address, "Kramps Security", 1, "public", "public"},
		{corp.Address, "Central Services", 1, "public", "public"},
		{abus.Address, "Association des Banques Unifiées Suisses", 1, "public", "public"},
		{greendata.Address, "Green Data, solution environnementale", 1, "public", "public"},
	},
	Detection: SEC2,
}

// serveur public de la kramps
var kramps = Server{
	Address: "kramps.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Targets: []Target{
		{kramps_priv.Address, "Serveur réservé au personnel", 3, "personnel", "kramps1234"},
	},
}

// serveur privé de la kramps
var kramps_priv = Server{
	Address: "priv.kramps.d22.eu",
	Credentials: []Cred{
		{"personnel", "kramps1234", 3},
	},
	Targets: []Target{
		{kramps_sec.Address, "Serveur central de sécurité", 5, "admin", "lkjqsod"},
	},
}

// serveur de sécurité de la kramps
var kramps_sec = Server{
	Address: "sec.kramps.d22.eu",
	Credentials: []Cred{
		{"admin", "lkjqsod", 5},
	},
}

// serveur des services corporatistes D22
var corp = Server{
	Address: "corp.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
}

// serveur bancaire du D22
var abus = Server{
	Address: "abus.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
}

// serveur public de Legba Voodoocom
var legba = Server{
	Address: "legba.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Targets: []Target{
		{legba_satcom.Address, "division sat-com", 5, "admin", "satcom9876"},
		{legba_archive.Address, "archives", 3, "personnel", "archive6543"},
	},
}

// serveur privé de la communication satellite
var legba_satcom = Server{
	Address: "satcom.legba.d22.eu",
	Credentials: []Cred{
		{"admin", "satcom9876", 5},
	},
}

// serveur archive de Silicon Spirit
var legba_archive = Server{
	Address: "archive.legba.d22.eu",
	Credentials: []Cred{
		{"personnel", "archive6543", 3},
	},
}

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
}

// serveur privé de Crunch
var leet = Server{
	Address: "l33t.darknet",
	Credentials: []Cred{
		{"crunch", "hacktheplanet", 5},
	},
}

// serveur privé de Céline
var lair = Server{
	Address: "celine.darknet",
	Credentials: []Cred{
		{"celine", "waytoocool", 5},
	},
}

// serveur mémoriel de Hope
var hope = Server{
	Address: "hope.local",
	Credentials: []Cred{
		{"hope", "tearsintherain", 5},
	},
}

var game = &Game{
	Network: []Server{
		dd,
		d22,
		kramps,
		kramps_priv,
		kramps_sec,
		corp,
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
