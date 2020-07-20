package edas

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

// DisableFlowControl invokes the edas.DisableFlowControl API synchronously
// api document: https://help.aliyun.com/api/edas/disableflowcontrol.html
func (client *Client) DisableFlowControl(request *DisableFlowControlRequest) (response *DisableFlowControlResponse, err error) {
	response = CreateDisableFlowControlResponse()
	err = client.DoAction(request, response)
	return
}

// DisableFlowControlWithChan invokes the edas.DisableFlowControl API asynchronously
// api document: https://help.aliyun.com/api/edas/disableflowcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableFlowControlWithChan(request *DisableFlowControlRequest) (<-chan *DisableFlowControlResponse, <-chan error) {
	responseChan := make(chan *DisableFlowControlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableFlowControl(request)
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

// DisableFlowControlWithCallback invokes the edas.DisableFlowControl API asynchronously
// api document: https://help.aliyun.com/api/edas/disableflowcontrol.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableFlowControlWithCallback(request *DisableFlowControlRequest, callback func(response *DisableFlowControlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableFlowControlResponse
		var err error
		defer close(result)
		response, err = client.DisableFlowControl(request)
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

// DisableFlowControlRequest is the request struct for api DisableFlowControl
type DisableFlowControlRequest struct {
	*requests.RoaRequest
	AppId  string `position:"Query" name:"AppId"`
	RuleId string `position:"Query" name:"RuleId"`
}

// DisableFlowControlResponse is the response struct for api DisableFlowControl
type DisableFlowControlResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableFlowControlRequest creates a request to invoke DisableFlowControl API
func CreateDisableFlowControlRequest() (request *DisableFlowControlRequest) {
	request = &DisableFlowControlRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DisableFlowControl", "/pop/v5/flowcontrol/disable", "Edas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateDisableFlowControlResponse creates a response to parse from DisableFlowControl response
func CreateDisableFlowControlResponse() (response *DisableFlowControlResponse) {
	response = &DisableFlowControlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
