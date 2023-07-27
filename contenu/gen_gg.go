package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type EntryGG struct {
	Code      string
	splitKeys []string
	Owner     string
	title     string
	inside    string
}

var servAddress = "leet.darknet"

var ggList = []EntryGG{
	// ID keywords restricted owner title content
	{"vlope20", []string{"flr", "pr0n"}, "", "vanessalope", `login: green pass: nait5zee`},
	{"bitecoin19", []string{"flr", "pr0n"}, "", "lebitecoin", `login: green pass: ohphe0cu`},
	{"qtf20", []string{"flr", "pr0n"}, "", "QueTesFans", `login: green pass: aesahm0l`},
	{"pndr20", []string{"flr", "pr0n"}, "", "Pinederer", `login: green pass: ohdaf9uo`},
	{"jm20", []string{"flr", "pr0n"}, "", "Jockey & Micheline", `login: green pass: eig0thob`},
	{"005672bR1An ", []string{"maman", "grocery"}, "greenglass", "grocetag458", `tiger power, cheetos`},
	{"005673bR1An", []string{"maman", "pr0n"}, "greenglass", "pr0n", `fistfukdenaines2mains.mov`},
	{"005674bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag006", `2006-2006 - The Solsburry Four - NeoLondon - LD-Aurora*Cobalt*Grim*Slice* - Validated contracts 17 - Estimated total number of contracts: 18 - Contract unfulfilled on 2006-06-07 Reason: Aurora Terminated On Duty | estimated error: 8.6%`},
	{"005675bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag459", `cotton tiges, beignets, chouquettes, semoule, cheetos`},
	{"005676bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag007", `2009-2014 - I Volteggiatori -NuevaRoma - LD-Phase*Pins*ShallowWater*Bull* - Validated contracts 7 - Estimated total number of contracts: 8 - Contract unfulfilled on 2009-09-14 Reason: Phase*ShallowWater Terminated On Duty | estimated error: 0.4%`},
	{"005677bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag460", `steaks hachés, poulet frit, dentifrice`},
	{"005678bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag461", `cheetos`},
	{"005679bR1An", []string{"maman", "pr0n", "movie"}, "greenglass", "pr0n", `Mummypees.mov`},
	{"005680bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag008", `2012-2012 - Kanibale - VecchiaFirenze - Ego*Stitches*ZenHook*BlankSheet - Validated contracts 2 - Estimated total number of contracts: 4 - Contract unfulfilled on 2012-06-18 Reason: Blanksheet*Ego*Zenhook Terminated On Duty | estimated error: 16.005%`},
	{"005681bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag009", `2013-2015 - Credo zabójcy - NowaWarszawa - LD-Design*DoubleFeature*ShinyBone*Arch*DeadZone - Validated contracts 19 - Estimated total number of contracts: 19 - Contract unfulfilled 0 `},
	{"005682bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag462", `lingettes, saucisson, mozzarella, cheetos`},
	{"005683bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag463", `tomates, sparadra, bacon, oeufs, chocolat dur, cheetos`},
	{"005684bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag010", `2015-2018 - Liberty-Damen - NordenBerlin - LD-FlukeAnt*LibertyBell*Pins*Design - Validated contracts 31 - Estimated total number of contracts: 32 - Contract unfulfilled on 2018-05-16 Reason: Design*FlukeAnt Terminated On Duty | estimated error: 23.8%`},
	{"005685bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag464", `chaussettes, ceinture, poivron, poires, mousse, lait, steak de petits pois, escalope de carottes, cheetos (les nouveaux au goût vanille)`},
	{"005686bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag011", `2015-2019 - Wonton Soup - OldParis - LD-Cobalt*ImpPulse*Ink*SamouraiShowdown*SliceNDice - Validated contracts 28 - Estimated total number of contracts: 29 - Contract unfulfilled on 2019-01-24 Reason: Ink*SamouraiShodown Terminated On Duty | estimated error: 2.93%`},
	{"005687bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag011", `2018-2019 - King Kong Five - Europole - LD-AlcoLine*Bull*Mirror*PoisonClock - Validated contracts 5 - Estimated total number of contracts: 6 - Contract unfulfilled on 2019-12-01 Reason: Alcoline Missing On Duty | estimated error: 44.5%`},
	{"005688bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag011", `2018-2019 - RoadRunners - Europole -  LD-Mamasita*Pins*ZenHook*BrashBeast - Validated contracts 14 - Estimated total number of contracts: 14 - Contract unfulfilled 0`},
	{"005689bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag465", `Tampons, lingettes, mascara noire, mascara bleu, lait démaquillant, cheetos`},
	{"005690bR1An", []string{"maman", "pr0n"}, "greenglass", "pr0n", `mummycums.mp3`},
	{"005691bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag011", `2018-2019 - BlastingRogues - Europole -  LD-Design*Silverpath*Stitches - Validated contracts 8 - Estimated total number of contracts: 9 - Contract unfulfilled on 2019-10-15 Reason: Cancelled contract | estimated error: 0.09%`},
	{"005692bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag466", `pain, crème pour le cul (hydratante, pas celle qui pique), yaourts hypoallergéniques, bananes roses, gin, whisky`},
	{"005693bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag467", `cheetos`},
	{"005694bR1An", []string{"maman", "music"}, "greenglass", "musicbadass", `blackwavefullalbum.7zip`},
	{"005695bR1An", []string{"maman", "pr0n"}, "greenglass", "pr0n", `fuckedbyherdogs.mov`},
	{"005696bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag468", `crème solaire, nougats, gants de vaisselle, frozeloops, petits sachets plastique, lait, cheetos`},
	{"005697bR1An", []string{"maman", "contract"}, "greenglass", "runteamtag011", `2019-12-31+ - Trigger Legacy - Europole - LD-Cobalt*Design*Pins*Stitches*Bull* - Validated contracts 4 - Estimated total number of contracts: 5 - Contract unfulfilled 0 - Current Contracts: Contract A = Castle Corp. [Disney-Dassault] / estimated ¥€$ = 50.000 end 2020-07-31 | Contract B = Ubink Inc. [MetaSoft] / estimated ¥€$ = 25.000 / end 2020-08-10 |`},
	{"005698bR1An", []string{"maman", "bot"}, "greenglass", "2020-07-05T00:00:00", `2020/07/05 - nb entrées 0 | total data: 0`},
	{"005698bR1An", []string{"maman", "bot"}, "greenglass", "2020-07-12T00:00:00", `2020/07/12 - nb entrées 0 | total data: 0`},
	{"005698bR1An", []string{"maman", "bot"}, "greenglass", "2020-07-19T00:00:00", `2020/07/19 - nb entrées 0 | total data: 0`},
	{"005698bR1An", []string{"maman", "bot"}, "greenglass", "2020-07-26T00:00:00", `2020/07/26 - nb entrées 1 | total data: 1 - The Team/LD-Mercury*Dozer - Contract: zone isolation -     starting 2020/07/29:20:00 -     Mr Johnson`},
	{"005699bR1An", []string{"maman", "grocery"}, "greenglass", "grocetag469", `cola, cheetos`},
	{"005700bR1An", []string{"maman", "pr0n"}, "greenglass", "pr0n", `MummyshitsonDaddy.mov`},
}

