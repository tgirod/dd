package main

import (
	"flag"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

const host = "0.0.0.0"
const portPlayer = 1337
const portMonitor = 7331

var app *App

func main() {
	init := flag.Bool("init", false, "réinitialise la base de données")
	fg_post := flag.String("post", "no_file", "load a YAML file of posts")
	fg_msg := flag.String("msg", "no_file", "load a YAML file of messages")
	fg_transaction := flag.String("trans", "no_file", "load a YAML file of transactions")
	flag.Parse()

	// if loading, then backup current DB
	if *fg_post != "no_file" {
		// using the Underlying Bolt Database to make a copy
		err := db.Bolt.View(func(tx *bolt.Tx) error {
			err := tx.CopyFile("dirtydistrict_before_post.db", os.FileMode(int(0600)))
			return err
		})

		if err != nil {
			fmt.Printf("WARRNING could not make a copy of DB for post: %v\n", err)
		}

		//SerializePosts()
		LoadPosts(*fg_post)
	}
	// if loading, then backup current DB
	if *fg_msg != "no_file" {
		// using the Underlying Bolt Database to make a copy
		err := db.Bolt.View(func(tx *bolt.Tx) error {
			err := tx.CopyFile("dirtydistrict_before_msg.db", os.FileMode(int(0600)))
			return err
		})

		if err != nil {
			fmt.Printf("WARRNING could not make a copy of DB for msg: %v\n", err)
		}

		//SerializeMessages()
		LoadMessages(*fg_msg)
	}

	if *fg_transaction != "no_file" {
		// using the Underlying Bolt Database to make a copy
		err := db.Bolt.View(func(tx *bolt.Tx) error {
			err := tx.CopyFile("dirtydistrict_befor_trans.db", os.FileMode(int(0600)))
			return err
		})

		if err != nil {
			fmt.Printf("WARRNING could not make a copy of DB for trans: %v\n", err)
		}

		//SerializeTransactions()
		LoadTransactions(*fg_transaction)
	}
	//os.Exit(0)
	app = NewApp(*init)
	app.Start()
}
