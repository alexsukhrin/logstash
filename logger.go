package logstash

import (
	"fmt"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"log"
	"net"
)

type Config struct {
	Host, Port, ServiceName, Protocol string
}

func (c *Config) Logger() *logrus.Logger {
	logger := logrus.New()

	logstashConn, err := net.Dial(c.Protocol, fmt.Sprintf("%s:%s", c.Host, c.Port))

	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't send response %v", err))
	}

	hook := logrustash.New(
		logstashConn,
		logrustash.DefaultFormatter(
			logrus.Fields{"type": c.ServiceName}))

	logger.Hooks.Add(hook)

	return logger
}
