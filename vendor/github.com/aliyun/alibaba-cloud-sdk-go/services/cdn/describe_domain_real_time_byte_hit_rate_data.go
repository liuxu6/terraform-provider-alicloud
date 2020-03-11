package cdn

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

// DescribeDomainRealTimeByteHitRateData invokes the cdn.DescribeDomainRealTimeByteHitRateData API synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimebytehitratedata.html
func (client *Client) DescribeDomainRealTimeByteHitRateData(request *DescribeDomainRealTimeByteHitRateDataRequest) (response *DescribeDomainRealTimeByteHitRateDataResponse, err error) {
	response = CreateDescribeDomainRealTimeByteHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainRealTimeByteHitRateDataWithChan invokes the cdn.DescribeDomainRealTimeByteHitRateData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimebytehitratedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRealTimeByteHitRateDataWithChan(request *DescribeDomainRealTimeByteHitRateDataRequest) (<-chan *DescribeDomainRealTimeByteHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRealTimeByteHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRealTimeByteHitRateData(request)
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

// DescribeDomainRealTimeByteHitRateDataWithCallback invokes the cdn.DescribeDomainRealTimeByteHitRateData API asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainrealtimebytehitratedata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainRealTimeByteHitRateDataWithCallback(request *DescribeDomainRealTimeByteHitRateDataRequest, callback func(response *DescribeDomainRealTimeByteHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRealTimeByteHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRealTimeByteHitRateData(request)
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

// DescribeDomainRealTimeByteHitRateDataRequest is the request struct for api DescribeDomainRealTimeByteHitRateData
type DescribeDomainRealTimeByteHitRateDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDomainRealTimeByteHitRateDataResponse is the response struct for api DescribeDomainRealTimeByteHitRateData
type DescribeDomainRealTimeByteHitRateDataResponse struct {
	*responses.BaseResponse
	RequestId string                                      `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDomainRealTimeByteHitRateData `json:"Data" xml:"Data"`
}

// CreateDescribeDomainRealTimeByteHitRateDataRequest creates a request to invoke DescribeDomainRealTimeByteHitRateData API
func CreateDescribeDomainRealTimeByteHitRateDataRequest() (request *DescribeDomainRealTimeByteHitRateDataRequest) {
	request = &DescribeDomainRealTimeByteHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainRealTimeByteHitRateData", "", "")
	return
}

// CreateDescribeDomainRealTimeByteHitRateDataResponse creates a response to parse from DescribeDomainRealTimeByteHitRateData response
func CreateDescribeDomainRealTimeByteHitRateDataResponse() (response *DescribeDomainRealTimeByteHitRateDataResponse) {
	response = &DescribeDomainRealTimeByteHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
