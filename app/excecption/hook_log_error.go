package excecption

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type CustomHook struct {
}

// handler jika level warm / error maka eksekusi hook
func (hook *CustomHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

// eksekusi perintah selanjutnya, contoh kirim email ke dev
func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	fmt.Println("sending email ustomHook : ", entry.Level, entry.Message)
	return nil
}
