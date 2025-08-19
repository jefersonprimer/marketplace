
package entities

import (
    "time"

    "github.com/google/uuid"
)

type PaymentStatus string

const (
    PaymentStatusPending  PaymentStatus = "pending"
    PaymentStatusPaid     PaymentStatus = "paid"
    PaymentStatusFailed   PaymentStatus = "failed"
    PaymentStatusRefunded PaymentStatus = "refunded"
)

type Payment struct {
    ID               uuid.UUID     `json:"id"`
    OrderGroupID     uuid.UUID     `json:"order_group_id"`
    Gateway          string        `json:"gateway"`
    GatewayPaymentID string        `json:"gateway_payment_id"`
    Amount           float64       `json:"amount"`
    Status           PaymentStatus `json:"status"`
    RawResponse      string        `json:"raw_response"`
    CreatedAt        time.Time     `json:"created_at"`
}
