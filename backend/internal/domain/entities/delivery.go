
package entities

import (
    "time"

    "github.com/google/uuid"
)

type DeliveryStatus string

const (
    DeliveryStatusPending   DeliveryStatus = "pending"
    DeliveryStatusShipped   DeliveryStatus = "shipped"
    DeliveryStatusDelivered DeliveryStatus = "delivered"
    DeliveryStatusCancelled DeliveryStatus = "cancelled"
)

type Delivery struct {
    ID                  uuid.UUID      `json:"id"`
    OrderID             uuid.UUID      `json:"order_id"`
    DeliveryPartnerID   uuid.UUID      `json:"delivery_partner_id"`
    Status              DeliveryStatus `json:"status"`
    TrackingCode        string         `json:"tracking_code"`
    EstimatedDeliveryAt time.Time      `json:"estimated_delivery_at"`
    DeliveredAt         time.Time      `json:"delivered_at"`
    CreatedAt           time.Time      `json:"created_at"`
}
