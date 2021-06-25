package server

import (
	"github.com/Aziz145/testaziz/tree/master/controller"
	"github.com/Aziz145/testaziz/tree/master/infra"
	"github.com/Aziz145/testaziz/tree/master/util"

	"github.com/gin-gonic/gin"

)

func Start() {
	server := gin.New()
	server.Use(gin.Recovery(), util.Logger())
	// db := infra.LoadPostgreSQLDB()
	// db := infra.LoadDatabase()
	db := infra.LoadMySQLDB()
	//db1 := infra.LoadOracleDB()
	//db2 := infra.LoadPostgreSQLDB()
	inDB := &controller.InDB{DB: db}
	// connDB := &controller.TwoDB{DB1: db1, DB2: db2}
	Routes(server, inDB)

	// Routes2(server, connDB)
	server.Run(":2416")
	// defer db.Close()
}