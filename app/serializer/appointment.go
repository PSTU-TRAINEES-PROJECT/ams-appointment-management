package serializer

import (
	"time"

	"github.com/lib/pq"
)

type Appointment struct {
	Id                   int           `json:"id" gorm:"primary_key"`
	Name                 string        `json:"name"`
	Subject              string        `json:"subject"`
	Mobile               string        `json:"mobile"`
	Email                string        `json:"email"`
	Age                  int           `json:"age"`
	OwnerMembershipID    int           `json:"owner_membership_id"`
	CustomerMembershipID int           `json:"customer_membership_id"`
	StartDate            time.Time     `json:"start_date"`
	EndDate              time.Time     `json:"end_date"`
	DocumentIDList       pq.Int64Array `gorm:"type:integer[]" json:"document_id_list"` // PostgreSQL integer array
	VisitType            string        `json:"visit_type"`
	ApprovalStatus       string        `json:"approval_status"`
	CreatedByID          int           `json:"created_by_id"`
	DeletedByID          *int          `json:"deleted_by_id"` // Nullable field
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
	AssignedMembershipID int           `json:"assigned_membership_id"`
}
