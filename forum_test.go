package main

import (
	"testing"
	"fmt"
	"os"
)

// Test Forum exists
func TestExists(t *testing.T) {
	// '.' HAS a forum
	fmt.Println("__forum '.'")
	_, err := GetForum( "." )
	if err != nil {
		t.Fatalf(`Should be a Forum in ".", %v`, err )
	}

	// 'noserver' DOES not have a Forum
	fmt.Println("__forum 'noserver'" )
	_, err = GetForum( "noserver" )
	if err == nil {
		t.Fatalf(`Should NOT be a Forum in "noserver", %v`, err )
	}
}
// Test listing of the forum topics
func TestList(t *testing.T)  {
	fmt.Println("__forum '.' => list")
	f, err := GetForum( "." )
	if err != nil {
		t.Fatalf(`Should be a Forum in ".", %v`, err )
	}
	f.DisplayTopics()
}


// Test enter subtopic
func TestSub(t *testing.T)  {
	fmt.Println("__forum '.' => enter 'news'")
	f, err := GetForum( "." )
	if err != nil {
		t.Fatalf(`Should be a Forum in ".", %v`, err )
	}
	err = f.EnterTopic( "news" )
	if err != nil {
		t.Fatalf(`Unable to enter %s from %s ".", %v`, "news", ".", err )
	}

	f.DisplayTopics()
}

// Test enter and leave
func TestEnterLeave(t *testing.T) {
	fmt.Println("__forum '.'")
	f, err := GetForum( "." )
	if err != nil {
		t.Fatalf(`Should be a Forum in ".", %v`, err )
	}
	startPoint := f.Address+f.Topic

	fmt.Println("__forum '.' => enter 'news'")
	err = f.EnterTopic( "news" )
	if err != nil {
		t.Fatalf(`Unable to enter %s from %s ".", %v`, "news", ".", err )
	}
	fmt.Println("__forum '.' => leaves 'news'")
	err = f.LeaveTopic()
	if err != nil {
		t.Fatalf(`Unable to leave from %s, %v`, "news", err )
	}
	endPoint := f.Address+f.Topic
	if startPoint != endPoint {
		t.Fatalf(`Should at startPoint start=%s ebd=%s, %v`,
			startPoint, endPoint, err)
	}
	fmt.Println("__forum '.' => leaves '.'")
	err = f.LeaveTopic()
	if err != nil {
		t.Fatalf(`Unable to leave from %s, %v`, ".", err )
	}
	if f.Address != "" || f.Topic != "" {
		t.Fatalf(`Should have left Forum %v`, err)
	}
}

// Test enter post
func TestEnterPost(t *testing.T) {
	fmt.Println("__forum '.'")
	f, err := GetForum( "." )
	if err != nil {
		t.Fatalf(`Should be a Forum in ".", %v`, err )
	}

	fmt.Println("__forum '.' => enter 'news'")
	err = f.EnterTopic( "news" )
	if err != nil {
		t.Fatalf(`Unable to enter %s from %s ".", %v`, "news", ".", err )
	}
	f.DisplayTopics()
	err = f.EnterTopicIndex(1)
	err = f.LeaveTopic()
	err = f.EnterTopicIndex(0)
}

// Test Add Post
func TestAddPost(t *testing.T) {
	fmt.Println("__forum '.'")
	f, err := GetForum( "." )
	if err != nil {
		t.Fatalf(`Should be a Forum in ".", %v`, err )
	}
	f.Show()
	//nbTopicHome := len(f.TopicList)

	fmt.Println("__forum '.' => enter 'news'")
	err = f.EnterTopic( "news" )
	if err != nil {
		t.Fatalf(`Unable to enter %s from %s ".", %v`, "news", ".", err )
	}
	f.Show()
	f.DisplayTopics()
	nbTopicNews := len(f.TopicList)

	err = f.AddPost("200319", "161027", "Essai N°1", "alain", "Premier essai\nPas forcément génial\n")
	if err != nil {
		t.Fatalf(`Unable to add new post, %v`, err )
	}
	f.Show()
	f.DisplayTopics()
	newTopicNews := len(f.TopicList)
	if newTopicNews != nbTopicNews+1 {
		t.Fatalf(`Post not added here, %v`, err)
	}
	f.EnterTopicIndex(newTopicNews-1)

	fmt.Println("__tries to add post whithout Leaving first")
	// Could not directly add. Must Leave first
	nbTopicNews = newTopicNews
	err = f.AddPost("200319", "162810", "Essai N°2", "alain", "bon, on va faire mieux\nUn jour\nsurement\n")
	if err == nil {
		t.Fatalf(`Should not be able to add new post, %v`, err )
	}
	f.LeaveTopic()
	f.Show()
	f.DisplayTopics()

	fmt.Println("__Removing new post BY HAND")
	err = os.Remove("./forum/news/200319_161027_Essai N°1_alain")
	if err != nil {
		t.Fatalf(`Unable to MANUALLY remove added post, %v`, err )
	}
}

// Test Add Topic
func TestAddTopic(t *testing.T) {
	fmt.Println("__forum '.'")
	f, err := GetForum( "." )
	if err != nil {
		t.Fatalf(`Should be a Forum in ".", %v`, err )
	}
	f.Show()
	//nbTopicHome := len(f.TopicList)

	fmt.Println("__forum '.' => add topic 'ajout'")
	err = f.AddTopic( "ajout" )
	if err != nil {
		t.Fatalf(`Unable to add topic 'ajout' in "." %v\n`, err )
	}
	f.DisplayTopics()

	fmt.Println("__forum '.' => enter 'ajout'")
	err = f.EnterTopic( "ajout" )
	if err != nil {
		t.Fatalf(`Unable to enter %s from %s ".", %v`, "ajout", ".", err )
	}
	f.Show()
	f.DisplayTopics()
}
