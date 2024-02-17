package main

import (
	"fmt"
	"time"
)

func main() {
	wal, err := NewWAL("wal.log")
	if err != nil {
		panic(err)
	}
	store := NewKeyValueStore(wal)
	store.Init()
	val, err := store.Get("abc")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func writeTest(store *KeyValueStore) {
	testVar := Entry{
		Timestamp: time.Now(),
		Operation: "INSERT",
		Key:       "abc",
		Value:     "123",
	}
	testVar2 := Entry{
		Timestamp: time.Now(),
		Operation: "INSERT",
		Key:       "abcd",
		Value:     "1234",
	}
	store.Insert(testVar)
	store.Insert(testVar2)
}
