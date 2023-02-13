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
func NewTaskRepo(db *bolt.DB) *TaskRepo {
	return &TaskRepo{
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

// CreateTask
func (tr *TaskRepo) CreateTask(t models.Task) error {
	tr.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("tasks"))
		err := b.Put([]byte(t.ID), []byte(t.Name))
		if err != nil {
			return err
		}
		return nil
	})
	log.Printf("created task %v", t.Name)
	return nil
}

// ViewDB
func (tr *TaskRepo) ViewDB(bucket string) error {

	tr.DB.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %v value: %v\n", string(k), string(v))
		}

		return nil
	})
	return nil
}
