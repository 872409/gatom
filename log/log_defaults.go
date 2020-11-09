package log

var defaultLog = New("default")

func Default() *Log {
	return defaultLog
}

func SetLevel(level Level) {
	defaultLog.SetLevel(level)
}

func SaveToFile(dir string, fileName string) {
	defaultLog.SaveToFile(dir, fileName)
}

func SaveToFileDefault(fileName ...string) {
	_fileName := "app"
	if len(fileName) > 0 {
		_fileName = fileName[0]
	}
	defaultLog.SaveToFile("log", _fileName)
}

func SetDefault(l *Log) {
	defaultLog = l
}

func Print(args ...interface{}) {
	defaultLog.Logger.Print(args...)
}
func Debug(args ...interface{}) {
	defaultLog.Logger.Debug(args...)
}

func Info(args ...interface{}) {
	defaultLog.Logger.Info(args...)
}

func Warn(args ...interface{}) {
	defaultLog.Logger.Warn(args...)
}

func Fatal(args ...interface{}) {
	defaultLog.Logger.Fatal(args...)
}
func Trace(args ...interface{}) {
	defaultLog.Logger.Trace(args...)
}
func Error(args ...interface{}) {
	defaultLog.Logger.Error(args...)
}
func Panic(args ...interface{}) {
	defaultLog.Logger.Panic(args...)
}

func Debugln(args ...interface{}) {
	defaultLog.Logger.Debugln(args...)
}

func Infoln(args ...interface{}) {
	defaultLog.Logger.Infoln(args...)
}

func Warnln(args ...interface{}) {
	defaultLog.Logger.Warnln(args...)
}

func Fatalln(args ...interface{}) {
	defaultLog.Logger.Fatalln(args...)
}
func Traceln(args ...interface{}) {
	defaultLog.Logger.Traceln(args...)
}
func Errorln(args ...interface{}) {
	defaultLog.Logger.Errorln(args...)
}

func Panicln(args ...interface{}) {
	defaultLog.Logger.Panicln(args...)
}

func Println(args ...interface{}) {
	defaultLog.Logger.Println(args...)
}

func Printf(format string, args ...interface{}) {
	defaultLog.Logger.Printf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLog.Logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLog.Logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLog.Logger.Warnf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLog.Logger.Fatalf(format, args...)
}
func Tracef(format string, args ...interface{}) {
	defaultLog.Logger.Tracef(format, args...)
}
func Errorf(format string, args ...interface{}) {
	defaultLog.Logger.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	defaultLog.Logger.Panicf(format, args...)
}
