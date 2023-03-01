// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists one or more AMIs that are currently in the Recycle Bin. For more
// information, see Recycle Bin
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/recycle-bin.html) in the
// Amazon EC2 User Guide.
func (c *Client) ListImagesInRecycleBin(ctx context.Context, params *ListImagesInRecycleBinInput, optFns ...func(*Options)) (*ListImagesInRecycleBinOutput, error) {
	if params == nil {
		params = &ListImagesInRecycleBinInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImagesInRecycleBin", params, optFns, c.addOperationListImagesInRecycleBinMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImagesInRecycleBinOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagesInRecycleBinInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IDs of the AMIs to list. Omit this parameter to list all of the AMIs that
	// are in the Recycle Bin. You can specify up to 20 IDs in a single request.
	ImageIds []string

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListImagesInRecycleBinOutput struct {

	// Information about the AMIs.
	Images []types.ImageRecycleBinInfo

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImagesInRecycleBinMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpListImagesInRecycleBin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpListImagesInRecycleBin{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImagesInRecycleBin(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListImagesInRecycleBinAPIClient is a client that implements the
// ListImagesInRecycleBin operation.
type ListImagesInRecycleBinAPIClient interface {
	ListImagesInRecycleBin(context.Context, *ListImagesInRecycleBinInput, ...func(*Options)) (*ListImagesInRecycleBinOutput, error)
}

var _ ListImagesInRecycleBinAPIClient = (*Client)(nil)

// ListImagesInRecycleBinPaginatorOptions is the paginator options for
// ListImagesInRecycleBin
type ListImagesInRecycleBinPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImagesInRecycleBinPaginator is a paginator for ListImagesInRecycleBin
type ListImagesInRecycleBinPaginator struct {
	options   ListImagesInRecycleBinPaginatorOptions
	client    ListImagesInRecycleBinAPIClient
	params    *ListImagesInRecycleBinInput
	nextToken *string
	firstPage bool
}

// NewListImagesInRecycleBinPaginator returns a new ListImagesInRecycleBinPaginator
func NewListImagesInRecycleBinPaginator(client ListImagesInRecycleBinAPIClient, params *ListImagesInRecycleBinInput, optFns ...func(*ListImagesInRecycleBinPaginatorOptions)) *ListImagesInRecycleBinPaginator {
	if params == nil {
		params = &ListImagesInRecycleBinInput{}
	}

	options := ListImagesInRecycleBinPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImagesInRecycleBinPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImagesInRecycleBinPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImagesInRecycleBin page.
func (p *ListImagesInRecycleBinPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImagesInRecycleBinOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListImagesInRecycleBin(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImagesInRecycleBin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ListImagesInRecycleBin",
	}
}
