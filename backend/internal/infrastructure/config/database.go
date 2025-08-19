package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (*gorm.DB, error) {
	// Carregar arquivo .env se existir
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Supabase connection string
	supabaseURL := os.Getenv("SUPABASE_URL")
	log.Printf("SUPABASE_URL: %s", supabaseURL)

	if supabaseURL == "" {
		log.Println("SUPABASE_URL não encontrada, usando configuração local do PostgreSQL")
		// Fallback para configuração manual do PostgreSQL
		host := os.Getenv("DB_HOST")
		if host == "" {
			host = "localhost"
		}

		port := os.Getenv("DB_PORT")
		if port == "" {
			port = "5432"
		}

		user := os.Getenv("DB_USER")
		if user == "" {
			user = "postgres"
		}

		password := os.Getenv("DB_PASSWORD")
		if password == "" {
			password = "postgres"
		}

		dbname := os.Getenv("DB_NAME")
		if dbname == "" {
			dbname = "marketplace"
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
			host, user, password, dbname, port)

		log.Printf("Tentando conectar ao PostgreSQL local: %s", dsn)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}

		return db, nil
	}

	// Usar Supabase connection string
	// O Supabase fornece uma URL de conexão completa
	dsn := supabaseURL
	if !strings.Contains(strings.ToLower(dsn), "sslmode=") {
		if strings.Contains(dsn, "?") {
			dsn = dsn + "&sslmode=require"
		} else {
			dsn = dsn + "?sslmode=require"
		}
	}

	log.Printf("Tentando conectar ao Supabase: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Supabase: %w", err)
	}

	log.Println("Conectado com sucesso ao Supabase!")
	return db, nil
}
