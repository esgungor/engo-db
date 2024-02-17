package main

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

type insertFn func(Entry) error

type WAL interface {
	AppendWAL(event Entry, lock *sync.Mutex) error
	RecoverFromWAL(lock *sync.Mutex, fn insertFn) error
}

type wal struct {
	path string
}

func NewWAL(path string) (WAL, error) {
	return &wal{
		path: path,
	}, nil
}

func (w wal) AppendWAL(event Entry, lock *sync.Mutex) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.OpenFile(w.path, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(f)
	enc.Encode(event)
	if err != nil {
		return err
	}
	return nil
}

func (w wal) RecoverFromWAL(lock *sync.Mutex, fn insertFn) error {
	f, err := os.OpenFile(w.path, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)
	for {
		var entry Entry
		err := dec.Decode(&entry)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fn(entry)
	}
	return nil
}
