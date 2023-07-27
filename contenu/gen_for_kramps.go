package main

// Un Forum "Personnel" où les gardiens sont listés
// Un Forum "Prisonner" om les prisonniers sont listé
import (
	"fmt"
	"math/rand"
	"time"
)

// to use data from "avant"
type Entry struct {
	Code      string
	splitKeys []string
	Owner     string
	title     string
	inside    string
}

var guards = []Entry{
	{"G-F5", []string{"Sig.", "Raffaellino", "Bombieri"}, "", "Raffaellino Bombieri", "Raffaellino Bombieri - 25/5/1963 - 1084dd5a-25ac-43c1-90f4-1bd3d7769680"},
	{"G-H292", []string{"Humberto", "Plana-Chacón"}, "", "Humberto Plana-Chacón", "Humberto Plana-Chacón - 2/5/1999 - 2e285752-3a57-4a6c-b4cc-556d3a9c873d"},
	{"G-A37", []string{"Jonathan", "Swift"}, "", "Jonathan Swift", "Jonathan Swift  - 11/4/1984 - 383de3ff-56eb-4745-b472-e046ff8e552e"},
	{"G-F254", []string{"Éric", "Huet"}, "", "Éric Huet", "Éric Huet - 14/7/1968 - 4ec26008-158e-433d-9fd0-7937a5ea4e13"},
	{"G-J74", []string{"Angelina", "Deledda"}, "", "Angelina Deledda", "Angelina Deledda - 21/12/1965 - d764f740-fca3-436e-b36f-8f4c63e923f1"},
	{"G-C50", []string{"Romana", "Tutino"}, "", "Romana Tutino", "Romana Tutino - 1/5/1986 - 199c14a3-ff9b-4dbe-9e27-1083c890d5d3"},
	{"G-L35", []string{"Azahar", "Marcos", "Piñol"}, "", "Azahar Marcos Piñol", "Azahar Marcos Piñol - 5/12/1990 - f68c0dd9-0a85-41f4-8a23-91dc30efd7b1"},
	{"G-C17", []string{"Rodolph", "Kievain"}, "", "Rodolph Kievain", "Rodolph Kievain - 11/3/1992 - 6e43fdc5-6518-247d-b927-6atfhu7be5c"},
	{"G-C3", []string{"Harvey", "Zimmermann"}, "", "Harvey Zimmermann", "Harvey Zimmermann  - 15/11/1961 - a37dd901-6526-4913-8900-daf9af5f8fab"},
	{"G-F196", []string{"Mary", "Flynn"}, "", "Mary Flynn", "Mary Flynn - 5/9/1983 - 18518deb-c136-4045-8f58-32ed11621e25"},
	{"G-N19", []string{"Ing.", "Constantin", "Briemer"}, "", "Constantin Briemer", "Constantin Briemer - 30/10/2000 - ed3f45c9-4396-4675-b2e3-545ce188bbe5"},
	{"G-C279", []string{"Dott.", "Cecilia", "Passalacqua"}, "", "Cecilia Passalacqua", "Cecilia Passalacqua - 7/2/1969 - e3eca928-532f-4281-ab6d-c3c8539bfff4"},
	{"G-S20", []string{"Alfredo", "Vendetti"}, "", "Alfredo Vendetti", "Alfredo Vendetti - 15/5/1958 - 5f945848-4682-4282-b8d7-65aebb52be5c"},
	{"S-01", []string{"Alfredo", "Vendetti"}, "", "Alexandre Kremmer", "Alexandre Kremmer - 02/9/1977 - 8def9643-e69b-46e4-3fdc5-b096-1b78ba7e71bb"},
	{"D-4B", []string{"Alfredo", "Vendetti"}, "", "Milton Davidson", "Milton Davidson - 29/10/1966 - aa6eb7d5be6b6-e43f-bd52-247d-b927-6bb753acde5c"},
	{"P-G42", []string{"Alfredo", "Vendetti"}, "", "Virginia Redmint", "Virginia Redmint - 1/6/1990 - 1f2275db-890f-465e-79d3-cd51-5a43dfe79c5d"},
}

