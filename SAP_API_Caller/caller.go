package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-purchase-scheduling-agreement-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchaseSchedulingAgreement(schedulingAgreement, schedulingAgreementItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(schedulingAgreement)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(schedulingAgreement, schedulingAgreementItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(schedulingAgreement string) {
	headerData, err := c.callPurchaseSchedulingAgreementSrvAPIRequirementHeader("A_SchAgrmtHeader", schedulingAgreement)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerData)
	}

	headerPartnerData, err := c.callToHeaderPartner(headerData[0].ToHeaderPartner)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerPartnerData)
	}

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)
	}

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemScheduleLineData)
	}
	
	itemDeliveryAddressData, err := c.callToItemDeliveryAddress(itemData[0].ToItemDeliveryAddress)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemDeliveryAddressData)
	}
	return
}

func (c *SAPAPICaller) callPurchaseSchedulingAgreementSrvAPIRequirementHeader(api, schedulingAgreement string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_SCHED_AGRMT_PROCESS_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, schedulingAgreement)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToHeaderPartner(url string) ([]sap_api_output_formatter.ToHeaderPartner, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemScheduleLine2(url string) ([]sap_api_output_formatter.ToItemScheduleLine, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemDeliveryAddresss2(url string) (*sap_api_output_formatter.ToItemDeliveryAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemDeliveryAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(schedulingAgreement, schedulingAgreementItem string) {
	itemData, err := c.callPurchaseSchedulingAgreementSrvAPIRequirementItem("A_SchAgrmtItem", schedulingAgreement, schedulingAgreementItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)
	}

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemScheduleLineData)
	}

	itemDeliveryAddressData, err := c.callToItemDeliveryAddress(itemData[0].ToItemDeliveryAddress)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemDeliveryAddressData)
	}
	return
}

func (c *SAPAPICaller) callPurchaseSchedulingAgreementSrvAPIRequirementItem(api, schedulingAgreement, schedulingAgreementItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_SCHED_AGRMT_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, schedulingAgreement, schedulingAgreementItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemScheduleLine(url string) ([]sap_api_output_formatter.ToItemScheduleLine, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemDeliveryAddress(url string) (*sap_api_output_formatter.ToItemDeliveryAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemDeliveryAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, schedulingAgreement string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SchedulingAgreement eq '%s'", schedulingAgreement)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, schedulingAgreement, schedulingAgreementItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SchedulingAgreement eq '%s' and SchedulingAgreementItem eq '%s'", schedulingAgreement, schedulingAgreementItem)
	return params
}
