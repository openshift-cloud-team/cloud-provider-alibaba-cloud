package alb

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

// DisableLoadBalancerIpv6Internet invokes the alb.DisableLoadBalancerIpv6Internet API synchronously
func (client *Client) DisableLoadBalancerIpv6Internet(request *DisableLoadBalancerIpv6InternetRequest) (response *DisableLoadBalancerIpv6InternetResponse, err error) {
	response = CreateDisableLoadBalancerIpv6InternetResponse()
	err = client.DoAction(request, response)
	return
}

// DisableLoadBalancerIpv6InternetWithChan invokes the alb.DisableLoadBalancerIpv6Internet API asynchronously
func (client *Client) DisableLoadBalancerIpv6InternetWithChan(request *DisableLoadBalancerIpv6InternetRequest) (<-chan *DisableLoadBalancerIpv6InternetResponse, <-chan error) {
	responseChan := make(chan *DisableLoadBalancerIpv6InternetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableLoadBalancerIpv6Internet(request)
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

// DisableLoadBalancerIpv6InternetWithCallback invokes the alb.DisableLoadBalancerIpv6Internet API asynchronously
func (client *Client) DisableLoadBalancerIpv6InternetWithCallback(request *DisableLoadBalancerIpv6InternetRequest, callback func(response *DisableLoadBalancerIpv6InternetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableLoadBalancerIpv6InternetResponse
		var err error
		defer close(result)
		response, err = client.DisableLoadBalancerIpv6Internet(request)
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

// DisableLoadBalancerIpv6InternetRequest is the request struct for api DisableLoadBalancerIpv6Internet
type DisableLoadBalancerIpv6InternetRequest struct {
	*requests.RpcRequest
	ClientToken    string           `position:"Query" name:"ClientToken"`
	DryRun         requests.Boolean `position:"Query" name:"DryRun"`
	LoadBalancerId string           `position:"Query" name:"LoadBalancerId"`
}

// DisableLoadBalancerIpv6InternetResponse is the response struct for api DisableLoadBalancerIpv6Internet
type DisableLoadBalancerIpv6InternetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateDisableLoadBalancerIpv6InternetRequest creates a request to invoke DisableLoadBalancerIpv6Internet API
func CreateDisableLoadBalancerIpv6InternetRequest() (request *DisableLoadBalancerIpv6InternetRequest) {
	request = &DisableLoadBalancerIpv6InternetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "DisableLoadBalancerIpv6Internet", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableLoadBalancerIpv6InternetResponse creates a response to parse from DisableLoadBalancerIpv6Internet response
func CreateDisableLoadBalancerIpv6InternetResponse() (response *DisableLoadBalancerIpv6InternetResponse) {
	response = &DisableLoadBalancerIpv6InternetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
