//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	"github.com/moemoe89/go-mongodb-gita/api/repository/mongo"
	"github.com/moemoe89/go-mongodb-gita/api/service"
	conf "github.com/moemoe89/go-mongodb-gita/config"
	"github.com/moemoe89/go-mongodb-gita/routers"

	"fmt"
	"log"

	"github.com/DeanThompson/ginpprof"
)

func main() {
	config, err := conf.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := conf.InitDB(config)
	if err != nil {
		log.Fatal(err)
	}

	logging := conf.InitLog()

	mongoRepo := mongo.NewMongoRepository(client, config.MongoDB.Database)
	svc := service.NewService(logging, mongoRepo)

	app := routers.GetRouter(logging, svc)
	ginpprof.Wrap(app)
	err = app.Run(":" + config.Port)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
