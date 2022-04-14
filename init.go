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
