
package entities

import (
    "time"

    "github.com/google/uuid"
)

type ProfileDocumentType string

const (
    ProfileDocumentTypeCPF          ProfileDocumentType = "cpf"
    ProfileDocumentTypeCNPJ         ProfileDocumentType = "cnpj"
    ProfileDocumentTypeRG           ProfileDocumentType = "rg"
    ProfileDocumentTypeCNH          ProfileDocumentType = "cnh"
    ProfileDocumentTypePassport     ProfileDocumentType = "passport"
    ProfileDocumentTypeAddressProof ProfileDocumentType = "address_proof"
)

type ProfileDocumentStatus string

const (
    ProfileDocumentStatusPending  ProfileDocumentStatus = "pending"
    ProfileDocumentStatusApproved ProfileDocumentStatus = "approved"
    ProfileDocumentStatusRejected ProfileDocumentStatus = "rejected"
)

type ProfileDocument struct {
    ID             uuid.UUID             `json:"id"`
    ProfileID      uuid.UUID             `json:"profile_id"`
    DocumentType   ProfileDocumentType   `json:"document_type"`
    DocumentNumber string                `json:"document_number"`
    DocumentURL    string                `json:"document_url"`
    Status         ProfileDocumentStatus `json:"status"`
    CreatedAt      time.Time             `json:"created_at"`
}
