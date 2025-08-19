//+build wireinject

package graph

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"backend/internal/application/usecases"
	"backend/internal/domain/repositories"
	infra "backend/internal/infrastructure/repositories"
	"backend/internal/infrastructure/config"
)

func provideDB() *gorm.DB {
	db, err := config.NewDatabase()
	if err != nil {
		panic(err)
	}
	return db
}

var repositorySet = wire.NewSet(
	provideDB,
	infra.NewProductRepository,
	infra.NewCategoryRepository,
	infra.NewStoreRepository,
	wire.Bind(new(repositories.ProductRepository), new(*infra.ProductRepositoryImpl)),
	wire.Bind(new(repositories.CategoryRepository), new(*infra.CategoryRepositoryImpl)),
	wire.Bind(new(repositories.StoreRepository), new(*infra.StoreRepositoryImpl)),
	wire.Bind(new(repositories.ProductPriceHistoryRepository), new(*infra.ProductPriceHistoryRepositoryImpl)),
	wire.Bind(new(repositories.ProductReviewRepository), new(*infra.ProductReviewRepositoryImpl)),
	wire.Bind(new(repositories.ProductVariantRepository), new(*infra.ProductVariantRepositoryImpl)),
	wire.Bind(new(repositories.ProfileDocumentRepository), new(*infra.ProfileDocumentRepositoryImpl)),
	wire.Bind(new(repositories.PromotionRepository), new(*infra.PromotionRepositoryImpl)),
	wire.Bind(new(repositories.ReturnRepository), new(*infra.ReturnRepositoryImpl)),
	wire.Bind(new(repositories.SellerPaymentRepository), new(*infra.SellerPaymentRepositoryImpl)),
	wire.Bind(new(repositories.ShippingMethodRepository), new(*infra.ShippingMethodRepositoryImpl)),
	wire.Bind(new(repositories.StockMovementRepository), new(*infra.StockMovementRepositoryImpl)),
	wire.Bind(new(repositories.StoreCategoryRepository), new(*infra.StoreCategoryRepositoryImpl)),
	wire.Bind(new(repositories.StoreCommissionRepository), new(*infra.StoreCommissionRepositoryImpl)),
	wire.Bind(new(repositories.StorePolicyRepository), new(*infra.StorePolicyRepositoryImpl)),
	wire.Bind(new(repositories.StoreReviewRepository), new(*infra.StoreReviewRepositoryImpl)),
	wire.Bind(new(repositories.StoreShippingMethodRepository), new(*infra.StoreShippingMethodRepositoryImpl)),
	wire.Bind(new(repositories.StoreTransactionRepository), new(*infra.StoreTransactionRepositoryImpl)),
	wire.Bind(new(repositories.UserActivityLogRepository), new(*infra.UserActivityLogRepositoryImpl)),
	wire.Bind(new(repositories.VariantOptionRepository), new(*infra.VariantOptionRepositoryImpl)),
	wire.Bind(new(repositories.WarehouseRepository), new(*infra.WarehouseRepositoryImpl)),
	wire.Bind(new(repositories.WishlistRepository), new(*infra.WishlistRepositoryImpl)),
)

var useCaseSet = wire.NewSet(
	usecases.NewProductUseCase,
	usecases.NewCategoryUseCase,
	usecases.NewStoreUseCase,
	usecases.NewProductPriceHistoryUseCase,
	usecases.NewProductReviewUseCase,
	usecases.NewProductVariantUseCase,
	usecases.NewProfileDocumentUseCase,
	usecases.NewPromotionUseCase,
	usecases.NewReturnUseCase,
	usecases.NewSellerPaymentUseCase,
	usecases.NewShippingMethodUseCase,
	usecases.NewStockMovementUseCase,
	usecases.NewStoreCategoryUseCase,
	usecases.NewStoreCommissionUseCase,
	usecases.NewStorePolicyUseCase,
	usecases.NewStoreReviewUseCase,
	usecases.NewStoreShippingMethodUseCase,
	usecases.NewStoreTransactionUseCase,
	usecases.NewUserActivityLogUseCase,
	usecases.NewVariantOptionUseCase,
	usecases.NewWarehouseUseCase,
	usecases.NewWishlistUseCase,
)

func InitializeResolver() (*Resolver, error) {
	wire.Build(repositorySet, useCaseSet, wire.Struct(new(Resolver), "*"))
	return nil, nil
}
