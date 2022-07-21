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
	Description: `Bienvenue sur le serveur public du District 22 d'Europole.`,
	Targets: []Target{
		{legba.Address, "Legba Voodoocom", 1, "public", "public"},
		{kramps.Address, "Kramps Security", 1, "public", "public"},
		{corp.Address, "Central Services", 1, "public", "public"},
		{abus.Address, "Association des Banques Unifiées Suisses", 1, "public", "public"},
		{greendata.Address, "Green Data, solution environnementale", 1, "public", "public"},
	},
	Detection: SEC2,
}

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
	Detection: SEC2,
}

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
}

// serveur de sécurité de la kramps
var kramps_sec = Server{
	Address: "sec.kramps.d22.eu",
	Credentials: []Cred{
		{"admin", "lkjqsod", 5},
	},
	Detection: SEC4,
}

// serveur des services corporatistes D22
var corp = Server{
	Address: "corp.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
}

// serveur bancaire du D22
var abus = Server{
	Address: "abus.d22.eu",
	Credentials: []Cred{
		{"public", "public", 1},
	},
}

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
}

// serveur privé de Crunch
var leet = Server{
	Address: "l33t.darknet",
	Credentials: []Cred{
		{"crunch", "hacktheplanet", 5},
	},
}

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
