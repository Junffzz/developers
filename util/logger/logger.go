package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

var (
	DebugLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

type Level int

var DefaultCallerDepth = 2
var LogRateLimit = 20

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {

	DebugLogger = log.New(os.Stdout, "", 0)
	//DebugLogger.SetPrefix(prefix(DEBUG))
	InfoLogger = log.New(os.Stdout, "", 0)
	//InfoLogger.SetPrefix(prefix(INFO))
	WarningLogger = log.New(os.Stdout, "", 0)
	//WarningLogger.SetPrefix(prefix(WARNING))
	ErrorLogger = log.New(os.Stdout, "", 0)
	//ErrorLogger.SetPrefix(prefix(ERROR))
	FatalLogger = log.New(os.Stdout, "", 0)
	//FatalLogger.SetPrefix(prefix(FATAL))

}

func Debug(v ...interface{}) {
	msgs := make([]interface{}, 0, 5)
	msgs = append(msgs, prefix(DEBUG))
	msgs = append(msgs, v...)

	DebugLogger.Println(msgs...)
}
func Debugf(format string, v ...interface{}) {

	DebugLogger.Printf(prefix(DEBUG)+" "+format, v...)
}

func Info(v ...interface{}) {
	msgs := make([]interface{}, 0, 5)
	msgs = append(msgs, prefix(INFO))
	msgs = append(msgs, v...)

	InfoLogger.Println(msgs...)
}

func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(prefix(INFO)+" "+format, v...)
}

func Warn(v ...interface{}) {
	msgs := make([]interface{}, 0, 5)
	msgs = append(msgs, prefix(WARNING))
	msgs = append(msgs, v...)

	WarningLogger.Println(msgs...)
}

func Warnf(format string, v ...interface{}) {

	WarningLogger.Printf(prefix(WARNING)+" "+format, v...)
}

type LoggerFunc func(format string, v ...interface{})

func WrapLog(ctx map[string]interface{}, f LoggerFunc, format string, v ...interface{}) {
	f("with context:%+v,"+format, ctx, v)
}

func Error(v ...interface{}) {
	msgs := make([]interface{}, 0, 5)
	msgs = append(msgs, prefix(ERROR))
	msgs = append(msgs, v...)

	ErrorLogger.Println(msgs...)
}

func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(prefix(ERROR)+" "+format, v...)
}

func Fatal(v ...interface{}) {
	msgs := make([]interface{}, 0, 5)
	msgs = append(msgs, prefix(FATAL))
	msgs = append(msgs, v...)

	FatalLogger.Fatalln(msgs...)
}
func Fatalf(format string, v ...interface{}) {

	FatalLogger.Fatalf(prefix(FATAL)+" "+format, v...)
}

func prefix(level Level) string {
	t := time.Now()

	data := fmt.Sprintf("%s.%06d", t.Format("2006-01-02 15:04:05"), t.Nanosecond()/1e3)

	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix := fmt.Sprintf("%s %5s %s:%d:", data, levelFlags[level], filepath.Base(file), line)
		return logPrefix
	} else {
		logPrefix := fmt.Sprintf("%s %5s:", data, levelFlags[level])
		return logPrefix
	}

}

//控制日志打印次数，减少日志打印量，外部传入计数，达到LogRateLimit返回true 不保证线程安全
func RateLimitAllow(counter *int) bool {
	if *counter%LogRateLimit == LogRateLimit-1 {
		*counter = 0
		return true
	}
	*counter++
	return false
}
