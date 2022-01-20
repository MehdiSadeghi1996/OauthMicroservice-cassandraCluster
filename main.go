package main

import (
	"fmt"
	"github.com/gocql/gocql"
)

func main() {

	cluster := gocql.NewCluster("192.168.21.128")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	fmt.Println(err)
	defer session.Close()

}
