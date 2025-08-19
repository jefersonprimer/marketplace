
package entities

import (
    "time"

    "github.com/google/uuid"
)

type ReturnStatus string

const (
    ReturnStatusRequested ReturnStatus = "requested"
    ReturnStatusApproved  ReturnStatus = "approved"
    ReturnStatusRejected  ReturnStatus = "rejected"
    ReturnStatusReturned  ReturnStatus = "returned"
    ReturnStatusRefunded  ReturnStatus = "refunded"
)

type Return struct {
    ID          uuid.UUID    `json:"id"`
    OrderID     uuid.UUID    `json:"order_id"`
    OrderItemID uuid.UUID    `json:"order_item_id"`
    Reason      string       `json:"reason"`
    Status      ReturnStatus `json:"status"`
    RequestedAt time.Time    `json:"requested_at"`
    ProcessedAt time.Time    `json:"processed_at"`
}
