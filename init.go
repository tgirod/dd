package main

import (
	"fmt"
	"log"
)

const (
	SEC1 int = iota
	SEC2
	SEC3
	SEC4
	SEC5
)

func InitNetwork(
	identities []Identity,
	transactions []Transaction,
) {
	log.Println("identités")
	for _, i := range identities {
		log.Println("\t", i.Login)
		if _, err := Save(i); err != nil {
			log.Fatalf("%v : %v\n", i, err)
		}
	}

	log.Println("transactions")
	for _, t := range transactions {
		log.Println("\t", t.From, t.To, t.Yes)
		if _, err := Save(t); err != nil {
			log.Fatalf("%v : %v\n", t, err)
		}
	}
}

func InitServer(
	s Server,
	users []User,
	links []Link,
	registers []Register,
	posts []Post,
) {
	addr := s.Address
	if addr == "" {
		panic("le serveur n'a pas d'adresse")
	}

	log.Println("server", s.Address)
	if _, err := Save(s); err != nil {
		log.Fatal(err)
	}

	log.Println("users")
	for _, a := range users {
		log.Println("\t", a.Login)
		a.Server = addr
		if _, err := Save(a); err != nil {
			log.Fatalf("%v : %v\n", a, err)
		}
	}
	log.Println("links")
	for _, l := range links {
		log.Println("\t", l.Address)
		l.Server = addr
		if _, err := Save(l); err != nil {
			log.Fatalf("%v : %v\n", l, err)
		}
	}
	log.Println("registers")
	for _, r := range registers {
		log.Println("\t", r.Description)
		r.Server = addr
		if _, err := Save(r); err != nil {
			log.Fatalf("%v : %v\n", r, err)
		}
	}
	log.Println("posts")
	for _, p := range posts {
		log.Println("\t", p.Subject)
		p.Server = addr
		if _, err := Save(p); err != nil {
			log.Fatalf("%v : %v\n", p, err)
		}
	}
}

func Reset() {
	db.Drop(Identity{})
	db.Drop(Message{})
	db.Drop(Server{})
	db.Drop(User{})
	db.Drop(Link{})
	db.Drop(Register{})
	db.Drop(Post{})
	db.Drop(Transaction{})
}

var ddDesc = `
 ____  _      _           ____  _     _        _      _
|  _ \(_)_ __| |_ _   _  |  _ \(_)___| |_ _ __(_) ___| |_
| | | | | '__| __| | | | | | | | / __| __| '__| |/ __| __|
| |_| | | |  | |_| |_| | | |_| | \__ \ |_| |  | | (__| |_
|____/|_|_|   \__|\__, | |____/|_|___/\__|_|  |_|\___|\__|
                  |___/
Bienvenue sur le serveur communautaire du Dirty District.

Ce serveur est connecté au Net par le biais d'un accès illégal. Merci de ne pas
faire n'importe quoi.

******************************************************************************
**** DÉSOLÉ : suite au crash de la semaine, on n'a pas réussi à tout récuperer
************************************************************* (fuck) *********

Tape "index" pour avoir la liste des services fournis par le serveur. Si tu as
besoin d'aide, demande à ton nerd préféré.`

var dd = Server{
	Address:     "dd.local",
	Description: ddDesc,
	Security:    SEC1,
}

var frozDesc = `
  █████▒██▀███   ▒█████  ▒███████▒ ██▓███   █    ██  ███▄    █  ██ ▄█▀
▓██   ▒▓██ ▒ ██▒▒██▒  ██▒▒ ▒ ▒ ▄▀░▓██░  ██▒ ██  ▓██▒ ██ ▀█   █  ██▄█▒
▒████ ░▓██ ░▄█ ▒▒██░  ██▒░ ▒ ▄▀▒░ ▓██░ ██▓▒▓██  ▒██░▓██  ▀█ ██▒▓███▄░
░▓█▒  ░▒██▀▀█▄  ▒██   ██░  ▄▀▒   ░▒██▄█▓▒ ▒▓▓█  ░██░▓██▒  ▐▌██▒▓██ █▄
░▒█░   ░██▓ ▒██▒░ ████▓▒░▒███████▒▒██▒ ░  ░▒▒█████▓ ▒██░   ▓██░▒██▒ █▄
 ▒ ░   ░ ▒▓ ░▒▓░░ ▒░▒░▒░ ░▒▒ ▓░▒░▒▒▓▒░ ░  ░░▒▓▒ ▒ ▒ ░ ▒░   ▒ ▒ ▒ ▒▒ ▓▒
 ░       ░▒ ░ ▒░  ░ ▒ ▒░ ░░▒ ▒ ░ ▒░▒ ░     ░░▒░ ░ ░ ░ ░░   ░ ▒░░ ░▒ ▒░
 ░ ░     ░░   ░ ░ ░ ░ ▒  ░ ░ ░ ░ ░░░        ░░░ ░ ░    ░   ░ ░ ░ ░░ ░
          ░         ░ ░    ░ ░                ░              ░ ░  ░
                         ░
Ici ça cause FrozPunk, et de rien d'autre !

                             Sauf si vous voulez.....
`
var frozdd = Server{
	Address:     "froz.dd.local",
	Private:     false,
	Description: frozDesc,
	Security:    SEC1,
}

var maravdd = Server{
	Address:     "marav.dd.local",
	Private:     false,
	Description: maravDesc,
	Security:    0,
}

var maravDesc = `
                LA MARAV' PRÉSENTE                                                                        

        █████████    █████████   ███████████                   
       ███░░░░░███  ███░░░░░███ ░█░░░███░░░█                   
      ░███    ░░░  ░███    ░███ ░   ░███  ░                    
      ░░█████████  ░███████████     ░███                       
       ░░░░░░░░███ ░███░░░░░███     ░███                       
       ███    ░███ ░███    ░███     ░███                       
      ░░█████████  █████   █████    █████    ██                
       ░░░░░░░░░  ░░░░░   ░░░░░    ░░░░░    ░░                 

                                                                        
     ███████    ███████████  ██████████ ██████   █████                   
   ███░░░░░███ ░░███░░░░░███░░███░░░░░█░░██████ ░░███                    
  ███     ░░███ ░███    ░███ ░███  █ ░  ░███░███ ░███                    
 ░███      ░███ ░██████████  ░██████    ░███░░███░███                    
 ░███      ░███ ░███░░░░░░   ░███░░█    ░███ ░░██████                    
 ░░███     ███  ░███         ░███ ░   █ ░███  ░░█████                    
  ░░░███████░   █████        ██████████ █████  ░░█████                   
    ░░░░░░░    ░░░░░        ░░░░░░░░░░ ░░░░░    ░░░░░                    
                                                                        
                                                                        
 ███████████ █████   █████████  █████   █████ ███████████               
░░███░░░░░░█░░███   ███░░░░░███░░███   ░░███ ░█░░░███░░░█               
 ░███   █ ░  ░███  ███     ░░░  ░███    ░███ ░   ░███  ░                
 ░███████    ░███ ░███          ░███████████     ░███                   
 ░███░░░█    ░███ ░███    █████ ░███░░░░░███     ░███                   
 ░███  ░     ░███ ░░███  ░░███  ░███    ░███     ░███                   
 █████       █████ ░░█████████  █████   █████    █████                  
░░░░░       ░░░░░   ░░░░░░░░░  ░░░░░   ░░░░░    ░░░░░                   

     TIGER DOLL       vs       NINO FIREGUN BENVENUTI

pour les parieurs locaux, voyez avec Oggy
`

