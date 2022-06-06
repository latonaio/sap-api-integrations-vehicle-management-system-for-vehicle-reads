package main

import (
	sap_api_caller "sap-api-integrations-vehicle-management-system-for-vehicle-reads/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-vehicle-management-system-for-vehicle-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs/SDC_Vehicle_Details_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"VehicleDetails",
		}
	}

	caller.AsyncGetVehicleDetails(
		inoutSDC.VehicleDetails.VMSVehicleInternalID,
		accepter,
	)
}
