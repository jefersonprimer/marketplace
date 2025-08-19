
package entities

import (
    "time"

    "github.com/google/uuid"
)

type StorePolicyType string

const (
    StorePolicyTypeReturn   StorePolicyType = "return"
    StorePolicyTypeShipping StorePolicyType = "shipping"
    StorePolicyTypePrivacy  StorePolicyType = "privacy"
    StorePolicyTypeTerms    StorePolicyType = "terms"
)

type StorePolicy struct {
    ID        uuid.UUID       `json:"id"`
    StoreID   uuid.UUID       `json:"store_id"`
    PolicyType StorePolicyType `json:"policy_type"`
    Content   string          `json:"content"`
    CreatedAt time.Time       `json:"created_at"`
}
