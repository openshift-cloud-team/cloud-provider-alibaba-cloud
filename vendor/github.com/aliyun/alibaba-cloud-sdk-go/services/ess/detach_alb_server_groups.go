package ess

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

// DetachAlbServerGroups invokes the ess.DetachAlbServerGroups API synchronously
func (client *Client) DetachAlbServerGroups(request *DetachAlbServerGroupsRequest) (response *DetachAlbServerGroupsResponse, err error) {
	response = CreateDetachAlbServerGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DetachAlbServerGroupsWithChan invokes the ess.DetachAlbServerGroups API asynchronously
func (client *Client) DetachAlbServerGroupsWithChan(request *DetachAlbServerGroupsRequest) (<-chan *DetachAlbServerGroupsResponse, <-chan error) {
	responseChan := make(chan *DetachAlbServerGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachAlbServerGroups(request)
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

// DetachAlbServerGroupsWithCallback invokes the ess.DetachAlbServerGroups API asynchronously
func (client *Client) DetachAlbServerGroupsWithCallback(request *DetachAlbServerGroupsRequest, callback func(response *DetachAlbServerGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachAlbServerGroupsResponse
		var err error
		defer close(result)
		response, err = client.DetachAlbServerGroups(request)
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

// DetachAlbServerGroupsRequest is the request struct for api DetachAlbServerGroups
type DetachAlbServerGroupsRequest struct {
	*requests.RpcRequest
	ClientToken          string                                 `position:"Query" name:"ClientToken"`
	ScalingGroupId       string                                 `position:"Query" name:"ScalingGroupId"`
	ResourceOwnerAccount string                                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer                       `position:"Query" name:"OwnerId"`
	AlbServerGroup       *[]DetachAlbServerGroupsAlbServerGroup `position:"Query" name:"AlbServerGroup"  type:"Repeated"`
	ForceDetach          requests.Boolean                       `position:"Query" name:"ForceDetach"`
}

// DetachAlbServerGroupsAlbServerGroup is a repeated param struct in DetachAlbServerGroupsRequest
type DetachAlbServerGroupsAlbServerGroup struct {
	AlbServerGroupId string `name:"AlbServerGroupId"`
	Port             string `name:"Port"`
}

// DetachAlbServerGroupsResponse is the response struct for api DetachAlbServerGroups
type DetachAlbServerGroupsResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ScalingActivityId string `json:"ScalingActivityId" xml:"ScalingActivityId"`
}

// CreateDetachAlbServerGroupsRequest creates a request to invoke DetachAlbServerGroups API
func CreateDetachAlbServerGroupsRequest() (request *DetachAlbServerGroupsRequest) {
	request = &DetachAlbServerGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DetachAlbServerGroups", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachAlbServerGroupsResponse creates a response to parse from DetachAlbServerGroups response
func CreateDetachAlbServerGroupsResponse() (response *DetachAlbServerGroupsResponse) {
	response = &DetachAlbServerGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
