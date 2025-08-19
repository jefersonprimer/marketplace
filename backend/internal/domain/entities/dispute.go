
package entities

import (
    "time"

    "github.com/google/uuid"
)

type DisputeStatus string

const (
    DisputeStatusOpen      DisputeStatus = "open"
    DisputeStatusInReview  DisputeStatus = "in_review"
    DisputeStatusResolved  DisputeStatus = "resolved"
    DisputeStatusRejected  DisputeStatus = "rejected"
)

type Dispute struct {
    ID         uuid.UUID     `json:"id"`
    OrderID    uuid.UUID     `json:"order_id"`
    OpenedBy   uuid.UUID     `json:"opened_by"`
    Reason     string        `json:"reason"`
    Status     DisputeStatus `json:"status"`
    Resolution string        `json:"resolution"`
    CreatedAt  time.Time     `json:"created_at"`
    UpdatedAt  time.Time     `json:"updated_at"`
}
