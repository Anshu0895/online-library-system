package models

type IssueRegistry struct {
	IssueID            uint   `gorm:"primaryKey"  json:"issue_id"`
	ISBN               string `json:"isbn"`
	ReaderID           uint   `json:"reader_id"`
	IssueApproverID    uint   `json:"issue_approver_id"`
	IssueStatus        string `json:"issue_status"`
	IssueDate          string `json:"issue_date"`
	ExpectedReturnDate string `json:"expected_return_date"`
	ReturnDate         string `json:"return_date"`
	ReturnApproverID   uint   `json:"return_approver_id"`
}
