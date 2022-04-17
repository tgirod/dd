package main

func (g Game) Init() error {
	// nettoyer les tables
	g.Drop(&Server{})
	g.Drop(&Console{})

	// serveur de Jésus
	if err := g.Save(&Server{
		Address: "jesus",
		Credentials: []Cred{
			{"public", "public", 1},
		},
		Gates: []Gate{
			{
				Service: Service{
					Name:        "dist22",
					Description: "accès au réseau du district 22",
					Restricted:  1,
				},
				TargetAddress: "dist22",
				Privilege:     1,
			},
			{
				Service: Service{
					Name:        "dist22-4",
					Description: "accès au réseau du district 22",
					Restricted:  4,
				},
				TargetAddress: "dist22",
				Privilege:     4,
			},
		},
		Databases: []Database{
			{
				Service: Service{
					Name:        "frostpunk",
					Description: "discographie",
					Restricted:  1,
				},
				Entries: []Entry{
					{
						Keywords: []string{"plastobéton"},
						Title:    "du plasto sous les plages",
						Content:  "le meilleur album du monde",
					},
					{
						Keywords: []string{"blackwave"},
						Title:    "The Black Wave - première sommation",
						Content:  "le meilleur album du monde",
					},
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
			{"public", "public", 1},
		},
	}); err != nil {
		return err
	}

	return nil
}
