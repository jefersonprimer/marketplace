
package entities

import (
    "time"

    "github.com/google/uuid"
)

type StoreTransactionType string

const (
    StoreTransactionTypeCredit StoreTransactionType = "credit"
    StoreTransactionTypeDebit  StoreTransactionType = "debit"
)

type StoreTransaction struct {
    ID             uuid.UUID            `json:"id"`
    StoreID        uuid.UUID            `json:"store_id"`
    Type           StoreTransactionType `json:"type"`
    Amount         float64              `json:"amount"`
    Description    string               `json:"description"`
    RelatedOrderID uuid.UUID            `json:"related_order_id"`
    CreatedAt      time.Time            `json:"created_at"`
}
