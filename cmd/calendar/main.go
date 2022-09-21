package main

import (
	"flag"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/racoon-devel/calendar/internal/config"
	"github.com/racoon-devel/calendar/internal/server"
	"github.com/racoon-devel/calendar/internal/service"
	"github.com/racoon-devel/calendar/internal/storage"
	"os"
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	log.SetLevel(log.InfoLevel)

	configFile := flag.String("config", "", "set path to config file")
	verbose := flag.Bool("verbose", false, "verbose mode")
	flag.Parse()

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	cfg, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("cannot load configuration: %s", err)
	}
	log.Infof("configuration loaded: %+v", cfg)

	conn, err := storage.New(&cfg.Database)
	if err != nil {
		log.Fatalf("connect to storage failed: %s", err)
	}

	srv := server.Server{
		Calendar: service.NewCalendar(conn),
	}
	if err := srv.ListenAndServer(cfg.Http.Endpoint); err != nil {
		log.Fatalf("server error: %s", err)
	}
}

func loadConfig(file string) (cfg config.Configuration, err error) {
	if file != "" {
		cfg, err = config.Load(file)
	} else {
		cfg, err = config.Load()
	}
	return
}
