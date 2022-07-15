package main

// serveur local du dirty district
var dd = Server{}

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
var green = Server{}

// serveur privé de Crunch
var leet = Server{}

// serveur privé de Céline
var celine = Server{}

// serveur mémoriel de Hope
var hope = Server{}

// fanpage The Black Wave
var tbw = Server{}

func (g Game) Init() error {
	// nettoyer les tables
	g.Drop(&Server{})
	g.Drop(&Console{})

	// serveur de Jésus
	if err := g.Save(&Server{
		Address: "jesus",
		Credentials: []Cred{
			{"", "", 1},
		},
		Detection: 0.1,
		Gate: Gate{
			Targets: []Target{
				{
					Address:     "dist22",
					Description: "lien vers la grille du district 22",
					Restricted:  1,
					Privilege:   1,
				},
				{
					Address:     "dist22",
					Description: "lien vers la grille du district 22",
					Restricted:  3,
					Privilege:   3,
				},
			},
		},
		Database: Database{
			Description: "anthologie du frostpunk",
			Entries: []Entry{
				{
					Key:      "plas0",
					Keywords: []string{"plastobeton"},
					Title:    "du plasto sous les plages",
					Content:  "le meilleur album du monde",
				},
				{
					Key:      "tbw0",
					Keywords: []string{"blackwave"},
					Title:    "The Black Wave - première sommation",
					Content:  "le meilleur album du monde",
				},
			},
		},
	}); err != nil {
		return err
	}

	// serveur du district 22
	if err := g.Save(&Server{
		Address: "dist22",
		Credentials: []Cred{
			{"", "", 1},
		},
	}); err != nil {
		return err
	}

	return nil
}
