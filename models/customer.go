package models

import "time"

//Customer is customer data type
type Customer struct {
	CustomerID         uint64
	Name               string
	Mobile             string
	Email              string
	Cif                string
	IDNo               string
	FullName           string
	DOB                time.Time
	PlaceOfBirth       string
	Gender             int32
	BloodType          string
	Address            string
	Province           uint64
	City               uint64
	District           uint64
	SubDistrict        uint64
	Religion           uint64
	MarritalStatus     string
	Employee           string
	PhotoURL           string
	IconURL            string
	Status             int64
	Type               int32
	KycLevel           int32
	VerificationStatus int32
	CreatedBy          uint64
	CreatedAt          time.Time
	UpdatedBy          uint64
	UpdatedAt          time.Time
	RequestedDate      *time.Time
	EktpStatus         int32
	KtpFile            string
	KtpFolder          string
	SelfieFile         string
	SelfieFolder       string
}

type GetCustomerReq struct {
	CustomerID   uint64
	MobileNumber string
}
