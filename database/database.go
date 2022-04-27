package database
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &grom.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
		os.Exit(2)
	}

	log.Println("database connected");
	db.Logger = logger.Default.LogMode(logger.Info)
}