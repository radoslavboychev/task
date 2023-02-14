package repo

import (
	"fmt"
	"io/fs"
	"log"

	"github.com/boltdb/bolt"
	"github.com/radoslavboychev/task/models"
)

// InitDB
func InitDB(name string, port int16) (*bolt.DB, error) {
	db, err := bolt.Open(name, fs.FileMode(port), nil)
	if err != nil {
		return db, err
	}
	return db, nil
}

type TaskRepo struct {
	DB *bolt.DB
}

// NewTaskRepo
func NewTaskRepo(db *bolt.DB) TaskRepo {
	return TaskRepo{
		DB: db,
	}
}

// CreateBucket
func (tr *TaskRepo) CreateBucket(name string) error {
	tr.DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})
	log.Printf("created bucket %v", name)
	return nil
}

// CreateTask creates a task
func (tr *TaskRepo) CreateTask(t models.Task) error {
	tr.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("tasks"))
		err := b.Put([]byte(t.ID), []byte(t.Name))
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

// ViewDB prints all of the tasks from the bucket
func (tr *TaskRepo) ViewDB(bucket string) error {

	tr.DB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("Task: %v %v \n", string(k), string(v))
			return nil
		})

		return nil
	})
	return nil
}

// DoTask removes a key/value pair from a specific bucket that matches the ID of the key
func (tr *TaskRepo) DoTask(id string) error {
	err := tr.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		err := b.Delete([]byte(id))
		if err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return err
	}
	return nil
}
