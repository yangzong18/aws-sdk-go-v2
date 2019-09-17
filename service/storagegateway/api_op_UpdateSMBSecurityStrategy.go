// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSMBSecurityStrategyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// Specifies the type of security strategy.
	//
	// ClientSpecified: if you use this option, requests are established based on
	// what is negotiated by the client. This option is recommended when you want
	// to maximize compatibility across different clients in your environment.
	//
	// MandatorySigning: if you use this option, file gateway only allows connections
	// from SMBv2 or SMBv3 clients that have signing enabled. This option works
	// with SMB clients on Microsoft Windows Vista, Windows Server 2008 or newer.
	//
	// MandatoryEncryption: if you use this option, file gateway only allows connections
	// from SMBv3 clients that have encryption enabled. This option is highly recommended
	// for environments that handle sensitive data. This option works with SMB clients
	// on Microsoft Windows 8, Windows Server 2012 or newer.
	//
	// SMBSecurityStrategy is a required field
	SMBSecurityStrategy SMBSecurityStrategy `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UpdateSMBSecurityStrategyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSMBSecurityStrategyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSMBSecurityStrategyInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if len(s.SMBSecurityStrategy) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SMBSecurityStrategy"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSMBSecurityStrategyOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateSMBSecurityStrategyOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateSMBSecurityStrategy = "UpdateSMBSecurityStrategy"

// UpdateSMBSecurityStrategyRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates the SMB security strategy on a file gateway. This action is only
// supported in file gateways.
//
// This API is called Security level in the User Guide.
//
// A higher security level can affect performance of the gateway.
//
//    // Example sending a request using UpdateSMBSecurityStrategyRequest.
//    req := client.UpdateSMBSecurityStrategyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateSMBSecurityStrategy
func (c *Client) UpdateSMBSecurityStrategyRequest(input *UpdateSMBSecurityStrategyInput) UpdateSMBSecurityStrategyRequest {
	op := &aws.Operation{
		Name:       opUpdateSMBSecurityStrategy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSMBSecurityStrategyInput{}
	}

	req := c.newRequest(op, input, &UpdateSMBSecurityStrategyOutput{})
	return UpdateSMBSecurityStrategyRequest{Request: req, Input: input, Copy: c.UpdateSMBSecurityStrategyRequest}
}

// UpdateSMBSecurityStrategyRequest is the request type for the
// UpdateSMBSecurityStrategy API operation.
type UpdateSMBSecurityStrategyRequest struct {
	*aws.Request
	Input *UpdateSMBSecurityStrategyInput
	Copy  func(*UpdateSMBSecurityStrategyInput) UpdateSMBSecurityStrategyRequest
}

// Send marshals and sends the UpdateSMBSecurityStrategy API request.
func (r UpdateSMBSecurityStrategyRequest) Send(ctx context.Context) (*UpdateSMBSecurityStrategyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSMBSecurityStrategyResponse{
		UpdateSMBSecurityStrategyOutput: r.Request.Data.(*UpdateSMBSecurityStrategyOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSMBSecurityStrategyResponse is the response type for the
// UpdateSMBSecurityStrategy API operation.
type UpdateSMBSecurityStrategyResponse struct {
	*UpdateSMBSecurityStrategyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSMBSecurityStrategy request.
func (r *UpdateSMBSecurityStrategyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}