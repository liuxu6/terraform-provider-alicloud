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

// UpdateJvmConfiguration invokes the edas.UpdateJvmConfiguration API synchronously
// api document: https://help.aliyun.com/api/edas/updatejvmconfiguration.html
func (client *Client) UpdateJvmConfiguration(request *UpdateJvmConfigurationRequest) (response *UpdateJvmConfigurationResponse, err error) {
	response = CreateUpdateJvmConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateJvmConfigurationWithChan invokes the edas.UpdateJvmConfiguration API asynchronously
// api document: https://help.aliyun.com/api/edas/updatejvmconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateJvmConfigurationWithChan(request *UpdateJvmConfigurationRequest) (<-chan *UpdateJvmConfigurationResponse, <-chan error) {
	responseChan := make(chan *UpdateJvmConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateJvmConfiguration(request)
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

// UpdateJvmConfigurationWithCallback invokes the edas.UpdateJvmConfiguration API asynchronously
// api document: https://help.aliyun.com/api/edas/updatejvmconfiguration.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateJvmConfigurationWithCallback(request *UpdateJvmConfigurationRequest, callback func(response *UpdateJvmConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateJvmConfigurationResponse
		var err error
		defer close(result)
		response, err = client.UpdateJvmConfiguration(request)
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

// UpdateJvmConfigurationRequest is the request struct for api UpdateJvmConfiguration
type UpdateJvmConfigurationRequest struct {
	*requests.RoaRequest
	MinHeapSize requests.Integer `position:"Query" name:"MinHeapSize"`
	AppId       string           `position:"Query" name:"AppId"`
	GroupId     string           `position:"Query" name:"GroupId"`
	Options     string           `position:"Query" name:"Options"`
	MaxPermSize requests.Integer `position:"Query" name:"MaxPermSize"`
	MaxHeapSize requests.Integer `position:"Query" name:"MaxHeapSize"`
}

// UpdateJvmConfigurationResponse is the response struct for api UpdateJvmConfiguration
type UpdateJvmConfigurationResponse struct {
	*responses.BaseResponse
	Code             int              `json:"Code" xml:"Code"`
	Message          string           `json:"Message" xml:"Message"`
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	JvmConfiguration JvmConfiguration `json:"JvmConfiguration" xml:"JvmConfiguration"`
}

// CreateUpdateJvmConfigurationRequest creates a request to invoke UpdateJvmConfiguration API
func CreateUpdateJvmConfigurationRequest() (request *UpdateJvmConfigurationRequest) {
	request = &UpdateJvmConfigurationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UpdateJvmConfiguration", "/pop/v5/app/app_jvm_config", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateJvmConfigurationResponse creates a response to parse from UpdateJvmConfiguration response
func CreateUpdateJvmConfigurationResponse() (response *UpdateJvmConfigurationResponse) {
	response = &UpdateJvmConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
