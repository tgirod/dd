package main

func (g Game) Init() error {
	s1 := Server{
		ID:      1,
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
