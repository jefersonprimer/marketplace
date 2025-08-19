
package entities

import (
    "time"

    "github.com/google/uuid"
)

type OrderStatusHistory struct {
    ID        uuid.UUID   `json:"id"`
    OrderID   uuid.UUID   `json:"order_id"`
    Status    OrderStatus `json:"status"`
    ChangedAt time.Time   `json:"changed_at"`
}
