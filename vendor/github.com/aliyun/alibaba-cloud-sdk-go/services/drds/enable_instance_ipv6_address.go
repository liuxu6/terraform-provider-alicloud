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

// EnableInstanceIpv6Address invokes the drds.EnableInstanceIpv6Address API synchronously
// api document: https://help.aliyun.com/api/drds/enableinstanceipv6address.html
func (client *Client) EnableInstanceIpv6Address(request *EnableInstanceIpv6AddressRequest) (response *EnableInstanceIpv6AddressResponse, err error) {
	response = CreateEnableInstanceIpv6AddressResponse()
	err = client.DoAction(request, response)
	return
}

// EnableInstanceIpv6AddressWithChan invokes the drds.EnableInstanceIpv6Address API asynchronously
// api document: https://help.aliyun.com/api/drds/enableinstanceipv6address.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableInstanceIpv6AddressWithChan(request *EnableInstanceIpv6AddressRequest) (<-chan *EnableInstanceIpv6AddressResponse, <-chan error) {
	responseChan := make(chan *EnableInstanceIpv6AddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableInstanceIpv6Address(request)
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

// EnableInstanceIpv6AddressWithCallback invokes the drds.EnableInstanceIpv6Address API asynchronously
// api document: https://help.aliyun.com/api/drds/enableinstanceipv6address.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableInstanceIpv6AddressWithCallback(request *EnableInstanceIpv6AddressRequest, callback func(response *EnableInstanceIpv6AddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableInstanceIpv6AddressResponse
		var err error
		defer close(result)
		response, err = client.EnableInstanceIpv6Address(request)
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

// EnableInstanceIpv6AddressRequest is the request struct for api EnableInstanceIpv6Address
type EnableInstanceIpv6AddressRequest struct {
	*requests.RpcRequest
	DrdsPassword   string `position:"Query" name:"DrdsPassword"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// EnableInstanceIpv6AddressResponse is the response struct for api EnableInstanceIpv6Address
type EnableInstanceIpv6AddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateEnableInstanceIpv6AddressRequest creates a request to invoke EnableInstanceIpv6Address API
func CreateEnableInstanceIpv6AddressRequest() (request *EnableInstanceIpv6AddressRequest) {
	request = &EnableInstanceIpv6AddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "EnableInstanceIpv6Address", "Drds", "openAPI")
	return
}

// CreateEnableInstanceIpv6AddressResponse creates a response to parse from EnableInstanceIpv6Address response
func CreateEnableInstanceIpv6AddressResponse() (response *EnableInstanceIpv6AddressResponse) {
	response = &EnableInstanceIpv6AddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
