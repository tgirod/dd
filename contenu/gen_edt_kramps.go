package main

import (
	"fmt"
	"math/rand"
 	"time"

	"gopkg.in/yaml.v3"
)
// Register représente registre mémoire qui peut être modifié pour contrôler quelque chose
type Register struct {
	Server      string
	Group       string
	Description string
	State       RegisterState   // état actuel
	Options     []RegisterState // valeurs possible
}
type RegisterState string

type Entry struct {
	Code string
	splitKeys []string
	Owner string
	title string
	inside string

}

var guards = []Entry{
		{"G-F5", []string{"Sig.", "Raffaellino", "Bombieri"}, "", "Sig. Raffaellino Bombieri", "Sig. Raffaellino Bombieri - 25/5/1963 - 1084dd5a-25ac-43c1-90f4-1bd3d7769680"},
		{"G-H292", []string{"Humberto", "Plana-Chacón"}, "", "Humberto Plana-Chacón", "Humberto Plana-Chacón - 2/5/1999 - 2e285752-3a57-4a6c-b4cc-556d3a9c873d"},
		{"G-A37", []string{"Jonathan", "Swift"}, "", "Jonathan Swift", "Jonathan Swift  - 11/4/1984 - 383de3ff-56eb-4745-b472-e046ff8e552e"},
		{"G-F254", []string{"Éric", "Huet"}, "", "Éric Huet", "Éric Huet - 14/7/1968 - 4ec26008-158e-433d-9fd0-7937a5ea4e13"},
		{"G-J74", []string{"Angelina", "Deledda"}, "", "Angelina Deledda", "Angelina Deledda - 21/12/1965 - d764f740-fca3-436e-b36f-8f4c63e923f1"},
		{"G-C50", []string{"Romana", "Tutino"}, "", "Romana Tutino", "Romana Tutino - 1/5/1986 - 199c14a3-ff9b-4dbe-9e27-1083c890d5d3"},
		{"G-L35", []string{"Azahar", "Marcos", "Piñol"}, "", "Azahar Marcos Piñol", "Azahar Marcos Piñol - 5/12/1990 - f68c0dd9-0a85-41f4-8a23-91dc30efd7b1"},
		{"G-C3", []string{"Harvey", "Zimmermann"}, "", "Harvey Zimmermann", "Harvey Zimmermann  - 15/11/1961 - a37dd901-6526-4913-8900-daf9af5f8fab"},
		{"G-F196", []string{"Mary", "Flynn"}, "", "Mary Flynn", "Mary Flynn - 5/9/1983 - 18518deb-c136-4045-8f58-32ed11621e25"},
		{"G-N19", []string{"Ing.", "Constantin", "Briemer"}, "", "Ing. Constantin Briemer", "Ing. Constantin Briemer - 30/10/2000 - ed3f45c9-4396-4675-b2e3-545ce188bbe5"},
		{"G-C279", []string{"Dott.", "Cecilia", "Passalacqua"}, "", "Dott. Cecilia Passalacqua", "Dott. Cecilia Passalacqua - 7/2/1969 - e3eca928-532f-4281-ab6d-c3c8539bfff4"},
		{"G-S20", []string{"Alfredo", "Vendetti"}, "", "Alfredo Vendetti", "Alfredo Vendetti - 15/5/1958 - 5f945848-4682-4282-b8d7-65aebb52be5c"},
	}

