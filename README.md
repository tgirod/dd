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
- [ ] deconnexion ne change pas le ForumView !!
- [ ] tester en changeant de server
- [~] commandes
- [ ] functions basiques internes pour
  - [X] forum list
  - [X] forum read (un post ou un topic)
    - [X] trier/lister les réponses
    - [X] voir plus haut/plus bas/index 
    - [X] changer commandes Ctrl quand lit un Post
      - [X] Next/Previous in Thread ?????
  - [X] forum up (go up one topic)
  - [~] forum post (write a post to a topic)
    - [~] générer une date compatible avec le GN
    - [X] générer l'heure
    - [X] Pk les lignes du message sont passés en Cmd (pd qu'on écrit)
    - [ ] Vérifier que, si en train de lire un Post => AnswerPost 
  - [X] Answer a POST
  - [ ] Fonction pour quiter un Mode (Writing/Reading) proprement
  - [X] create new topic
  - [ ] ADMIN remove post ? (rm is enough ?)
  - [ ] ADMIN remove topic ? (rm -rf is enough ?)
