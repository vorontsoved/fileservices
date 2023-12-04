//package main
//
//import (
//	"fileservices/internal/config"
//	"fmt"
//	"gopkg.in/gormigrate.v1"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"io/ioutil"
//	"log"
//	"path/filepath"
//)
//
//func main() {
//	cfg := config.MustLoad()
//
//	dbConnStr := fmt.Sprintf(
//		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
//		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Database, cfg.Postgres.Password,
//	)
//
//	// Подключение к базе данных
//	db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		panic("Error connecting to PostgreSQL database")
//	}
//	sqlDB, err := db.DB()
//	if err != nil {
//		panic("Error getting SQL DB from GORM")
//	}
//	defer sqlDB.Close()
//
//	// Создание мигратора
//	m := gormigrate.New(db, gormigrate.DefaultOptions, getMigrations(cfg.MigrationPath))
//
//	// Выполнение миграций
//	if err := m.Migrate(); err != nil {
//		log.Fatal("Error applying migrations: ", err)
//	}
//	log.Println("Migrations applied successfully")
//}
//
//func getMigrations(migrationPath string) []*gormigrate.Migration {
//	files, err := filepath.Glob(filepath.Join(migrationPath, "*.sql"))
//	if err != nil {
//		log.Fatal("Error listing migration files: ", err)
//	}
//
//	var migrations []*gormigrate.Migration
//
//	for _, file := range files {
//		// Чтение содержимого файла миграции
//		content, err := ioutil.ReadFile(file)
//		if err != nil {
//			log.Fatal("Error reading migration file: ", err)
//		}
//
//		// Создание миграции
//		migrations = append(migrations, &gormigrate.Migration{
//			ID: filepath.Base(file),
//			Migrate: func(tx *gorm.DB) error {
//				// Выполнение SQL-запросов из файла
//				if err := tx.Exec(string(content)).Error; err != nil {
//					return err
//				}
//				return nil
//			},
//			Rollback: func(tx *gorm.DB) error {
//				// Здесь выполните откат миграции
//				return nil
//			},
//		})
//	}
//
//	return migrations
//}
