package main

// Génère des dates de concerts : Plastobéton et Black Wave
// sous forme de post pour le Forum frozpunk

import (
	"fmt"
	"math/rand"
	"time"
)

var servAddress = "froz.dd.local"

var concertHall = []string{"Neon Jungle", "Blade Runner", "Dark Light Matters",
	"Cyberpunk Cafe", "Cyberdome", "Technodrome", "Cyberspace",
	"VR Club", "AR Lounge", "Holodeck", "The Darkest Nightclub",
	"The Rainy Day Bar", "Stormy Night", "Ghost Cafe",
	"The Abandoned Warehouse", "Dark Alley Club", "The Underground Bunker",
	"Secret Hideout", "Forbidden Zone", "End of the World"}
var districts = []int{9, 13, 16, 17, 2, 23, 24, 25, 32}
var places []string

func genBarNames() []string {
	var res []string

	for _, ch := range concertHall {
		// chance outside D22
		district := 22
		if rnd.Float32() < 0.15 {
			district = districts[rnd.Intn(len(districts))]
		}
		res = append(res, fmt.Sprintf("%s (D%d)", ch, district))
	}
	return res
}

var rnd *rand.Rand

func genDate(start time.Time, interval int) time.Time {
	nbDay := rnd.Intn(interval)
	tmpDate := start.AddDate(0, 0, nbDay)
	nbMin := rnd.Intn(60 * 14)
	nbSec := rnd.Intn(60)
	res := tmpDate.Add(time.Minute*time.Duration(nbMin) +
		time.Second*time.Duration(nbSec))
	return res
}
func genPlace() string {
	// chance for Termina
	res := "Terminal - DZ"
	if rnd.Float32() < 0.90 {
		res = places[rnd.Intn(len(places))]
	}
	return res
}
func genConcertDates(nb int,
	subject string, content string, author string,
	start string, interval int, end string) {

	dateEnd, _ := time.Parse("2006-01-02T15:04:05", end)

	fmt.Printf("- server: %s\n", servAddress)
	fmt.Printf("  group: \"\"\n")
	fmt.Printf("  date: %s\n", start)
	fmt.Printf("  author: %s\n", author)
	fmt.Printf("  subject: %s\n", subject)
	fmt.Printf("  content: \"%s\"\n", content)
	fmt.Printf("  answers:\n")

	datePost, _ := time.Parse("2006-01-02T15:04:05", start)
	dateConcert := datePost
	dateOk := true
	for n := 0; n < nb && dateOk; n = n + 1 {
		// Pick a new date
		dateConcert = genDate(datePost, 30)
		if dateConcert.After(dateEnd) {
			dateConcert = dateEnd
			dateOk = false
		}
		// and a place
		place := genPlace()
		fmt.Printf("  - date: %s\n", datePost.Format("2006-01-02T15:04:05"))
		fmt.Printf("    author: %s\n", author)
		fmt.Printf("    content: \"Prochaine date %s au  %s\"\n",
			dateConcert.Format("02 Jan 2006"),
			place,
		)
		// date for the next Post
		datePost = genDate(dateConcert, 30)
		if datePost.After(dateEnd) {
			datePost = dateEnd
			dateOk = false
		}
	}
}

func main() {

	rnd = rand.New(rand.NewSource(time.Now().Unix()))
	places = append(places, genBarNames()...)

	genConcertDates(100, "«Plastobéton», tous les concerts (enfin, on essaie)",
		"Ca frappe, ça jette, ça envoie, ça te r'tourne et t'en red'mande... Ils sont là !",
		"zilmir",
		"2001-04-15T01:12:47", 15, "2007-11-30T00:00:00")
	fmt.Printf("  - date: 2007-12-21T12:00:31\n")
	fmt.Printf("    author: zilmir\n")
	fmt.Printf("    content: \"L'aventure s'est arrêtée trop tôt...\"\n")
	genConcertDates(50, "Les Concerts à venir du Black Wave",
		"L'événement à ne pas manquer. Ramène ta fraise et ta vibe",
		"bettyb",
		"2018-02-27T22:54:14", 15, "2020-07-15T00:00:00")
	fmt.Printf("  - date: 2020-06-06T22:15:52\n")
	fmt.Printf("    author: bettyb\n")
	fmt.Printf("    content: \"Prochaine date 29 Jul 2020 au Terminal - DZ\"\n")

}
