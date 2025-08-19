
package entities

import (
    "time"

    "github.com/google/uuid"
)

type PaymentAttemptStatus string

const (
    PaymentAttemptStatusPending PaymentAttemptStatus = "pending"
    PaymentAttemptStatusSuccess PaymentAttemptStatus = "success"
    PaymentAttemptStatusFailed  PaymentAttemptStatus = "failed"
)

type PaymentAttempt struct {
    ID            uuid.UUID          `json:"id"`
    PaymentID     uuid.UUID          `json:"payment_id"`
    AttemptNumber int                `json:"attempt_number"`
    Status        PaymentAttemptStatus `json:"status"`
    RawResponse   string             `json:"raw_response"`
    CreatedAt     time.Time          `json:"created_at"`
}
