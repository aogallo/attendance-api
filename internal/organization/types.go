package organization

import (
	"context"
	"time"
)

// Identity identifies an authenticated administrative account, not an attendance Person.
type Identity struct{ Issuer, Subject string }

type Organization struct {
	ID                        string    `json:"id"`
	Name                      string    `json:"name"`
	InitialAdministratorEmail string    `json:"initialAdministratorEmail"`
	AttendanceConfiguration   any       `json:"attendanceConfiguration"`
	CreatedAt                 time.Time `json:"createdAt"`
}

// Repository creates an organization only while the identity holds platform authority.
// Implementations must check authority and commit the insertion in one transaction.
type Repository interface {
	Create(context.Context, Identity, string, string) (Organization, error)
}
