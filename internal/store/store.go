package store

import (
	"encoding/json"
	"fmt"
	"frontend_go/internal/model"
	"go.etcd.io/bbolt"
	"sync"
)

var buckets = [][]byte{[]byte("records"), []byte("profiles"), []byte("events"), []byte("audits"), []byte("files")}

type Store struct {
	db      *bbolt.DB
	mu      sync.RWMutex
	writeMu sync.Mutex
}

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, nil)
	if e != nil {
		return nil, e
	}
	s := &Store{db: db}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func encode(v any) ([]byte, error) { return json.Marshal(v) }
func decode(b []byte, v any) error { return json.Unmarshal(b, v) }
func (s *Store) put(bucket, key string, v any) error {
	data, e := encode(v)
	if e != nil {
		return e
	}
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, v any) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if b == nil {
			return fmt.Errorf("not found")
		}
		return decode(b, v)
	})
}
func (s *Store) PutRecord(v model.Record) error { return s.put("records", v.ID, v) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var v model.Record
	e := s.get("records", id, &v)
	return v, e
}
func (s *Store) PutProfile(v model.Profile) error { return s.put("profiles", v.ID, v) }
func (s *Store) GetProfile(id string) (model.Profile, error) {
	var v model.Profile
	e := s.get("profiles", id, &v)
	return v, e
}
func (s *Store) PutEvent(v model.Event) error { return s.put("events", v.ID, v) }
func (s *Store) GetEvent(id string) (model.Event, error) {
	var v model.Event
	e := s.get("events", id, &v)
	return v, e
}
func (s *Store) PutAudit(v model.Audit) error { return s.put("audits", v.ID, v) }
func (s *Store) GetAudit(id string) (model.Audit, error) {
	var v model.Audit
	e := s.get("audits", id, &v)
	return v, e
}
func (s *Store) PutFile(v model.TrainingFile) error { return s.put("files", v.ID, v) }
func (s *Store) GetFile(id string) (model.TrainingFile, error) {
	var v model.TrainingFile
	e := s.get("files", id, &v)
	return v, e
}
func (s *Store) ListRecords() ([]model.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out []model.Record
	if s.db == nil {
		return nil, fmt.Errorf("store closed")
	}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := decode(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
