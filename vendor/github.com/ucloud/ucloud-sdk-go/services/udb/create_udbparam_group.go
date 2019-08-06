//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB CreateUDBParamGroup

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CreateUDBParamGroupRequest is request schema for CreateUDBParamGroup action
type CreateUDBParamGroupRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 新配置参数组名称
	GroupName *string `required:"true"`

	// 参数组描述
	Description *string `required:"true"`

	// 源参数组id
	SrcGroupId *int `required:"true"`

	// DB类型id，mysql/mongodb/postgesql按版本细分 1：mysql-5.1，2：mysql-5.5，3：percona-5.5，4：mysql-5.6，5：percona-5.6，6：mysql-5.7，7：percona-5.7，8：mariadb-10.0，9：mongodb-2.4，10：mongodb-2.6，11：mongodb-3.0，12：mongodb-3.2,13：postgresql-9.4，14：postgresql-9.6
	DBTypeId *string `required:"true"`

	// 是否是地域级别的配置文件，默认是false
	RegionFlag *bool `required:"false"`
}

// CreateUDBParamGroupResponse is response schema for CreateUDBParamGroup action
type CreateUDBParamGroupResponse struct {
	response.CommonBase

	// 新配置参数组id
	GroupId int
}

// NewCreateUDBParamGroupRequest will create request of CreateUDBParamGroup action.
func (c *UDBClient) NewCreateUDBParamGroupRequest() *CreateUDBParamGroupRequest {
	req := &CreateUDBParamGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// CreateUDBParamGroup - 从已有配置文件创建新配置文件
func (c *UDBClient) CreateUDBParamGroup(req *CreateUDBParamGroupRequest) (*CreateUDBParamGroupResponse, error) {
	var err error
	var res CreateUDBParamGroupResponse

	err = c.Client.InvokeAction("CreateUDBParamGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
