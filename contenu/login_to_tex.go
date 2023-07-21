package main

import (
	"fmt"
)

// Utilisation pour générer login_card.pdf
// go run login_to_tex.go > login_list.tex
// latexmk -pdf login_card.tex

type Identity struct {
	Login    string
	Password string
	Name     string
}

// identités corpo recopiées depuis l'ancienne version
var (
	amathison     = Identity{"amathison", "hai3ja", "Alan Mathison"}            // alan
	mmathison     = Identity{"mmathison", "mie6oo", "Mélody Mathison"}          // mel
	mbellamy      = Identity{"mbellamy", "ahng7e", "Margherita Bellamy"}        // rita
	sbronner      = Identity{"sbronner", "rahk0u", "Sebastian Bronner"}         // styx
	cbellamy      = Identity{"cbellamy", "xoh7sh", "Camélia Bellamy"}           // kapo
	jvillanova    = Identity{"jvillanova", "ay9aef", "Julius Villanova"}        // scalpel
	ecanto        = Identity{"ecanto", "ti3eim", "Eddy Canto"}                  // greko
	ejohannesen   = Identity{"ejohannesen", "obo4ie", "Edwin Johannesen"}       // jesus
	jbranson      = Identity{"jbranson", "aich8g", "Jonathan Branson"}          // escobar
	jmfright      = Identity{"jmfright", "uruw5g", "John Mac Fright"}           // cageot
	skmihalec     = Identity{"skmihalec", "paeh3l", "Sylvia Kemija Mihalec"}    // la fouine
	emartin       = Identity{"emartin", "thooy1", "Eva Martin"}                 // eva
	mdubian       = Identity{"mdubian", "iup1ie", "Michael Dubian"}             // fat mike
	cmihalec      = Identity{"cmihalec", "uequ8u", "Carlotta Mihalec"}          // kennedy
	sjohannesen   = Identity{"sjohannesen", "aiphu4", "Sabrina Johannesen"}     // savage girl
	rmichu        = Identity{"rmichu", "ool7ch", "Raoul Michu"}                 // raoul cool
	rglass        = Identity{"rglass", "ahzae2", "Rupert Glass"}                // green glass
	sglass        = Identity{"sglass", "si6aeb", "Stefie Glass"}                // stefie
	djohannesen   = Identity{"djohannesen", "loh1ie", "Daisy Johannesen"}       // chilly daisy
	dbonenfant    = Identity{"dbonenfant", "de4oiv", "Désiré Bonenfant"}        // frère ping
	hproskychev   = Identity{"hproskychev", "ooj4an", "Harald Proskychev"}      // papa proxy
	njasinski     = Identity{"njasinski", "eveth3", "Nikole Jasinski"}          // nikki
	sjasinski     = Identity{"sjasinski", "ie7uo2", "Stefan Jasinski"}          // sasquatch
	ffceline      = Identity{"ffceline", "boh6ay", "Franz-Ferdinand Celine"}    // celine
	cmills        = Identity{"cmills", "thue1d", "Camélia Mills"}               // cramille
	lseptembre    = Identity{"lseptembre", "cul1ol", "Lilas Septembre"}         // tiger doll
	edubian       = Identity{"edubian", "rooch7", "Eloïse Dubian"}              // sister morphine
	zabasolo      = Identity{"zabasolo", "aipho0", "Zilmir Abasolo"}            // zilmir
	ebranson      = Identity{"ebranson", "rae2ie", "Elisabeth Branson"}         // betty b
	jkievain      = Identity{"jkievain", "nie3oo", "Jordan Kievain"}            // abraham
	fmanson       = Identity{"fmanson", "tiuf0y", "Frédéric Manson"}            // crunch
	rkievain      = Identity{"rkievain", "aso2qu", "Rodolph Kievain"}           // one kick
	pdoberty      = Identity{"pdoberty", "aivei1", "Pete Doberty"}              // jacob
	rwhite        = Identity{"rwhite", "ies2su", "Richard White"}               // oggy
	ajolivet      = Identity{"ajolivet", "quai1a", "Adrien Jolivet"}            // cyrano
	mklebert      = Identity{"mklebert", "eis6ku", "Mickael Klebert"}           // iron mike
	jvazzanna     = Identity{"jvazzanna", "ueth4k", "Joseph Vazzanna"}          // small joe
	jbatista      = Identity{"jbatista", "yah6ae", "Johaquim Batista"}          // joe-rez
	gsuleymanoglu = Identity{"gsuleymanoglu", "zo1daa", "Georges Suleymanoglu"} // georges

	// PNJs
	afrieman = Identity{"afrieman", "far3ik", "Anton Frieman"} // PNJ fan blackwave
	// TODO quelques employé•e•s de la kramps
	// TODO quelques employé•e•s de legba voodoocom
)

