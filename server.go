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
