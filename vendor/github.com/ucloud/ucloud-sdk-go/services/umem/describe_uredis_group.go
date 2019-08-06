//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeURedisGroup

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeURedisGroupRequest is request schema for DescribeURedisGroup action
type DescribeURedisGroupRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 组的ID,如果指定则获取描述，否则为列表操 作,需指定Offset/Limit
	GroupId *string `required:"false"`

	// 分页显示的起始偏移, 默认值为0
	Offset *int `required:"false"`

	// 分页显示的条目数, 默认值为20
	Limit *int `required:"false"`
}

// DescribeURedisGroupResponse is response schema for DescribeURedisGroup action
type DescribeURedisGroupResponse struct {
	response.CommonBase

	// 组的总的节点个数
	TotalCount int

	// 组列表 参见 URedisGroupSet
	DataSet []URedisGroupSet
}

// NewDescribeURedisGroupRequest will create request of DescribeURedisGroup action.
func (c *UMemClient) NewDescribeURedisGroupRequest() *DescribeURedisGroupRequest {
	req := &DescribeURedisGroupRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeURedisGroup - 查询主备Redis
func (c *UMemClient) DescribeURedisGroup(req *DescribeURedisGroupRequest) (*DescribeURedisGroupResponse, error) {
	var err error
	var res DescribeURedisGroupResponse

	err = c.Client.InvokeAction("DescribeURedisGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
