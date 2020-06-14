//
//  Practicing MongoDB
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	ap "github.com/moemoe89/go-mongodb-gita/api"
	conf "github.com/moemoe89/go-mongodb-gita/config"
	"github.com/moemoe89/go-mongodb-gita/routers"

	"fmt"

	"github.com/DeanThompson/ginpprof"
)

func main() {
	client, err := conf.InitDB()
	if err != nil {
		panic(err)
	}

	log := conf.InitLog()

	repo := ap.NewMongoDBRepository(client)
	svc := ap.NewService(log, repo)

	app := routers.GetRouter(log, svc)
	ginpprof.Wrap(app)
	err = app.Run(":" + conf.Configuration.Port)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
