package main

// Server représente un serveur sur le Net
type Server struct {
	// informations générales
	ID      int    // ID du serveur (interne)
	Address string `storm:"unique"` // Addresse du serveur sur le réseau

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials []Cred

	// TODO backdoors
}

type Cred struct {
	login     string
	password  string
	privilege int
}

func (s Server) CheckCredentials(login, password string) (int, error) {
	var cred Cred
	for _, c := range s.Credentials {
		if c.login == login && c.password == password {
			return cred.privilege, nil
		}
	}
	return 0, errInvalidCredentials
}