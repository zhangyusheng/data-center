package logging

import (
	rotates "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var Logger *logrus.Logger

// Setup initialize the log instance
func Setup() {
	Logger = logrus.New()
	path := "/Users/huangmei/go/project/go.log"

	//下面配置日志每隔1小时轮转一个新文件，保留最近7天的日志文件，多余的自动清理掉。
	writer, _ := rotates.New(
		path+".%Y%m%d%H",
		rotates.WithLinkName(path),
		rotates.WithMaxAge(time.Duration(7 * 24) * time.Hour),
		rotates.WithRotationTime(time.Hour),
	)

	//同时写文件和屏幕
	writers := []io.Writer{writer, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)

	Logger.SetOutput(fileAndStdoutWriter)
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&logrus.JSONFormatter{})
}
