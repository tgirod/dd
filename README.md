# Comment utiliser ce programme ?

installer go 1.18+

récupérer le code :

> git clone github.com/tgirod/dd

pour lancer le serveur et initialiser la base de données :

> make init

pour ouvrir un client :

> ssh localhost -p 1337

pour ouvrir, en même temps, un "monitor" avec des droits admin et des commandes supplémentaires :
> ssh localhost -p 7331

dirtydistricy_afterGN.db contient la base de donnée sauvée en fin de GN.

Dans le répertoire 'contenu', il a des utilitaires (faits à la va vite) pour produire des fichiers .yaml qui ont servi à peupler le "net" en début de GN.

