  
package main

import (
	"github.com/Aziz145/testaziz/tree/master/util"
	"github.com/Aziz145/testaziz/tree/master/server"


)

func main() {
	util.LoadEnv()
	server.Start()
}
