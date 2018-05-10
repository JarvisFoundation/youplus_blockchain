package db

import (
	"fmt"
	"sync"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

const dbFile = "jarvis.db"

var _initCtx sync.Once
var _instance *LDB

//LDB ...
type LDB struct {
	file     string
	Database *leveldb.DB
}

//getLevelDatabase ...
func getLevelDatabase() *leveldb.DB {
	_initCtx.Do(func() {
		fmt.Println("Connecting to database...")
		_instance = new(LDB)
		dbObj, err := leveldb.OpenFile(dbFile, nil)

		if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
			dbObj, err = leveldb.RecoverFile(dbFile, nil)
		}
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("Failed to connect...")
			return
		}
		_instance.Database = dbObj
		fmt.Println("Connected to database...")
	})

	if _instance != nil {
		return _instance.Database
	}
	return nil
}

// Get Returns key value if present
func (db *LDB) Get(key []byte) []byte {
	ldb := getLevelDatabase()
	if ldb != nil {
		data, err := ldb.Get(key, nil)
		if err != nil {
			fmt.Println(err.Error())
		}
		return data
	}
	return nil
}

// Put Insert key value pair to database
func (db *LDB) Put(key []byte, value []byte) error {
	ldb := getLevelDatabase()
	if ldb != nil {
		return ldb.Put(key, value, nil)
	}
	return nil
}

// Delete Delete key value pair from database
func (db *LDB) Delete(key []byte) error {
	ldb := getLevelDatabase()
	if ldb != nil {
		return ldb.Delete(key, nil)
	}
	return nil
}
