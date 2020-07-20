package cbn

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

// DescribeCenRouteMaps invokes the cbn.DescribeCenRouteMaps API synchronously
// api document: https://help.aliyun.com/api/cbn/describecenroutemaps.html
func (client *Client) DescribeCenRouteMaps(request *DescribeCenRouteMapsRequest) (response *DescribeCenRouteMapsResponse, err error) {
	response = CreateDescribeCenRouteMapsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCenRouteMapsWithChan invokes the cbn.DescribeCenRouteMaps API asynchronously
// api document: https://help.aliyun.com/api/cbn/describecenroutemaps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCenRouteMapsWithChan(request *DescribeCenRouteMapsRequest) (<-chan *DescribeCenRouteMapsResponse, <-chan error) {
	responseChan := make(chan *DescribeCenRouteMapsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCenRouteMaps(request)
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

// DescribeCenRouteMapsWithCallback invokes the cbn.DescribeCenRouteMaps API asynchronously
// api document: https://help.aliyun.com/api/cbn/describecenroutemaps.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCenRouteMapsWithCallback(request *DescribeCenRouteMapsRequest, callback func(response *DescribeCenRouteMapsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCenRouteMapsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCenRouteMaps(request)
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

// DescribeCenRouteMapsRequest is the request struct for api DescribeCenRouteMaps
type DescribeCenRouteMapsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer              `position:"Query" name:"ResourceOwnerId"`
	CenId                string                        `position:"Query" name:"CenId"`
	PageNumber           requests.Integer              `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer              `position:"Query" name:"PageSize"`
	TransmitDirection    string                        `position:"Query" name:"TransmitDirection"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                        `position:"Query" name:"OwnerAccount"`
	RouteMapId           string                        `position:"Query" name:"RouteMapId"`
	OwnerId              requests.Integer              `position:"Query" name:"OwnerId"`
	Filter               *[]DescribeCenRouteMapsFilter `position:"Query" name:"Filter"  type:"Repeated"`
	CenRegionId          string                        `position:"Query" name:"CenRegionId"`
}

// DescribeCenRouteMapsFilter is a repeated param struct in DescribeCenRouteMapsRequest
type DescribeCenRouteMapsFilter struct {
	Value *[]string `name:"Value" type:"Repeated"`
	Key   string    `name:"Key"`
}

// DescribeCenRouteMapsResponse is the response struct for api DescribeCenRouteMaps
type DescribeCenRouteMapsResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	RouteMaps  RouteMaps `json:"RouteMaps" xml:"RouteMaps"`
}

// CreateDescribeCenRouteMapsRequest creates a request to invoke DescribeCenRouteMaps API
func CreateDescribeCenRouteMapsRequest() (request *DescribeCenRouteMapsRequest) {
	request = &DescribeCenRouteMapsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenRouteMaps", "Cbn", "openAPI")
	return
}

// CreateDescribeCenRouteMapsResponse creates a response to parse from DescribeCenRouteMaps response
func CreateDescribeCenRouteMapsResponse() (response *DescribeCenRouteMapsResponse) {
	response = &DescribeCenRouteMapsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
