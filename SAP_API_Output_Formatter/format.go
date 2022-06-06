package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-vehicle-management-system-for-vehicle-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToVehicleDetails(raw []byte, l *logger.Logger) ([]VehicleDetails, error) {
	pm := &responses.VehicleDetails{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to VehicleDetails. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	vehicleDetails := make([]VehicleDetails, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		vehicleDetails = append(vehicleDetails, VehicleDetails{
			VMSVehicleUUID:                    data.VMSVehicleUUID,
			Material:                          data.Material,
			VMSDigitalVehicleHubModel:         data.VMSDigitalVehicleHubModel,
			VMSVehicleInternalID:              data.VMSVehicleInternalID,
			ReferenceVendor:                   data.ReferenceVendor,
			VMSVehicleBatch:                   data.VMSVehicleBatch,
			Equipment:                         data.Equipment,
			VMSVehicleActionDocType:           data.VMSVehicleActionDocType,
			Plant:                             data.Plant,
			PlantName:                         data.PlantName,
			StorageLocation:                   data.StorageLocation,
			StorageLocationName:               data.StorageLocationName,
			ValuationType:                     data.ValuationType,
			ConfigurationNumber:               data.ConfigurationNumber,
			VMSVehiclePrimActionControl:       data.VMSVehiclePrimActionControl,
			VMSVehiclePrimaryStatus:           data.VMSVehiclePrimaryStatus,
			VMSVehiclePrimaryStatus_Text:      data.VMSVehiclePrimaryStatus_Text,
			MSVehiclePrimStatusDateTime:       data.MSVehiclePrimStatusDateTime,
			VMSVehicleSecdryActionControl:     data.VMSVehicleSecdryActionControl,
			VMSVehicleSecondaryStatus:         data.VMSVehicleSecondaryStatus,
			VMSVehicleSecondaryStatus_Text:    data.VMSVehicleSecondaryStatus_Text,
			VMSVehicleSecdryStatusDateTime:    data.VMSVehicleSecdryStatusDateTime,
			VMSVehicleConfdDelivDateTime:      data.VMSVehicleConfdDelivDateTime,
			VMSVehicleProductionDateTime:      data.VMSVehicleProductionDateTime,
			VMSVehiclePlndDelivDateTime:       data.VMSVehiclePlndDelivDateTime,
			VMSVehicleIsArchived:              data.VMSVehicleIsArchived,
			VehicleIdentificationNumber:       data.VehicleIdentificationNumber,
			VMSVehicleExternalID:              data.VMSVehicleExternalID,
			VMSVehicleAvailabilityStatus:      data.VMSVehicleAvailabilityStatus,
			VMSVehicleAvailabilityStatus_Text: data.VMSVehicleAvailabilityStatus_Text,
			VMSVehicleIsVisible:               data.VMSVehicleIsVisible,
			VMSVehicleLocation:                data.VMSVehicleLocation,
			VMSVehicleLocation_Text:           data.VMSVehicleLocation_Text,
			VMSVehicleCustomer:                data.VMSVehicleCustomer,
			VMSVehicleEndCustomer:             data.VMSVehicleEndCustomer,
			CreatedByUserName:                 data.CreatedByUserName,
			VMSVehicleUsageStatus:             data.VMSVehicleUsageStatus,
			VMSVehicleUsageStatus_Text:        data.VMSVehicleUsageStatus_Text,
			VMSVehicleSalesCampaign:           data.VMSVehicleSalesCampaign,
			VMSVehicleSalesCampaign_Text:      data.VMSVehicleSalesCampaign_Text,
			VMSVehicleMileageQuantity:         data.VMSVehicleMileageQuantity,
			VMSVehicleMileageUsageUnit:        data.VMSVehicleMileageUsageUnit,
			VMSVehicleRegistrationDate:        data.VMSVehicleRegistrationDate,
			VMSVehicleSearchArea:              data.VMSVehicleSearchArea,
			VMSVehicleSearchArea_Text:         data.VMSVehicleSearchArea_Text,
			VMSVehicleOrderingParty:           data.VMSVehicleOrderingParty,
			VMSVehGrossPrcAmtInVehPrcCrcy:     data.VMSVehGrossPrcAmtInVehPrcCrcy,
			VMSVehiclePriceCrcy:               data.VMSVehiclePriceCrcy,
			VMSVehicleSharingLevel:            data.VMSVehicleSharingLevel,
			VMSVehicleSharingLevel_Text:       data.VMSVehicleSharingLevel_Text,
			VMSVehicleIsUsed:                  data.VMSVehicleIsUsed,
			LogicalSystem:                     data.LogicalSystem,
		})
	}

	return vehicleDetails, nil
}
