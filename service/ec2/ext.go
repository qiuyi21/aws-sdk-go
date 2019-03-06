package ec2

import (
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

type CreateVipInput struct {
	_                struct{} `type:"structure"`
	SubnetId         *string  `type:"string" required:"true"`
	VipType          *string  `type:"string"`
	Description      *string  `type:"string"`
	PrivateIpAddress *string  `type:"string"`
	UserScriptUrl    *string  `type:"string"`
	OwnerId          *string  `type:"string"`
}

type CreateVipResult struct {
	_      struct{} `type:"structure"`
	VipId  *string  `locationName:"vipId" type:"string"`
	Return *bool    `locationName:"return" type:"boolean"`
}

type CreateVipOutput struct {
	_      struct{}         `type:"structure"`
	Result *CreateVipResult `locationName:"CreateVipResult" type:"structure"`
}

func (s CreateVipOutput) String() string {
	return awsutil.Prettify(s)
}

func (s CreateVipOutput) GoString() string {
	return s.String()
}

func (c *EC2) CreateVip(input *CreateVipInput) (*CreateVipOutput, error) {
	req, out := c.CreateVipRequest(input)
	return out, req.Send()
}

func (c *EC2) CreateVipRequest(input *CreateVipInput) (req *request.Request, output *CreateVipOutput) {
	op := &request.Operation{
		Name:       "CreateVip",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVipInput{}
	}

	output = &CreateVipOutput{}
	req = c.newRequest(op, input, output)
	return
}

type DeleteVipInput struct {
	_                struct{} `type:"structure"`
	VipId            *string  `type:"string"`
	PrivateIpAddress *string  `type:"string"`
	SubnetId         *string  `type:"string"`
}

type DeleteVipResult struct {
	_      struct{} `type:"structure"`
	Return *bool    `locationName:"return" type:"boolean"`
}

type DeleteVipOutput struct {
	_      struct{}         `type:"structure"`
	Result *DeleteVipResult `locationName:"DeleteVipResult" type:"structure"`
}

func (s DeleteVipOutput) String() string {
	return awsutil.Prettify(s)
}

func (s DeleteVipOutput) GoString() string {
	return s.String()
}

func (c *EC2) DeleteVip(input *DeleteVipInput) (*DeleteVipOutput, error) {
	req, out := c.DeleteVipRequest(input)
	return out, req.Send()
}

func (c *EC2) DeleteVipRequest(input *DeleteVipInput) (req *request.Request, output *DeleteVipOutput) {
	op := &request.Operation{
		Name:       "DeleteVip",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVipInput{}
	}

	output = &DeleteVipOutput{}
	req = c.newRequest(op, input, output)
	return
}

type DescribeVipsInput struct {
	_                struct{} `type:"structure"`
	VipId            *string  `type:"string"`
	PrivateIpAddress *string  `type:"string"`
	SubnetId         *string  `type:"string"`
}

type VipInstance struct {
	_                  struct{} `type:"structure"`
	InstanceId         *string  `locationName:"instanceId" type:"string"`
	NetworkInterfaceId *string  `locationName:"networkInterfaceId" type:"string"`
	IsHeader           *bool    `locationName:"isHeader" type:"boolean"`
	Mac                *string  `locationName:"mac" type:"string"`
	Status             *string  `locationName:"status" type:"string"`
	StateReason        *string  `locationName:"stateReason" type:"string"`
}

type VipDescription struct {
	_                struct{}       `type:"structure"`
	VipId            *string        `locationName:"vipId" type:"string"`
	PrivateIpAddress *string        `locationName:"privateIpAddress" type:"string"`
	PubIpAddress     *string        `locationName:"pubIpAddress" type:"string"`
	SubnetId         *string        `locationName:"subnetId" type:"string"`
	CurrentTarget    *string        `locationName:"currentTarget" type:"string"`
	OwnerId          *string        `locationName:"ownerId" type:"string"`
	CurrentNetifId   *string        `locationName:"currentNetifId" type:"string"`
	VipType          *string        `locationName:"vipType" type:"string"`
	Status           *string        `locationName:"status" type:"string"`
	Description      *string        `locationName:"description" type:"string"`
	Instances        []*VipInstance `locationName:"instanceSet" locationNameList:"instanceInfo" type:"list"`
}

type DescribeVipsResult struct {
	_    struct{}          `type:"structure"`
	Vips []*VipDescription `locationName:"vipSet" locationNameList:"item" type:"list"`
}

type DescribeVipsOutput struct {
	_      struct{}            `type:"structure"`
	Result *DescribeVipsResult `locationName:"DescribeVipsResult" type:"structure"`
}

func (s DescribeVipsOutput) String() string {
	return awsutil.Prettify(s)
}

func (s DescribeVipsOutput) GoString() string {
	return s.String()
}

func (c *EC2) DescribeVips(input *DescribeVipsInput) (*DescribeVipsOutput, error) {
	req, out := c.DescribeVipsRequest(input)
	return out, req.Send()
}

func (c *EC2) DescribeVipsRequest(input *DescribeVipsInput) (req *request.Request, output *DescribeVipsOutput) {
	op := &request.Operation{
		Name:       "DescribeVips",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVipsInput{}
	}

	output = &DescribeVipsOutput{}
	req = c.newRequest(op, input, output)
	return
}

type AddInstanceToVipInput struct {
	_                  struct{} `type:"structure"`
	VipId              *string  `type:"string"`
	PrivateIpAddress   *string  `type:"string"`
	SubnetId           *string  `type:"string"`
	InstanceId         *string  `type:"string" required:"true"`
	NetworkInterfaceId *string  `type:"string"`
	IsHeader           *bool    `type:"boolean"`
}

type AddInstanceToVipResult struct {
	_      struct{} `type:"structure"`
	Return *bool    `locationName:"return" type:"boolean"`
}

type AddInstanceToVipOutput struct {
	_      struct{}                `type:"structure"`
	Result *AddInstanceToVipResult `locationName:"AddInstanceToVipResult" type:"structure"`
}

func (s AddInstanceToVipOutput) String() string {
	return awsutil.Prettify(s)
}

func (s AddInstanceToVipOutput) GoString() string {
	return s.String()
}

func (c *EC2) AddInstanceToVip(input *AddInstanceToVipInput) (*AddInstanceToVipOutput, error) {
	req, out := c.AddInstanceToVipRequest(input)
	return out, req.Send()
}

func (c *EC2) AddInstanceToVipRequest(input *AddInstanceToVipInput) (req *request.Request, output *AddInstanceToVipOutput) {
	op := &request.Operation{
		Name:       "AddInstanceToVip",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddInstanceToVipInput{}
	}

	output = &AddInstanceToVipOutput{}
	req = c.newRequest(op, input, output)
	return
}

type RemoveInstanceFromVipInput struct {
	_                  struct{} `type:"structure"`
	VipId              *string  `type:"string"`
	PrivateIpAddress   *string  `type:"string"`
	SubnetId           *string  `type:"string"`
	InstanceId         *string  `type:"string" required:"true"`
	NetworkInterfaceId *string  `type:"string"`
}

type RemoveInstanceFromVipResult struct {
	_      struct{} `type:"structure"`
	Return *bool    `locationName:"return" type:"boolean"`
}

type RemoveInstanceFromVipOutput struct {
	_      struct{}                     `type:"structure"`
	Result *RemoveInstanceFromVipResult `locationName:"RemoveInstanceFromVipResult" type:"structure"`
}

func (s RemoveInstanceFromVipOutput) String() string {
	return awsutil.Prettify(s)
}

func (s RemoveInstanceFromVipOutput) GoString() string {
	return s.String()
}

func (c *EC2) RemoveInstanceFromVip(input *RemoveInstanceFromVipInput) (*RemoveInstanceFromVipOutput, error) {
	req, out := c.RemoveInstanceFromVipRequest(input)
	return out, req.Send()
}

func (c *EC2) RemoveInstanceFromVipRequest(input *RemoveInstanceFromVipInput) (req *request.Request, output *RemoveInstanceFromVipOutput) {
	op := &request.Operation{
		Name:       "RemoveInstanceFromVip",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveInstanceFromVipInput{}
	}

	output = &RemoveInstanceFromVipOutput{}
	req = c.newRequest(op, input, output)
	return
}

type ExecuteInstanceTaskInput struct {
	_           struct{} `type:"structure"`
	InstanceId  *string  `type:"string" required:"true"`
	CommandType *string  `type:"string" required:"true"`
	CommandName *string  `type:"string" required:"true"`
	CommandArgs *string  `type:"string"`
}

type ExecuteInstanceTaskResult struct {
	_      struct{} `type:"structure"`
	Result *string  `locationName:"result" type:"string"`
}

type ExecuteInstanceTaskOutput struct {
	_      struct{}                   `type:"structure"`
	Result *ExecuteInstanceTaskResult `locationName:"ExecuteInstanceTaskResult" type:"structure"`
}

func (s ExecuteInstanceTaskOutput) String() string {
	return awsutil.Prettify(s)
}

func (s ExecuteInstanceTaskOutput) GoString() string {
	return s.String()
}

func (c *EC2) ExecuteInstanceTask(input *ExecuteInstanceTaskInput) (*ExecuteInstanceTaskOutput, error) {
	req, out := c.ExecuteInstanceTaskRequest(input)
	return out, req.Send()
}

func (c *EC2) ExecuteInstanceTaskRequest(input *ExecuteInstanceTaskInput) (req *request.Request, output *ExecuteInstanceTaskOutput) {
	op := &request.Operation{
		Name:       "ExecuteInstanceTask",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ExecuteInstanceTaskInput{}
	}

	output = &ExecuteInstanceTaskOutput{}
	req = c.newRequest(op, input, output)
	return
}

type DescribeInstanceTypesInput struct {
	_                struct{}  `type:"structure"`
	InstanceType     []*string `type:"list"`
	ImageId          *string   `type:"string"`
	NetworkName      *string   `type:"string"`
	AvailabilityZone *string   `type:"string"`
	VpcId            *string   `type:"string"`
	StorageId        *string   `type:"string"`
}

type InstanceTypeInfo struct {
	_            struct{} `type:"structure"`
	InstanceType *string  `locationName:"instanceType" type:"string"`
	DisplayName  *string  `locationName:"displayName" type:"string"`
	CPU          *int64   `locationName:"cpu" type:"integer"`
	RAM          *int64   `locationName:"ram" type:"integer"`
	Disk         *int64   `locationName:"disk" type:"integer"`
	GPU          *int64   `locationName:"gpu" type:"integer"`
	SSD          *int64   `locationName:"ssd" type:"integer"`
	HDD          *int64   `locationName:"hdd" type:"integer"`
	HBA          *int64   `locationName:"hba" type:"integer"`
	SRIOV        *int64   `locationName:"sriov" type:"integer"`
	Max          *int64   `locationName:"max" type:"integer"`
	Available    *int64   `locationName:"available" type:"integer"`
}

type DescribeInstanceTypesOutput struct {
	_             struct{}            `type:"structure"`
	InstanceTypes []*InstanceTypeInfo `locationName:"instanceTypeInfo" locationNameList:"item" type:"list"`
}

func (s DescribeInstanceTypesOutput) String() string {
	return awsutil.Prettify(s)
}

func (s DescribeInstanceTypesOutput) GoString() string {
	return s.String()
}

func (c *EC2) DescribeInstanceTypes(input *DescribeInstanceTypesInput) (*DescribeInstanceTypesOutput, error) {
	req, out := c.DescribeInstanceTypesRequest(input)
	return out, req.Send()
}

func (c *EC2) DescribeInstanceTypesRequest(input *DescribeInstanceTypesInput) (req *request.Request, output *DescribeInstanceTypesOutput) {
	op := &request.Operation{
		Name:       "DescribeInstanceTypes",
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceTypesInput{}
	}

	output = &DescribeInstanceTypesOutput{}
	req = c.newRequest(op, input, output)
	return
}
