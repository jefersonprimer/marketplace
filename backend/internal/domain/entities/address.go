
package entities

import (
    "time"

    "github.com/google/uuid"
)

type AddressType string

const (
    ShippingAddress AddressType = "shipping"
    BillingAddress  AddressType = "billing"
    HomeAddress     AddressType = "home"
    WorkAddress     AddressType = "work"
)

type Address struct {
    ID           uuid.UUID   `json:"id"`
    UserID       uuid.UUID   `json:"user_id"`
    StoreID      uuid.UUID   `json:"store_id"`
    FullName     string      `json:"full_name"`
    Phone        string      `json:"phone"`
    Street       string      `json:"street"`
    Number       string      `json:"number"`
    Complement   string      `json:"complement"`
    Neighborhood string      `json:"neighborhood"`
    City         string      `json:"city"`
    State        string      `json:"state"`
    ZipCode      string      `json:"zip_code"`
    Country      string      `json:"country"`
    IsDefault    bool        `json:"is_default"`
    Latitude     float64     `json:"latitude"`
    Longitude    float64     `json:"longitude"`
    Type         AddressType `json:"type"`
    CreatedAt    time.Time   `json:"created_at"`
}
