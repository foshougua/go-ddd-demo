package logger

import (
	"ddd-test/infrastructure/config"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var GLog *logrus.Logger

func init() {
	GLog = GetDefaultGinLog()
}

func GetDefaultGinLog() *logrus.Logger {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		FieldMap:         nil,
		CallerPrettyfier: LogCallerFunc,
		PrettyPrint:      false,
	}
	l.Out = os.Stdout
	l.ReportCaller = true
	l.Level = getModeByStr(config.Cfg.Mode)
	return l
}

func NewDefaultGinLog(traceId string) *logrus.Logger {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		FieldMap:         nil,
		CallerPrettyfier: LogCallerFunc,
		PrettyPrint:      false,
	}
	l.Out = os.Stdout
	l.ReportCaller = true
	l.Level = getModeByStr(config.Cfg.Mode)
	return l
}

func getModeByStr(s string) logrus.Level {
	switch s {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.DebugLevel
	}
}

//处理调用信息
func LogCallerFunc(r *runtime.Frame) (function string, file string) {
	if r != nil {
		s := strings.Split(r.File, "/")
		sLen := len(s)
		if sLen >= 2 {
			return "/" + s[sLen-1] + "/" + s[sLen-2] + ":" + strconv.Itoa(r.Line), ""
		}
	}
	return "", ""
}
