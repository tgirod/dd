Architecture générale du projet

# App

l'application principale, créée par `main`

- gérer les connexions entrantes
- stocker l'état interne du jeu

# Console

la connexion d'un joueur via un terminal

- la connexion ssh avec le client
- l'interface utilisateur présentée au client
- le parsing des commandes
- des infos de jeu sur l'état de la connexion (alerte, dégats etc)

# Network

le réseau auquel les joueurs se connectent

- les serveurs et leurs adresses associées

# Server

un serveur dans le réseau, fournissant des services

- les services disponibles
- les méthodes correspondant aux différentes commandes utilisateur
- la gestion des droits d'accès au serveur
- les fonctions de sécurité

# Service

un service accessible aux utilisateurs
