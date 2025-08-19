
package entities

import (
    "time"

    "github.com/google/uuid"
)

type OrderGroupStatus string

const (
    OrderGroupStatusPending   OrderGroupStatus = "pending"
    OrderGroupStatusPaid      OrderGroupStatus = "paid"
    OrderGroupStatusCancelled OrderGroupStatus = "cancelled"
    OrderGroupStatusRefunded  OrderGroupStatus = "refunded"
)

type OrderGroup struct {
    ID              uuid.UUID        `json:"id"`
    CustomerID      uuid.UUID        `json:"customer_id"`
    TotalAmount     float64          `json:"total_amount"`
    Status          OrderGroupStatus `json:"status"`
    PaymentIntentID string           `json:"payment_intent_id"`
    CreatedAt       time.Time        `json:"created_at"`
}
