Architecture générale du projet

Sur cette nouvelle version, je m'appuie sur bubbletea, wish et storm

- bubbletea est une lib pour fabriquer des TUI 
- wish est un serveur ssh permettant de servir ladite TUI 
- storm est une base de données clef-valeur simple

# App

L'application principale contient le serveur SSH, et l'état interne
du jeu.

# Game

L'état interne du jeu, sous la forme du base de données
storm. Encapsuler la base de données sert surtout à ajouter quelques
méthodes pour simplifier certaines requêtes sur la BDD.

# Console

Lors de chaque connexion, une interface bubbletea est créée, ainsi
qu'un nouvel objet `Console`. Celui-ci est par la même occasion ajouté
à la BDD.

La console contient les commandes disponibles, ainsi que quelques
informations sur l'état courant de sa connexion.

Elle prend en charge le parsing des commandes saisies par l'utilisateur
et leur exécution sur le réseau.

Ce coup-ci, je me suis forcé à rester sur un truc un peu verbeux mais
simple pour ce qui est de la définition des commandes.

# Server

un serveur dans le réseau, fournissant des services.

- les services disponibles - les méthodes correspondant aux différentes
commandes utilisateur - la gestion des droits d'accès au serveur -
les fonctions de sécurité

# Service

un service accessible aux utilisateurs
