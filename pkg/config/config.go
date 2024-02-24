package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Парсим .env файл
func GetConfigENV(path, name string, cfg any) error {
	// Задаем настройки конфигурации
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	// Читаем конфиг
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("read in config: %v", err)
	}

	// Распарсинг конфига в структуру
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("unmarshall config: %v", err)
	}
	return nil
}
