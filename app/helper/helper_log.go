package helper

import (
	"github.com/denpaden/go-belajar-restfull-api/app/feature/user"
	"github.com/sirupsen/logrus"
	"os"
)

// seharusnya automatis memaksi func dari go untuk membaca workinf direktory project
func loggerPath() string {
	return "D:/GOPATH/training/go-belajar-restfull-api/logs/"
}

func NewLoggerFile() *logrus.Logger {
	path := loggerPath()
	dirName := path + GetCurrentYear()
	CreateDirektori(dirName)

	logger := logrus.New()
	//logger.SetFormatter(&logrus.JSONFormatter{})
	date := GetCurrentDateYYYYMMDD()
	fileName := dirName + "/" + date + "_applog.log"
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	return logger
}

func getUserLogin() *user.User {
	return user.NewUser("userdeni", "Deni")
}

func loggerField(user *user.User) logrus.Fields {
	return logrus.Fields{
		"username": user.Username,
		"name":     user.Name,
	}
}
func LoggerInfo(msg string) {
	logger := NewLoggerFile()
	//logger.AddHook(&excecption.CustomHook{})
	logger.WithFields(loggerField(getUserLogin())).Info(msg)

}

func LoggerWarn(msg string) {
	logger := NewLoggerFile()
	//logger.AddHook(&excecption.CustomHook{})
	logger.WithFields(loggerField(getUserLogin())).Warn(msg)

}

func LoggerError(msg string) {
	logger := NewLoggerFile()
	//logger.AddHook(&excecption.CustomHook{})
	logger.WithFields(loggerField(getUserLogin())).Error(msg)

}
