// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetComplianceDetailsByConfigRuleRequest
type GetComplianceDetailsByConfigRuleInput struct {
	_ struct{} `type:"structure"`

	// Filters the results by compliance.
	//
	// The allowed values are COMPLIANT, NON_COMPLIANT, and NOT_APPLICABLE.
	ComplianceTypes []ComplianceType `type:"list"`

	// The name of the AWS Config rule for which you want compliance information.
	//
	// ConfigRuleName is a required field
	ConfigRuleName *string `min:"1" type:"string" required:"true"`

	// The maximum number of evaluation results returned on each page. The default
	// is 10. You cannot specify a number greater than 100. If you specify 0, AWS
	// Config uses the default.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetComplianceDetailsByConfigRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetComplianceDetailsByConfigRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetComplianceDetailsByConfigRuleInput"}

	if s.ConfigRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigRuleName"))
	}
	if s.ConfigRuleName != nil && len(*s.ConfigRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigRuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetComplianceDetailsByConfigRuleResponse
type GetComplianceDetailsByConfigRuleOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the AWS resource complies with the specified AWS Config
	// rule.
	EvaluationResults []EvaluationResult `type:"list"`

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetComplianceDetailsByConfigRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetComplianceDetailsByConfigRule = "GetComplianceDetailsByConfigRule"

// GetComplianceDetailsByConfigRuleRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the evaluation results for the specified AWS Config rule. The results
// indicate which AWS resources were evaluated by the rule, when each resource
// was last evaluated, and whether each resource complies with the rule.
//
//    // Example sending a request using GetComplianceDetailsByConfigRuleRequest.
//    req := client.GetComplianceDetailsByConfigRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetComplianceDetailsByConfigRule
func (c *Client) GetComplianceDetailsByConfigRuleRequest(input *GetComplianceDetailsByConfigRuleInput) GetComplianceDetailsByConfigRuleRequest {
	op := &aws.Operation{
		Name:       opGetComplianceDetailsByConfigRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetComplianceDetailsByConfigRuleInput{}
	}

	req := c.newRequest(op, input, &GetComplianceDetailsByConfigRuleOutput{})
	return GetComplianceDetailsByConfigRuleRequest{Request: req, Input: input, Copy: c.GetComplianceDetailsByConfigRuleRequest}
}

// GetComplianceDetailsByConfigRuleRequest is the request type for the
// GetComplianceDetailsByConfigRule API operation.
type GetComplianceDetailsByConfigRuleRequest struct {
	*aws.Request
	Input *GetComplianceDetailsByConfigRuleInput
	Copy  func(*GetComplianceDetailsByConfigRuleInput) GetComplianceDetailsByConfigRuleRequest
}

// Send marshals and sends the GetComplianceDetailsByConfigRule API request.
func (r GetComplianceDetailsByConfigRuleRequest) Send(ctx context.Context) (*GetComplianceDetailsByConfigRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetComplianceDetailsByConfigRuleResponse{
		GetComplianceDetailsByConfigRuleOutput: r.Request.Data.(*GetComplianceDetailsByConfigRuleOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetComplianceDetailsByConfigRuleResponse is the response type for the
// GetComplianceDetailsByConfigRule API operation.
type GetComplianceDetailsByConfigRuleResponse struct {
	*GetComplianceDetailsByConfigRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetComplianceDetailsByConfigRule request.
func (r *GetComplianceDetailsByConfigRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}