var prisoners = []Entry{
	{"PR-289", []string{"Josué", "Cobos"}, "", "Josué Cobos", "Josué Cobos - 1/3/1989 - 1e80def0-56e2-4334-a5e9-3900c68e924a"},
	{"PX-149", []string{"Iker", "Diaz-Figuerola"}, "", "Iker Diaz-Figuerola", "Iker Diaz-Figuerola - 19/1/1960 - ecf774b2-d824-4e39-83a6-71a3ea504d48"},
	{"PO-235", []string{"Shannon", "Ward"}, "", "Shannon Ward", "Shannon Ward - 19/9/1965 - 5b5e9a8f-7059-41fa-abc1-834fad65dcae"},
	{"PS-161", []string{"Dipl.-Ing.", "Rebecca", "Briemer"}, "", "Rebecca Briemer", "Rebecca Briemer - 26/3/1992 - a62cc240-954f-4731-a2d9-fcf1428a6deb"},
	{"PB-32", []string{"Rembrandt", "Gatto-Togliatti"}, "", "Rembrandt Gatto-Togliatti", "Rembrandt Gatto-Togliatti - 2/3/1955 - 5b073a69-cf82-41f4-a69b-766aa8dabb71"},
	{"PR-58", []string{"Gustavo", "Soto", "Blanca"}, "", "Gustavo Soto Blanca", "Gustavo Soto Blanca - 19/11/1965 - 9a5947a2-b306-4927-987f-41d0171fd21b"},
	{"PK-224", []string{"Sig.", "Sabatino", "Zecchini"}, "", "abatino Zecchini", "Sabatino Zecchini - 29/9/1980 - ca15d1f5-d2b5-471e-b7ed-bf374a65fe9f"},
	{"PF-47", []string{"Stefan", "Jasinski"}, "", "Stefan Jasinski", "Stefan Jasinski - 15/3/1988 - c54dd982-6526-cd42-8900-cbe27f5f8fab"},
	{"PC-261", []string{"Eusebia", "Pozo-Botella"}, "", "Eusebia Pozo-Botella", "Eusebia Pozo-Botella - 21/7/1967 - 715c2159-72ea-4738-a401-b5a58506e694"},
	{"PN-94", []string{"Nieves", "Vázquez", "Franco"}, "", "Nieves Vázquez Franco", "Nieves Vázquez Franco - 30/9/1961 - 8def9643-e69b-45ba-81fe-41ffbfddb4ef"},
	{"PR-1", []string{"Grégoire", "Menard"}, "", "Grégoire Menard", "Grégoire Menard - 24/3/1955 - 44a26bdf-0210-4607-abd7-aa9cadbbdf1c"},
	{"PO-114", []string{"Sébastien", "Albert"}, "", "Sébastien Albert", "Sébastien Albert - 3/3/2002 - 868b22aa-370e-4d7e-ae0f-26c22eb0ec1c"},
	{"PX-280", []string{"Ing.", "Eveline", "Buchholz"}, "", "Eveline Buchholz", "Eveline Buchholz - 9/2/1979 - 39de7fda-a12d-4da7-ac74-4d8802095f69"},
	{"PS-197", []string{"Bruce", "Hoffman"}, "", "Bruce Hoffman", "Bruce Hoffman - 28/4/1978 - a1e241a4-6b8d-4e2e-bb77-3f93242fa942"},
	{"PV-144", []string{"Stefania", "Cagnin"}, "", "Stefania Cagnin", "Stefania Cagnin - 22/8/1994 - 1e6d16c6-2f9c-4d07-804a-1b44d042cc1d"},
	{"PR-25", []string{"Adelgunde", "Krebs", "B.Eng."}, "", "Adelgunde Krebs B.Eng.", "Adelgunde Krebs B.Eng. - 30/10/1974 - fe435f6a-84ae-41c5-80ae-e4239a270b00"},
	{"PN-173", []string{"Sig.", "Cesare", "Mazzi"}, "", "Cesare Mazzi", "Cesare Mazzi - 24/6/1956 - ad980090-d686-4637-a204-90a0b956a68b"},
	{"PD-256", []string{"Univ.Prof.", "Wolf-Rüdiger", "Cichorius"}, "", "Wolf-Rüdiger Cichorius", "Wolf-Rüdiger Cichorius - 23/6/1950 - fe2944cf-bfa1-404f-8d72-4bfcf1129376"},
	{"PA-186", []string{"Isidora", "del", "Olivera"}, "", "Isidora del Olivera", "Isidora del Olivera - 30/9/1952 - e713a1d1-fbdc-4048-a12e-6f1f4961537a"},
	{"PK-222", []string{"Jolanda", "Guinizzelli-Panicucci"}, "", "Jolanda Guinizzelli-Panicucci", "Jolanda Guinizzelli-Panicucci - 21/4/1953 - 9363dd30-7e40-4d91-8ea3-db67ea58c45b"},
	{"PV-28", []string{"Swen", "Ullrich"}, "", "Swen Ullrich", "Swen Ullrich - 23/12/1990 - 6f2287cf-036e-4372-bc3e-ad5d408283d2"},
	{"PZ-176", []string{"Galo", "Santamaria", "Carro"}, "", "Galo Santamaria Carro", "Galo Santamaria Carro - 8/6/1959 - 2b19228c-3c84-4031-a3b6-ff4be465e1b5"},
	{"PS-58", []string{"Philippe", "Guillaume", "de", "Pinto"}, "", "Philippe Guillaume de Pinto", "Philippe Guillaume de Pinto - 23/9/2000 - 167b7a42-4fd6-4d12-aaf1-4f32c3704705"},
	{"PH-270", []string{"André", "de", "la", "Ribeiro"}, "", "André de la Ribeiro", "André de la Ribeiro - 14/2/2002 - 4000c990-9824-49e3-b53e-30d4b9c5bf6c"},
	{"PT-19", []string{"Saverio", "Trotta-Pontecorvo"}, "", "Savrio Trotta-Pontecorvo", "Saverio Trotta-Pontecorvo - 16/12/1953 - b40ca3bc-6f1b-480e-81b7-f6b744ab6d32"},
	{"PI-135", []string{"Erin", "Kelly"}, "", "Erin Kelly", "Erin Kelly - 10/8/1978 - 13496acb-9112-4b34-9890-413f9f6d3343"},
	{"PP-250", []string{"Renata", "Tomasetti-Camicione"}, "", "Renata Tomasetti-Camicione", "Renata Tomasetti-Camicione - 25/11/1953 - 4c0f783a-59fd-46bf-be87-e97b97db9460"},
	{"PE-15", []string{"Isidro", "Cabezas", "Dalmau"}, "", "Isidro Cabezas Dalmau", "Isidro Cabezas Dalmau - 7/12/2000 - 8ec8657f-23e0-4114-b096-1b78ba7e71bb"},
	{"PN-44", []string{"Gregory", "Thomas"}, "", "Gregory Thomas", "Gregory Thomas - 2/5/1969 - fd62a7e7-6a87-4ca8-a629-b03e1cbe3ecf"},
	{"PL-187", []string{"Rocío", "Carbajo"}, "", "Rocío Carbajo", "Rocío Carbajo - 3/3/1983 - f23845bc-2952-4898-a165-99dd7dcfcb87"},
	{"PV-240", []string{"Amanda", "Pinamonte"}, "", "Amanda Pinamonte", "Amanda Pinamonte - 14/9/2000 - 1bd20f1a-9eea-47a3-a0b8-c46e39c762fe"},
	{"PO-270", []string{"Adelia", "de", "Marquez"}, "", "Adelia de Marquez", "Adelia de Marquez - 10/8/1961 - ed4529cc-118e-41b1-b8d5-d8b8f5e824b1"},
	{"PM-58", []string{"Jochem", "Scheibe-Köhler"}, "", "Jochem Scheibe-Köhler", "Jochem Scheibe-Köhler - 24/11/1959 - fa641ac5-2b4b-42b4-9ec7-0fe1a6241221"},
	{"PE-66", []string{"Fabrizia", "Agostini"}, "", "Fabrizia Agostini", "Fabrizia Agostini - 7/2/1981 - b8fa7024-2342-49f1-a94c-aa6eb7d5be6b"},
	{"PW-168", []string{"Jennifer", "Martin"}, "", "Jennifer Martin", "Jennifer Martin - 19/1/1959 - 0471c610-34c5-49cf-aaa1-a9e06584eb88"},
	{"PC-263", []string{"Gülsen", "Trapp"}, "", "Gülsen Trapp", "Gülsen Trapp - 9/10/1968 - 7554027a-4841-42d6-9ee3-222f16518ef8"},
	{"PX-217", []string{"Elizabeth", "Wilson"}, "", "Elizabeth Wilson", "Elizabeth Wilson - 18/11/1987 - 1f2275db-890f-465e-8a66-9ba5e59358c5"},
	{"PI-40", []string{"Che", "Rozas", "Montesinos"}, "", "Che Rozas Montesinos", "Che Rozas Montesinos - 12/12/2000 - 5ac440eb-c586-4051-9abd-736ee40c9ada"},
	{"PK-4", []string{"Gilles", "Pereira"}, "", "Gilles Pereira", "Gilles Pereira - 26/5/2002 - c378d251-d729-4a88-9187-70570b2ddb89"},
	{"PC-52", []string{"Marcial", "Rodrigo"}, "", "Marcial Rodrigo", "Marcial Rodrigo - 19/11/1953 - 3f5fe026-7da2-40b1-a799-159a0c2322d5"},
	{"PA-21", []string{"Pedro", "Ramirez"}, "", "Pedro Ramirez", "Pedro Ramirez - 8/7/1991 - 383ca3ff-58eb-4745-efff-e046ff8e552e"},
	{"PL-25", []string{"Valentine", "Clerc"}, "", "Valentine Clerc", "Valentine Clerc - 30/5/1970 - 0d4dfc82-e8f7-47e2-9c48-3726d2f8c064"},
	{"PG-11", []string{"Miguel", "Williams"}, "", "Miguel Williams", "Miguel Williams - 1/2/1970 - 937859bc-9d1d-4b8d-98ae-08e541df9159"},
	{"PN-144", []string{"Vera", "Pont", "Avilés"}, "", "Vera Pont Avilés", "Vera Pont Avilés - 12/7/1987 - 03c9b13c-f632-4fb6-b0cb-8aefb17e7c74"},
	{"PH-148", []string{"Robert", "Ellis"}, "", "Robert Ellis", "Robert Ellis - 5/3/1971 - 007862ec-20c8-48bf-aeef-43c4ba76b8c6"},
	{"PR-8", []string{"David", "Stewart"}, "", "David Stewart", "David Stewart - 1/2/1964 - acc87482-6c76-4212-a9e1-ab5060d4a6be"},
	{"PV-35", []string{"Marino", "Torre", "Canet"}, "", "Marino Torre Canet", "Marino Torre Canet - 28/6/1956 - d1648b3a-51a5-4239-9da9-edf3c3e8539e"},
	{"PH-169", []string{"Jeanne", "Pages", "de", "la", "Grondin"}, "", "Jeanne Pages de la Grondin", "Jeanne Pages de la Grondin - 29/6/1994 - 25494162-e01b-4b15-b7de-990648ca8d76"},
	{"PO-276", []string{"Sophia", "Nunez"}, "", "Sophia Nunez", "Sophia Nunez - 5/8/1991 - 0f84aada-680f-4450-b21b-07ae9f16f047"},
	{"PO-297", []string{"Gabriela", "Águila-Viana"}, "", "Gabriela Águila-Viana", "Gabriela Águila-Viana - 26/10/1957 - ec29fc1b-c977-4fce-bfba-b07e441aa5e9"},
	{"PJ-1", []string{"Rosa", "María", "Noguera", "Sastre"}, "", "Rosa María Noguera Sastre", "Rosa María Noguera Sastre - 25/12/1956 - 7284cbfb-39a8-4511-a4f9-0f52dc834fe1"},
	{"PW-12", []string{"Lisa", "Steele"}, "", "Lisa Steele", "Lisa Steele - 7/12/1952 - a7580a18-00ed-42ee-90d6-192bed734e86"},
}

