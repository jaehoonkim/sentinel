package main

import (
	"flag"
	"log"

	"github.com/NexClipper/sudory/pkg/server/config"
	"github.com/NexClipper/sudory/pkg/server/database"
	"github.com/NexClipper/sudory/pkg/server/events"
	"github.com/NexClipper/sudory/pkg/server/macro/channels"
	"github.com/NexClipper/sudory/pkg/server/route"
)

func main() {
	cfg := &config.Config{}
	flag.StringVar(&cfg.Database.Host, "db-host", "127.0.0.1", "Database's host")
	flag.StringVar(&cfg.Database.Port, "db-port", "3306", "Database's port")
	flag.StringVar(&cfg.Database.Username, "db-user", "", "Database's username")
	flag.StringVar(&cfg.Database.Password, "db-passwd", "", "Database's password")
	flag.StringVar(&cfg.Database.DBName, "db-dbname", "", "Database's dbname")

	configPath := flag.String("config", "../../conf/sudory-server.yml", "Path to sudory-server's config file")
	flag.Parse()

	cfg, err := config.New(cfg, *configPath)
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg)
	if err != nil {
		panic(err)
	}

	//events
	var eventContexts []events.EventContexter
	var eventConfig *events.Config
	//event config
	if eventConfig, err = events.NewConfig(*configPath); err != nil { //config file load
		panic(err)
	}
	//event config vaild
	if err = eventConfig.Vaild(); err != nil { //config vaild
		panic(err)
	}
	//event config make listener
	if eventContexts, err = eventConfig.MakeEventListener(); err != nil { //events regist listener
		panic(err)
	}
	//event manager
	eventInvoke := channels.NewSafeChannel(0)
	manager := events.NewManager(eventContexts, log.Printf)
	deactivate := manager.Activate(eventInvoke, len(eventContexts)) //manager activate
	defer func() {
		deactivate() //stop when closing
	}()
	events.Invoke = func(v *events.EventArgs) { eventInvoke.SafeSend(v) } //setting invoker

	r := route.New(cfg, db)

	r.Start(cfg.Host.Port)
}
