package db

import (
	"fmt"
	"log"

	"github.com/Kartik30R/Tiket.git/config"
 	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(cfg *config.EnvConfig, DBMigrator func(*gorm.DB) error) *gorm.DB {
		uri := fmt.Sprintf(`
		host=%s user=%s dbname=%s password=%s sslmode=%s port=5432`,
		cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode,
	)


	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("couldn't connect to database %e",err )
	}

	if err:= DBMigrator(db); err!=nil{
		log.Fatalf("migration failed %e", err)
	}
	

	log.Println("connected to database")

	return db

}
