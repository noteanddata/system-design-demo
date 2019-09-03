package main

import "github.com/gocql/gocql"
import "log"

// global session,  safe for concurrent goroutine calls
var session *gocql.Session


func createCqlSession(cassandraConfig CassandraConfig) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cassandraConfig.hostname)
	cluster.Keyspace = cassandraConfig.keyspace
	cluster.Consistency = cassandraConfig.consistency

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return session, err
}
