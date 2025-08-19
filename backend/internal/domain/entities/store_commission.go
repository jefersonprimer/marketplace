
package entities

import (
    "time"

    "github.com/google/uuid"
)

type StoreCommissionType string

const (
    StoreCommissionTypePercentage StoreCommissionType = "percentage"
    StoreCommissionTypeFixed      StoreCommissionType = "fixed"
)

type StoreCommission struct {
    ID              uuid.UUID           `json:"id"`
    StoreID         uuid.UUID           `json:"store_id"`
    CommissionType  StoreCommissionType `json:"commission_type"`
    CommissionValue float64             `json:"commission_value"`
    ValidFrom       time.Time           `json:"valid_from"`
    ValidTo         time.Time           `json:"valid_to"`
}
