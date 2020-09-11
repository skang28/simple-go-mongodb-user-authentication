package db

import(
	"os"
	"time"
	"gopkg.in/mgo.v2"
)

//Connection defines connection structure
type Connection struct {
	session *mgo.Session
}

//NewConnection handles connecting to a mongo database
func NewConnection(host string, dbName string) (conn *Connection) {
	info := &mgo.DialInfo{
		Addrs: []string{host},
		Timeout: 60 *time.Second,
		Database: dbName,
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &Connection{session}
	return conn
}

//Use connects to a certain collection/table
func (conn *Connection) Use(dbName, tableName string) (collection *mgo.Collection) {
	return conn.session.DB(dbName).C(tableName)
}

//Close handles the closing of database connection
func (conn *Connection) Close() {
	conn.session.Close()
	return
}