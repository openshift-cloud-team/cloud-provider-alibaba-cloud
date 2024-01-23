package cas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetCertWarehouseQuota invokes the cas.GetCertWarehouseQuota API synchronously
func (client *Client) GetCertWarehouseQuota(request *GetCertWarehouseQuotaRequest) (response *GetCertWarehouseQuotaResponse, err error) {
	response = CreateGetCertWarehouseQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// GetCertWarehouseQuotaWithChan invokes the cas.GetCertWarehouseQuota API asynchronously
func (client *Client) GetCertWarehouseQuotaWithChan(request *GetCertWarehouseQuotaRequest) (<-chan *GetCertWarehouseQuotaResponse, <-chan error) {
	responseChan := make(chan *GetCertWarehouseQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCertWarehouseQuota(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetCertWarehouseQuotaWithCallback invokes the cas.GetCertWarehouseQuota API asynchronously
func (client *Client) GetCertWarehouseQuotaWithCallback(request *GetCertWarehouseQuotaRequest, callback func(response *GetCertWarehouseQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCertWarehouseQuotaResponse
		var err error
		defer close(result)
		response, err = client.GetCertWarehouseQuota(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetCertWarehouseQuotaRequest is the request struct for api GetCertWarehouseQuota
type GetCertWarehouseQuotaRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
}

// GetCertWarehouseQuotaResponse is the response struct for api GetCertWarehouseQuota
type GetCertWarehouseQuotaResponse struct {
	*responses.BaseResponse
	TotalQuota int64  `json:"TotalQuota" xml:"TotalQuota"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	UseCount   int64  `json:"UseCount" xml:"UseCount"`
}

// CreateGetCertWarehouseQuotaRequest creates a request to invoke GetCertWarehouseQuota API
func CreateGetCertWarehouseQuotaRequest() (request *GetCertWarehouseQuotaRequest) {
	request = &GetCertWarehouseQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cas", "2020-04-07", "GetCertWarehouseQuota", "cas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetCertWarehouseQuotaResponse creates a response to parse from GetCertWarehouseQuota response
func CreateGetCertWarehouseQuotaResponse() (response *GetCertWarehouseQuotaResponse) {
	response = &GetCertWarehouseQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