var dd22Desc = `
                    _____                            _                       
                   |  ___|                          | |                      
                   | |__ _   _ _ __ ___  _ __   ___ | | ___                  
                   |  __| | | | '__/ _ \| '_ \ / _ \| |/ _ \                 
                   | |__| |_| | | | (_) | |_) | (_) | |  __/                 
                   \____/\__,_|_|  \___/| .__/ \___/|_|\___|                 
                                        | |                                  
                                        |_|                                  
               ______ _     _        _      _     _____  _____               
               |  _  (_)   | |      (_)    | |   / __  \/ __  \              
  ______ ___   | | | |_ ___| |_ _ __ _  ___| |_   ' / /' ' / /'   ___ ______ 
 |______/ _ \  | | | | / __| __| '__| |/ __| __|   / /    / /    / _ \______|
       | (_) | | |/ /| \__ \ |_| |  | | (__| |_  ./ /___./ /___ | (_) |      
        \___/  |___/ |_|___/\__|_|  |_|\___|\__| \_____/\_____/  \___/       
                                                                             

           Bienvenue sur le serveur public du District 22 d'Europole.
           Un noeud du plus grand fournisseur d'accès de Méga-Europe.`

var d22 = Server{
	Address:     "d22.eu",
	Private:     false,
	Description: dd22Desc,
	Security:    SEC3,
}

// serveur public de la kramps
var kramps = Server{
	Address:     "kramps.d22.eu",
	Private:     false,
	Description: kpubDesc,
	Security:    SEC2,
}

var kpubDesc = `
      ___           ___           ___           ___           ___         ___
     /__/|         /  /\         /  /\         /__/\         /  /\       /  /\
    |  |:|        /  /::\       /  /::\       |  |::\       /  /::\     /  /:/_
    |  |:|       /  /:/\:\     /  /:/\:\      |  |:|:\     /  /:/\:\   /  /:/ /\
  __|  |:|      /  /:/~/:/    /  /:/~/::\   __|__|:|\:\   /  /:/~/:/  /  /:/ /::\
 /__/\_|:|____ /__/:/ /:/___ /__/:/ /:/\:\ /__/::::| \:\ /__/:/ /:/  /__/:/ /:/\:\
 \  \:\/:::::/ \  \:\/:::::/ \  \:\/:/__\/ \  \:\~~\__\/ \  \:\/:/   \  \:\/:/~/:/
  \  \::/~~~~   \  \::/~~~~   \  \::/       \  \:\        \  \::/     \  \::/ /:/
   \  \:\        \  \:\        \  \:\        \  \:\        \  \:\      \__\/ /:/
    \  \:\        \  \:\        \  \:\        \  \:\        \  \:\       /__/:/
     \__\/         \__\/         \__\/         \__\/         \__\/       \__\/

                                                                +-+-+-+-+-+-+-+-+
                                                                |S|e|c|u|r|i|t|y|
   _|_  | | _   |\/| _  _  _| _    _ |    _  (~ ∧  _  _|_       +-+-+-+-+-+-+-+-+
    !   |_|| |  |  |(_)| |(_|(/_  |_)||_|_\  _)|_||    !        | | | | | | | | |
                                  |                             | | | | | | | | |
                                                                | | | | | | | | |
                                                                | | | | | | | | |
                                                                | | | | | | | | |
                    Plus de 4.000.000.000 ¥€$ de capital        | | | | | | | | |
`

var kperDesc = `
      ___           ___           ___           ___           ___         ___
     /__/|         /  /\         /  /\         /__/\         /  /\       /  /\
    |  |:|        /  /::\       /  /::\       |  |::\       /  /::\     /  /:/_
    |  |:|       /  /:/\:\     /  /:/\:\      |  |:|:\     /  /:/\:\   /  /:/ /\
  __|  |:|      /  /:/~/:/    /  /:/~/::\   __|__|:|\:\   /  /:/~/:/  /  /:/ /::\
 /__/\_|:|____ /__/:/ /:/___ /__/:/ /:/\:\ /__/::::| \:\ /__/:/ /:/  /__/:/ /:/\:\
 \  \:\/:::::/ \  \:\/:::::/ \  \:\/:/__\/ \  \:\~~\__\/ \  \:\/:/   \  \:\/:/~/:/
  \  \::/~~~~   \  \::/~~~~   \  \::/       \  \:\        \  \::/     \  \::/ /:/
   \  \:\        \  \:\        \  \:\        \  \:\        \  \:\      \__\/ /:/
    \  \:\        \  \:\        \  \:\        \  \:\        \  \:\       /__/:/
     \__\/         \__\/         \__\/         \__\/         \__\/       \__\/

  [-> Serveur du personnel <-]++ toutes les transaction sont loguées (SecLvl 4)

  20/07: Rappel ! Les EdT sont calculés chaque nuit pour le lendemain.
         Aucune réclamation n'est recevable.

  18/07: StuKa est la 3° organisation du personnel (spécialisée Tech niv 1-3) à
         déposer le bilan cette année.

  02/07: Rappel ! Les dossiers de promotions pour SecAgent, tout niveau, sont à
         renvoyer avant le 31/07. Tarif habituel, voir agence comptable.
`

// serveur privé de la kramps
var persKramps = Server{
	Address:     "priv.kramps.d22.eu",
	Private:     true,
	Description: kperDesc,
	Security:    SEC3,
}

// Accounts: []Account{
// 	{Login: "akremmer"},
// 	{Login: "haxxor", Backdoor: true},
// },
// Links: []Link{
// 	{kramps_inmates.Address, "Gestion des prisonniers"},
// 	{kramps_sec.Address, "Sécurité des installations"},
// },

var ksecDesc = `
      ___           ___           ___           ___           ___         ___
     /__/|         /  /\         /  /\         /__/\         /  /\       /  /\
    |  |:|        /  /::\       /  /::\       |  |::\       /  /::\     /  /:/_
    |  |:|       /  /:/\:\     /  /:/\:\      |  |:|:\     /  /:/\:\   /  /:/ /\
  __|  |:|      /  /:/~/:/    /  /:/~/::\   __|__|:|\:\   /  /:/~/:/  /  /:/ /::\
 /__/\_|:|____ /__/:/ /:/___ /__/:/ /:/\:\ /__/::::| \:\ /__/:/ /:/  /__/:/ /:/\:\
 \  \:\/:::::/ \  \:\/:::::/ \  \:\/:/__\/ \  \:\~~\__\/ \  \:\/:/   \  \:\/:/~/:/
  \  \::/~~~~   \  \::/~~~~   \  \::/       \  \:\        \  \::/     \  \::/ /:/
   \  \:\        \  \:\        \  \:\        \  \:\        \  \:\      \__\/ /:/
    \  \:\        \  \:\        \  \:\        \  \:\        \  \:\       /__/:/
     \__\/         \__\/         \__\/         \__\/         \__\/       \__\/

  [-> Serveur Sécurité <-]        ++ toutes les transaction sont loguées (SecLvl 4)

  Vous avez suivi la formation <sécurité> obligatoire, mais nous vous rappelons les
  règles essentielles :
       #1) Respectez la vie privée des autres.
       #2) Réfléchissez avant de taper.
       #3) De grands pouvoirs impliquent de grandes responsabilités.
`
var secKramps = Server{
	Address:     "sec.kramps.d22.eu",
	Private:     true,
	Description: ksecDesc,
	Security:    SEC4,
}

