package models

import (
	"os"
	"github.com/skang28/golang_gin_mongo/db"
)


//put global variables here

// mongo server ip
var server = os.Getenv("DATABASE")

// database name
var databaseName = os.Getenv("DATABASE_NAME")

// connect to database
var dbConnect = db.NewConnection(server, databaseName)