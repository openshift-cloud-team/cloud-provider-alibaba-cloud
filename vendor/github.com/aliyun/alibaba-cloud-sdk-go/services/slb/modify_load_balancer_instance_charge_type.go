package slb

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

// ModifyLoadBalancerInstanceChargeType invokes the slb.ModifyLoadBalancerInstanceChargeType API synchronously
func (client *Client) ModifyLoadBalancerInstanceChargeType(request *ModifyLoadBalancerInstanceChargeTypeRequest) (response *ModifyLoadBalancerInstanceChargeTypeResponse, err error) {
	response = CreateModifyLoadBalancerInstanceChargeTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLoadBalancerInstanceChargeTypeWithChan invokes the slb.ModifyLoadBalancerInstanceChargeType API asynchronously
func (client *Client) ModifyLoadBalancerInstanceChargeTypeWithChan(request *ModifyLoadBalancerInstanceChargeTypeRequest) (<-chan *ModifyLoadBalancerInstanceChargeTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyLoadBalancerInstanceChargeTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLoadBalancerInstanceChargeType(request)
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

// ModifyLoadBalancerInstanceChargeTypeWithCallback invokes the slb.ModifyLoadBalancerInstanceChargeType API asynchronously
func (client *Client) ModifyLoadBalancerInstanceChargeTypeWithCallback(request *ModifyLoadBalancerInstanceChargeTypeRequest, callback func(response *ModifyLoadBalancerInstanceChargeTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLoadBalancerInstanceChargeTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyLoadBalancerInstanceChargeType(request)
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

// ModifyLoadBalancerInstanceChargeTypeRequest is the request struct for api ModifyLoadBalancerInstanceChargeType
type ModifyLoadBalancerInstanceChargeTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	LoadBalancerSpec     string           `position:"Query" name:"LoadBalancerSpec"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	InternetChargeType   string           `position:"Query" name:"InternetChargeType"`
}

// ModifyLoadBalancerInstanceChargeTypeResponse is the response struct for api ModifyLoadBalancerInstanceChargeType
type ModifyLoadBalancerInstanceChargeTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyLoadBalancerInstanceChargeTypeRequest creates a request to invoke ModifyLoadBalancerInstanceChargeType API
func CreateModifyLoadBalancerInstanceChargeTypeRequest() (request *ModifyLoadBalancerInstanceChargeTypeRequest) {
	request = &ModifyLoadBalancerInstanceChargeTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "ModifyLoadBalancerInstanceChargeType", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLoadBalancerInstanceChargeTypeResponse creates a response to parse from ModifyLoadBalancerInstanceChargeType response
func CreateModifyLoadBalancerInstanceChargeTypeResponse() (response *ModifyLoadBalancerInstanceChargeTypeResponse) {
	response = &ModifyLoadBalancerInstanceChargeTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}