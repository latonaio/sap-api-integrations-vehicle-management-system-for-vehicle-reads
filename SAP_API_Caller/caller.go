package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-vehicle-management-system-for-vehicle-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetVehicleDetails(vMSVehicleInternalID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "VehicleDetails":
			func() {
				c.VehicleDetails(vMSVehicleInternalID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) VehicleDetails(vMSVehicleInternalID string) {
	data, err := c.callVMSForVehicleSrvAPIRequirementVehicleDetails("A_VMSVehicle", vMSVehicleInternalID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callVMSForVehicleSrvAPIRequirementVehicleDetails(api, vMSVehicleInternalID string) ([]sap_api_output_formatter.VehicleDetails, error) {
	url := strings.Join([]string{c.baseURL, "API_VMSVEHICLE", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithVehicleDetails(req, vMSVehicleInternalID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToVehicleDetails(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithVehicleDetails(req *http.Request, vMSVehicleInternalID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("VMSVehicleInternalID eq '%s'", vMSVehicleInternalID))
	req.URL.RawQuery = params.Encode()
}
