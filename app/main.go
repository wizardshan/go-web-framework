package main

import (
	"app/controller"
	"app/repository"
	"app/repository/ent"
	"app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"pkg/cache"
)

func main() {
	host := "127.0.0.1:3306"
	name := "test"
	username := "root"
	password := ""

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		name,
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	cache := cache.New()

	engine := gin.Default()

	repoUser := repository.NewUser(client, cache)
	servUser := service.NewUser(repoUser, cache)

	ctrUser := controller.NewUser(repoUser, servUser, cache)
	user := engine.Group("/user")
	{
		user.PUT("/{id}", controller.Wrapper(ctrUser.EditProfile))
	}

	engine.Run()
}
