package main

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
					Keywords: []string{"plastobéton"},
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
