package hbase

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

// CreateMultiZoneCluster invokes the hbase.CreateMultiZoneCluster API synchronously
// api document: https://help.aliyun.com/api/hbase/createmultizonecluster.html
func (client *Client) CreateMultiZoneCluster(request *CreateMultiZoneClusterRequest) (response *CreateMultiZoneClusterResponse, err error) {
	response = CreateCreateMultiZoneClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMultiZoneClusterWithChan invokes the hbase.CreateMultiZoneCluster API asynchronously
// api document: https://help.aliyun.com/api/hbase/createmultizonecluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMultiZoneClusterWithChan(request *CreateMultiZoneClusterRequest) (<-chan *CreateMultiZoneClusterResponse, <-chan error) {
	responseChan := make(chan *CreateMultiZoneClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMultiZoneCluster(request)
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

// CreateMultiZoneClusterWithCallback invokes the hbase.CreateMultiZoneCluster API asynchronously
// api document: https://help.aliyun.com/api/hbase/createmultizonecluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMultiZoneClusterWithCallback(request *CreateMultiZoneClusterRequest, callback func(response *CreateMultiZoneClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMultiZoneClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateMultiZoneCluster(request)
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

// CreateMultiZoneClusterRequest is the request struct for api CreateMultiZoneCluster
type CreateMultiZoneClusterRequest struct {
	*requests.RpcRequest
	ArchVersion          string           `position:"Query" name:"ArchVersion"`
	ClusterName          string           `position:"Query" name:"ClusterName"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	LogDiskType          string           `position:"Query" name:"LogDiskType"`
	PrimaryVSwitchId     string           `position:"Query" name:"PrimaryVSwitchId"`
	LogInstanceType      string           `position:"Query" name:"LogInstanceType"`
	AutoRenewPeriod      requests.Integer `position:"Query" name:"AutoRenewPeriod"`
	Period               requests.Integer `position:"Query" name:"Period"`
	LogNodeCount         requests.Integer `position:"Query" name:"LogNodeCount"`
	SecurityIPList       string           `position:"Query" name:"SecurityIPList"`
	PeriodUnit           string           `position:"Query" name:"PeriodUnit"`
	CoreDiskType         string           `position:"Query" name:"CoreDiskType"`
	ArbiterZoneId        string           `position:"Query" name:"ArbiterZoneId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	MultiZoneCombination string           `position:"Query" name:"MultiZoneCombination"`
	PrimaryZoneId        string           `position:"Query" name:"PrimaryZoneId"`
	Engine               string           `position:"Query" name:"Engine"`
	StandbyVSwitchId     string           `position:"Query" name:"StandbyVSwitchId"`
	StandbyZoneId        string           `position:"Query" name:"StandbyZoneId"`
	MasterInstanceType   string           `position:"Query" name:"MasterInstanceType"`
	CoreNodeCount        requests.Integer `position:"Query" name:"CoreNodeCount"`
	LogDiskSize          requests.Integer `position:"Query" name:"LogDiskSize"`
	CoreInstanceType     string           `position:"Query" name:"CoreInstanceType"`
	CoreDiskSize         requests.Integer `position:"Query" name:"CoreDiskSize"`
	VpcId                string           `position:"Query" name:"VpcId"`
	PayType              string           `position:"Query" name:"PayType"`
	ArbiterVSwitchId     string           `position:"Query" name:"ArbiterVSwitchId"`
}

// CreateMultiZoneClusterResponse is the response struct for api CreateMultiZoneCluster
type CreateMultiZoneClusterResponse struct {
	*responses.BaseResponse
	ClusterId string `json:"ClusterId" xml:"ClusterId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateMultiZoneClusterRequest creates a request to invoke CreateMultiZoneCluster API
func CreateCreateMultiZoneClusterRequest() (request *CreateMultiZoneClusterRequest) {
	request = &CreateMultiZoneClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "CreateMultiZoneCluster", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMultiZoneClusterResponse creates a response to parse from CreateMultiZoneCluster response
func CreateCreateMultiZoneClusterResponse() (response *CreateMultiZoneClusterResponse) {
	response = &CreateMultiZoneClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