var elecKramps = Server{
	Address:     "elec.kramps.d22.eu",
	Private:     true,
	Description: kElecDesc,
	Security:    SEC2,
}

var kElecDesc = `
██╗  ██╗██████╗  █████╗ ███╗   ███╗██████╗ ███████╗
██║ ██╔╝██╔══██╗██╔══██╗████╗ ████║██╔══██╗██╔════╝
█████╔╝ ██████╔╝███████║██╔████╔██║██████╔╝███████╗
██╔═██╗ ██╔══██╗██╔══██║██║╚██╔╝██║██╔═══╝ ╚════██║
██║  ██╗██║  ██║██║  ██║██║ ╚═╝ ██║██║     ███████║
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝     ╚══════╝
                                                   
███████╗███╗   ██╗███████╗██████╗  ██████╗██╗   ██╗
██╔════╝████╗  ██║██╔════╝██╔══██╗██╔════╝╚██╗ ██╔╝
█████╗  ██╔██╗ ██║█████╗  ██████╔╝██║  ███╗╚████╔╝ 
██╔══╝  ██║╚██╗██║██╔══╝  ██╔══██╗██║   ██║ ╚██╔╝  
███████╗██║ ╚████║███████╗██║  ██║╚██████╔╝  ██║   
╚══════╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   
	     Gestion du réseau électrique
	L'énergie c'est le pouvoir, utilisons-là
`

// serveur des services corporatistes D22
var corp = Server{
	Address:     "corp.d22.eu",
	Private:     false,
	Description: cd22Desc,
	Security:    SEC3,
}

var cd22Desc = `

   ((ervices  ((orporatistes
   ''
    _
   [|)istrict  22

   Ce service **public** vous est proposé **gratuitement** par la Cour Corporatiste.

   Ce service est livré en l'état, et la Cours Corporatiste décline toute responsabilité
   en ce qui concerne les données présentes et l'usage qui en est fait.

   Ce site existe gràce à la généreuse participation de Weyland-Yutani Corp,
   Tyrel Corp, Tessier-Ashpool SA, Disney Dassault, Arasaka, Renraku, Ubik,
   Legba Voodoocom, Avalon, Association des Banques Unifiées Suisses (ABUS).
`

// serveur judiciaire
var justice = Server{
	Address:     "justice.corp.d22.eu",
	Private:     false,
	Description: cd22justDesc,
	Security:    SEC5,
}

var cd22justDesc = `
   ((ervices  ((orporatistes
   ''
    _
   [|)istrict  22   === Département JUDICIAIRE ===

   Ce service **public** vous est proposé **gratuitement** par la Cours Corporatiste.

   Ce service est livré en l'état, et la Cours Corporatiste décline toute responsabilité
   en ce qui concerne les données présentes et l'usage qui en est fait.

   Ce site existe gràce à la généreuse participation de Weyland-Yutani Corp,
   Tyrel Corp, Tessier-Ashpool SA, Disney Dassault, Arasaka, Renraku, Ubik,
   Legba Voodoocom, Avalon, Association des Banques Unifiées Suisses (ABUS).
`

// serveur bancaire du D22
var abus = Server{
	Address:     "abus.d22.eu",
	Private:     false,
	Security:    SEC5,
	Description: cd22bankDesc,
}

var cd22bankDesc = `
   ((ervices  ((orporatistes
   ''
    _
   [|)istrict  22   === Département BANCAIRE ===

   Ce service **public** est assuré par
                                          _______     ______                 _______
                       Association       (  ___  )   (  ___ \   |\     /|   (  ____ \
                                         | (   ) |   | (   ) )  | )   ( |   | (    \/
                       des Banques       | (___) |   | (__/ /   | |   | |   | (_____
                                         |  ___  |   |  __ (    | |   | |   (_____  )
                          Unifiées       | (   ) |   | (  \ \   | |   | |         ) |
                                         | )   ( | _ | )___) )_ | (___) | _ /\____) |
                           Suisses       |/     \|(¥)|/ \___/(€)(_______)($)\_______)
`

var (
	lbDesc = `
                 ......
           .',,,,,,,,,,,,,,,.
        .;;,'.            .',;;'
      ':;.   ...  ,odooodkxxdc;;:,
    .::..;:oOKXXOOXWWWWNXNWN0xd:';c'     Entrez dans une nouvelle réalité
   .c,  :XMWWWWWKkKK0KKdck0Okdol. .c,
  .c'   'dxkKNWWO,...;kklxdOWWWXl  .c,       Soyez l'acteur du changement
  :;        .:ollc,cxOOooxkXMWO;.   'l.
 'c.         'dc.';ox0O,lWWWWKl''.   :;            Matérialisez vos rêves
 ;c         .;l:'..;coOxdKNXKc,kXd.  ;:
 ;c         .,:od:.   'kKOdl:'cdc.   ;:              Obtenez l'impossible
 'c.           .';c;',l0Oolc:c,.     :;
  :;              .xXN0c',;,'.      'c.                    Et plus encore
  .c,            'xKKd.            .c,
   .c,     .',;:lKMKxo:;,'..      .c,
    .::. 'xKNWWMMMWMMMMWWNXk:   .;c. .____                ___.
      .:;;,,:clloooooollc:;'. .;:,   |    |    ____   ____\_ |__ _____
        .,;;,.            .',;;'     |    |  _/ __ \ / ___\| __ \\__  \
           .',,,,,,,,,,,',,'.        |    |__\  ___// /_/  > \_\ \/ __ \_
                 .......             |_______ \___  >___  /|___  (____  /
          ____   ___                         \/   \/_____/     \/     \/
          \   \ /   /___   ____   __| _/____   ____   ____  ____   _____
           \   Y   /  _ \ /  _ \ / __ |/  _ \ /  _ \_/ ___\/  _ \ /     \
            \     (  <_> |  <_> ) /_/ (  <_> |  <_> )  \__(  <_> )  Y Y  \
             \___/ \____/ \____/\____ |\____/ \____/ \___  >____/|__|_|  /
                                     \/                  \/            \/
`
	satDesc = `
SATCOM, une division externalisée de
 ____                ___
|    |    ____   ____\_ |__ _____
|    |  _/ __ \ / ___\| __ \\__  \
|    |__\  ___// /_/  > \_\ \/ __ \_
|_______ \___  >___  /|___  (____  /
 ____   \/_  \/_____/     \/___  \/
 \   \ /   /___   ____   __| _/____   ____   ____  ____   _____
  \   Y   /  _ \ /  _ \ / __ |/  _ \ /  _ \_/ ___\/  _ \ /     \
   \     (  <_> |  <_> ) /_/ (  <_> |  <_> )  \__(  <_> )  Y Y  \
    \___/ \____/ \____/\____ |\____/ \____/ \___  >____/|__|_|  /
                            \/                  \/            \/

[Accès Restreint]         >>>>>>> entrez vos identifiants <<<<<<<
`
	arcDesc = `

*********************************************************************************
Legba Voodoocom ne peut être tenu responsable de l'usage et des données stockées.
**** WARNING **** : ce service n'est plus maintenu.
*********************************************************************************
━━━╮╭╮╱╱╱╱╱╱╱╱╱╱╱╭━━━╮╱╱╱╱╱╱╭╮
┃╭━╮┃┃┃╱╱╱╱╱╱╱╱╱╱╱┃╭━╮┃╱╱╱╱╱╭╯╰╮     Division: R&D, Unité 2772
┃╰━━┳┫┃╭┳━━┳━━┳━╮╱┃╰━━┳━━┳┳━╋╮╭╯               Projets spéciaux
╰━━╮┣┫┃┣┫╭━┫╭╮┃╭╮╮╰━━╮┃╭╮┣┫╭╋┫┃                (dir: A.M)
┃╰━╯┃┃╰┫┃╰━┫╰╯┃┃┃┃┃╰━╯┃╰╯┃┃┃┃┃╰╮
╰━━━┻┻━┻┻━━┻━━┻╯╰╯╰━━━┫╭━┻┻╯╰┻━╯
╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱┃┃
╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╱╰╯
*********************************************************************************
**** WARNING **** : ce service n'est plus maintenu.
Legba Voodoocom ne peut être tenu responsable de l'usage et des données stockées.
*********************************************************************************
[Beware MalvolentKiIA, hack@45EBG56#EACD M@dJ0k3r;3/4/206]
`
)

