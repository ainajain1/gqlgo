// // db.go

// package graph

// import (
// 	"os"

// 	"github.com/go-pg/pg"
// )

// func Connect() *pg.DB {
// 	connStr := os.Getenv("DB_URL")
// 	opt, err := pg.ParseURL(connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	db := pg.Connect(opt)
// 	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
// 		panic("PostgreSQL is down")
// 	}
// 	return db
// }

package graph

import (
	"os"
	"time"

	"github.com/ainajain1/gqlgo/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	connStr := os.Getenv("DB_URL")
	logLevel := logger.Info
	logger.Default.LogMode(logger.Warn)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NowFunc: func() time.Time {
			currentTime := time.Now()
			name, offset := currentTime.Zone()
			if name != "IST" {
				mysqlTime := currentTime.Add(time.Second * time.Duration(offset))
				return mysqlTime
			}
			return currentTime
		},

		// Values can be: Silent, Error, Warn, Info
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		panic("Failed to initialise database. Reason: " + err.Error())
	}
	DB = db
	// postgresDB, err := db.DB()
	if err != nil {
		panic("Failed to connect with DB" + err.Error())
	}
	autoMigrate()
	return db
}

// Auto migration should only be enabled in test env.
func autoMigrate() {
	// if !utils.IsProductionEnvironment() {
	DB.AutoMigrate(&model.Movie{})
	// }
}
