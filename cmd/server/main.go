package main

import (
	"fmt"
	"time"

	"github.com/esgungor/engo-db/pkg/store"
	"github.com/esgungor/engo-db/pkg/wal"
)

func main() {
	wal, err := wal.NewWAL("wal.log")
	if err != nil {
		panic(err)
	}
	store := store.NewKeyValueStore(wal)
	store.Init()
	val, err := store.Get("abc")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func writeTest(s *store.KeyValueStore) {
	testVar := store.Entry{
		Timestamp: time.Now(),
		Operation: "INSERT",
		Key:       "abc",
		Value:     "123",
	}
	testVar2 := store.Entry{
		Timestamp: time.Now(),
		Operation: "INSERT",
		Key:       "abcd",
		Value:     "1234",
	}
	s.Insert(testVar)
	s.Insert(testVar2)
}