// serveur public de Legba Voodoocom
var legba = Server{
	Address:     "legba.d22.eu",
	Private:     false,
	Security:    SEC4,
	Description: lbDesc,
}

// serveur privé de la communication satellite
var legba_satcom = Server{
	Address:     "satcom.legba.d22.eu",
	Private:     true,
	Description: satDesc,
	Security:    SEC4,
}

// serveur archive de Silicon Spirit
var legba_archive = Server{
	Address:     "archive.legba.d22.eu",
	Private:     true,
	Description: arcDesc,
	Security:    SEC4,
}

var gdDesc = `
 _                _____                    _____   _        _          _         _
(_)      ____    (_____)         _        (_____) (_) ____ (_)_  _    (_)       (_)_
(_)     (____)   (_)__(_)  ___  (_)__     (_)  (_) _ (____)(___)(_)__  _    ___ (___)
(_)    (_)_(_)   (_____)  (___) (____)    (_)  (_)(_)(_)__ (_)  (____)(_) _(___)(_)
(_)____(__)__    (_)__(_)(_)_(_)(_) (_)   (_)__(_)(_) _(__)(_)_ (_)   (_)(_)___ (_)_
(______)(____)   (_____)  (___) (_) (_)   (_____) (_)(____) (__)(_)   (_) (____) (__)

   ** On vend de tout - s tout'achète..... **
`

// serveur le bon district
var lbd = Server{
	Address:     "lebondistrict.d22.eu",
	Private:     false,
	Description: gdDesc,
	Security:    SEC2,
}

// green data
var greendata = Server{
	Address:     "greendata.d22.eu",
	Private:     false,
	Description: greenDesc,
	Security:    SEC3,
}

var greenDesc = `
                             %      ____                        __  __
                         %%%%%%    /\  _ \                     /\ \/\ \
                %%%%%%%%%%%%%%%%   \ \ \L\_\  _ __    __     __\ \  \\ \
            %%%%%%%%%%%%%% %%%%%    \ \ \L_L /\  __\/ __ \ / __ \ \ ,   \
          %%%%%%%%%%%%% *%%%%%%%     \ \ \/, \ \ \//\  __//\  __/\ \ \ \ \
         %%%%%%%%%   %%%%%%%%%%       \ \____/\ \_\\ \____\ \____\\ \_\ \_\
         %%%%   %%%%%%%%%%%%%%         \/___/  \/_/ \/____/\/____/ \/_/\/_/
            .%%%%%%%%%%%%%%%
          %%%%%%%%%%%%%%%               ____              __    ______
        %%%                            /\  _ \           /\ \__/\  _  \
        %%                             \ \ \/\ \     __  \ \ ,_\ \ \L\ \
                                        \ \ \ \ \  / __ \ \ \ \/\ \  __ \
       Analyses / Diagnostics            \ \ \_\ \/\ \L\ \_\ \ \_\ \ \/\ \
                                          \ \____/\ \__/ \_\\ \__\\ \_\ \_\
       Certifié ISO-56-52-100              \/___/  \/__/\/_/ \/__/ \/_/\/_/
`

var invertedLeaf = `
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@.@@@@@
@@@@@@@@@@@@@@@@@@@@......@@@@
@@@@@@@@@@@................@@@
@@@@@@@..............@.....@@@
@@@@@.............@&.......@@@
@@@@.........@@@..........@@@@
@@@@....@@@..............@@@@@
@@@@@@@@...............@@@@@@@
@@@@@...............@@@@@@@@@@
@@@...@@@@@@@@@@@@@@@@@@@@@@@@
@@@..@@@@@@@@@@@@@@@@@@@@@@@@@
`

// serveur privé de Crunch
var leet = Server{
	Address:     "leet.darknet",
	Private:     true,
	Description: cruDesc,
	Security:    SEC3,
}

var cruDesc = `
         _/_/_/  _/_/_/    _/    _/  _/      _/    _/_/_/  _/    _/
      _/        _/    _/  _/    _/  _/_/    _/  _/        _/    _/
     _/        _/_/_/    _/    _/  _/  _/  _/  _/        _/_/_/_/
    _/        _/    _/  _/    _/  _/    _/_/  _/        _/    _/
     _/_/_/  _/    _/    _/_/    _/      _/    _/_/_/  _/    _/

    is NOT watching you... No need for that.... :o)
`

// serveur de Hope
var hopeServ = Server{
	Address:     "hope.local",
	Private:     true,
	Description: hopDesc,
	Security:    SEC3,
}

// Project "Hope"
// Dépot mémoriel
// - fenêtre temporelle glissante v12.5
// - compression McVaugh-Korba-Yang
// - contenu infix normalisé
// - (c) A.M
var hopDesc = `
01010000 01110010 01101111 01101010 01100101 01100011
01110100  00100010 01001000 01101111 01110000 01100101 00100010
01000100 11101001 01110000 01101111 01110100  01101101
11101001 01101101 01101111 01110010 01101001 01100101 01101100
00101101  01100110 01100101 01101110 11101010 01110100 01110010
01100101  01110100 01100101 01101101 01110000 01101111 01110010 01100101
01101100 01101100 01100101  01100111 01101100 01101001 01110011 01110011
01100001 01101110 01110100 01100101  01110110 00110001 00110010 00101110
00110101  00101101  01100011 01101111 01101101 01110000 01110010
01100101 01110011 01110011 01101001 01101111 01101110  01001101 01100011
01010110 01100001 01110101 01100111 01101000 00101101 01001011 01101111
01110010 01100010 01100001 00101101 01011001 01100001 01101110 01100111
00101101  01100011 01101111 01101110 01110100 01100101 01101110 01110101
01101001 01101110 01100110 01101001 01111000  01101110 01101111 01110010
01101101 01100001 01101100 01101001 01110011 11101001
00101101  00101000 01100011 00101001  01000001 00101110 01001101
`

