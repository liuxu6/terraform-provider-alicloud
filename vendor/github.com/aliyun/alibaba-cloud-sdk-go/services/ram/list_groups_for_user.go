package ram

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

// ListGroupsForUser invokes the ram.ListGroupsForUser API synchronously
// api document: https://help.aliyun.com/api/ram/listgroupsforuser.html
func (client *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (response *ListGroupsForUserResponse, err error) {
	response = CreateListGroupsForUserResponse()
	err = client.DoAction(request, response)
	return
}

// ListGroupsForUserWithChan invokes the ram.ListGroupsForUser API asynchronously
// api document: https://help.aliyun.com/api/ram/listgroupsforuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupsForUserWithChan(request *ListGroupsForUserRequest) (<-chan *ListGroupsForUserResponse, <-chan error) {
	responseChan := make(chan *ListGroupsForUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGroupsForUser(request)
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

// ListGroupsForUserWithCallback invokes the ram.ListGroupsForUser API asynchronously
// api document: https://help.aliyun.com/api/ram/listgroupsforuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListGroupsForUserWithCallback(request *ListGroupsForUserRequest, callback func(response *ListGroupsForUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGroupsForUserResponse
		var err error
		defer close(result)
		response, err = client.ListGroupsForUser(request)
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

// ListGroupsForUserRequest is the request struct for api ListGroupsForUser
type ListGroupsForUserRequest struct {
	*requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

// ListGroupsForUserResponse is the response struct for api ListGroupsForUser
type ListGroupsForUserResponse struct {
	*responses.BaseResponse
	RequestId string                    `json:"RequestId" xml:"RequestId"`
	Groups    GroupsInListGroupsForUser `json:"Groups" xml:"Groups"`
}

// CreateListGroupsForUserRequest creates a request to invoke ListGroupsForUser API
func CreateListGroupsForUserRequest() (request *ListGroupsForUserRequest) {
	request = &ListGroupsForUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListGroupsForUser", "Ram", "openAPI")
	return
}

// CreateListGroupsForUserResponse creates a response to parse from ListGroupsForUser response
func CreateListGroupsForUserResponse() (response *ListGroupsForUserResponse) {
	response = &ListGroupsForUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
