package emr

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

// DescribeFlowNodeInstanceContainerLog invokes the emr.DescribeFlowNodeInstanceContainerLog API synchronously
// api document: https://help.aliyun.com/api/emr/describeflownodeinstancecontainerlog.html
func (client *Client) DescribeFlowNodeInstanceContainerLog(request *DescribeFlowNodeInstanceContainerLogRequest) (response *DescribeFlowNodeInstanceContainerLogResponse, err error) {
	response = CreateDescribeFlowNodeInstanceContainerLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowNodeInstanceContainerLogWithChan invokes the emr.DescribeFlowNodeInstanceContainerLog API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflownodeinstancecontainerlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowNodeInstanceContainerLogWithChan(request *DescribeFlowNodeInstanceContainerLogRequest) (<-chan *DescribeFlowNodeInstanceContainerLogResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowNodeInstanceContainerLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowNodeInstanceContainerLog(request)
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

// DescribeFlowNodeInstanceContainerLogWithCallback invokes the emr.DescribeFlowNodeInstanceContainerLog API asynchronously
// api document: https://help.aliyun.com/api/emr/describeflownodeinstancecontainerlog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFlowNodeInstanceContainerLogWithCallback(request *DescribeFlowNodeInstanceContainerLogRequest, callback func(response *DescribeFlowNodeInstanceContainerLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowNodeInstanceContainerLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowNodeInstanceContainerLog(request)
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

// DescribeFlowNodeInstanceContainerLogRequest is the request struct for api DescribeFlowNodeInstanceContainerLog
type DescribeFlowNodeInstanceContainerLogRequest struct {
	*requests.RpcRequest
	Offset         requests.Integer `position:"Query" name:"Offset"`
	LogName        string           `position:"Query" name:"LogName"`
	Length         requests.Integer `position:"Query" name:"Length"`
	NodeInstanceId string           `position:"Query" name:"NodeInstanceId"`
	AppId          string           `position:"Query" name:"AppId"`
	ContainerId    string           `position:"Query" name:"ContainerId"`
	ProjectId      string           `position:"Query" name:"ProjectId"`
}

// DescribeFlowNodeInstanceContainerLogResponse is the response struct for api DescribeFlowNodeInstanceContainerLog
type DescribeFlowNodeInstanceContainerLogResponse struct {
	*responses.BaseResponse
	RequestId string                                          `json:"RequestId" xml:"RequestId"`
	LogEnd    bool                                            `json:"LogEnd" xml:"LogEnd"`
	LogEntrys LogEntrysInDescribeFlowNodeInstanceContainerLog `json:"LogEntrys" xml:"LogEntrys"`
}

// CreateDescribeFlowNodeInstanceContainerLogRequest creates a request to invoke DescribeFlowNodeInstanceContainerLog API
func CreateDescribeFlowNodeInstanceContainerLogRequest() (request *DescribeFlowNodeInstanceContainerLogRequest) {
	request = &DescribeFlowNodeInstanceContainerLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeFlowNodeInstanceContainerLog", "emr", "openAPI")
	return
}

// CreateDescribeFlowNodeInstanceContainerLogResponse creates a response to parse from DescribeFlowNodeInstanceContainerLog response
func CreateDescribeFlowNodeInstanceContainerLogResponse() (response *DescribeFlowNodeInstanceContainerLogResponse) {
	response = &DescribeFlowNodeInstanceContainerLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
