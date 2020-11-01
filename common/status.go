package common

const (
	CustomerDetailNotFound = 250
	SearchCustomerNoParam  = 251
)

func GetMessage(id uint32) string {
	keyValueMap := map[uint32]string{
		CustomerDetailNotFound: "Customer Details not found",
		SearchCustomerNoParam:  "mobile_number or cif or saving_account_id is required",
	}
	value, _ := keyValueMap[id]
	return value
}
