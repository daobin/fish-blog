package db

type iDb interface {
	loadFile() string
	InitFile() error
}

var Mongo = new(mongo)
