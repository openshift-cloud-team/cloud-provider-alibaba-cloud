package vpc

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

// CreateVpnAttachment invokes the vpc.CreateVpnAttachment API synchronously
func (client *Client) CreateVpnAttachment(request *CreateVpnAttachmentRequest) (response *CreateVpnAttachmentResponse, err error) {
	response = CreateCreateVpnAttachmentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpnAttachmentWithChan invokes the vpc.CreateVpnAttachment API asynchronously
func (client *Client) CreateVpnAttachmentWithChan(request *CreateVpnAttachmentRequest) (<-chan *CreateVpnAttachmentResponse, <-chan error) {
	responseChan := make(chan *CreateVpnAttachmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpnAttachment(request)
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

// CreateVpnAttachmentWithCallback invokes the vpc.CreateVpnAttachment API asynchronously
func (client *Client) CreateVpnAttachmentWithCallback(request *CreateVpnAttachmentRequest, callback func(response *CreateVpnAttachmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpnAttachmentResponse
		var err error
		defer close(result)
		response, err = client.CreateVpnAttachment(request)
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

// CreateVpnAttachmentRequest is the request struct for api CreateVpnAttachment
type CreateVpnAttachmentRequest struct {
	*requests.RpcRequest
	IkeConfig                    string                     `position:"Query" name:"IkeConfig"`
	AutoConfigRoute              requests.Boolean           `position:"Query" name:"AutoConfigRoute"`
	ResourceOwnerId              requests.Integer           `position:"Query" name:"ResourceOwnerId"`
	CenId                        string                     `position:"Query" name:"CenId"`
	AttachType                   string                     `position:"Query" name:"AttachType"`
	ClientToken                  string                     `position:"Query" name:"ClientToken"`
	IpsecConfig                  string                     `position:"Query" name:"IpsecConfig"`
	BgpConfig                    string                     `position:"Query" name:"BgpConfig"`
	RouteTableAssociationEnabled requests.Boolean           `position:"Query" name:"RouteTableAssociationEnabled"`
	NetworkType                  string                     `position:"Query" name:"NetworkType"`
	HealthCheckConfig            string                     `position:"Query" name:"HealthCheckConfig"`
	CustomerGatewayId            string                     `position:"Query" name:"CustomerGatewayId"`
	LocalSubnet                  string                     `position:"Query" name:"LocalSubnet"`
	RemoteCaCert                 string                     `position:"Query" name:"RemoteCaCert"`
	AutoPublishRouteEnabled      requests.Boolean           `position:"Query" name:"AutoPublishRouteEnabled"`
	RouteTablePropagationEnabled requests.Boolean           `position:"Query" name:"RouteTablePropagationEnabled"`
	RemoteSubnet                 string                     `position:"Query" name:"RemoteSubnet"`
	EffectImmediately            requests.Boolean           `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount         string                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                 string                     `position:"Query" name:"OwnerAccount"`
	EnableDpd                    requests.Boolean           `position:"Query" name:"EnableDpd"`
	Tags                         *[]CreateVpnAttachmentTags `position:"Query" name:"Tags"  type:"Repeated"`
	Name                         string                     `position:"Query" name:"Name"`
	ZoneId                       string                     `position:"Query" name:"ZoneId"`
	EnableNatTraversal           requests.Boolean           `position:"Query" name:"EnableNatTraversal"`
}

// CreateVpnAttachmentTags is a repeated param struct in CreateVpnAttachmentRequest
type CreateVpnAttachmentTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateVpnAttachmentResponse is the response struct for api CreateVpnAttachment
type CreateVpnAttachmentResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	VpnConnectionId string `json:"VpnConnectionId" xml:"VpnConnectionId"`
	Name            string `json:"Name" xml:"Name"`
	CreateTime      int64  `json:"CreateTime" xml:"CreateTime"`
	Code            string `json:"Code" xml:"Code"`
	Success         bool   `json:"Success" xml:"Success"`
	Message         string `json:"Message" xml:"Message"`
}

// CreateCreateVpnAttachmentRequest creates a request to invoke CreateVpnAttachment API
func CreateCreateVpnAttachmentRequest() (request *CreateVpnAttachmentRequest) {
	request = &CreateVpnAttachmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateVpnAttachment", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpnAttachmentResponse creates a response to parse from CreateVpnAttachment response
func CreateCreateVpnAttachmentResponse() (response *CreateVpnAttachmentResponse) {
	response = &CreateVpnAttachmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
