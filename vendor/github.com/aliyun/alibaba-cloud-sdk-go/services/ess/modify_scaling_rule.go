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

// ModifyScalingRule invokes the ess.ModifyScalingRule API synchronously
func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (response *ModifyScalingRuleResponse, err error) {
	response = CreateModifyScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyScalingRuleWithChan invokes the ess.ModifyScalingRule API asynchronously
func (client *Client) ModifyScalingRuleWithChan(request *ModifyScalingRuleRequest) (<-chan *ModifyScalingRuleResponse, <-chan error) {
	responseChan := make(chan *ModifyScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyScalingRule(request)
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

// ModifyScalingRuleWithCallback invokes the ess.ModifyScalingRule API asynchronously
func (client *Client) ModifyScalingRuleWithCallback(request *ModifyScalingRuleRequest, callback func(response *ModifyScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.ModifyScalingRule(request)
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

// ModifyScalingRuleRequest is the request struct for api ModifyScalingRule
type ModifyScalingRuleRequest struct {
	*requests.RpcRequest
	ResourceOwnerId          requests.Integer                   `position:"Query" name:"ResourceOwnerId"`
	StepAdjustment           *[]ModifyScalingRuleStepAdjustment `position:"Query" name:"StepAdjustment"  type:"Repeated"`
	DisableScaleIn           requests.Boolean                   `position:"Query" name:"DisableScaleIn"`
	ScalingRuleId            string                             `position:"Query" name:"ScalingRuleId"`
	InitialMaxSize           requests.Integer                   `position:"Query" name:"InitialMaxSize"`
	ScalingRuleName          string                             `position:"Query" name:"ScalingRuleName"`
	Cooldown                 requests.Integer                   `position:"Query" name:"Cooldown"`
	PredictiveValueBehavior  string                             `position:"Query" name:"PredictiveValueBehavior"`
	ScaleInEvaluationCount   requests.Integer                   `position:"Query" name:"ScaleInEvaluationCount"`
	MetricName               string                             `position:"Query" name:"MetricName"`
	PredictiveScalingMode    string                             `position:"Query" name:"PredictiveScalingMode"`
	ResourceOwnerAccount     string                             `position:"Query" name:"ResourceOwnerAccount"`
	AdjustmentValue          requests.Integer                   `position:"Query" name:"AdjustmentValue"`
	EstimatedInstanceWarmup  requests.Integer                   `position:"Query" name:"EstimatedInstanceWarmup"`
	OwnerAccount             string                             `position:"Query" name:"OwnerAccount"`
	PredictiveTaskBufferTime requests.Integer                   `position:"Query" name:"PredictiveTaskBufferTime"`
	AdjustmentType           string                             `position:"Query" name:"AdjustmentType"`
	OwnerId                  requests.Integer                   `position:"Query" name:"OwnerId"`
	PredictiveValueBuffer    requests.Integer                   `position:"Query" name:"PredictiveValueBuffer"`
	ScaleOutEvaluationCount  requests.Integer                   `position:"Query" name:"ScaleOutEvaluationCount"`
	MinAdjustmentMagnitude   requests.Integer                   `position:"Query" name:"MinAdjustmentMagnitude"`
	TargetValue              requests.Float                     `position:"Query" name:"TargetValue"`
}

// ModifyScalingRuleStepAdjustment is a repeated param struct in ModifyScalingRuleRequest
type ModifyScalingRuleStepAdjustment struct {
	MetricIntervalUpperBound string `name:"MetricIntervalUpperBound"`
	MetricIntervalLowerBound string `name:"MetricIntervalLowerBound"`
	ScalingAdjustment        string `name:"ScalingAdjustment"`
}

// ModifyScalingRuleResponse is the response struct for api ModifyScalingRule
type ModifyScalingRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyScalingRuleRequest creates a request to invoke ModifyScalingRule API
func CreateModifyScalingRuleRequest() (request *ModifyScalingRuleRequest) {
	request = &ModifyScalingRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "ModifyScalingRule", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyScalingRuleResponse creates a response to parse from ModifyScalingRule response
func CreateModifyScalingRuleResponse() (response *ModifyScalingRuleResponse) {
	response = &ModifyScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}