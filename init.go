package main

func (g Game) Init() error {
	// nettoyer les tables
	g.Drop(&Server{})
	g.Drop(&Console{})
	g.Drop(&Link{})

	s1 := Server{
		Address: "jesus",
		Credentials: []Cred{
			{"invite", "invite", 1},
		},
	}

	if err := g.Save(&s1); err != nil {
		return err
	}

	return nil
}
