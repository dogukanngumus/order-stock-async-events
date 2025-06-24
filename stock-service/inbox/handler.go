package inbox

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	"stock-service/models"
	"stock-service/service"

	"gorm.io/gorm"
)

func generateMessageID(body []byte) string {
	hash := sha256.Sum256(body)
	return hex.EncodeToString(hash[:])
}

func HandleInboxMessage(db *gorm.DB, body []byte) {
	messageID := generateMessageID(body)

	var existing models.InboxMessage
	if err := db.Where("message_id = ?", messageID).First(&existing).Error; err == nil {
		if existing.Processed {
			log.Printf("🔁 Mesaj zaten işlenmiş, atlanıyor. ID: %s", messageID)
			return
		}
	}

	msg := models.InboxMessage{
		MessageID: messageID,
		Payload:   body,
		CreatedAt: time.Now(),
	}

	if err := db.Create(&msg).Error; err != nil {
		log.Printf("❌ InboxMessage veritabanına yazılamadı: %v", err)
		return
	}

	var order models.Order
	if err := json.Unmarshal(body, &order); err != nil {
		log.Printf("❌ Mesaj parse edilemedi: %v", err)
		return
	}

	service.ProcessOrder(order)

	now := time.Now()
	msg.Processed = true
	msg.ProcessedAt = &now

	if err := db.Save(&msg).Error; err != nil {
		log.Printf("❌ Mesaj işlendikten sonra güncellenemedi: %v", err)
	}
}
