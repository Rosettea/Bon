package main

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func InitDB() {
	_db, err := bolt.Open("bon.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	db = _db

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("links"))
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})
}

func AddLink(id, link string) {
	if err := db.Update(func(tx *bolt.Tx) error {	
		b := tx.Bucket([]byte("links"))
		if err := b.Put([]byte(id), []byte(link)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

func GetLink(id string) string {
	link := ""
	if err := db.View(func(tx *bolt.Tx) error {
		link = string(tx.Bucket([]byte("links")).Get([]byte(id)))
		return nil
	}); err != nil {
		fmt.Println(err)
	}

	return link
}

