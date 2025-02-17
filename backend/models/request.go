package models

import "time"

type RequestEvent struct {
	ReqID        uint      `gorm:"primaryKey" json:"req_id"`
	BookID       string    `json:"book_id"`
	ReaderID     uint      `json:"reader_id"`
	RequestDate  time.Time `json:"request_date"`
	ApprovalDate time.Time `json:"approval_date"`
	ApproverID   uint      `json:"approver_id"`
	RequestType  string    `json:"request_type"`
}