var prisoners = []Entry{
		{"PR-289", []string{"Josué", "Cobos"}, "", "Josué Cobos", "Josué Cobos - 1/3/1989 - 1e80def0-56e2-4334-a5e9-3900c68e924a"},
		{"PX-149", []string{"Iker", "Diaz-Figuerola"}, "", "Iker Diaz-Figuerola", "Iker Diaz-Figuerola - 19/1/1960 - ecf774b2-d824-4e39-83a6-71a3ea504d48"},
		{"PO-235", []string{"Shannon", "Ward"}, "", "Shannon Ward", "Shannon Ward - 19/9/1965 - 5b5e9a8f-7059-41fa-abc1-834fad65dcae"},
		{"PS-161", []string{"Dipl.-Ing.", "Rebecca", "Briemer"}, "", "Dipl.-Ing. Rebecca Briemer", "Dipl.-Ing. Rebecca Briemer - 26/3/1992 - a62cc240-954f-4731-a2d9-fcf1428a6deb"},
		{"PB-32", []string{"Rembrandt", "Gatto-Togliatti"}, "", "Rembrandt Gatto-Togliatti", "Rembrandt Gatto-Togliatti - 2/3/1955 - 5b073a69-cf82-41f4-a69b-766aa8dabb71"},
		{"PR-58", []string{"Gustavo", "Soto", "Blanca"}, "", "Gustavo Soto Blanca", "Gustavo Soto Blanca - 19/11/1965 - 9a5947a2-b306-4927-987f-41d0171fd21b"},
		{"PK-224", []string{"Sig.", "Sabatino", "Zecchini"}, "", "Sig. Sabatino Zecchini", "Sig. Sabatino Zecchini - 29/9/1980 - ca15d1f5-d2b5-471e-b7ed-bf374a65fe9f"},
		{"PF-47", []string{"Stefan", "Jasinski"}, "", "Stefan Jasinski", "Stefan Jasinski - 15/3/1988 - c54dd982-6526-cd42-8900-cbe27f5f8fab"},
		{"PC-261", []string{"Eusebia", "Pozo-Botella"}, "", "Eusebia Pozo-Botella", "Eusebia Pozo-Botella - 21/7/1967 - 715c2159-72ea-4738-a401-b5a58506e694"},
		{"PN-94", []string{"Nieves", "Vázquez", "Franco"}, "", "Nieves Vázquez Franco", "Nieves Vázquez Franco - 30/9/1961 - 8def9643-e69b-45ba-81fe-41ffbfddb4ef"},
		{"PR-1", []string{"Grégoire", "Menard"}, "", "Grégoire Menard", "Grégoire Menard - 24/3/1955 - 44a26bdf-0210-4607-abd7-aa9cadbbdf1c"},
		{"PO-114", []string{"Sébastien", "Albert"}, "", "Sébastien Albert", "Sébastien Albert - 3/3/2002 - 868b22aa-370e-4d7e-ae0f-26c22eb0ec1c"},
		{"PX-280", []string{"Ing.", "Eveline", "Buchholz"}, "", "Ing. Eveline Buchholz", "Ing. Eveline Buchholz - 9/2/1979 - 39de7fda-a12d-4da7-ac74-4d8802095f69"},
		{"PS-197", []string{"Bruce", "Hoffman"}, "", "Bruce Hoffman", "Bruce Hoffman - 28/4/1978 - a1e241a4-6b8d-4e2e-bb77-3f93242fa942"},
		{"PV-144", []string{"Stefania", "Cagnin"}, "", "Stefania Cagnin", "Stefania Cagnin - 22/8/1994 - 1e6d16c6-2f9c-4d07-804a-1b44d042cc1d"},
		{"PR-25", []string{"Adelgunde", "Krebs", "B.Eng."}, "", "Adelgunde Krebs B.Eng.", "Adelgunde Krebs B.Eng. - 30/10/1974 - fe435f6a-84ae-41c5-80ae-e4239a270b00"},
		{"PN-173", []string{"Sig.", "Cesare", "Mazzi"}, "", "Sig. Cesare Mazzi", "Sig. Cesare Mazzi - 24/6/1956 - ad980090-d686-4637-a204-90a0b956a68b"},
		{"PD-256", []string{"Univ.Prof.", "Wolf-Rüdiger", "Cichorius"}, "", "Univ.Prof. Wolf-Rüdiger Cichorius", "Univ.Prof. Wolf-Rüdiger Cichorius - 23/6/1950 - fe2944cf-bfa1-404f-8d72-4bfcf1129376"},
		{"PA-186", []string{"Isidora", "del", "Olivera"}, "", "Isidora del Olivera", "Isidora del Olivera - 30/9/1952 - e713a1d1-fbdc-4048-a12e-6f1f4961537a"},
		{"PK-222", []string{"Jolanda", "Guinizzelli-Panicucci"}, "", "Jolanda Guinizzelli-Panicucci", "Jolanda Guinizzelli-Panicucci - 21/4/1953 - 9363dd30-7e40-4d91-8ea3-db67ea58c45b"},
		{"PV-28", []string{"Swen", "Ullrich"}, "", "Swen Ullrich", "Swen Ullrich - 23/12/1990 - 6f2287cf-036e-4372-bc3e-ad5d408283d2"},
		{"PZ-176", []string{"Galo", "Santamaria", "Carro"}, "", "Galo Santamaria Carro", "Galo Santamaria Carro - 8/6/1959 - 2b19228c-3c84-4031-a3b6-ff4be465e1b5"},
		{"PS-58", []string{"Philippe", "Guillaume", "de", "Pinto"}, "", "Philippe Guillaume de Pinto", "Philippe Guillaume de Pinto - 23/9/2000 - 167b7a42-4fd6-4d12-aaf1-4f32c3704705"},
		{"PH-270", []string{"André", "de", "la", "Ribeiro"}, "", "André de la Ribeiro", "André de la Ribeiro - 14/2/2002 - 4000c990-9824-49e3-b53e-30d4b9c5bf6c"},
		{"PT-19", []string{"Saverio", "Trotta-Pontecorvo"}, "", "Saverio Trotta-Pontecorvo", "Saverio Trotta-Pontecorvo - 16/12/1953 - b40ca3bc-6f1b-480e-81b7-f6b744ab6d32"},
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

var places = []string{ "AC1", "AC2", "AC3", "DZ", "RR", "CE"}

var timeSlots = []string{ "10", "14", "16", "20"}

var days = []string{ "2507", "2607", "2707", "2807", "2907", "3007"}

var regGuards []Register
var regPriso  []Register

var rnd *rand.Rand

func rndPlace() string {
	return places[rnd.Intn(len(places))]
}

func gen_registries( peoples []string) []Register {
	var res []Register

	for _, p := range peoples {
		for _, d := range days {
			for _, t := range timeSlots {
				key := fmt.Sprintf("%s_%s_%s", d, t, p)
				val := rndPlace()


				var options []RegisterState
				for _, rs := range places {
					options = append(options, RegisterState(rs))
				}
				reg := Register{
					Server: "priv.kramps.d22.eu",
					Group: "",
					Description: key,
					State: RegisterState(val),
					Options: options,
				}
				//fmt.Printf("REG %v\n", reg)
				res = append(res, reg)
			}
		}
	}
	return res
}

// génère du YAML pour les registres de la Kramps
// go run contenu/gen_edt_kramps.go > reg_kramps.yaml
func main() {
	rnd = rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Printf("p=%s\n", rndPlace())
	// fmt.Printf("p=%s\n", rndPlace())
	// fmt.Printf("p=%s\n", rndPlace())

	var guardKeys []string
	for _, g := range guards {
		guardKeys = append(guardKeys, g.Code)
	}
	regGuards = append(regGuards, gen_registries( guardKeys)...)

	// serialize
	yamlGuards, err := yaml.Marshal(regGuards)
	if err != nil {
		panic(err)
	}
	fmt.Printf("## Guards *****\n%s\n", yamlGuards)


	var prisoKeys []string
	for _, p := range prisoners {
		prisoKeys = append(prisoKeys, p.Code)
	}
	regPriso = append(regPriso, gen_registries( prisoKeys)...)
	yamlPriso, err := yaml.Marshal(regGuards)
	if err != nil {
		panic(err)
	}
	fmt.Printf("## Priso *****\n%s\n", yamlPriso)
}