// *****************************************************************************
// Server pour les Hacker, acyclic graph
// *****************************************************************************
// TODO FIXME allow some connected node to Nikki and Celine
// connected nodes can lead to d22
var connectedA = Server{
	Address:     "kommunishky.eu",
	Description: "Главный узел связи",
	Security:    SEC2,
}
var connectedB = Server{
	Address:     "kashik1842.eu",
	Description: "Вычислительный кластер - стойка 18 - банк 42",
	Security:    SEC3,
}
var connectedC = Server{
	Address:     "watchers.free.eu",
	Description: "Big Brother is watching them.....",
	Security:    SEC5,
}
var connectedD = Server{
	Address:     "kashik1851.eu",
	Description: "Вычислительный кластер - стойка 18 - банк 51",
	Security:    SEC3,
}
var connectedE = Server{
	Address:     "trace.net.eu",
	Description: "Backbone node for traffic surveillance. Keep out !",
	Security:    SEC4,
}
var connectedF = Server{
	Address:     "kashik1874.eu",
	Description: "Вычислительный кластер - стойка 18 - банк 74",
	Security:    SEC3,
}
var connectedG = Server{
	Address:     "pb5-34b.eu",
	Description: "Политбюро. Узел, занимающийся верификацией и валидацией. ФСБ, департамент 5, район 34Б.",
	Security:    SEC5,
}
var connectedH = Server{
	Address:     "route.global.ko",
	Description: "라우팅 인프라. 패킷 확인.",
	Security:    SEC1,
}
var connectedI = Server{
	Address:     "backup22.main.eu",
	Description: "Infrastructure de backup, district 22",
	Security:    SEC3,
}
var connectedJ = Server{
	Address:     "route.global.eu",
	Description: "Noeud de routage. Infrastructure de controle.",
	Security:    SEC1,
}
var connectedK = Server{
	Address:     "backup64.main.eu",
	Description: "Infrastructure de backup, district 64",
	Security:    SEC3,
}

// unconnected node have no toute to d22, and allow cycles :o)
var unconA = Server{
	Address:     "backup31.main.eu",
	Description: "Infrastructure de backup, district 31",
	Security:    SEC3,
}
var unconB = Server{
	Address:     "pb5-11b.eu",
	Description: "Политбюро. Узел, занимающийся верификацией и валидацией. ФСБ, департамент 5, район 11Б.",
	Security:    SEC5,
}
var unconC = Server{
	Address:     "kashik1953.eu",
	Description: "Вычислительный кластер - стойка 19 - банк 53",
	Security:    SEC3,
}
var unconD = Server{
	Address:     "cl52.cern.eu",
	Description: "Cluster de calcul, dédié muon-gamma. Réservation par OAR. oar-schedule@cern.eu.",
	Security:    SEC2,
}
var unconE = Server{
	Address:     "reservior12.eu",
	Description: "Резервное копирование и архивирование.",
	Security:    SEC2,
}
var unconF = Server{
	Address:     "backup02.main.eu",
	Description: "Infrastructure de backup, district 02",
	Security:    SEC3,
}
var unconG = Server{
	Address:     "cl53.cern.eu",
	Description: "Cluster de calcul, dédié muon-gamma. Réservation par OAR. oar-schedule@cern.eu.",
	Security:    SEC2,
}
var unconH = Server{
	Address:     "cl54.cern.eu",
	Description: "Cluster de calcul, dédié muon-gamma. Réservation par OAR. oar-schedule@cern.eu.",
	Security:    SEC2,
}

type Graph struct {
	Node     *Server
	LinkDesc string
	Links    []*Server
}

var allFLR = []*Identity{&crunch, &celine, &nikki}

var dag = []Graph{
	{&connectedA, "ссылка на ",
		[]*Server{&d22, &unconE}},
	{&connectedB, "ссылка на ",
		[]*Server{&unconG, &connectedA}},
	{&connectedC, "I got an eye on ",
		[]*Server{&d22, &connectedB, &unconH, &connectedD, &connectedE, &connectedH, &unconD}},
	{&connectedD, "ссылка на ",
		[]*Server{&unconB, &connectedA, &unconC, &connectedB, &connectedE}},
	{&connectedE, "Monitoring ",
		[]*Server{&d22, &unconA}},
	{&connectedF, "ссылка на ",
		[]*Server{&connectedA, &connectedB, &unconC, &connectedG}},
	{&connectedG, "ссылка на ",
		[]*Server{&connectedD, &unconB}},
	{&connectedH, "경로 ",
		[]*Server{&connectedF, &unconG, &connectedJ, &unconA}},
	{&connectedI, "Backup de ",
		[]*Server{&d22, &connectedK}},
	{&connectedJ, "Backup de ",
		[]*Server{&unconE, &connectedE, &connectedI}},
	{&connectedK, "Backup de ",
		[]*Server{&unconF, &connectedE}},
	{&unconA, "Backup de ",
		[]*Server{&unconB, &unconF}},
	{&unconB, "ссылка на ",
		[]*Server{&unconD}},
	{&unconC, "ссылка на ",
		[]*Server{&unconB, &unconG}},
	{&unconD, "Liens vers ",
		[]*Server{&unconC}},
	{&unconE, "ссылка на ",
		[]*Server{&unconA, &unconF}},
	{&unconF, "Backup de ",
		[]*Server{&unconA, &unconD}},
	{&unconG, "Liens vers ",
		[]*Server{&unconD, &unconH}},
	{&unconH, "Liens vers ",
		[]*Server{&unconF, &unconA, &unconG}},
}

// unconnected nodes can make cycle because not connected

