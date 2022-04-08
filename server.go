package main

// Server représente un serveur sur le Net
type Server struct {
	// informations générales
	ID      int    // ID du serveur (interne)
	Address string `storm:"unique"` // Addresse du serveur sur le réseau

	// services
	Links     []Link
	Databases []Database

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials map[string]struct {
		password  string
		privilege int
	}

	// TODO backdoors
}

// Connect vérifie que la tentative de connexion est valide
func (s Server) Connect(login, password string) (int, error) {
	cred, ok := s.Credentials[login]
	if !ok {
		return 0, errInvalidLogin
	}

	if cred.password != password {
		return 0, errInvalidPassword
	}

	return cred.privilege, nil
}

func (s Server) Link(name string) (int, error) {
	return 0, errInvalidLink
}
