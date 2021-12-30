package logstash

import (
	"fmt"
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"net"
)

type Logger struct {
	Host, Port, ServiceName, Protocol string
}

func (l *Logger) GetLogger() *logrus.Logger {
	log := logrus.New()

	logstashConn, err := net.Dial(l.Protocol, fmt.Sprintf("%s:%s", l.Host, l.Port))

	if err != nil {
		log.Fatal(fmt.Sprintf("Couldn't send response %v", err))
	}

	hook := logrustash.New(
		logstashConn,
		logrustash.DefaultFormatter(
			logrus.Fields{"type": l.ServiceName}))

	log.Hooks.Add(hook)

	return log
}
