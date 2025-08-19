
package entities

import (
    "time"

    "github.com/google/uuid"
)

type ProductStatus string

const (
    ProductStatusActive      ProductStatus = "active"
    ProductStatusInactive    ProductStatus = "inactive"
    ProductStatusOutOfStock ProductStatus = "out_of_stock"
)

type Product struct {
    ID          uuid.UUID     `json:"id"`
    StoreID     uuid.UUID     `json:"store_id"`
    Name        string        `json:"name"`
    Slug        string        `json:"slug"`
    Description string        `json:"description"`
    Price       float64       `json:"price"`
    Stock       int           `json:"stock"`
    Status      ProductStatus `json:"status"`
    ImageURL    string        `json:"image_url"`
    CategoryID  uuid.UUID     `json:"category_id"`
    SKU         string        `json:"sku"`
    GTIN        string        `json:"gtin"`
    Weight      float64       `json:"weight"`
    Length      float64       `json:"length"`
    Width       float64       `json:"width"`
    Height      float64       `json:"height"`
    CreatedAt   time.Time     `json:"created_at"`
    UpdatedAt   time.Time     `json:"updated_at"`
}
