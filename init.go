package main

const (
	SEC1 = 0.05
	SEC2 = 0.1
	SEC3 = 0.2
	SEC4 = 0.4
	SEC5 = 0.8
)

const (
	mel            = "GGCCAAAGCTCCTTCGGAGC"
	rocky          = "CCGCGCAGAATCATAGCTGT"
	rita           = "CAAAGTTCTAGGCATAGGGA"
	styx           = "TTAGCTCGATATCCTAACCC"
	kapo           = "GAACTGCTTTAGTTGACGGA"
	scalpel        = "TGAAAGAGACATGATGCCTT"
	greko          = "TCTGAGGTTTATTGATTTCG"
	jesus          = "TTCGGGATTACTGCGTGCTG"
	escobar        = "GGAGGACACCCCAAACGCAT"
	cageot         = "GCCCTTGTCATGTACTTAGT"
	lafouine       = "CTGTCACCCAATCTACAGCG"
	eva            = "CTGTTGTAGTGACATGTTTC"
	fatmike        = "AACCTTGGGCACGGTCGGTA"
	kennedy        = "CCCGCGGGCAAAGCTGACAG"
	savagegirl     = "GGGTCTATAGGTCAAACGGT"
	raoulcool      = "GTCACAAGGTTGTTTAATGG"
	greenglass     = "ATGCCTACCTCCAATGATTA"
	chillydaisy    = "CGGGAGACACGTTCAGTCTT"
	frereping      = "GCATGGCCGAATTCCTCATT"
	papaproxy      = "CGATTTGTATTGGATACGGA"
	nikki          = "ACGAACCTAGAGCCGCACGC"
	celine         = "CGCTCCCATTTCATGTCAGC"
	cramille       = "TTTGGGAGAAGCTTATGCAC"
	tigerdoll      = "ATATGTTGAGCGTAAAGGCG"
	sistermorphine = "CCATCCGGCGGACCTTATGC"
	zilmir         = "GACGGGATACCTACTCTCGA"
	bettyb         = "ATTCCGACTCAGGGTACCGG"
	abraham        = "TGGCGTCTCTAATTCTTGCC"
	crunch         = "TTCAAGCTGAATATGAAAGG"
	onekick        = "GTCAAATCTGAGACTCTTGC"
	jacob          = "TGAAAGAGACAGTATGCCGT"
	gang1          = "TTCGACTGAATGTTTGATGT"
	gang2          = "TATCGACGCACGGGACTTGG"
	gang3          = "CGAGAAATGACAGAGTTGTA"
	paula          = "GGGTGATCTGTTGCCCCCTG"
	ringo          = "AACTGACGGATTCGATCATG"
	georges        = "GTTTGCACGGAACATGCAAC"
	jeanne         = "GACCCGTATTTCGCTGATTG"
	oggy           = "TCAGCTTCTAACGTTCGGGA"
)