// identités corpo recopiées depuis l'ancienne version
var (
	amathison     = Identity{"amathison", "hai3ja", "Alan Mathison", true}            // alan
	mmathison     = Identity{"mmathison", "mie6oo", "Mélody Mathison", true}          // mel
	mbellamy      = Identity{"mbellamy", "ahng7e", "Margherita Bellamy", true}        // rita
	sbronner      = Identity{"sbronner", "rahk0u", "Sebastian Bronner", true}         // styx
	cbellamy      = Identity{"cbellamy", "xoh7sh", "Camélia Bellamy", true}           // kapo
	jvillanova    = Identity{"jvillanova", "ay9aef", "Julius Villanova", true}        // scalpel
	ecanto        = Identity{"ecanto", "ti3eim", "Eddy Canto", true}                  // greko
	ejohannesen   = Identity{"ejohannesen", "obo4ie", "Edwin Johannesen", true}       // jesus
	jbranson      = Identity{"jbranson", "aich8g", "Jonathan Branson", true}          // escobar
	jmfright      = Identity{"jmfright", "uruw5g", "John Mac Fright", true}           // cageot
	skmihalec     = Identity{"skmihalec", "paeh3l", "Sylvia Kemija Mihalec", true}    // la fouine
	emartin       = Identity{"emartin", "thooy1", "Eva Martin", true}                 // eva
	mdubian       = Identity{"mdubian", "iup1ie", "Michael Dubian", true}             // fat mike
	cmihalec      = Identity{"cmihalec", "uequ8u", "Carlotta Mihalec", true}          // kennedy
	sjohannesen   = Identity{"sjohannesen", "aiphu4", "Sabrina Johannesen", true}     // savage girl
	rmichu        = Identity{"rmichu", "ool7ch", "Raoul Michu", true}                 // raoul cool
	rglass        = Identity{"rglass", "ahzae2", "Rupert Glass", true}                // green glass
	sglass        = Identity{"sglass", "si6aeb", "Stefie Glass", true}                // stefie
	djohannesen   = Identity{"djohannesen", "loh1ie", "Daisy Johannesen", true}       // chilly daisy
	dbonenfant    = Identity{"dbonenfant", "de4oiv", "Désiré Bonenfant", true}        // frère ping
	hproskychev   = Identity{"hproskychev", "ooj4an", "Harald Proskychev", true}      // papa proxy
	njasinski     = Identity{"njasinski", "eveth3", "Nikole Jasinski", true}          // nikki
	sjasinski     = Identity{"sjasinski", "ie7uo2", "Stefan Jasinski", true}          // sasquatch
	ffceline      = Identity{"ffceline", "boh6ay", "Franz-Ferdinand Celine", true}    // celine
	cmills        = Identity{"cmills", "thue1d", "Camélia Mills", true}               // cramille
	lseptembre    = Identity{"lseptembre", "cul1ol", "Lilas Septembre", true}         // tiger doll
	edubian       = Identity{"edubian", "rooch7", "Eloïse Dubian", true}              // sister morphine
	zabasolo      = Identity{"zabasolo", "aipho0", "Zilmir Abasolo", true}            // zilmir
	ebranson      = Identity{"ebranson", "rae2ie", "Elisabeth Branson", true}         // betty b
	jkievain      = Identity{"jkievain", "nie3oo", "Jordan Kievain", true}            // abraham
	fmanson       = Identity{"fmanson", "tiuf0y", "Frédéric Manson", true}            // crunch
	rkievain      = Identity{"rkievain", "aso2qu", "Rodolph Kievain", true}           // one kick
	pdoberty      = Identity{"pdoberty", "aivei1", "Pete Doberty", true}              // jacob
	rwhite        = Identity{"rwhite", "ies2su", "Richard White", true}               // oggy
	ajolivet      = Identity{"ajolivet", "quai1a", "Adrien Jolivet", true}            // cyrano
	mklebert      = Identity{"mklebert", "eis6ku", "Mickael Klebert", true}           // iron mike
	jvazzanna     = Identity{"jvazzanna", "ueth4k", "Joseph Vazzanna", true}          // small joe
	jbatista      = Identity{"jbatista", "yah6ae", "Johaquim Batista", true}          // joe-rez
	gsuleymanoglu = Identity{"gsuleymanoglu", "zo1daa", "Georges Suleymanoglu", true} // georges

	// PNJs
	afrieman  = Identity{"afrieman", "far3ik", "Anton Frieman", true} // PNJ fan blackwave
	ifrancher = Identity{"ifrancher", "asu62k", "Isabella Francher", true}
	yfrancher = Identity{"yfrancher", "oegy8s", "Yves Francher", true}
	svox      = Identity{"svox", "eg76wn", "Samantha Vox", true}
	// TODO quelques employé•e•s de la kramps
	akremmer  = Identity{"akremmer", "sexgod22", "Alexandre Kremmer", true}   // security Kramps
	mdavidson = Identity{"mdavidson", "allbitches", "Milton Davidson", true}  // dir adjoint Kramps
	vredmint  = Identity{"vredmint", "lily-dorian", "Virginia Redmint", true} // assistante Kramps
	taugusto  = Identity{"taugusto", "tde54e", "Terry Augusto", true}
	// TODO quelques employé•e•s de legba voodoocom
	atrebinsky = Identity{"atrebinsky", "56raz8", "Anthon Trebinsky", true}  // proj. Mandrake
	dyuong     = Identity{"dyuong", "gd86rw", "Dyop Yuong", true}            // proj. Mandrake
	eherswing  = Identity{"eherswing", "oh7fd4", "Emmet Herswing", true}     // proj. Mandrake
	jkuipers   = Identity{"jkuipers", "azgh4d", "Jordan Kuipers", true}      // proj. Mandrake
	jmfusion   = Identity{"jmfusion", "sg7vf4", "John-Mickael Fusion", true} // Manager LegbaV
	yblansein  = Identity{"yblansein", "tyg45g", "Youri Blansein", true}     // satcom
	// Employé de GreenData
	cyolinaro = Identity{"cyolinaro", "rtd98y", "Consuella Yolinaro", true} // gère contrats chez Green Data
	// Cours de justice
	agargan = Identity{"agargan", "tdg5df", "Armand Gargan", true}
	// quelques boites mails
	contKitsu  = Identity{"cont4yes@kitsune", "bosskitsu", "Contact Famille Kitsune", true}
	contMills  = Identity{"mills.contact@weyland.eu", "bossweyland", "Contact Weyland", true}
	contKramps = Identity{"contact@kramps.d22.eu", "bosskramps", "Contact Kramps Security", true}
)

// identités virtuelles fournies par Jésus et le FLR
var (
	hope           = Identity{"hope", "011011011011", "Hope", false}
	mel            = Identity{"mel", "hope4ever", "Mel", false}
	rocky          = Identity{"rocky", "pourquoi", "Rocky", false}
	rita           = Identity{"rita", "1wantM0re", "Rita", false}
	styx           = Identity{"styx", "plastobeton", "Styx", false}
	kapo           = Identity{"kapo", "touspour1", "Kapo", false}
	scalpel        = Identity{"scalpel", "m3dic!!", "Scalpel", false}
	greko          = Identity{"greko", "FuckY00", "Greko", false}
	jesus          = Identity{"jesus", "ZtqCtdlb42", "Jesus", false}
	escobar        = Identity{"escobar", "M0n3y++", "Escobar", false}
	cageot         = Identity{"cageot", "fr33dom", "Cageot", false}
	lafouine       = Identity{"lafouine", "cplvfh3h3", "La Fouine", false}
	eva            = Identity{"eva", "n3verAgain", "Eva", false}
	fatmike        = Identity{"fatmike", "tamereenshort", "Fat", false}
	kennedy        = Identity{"kennedy", "jensaisrien", "Kennedy", false}
	savagegirl     = Identity{"savagegirl", "zeStyle!", "Savage", false}
	raoulcool      = Identity{"raoulcool", "barb3rKing", "Raoul Cool", false}
	greenglass     = Identity{"greenglass", "il0veM0m", "Green Glass", false}
	chillydaisy    = Identity{"chillydaisy", "rb0nesQueen", "Chilly Daisy", false}
	ping           = Identity{"ping", "n0tp0ng!!", "Frère Ping", false}
	papaproxy      = Identity{"papaproxy", "4ragnar!", "Papa Proxy", false}
	nikki          = Identity{"nikki", "3141593", "Nikki", false}
	celine         = Identity{"celine", "f0rg3tme", "Céline", false}
	cramille       = Identity{"cramille", "onsenbalance", "Cramille", false}
	tigerdoll      = Identity{"tigerdoll", "karateGirl", "Tiger Doll", false}
	sistermorphine = Identity{"sistermorphine", "Icanfly", "Sister Morphine", false}
	zilmir         = Identity{"zilmir", "cl3v3r", "Zilmir", false}
	bettyb         = Identity{"bettyb", "ZeK0nsol", "Betty B", false}
	abraham        = Identity{"abraham", "killerSolo", "Abraham", false}
	crunch         = Identity{"crunch", "umdpcfpnp3o", "Crunch", false}
	onekick        = Identity{"onekick", "faitchier", "One Kick", false}
	jacob          = Identity{"jacob", "el01se", "Jacob", false}
	oggy           = Identity{"oggy", "y0pasC0ul", "Oggy", false}
	cyrano         = Identity{"cyrano", "rbNbOne", "Cyrano", false}
	ironmike       = Identity{"ironmike", "deadlymike", "Iron Mike", false}
	smallbob       = Identity{"smallbob", "smallbob", "Small Bob", false}
	joerez         = Identity{"joerez", "mfuck3r", "Joe-Rez", false}
	jeanne         = Identity{"jeanne", "j", "Jeanne", false}
	paula          = Identity{"paula", "mdpH@rd", "Paula", false}
	georges        = Identity{"georges", "devine!", "Georges", false}
	ringo          = Identity{"ringo", "l@cherien!", "Ringo", false}
	schwartz       = Identity{"schwartz", "noircnoir", "Schwartz", false}
)

