package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	SetFormater(formatter logrus.Formatter)
	SetLevel(level uint32)
	//Log() *logrus.Logger
	WriteLog(level logrus.Level, message string, fields logrus.Fields)
	SetHooks(hooks []logrus.Hook)
}

type log struct {
	log *logrus.Logger
}

// Создаем логгер
func New() Logger {
	return &log{log: logrus.New()}
}

// Установка хуков
func (l *log) SetHooks(hooks []logrus.Hook) {
	for _, hook := range hooks {
		l.log.AddHook(hook)
	}
}

// Устанавливаем логгеру формат
func (l *log) SetFormater(formatter logrus.Formatter) {
	// Установка формата
	l.log.Formatter = formatter
}

// Устанавливаем уровень логгирования
func (l *log) SetLevel(level uint32) {
	l.log.SetLevel(logrus.Level(level))
}

// Получить логгер
/*func (l *log) Log() *logrus.Logger {
	return l.log
}*/

// Запись информации в логгер
func (l *log) WriteLog(level logrus.Level, message string, fields logrus.Fields) {
	entry := l.log.WithFields(fields)

	switch level {
	case logrus.DebugLevel:
		entry.Debug(message)
	case logrus.InfoLevel:
		entry.Info(message)
	case logrus.WarnLevel:
		entry.Warn(message)
	case logrus.ErrorLevel:
		entry.Error(message)
	case logrus.FatalLevel:
		entry.Fatal(message)
	case logrus.PanicLevel:
		entry.Panic(message)
	default:
		entry.Info(message)
	}
}
