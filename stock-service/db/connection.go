package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithRetry(dsn string, maxRetries int, retryDelay time.Duration) *gorm.DB {
	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("✅ Veritabanına bağlanıldı")
			return db
		}
		log.Printf("Veritabanına bağlanılamadı: %v. %d. deneme, %v sonra tekrar denenecek.", err, i+1, retryDelay)
		time.Sleep(retryDelay)
	}

	log.Fatalf("Veritabanına bağlanılamadı, program kapanıyor: %v", err)
	return nil
}