// serveur local du dirty district
var dd = Server{
	Address: "dd.local",
	Credentials: []Cred{
		{"public", "public", 1},
		{"jesus", "roxor", 5},
	},
	Description: ddDesc,
	Targets: []Target{
		{d22.Address, "serveur public du District 22", 1, "public", "public"},
	},
	Registers: []Register{
		{"cafe", false, "machine à café", 1},
		{"sono", true, "et je coupe le son ...", 3},
	},
	Entries: []Entry{
		{"bluemars", []string{"boisson"}, 1, "jesus", "Blue Mars - le cocktail parfait", "la recette"},
		{"mojito", []string{"boisson"}, 1, "", "Mojito - le cocktail classique", "Menthe, glace pilée, citron vert et plein de rhum"},
	},
	Detection: SEC1,
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

Tape "index" pour avoir la liste des services fournis par le serveur. Si tu as
besoin d'aide, demande à ton nerd préféré.
`

// serveur public du district 22
var d22 = Server{
	Address: "d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: dd22Desc,
	Targets: []Target{
		{legba.Address, "Legba Voodoocom", 1, "public", "public"},
		{kramps.Address, "Kramps Security", 1, "public", "public"},
		{corp.Address, "Central Services", 1, "public", "public"},
		{abus.Address, "Association des Banques Unifiées Suisses", 1, "public", "public"},
		{greendata.Address, "Green Data, solution environnementale", 1, "public", "public"},
	},
	Detection: SEC2,
}
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
           Un noeud du plus grand fournisseur d'accès de Méga-Europe. 
                                                                                                
`

// serveur public de la kramps
var kramps = Server{
	Address: "kramps.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
		{"joe", "password", 3}, // utilisateur lambda, accès direct
	},
	Targets: []Target{
		{kramps_priv.Address, "Serveur réservé au personnel", 3, "personnel", "kramps1234"},
	},
	Description: kpubDesc,
	Detection: SEC2,
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
   _|_  | | _   |\/| _  _  _| _    _ |    _  (~    _  _|_       +-+-+-+-+-+-+-+-+
    !   |_|| |  |  |(_)| |(_|(/_  |_)||_|_\  _)|_||    !        | | | | | | | | |
                                  |                             | | | | | | | | |
                                                                | | | | | | | | |
                                                                | | | | | | | | |
                                                                | | | | | | | | |
                    Plus de 4.000.000.000 ¥€$ de capital        | | | | | | | | |
`

// serveur privé de la kramps
var kramps_priv = Server{
	Address: "priv.kramps.d22.eu",
	Credentials: []Cred{
		{"personnel", "kramps1234", 3},
	},
	Targets: []Target{
		{kramps_sec.Address, "Serveur central de sécurité", 5, "admin", "lkjqsod"},
	},
	Detection: SEC3,
	Description: kperDesc,
}

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

  18/07: StuKa est la 3° organisation du personnel (spécialisée Tech niv 1-3) à 
         déposer le bilan cette année.

  02/07: Rappel ! Les dossiers de promotions pour SecAgent, tout niveau, sont à
         renvoyer avant le 31/07. Tarif habituel, voir agence comptable.
`

// serveur de sécurité de la kramps
var kramps_sec = Server{
	Address: "sec.kramps.d22.eu",
	Credentials: []Cred{
		{"admin", "lkjqsod", 5},
	},
	Detection: SEC4,
	Description: ksecDesc,
}

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
// serveur des services corporatistes D22
var corp = Server{
	Address: "corp.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: cd22Desc,
	Targets: []Target{
		{justice.Address, "services judiciaires", 1, "public", "public"},
	},
}

var cd22Desc = `

   ((ervices  ((orporatistes
   ''                       
    _                       
   [|)istrict  22           

   Ce service **public** vous est proposé **gratuitement** par la Cours Corporatiste.

   Ce service est livré en l'état, et la Cours Corporatiste décline toute responsabilité
   en ce qui concerne les données présentes et l'usage qui en est fait.

   Ce site existe gràce à la généreuse participation de Weyland-Yutani Corp,
   Tyrel Corp, Tessier-Ashpool SA, Disney Dassault, Arasaka, Renraku, Ubik,
   Legba Voodoocom, Avalon, Association des Banques Unifiées Suisses (ABUS).

                      
`

// serveur judiciaire
var justice = Server {
	Address: "justice.corp.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: cd22justDesc,	
	Entries: []Entry{
		{"@mel", []string{mel}, 1, "", "Mélody MATHISON", "Disparue - Incident 16485-4346B, Nexkemia Petrochemicals, 07/07/2000"},
		{"@rocky", []string{rocky}, 1, "", "TODO", "- D22/de#867533654: encours de dettes, cumul 4.463 ¥€$\n- D22/ou#7578538765: outrage et rébellion, EuroPol\n- D22/ou#65432446543: outrage et rébellion, LegbaSecurity" },
		{"@rita", []string{rita}, 1, "", "Margherita BELLAMY", "- néant"},
		{"@styx", []string{styx}, 1, "", "Sébastien BRONNER", "TODO"},
		{"@kapo", []string{kapo}, 1, "", "Carmélia BELLAMY", "TODO"},
		{"@scalpel", []string{scalpel}, 1, "", "Julius VILLANOVA", "***** Personne recherchée, mandat inter-district PJ/676/ER/65534 *****\n- D22/cm#5674243: complicité de meurtre"},
		{"@greko", []string{greko}, 1, "", "Eddy TODO", "- néant"},
		{"@jesus", []string{jesus}, 1, "", "Edwin JOHANNESEN", "- néant"},
		{"@escobar", []string{escobar}, 1, "", "Jonathan BRANSON", "- néant"},
		{"@cageot", []string{cageot}, 1, "", "John MacFRIGHT", "***** Personne recherchée, mandat inter-district PF/0865/EP/55463 *****\n- D21/rc#12785234452 rupture contrat\n\n\n$$$SPECIAL$$$ contacter cont4yes@kitsu.eu, ¥€$ en rapport."},	
		{"@lafouine", []string{lafouine}, 1, "", "Sylvia Kemija MIHALEC", "- néant"},
		{"@eva", []string{eva}, 1, "", "Pamela TODO", "***** Personne recherchée, mandat inter-district PF/1437/PM/02 *****\n- D21/rc#6542867 rupture contrat"},
		{"@fatmike", []string{fatmike}, 1, "", "Michael DUBIAN", "- D22/vm#23842834: vol à l'étalage\n- D22/vm#54327653: vol recette épicerie nuit\n- D22/vm#543299873: vol simple\n- D22/vm#547699823: vol graviscooter\n- D22/vm#753296671: vol à l'étalage"},
		{"@kennedy", []string{kennedy}, 1, "", "Carlotta MIHALEC", "***** Personne recherchée, mandat inter-district PF/0865/EP/55463 *****\n- D22/vd#765428736: vol données confidentielles "},
		{"@savagegirl", []string{savagegirl}, 1, "", "Sabrina JOHANNESEN", "- néant TODO"},
		{"@raoulcool", []string{raoulcool}, 1, "", "Raoul MICHU", "- néant TODO"},
		{"@greenglass", []string{greenglass}, 1, "", "Rupert GLASS", "- néant"},
		{"@chillydaisy", []string{chillydaisy}, 1, "", "Daisy JOHANNESEN", "TODO"},
		{"@frereping", []string{frereping}, 1, "", "Désiré BONENFANT", "- néant"},
		{"@papaproxy", []string{papaproxy}, 1, "", "Harald PROSKYCHEV", "***** Personne recherchée, mandat inter-district PF/2964/EP/98254 *****\n- D22/vd#89875357678: vol données avec copyright"},
		{"@nikki", []string{nikki}, 1, "", "Nicole JASINSKI", "***** Personne recherchée, mandat inter-district PF/7253/EP/90271 *****\n- D22/vd#1100298735: vol données sous brevet"},
		{"@celine", []string{celine}, 1, "", "Franz-Ferdinand CÉLINE", "***** Personne recherchée, mandat inter-district PF/1001/EP/98682 *****\n- D22/pi#9867356873: piratage informatique\n- D22/am#18763725: association malfaiteurs"},
		{"@cramille", []string{cramille}, 1, "", "Camelia MILLS", "- néant"},
		{"@tigerdoll", []string{tigerdoll}, 1, "", "Mei-Li Lilas TODO", "- néant"},
		{"@sistermorphine", []string{sistermorphine}, 1, "", "Eloïse DUBIAN", "- D22/ou#7578538765: outrage et rébellion, EuroPol\n- D22/va#325363552: vandalisme\n- D22/td#89765363: tapage diurne répété\n- D22/tn#101002543: tapage nocturne"},
		{"@zilmir", []string{zilmir}, 1, "", "Zilmir Abasolo", "- néant"},
		{"@bettyb", []string{bettyb}, 1, "", "Elisabeth BRANSON", "- néant"},
		{"@abraham", []string{abraham}, 1, "", "TODO", "- néant"},
		{"@crunch", []string{crunch}, 1, "", "TODO", "- néant"},
		{"@onekick", []string{onekick}, 1, "", "Rodolphe KIÉVAIN", "- néant\n>>> automated procedure: contact@kramps.eu | #line>2"},
		{"@jacob", []string{jacob}, 1, "", "Pete TODO", "- néant"},
		{"@cyrano", []string{gang1}, 1, "", "Adrien JOLIVET", "TODO"},
		{"@smalljoe", []string{gang2}, 1, "", "Mickael KLEBERT", "TODO"},
		{"@ironmike", []string{gang3}, 1, "", "Joseph VAZZANNA", "TODO"},
		{"@paula", []string{paula}, 1, "", "Paula TODO", "TODO NON"},
		{"@ringo", []string{ringo}, 1, "", "Ringo TODO", "TODO NON"},
		{"@georges", []string{georges}, 1, "", "Georges TODO", "TODO NON"},
		{"@jeanne", []string{jeanne}, 1, "", "Jeanne TODO", "TODO NON"},
		{"@joggy", []string{oggy}, 1, "", "Richard WHITE", "- néant"},
	},
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
	Address: "abus.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
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

// serveur public de Legba Voodoocom
var legba = Server{
	Address: "legba.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: lbDesc,
	Targets: []Target{
		{legba_satcom.Address, "division sat-com", 5, "admin", "satcom9876"},
		{legba_archive.Address, "archives", 3, "personnel", "archive6543"},
	},
}
var lbDesc = `
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

// serveur privé de la communication satellite
var legba_satcom = Server{
	Address: "satcom.legba.d22.eu",
	Credentials: []Cred{
		{"admin", "satcom9876", 5},
	},
	Description: satDesc,
}
var satDesc = `
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

// serveur archive de Silicon Spirit
var legba_archive = Server{
	Address: "archive.legba.d22.eu",
	Credentials: []Cred{
		{"personnel", "archive6543", 3},
	},
	Description: arcDesc,
}
var arcDesc = `
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

// serveur le bon district
var lbd = Server{
	Address: "lebondistrict.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
}

// green data
var greendata = Server{
	Address: "greendata.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
	Description: greenDesc,
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
	Address: "l33t.darknet",	       
	Credentials: []Cred{		       
		{"crunch", "hacktheplanet", 5},
	},
	Description: cruDesc,
}                                        


var cruDesc = `

                                                                      
         _/_/_/  _/_/_/    _/    _/  _/      _/    _/_/_/  _/    _/   
      _/        _/    _/  _/    _/  _/_/    _/  _/        _/    _/    
     _/        _/_/_/    _/    _/  _/  _/  _/  _/        _/_/_/_/     
    _/        _/    _/  _/    _/  _/    _/_/  _/        _/    _/      
     _/_/_/  _/    _/    _/_/    _/      _/    _/_/_/  _/    _/       
                                                                      
    is NOT watching you... No need for that.... :o) 


`

// serveur privé de Céline
var lair = Server{
	Address: "celine.darknet",
	Credentials: []Cred{
		{"celine", "waytoocool", 5},
	},
	Description: celDesc,
}
var celDesc = `
  ******               *******                    **    
  **////**   **    **  /**////**                  /**    
 **    //  ************/**    /**  ******   ******/**  **
/**       ///**////**/ /**    /** //////** //**//*/** ** 
/**         /**   /**  /**    /**  *******  /** / /****  
//**    ** ************/**    **  **////**  /**   /**/** 
 //****** ///**////**/ /*******  //********/***   /**//**
  //////    //    //   ///////    //////// ///    //  // 

...... Dernier avertissement ............................
`

// serveur mémoriel de Hope
var hope = Server{
	Address: "hope.local",
	Credentials: []Cred{
		{"hope", "tearsintherain", 5},
	},
	Description: hopDesc,
}

// Project "Hope"
// Dépot mémoriel
// - fenêtre temporelle glissante v12.5
// - compression McVaugh-Korba-Yang
// - contenu infix normalisé
// - (c) A.M
var hopDesc = `
01010000 01110010 01101111 01101010 01100101 01100011 01110100  00100010 01001000 01101111 01110000 01100101 00100010                                                                                                                                                                               
01000100 11101001 01110000 01101111 01110100  01101101 11101001 01101101 01101111 01110010 01101001 01100101 01101100                                                                                                                                                                               
00101101  01100110 01100101 01101110 11101010 01110100 01110010 01100101  01110100 01100101 01101101 01110000 01101111 01110010 01100101 01101100 01101100 01100101  01100111 01101100 01101001 01110011 01110011 01100001 01101110 01110100 01100101  01110110 00110001 00110010 00101110 00110101 
00101101  01100011 01101111 01101101 01110000 01110010 01100101 01110011 01110011 01101001 01101111 01101110  01001101 01100011 01010110 01100001 01110101 01100111 01101000 00101101 01001011 01101111 01110010 01100010 01100001 00101101 01011001 01100001 01101110 01100111                     
00101101  01100011 01101111 01101110 01110100 01100101 01101110 01110101  01101001 01101110 01100110 01101001 01111000  01101110 01101111 01110010 01101101 01100001 01101100 01101001 01110011 11101001                                                                                            
00101101  00101000 01100011 00101001  01000001 00101110 01001101                                                                                                                                                                                                                                    
`

var game = &Game{
	Network: []Server{
		dd,
		d22,
		kramps,
		kramps_priv,
		kramps_sec,
		corp,
		justice,
		abus,
		legba,
		legba_satcom,
		legba_archive,
		lbd,
		greendata,
		leet,
		lair,
		hope,
	},
}
