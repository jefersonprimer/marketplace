
package entities

import (
    "time"

    "github.com/google/uuid"
)

type NotificationType string

const (
    NotificationTypeEmail    NotificationType = "email"
    NotificationTypeWhatsapp NotificationType = "whatsapp"
    NotificationTypeSystem   NotificationType = "system"
)

type NotificationStatus string

const (
    NotificationStatusPending NotificationStatus = "pending"
    NotificationStatusSent    NotificationStatus = "sent"
    NotificationStatusFailed  NotificationStatus = "failed"
)

type Notification struct {
    ID        uuid.UUID          `json:"id"`
    UserID    uuid.UUID          `json:"user_id"`
    Type      NotificationType   `json:"type"`
    Content   string             `json:"content"`
    Status    NotificationStatus `json:"status"`
    CreatedAt time.Time          `json:"created_at"`
}
