package main

// serveur local du dirty district
var dd = Server{
	Address: "dd.local",
	Credentials: []Cred{
		{"invite", "invite", 1},
		{"jesus", "roxor", 5},
	},
	Description: ddDesc,
	Targets: []Target{
		{d22.Address, "serveur public du District 22", 1, 1},
	},
	Registers: []Register{
		{"cafe", false, "machine à café", 1},
		{"sono", true, "et je coupe le son ...", 3},
	},
	Detection: 0.1,
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
		{"invite", "invite", 1},
	},
	Description: `Bienvenue sur le serveur public du District 22 d'Europole.`,
	Targets: []Target{
		{legba.Address, "Legba Voodocom", 1, 1},
		{kramps.Address, "Kramps Securty", 1, 1},
		{corp.Address, "Central Services", 1, 1},
		{abus.Address, "Association des Banques Unifiées Suisses", 1, 1},
		{greendata.Address, "Green Data, solution environnementale", 1, 1},
	},
}

// serveur public de la kramps
var kramps = Server{
	Address: "kramps.d22.eu",
	Targets: []Target{
		{kramps_priv.Address, "Serveur réservé au personnel", 3, 1},
	},
}

// serveur privé de la kramps
var kramps_priv = Server{
	Address: "priv.kramps.d22.eu",
	Targets: []Target{
		{kramps_sec.Address, "Serveur central de sécurité", 5, 1},
	},
}

// serveur de sécurité de la kramps
var kramps_sec = Server{
	Address: "sec.kramps.d22.eu",
}

// serveur des services corporatistes D22
var corp = Server{
	Address: "corp.d22.eu",
}

// serveur bancaire du D22
var abus = Server{
	Address: "abus.d22.eu",
}

// serveur public de Legba Voodoocom
var legba = Server{
	Address: "legba.d22.eu",
	Targets: []Target{
		{legba_satcom.Address, "division sat-com", 3, 1},
		{legba_archive.Address, "archives", 3, 1},
	},
}

// serveur privé de la communication satellite
var legba_satcom = Server{
	Address: "satcom.legba.d22.eu",
}

// serveur archive de Silicon Spirit
var legba_archive = Server{
	Address: "archive.legba.d22.eu",
}

// serveur le bon district
var lbd = Server{
	Address: "lebondistrict.d22.eu",
}

// green data
var greendata = Server{
	Address: "greendata.d22.eu",
}

// serveur privé de Crunch
var leet = Server{
	Address: "l33t.darknet",
}

// serveur privé de Céline
var celine = Server{
	Address: "celine.darknet",
}

// serveur mémoriel de Hope
var hope = Server{
	Address: "hope.local",
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
		celine,
		hope,
	},
}
