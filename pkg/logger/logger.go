package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	SetFormater(formatter logrus.Formatter)
	SetLevel(level uint32)
	WriteLog(level logrus.Level, message string, fields logrus.Fields)
	ReturnPrintf() func(string, ...interface{})
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

// Отдаем функцию для логгирования
func (l *log) ReturnPrintf() func(string, ...interface{}) {
	return l.log.Printf
}

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
