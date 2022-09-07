package initialize

import (
	"example/sample/global"
	"time"

	"github.com/rs/zerolog"
)

func InitComponents() {
	zerolog.TimeFieldFormat = time.RFC3339
	global.Viper = initViper()
	global.Db = initDb()
}
