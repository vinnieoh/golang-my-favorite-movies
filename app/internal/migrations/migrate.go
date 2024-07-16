package migrations

import (
    "log"
    "os"

    "github.com/golang-migrate/migrate/v4"
    migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    gormpg "gorm.io/driver/postgres"
    "gorm.io/gorm"
    _ "github.com/lib/pq" // Importa o driver Postgres
)

func Migrate() {
    dsn := os.Getenv("DB_URL")
    db, err := gorm.Open(gormpg.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get DB instance: %v", err)
    }

    driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
    if err != nil {
        log.Fatalf("Failed to create migration driver: %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://app/db/migrations",
        "postgres", driver)
    if err != nil {
        log.Fatalf("Failed to create migrate instance: %v", err)
    }

    err = m.Up()
    if err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Failed to apply migrations: %v", err)
    }

    log.Println("Database migrations applied successfully")
}
