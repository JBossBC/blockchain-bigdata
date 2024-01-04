package database

import (
	"context"
	"fmt"
	"neo4j-producer/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	_Neo4jURLName  = "neo4j.url"
	_Neo4jUserName = "neo4j.username"
	_Neo4jPassword = "neo4j.password"
)

var (
	driver neo4j.DriverWithContext
)

func init() {
	var err error
	driver, err = neo4j.NewDriverWithContext(config.GetString(_Neo4jURLName), neo4j.BasicAuth(config.GetString(_Neo4jUserName), config.GetString(_Neo4jPassword), ""), func(c *neo4j.Config) {
		c.MaxConnectionPoolSize = 15
	})
	if err != nil {
		panic(fmt.Sprintf("neo4j连接出错:%s", err.Error()))
	}
	err = driver.VerifyConnectivity(context.TODO())
	if err != nil {
		panic(fmt.Sprintf("neo4j连接出错:%s", err.Error()))
	}
}

func getNeo4j() neo4j.SessionWithContext {
	return driver.NewSession(context.TODO(), neo4j.SessionConfig{})
}
