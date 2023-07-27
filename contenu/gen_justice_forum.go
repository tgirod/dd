package main

import (
	"fmt"
	"strings"
)

type EntryCasier struct {
	Code      string
	splitKeys []string
	Owner     string
	title     string
	inside    string
}

var servAddress = "justice.corp.d22.eu"

type Identity struct {
	Login    string `storm:"id"`
	Password string
	Name     string
	Bank     bool
}

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
	afrieman = Identity{"afrieman", "far3ik", "Anton Frieman", true} // PNJ fan blackwave
	// TODO quelques employé•e•s de la kramps
	akremmer  = Identity{"akremmer", "sexgod22", "Alexandre Kremmer", true}   // security Kramps
	mdavidson = Identity{"mdavidson", "allbitches", "Milton Davidson", true}  // dir adjoint Kramps
	vredmint  = Identity{"vredmint", "lily-dorian", "Virginia Redmint", true} // assistante Kramps
	// TODO quelques employé•e•s de legba voodoocom
	atrebinsky = Identity{"atrebinsky", "56raz8", "Anthon Trebinsky", true}  // proj. Mandrake
	dyuong     = Identity{"dyuong", "gd86rw", "Dyop Yuong", true}            // proj. Mandrake
	eherswing  = Identity{"eherswing", "oh7fd4", "Emmet Herswing", true}     // proj. Mandrake
	jkuipers   = Identity{"jkuipers", "azgh4d", "Jordan Kuipers", true}      // proj. Mandrake
	jmfusion   = Identity{"jmfusion", "sg7vf4", "John-Mickael Fusion", true} // Manager LegbaV
	yblansein  = Identity{"yblansein", "tyg45g", "Youri Blansein", true}     // satcom
	// Employé de GreenData
	cyolinaro = Identity{"cyolinaro", "rtd98y", "Consuella Yolinaro", true} // gère contrats chez Green Data
)

var persCasiers = []EntryCasier{
	{mmathison.Login, nil, "", mmathison.Name, "Disparue\n- Incident 16485-4346B, Nexkemia Petrochemicals, 07/07/2000\n"},
	{"jd2051", nil, "", "John Doe (2051)", "***** Personne recherchée, mandat interdistrict PG/634/ID/765387 *****\n- D22/ag#867533654: agression à main armée (victime Sony HAARTZ)\n"},
	{mbellamy.Login, nil, "", mbellamy.Name, "- néant\n"},
	{sbronner.Login, nil, "", sbronner.Name, "- néant\n"},
	{cbellamy.Login, nil, "", cbellamy.Name, "- néant\n"},
	{jvillanova.Login, nil, "", jvillanova.Name, "***** Personne recherchée, mandat interdistrict PJ/676/ER/65534 *****\n- D22/cm#5674243: complicité de meurtre\n"},
	{ecanto.Login, nil, "", ecanto.Name, "- néant\n"},
	{ejohannesen.Login, nil, "", ejohannesen.Name, "- néant\n"},
	{jbranson.Login, nil, "", jbranson.Name, "- néant\n"},
	{jmfright.Login, nil, "", jmfright.Name, "***** Personne recherchée, mandat interdistrict PF/0865/EP/55463 *****\n- D21/rc#12785234452 rupture de contrat\n- \n- $$$SPECIAL$$$ contacter cont4yes@kitsune ¥€$ en rapport\n- \n"},
	{skmihalec.Login, nil, "", skmihalec.Name, "- néant\n"},
	{emartin.Login, nil, "", emartin.Name, "***** Personne recherchée, mandat interdistrict PF/1437/PM/02 *****\n- D21/rc#6542867 rupture contrat\n"},
	{mdubian.Login, nil, "", mdubian.Name, "***** Personne recherchée, mandat interdistrict PA/172/PD/945337 *****\n- D22/vm#23842834: vol à l'étalage\n- D22/vm#54327653: vol recette épicerie nuit\n- D22/vm#543299873: vol simple\n- D22/vm#547699823: vol graviscooter\n- D22/vm#753296671: vol à l'étalage\n"},
	{cmihalec.Login, nil, "", cmihalec.Name, "***** Personne recherchée, mandat interdistrict PF/0865/EP/55463 *****\n- D22/vd#765428736: vol données confidentielles \n"},
	{sjohannesen.Login, nil, "", sjohannesen.Name, "- néant\n"},
	{rmichu.Login, nil, "", rmichu.Name, "- néant\n"},
	{rglass.Login, nil, "", rglass.Name, "- néant\n"},
	{djohannesen.Login, nil, "", djohannesen.Name, "***** Personne recherchée, mandat interdistrict PF/0415/EG/55323 *****\n- D22/me#1275436253: double meurtre, arme à feu\n"},
	{dbonenfant.Login, nil, "", dbonenfant.Name, "- néant\n"},
	{hproskychev.Login, nil, "", hproskychev.Name, "***** Personne recherchée, mandat interdistrict PF/2964/EP/98254 *****\n- D22/vd#89875357678: vol données avec copyright\n"},
	{njasinski.Login, nil, "", njasinski.Name, "- néant\n"},
	{sjasinski.Login, nil, "", sjasinski.Name, "***** Personne recherchée, mandat interdistrict PF/7253/EP/90271 *****\n- D22/vd#1100298735: vol données sous brevet\n"},
	{ffceline.Login, nil, "", ffceline.Name, "***** Personne recherchée, mandat interdistrict PF/1001/EP/98682 *****\n- D22/pi#9867356873: piratage informatique\n- D22/am#18763725: association malfaiteurs\n"},
	{cmills.Login, nil, "", cmills.Name, "***** Personne recherchée\n- Disparue : main courante du 05/05/2013\n\n- $$$SPECIAL$$$ contacter mills.contact@weyland.eu \n- \n"},
	{lseptembre.Login, nil, "", lseptembre.Name, "- néant\n"},
	{edubian.Login, nil, "", edubian.Name, "***** Personne recherchée *****\n- D22/ou#7578538765: outrage et rébellion, EuroPol\n- D22/va#325363552: vandalisme\n- D22/td#89765363: tapage diurne répété\n- D22/tn#101002543: tapage nocturne\n"},
	{zabasolo.Login, nil, "", zabasolo.Name, "- néant\n"},
	{ebranson.Login, nil, "", ebranson.Name, "- néant\n"},
	{jkievain.Login, nil, "", jkievain.Name, "- néant\n"},
	{fmanson.Login, nil, "", fmanson.Name, "- néant\n"},
	{rkievain.Login, nil, "", rkievain.Name, "- néant\n      - >>> automated procedure: contact@kramps.d22.eu | #line>2\n"},
	{pdoberty.Login, nil, "", pdoberty.Name, "- néant\n"},
	{ajolivet.Login, nil, "", ajolivet.Name, "- néant\n"},
	{jvazzanna.Login, nil, "", jvazzanna.Name, "- néant\n"},
	{mklebert.Login, nil, "", mklebert.Name, "- néant\n"},
	{jbatista.Login, nil, "", jbatista.Name, "- néant\n"},
	{gsuleymanoglu.Login, nil, "", gsuleymanoglu.Name, "- néant\n"},
	{rwhite.Login, nil, "", rwhite.Name, "***** Personne recherchée, mandat interdistrict PF/3151/FZ/76429 *****\n-D22/tr#8563427735: traffic illégal\n-D22/re#16753823: recel répété\n"},
}

