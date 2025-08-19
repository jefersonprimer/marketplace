
package entities

import (
    "time"

    "github.com/google/uuid"
)

type SenderRole string

const (
    SenderRoleCustomer SenderRole = "customer"
    SenderRoleSeller   SenderRole = "seller"
    SenderRoleSystem   SenderRole = "system"
)

type OrderMessage struct {
    ID         uuid.UUID  `json:"id"`
    OrderID    uuid.UUID  `json:"order_id"`
    SenderRole SenderRole `json:"sender_role"`
    Message    string     `json:"message"`
    CreatedAt  time.Time  `json:"created_at"`
}
