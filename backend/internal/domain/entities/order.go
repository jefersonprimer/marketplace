
package entities

import (
    "time"

    "github.com/google/uuid"
)

type OrderStatus string

const (
    OrderStatusPending   OrderStatus = "pending"
    OrderStatusPaid      OrderStatus = "paid"
    OrderStatusShipped   OrderStatus = "shipped"
    OrderStatusDelivered OrderStatus = "delivered"
    OrderStatusCancelled OrderStatus = "cancelled"
)

type Order struct {
    ID                uuid.UUID   `json:"id"`
    StoreID           uuid.UUID   `json:"store_id"`
    TotalAmount       float64     `json:"total_amount"`
    Status            OrderStatus `json:"status"`
    ShippingAddressID uuid.UUID   `json:"shipping_address_id"`
    TransactionID     uuid.UUID   `json:"transaction_id"`
    OrderGroupID      uuid.UUID   `json:"order_group_id"`
    CommissionFee     float64     `json:"commission_fee"`
    NetAmount         float64     `json:"net_amount"`
    CreatedAt         time.Time   `json:"created_at"`
    UpdatedAt         time.Time   `json:"updated_at"`
}