type PhraseTemplate struct {
	NbArg    int // 1er guard, 2+ prisoners
	Template string
}

var phraseTypique = []PhraseTemplate{
	{1, "%s faites attention aux détails et assurez-vous que votre travail est toujours précis."},
	{1, "%s soyez efficace dans votre travail et essayez toujours de trouver des moyens d'améliorer votre productivité."},
	{1, "%s soyez professionnel dans votre comportement et dans votre tenue vestimentaire."},
	{1, "%s faut être plus flexible, s'adapter au changement."},
	{1, "%s pas assez motivé et déterminé à réussir, il faut se donner à fond."},
	{2, "%s a donné au prisonnier %s un livre."},
	{2, "%s a fait passer à %s une lettre non officielle."},
	{3, "%s a mal supervisé le résident %s pendant sa sortie, qui s'est battu avec %s."},
	{2, "%s aide le résident %s à préparer son dossier de libération."},
	{2, "%s devient bien trop proche de %s."},
	{3, "Lors de la pause, %s a laissé %s et %s seuls trop longtemps."},
	{2, "Insulte envers %s, %s se retrouve au trou."},
	{3, "%s isolé par %s et %s. Séjour à l'infirmerie."},
	{1, "%s a encore oublié de fermer correctement derrière lui."},
	{1, "%s a détourné une partie de la nourriture des résidents"},
	{1, "%s se comporte en petit chef. A approcher"},
}

