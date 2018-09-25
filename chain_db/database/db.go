package database

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/vitelabs/go-vite/log15"
)

var dbLog = log15.New("module", "chain_db/database")

func NewLevelDb(dbDir string) *leveldb.DB {
	db, err := leveldb.OpenFile(dbDir, nil)
	if err != nil {
		dbLog.Error(err.Error())
		return nil
	}
	return db
}