package database

import (
	"fmt"
	"github.com/subosito/gotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

var DB *gorm.DB = nil
var err error = nil

func ConnectToDatabase() {
	err = gotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	newLogger := getNewGormLogger()
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG_DATABASE"))
	if debug {
		newLogger = newLogger.LogMode(logger.Info)
	}
	if DB == nil {
		connectToMySQL(newLogger)
	}
	if err != nil {
		panic(err)
	}
	if DB == nil {
		panic("failed to connect database!")
	}
}

func connectToMySQL(logger logger.Interface) {
	DB, err = gorm.Open(mysql.Open(os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+
		"@tcp("+os.Getenv("DATABASE_HOST")+":"+os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+
		"?charset=utf8mb4&collation=utf8mb4_bin&parseTime=True"), &gorm.Config{
		Logger:                                   logger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}

func getNewGormLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			Colorful: true,
		},
	)
}

func Migrate(object interface{}) {
	db := DB
	err := db.AutoMigrate(object)
	if err != nil {
		panic(err.Error())
	}
}