func genCasiers() {
	// First the topic
	fmt.Printf("- server: %s\n", servAddress)
	fmt.Printf("  group: allowed\n")
	fmt.Printf("  date: 2020-02-01T18:00:00\n")
	fmt.Printf("  author: greff03\n")
	fmt.Printf("  subject: \"**** Wanted List *****\"\n")
	fmt.Printf("  content: MAJ 2020 - A.G - bull-03.64a-2020\n")
	fmt.Printf("  answers:\n")

	for _, pers := range persCasiers {
		if strings.HasPrefix(pers.inside, "- néant") == false {

			fmt.Printf("  - date: 2020-02-01T18:00:00\n")
			fmt.Printf("    author: greff03\n")
			fmt.Printf("    content: |\n")
			fmt.Printf("      Etat Civil: %s\n", pers.title)
			// split inside at '-'
			lines := strings.Split(pers.inside, "-")
			for _, l := range lines {
				fmt.Printf("      %s", l)
			}
		}
	}
}
func genOk() {
	// First the topic
	fmt.Printf("- server: %s\n", servAddress)
	fmt.Printf("  group: allowed\n")
	fmt.Printf("  date: 2020-02-01T18:00:00\n")
	fmt.Printf("  author: greff03\n")
	fmt.Printf("  subject: Citoyen.e.s en règle\n")
	fmt.Printf("  content: MAJ 2020 - A.G - bull-03.63-2020\n")
	fmt.Printf("  answers:\n")

	for _, pers := range persCasiers {
		if strings.HasPrefix(pers.inside, "- néant") == true {

			fmt.Printf("  - date: 2020-02-01T18:00:00\n")
			fmt.Printf("    author: greff03\n")
			fmt.Printf("    content: |\n")
			fmt.Printf("      Etat Civil: %s\n", pers.title)
			fmt.Printf("      %s\n", pers.inside)
		}
	}
}
func main() {
	genCasiers()
	genOk()
}
