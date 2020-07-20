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

// Resource is a nested struct in cms response
type Resource struct {
	TemplateId     string                                      `json:"TemplateId" xml:"TemplateId"`
	Name           string                                      `json:"Name" xml:"Name"`
	Category       string                                      `json:"Category" xml:"Category"`
	Unit           string                                      `json:"Unit" xml:"Unit"`
	NameDesc       string                                      `json:"NameDesc" xml:"NameDesc"`
	Desc           string                                      `json:"Desc" xml:"Desc"`
	Dimensions     string                                      `json:"Dimensions" xml:"Dimensions"`
	RestVersion    string                                      `json:"RestVersion" xml:"RestVersion"`
	RegionId       string                                      `json:"RegionId" xml:"RegionId"`
	InstanceId     string                                      `json:"InstanceId" xml:"InstanceId"`
	NetworkType    string                                      `json:"NetworkType" xml:"NetworkType"`
	Description    string                                      `json:"Description" xml:"Description"`
	Periods        string                                      `json:"Periods" xml:"Periods"`
	Product        string                                      `json:"Product" xml:"Product"`
	InstanceName   string                                      `json:"InstanceName" xml:"InstanceName"`
	Level          string                                      `json:"Level" xml:"Level"`
	Dimension      string                                      `json:"Dimension" xml:"Dimension"`
	Id             int64                                       `json:"Id" xml:"Id"`
	EventType      string                                      `json:"EventType" xml:"EventType"`
	Namespace      string                                      `json:"Namespace" xml:"Namespace"`
	GroupId        int64                                       `json:"GroupId" xml:"GroupId"`
	ServiceId      int64                                       `json:"ServiceId" xml:"ServiceId"`
	MetricName     string                                      `json:"MetricName" xml:"MetricName"`
	StatusDesc     string                                      `json:"StatusDesc" xml:"StatusDesc"`
	Labels         string                                      `json:"Labels" xml:"Labels"`
	Status         string                                      `json:"Status" xml:"Status"`
	Statistics     string                                      `json:"Statistics" xml:"Statistics"`
	Vpc            Vpc                                         `json:"Vpc" xml:"Vpc"`
	Region         Region                                      `json:"Region" xml:"Region"`
	Tags           TagsInDescribeMonitorGroupInstanceAttribute `json:"Tags" xml:"Tags"`
	AlertResults   []Result                                    `json:"AlertResults" xml:"AlertResults"`
	AlertTemplates AlertTemplates                              `json:"AlertTemplates" xml:"AlertTemplates"`
}