type InfoPlayer struct {
	Perso  string
	IdCorp *Identity
	IdVirt Identity
	Wanted bool
	Known  bool
}

var allPlayers = []InfoPlayer{
	{"Hope", nil, hope, false, false},
	{"Mel", &mmathison, mel, false, true},
	{"Rocky", nil, rocky, false, false},
	{"Rita", &mbellamy, rita, false, true},
	{"Styx", &sbronner, styx, false, true},
	{"Kapo", &cbellamy, kapo, false, true},
	{"Scalpel", &jvillanova, scalpel, false, true},
	{"Greko", &ecanto, greko, false, true},
	{"jesus", &ejohannesen, jesus, false, true},
	{"Escobar", &jbranson, escobar, false, true},
	{"Cageot", &jmfright, cageot, false, true},
	{"La Fouine", &skmihalec, lafouine, false, true},
	{"Eva", &emartin, eva, false, true},
	{"Fat Mike", &mdubian, fatmike, false, true},
	{"Kenndy", &cmihalec, kennedy, false, true},
	{"Savage Girl", &sjohannesen, savagegirl, false, true},
	{"Raoul Cool", &rmichu, raoulcool, false, true},
	{"Green Glass", &rglass, greenglass, false, true},
	{"Chilly Daisy", &djohannesen, chillydaisy, false, true},
	{"Frère Ping", &dbonenfant, ping, false, true},
	{"Papa Proxy", &hproskychev, papaproxy, false, true},
	{"Nikki", &njasinski, nikki, false, true},
	{"Céline", &ffceline, celine, false, true},
	{"Cramille", &cmills, cramille, false, true},
	{"Tiger Doll", &lseptembre, tigerdoll, false, true},
	{"Sister Morphine", &edubian, sistermorphine, false, true},
	{"Zilmir", &zabasolo, zilmir, false, true},
	{"Betty B", &ebranson, bettyb, false, true},
	{"Abraham", &jkievain, abraham, false, true},
	{"Crunch", &fmanson, crunch, false, true},
	{"One Kick", &rkievain, onekick, false, true},
	{"Jacob", &pdoberty, jacob, false, true},
	{"Oggy", &rwhite, oggy, false, true},
	{"Iron Mike", &mklebert, ironmike, false, true},
	{"Joe-Rez", &jbatista, joerez, false, true},
	{"Cyrano", &ajolivet, cyrano, false, true},
	{"Small Bob", &jvazzanna, smallbob, false, true},
	{"Jeanne", nil, jeanne, false, true},
	{"Ringo", nil, ringo, false, true},
	{"Georges", &gsuleymanoglu, georges, false, false},
	{"Paula", nil, paula, false, true},
	{schwartz.Name, nil, schwartz, false, true},
}
var otherIds = []Identity{
	// fan blackwave
	afrieman, ifrancher, yfrancher, svox,
	// sasquatch
	sjasinski,
	// Kramps
	akremmer, mdavidson, vredmint,
	// Mandrake, Legba, SatCom
	amathison, atrebinsky, dyuong, eherswing, jkuipers, jmfusion, yblansein,
	// Green Data
	cyolinaro,
	// cours justice
	agargan,
}

// All the IDs, Corpo and Virtuelles
func AllIds() []Identity {
	var allIds []Identity

	for _, perso := range allPlayers {
		if perso.IdCorp != nil {
			allIds = append(allIds, *perso.IdCorp)
		}
		allIds = append(allIds, perso.IdVirt)
	}
	allIds = append(allIds, otherIds...)

	return allIds
}

// Add all the otherIds (group=public, backdoor=false) that are not 'alreadyOnServer'
// PAS SUR QUE UTILE
// func AllOtherUser( otherIds []Identity, alreadyOnServer []User) []User {
// 	var allUsers []User

// 	for _, id := range otherIds {
// 		// Check not alreadyIn
// 		already := false
// 		for _, u := range alreadyOnServer {
// 			if u.Login == id.Login {
// 				already = true
// 				break
// 			}
// 		}
// 		if !already {
// 			user := User{
// 				Login: id.Login,
// 				Server: "",
// 				Backdoor: false,
// 			}
// 			allUsers = append(allUsers, user)
// 		}
// 	}
// 	return allUsers
// }

