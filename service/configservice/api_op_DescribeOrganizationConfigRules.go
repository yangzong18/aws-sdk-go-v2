// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOrganizationConfigRulesInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `type:"integer"`

	NextToken *string `type:"string"`

	OrganizationConfigRuleNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeOrganizationConfigRulesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeOrganizationConfigRulesOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	OrganizationConfigRules []OrganizationConfigRule `type:"list"`
}

// String returns the string representation
func (s DescribeOrganizationConfigRulesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeOrganizationConfigRules = "DescribeOrganizationConfigRules"

// DescribeOrganizationConfigRulesRequest returns a request value for making API operation for
// AWS Config.
//
//    // Example sending a request using DescribeOrganizationConfigRulesRequest.
//    req := client.DescribeOrganizationConfigRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeOrganizationConfigRules
func (c *Client) DescribeOrganizationConfigRulesRequest(input *DescribeOrganizationConfigRulesInput) DescribeOrganizationConfigRulesRequest {
	op := &aws.Operation{
		Name:       opDescribeOrganizationConfigRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOrganizationConfigRulesInput{}
	}

	req := c.newRequest(op, input, &DescribeOrganizationConfigRulesOutput{})
	return DescribeOrganizationConfigRulesRequest{Request: req, Input: input, Copy: c.DescribeOrganizationConfigRulesRequest}
}

// DescribeOrganizationConfigRulesRequest is the request type for the
// DescribeOrganizationConfigRules API operation.
type DescribeOrganizationConfigRulesRequest struct {
	*aws.Request
	Input *DescribeOrganizationConfigRulesInput
	Copy  func(*DescribeOrganizationConfigRulesInput) DescribeOrganizationConfigRulesRequest
}

// Send marshals and sends the DescribeOrganizationConfigRules API request.
func (r DescribeOrganizationConfigRulesRequest) Send(ctx context.Context) (*DescribeOrganizationConfigRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOrganizationConfigRulesResponse{
		DescribeOrganizationConfigRulesOutput: r.Request.Data.(*DescribeOrganizationConfigRulesOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOrganizationConfigRulesResponse is the response type for the
// DescribeOrganizationConfigRules API operation.
type DescribeOrganizationConfigRulesResponse struct {
	*DescribeOrganizationConfigRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOrganizationConfigRules request.
func (r *DescribeOrganizationConfigRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}