package main

import (
	"go-couchbase/app"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {

	// init couchbase
	app.NewCouchBaseInstance()

	// init router
	router := gin.Default()
	app.InitRoute(router)

	go func() {
		_ = router.Run(":8080")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
