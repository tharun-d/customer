package http

import (
	rh "customer/common"
	"customer/models"
	"customer/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func httpResponseWrite(rw http.ResponseWriter, response interface{}, statusCode int) {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(response)
}

// CustomerHandler  represent the http handler for Customer
type CustomerHandler struct {
	CUsecase service.CustomerService
}

func NewCustomerHandler(customerSrv service.CustomerService) {
	handler := &CustomerHandler{
		CUsecase: customerSrv}

	rh.Route(http.MethodGet, "/r/customer", handler.GetCustomerSearch)

}

func (handler *CustomerHandler) GetCustomerSearch(rw http.ResponseWriter, req *http.Request) {

	mobileNumber := strings.TrimSpace(req.URL.Query().Get("mobile_number"))
	cid := strings.TrimSpace(req.URL.Query().Get("cid"))
	var (
		customerID uint64
		err        error
	)

	if mobileNumber == "" && cid == "" {
		response := &rh.DataResponse{Message: rh.GetMessage(rh.SearchCustomerNoParam)}
		httpResponseWrite(rw, response, http.StatusBadRequest)
		return
	}
	if cid != "" {
		customerID, err = strconv.ParseUint(rh.Param(req, 0), 0, 64)
		if err != nil {
			log.Printf("[delivery:http:handler] GetCustomerSearch Error in converting cid string to integer. Error: %s \n", err)
			response := &rh.DataResponse{Message: "Unable to Parse cid from URL Param"}
			httpResponseWrite(rw, response, http.StatusBadRequest)
			return
		}
	}

	customerDtlRq := models.GetCustomerReq{MobileNumber: mobileNumber, CustomerID: customerID}
	result, err := handler.CUsecase.GetCustomerData(req.Context(), customerDtlRq)

	switch e := err.(type) {
	case nil:
		httpResponseWrite(rw, result, http.StatusOK)
	case *rh.AppError:
		if e.Code == rh.CustomerDetailNotFound {
			response := &rh.DataResponse{Message: e.Message, Data: []int{}}
			httpResponseWrite(rw, response, http.StatusNotFound)
		} else {
			response := &rh.DataResponse{Message: "Internal server error"}
			httpResponseWrite(rw, response, http.StatusInternalServerError)
		}
	default:
		response := &rh.DataResponse{Message: "Internal server error"}
		httpResponseWrite(rw, response, http.StatusInternalServerError)
	}

}
