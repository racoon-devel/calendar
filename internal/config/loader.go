package config

import "github.com/jinzhu/configor"

// Load загружает конфигурацию из нескольких файлов, значения из переменных окружения переписывают значения из файлов.
// Если не указано файлов загрузка происходит из os.Environ
func Load(sources ...string) (config Configuration, err error) {
	err = configor.New(&configor.Config{ErrorOnUnmatchedKeys: true, ENVPrefix: "-"}).Load(&config, sources...)
	return
}