func Init() {
	log.Println("réinitialisation de la base de données")

	Reset()

	InitNetwork(
		AllIds(),
		[]Transaction{
			{
				To:      "jesus",
				Yes:     100,
				Comment: "report du solde",
			},
			{
				To:      "admin",
				Yes:     1337,
				Comment: "report du solde",
			},
		},
	)

	InitServer(dd,
		[]User{
			{
				Login:    "jesus",
				Server:   "",
				Backdoor: false,
				Groups:   []string{"elec", "admin"},
			},
			{
				Login:    "ping",
				Server:   "",
				Backdoor: false,
				Groups:   []string{"elec"},
			},
			{
				Login:    "nikki",
				Server:   "",
				Backdoor: false,
				Groups:   []string{"elec"},
			},
			{
				Login:    "schwartz",
				Server:   "",
				Backdoor: false,
				Groups:   []string{"elec"},
			},
		},
		[]Link{
			{Address: d22.Address, Description: "serveur public du District 22"},
			{Address: frozdd.Address, Description: "pour les fans de Frozpunk"},
			{Address: maravdd.Address, Description: "paris en ligne pour le Saturday Open Fight"},
		},
		[]Register{
			// FIXME TODO registres elec
			{Description: "machine à café", State: "on", Options: []RegisterState{"on", "off", "overdrive"}},
		},
		[]Post{},
	)
	InitServer(frozdd,
		[]User{},
		[]Link{},
		[]Register{},
		[]Post{},
	)
	InitServer(maravdd,
		[]User{},
		[]Link{},
		[]Register{},
		[]Post{},
	)

	InitServer(d22,
		[]User{},
		[]Link{
			{Address: legba.Address, Description: "Legba Voodoocom"},
			{Address: kramps.Address, Description: "Kramps Security"},
			{Address: corp.Address, Description: "Central Services"},
			{Address: abus.Address, Description: "Association des Banques Unifiées Suisses"},
			{Address: greendata.Address, Description: "Green Data, solution environementale"},
			{Address: lbd.Address, Description: "Le Bon District, affaires à faire !"},
		},
		[]Register{},
		[]Post{},
	)
	InitServer(kramps,
		[]User{
			{Login: "akremmer", Groups: []string{"pers", "sec"}},
			{Login: "mdavidson", Groups: []string{"pers", "sec", "diradj"}},
		},
		[]Link{
			{Address: persKramps.Address, Description: "Serveur réservé au personnel"},
			{Address: elecKramps.Address, Description: "Gestion de l'énergie"},
		},
		[]Register{},
		[]Post{},
	)
	InitServer(persKramps,
		[]User{
			{Login: "akremmer", Groups: []string{"pers", "sec"}},
			{Login: "mdavidson", Groups: []string{"pers", "sec", "diradj"}},
			{Login: "rkievain", Groups: []string{"pers"}},
			{Login: "vredmint", Groups: []string{"pers"}},
			{Login: "taugusto", Groups: []string{"pers", "elec"}},
		},
		[]Link{
			//{Address: kramps.Address, Description: "Accueil"},
			{Address: secKramps.Address, Description: "Sécurité des installations", Group: "sec"},
		},
		[]Register{},
		[]Post{},
	)
	InitServer(secKramps,
		[]User{
			{Login: "akremmer", Groups: []string{"pers", "sec"}},
			{Login: "mdavidson", Groups: []string{"pers", "sec", "diradj"}},
		},
		[]Link{
			//{Address: kramps.Address, Description: "Accueil"},
		},
		[]Register{},
		[]Post{},
	)
	InitServer(elecKramps,
		[]User{
			{Login: "akremmer", Groups: []string{"pers", "sec"}},
			{Login: "mdavidson", Groups: []string{"pers", "sec", "diradj", "elec"}},
			{Login: "taugusto", Groups: []string{"pers", "elec"}},
		},
		[]Link{},
		[]Register{},
		[]Post{},
	)
	InitServer(corp,
		[]User{},
		[]Link{
			{Address: justice.Address, Description: "Services judiciairesé"},
		},
		[]Register{},
		[]Post{},
	)
	InitServer(justice,
		[]User{
			{Login: "agargan", Groups: []string{"allowed"}},
		},
		[]Link{},
		[]Register{},
		[]Post{},
	)
	InitServer(abus,
		[]User{},
		[]Link{},
		[]Register{},
		[]Post{},
	)

	InitServer(legba,
		[]User{
			{Login: "jkuipers", Groups: []string{"pers", "admin"}},
			{Login: "jmfusion", Groups: []string{"pers", "admin"}},
		},
		[]Link{
			{Address: legba_satcom.Address, Description: "Division Sat-Com"},
			{Address: legba_archive.Address, Description: "Archives", Group: "admin"},
		},
		[]Register{},
		[]Post{},
	)
	InitServer(legba_satcom,
		[]User{
			{Login: "yblansein", Groups: []string{"sat"}},
		},
		[]Link{},
		[]Register{},
		[]Post{},
	)
	InitServer(legba_archive,
		[]User{
			{Login: "jkuipers", Groups: []string{"pers", "admin"}},
			{Login: "jmfusion", Groups: []string{"pers", "admin"}},
			{Login: "atrebinsky", Groups: []string{"pers"}},
			{Login: "dyuong", Groups: []string{"pers"}},
			{Login: "eherswing", Groups: []string{"pers"}},
		},
		[]Link{},
		[]Register{},
		[]Post{},
	)

	InitServer(lbd,
		[]User{},
		[]Link{},
		[]Register{},
		[]Post{},
	)

	InitServer(greendata,
		[]User{
			{
				Login:    "cyolinaro",
				Server:   "",
				Backdoor: false,
				Groups:   []string{"contract"},
			},
		},
		[]Link{},
		[]Register{},
		[]Post{},
	)

	InitServer(leet,
		[]User{
			{Login: "crunch", Backdoor: false, Groups: []string{"admin", "flr", "crunch"}},
			{Login: "celine", Backdoor: false, Groups: []string{"celine", "flr", "pornloverz"}},
			{Login: "nikki", Backdoor: false, Groups: []string{"flr", "nikki"}},
			{Login: "greenglass", Backdoor: false, Groups: []string{"pornloverz", "gg"}},
			{Login: "bettyb", Backdoor: false, Groups: []string{"bettyb"}},
		},
		[]Link{
			{
				Group:       "crunch",
				Address:     connectedJ.Address,
				Description: "réseau redirections",
			},
			{
				Group:       "celine",
				Address:     connectedI.Address,
				Description: "réseau redirections",
			},
			{
				Group:       "nikki",
				Address:     connectedD.Address,
				Description: "réseau redirections",
			},
		},
		[]Register{},
		[]Post{},
	)
	InitServer(hopeServ,
		[]User{
			// TODO FIXME ajouter amathison
			{Login: "hope", Backdoor: false, Groups: []string{"admin"}},
			{Login: "amathison", Backdoor: false},
		},
		[]Link{},
		[]Register{},
		[]Post{},
	)

	// acyclic graph for hackers
	for _, g := range dag {
		var links []Link
		for _, l := range g.Links {
			links = append(links, Link{
				Address:     l.Address,
				Description: fmt.Sprintf("%s %s", g.LinkDesc, l.Address),
			})
		}

		InitServer(*(g.Node), []User{}, links, []Register{}, []Post{})
	}

	// chargement YAML
	LoadTransactions("contenu/trans_balance.yaml")
	LoadPosts("contenu/for_lbd.yaml")
	LoadPosts("contenu/for_gg_grocery_contract.yaml")
	LoadPosts("contenu/for_justice.yaml")
	LoadTransactions("contenu/trans_gg.yaml")
	LoadPosts("contenu/for_hope_greendata_final.yaml")
	LoadPosts("contenu/for_legba_final.yaml")
	LoadPosts("contenu/for_kramps_final.yaml")
	LoadPosts("contenu/for_frozpunk_BY_HAND.yaml")
	LoadPosts("contenu/for_leet.yaml")
	LoadPosts("contenu/for_nrj.yaml")
	LoadPosts("contenu/for_ddlocal.yaml")
	LoadRegistries("contenu/reg_kramps.yaml")
	LoadRegistries("contenu/reg_satcom.yaml")
	LoadRegistries("contenu/reg_nrj.yaml")
	LoadRegistries("contenu/reg_ddlocal.yaml")
	LoadMessages("contenu/msg_francher.yaml")
}
