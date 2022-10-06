package app

import (
	"log"
	"os"

	"github.com/couchbase/gocb/v2"
)

var cluster *gocb.Cluster

func NewCouchBaseInstance() {
	c, err := gocb.Connect(os.Getenv("COUCHBASE_HOST"), gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: os.Getenv("COUCHBASE_USERNAME"),
			Password: os.Getenv("COUCHBASE_PASSWORD"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	cluster = c
}

func GetCouchBaseInstance() *gocb.Cluster {
	return cluster
}
