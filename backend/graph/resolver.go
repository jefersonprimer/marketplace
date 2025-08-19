package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "backend/internal/application/usecases"

type Resolver struct {
	ProductUseCase             usecases.ProductUseCase
	CategoryUseCase            usecases.CategoryUseCase
	StoreUseCase               usecases.StoreUseCase
	ProductPriceHistoryUseCase *usecases.ProductPriceHistoryUseCase
	ProductReviewUseCase       *usecases.ProductReviewUseCase
	ProductVariantUseCase      *usecases.ProductVariantUseCase
	ProfileDocumentUseCase     *usecases.ProfileDocumentUseCase
	PromotionUseCase           *usecases.PromotionUseCase
	ReturnUseCase              *usecases.ReturnUseCase
	SellerPaymentUseCase       *usecases.SellerPaymentUseCase
	ShippingMethodUseCase      *usecases.ShippingMethodUseCase
	StockMovementUseCase       *usecases.StockMovementUseCase
	StoreCategoryUseCase       *usecases.StoreCategoryUseCase
	StoreCommissionUseCase     *usecases.StoreCommissionUseCase
	StorePolicyUseCase         *usecases.StorePolicyUseCase
	StoreReviewUseCase         *usecases.StoreReviewUseCase
	StoreShippingMethodUseCase *usecases.StoreShippingMethodUseCase
	StoreTransactionUseCase    *usecases.StoreTransactionUseCase
	UserActivityLogUseCase     *usecases.UserActivityLogUseCase
	VariantOptionUseCase       *usecases.VariantOptionUseCase
	WarehouseUseCase           *usecases.WarehouseUseCase
	WishlistUseCase            *usecases.WishlistUseCase
}
