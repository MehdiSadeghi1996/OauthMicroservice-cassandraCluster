package cassandra

import "github.com/gocql/gocql"

var (
	cluster *gocql.ClusterConfig
)

func init() {
	cluster = gocql.NewCluster("192.168.21.128")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
