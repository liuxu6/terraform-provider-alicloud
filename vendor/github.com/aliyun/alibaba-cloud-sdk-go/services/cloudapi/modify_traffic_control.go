package cloudapi

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

// ModifyTrafficControl invokes the cloudapi.ModifyTrafficControl API synchronously
// api document: https://help.aliyun.com/api/cloudapi/modifytrafficcontrol.html
func (client *Client) ModifyTrafficControl(request *ModifyTrafficControlRequest) (response *ModifyTrafficControlResponse, err error) {
	response = CreateModifyTrafficControlResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTrafficControlWithChan invokes the cloudapi.ModifyTrafficControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifytrafficcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTrafficControlWithChan(request *ModifyTrafficControlRequest) (<-chan *ModifyTrafficControlResponse, <-chan error) {
	responseChan := make(chan *ModifyTrafficControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTrafficControl(request)
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

// ModifyTrafficControlWithCallback invokes the cloudapi.ModifyTrafficControl API asynchronously
// api document: https://help.aliyun.com/api/cloudapi/modifytrafficcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTrafficControlWithCallback(request *ModifyTrafficControlRequest, callback func(response *ModifyTrafficControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTrafficControlResponse
		var err error
		defer close(result)
		response, err = client.ModifyTrafficControl(request)
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

// ModifyTrafficControlRequest is the request struct for api ModifyTrafficControl
type ModifyTrafficControlRequest struct {
	*requests.RpcRequest
	TrafficControlId   string           `position:"Query" name:"TrafficControlId"`
	TrafficControlName string           `position:"Query" name:"TrafficControlName"`
	Description        string           `position:"Query" name:"Description"`
	UserDefault        requests.Integer `position:"Query" name:"UserDefault"`
	ApiDefault         requests.Integer `position:"Query" name:"ApiDefault"`
	SecurityToken      string           `position:"Query" name:"SecurityToken"`
	TrafficControlUnit string           `position:"Query" name:"TrafficControlUnit"`
	AppDefault         requests.Integer `position:"Query" name:"AppDefault"`
}

// ModifyTrafficControlResponse is the response struct for api ModifyTrafficControl
type ModifyTrafficControlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyTrafficControlRequest creates a request to invoke ModifyTrafficControl API
func CreateModifyTrafficControlRequest() (request *ModifyTrafficControlRequest) {
	request = &ModifyTrafficControlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyTrafficControl", "apigateway", "openAPI")
	return
}

// CreateModifyTrafficControlResponse creates a response to parse from ModifyTrafficControl response
func CreateModifyTrafficControlResponse() (response *ModifyTrafficControlResponse) {
	response = &ModifyTrafficControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
