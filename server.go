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
	Login     string
	Password  string
	Privilege int
}

func (s Server) CheckCredentials(login, password string) (int, error) {
	var cred Cred
	for _, c := range s.Credentials {
		if c.Login == login && c.Password == password {
			return cred.Privilege, nil
		}
	}
	return 0, errInvalidCredentials
}