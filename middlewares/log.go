package middlewares

import (
	"github.com/astaxie/beego/config"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	AccessLog     *logrus.Logger
	ErrorLog      *logrus.Logger
	accessLogFile = "./access.log"
	errorLogFile  = "./error.log"
)

func Initialize(cfg config.Configer) {
	accessLogFile = cfg.String("log::access")
	initAccessLog()
}

func initErrorLog() {
	ErrorLog = logrus.New()
	ErrorLog.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(errorLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	ErrorLog.SetOutput(file)
}

func initAccessLog() {
	AccessLog = logrus.New()
	AccessLog.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(accessLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	AccessLog.SetOutput(file)
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)

	return
}