# Les fichiers

main.go : point de départ de l'application, pas grand chose dedans
pour l'instant.

app.go : contient l'objet `App`, qui est le coeur de l'application :
gestion des connexions entrantes et de l'état interne du jeu.

game.go : l'état interne du jeu, stocké dans une BDD. Il y a quelques
autres trucs définis ici, mais ils pourraient aller dans des fichiers
séparés.

init.go : l'initialisation de l'état du jeu. A terme, l'exécution de
ce code devrait être optionnelle, pour pouvoir conserver l'état du
jeu entre deux démarrages.

client.go : le code de l'interface graphique exposée aux joueurs lors
de la connexion.

console.go : l'objet représentant la console utilisée par le joueur :
ses commandes, l'état de sa connexion ... Cet objet est ajouté à la
BDD lors de la création d'un client, et une copie de cet objet réside
dans le client.

command.go : le code utilisé pour le parsing des commandes.

cmd_connect.go cmd_help.go cmd_index.go cmd_quit.go : le code des
commandes accessibles au joueur

server.go : la représentation d'un serveur avec ses services et ses
mesures de sécurité.

# Outils

Sur cette nouvelle version, je m'appuie sur bubbletea, wish et storm

- bubbletea est une lib pour fabriquer des TUI
- wish est un serveur ssh permettant de servir ladite TUI
- storm est une base de données clef-valeur simple

