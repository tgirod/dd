# Comment utiliser ce programme ?

installer go 1.18

1. go build
2. ./dd
3. ssh localhost -p 1337

# TODO

fonctionnalités les plus importantes

- [x] se déplacer
- [x] faire des recherches
- [x] restreindre l'accès à certaines ressources
- [x] distinguer un hacker d'un personnage normal (load ?)
- [x] forcer une connexion
- [x] augmenter ses privilèges
- [x] ralentir une glace
- [ ] envoyer des messages (annulé, ou remplacé par un BBS)
- [ ] interaction entre plusieurs joueurs (annulé)

le jeu

- [x] topologie du réseau
- [x] accréditations pour accéder aux divers serveurs
- [ ] modifier les codes pour charger les logiciels de hack (et matérialiser les logiciels avec des clefs USB)

## Interface

- [x] interface de base : barre de statut + output + prompt
- [ ] affichage progressif à l'écran (annulé)
- [ ] scroll quand il y a trop de choses à afficher

## Forum
- fichiers dans rep serveur/forum/topic/subtopic/...
- fichier nom = date_hhmmss_title_origin
- [ ] tester title avec des '_' dedans !!!!
- [X] Esc/q pour quitter le Forum
- [ ] deconnexion ne change pas le ForumView !!
- [ ] tester en changeant de server
- [ ] Forum : module autonome dans fenêre modale de client
  - [ ] readPost : fenêtre modale dans la fenêtre modale
  - [X ] une seule commande : forum
- [ ] commandes
  - [ ] read Post
  - [ ] write Post
  - [ ] answer Post
  - [ ] add Topic
- [ ] ADMIN remove post ? (rm is enough ?)
- [ ] ADMIN remove topic ? (rm -rf is enough ?)
