package drds

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

// RemoveRecycleBinTable invokes the drds.RemoveRecycleBinTable API synchronously
// api document: https://help.aliyun.com/api/drds/removerecyclebintable.html
func (client *Client) RemoveRecycleBinTable(request *RemoveRecycleBinTableRequest) (response *RemoveRecycleBinTableResponse, err error) {
	response = CreateRemoveRecycleBinTableResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveRecycleBinTableWithChan invokes the drds.RemoveRecycleBinTable API asynchronously
// api document: https://help.aliyun.com/api/drds/removerecyclebintable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveRecycleBinTableWithChan(request *RemoveRecycleBinTableRequest) (<-chan *RemoveRecycleBinTableResponse, <-chan error) {
	responseChan := make(chan *RemoveRecycleBinTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveRecycleBinTable(request)
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

// RemoveRecycleBinTableWithCallback invokes the drds.RemoveRecycleBinTable API asynchronously
// api document: https://help.aliyun.com/api/drds/removerecyclebintable.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveRecycleBinTableWithCallback(request *RemoveRecycleBinTableRequest, callback func(response *RemoveRecycleBinTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveRecycleBinTableResponse
		var err error
		defer close(result)
		response, err = client.RemoveRecycleBinTable(request)
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

// RemoveRecycleBinTableRequest is the request struct for api RemoveRecycleBinTable
type RemoveRecycleBinTableRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	TableName      string `position:"Query" name:"TableName"`
}

// RemoveRecycleBinTableResponse is the response struct for api RemoveRecycleBinTable
type RemoveRecycleBinTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateRemoveRecycleBinTableRequest creates a request to invoke RemoveRecycleBinTable API
func CreateRemoveRecycleBinTableRequest() (request *RemoveRecycleBinTableRequest) {
	request = &RemoveRecycleBinTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "RemoveRecycleBinTable", "Drds", "openAPI")
	return
}

// CreateRemoveRecycleBinTableResponse creates a response to parse from RemoveRecycleBinTable response
func CreateRemoveRecycleBinTableResponse() (response *RemoveRecycleBinTableResponse) {
	response = &RemoveRecycleBinTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
