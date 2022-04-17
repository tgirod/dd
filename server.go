package main

// Server représente un serveur sur le Net
type Server struct {
	// informations générales
	Address string `storm:"id"` // Addresse du serveur sur le réseau

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials []Cred

	// les services hébergés
	Gates     []Gate
	Databases []Database

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

func (s Server) FindGate(name string) (Gate, error) {
	for _, l := range s.Gates {
		if l.Name == name {
			return l, nil
		}
	}
	return Gate{}, errGateNotFound
}

func (s Server) FindDatabase(name string) (Database, error) {
	for _, l := range s.Databases {
		if l.Name == name {
			return l, nil
		}
	}
	return Database{}, errDatabaseNotFound
}