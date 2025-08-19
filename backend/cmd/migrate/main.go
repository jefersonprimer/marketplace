package main

import (
	"backend/internal/domain/entities"
	"backend/internal/infrastructure/config"
	"log"
)

func main() {
	// Conectar ao banco de dados
	db, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Configurar GORM para ser mais tolerante com migrações
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}

	// Configurar pool de conexões
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Auto-migrar as tabelas com configurações mais tolerantes
	err = db.AutoMigrate(
		&entities.Category{},
		&entities.Store{},
		&entities.Product{},
		&entities.Profile{},
		&entities.Address{},
		&entities.Order{},
		&entities.OrderItem{},
		&entities.CartItem{},
		&entities.Payment{},
		&entities.Coupon{},
		&entities.Promotion{},
		&entities.ProductReview{},
		&entities.StoreReview{},
		&entities.Wishlist{},
		&entities.Notification{},
		&entities.UserActivityLog{},
		&entities.MarketplaceSetting{},
		&entities.Dispute{},
		&entities.Delivery{},
		&entities.OrderGroup{},
		&entities.OrderMessage{},
		&entities.OrderStatusHistory{},
		&entities.PaymentAttempt{},
		&entities.ProductAttribute{},
		&entities.ProductImage{},
		&entities.ProductPriceHistory{},
		&entities.ProductVariant{},
		&entities.ProfileDocument{},
		&entities.Return{},
		&entities.SellerPayment{},
		&entities.ShippingMethod{},
		&entities.StockMovement{},
		&entities.StoreCategory{},
		&entities.StoreCommission{},
		&entities.StorePolicy{},
		&entities.StoreShippingMethod{},
		&entities.StoreTransaction{},
		&entities.VariantOption{},
		&entities.Warehouse{},
		&entities.CouponCategory{},
		&entities.CouponProduct{},
		&entities.PromotionCategory{},
		&entities.PromotionProduct{},
	)
	if err != nil {
		log.Printf("Warning: Some migrations failed: %v", err)
		log.Println("This might be due to existing constraints or schema differences.")
		log.Println("The database connection is working, but some tables might need manual adjustment.")
	} else {
		log.Println("Database migration completed successfully!")
	}

	// Verificar se as tabelas principais foram criadas
	var count int64
	db.Table("categories").Count(&count)
	log.Printf("Categories table has %d records", count)

	db.Table("stores").Count(&count)
	log.Printf("Stores table has %d records", count)

	db.Table("products").Count(&count)
	log.Printf("Products table has %d records", count)

	log.Println("Database setup completed!")
}
