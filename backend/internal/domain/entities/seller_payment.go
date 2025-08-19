
package entities

import (
    "time"

    "github.com/google/uuid"
)

type SellerPaymentStatus string

const (
    SellerPaymentStatusPending    SellerPaymentStatus = "pending"
    SellerPaymentStatusProcessing SellerPaymentStatus = "processing"
    SellerPaymentStatusPaid       SellerPaymentStatus = "paid"
    SellerPaymentStatusFailed     SellerPaymentStatus = "failed"
)

type SellerPayment struct {
    ID          uuid.UUID           `json:"id"`
    OrderID     uuid.UUID           `json:"order_id"`
    StoreID     uuid.UUID           `json:"store_id"`
    Amount      float64             `json:"amount"`
    Status      SellerPaymentStatus `json:"status"`
    PayoutDate  time.Time           `json:"payout_date"`
    CreatedAt   time.Time           `json:"created_at"`
}
