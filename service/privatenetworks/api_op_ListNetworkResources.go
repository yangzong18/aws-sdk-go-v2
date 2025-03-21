// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/privatenetworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists network resources. Add filters to your request to return a more specific
// list of results. Use filters to match the Amazon Resource Name (ARN) of an order
// or the status of network resources.
//
// If you specify multiple filters, filters are joined with an OR, and the request
// returns results that match all of the specified filters.
func (c *Client) ListNetworkResources(ctx context.Context, params *ListNetworkResourcesInput, optFns ...func(*Options)) (*ListNetworkResourcesOutput, error) {
	if params == nil {
		params = &ListNetworkResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNetworkResources", params, optFns, c.addOperationListNetworkResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNetworkResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNetworkResourcesInput struct {

	// The Amazon Resource Name (ARN) of the network.
	//
	// This member is required.
	NetworkArn *string

	// The filters.
	//
	//   - ORDER - The Amazon Resource Name (ARN) of the order.
	//
	//   - STATUS - The status ( AVAILABLE | DELETED | DELETING | PENDING |
	//   PENDING_RETURN | PROVISIONING | SHIPPED ).
	//
	// Filter values are case sensitive. If you specify multiple values for a filter,
	// the values are joined with an OR , and the request returns all results that
	// match any of the specified values.
	Filters map[string][]string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	StartToken *string

	noSmithyDocumentSerde
}

type ListNetworkResourcesOutput struct {

	// Information about network resources.
	NetworkResources []types.NetworkResource

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNetworkResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNetworkResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNetworkResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListNetworkResources"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpListNetworkResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNetworkResources(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

// ListNetworkResourcesPaginatorOptions is the paginator options for
// ListNetworkResources
type ListNetworkResourcesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNetworkResourcesPaginator is a paginator for ListNetworkResources
type ListNetworkResourcesPaginator struct {
	options   ListNetworkResourcesPaginatorOptions
	client    ListNetworkResourcesAPIClient
	params    *ListNetworkResourcesInput
	nextToken *string
	firstPage bool
}

// NewListNetworkResourcesPaginator returns a new ListNetworkResourcesPaginator
func NewListNetworkResourcesPaginator(client ListNetworkResourcesAPIClient, params *ListNetworkResourcesInput, optFns ...func(*ListNetworkResourcesPaginatorOptions)) *ListNetworkResourcesPaginator {
	if params == nil {
		params = &ListNetworkResourcesInput{}
	}

	options := ListNetworkResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNetworkResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.StartToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNetworkResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNetworkResources page.
func (p *ListNetworkResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNetworkResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.StartToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	optFns = append([]func(*Options){
		addIsPaginatorUserAgent,
	}, optFns...)
	result, err := p.client.ListNetworkResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ListNetworkResourcesAPIClient is a client that implements the
// ListNetworkResources operation.
type ListNetworkResourcesAPIClient interface {
	ListNetworkResources(context.Context, *ListNetworkResourcesInput, ...func(*Options)) (*ListNetworkResourcesOutput, error)
}

var _ ListNetworkResourcesAPIClient = (*Client)(nil)

func newServiceMetadataMiddleware_opListNetworkResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListNetworkResources",
	}
}
