package models

type IssueRegistry struct {
	IssueID            uint `gorm:"primaryKey"`
	ISBN               string
	ReaderID           uint
	IssueApproverID    uint
	IssueStatus        string
	IssueDate          string
	ExpectedReturnDate string
	ReturnDate         string
	ReturnApproverID   uint
}
