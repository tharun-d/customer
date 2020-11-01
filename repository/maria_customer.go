package repository

import (
	"customer/common"
	"customer/models"
	"database/sql"
)

const (
	selectUserByMobileNumber = `SELECT id, full_name, mobile, email, type, kyc_level, status, photo_url,verification_status FROM customer
		FROM customerWHERE mobile = ?`
	selectUserByCustomerID = `SELECT id, full_name, mobile, email, type, kyc_level, status, photo_url,verification_status 
		FROM customer WHERE id = ?`
)

func (m *mariaCustomerRepository) GetCustomerData(req models.GetCustomerReq) (models.Customer, error) {
	cus := models.Customer{}

	var (
		query = selectUserByMobileNumber
		param interface{}
	)

	param = req.MobileNumber
	if req.CustomerID > 0 {
		query = selectUserByCustomerID
		param = req.CustomerID
	}

	err := m.Conn.QueryRow(query, param).Scan(&cus.CustomerID, &cus.FullName, &cus.Mobile,
		&cus.Email, &cus.Type, &cus.KycLevel, &cus.Status, &cus.PhotoURL, &cus.VerificationStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Customer{}, common.AppErrorCode(common.CustomerDetailNotFound)
		}
		return models.Customer{}, err
	}

	return cus, nil
}
