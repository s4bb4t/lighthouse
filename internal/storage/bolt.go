package storage

import (
	"encoding/binary"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"go.etcd.io/bbolt"
)

type Storage struct {
	db *bbolt.DB
}

const (
	subs       = "subs.db"
	subsBucket = "subs"
)

func New() (*Storage, error) {
	bb, err := bbolt.Open(subs, 0600, bbolt.DefaultOptions)
	if err != nil {
		return nil, sp.New(sp.Sample{
			Messages: map[string]string{
				"en": "failed to open subs database",
			},
			Desc:  "Failed to open subs database",
			Hint:  "Check subs.db in root or your права на папку",
			Level: levels.LevelError,
			Cause: err,
		})
	}

	err = bb.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(subsBucket))
		if err != nil {
			return sp.New(sp.Sample{
				Messages: map[string]string{
					"en": "failed to create subs bucket",
				},
				Hint:  "Follow bbolt instructions",
				Desc:  "Failed to create subs bucket",
				Level: levels.LevelError,
				Cause: err,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Storage{db: bb}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Put(group string, id int64) error {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, id)
	b := buf[:n]

	return s.db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket([]byte(subsBucket)).Put(b, []byte(group))
		if err != nil {
			return sp.New(sp.Sample{
				Messages: map[string]string{
					"en": "failed to put subs bucket",
				},
				Desc:  "Failed to put subscribed user's id to bucket",
				Hint:  "Check db's mode",
				Level: levels.LevelError,
				Cause: err,
				Meta: map[string]any{
					"group": group,
					"id":    id,
				},
			})
		}
		return nil
	})
}

func (s *Storage) Read(group string) ([]int64, error) {
	var users []int64
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(subsBucket))
		if b == nil {
			return sp.New(sp.Sample{
				Messages: map[string]string{
					"en": "failed to read subs bucket",
				},
				Desc:  "Subs bucket not found. Seems like you deleted the bucket or subs.db file",
				Hint:  "Reload Bolt storage and do not remove or change the bucket or .db file",
				Level: levels.LevelError,
			})
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if group != "" && string(v) != group {
				continue
			}

			id, n := binary.Varint(k)
			if n == 0 {
				return sp.New(sp.Sample{
					Messages: map[string]string{
						"en": "failed to read user's id",
					},
					Desc:  "Invalid user id",
					Hint:  "Your user id is invalid - check what you tries to save",
					Level: levels.LevelError,
					Meta: map[string]any{
						"group": string(k),
						"id":    id,
						"bytes": v,
					},
				})
			}
			users = append(users, id)
		}
		return nil
	})
	return users, err
}
