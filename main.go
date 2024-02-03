package main

import (
	"fmt"
	"log"

	"go-mnc/database"
	"go-mnc/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
    db  *gorm.DB
    err error
)

func main() {
    db, err = database.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    router := gin.Default()
    routes.UserRoutes(router)


    port := 8080
    log.Printf("Server started on :%d\n", port)
    log.Fatal(router.Run(fmt.Sprintf(":%d", port)))
}
