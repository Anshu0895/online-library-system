package models

type RequestEvent struct {
	ReqID        uint `gorm:"primaryKey"`
	BookID       string
	ReaderID     uint
	RequestDate  string
	ApprovalDate string
	ApproverID   uint
	RequestType  string
}
