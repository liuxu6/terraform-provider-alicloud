package cms

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

// DescribeUnhealthyHostAvailability invokes the cms.DescribeUnhealthyHostAvailability API synchronously
// api document: https://help.aliyun.com/api/cms/describeunhealthyhostavailability.html
func (client *Client) DescribeUnhealthyHostAvailability(request *DescribeUnhealthyHostAvailabilityRequest) (response *DescribeUnhealthyHostAvailabilityResponse, err error) {
	response = CreateDescribeUnhealthyHostAvailabilityResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUnhealthyHostAvailabilityWithChan invokes the cms.DescribeUnhealthyHostAvailability API asynchronously
// api document: https://help.aliyun.com/api/cms/describeunhealthyhostavailability.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUnhealthyHostAvailabilityWithChan(request *DescribeUnhealthyHostAvailabilityRequest) (<-chan *DescribeUnhealthyHostAvailabilityResponse, <-chan error) {
	responseChan := make(chan *DescribeUnhealthyHostAvailabilityResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUnhealthyHostAvailability(request)
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

// DescribeUnhealthyHostAvailabilityWithCallback invokes the cms.DescribeUnhealthyHostAvailability API asynchronously
// api document: https://help.aliyun.com/api/cms/describeunhealthyhostavailability.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUnhealthyHostAvailabilityWithCallback(request *DescribeUnhealthyHostAvailabilityRequest, callback func(response *DescribeUnhealthyHostAvailabilityResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUnhealthyHostAvailabilityResponse
		var err error
		defer close(result)
		response, err = client.DescribeUnhealthyHostAvailability(request)
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

// DescribeUnhealthyHostAvailabilityRequest is the request struct for api DescribeUnhealthyHostAvailability
type DescribeUnhealthyHostAvailabilityRequest struct {
	*requests.RpcRequest
	Id *[]string `position:"Query" name:"Id"  type:"Repeated"`
}

// DescribeUnhealthyHostAvailabilityResponse is the response struct for api DescribeUnhealthyHostAvailability
type DescribeUnhealthyHostAvailabilityResponse struct {
	*responses.BaseResponse
	Code          string        `json:"Code" xml:"Code"`
	Message       string        `json:"Message" xml:"Message"`
	Success       bool          `json:"Success" xml:"Success"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	UnhealthyList UnhealthyList `json:"UnhealthyList" xml:"UnhealthyList"`
}

// CreateDescribeUnhealthyHostAvailabilityRequest creates a request to invoke DescribeUnhealthyHostAvailability API
func CreateDescribeUnhealthyHostAvailabilityRequest() (request *DescribeUnhealthyHostAvailabilityRequest) {
	request = &DescribeUnhealthyHostAvailabilityRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeUnhealthyHostAvailability", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUnhealthyHostAvailabilityResponse creates a response to parse from DescribeUnhealthyHostAvailability response
func CreateDescribeUnhealthyHostAvailabilityResponse() (response *DescribeUnhealthyHostAvailabilityResponse) {
	response = &DescribeUnhealthyHostAvailabilityResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
