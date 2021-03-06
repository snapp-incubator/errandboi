package config

import (
	"time"

	"github.com/snapp-incubator/errandboi/internal/db/mongodb"
	"github.com/snapp-incubator/errandboi/internal/db/rdb"
	"github.com/snapp-incubator/errandboi/internal/logger"
	"github.com/snapp-incubator/errandboi/internal/services/emq"
	"github.com/snapp-incubator/errandboi/internal/services/nats"
)

func Default() Config {
	const t = 10

	const p = 1883

	return Config{
		Redis: rdb.Config{
			Address:  "localhost:6379",
			Password: "",
			DB:       0,
		},
		Mongo: mongodb.Config{
			Name:    "errandboi",
			URL:     "mongodb://localhost:27017",
			Timeout: t * time.Second,
		},
		Emq: emq.Config{
			Broker:   "localhost",
			Port:     p,
			ClientID: "go_mqtt_client",
			Username: "emqx",
			Password: "",
		},
		Nats: nats.Config{
			URL: "nats://0.0.0.0:4222",
		},
		Logger: logger.Config{
			Level: "debug",
			Syslog: logger.Syslog{
				Enabled: false,
				Network: "",
				Address: "",
				Tag:     "",
			},
		},
	}
}