// identités virtuelles fournies par Jésus et le FLR
var (
	hope           = Identity{"hope", "011011011011", "Hope"}
	mel            = Identity{"mel", "hope4ever", "Mel"}
	rocky          = Identity{"rocky", "pourquoi", "Rocky"}
	rita           = Identity{"rita", "1wantM0re", "Rita"}
	styx           = Identity{"styx", "plastobeton", "Styx"}
	kapo           = Identity{"kapo", "touspour1", "Kapo"}
	scalpel        = Identity{"scalpel", "m3dic!!", "Scalpel"}
	greko          = Identity{"greko", "FuckY00", "Greko"}
	jesus          = Identity{"jesus", "ZtqCtdlb42", "Jesus"}
	escobar        = Identity{"escobar", "M0n3y++", "Escobar"}
	cageot         = Identity{"cageot", "fr33dom", "Cageot"}
	lafouine       = Identity{"lafouine", "cplvfh3h3", "La Fouine"}
	eva            = Identity{"eva", "n3verAgain", "Eva"}
	fatmike        = Identity{"fatmike", "tamereenshort", "Fat"}
	kennedy        = Identity{"kennedy", "jensaisrien", "Kennedy"}
	savagegirl     = Identity{"savagegirl", "zeStyle!", "Savage"}
	raoulcool      = Identity{"raoulcool", "barb3rKing", "Raoul Cool"}
	greenglass     = Identity{"greenglass", "il0veM0m", "Green Glass"}
	chillydaisy    = Identity{"chillydaisy", "rb0nesQueen", "Chilly Daisy"}
	ping           = Identity{"ping", "n0tp0ng!!", "Frère Ping"}
	papaproxy      = Identity{"papaproxy", "4ragnar!", "Papa Proxy"}
	nikki          = Identity{"nikki", "3141593", "Nikki"}
	celine         = Identity{"celine", "f0rg3tme", "Céline"}
	cramille       = Identity{"cramille", "onsenbalance", "Cramille"}
	tigerdoll      = Identity{"tigerdoll", "karateGirl", "Tiger Doll"}
	sistermorphine = Identity{"sistermorphine", "Icanfly", "Sister Morphine"}
	zilmir         = Identity{"zilmir", "cl3v3r", "Zilmir"}
	bettyb         = Identity{"bettyb", "ZeK0nsol", "Betty B"}
	abraham        = Identity{"abraham", "killerSolo", "Abraham"}
	crunch         = Identity{"crunch", "umdpcfpnp3o", "Crunch"}
	onekick        = Identity{"onekick", "faitchier", "One Kick"}
	jacob          = Identity{"jacob", "el01se", "Jacob"}
	oggy           = Identity{"oggy", "y0pasC0ul", "Oggy"}
	cyrano         = Identity{"cyrano", "rbNbOne", "Cyrano"}
	ironmike       = Identity{"ironmike", "deadlymike", "Iron Mike"}
	smallbob       = Identity{"smallbob", "smallbob", "Small Bob"}
	joerez         = Identity{"joerez", "mfuck3r", "Joe-Rez"}
	jeanne         = Identity{"jeanne", "j", "Jeanne"}
	paula          = Identity{"paula", "mdpH@rd", "Paula"}
	georges        = Identity{"georges", "devine!", "Georges"}
	ringo          = Identity{"ringo", "l@cherien!", "Ringo"}
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
	{"Rocky", nil, rocky, true, false},
	{"Rita", &mbellamy, rita, false, true},
	{"Styx", &sbronner, styx, false, true},
	{"Kapo", &cbellamy, kapo, false, true},
	{"Scalpel", &jvillanova, scalpel, true, true},
	{"Greko", &ecanto, greko, false, true},
	{"jesus", &ejohannesen, jesus, false, true},
	{"Escobar", &jbranson, escobar, false, true},
	{"Cageot", &jmfright, cageot, true, true},
	{"La Fouine", &skmihalec, lafouine, false, true},
	{"Eva", &emartin, eva, true, true},
	{"Fat Mike", &mdubian, fatmike, true, true},
	{"Kenndy", &cmihalec, kennedy, true, true},
	{"Savage Girl", &sjohannesen, savagegirl, false, true},
	{"Raoul Cool", &rmichu, raoulcool, false, true},
	{"Green Glass", &rglass, greenglass, false, true},
	{"Chilly Daisy", &djohannesen, chillydaisy, false, true},
	{"Frère Ping", &dbonenfant, ping, false, true},
	{"Papa Proxy", &hproskychev, papaproxy, true, true},
	{"Nikki", &njasinski, nikki, false, true},
	{"Céline", &ffceline, celine, true, true},
	{"Cramille", &cmills, cramille, true, true},
	{"Tiger Doll", &lseptembre, tigerdoll, false, true},
	{"Sister Morphine", &edubian, sistermorphine, true, true},
	{"Zilmir", &zabasolo, zilmir, false, true},
	{"Betty B", &ebranson, bettyb, false, true},
	{"Abraham", &jkievain, abraham, false, true},
	{"Crunch", &fmanson, crunch, false, true},
	{"One Kick", &rkievain, onekick, false, true},
	{"Jacob", &pdoberty, jacob, false, true},
	{"Oggy", &rwhite, oggy, true, true},
	{"Iron Mike", &mklebert, ironmike, false, true},
	{"Joe-Rez", &jbatista, joerez, false, true},
	{"Cyrano", &ajolivet, cyrano, false, true},
	{"Small Bob", &jvazzanna, smallbob, false, true},
	{"Jeanne", nil, jeanne, false, false},
	{"Ringo", nil, ringo, false, false},
	{"Georges", &gsuleymanoglu, georges, false, false},
	{"Paula", nil, paula, false, false},
	// Only for printing
	{"Oggy YES", &gsuleymanoglu, oggy, false, true},
}

func main() {
	// list Ids
	for _, perso := range allPlayers {
		wanted := ""
		if perso.Wanted {
			wanted = "\\color{red} - Recherché·e"
		}
		idc := "??? "
		idclog := "--"
		idcmdp := "--"
		if perso.Known {
			idc = perso.IdCorp.Name
			idclog = perso.IdCorp.Login
			idcmdp = perso.IdCorp.Password
		}

		fmt.Printf("\\confpin{%s}{%s}{%s}{%s}{%s}{%s}\n",
			perso.Perso,
			idc+wanted,
			idclog, idcmdp,
			perso.IdVirt.Login, perso.IdVirt.Password)
	}
}
