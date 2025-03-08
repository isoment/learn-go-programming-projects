package db

import "github.com/boltdb/bolt"

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

// This is not the same as the built in init() function. It will not get called autmatically.
func Init(dbPath string) error {
	// Declare the error ahead of time so we can assign the the bold instance to the package
	// level db variable.
	var err error

	db, err = bolt.Open(dbPath, 0600, &bolt.Options{})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}
