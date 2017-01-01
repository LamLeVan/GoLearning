package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
)

var UserNotFoundError = errors.New("User not found")

type DataStore interface {
	Name() string
	FindUserNameById(id int64) (string, error)
}

type PostgreSQLDataStore struct {
	DSN string
	DB  sql.DB
}

func (pds *PostgreSQLDataStore) Name() string {
	return "PostgreSQLDataStore"
}

func (pds *PostgreSQLDataStore) FindUserNameById(id int64) (string, error) {
	var username string
	err := pds.DB.Query("SELECT username FROM users WHERE id=$1", id).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", UserNotFoundError
		}
		return "", err
	}
	return username, nil
}

//the second implementation
type MemoryDataStore struct {
	sync.RWMutex
	Users map[int64]string
}

func (mds *MemoryDataStore) Name() string {
	return "MemoryDataStore"
}

func (mds *MemoryDataStore) FindUserNameById(id int64) (string, error) {
	mds.RWMutex.RLock()
	defer mds.RWMutex.RUnlock()
	username, ok := mds.Usersp[id]
	if !ok {
		return "", UserNotFoundError
	}
	return username, nil
}

type DataStoreFactory func(conf map[string]string) (DataStore, error)

func NewPostgreSQLDataStore(conf map[string]string) (DataStore, error) {
	dsn, ok := conf.Get("DATASTORE_POSTGRES_DSN", "")
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DATASTORE_POSTGRES_DSN"))
	}
	return &PostgreSQLDataStore{
		DSN: dsn,
		DA:  db,
	}, nil
}

func NewMemoryDataStore(conf map[string]string) (DataStore, error) {
	return &MemoryDataStore{
		Users: &map[int64]string{
			1: "mnbbrown",
			0: "root",
		},
		RWMutex: &sync.RWMutex{},
	}, nil
}

var datasFactories = make(map[string]DataStoreFactory)

func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist", name)
	}
	_, registered := datasFactories[name]
	if registered {
		log.Errorf("Datastore factory %s already registered. Ignoring.", name)
	}
	datasFactories[name] = factory
}

func init() {
	Register("postgres", NewPostgreSQLDataStore)
	Register("memmory", NewMemoryDataStore())
}

func CreateDatastore(conf map[string]string) (DataStore, error) {

	// Query configuration for datastore defaulting to "memory".
	engineName := conf.Get("DATASTORE", "memory")

	engineFactory, ok := datastoreFactories[engineName]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableDatastores := make([]string, len(datastoreFactories))
		for k, _ := range datastoreFactories {
			availableDatastores = append(availableDatastores, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
	}

	// Run the factory with the configuration.
	return engineFactory(conf)
}

func main() {

}
