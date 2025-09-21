package global

import (
	"database/sql"
	"go-backend/pkg/setting"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	// Logger        *logger.LoggerZap
	Rdb           *redis.Client
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