var servAdrr = "priv.kramps.d22.eu"
var groupPers = "pers"
var groupDA = "diradj"

func rndCode(people []Entry) string {
	return people[rand.Intn(len(people))].Code
}
func genDate(start time.Time, interval int) time.Time {
	nbDay := rand.Intn(interval)
	tmpDate := start.AddDate(0, 0, nbDay)
	nbMin := rand.Intn(60 * 14)
	nbSec := rand.Intn(60)
	res := tmpDate.Add(time.Minute*time.Duration(nbMin) +
		time.Second*time.Duration(nbSec))
	return res
}

func genPhrases(nb int) {
	startDateStr := "2017-01-27T09:00:00"
	startDate, _ := time.Parse("2006-01-02T15:04:05", startDateStr)

	fmt.Printf("- server: %s\n", servAdrr)
	fmt.Printf("  group: %s\n", groupDA)
	fmt.Printf("  date: %s\n", startDateStr)
	fmt.Printf("  author: mdavidson\n")
	fmt.Printf("  subject: Toujours utile\n")
	fmt.Printf("  content: allez, c'est parti...\n")
	fmt.Printf("  answers:\n")

	for n := 0; n < nb; n = n + 1 {
		pt := phraseTypique[rand.Intn(len(phraseTypique))]
		date := genDate(startDate, 1000)
		guard := rndCode(guards)
		var phrase string
		if pt.NbArg == 1 {
			phrase = fmt.Sprintf(pt.Template, guard)
		} else if pt.NbArg == 2 {
			priso1 := rndCode(prisoners)
			phrase = fmt.Sprintf(pt.Template, guard, priso1)
		} else if pt.NbArg == 3 {
			priso1 := rndCode(prisoners)
			priso2 := rndCode(prisoners)
			for priso2 == priso1 {
				priso2 = rndCode(prisoners)
			}
			phrase = fmt.Sprintf(pt.Template, guard, priso1, priso2)
		}
		// one answer
		fmt.Printf("  - date: %s\n", date.Format("2006-01-02T15:04:05"))
		fmt.Printf("    author: mdavidson\n")
		fmt.Printf("    content: %s\n", phrase)
	}
}

func genPeople(subject string, content string, people []Entry) {

	fmt.Printf("- server: %s\n", servAdrr)
	fmt.Printf("  group: %s\n", groupPers)
	fmt.Printf("  date: 2020-02-01T19:23:00\n")
	fmt.Printf("  author: vredmint\n")
	fmt.Printf("  subject: %s\n", subject)
	fmt.Printf("  content: \"%s\"\n", content)
	fmt.Printf("  answers:\n")

	// one answer for each people
	for _, p := range people {
		fmt.Printf("  - date: 2020-02-02T10:00:00\n")
		fmt.Printf("    author: vredmint\n")
		fmt.Printf("    content: \"%s : %s\"\n", p.Code, p.inside)
	}
}

func main() {
	genPeople("Liste du Personnel", "-- m-à-j 2020 -- ", guards)
	genPeople("Liste des Résidents", "-- m-à-j 2020 -- ", prisoners)
	genPhrases(174)
}
