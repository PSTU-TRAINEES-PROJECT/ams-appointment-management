package db

import (
	"ams-appointment-management/app/config"
	"ams-appointment-management/app/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func NewGormDb(config config.Config) *gorm.DB {
	dbUrl := config.Db.DbUrl
	opts := &gorm.Config{}

	// if config.App.LogLevel == "debug" {
	// 	newLogger := logger.New(
	// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 		logger.Config{
	// 			SlowThreshold:             time.Second, // Slow SQL threshold
	// 			LogLevel:                  logger.Info, // Log level
	// 			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 			Colorful:                  true,        // Disable color
	// 		},
	// 	)
	// 	opts.Logger = newLogger
	// }

	db, err = gorm.Open(postgres.Open(dbUrl), opts)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(config.Db.DbMaxConnections)
	sqlDb.SetMaxOpenConns(config.Db.DbMaxConnections)

	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Appointment{})
}

func DbPing() error {
	sqlDb, _ := db.DB()
	err = sqlDb.Ping()
	return err
}
