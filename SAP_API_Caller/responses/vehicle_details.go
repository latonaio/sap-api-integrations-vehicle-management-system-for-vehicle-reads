package responses

type VehicleDetails struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			VMSVehicleUUID                    string `json:"VMSVehicleUUID"`
			Material                          string `json:"Material"`
			VMSDigitalVehicleHubModel         string `json:"VMSDigitalVehicleHubModel"`
			VMSVehicleInternalID              string `json:"VMSVehicleInternalID"`
			ReferenceVendor                   string `json:"ReferenceVendor"`
			VMSVehicleBatch                   string `json:"VMSVehicleBatch"`
			Equipment                         string `json:"Equipment"`
			VMSVehicleActionDocType           string `json:"VMSVehicleActionDocType"`
			Plant                             string `json:"Plant"`
			PlantName                         string `json:"PlantName"`
			StorageLocation                   string `json:"StorageLocation"`
			StorageLocationName               string `json:"StorageLocationName"`
			ValuationType                     string `json:"ValuationType"`
			ConfigurationNumber               string `json:"ConfigurationNumber"`
			VMSVehiclePrimActionControl       string `json:"VMSVehiclePrimActionControl"`
			VMSVehiclePrimaryStatus           string `json:"VMSVehiclePrimaryStatus"`
			VMSVehiclePrimaryStatus_Text      string `json:"VMSVehiclePrimaryStatus_Text"`
			MSVehiclePrimStatusDateTime       string `json:"MSVehiclePrimStatusDateTime"`
			VMSVehicleSecdryActionControl     string `json:"VMSVehicleSecdryActionControl"`
			VMSVehicleSecondaryStatus         string `json:"VMSVehicleSecondaryStatus"`
			VMSVehicleSecondaryStatus_Text    string `json:"VMSVehicleSecondaryStatus_Text"`
			VMSVehicleSecdryStatusDateTime    string `json:"VMSVehicleSecdryStatusDateTime"`
			VMSVehicleConfdDelivDateTime      string `json:"VMSVehicleConfdDelivDateTime"`
			VMSVehicleProductionDateTime      string `json:"VMSVehicleProductionDateTime"`
			VMSVehiclePlndDelivDateTime       string `json:"VMSVehiclePlndDelivDateTime"`
			VMSVehicleIsArchived              string `json:"VMSVehicleIsArchived"`
			VehicleIdentificationNumber       string `json:"VehicleIdentificationNumber"`
			VMSVehicleExternalID              string `json:"VMSVehicleExternalID"`
			VMSVehicleAvailabilityStatus      string `json:"VMSVehicleAvailabilityStatus"`
			VMSVehicleAvailabilityStatus_Text string `json:"VMSVehicleAvailabilityStatus_Text"`
			VMSVehicleIsVisible               string `json:"VMSVehicleIsVisible"`
			VMSVehicleLocation                string `json:"VMSVehicleLocation"`
			VMSVehicleLocation_Text           string `json:"VMSVehicleLocation_Text"`
			VMSVehicleCustomer                string `json:"VMSVehicleCustomer"`
			VMSVehicleEndCustomer             string `json:"VMSVehicleEndCustomer"`
			CreatedByUserName                 string `json:"CreatedByUserName"`
			VMSVehicleUsageStatus             string `json:"VMSVehicleUsageStatus"`
			VMSVehicleUsageStatus_Text        string `json:"VMSVehicleUsageStatus_Text"`
			VMSVehicleSalesCampaign           string `json:"VMSVehicleSalesCampaign"`
			VMSVehicleSalesCampaign_Text      string `json:"VMSVehicleSalesCampaign_Text"`
			VMSVehicleMileageQuantity         string `json:"VMSVehicleMileageQuantity"`
			VMSVehicleMileageUsageUnit        string `json:"VMSVehicleMileageUsageUnit"`
			VMSVehicleRegistrationDate        string `json:"VMSVehicleRegistrationDate"`
			VMSVehicleSearchArea              string `json:"VMSVehicleSearchArea"`
			VMSVehicleSearchArea_Text         string `json:"VMSVehicleSearchArea_Text"`
			VMSVehicleOrderingParty           string `json:"VMSVehicleOrderingParty"`
			VMSVehGrossPrcAmtInVehPrcCrcy     string `json:"VMSVehGrossPrcAmtInVehPrcCrcy"`
			VMSVehiclePriceCrcy               string `json:"VMSVehiclePriceCrcy"`
			VMSVehicleSharingLevel            string `json:"VMSVehicleSharingLevel"`
			VMSVehicleSharingLevel_Text       string `json:"VMSVehicleSharingLevel_Text"`
			VMSVehicleIsUsed                  string `json:"VMSVehicleIsUsed"`
			LogicalSystem                     string `json:"LogicalSystem"`
		} `json:results"`
	} `json:"d"`
}
