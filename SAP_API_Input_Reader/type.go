package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	VehicleDetails struct {
		VMSVehicleUUID                   string `json:"VMSVehicleUUID"`
		Material                         string `json:"Material"`
		VMSDigitalVehicleHubModel        string `json:"VMSDigitalVehicleHubModel"`
		VMSVehicleInternalID             string `json:"VMSVehicleInternalID"`
		ReferenceVendor                  string `json:"ReferenceVendor"`
		VMSVehicleBatch                  string `json:"VMSVehicleBatch"`
		Equipment                        string `json:"Equipment"`
		VMSVehicleActionDocType          string `json:"VMSVehicleActionDocType"`
		Plant                            string `json:"Plant"`
		PlantName                        string `json:"PlantName"`
		StorageLocation                  string `json:"StorageLocation"`
		StorageLocationName              string `json:"StorageLocationName"`
		ValuationType                    string `json:"ValuationType"`
		ConfigurationNumber              string `json:"ConfigurationNumber"`
		VMSVehiclePrimActionControl      string `json:"VMSVehiclePrimActionControl"`
		VMSVehiclePrimaryStatus          string `json:"VMSVehiclePrimaryStatus"`
		VMSVehiclePrimaryStatusText      string `json:"VMSVehiclePrimaryStatus_Text"`
		MSVehiclePrimStatusDateTime      string `json:"MSVehiclePrimStatusDateTime"`
		VMSVehicleSecdryActionControl    string `json:"VMSVehicleSecdryActionControl"`
		VMSVehicleSecondaryStatus        string `json:"VMSVehicleSecondaryStatus"`
		VMSVehicleSecondaryStatusText    string `json:"VMSVehicleSecondaryStatus_Text"`
		VMSVehicleSecdryStatusDateTime   string `json:"VMSVehicleSecdryStatusDateTime"`
		VMSVehicleConfdDelivDateTime     string `json:"VMSVehicleConfdDelivDateTime"`
		VMSVehicleProductionDateTime     string `json:"VMSVehicleProductionDateTime"`
		VMSVehiclePlndDelivDateTime      string `json:"VMSVehiclePlndDelivDateTime"`
		VMSVehicleIsArchived             string `json:"VMSVehicleIsArchived"`
		VehicleIdentificationNumber      string `json:"VehicleIdentificationNumber"`
		VMSVehicleExternalID             string `json:"VMSVehicleExternalID"`
		VMSVehicleAvailabilityStatus     string `json:"VMSVehicleAvailabilityStatus"`
		VMSVehicleAvailabilityStatusText string `json:"VMSVehicleAvailabilityStatus_Text"`
		VMSVehicleIsVisible              string `json:"VMSVehicleIsVisible"`
		VMSVehicleLocation               string `json:"VMSVehicleLocation"`
		VMSVehicleLocationText           string `json:"VMSVehicleLocation_Text"`
		VMSVehicleCustomer               string `json:"VMSVehicleCustomer"`
		VMSVehicleEndCustomer            string `json:"VMSVehicleEndCustomer"`
		CreatedByUserName                string `json:"CreatedByUserName"`
		VMSVehicleUsageStatus            string `json:"VMSVehicleUsageStatus"`
		VMSVehicleUsageStatusText        string `json:"VMSVehicleUsageStatus_Text"`
		VMSVehicleSalesCampaign          string `json:"VMSVehicleSalesCampaign"`
		VMSVehicleSalesCampaignText      string `json:"VMSVehicleSalesCampaign_Text"`
		VMSVehicleMileageQuantity        string `json:"VMSVehicleMileageQuantity"`
		VMSVehicleMileageUsageUnit       string `json:"VMSVehicleMileageUsageUnit"`
		VMSVehicleRegistrationDate       string `json:"VMSVehicleRegistrationDate"`
		VMSVehicleSearchArea             string `json:"VMSVehicleSearchArea"`
		VMSVehicleSearchAreaText         string `json:"VMSVehicleSearchArea_Text"`
		VMSVehicleOrderingParty          string `json:"VMSVehicleOrderingParty"`
		VMSVehGrossPrcAmtInVehPrcCrcy    string `json:"VMSVehGrossPrcAmtInVehPrcCrcy"`
		VMSVehiclePriceCrcy              string `json:"VMSVehiclePriceCrcy"`
		VMSVehicleSharingLevel           string `json:"VMSVehicleSharingLevel"`
		VMSVehicleSharingLevelText       string `json:"VMSVehicleSharingLevel_Text"`
		VMSVehicleIsUsed                 string `json:"VMSVehicleIsUsed"`
		LogicalSystem                    string `json:"LogicalSystem"`
	} `json:"VehicleDetails"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	VehicleDatailesCode string   `json:"vehicle_datailes_code"`
	Deleted             bool     `json:"deleted"`
}
