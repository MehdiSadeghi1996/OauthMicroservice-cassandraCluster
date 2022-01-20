package cassandra

import "github.com/gocql/gocql"

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("192.168.21.128")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() (*gocql.Session, error) {
	return session, nil
}
