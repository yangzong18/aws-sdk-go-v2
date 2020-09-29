// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the platform branches available for your account in an AWS Region.
// Provides summary information about each platform branch. For definitions of
// platform branch and other platform-related terms, see AWS Elastic Beanstalk
// Platforms Glossary
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html).
func (c *Client) ListPlatformBranches(ctx context.Context, params *ListPlatformBranchesInput, optFns ...func(*Options)) (*ListPlatformBranchesOutput, error) {
	stack := middleware.NewStack("ListPlatformBranches", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListPlatformBranchesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPlatformBranches(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListPlatformBranches",
			Err:           err,
		}
	}
	out := result.(*ListPlatformBranchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlatformBranchesInput struct {

	// Criteria for restricting the resulting list of platform branches. The filter is
	// evaluated as a logical conjunction (AND) of the separate SearchFilter terms. The
	// following list shows valid attribute values for each of the SearchFilter terms.
	// Most operators take a single value. The in and not_in operators can take
	// multiple values.
	//
	//     * Attribute = BranchName:
	//
	//         * Operator: = | != |
	// begins_with | ends_with | contains | in | not_in
	//
	//     * Attribute =
	// LifecycleState:
	//
	//         * Operator: = | != | in | not_in
	//
	//         * Values:
	// beta | supported | deprecated | retired
	//
	//     * Attribute = PlatformName:
	//
	//
	// * Operator: = | != | begins_with | ends_with | contains | in | not_in
	//
	//     *
	// Attribute = TierType:
	//
	//         * Operator: = | !=
	//
	//         * Values:
	// WebServer/Standard | Worker/SQS/HTTP
	//
	// Array size: limited to 10 SearchFilter
	// objects. Within each SearchFilter item, the Values array is limited to 10 items.
	Filters []*types.SearchFilter

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string

	// The maximum number of platform branch values returned in one call.
	MaxRecords *int32
}

type ListPlatformBranchesOutput struct {

	// Summary information about the platform branches.
	PlatformBranchSummaryList []*types.PlatformBranchSummary

	// In a paginated request, if this value isn't null, it's the token that you can
	// pass in a subsequent request to get the next response page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListPlatformBranchesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListPlatformBranches{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListPlatformBranches{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPlatformBranches(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ListPlatformBranches",
	}
}