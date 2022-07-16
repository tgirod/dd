package main

// serveur local du dirty district
var dd = Server{
	Address: "dd.local",
	Credentials: []Cred{
		{"invite", "invite", 1},
		{"jesus", "roxor", 5},
	},
}

// serveur public du district 22
var d22 = Server{}

// serveur public de la kramps
var kramps = Server{}

// serveur privé de la kramps
var kramps_priv = Server{}

// serveur de sécurité de la kramps
var kramps_sec = Server{}

// serveur des services corporatistes D22
var central = Server{}

// serveur bancaire du D22
var abus = Server{}

// serveur public de Legba Voodoocom
var legba = Server{}

// serveur privé de la communication satellite
var legba_satcom = Server{}

// serveur archive de Silicon Spirit
var legba_archive = Server{}

// serveur le bon district
var lbd = Server{}

// green data
var greendata = Server{}

// serveur privé de Crunch
var leet = Server{}

// serveur privé de Céline
var celine = Server{}

// serveur mémoriel de Hope
var hope = Server{}

// fanpage The Black Wave
var tbw = Server{}

// TODO remplir le jeu
var game = &Game{
	Network: []Server{
		dd,
		d22,
		kramps,
		kramps_priv,
		kramps_sec,
		central,
		abus,
		legba,
		legba_satcom,
		legba_archive,
		lbd,
		greendata,
		leet,
		celine,
		hope,
		tbw,
	},
}
