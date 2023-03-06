package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
	"time"
	//"errors"
)

type Forum struct {
	// adresse du Forum
	Address string

	// Topic en train d'être visité
	Topic string

	// InPost ?
	InPost bool

	// Tous les Topics actuellement accessibles
	TopicList []fs.FileInfo

	// Titre du Post/Topic en train d'être ajouté
	CurrentTitle string

	// Index du Post/Topic qui est en "haut de page"
	IndexShow int
}

// TODO garder filename ouvert, comme ça Topic ne signale que les Topic
// et InPost est plus simple (une propriété)
// et permet d'ajouter une commande Forum show

// var (
// 	errAlreadyAtForumRoot = errors.New("")
// )

func GetForum( serverAdress string ) (Forum, error) {
	_, err := os.Open(serverAdress+"/forum")
    if err != nil {
        fmt.Println(err)
		forum := Forum{}
        return forum, err
    }
	forum := Forum{serverAdress+"/forum", "", false, nil, "", 0}
	err = forum.GetFiles( "" )
	return forum, err
}

func (f Forum) Show() {
	fmt.Printf("Forum @%s:%s (%t) %d topics\n", f.Address, f.Topic, f.InPost, len(f.TopicList))
}

func (f *Forum) GetFiles( topicStr string ) error {
	ff, err := os.Open( f.Address+"/"+topicStr )
    if err != nil {
        return err
    }

	f.TopicList, err = ff.Readdir(0); // all entries
	f.InPost = false
    if err != nil {
        return err
    }

	// Sort files: Topics, the InPost
	// From Old to New
	sort.Slice(f.TopicList,
		func(i, j int) bool {
			fOne := f.TopicList[i]
			fTwo := f.TopicList[j]

			if fOne.IsDir() {
				if fTwo.IsDir() {
					return fOne.Name() < fTwo.Name()
				} else {
					return true
				}
			} else {
				if fTwo.IsDir() {
					return false
				} else {
					// Both are PostName, need to compare date, then Name
					t1,n1,_,_ := GetElements(fOne.Name())
					t2,n2,_,_ := GetElements(fTwo.Name())
					if t1.Equal(t2) {
						return n1 < n2
					} else {
						return t1.Before(t2)
					}
				}
			}

		})

	return err
}

// Updates f.Topic if successful
// EnterTopic will update f.TopicList
func (f *Forum) EnterTopicIndex( index int ) error {
	finfo := f.TopicList[index]
	if finfo.IsDir() {
		return f.EnterTopic( finfo.Name() )
	} else {
		return f.EnterPost( finfo.Name() )
	}
}
// Updates f.TopicList if successful
// Updates f.Topic if successful
func (f *Forum) EnterTopic( name string ) error {
	ff, err := os.Open( f.Address+"/"+f.Topic+"/"+name )
	if err != nil {
		return err
	}

	f.Topic = f.Topic+"/"+name
	f.TopicList, err = ff.Readdir(0); // all entries
	f.InPost = false
	return err
}
func (f *Forum) LeaveTopic() error {
	// cannot leave "server/forum"
	if f.Topic == "" {
		fmt.Println( "[LeaveTopic] Leaving Forum")
		*f = Forum{}
		return nil
	}

	// Remove last from f.Topic
	tokens := strings.Split( f.Topic, "/" )
	f.Topic = strings.Join( tokens[:len(tokens)-1], "/" )

	return f.GetFiles(f.Topic)
}
func (f *Forum) AddTopic( name string ) error {
	err := os.Mkdir( f.Address+"/"+f.Topic+"/"+name, 0755 )

	if err != nil {
		if os.IsExist(err) {
			return errTopicExists
		}
		return err
	}

	return f.GetFiles(f.Topic)
}

// Does not change f.TopicList
// Updates f.Topic if successful
func (f *Forum) EnterPost( name string ) error {
	// Display TitleBAR
	fmt.Println(DecodePostTitle(name))
	fmt.Println("---")

	// Display file
	dat, err := os.ReadFile( f.Address+"/"+f.Topic+"/"+name )
	if err != nil {
		return err
	}
	f.Topic = f.Topic+"/"+name
	f.InPost = true
	fmt.Print(string(dat))
	return nil
}

