package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type TaskBody struct {
	Description string     `json:"description"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type Task struct {
	Key   int
	Value TaskBody
}

// This is not the same as the built in init() function. It will not get called autmatically.
func Init(dbPath string) error {
	// Declare the error ahead of time so we can assign the the bolt instance to the package
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

func CreateTask(taskDescription string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		// Create an id for the task which we can use as the key
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(int(id64))

		taskBody := TaskBody{
			Description: taskDescription,
		}

		// Marshal the TaskBody struct to JSON
		taskJSON, err := json.Marshal(taskBody)
		if err != nil {
			return err // Return error if marshaling fails
		}
		return b.Put(key, taskJSON) // Store JSON in BoltDB
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var taskBody TaskBody
			_ = json.Unmarshal(v, &taskBody)
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: taskBody,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
	return err
}

// Int to byte slice, BoldDB only works with byte slices
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Convert byte slice to int.
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
