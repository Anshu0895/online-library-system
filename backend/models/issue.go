package models

import "time"

type IssueRegistry struct {
	IssueID            uint      `gorm:"primaryKey"  json:"issue_id"`
	ISBN               string    `json:"isbn"`
	ReaderID           uint      `json:"reader_id"`
	IssueApproverID    uint      `json:"issue_approver_id"`
	IssueStatus        string    `json:"issue_status"`
	IssueDate          time.Time `json:"issue_date"`
	ExpectedReturnDate time.Time `json:"expected_return_date"`
	ReturnDate         time.Time `json:"return_date"`
	ReturnApproverID   uint      `json:"return_approver_id"`
}
