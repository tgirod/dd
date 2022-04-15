# Architecture du jeu

Le réseau est composé de serveurs contenant des services. Les
joueurs se déplacent dans le réseau pour accéder aux services qui
les intéressent. Certaines parties du réseau sont protégées, et
accessibles uniquement aux hackers.

## Que peut-on faire dans le Net ?

**Se déplacer**, soit en se connectant directement à un serveur (avec
son adresse), soit en suivant des liens qui permettent de passer d'un
serveur à un autre.

**Faire des recherches** dans des services de base de donnée. La fonction
de recherche se base sur des mots-clefs, ce qui est intéressant en
terme de jeu.

**Communiquer** via des comptes de messagerie. Depuis un service de
messagerie, on peut envoyer et recevoir des messages texte à n'importe
quel autre messagerie.

**Payer en ligne** pour des services divers et variés.

**Opérer des périphériques** connectés au réseau.

## Quel rôle pour les hackers ?

Tout le monde ne peut pas tout faire sur le Net - il existe des mesures
de sécurité pour que chacun reste à sa place. Le rôle d'un hacker
est de contourner ces mesures de protection pour faire des choses qui
seraient normalement interdites.

**L'accès à un serveur** nécessite d'avoir des accréditations valides
pour entrer.

**Le niveau de privilège** restreint l'accès à certains services ou
à certaines commandes.

**Les glaces** surveillent les activités illégales sur le serveur et
réagissent en conséquence.

On pourrait alors imaginer trois actions de hacker correspondantes :

**Accès illégal** : exploiter des vulnérabilités pour se connecter
illégalement à un serveur.

**Escalade de privilège** : obtenir un niveau d'accréditation plus
élevé.

**Cybercombat** : éviter ou planter les glaces pour rester plus longtemps
connecté.

## Pistes d'utilisation pendant le GN

Le serveur monté par Jésus contient un système de messagerie pour les
habitants du quartier ainsi que quelques liens permettant d'accéder à
la grille du district voisin.

La mémoire externe de Hope est stockée dans une base de donnée. On
peut y accéder si on en connait l'adresse (Hope utilise parfois les
terminaux pour y faire des recherches).

La grande évasion nécessite un soutien dans le Net. Un ou plusieurs
hackers doivent pénétrer les serveurs de la Kramps pour désactiver
certaines sécurités au bon moment.

La messagerie pourrait permettre des échanges avec des PNJ extérieurs
au quartier. Les contacts mafieux d'Escobar et Chilly Daisy, Peter Rabit,
les parents de Cramille ...

Le samedi, la connexion au Net est coupée, et elle est nécessaire pour
diffuser en direct l'éviction des habitants.

Toute personne possédant une ID peut l'utiliser pour se logger sur le
serveur bancaire, et accéder à son compte.

Les hackers ont des adresses de serveurs du darknet, dans lesquels
ils peuvent acheter toutes sortes de choses (par exemple des accès à
certains serveurs).

## Comment se déroule un hack ?

Un hacker arrive légalement dans un serveur corpo, comme visiteur. Il
souhaite accéder à la base de données du personnel, qui se trouve
dans le serveur administratif.

En tant que visiteur (cred=0), il n'a pas le niveau d'accréditation
nécessaire pour utiliser le lien qui mène au serveur administratif
(cred=2). Il injecte donc un programme illégal dans le serveur pour
augmenter son accréditation.

L'injection du programme illégal déclenche une alerte dans le
serveur. Ce n'est qu'une question de temps avant que le hacker ne soit
repéré. Le niveau d'alerte augmente progressivement.

Arrivé à un premier seuil, le programme illégal est repéré et
effacé du serveur. La recherche se poursuit pour trouver l'utilisateur
responsable de l'injection.

Le deuxième seuil atteint, le hacker est repéré. Selon le serveur,
il se fait déconnecter de force, ou se prend des dégats.

Le hacker profite rapidement de son niveau d'accréditation pour suivre
le lien qui mène au serveur administratif. Il arrive dans le serveur
souhaité, avec le même niveau d'accréditation. Le niveau d'alerte
est remis à zéro, mais l'alerte est toujours active. Le hacker a un
peu de temps pour accéder à la base de données du personnel.

Le niveau d'alerte commence à être dangereusement élevé, le hacker
sera bientôt repéré. Il charge un logiciel d'évasion qui lui permet
de brouiller les pistes. Le niveau d'alerte redescend un peu.

> comment faire pour que le joueur ne se contente pas de juste relancer
le programme d'évasion en boucle ? Peut être qu'on ne peut utiliser
un programme qu'une fois par run ?

Bientôt l'évasion ne suffit plus. Le hacker sort l'arme lourde :
il repère une glace en charge de la protection du serveur et la fait
planter. Mais elle ne tardera pas à revenir, et là ce sera la fin ...

# Développement

- architecture
	- [x] serveur SSH qui expose l'interface utilisateur (wish + bubbletea)
	- [x] stocker l'état du jeu dans la BDD
	- [ ] quand un client met à jour l'état du jeu, pousser cette
	mise à jour vers les autres clients.
	- [ ] supprimer l'objet `Console` de la BDD en cas de plantage du client
	- [ ] accès concurrent à la BDD - utiliser des transactions ?
- interface
	- [x] interface de base : barre de statut + output + prompt
	- [x] fenêtre modale, quand une commande nécessite d'afficher une interface spéciale
	- [ ] effet "typewriter"
	- [ ] effet "corruption de l'affichage"
	- [ ] scroll quand il y a trop de choses à afficher
- jeu
	- [ ] simuler une coupure complète de la connexion
	- [x] console, serveur
	- [x] ajouter les restrictions basées sur les accréditations
	- commandes
		- [x] connect : se connecter à un serveur
		- [x] help : aide sur l'usage des commandes
		- [x] index : lister les services
			- [ ] afficher uniquement le services pour lesquels on a les bons privilèges ?
		- [x] quit : déconnexion du serveur courant
		- [x] link : suit un lien de connexion vers un autre serveur
		- [ ] data : faire des recherches dans une BDD
		- [ ] msg : messagerie
		- [ ] pay : effectuer des paiements
		- [ ] edit : manipuler les registres d'un device branché
	- hacking
		- [ ] jack : force la connexion a un serveur dont on connait l'adresse
		- [ ] priv : monte le niveau de privilège dans un serveur
		- [ ] hide : retarde la traque des glaces
		- [ ] bomb : fait planter une glace

# Idées

Mini-jeu pour le hacking, que le joueur ait un peu plus à faire que d'entrer une commande ?