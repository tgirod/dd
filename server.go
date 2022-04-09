package main

// Server représente un serveur sur le Net
type Server struct {
	// informations générales
	Address string `storm:"id"` // Addresse du serveur sur le réseau

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials []Cred

	// les services hébergés
	Links []Link

	// TODO backdoors
}

type Cred struct {
	Login     string
	Password  string
	Privilege int
}

func (s Server) CheckCredentials(login, password string) (int, error) {
	for _, c := range s.Credentials {
		if c.Login == login && c.Password == password {
			return c.Privilege, nil
		}
	}

	return 0, errInvalidCredentials
}

func (s Server) FindLink(name string) (Link, error) {
	for _, l := range s.Links {
		if l.Name == name {
			return l, nil
		}
	}
	return Link{}, errLinkNotFound
}