var startDate, _ = time.Parse("2006-01-02T15:04:05", "2020-04-06T12:42:23")

func genNextWeek(start time.Time) time.Time {
	nbDay := rand.Intn(10)
	tmpDate := start.AddDate(0, 0, nbDay)
	//alter the time
	nbMin := rand.Intn(60) - 30
	nbSec := rand.Intn(60)
	res := tmpDate.Add(time.Minute*time.Duration(nbMin) +
		time.Second*time.Duration(nbSec))
	return res
}
func genNextDate(start time.Time) time.Time {
	nbDay := 7
	tmpDate := start.AddDate(0, 0, nbDay)
	//alter the time
	nbMin := rand.Intn(60) - 30
	nbSec := rand.Intn(60)
	res := tmpDate.Add(time.Minute*time.Duration(nbMin) +
		time.Second*time.Duration(nbSec))
	return res
}
func genGGGrocery() {
	// First the topic
	fmt.Printf("- server: %s\n", servAddress)
	fmt.Printf("  group: gg\n")
	fmt.Printf("  date: %s\n", startDate.Format("2006-01-02T15:04:05"))
	fmt.Printf("  author: greenglass\n")
	fmt.Printf("  subject: List de courses pour la relou\n")
	fmt.Printf("  content: gnagnagna\n")
	fmt.Printf("  answers:\n")

	newDate := startDate
	for _, truc := range ggList {
		for _, key := range truc.splitKeys {
			if key == "grocery" {
				newDate = genNextWeek(newDate)
				fmt.Printf("  - date: %s\n", newDate.Format("2006-01-02T15:04:05"))
				fmt.Printf("    author: greenglass\n")
				fmt.Printf("    content: |\n")
				// split inside at '-'
				lines := strings.Split(truc.inside, "-")
				for _, l := range lines {
					fmt.Printf("      %s\n", l)
				}
			}
		}
	}
}
func genGGContract() {
	// First the topic
	fmt.Printf("- server: %s\n", servAddress)
	fmt.Printf("  group: gg\n")
	fmt.Printf("  date: %s\n", startDate.Format("2006-01-02T15:04:05"))
	fmt.Printf("  author: greenglass\n")
	fmt.Printf("  subject: Pour FatMike\n")
	fmt.Printf("  content: \"note: penser à augmenter mes tarifs\"\n")
	fmt.Printf("  answers:\n")

	newDate := startDate
	for _, truc := range ggList {
		for _, key := range truc.splitKeys {
			if key == "contract" {
				newDate = genNextDate(newDate)
				fmt.Printf("  - date: %s\n", newDate.Format("2006-01-02T15:04:05"))
				fmt.Printf("    author: greenglass\n")
				fmt.Printf("    content: |\n")
				fmt.Printf("      tag=%s\n", truc.title)
				// split inside at '-'
				lines := strings.Split(truc.inside, " -")
				for _, l := range lines {
					fmt.Printf("      %s\n", l)
				}
			} else if key == "bot" {
				fmt.Printf("  - date: %s\n", truc.title)
				fmt.Printf("    author: botGG742\n")
				fmt.Printf("    content: |\n")
				// split inside at '-'
				lines := strings.Split(truc.inside, " -")
				for _, l := range lines {
					fmt.Printf("      %s\n", l)
				}
			}
		}
	}
}

func main() {
	startDate, _ = time.Parse("2006-01-02T15:04:05", "2020-04-06T12:42:23")
	rand.Seed(0)
	genGGGrocery()

	startDate, _ = time.Parse("2006-01-02T15:04:05", "2020-04-06T12:42:23")
	genGGContract()
}
