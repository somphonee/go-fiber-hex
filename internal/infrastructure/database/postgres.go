package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"github.com/somphonee/go-fiber-hex/internal/core/domain"
	"github.com/somphonee/go-fiber-hex/config"
)

// NewPostgresDB สร้างการเชื่อมต่อใหม่กับฐานข้อมูล PostgreSQL และคืนค่า *gorm.DB
func NewPostgresDB(cfg *config.Config) *gorm.DB {
	// สร้าง DSN (Data Source Name) string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// กำหนดค่า GORM configuration
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // ใช้ชื่อตารางเป็นเอกพจน์
		},
		PrepareStmt:            true, // เตรียม statement สำหรับการ reuse
		SkipDefaultTransaction: true, // ข้าม default transaction เพื่อเพิ่ม performance
	}

	// เชื่อมต่อกับฐานข้อมูล
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ตั้งค่า connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}

	// กำหนดจำนวน connection สูงสุดที่จะเปิดพร้อมกัน
	sqlDB.SetMaxOpenConns(100)
	
	// กำหนดจำนวน connection สูงสุดที่จะ idle ไว้ในพูล
	sqlDB.SetMaxIdleConns(10)
	
	// กำหนดเวลาที่ connection จะ idle ก่อนที่จะถูกปิด
	sqlDB.SetConnMaxIdleTime(time.Hour)
	
	// กำหนดอายุสูงสุดของ connection
	sqlDB.SetConnMaxLifetime(time.Hour * 2)

	// ทดสอบการเชื่อมต่อ
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL database successfully")

	// Auto Migrate - สร้างตารางอัตโนมัติจาก struct models
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrations completed successfully")

	return db
}

// CloseDB ปิดการเชื่อมต่อกับฐานข้อมูล
func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting DB instance: %v", err)
		return
	}
	
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
		return
	}
	
	log.Println("Database connection closed successfully")
}