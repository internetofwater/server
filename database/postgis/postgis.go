package postgis

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/geodan/gost/sensorthings/models"
	_ "github.com/lib/pq" // postgres driver
)

// GostDatabase implementation
type GostDatabase struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Schema   string
	Ssl      bool
	Db       *sql.DB
}

// NewDatabase initialises the PostgreSQL database
//	host = TCP host:port or Unix socket depending on Network.
//	user = database user
//	password = database password
//	database = name of database
//	ssl = Whether to use secure TCP/IP connections (TLS).
func NewDatabase(host string, port int, user string, password string, database string, schema string, ssl bool) models.Database {
	return &GostDatabase{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
		Schema:   schema,
		Ssl:      ssl,
	}
}

// Start the database
func (gdb *GostDatabase) Start() {
	//ToDo: implement SSL
	log.Println("Creating database connection...")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", gdb.Host, gdb.User, gdb.Password, gdb.Database)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	gdb.Db = db
	err2 := gdb.Db.Ping()
	if err2 != nil {
		log.Fatal("Unable to connect to database, check your network connection.")
	}

	log.Printf("Connected to database, host: \"%v\", port: \"%v\" user: \"%v\", database: \"%v\", schema: \"%v\" ssl: \"%v\"", gdb.Host, gdb.Port, gdb.User, gdb.Database, gdb.Schema, gdb.Ssl)
}

// CreateSchema creates the needed schema in the database
func (gdb *GostDatabase) CreateSchema(location string) error {
	create, err := GetCreateDatabaseQuery(location, gdb.Schema)
	if err != nil {
		return err
	}

	c := *create
	_, err2 := gdb.Db.Exec(c)
	if err2 != nil {
		return err2
	}

	return nil
}

// GetCreateDatabaseQuery returns the database creation script for PostgreSQL
func GetCreateDatabaseQuery(location string, schema string) (*string, error) {
	bytes, err := ioutil.ReadFile(location)
	if err != nil {
		return nil, err
	}

	content := string(bytes[:])
	formatted := fmt.Sprintf(content, schema, schema)
	return &formatted, nil
}
