  
package main

import (
	"github.com/Aziz145/testaziz/util"
	"github.com/Aziz145/testaziz/server"


)

func main() {
	util.LoadEnv()
	server.Start()
}