func (f *Forum) AddPost( date string,
	time string,
	title string,
	author string,
	body string ) error {
	filename := date+"_"+time+"_"+title+"_"+author

	err := os.WriteFile(f.Address+"/"+f.Topic+"/"+filename, []byte(body), 0644)
	if err != nil {
		return err
	}
	// We must rebuild the current TopicList
	f.GetFiles(f.Topic)
	return err
}

// Enter a topic
// WIP: read and print files+dir inpath'
// index is optional (0 by default)
func (f *Forum) DisplayTopics(index int) {

	if index < 0 {
		index = 0
	}
	if index > len(f.TopicList) {
		index = len(f.TopicList)
	}
	// display AT MOST maxEntryDisplay entries
	lastIndex := index+maxEntryDisplay
	if lastIndex > len(f.TopicList) {
		lastIndex = len(f.TopicList)
	}

	for i, v := range f.TopicList[index:lastIndex] {
		if v.IsDir() {
			fmt.Printf( "%2d<T>: %s\n", i, v.Name())
		} else {
			fmt.Printf( "%2d<P>: %s\n", i, DecodePostTitle(v.Name()))
		}
	}
}
// Display start at index, AT MOST maxEntryDisplay entries
func (f *Forum) ListTopics(index int) []string {
	topics := make([]string, 0, len(f.TopicList))

	if index < 0 {
		index = 0
	}
	if index > len(f.TopicList) {
		index = len(f.TopicList)
	}
	f.IndexShow = index

	// display AT MOST maxEntryDisplay entries
	lastIndex := index+maxEntryDisplay
	if lastIndex > len(f.TopicList) {
		lastIndex = len(f.TopicList)
	}

	// Not at the beginning ?
	if index > 0 {
		topics = append(topics, "... : il y a des entrée avant !")
	}
	for i, v := range f.TopicList[index:lastIndex] {
		if v.IsDir() {
			topics = append(topics, fmt.Sprintf( "%2d<T>: %s",
				i+index, v.Name()))
		} else {
			topics = append(topics, fmt.Sprintf( "%2d<P>: %s",
				i+index, DecodePostTitle(v.Name())))
		}
	}

	if lastIndex < len(f.TopicList) {
		topics = append(topics, "... : la liste n'est pas finie !")
	}
	return topics
}
func (f *Forum) DisplayPost() []string {
	msg := make([]string, 0, 3)

	// name of file
	tokens := strings.Split( f.Topic, "/" )
	name := tokens[len(tokens)-1]

	msg = append(msg, DecodePostTitle(name))
	msg = append(msg, "---")

	// Display file
	dat, _:= os.ReadFile( f.Address+"/"+f.Topic )
	msg = append(msg, string(dat))

	return msg
}
// Display, either as a list of Topics/Post or content of Post
// index < 0 => use f.IndexShow
func (f *Forum) Display(index int) []string {
	if f.InPost {
		return f.DisplayPost()
	} else {
		if index < 0 {
			return f.ListTopics(f.IndexShow)
		} else {
			return f.ListTopics(index)
		}
	}
}
// WIP: decode filename
func DecodePostTitle(name string) string {
	tokens := strings.Split(name, "_")

	date := tokens[0]
	time := tokens[1]
	topic := tokens[2]
	author := tokens[3]

	msg := topic + "  [" + author
	msg += ", le "+ FormatDate(date)
	msg += " à " + FormatTime(time) + "]"
	return msg
}

func GetElements(name string) (time.Time, string, string, error) {
	tokens := strings.Split(name, "_")

	date := tokens[0]
	timeStr := tokens[1]
	topic := tokens[2]
	author := tokens[3]

	t, err := time.Parse("060102150405", date+timeStr)
	return t, topic, author, err
}

func FormatDate(rawdate string) string {
	msg := rawdate[0:2] + "/" + rawdate[2:4] + "/" + rawdate[4:6]
	return msg
}

func FormatTime(rawtime string) string {
	msg := rawtime[0:2] + ":" + rawtime[2:4] + ":" + rawtime[4:6]
	return msg
}
