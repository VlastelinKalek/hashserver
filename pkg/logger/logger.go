package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	SetFormater(formatter logrus.Formatter)
	SetLevel(level uint32)
	Log() *logrus.Logger
}

type log struct {
	log *logrus.Logger
}

// Создаем логгер
func New() Logger {
	return &log{log: logrus.New()}
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
func (l *log) Log() *logrus.Logger {
	return l.log
}
