// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type AuditResourceType string

const (
	AuditResourceTypeCloudAccount       AuditResourceType = "cloud_account"
	AuditResourceTypeTag                AuditResourceType = "tag"
	AuditResourceTypePolicy             AuditResourceType = "policy"
	AuditResourceTypeCloudTenant        AuditResourceType = "cloud_tenant"
	AuditResourceTypeUser               AuditResourceType = "user"
	AuditResourceTypeGroup              AuditResourceType = "group"
	AuditResourceTypeApiKey             AuditResourceType = "api_key"
	AuditResourceTypeOrganizationalUnit AuditResourceType = "organizational_unit"
)

func (e *AuditResourceType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuditResourceType(s)
	case string:
		*e = AuditResourceType(s)
	default:
		return fmt.Errorf("unsupported scan type for AuditResourceType: %T", src)
	}
	return nil
}

type NullAuditResourceType struct {
	AuditResourceType AuditResourceType
	Valid             bool // Valid is true if AuditResourceType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAuditResourceType) Scan(value interface{}) error {
	if value == nil {
		ns.AuditResourceType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AuditResourceType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAuditResourceType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.AuditResourceType, nil
}

type CallerType string

const (
	CallerTypeUser   CallerType = "user"
	CallerTypeApiKey CallerType = "api_key"
)

func (e *CallerType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CallerType(s)
	case string:
		*e = CallerType(s)
	default:
		return fmt.Errorf("unsupported scan type for CallerType: %T", src)
	}
	return nil
}

type NullCallerType struct {
	CallerType CallerType
	Valid      bool // Valid is true if CallerType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCallerType) Scan(value interface{}) error {
	if value == nil {
		ns.CallerType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CallerType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCallerType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.CallerType, nil
}

type ApiKey struct {
	ID          uuid.UUID
	Encodedhash string
	Owner       string
	Description sql.NullString
	System      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AuditLog struct {
	ID           uuid.UUID
	CallerID     uuid.UUID
	CallerType   CallerType
	ResourceType AuditResourceType
	ResourceID   uuid.UUID
	Message      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CloudAccount struct {
	ID                uuid.UUID
	Cloud             string
	TenantID          string
	AccountID         string
	Name              string
	Active            bool
	TagsCurrent       pgtype.JSONB
	TagsDesired       pgtype.JSONB
	TagsDriftDetected bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type CloudTenant struct {
	ID        uuid.UUID
	Cloud     string
	TenantID  string
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Group struct {
	ID          uuid.UUID
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GroupApiKey struct {
	GroupID  uuid.UUID
	ApiKeyID uuid.UUID
}

type GroupUser struct {
	GroupID uuid.UUID
	UserID  uuid.UUID
}

type OrganizationalUnit struct {
	ID          uuid.UUID
	ParentID    uuid.NullUUID
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrganizationalUnitsCloudAccount struct {
	CloudAccountID       uuid.UUID
	OrganizationalUnitID uuid.UUID
}

type OrganizationalUnitsGroup struct {
	GroupID              uuid.UUID
	OrganizationalUnitID uuid.UUID
}

type ScimApiKey struct {
	ID        uuid.UUID
	Domain    string
	ApiKeyID  uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StandardTag struct {
	ID          uuid.UUID
	DisplayName string
	Description string
	Key         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID          uuid.UUID
	Username    string
	ExternalID  sql.NullString
	Name        pgtype.JSONB
	DisplayName sql.NullString
	Locale      sql.NullString
	Active      bool
	Emails      pgtype.JSONB
	CreatedAt   time.Time
	UpdatedAt   time.Time
}