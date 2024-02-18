package store

import (
	"errors"
	"sync"
	"time"

	"github.com/esgungor/engo-db/pkg/wal"
	"github.com/mitchellh/mapstructure"
)

type Entry struct {
	Timestamp time.Time
	Operation string
	Key       string
	Value     string
}

type KeyValueStore struct {
	Kv   map[string]string
	wal  wal.WAL
	lock sync.Mutex
}

func NewKeyValueStore(wal wal.WAL) KeyValueStore {

	return KeyValueStore{
		Kv:   map[string]string{},
		wal:  wal,
		lock: sync.Mutex{},
	}
}

func (k *KeyValueStore) recoverFromWAL() {
	k.lock.Lock()
	defer k.lock.Unlock()
	k.wal.RecoverFromWAL(&k.lock, func(e any) error {
		var entry Entry
		mapstructure.Decode(e, &entry)
		err := k.applyCommand(entry)
		return err
	})
}
func (k *KeyValueStore) Init() {
	k.recoverFromWAL()
}

func (k *KeyValueStore) applyCommand(e Entry) error {
	switch e.Operation {
	case "INSERT":
		k.Kv[e.Key] = e.Value
	case "UPDATE":
		_, ok := k.Kv[e.Key]
		if ok {
			k.Kv[e.Key] = e.Value

		} else {
			return errors.New("record not found")
		}
	case "DELETE":
		delete(k.Kv, e.Key)
	}
	return nil
}

func (k *KeyValueStore) Insert(entry Entry) error {
	err := k.wal.AppendWAL(entry, &k.lock)
	if err != nil {
		return err
	}
	err = k.applyCommand(entry)
	if err != nil {
		return err
	}
	return nil
}

func (k *KeyValueStore) Get(key string) (string, error) {
	_, ok := k.Kv[key]
	if ok {
		return k.Kv[key], nil
	}
	return "", errors.New("key not found")
}
