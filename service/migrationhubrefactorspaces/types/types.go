// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A wrapper object holding the Amazon API Gateway proxy configuration.
type ApiGatewayProxyConfig struct {

	// The resource ID of the API Gateway for the proxy.
	ApiGatewayId *string

	// The type of API Gateway endpoint created.
	EndpointType ApiGatewayEndpointType

	// The Amazon Resource Name (ARN) of the Network Load Balancer configured by the
	// API Gateway proxy.
	NlbArn *string

	// The name of the Network Load Balancer that is configured by the API Gateway
	// proxy.
	NlbName *string

	// The endpoint URL of the API Gateway proxy.
	ProxyUrl *string

	// The name of the API Gateway stage. The name defaults to prod .
	StageName *string

	// The VpcLink ID of the API Gateway proxy.
	VpcLinkId *string

	noSmithyDocumentSerde
}

// A wrapper object holding the Amazon API Gateway endpoint input.
type ApiGatewayProxyInput struct {

	// The type of endpoint to use for the API Gateway proxy. If no value is specified
	// in the request, the value is set to REGIONAL by default. If the value is set to
	// PRIVATE in the request, this creates a private API endpoint that is isolated
	// from the public internet. The private endpoint can only be accessed by using
	// Amazon Virtual Private Cloud (Amazon VPC) endpoints for Amazon API Gateway that
	// have been granted access.
	EndpointType ApiGatewayEndpointType

	// The name of the API Gateway stage. The name defaults to prod .
	StageName *string

	noSmithyDocumentSerde
}

// A wrapper object holding the Amazon API Gateway proxy summary.
type ApiGatewayProxySummary struct {

	// The resource ID of the API Gateway for the proxy.
	ApiGatewayId *string

	// The type of API Gateway endpoint created.
	EndpointType ApiGatewayEndpointType

	// The Amazon Resource Name (ARN) of the Network Load Balancer configured by the
	// API Gateway proxy.
	NlbArn *string

	// The name of the Network Load Balancer that is configured by the API Gateway
	// proxy.
	NlbName *string

	// The endpoint URL of the API Gateway proxy.
	ProxyUrl *string

	// The name of the API Gateway stage. The name defaults to prod .
	StageName *string

	// The VpcLink ID of the API Gateway proxy.
	VpcLinkId *string

	noSmithyDocumentSerde
}

// The list of ApplicationSummary objects.
type ApplicationSummary struct {

	// The endpoint URL of the Amazon API Gateway proxy.
	ApiGatewayProxy *ApiGatewayProxySummary

	// The unique identifier of the application.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the application.
	Arn *string

	// The Amazon Web Services account ID of the application creator.
	CreatedByAccountId *string

	// A timestamp that indicates when the application is created.
	CreatedTime *time.Time

	// The unique identifier of the environment.
	EnvironmentId *string

	// Any error associated with the application resource.
	Error *ErrorResponse

	// A timestamp that indicates when the application was last updated.
	LastUpdatedTime *time.Time

	// The name of the application.
	Name *string

	// The Amazon Web Services account ID of the application owner (which is always
	// the same as the environment owner account ID).
	OwnerAccountId *string

	// The proxy type of the proxy created within the application.
	ProxyType ProxyType

	// The current state of the application.
	State ApplicationState

	// The tags assigned to the application.
	Tags map[string]string

	// The ID of the virtual private cloud (VPC).
	VpcId *string

	noSmithyDocumentSerde
}

// The configuration for the default route type.
type DefaultRouteInput struct {

	// If set to ACTIVE , traffic is forwarded to this route’s service after the route
	// is created.
	ActivationState RouteActivationState

	noSmithyDocumentSerde
}

// The summary information for environments as a response to ListEnvironments .
type EnvironmentSummary struct {

	// The Amazon Resource Name (ARN) of the environment.
	Arn *string

	// A timestamp that indicates when the environment is created.
	CreatedTime *time.Time

	// A description of the environment.
	Description *string

	// The unique identifier of the environment.
	EnvironmentId *string

	// Any error associated with the environment resource.
	Error *ErrorResponse

	// A timestamp that indicates when the environment was last updated.
	LastUpdatedTime *time.Time

	// The name of the environment.
	Name *string

	// The network fabric type of the environment.
	NetworkFabricType NetworkFabricType

	// The Amazon Web Services account ID of the environment owner.
	OwnerAccountId *string

	// The current state of the environment.
	State EnvironmentState

	// The tags assigned to the environment.
	Tags map[string]string

	// The ID of the transit gateway set up by the environment.
	TransitGatewayId *string

	noSmithyDocumentSerde
}

// Provides summary information for the EnvironmentVpc resource as a response to
// ListEnvironmentVpc .
type EnvironmentVpc struct {

	// The Amazon Web Services account ID of the virtual private cloud (VPC) owner.
	AccountId *string

	// The list of Amazon Virtual Private Cloud (Amazon VPC) CIDR blocks.
	CidrBlocks []string

	// A timestamp that indicates when the VPC is first added to the environment.
	CreatedTime *time.Time

	// The unique identifier of the environment.
	EnvironmentId *string

	// A timestamp that indicates when the VPC was last updated by the environment.
	LastUpdatedTime *time.Time

	// The ID of the VPC.
	VpcId *string

	// The name of the VPC at the time it is added to the environment.
	VpcName *string

	noSmithyDocumentSerde
}

// Error associated with a resource returned for a Get or List resource response.
type ErrorResponse struct {

	// The Amazon Web Services account ID of the resource owner.
	AccountId *string

	// Additional details about the error.
	AdditionalDetails map[string]string

	// The error code associated with the error.
	Code ErrorCode

	// The message associated with the error.
	Message *string

	// The ID of the resource.
	ResourceIdentifier *string

	// The type of resource.
	ResourceType ErrorResourceType

	noSmithyDocumentSerde
}

// The configuration for the Lambda endpoint type.
type LambdaEndpointConfig struct {

	// The Amazon Resource Name (ARN) of the Lambda endpoint.
	Arn *string

	noSmithyDocumentSerde
}

// The input for the Lambda endpoint type.
type LambdaEndpointInput struct {

	// The Amazon Resource Name (ARN) of the Lambda function or alias.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

// The summary for the Lambda endpoint type.
type LambdaEndpointSummary struct {

	// The Amazon Resource Name (ARN) of the Lambda endpoint.
	Arn *string

	noSmithyDocumentSerde
}

// The summary information for the routes as a response to ListRoutes .
type RouteSummary struct {

	// The unique identifier of the application.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the route.
	Arn *string

	// The Amazon Web Services account ID of the route creator.
	CreatedByAccountId *string

	// A timestamp that indicates when the route is created.
	CreatedTime *time.Time

	// The unique identifier of the environment.
	EnvironmentId *string

	// Any error associated with the route resource.
	Error *ErrorResponse

	// Indicates whether to match all subpaths of the given source path. If this value
	// is false , requests must match the source path exactly before they are forwarded
	// to this route's service.
	IncludeChildPaths *bool

	// A timestamp that indicates when the route was last updated.
	LastUpdatedTime *time.Time

	// A list of HTTP methods to match. An empty list matches all values. If a method
	// is present, only HTTP requests using that method are forwarded to this route’s
	// service.
	Methods []HttpMethod

	// The Amazon Web Services account ID of the route owner.
	OwnerAccountId *string

	// A mapping of Amazon API Gateway path resources to resource IDs.
	PathResourceToId map[string]string

	// The unique identifier of the route.
	RouteId *string

	// The route type of the route.
	RouteType RouteType

	// The unique identifier of the service.
	ServiceId *string

	// The path to use to match traffic. Paths must start with / and are relative to
	// the base of the application.
	SourcePath *string

	// The current state of the route.
	State RouteState

	// The tags assigned to the route.
	Tags map[string]string

	noSmithyDocumentSerde
}

// A summary for the service as a response to ListServices .
type ServiceSummary struct {

	// The unique identifier of the application.
	ApplicationId *string

	// The Amazon Resource Name (ARN) of the service.
	Arn *string

	// The Amazon Web Services account ID of the service creator.
	CreatedByAccountId *string

	// A timestamp that indicates when the service is created.
	CreatedTime *time.Time

	// A description of the service.
	Description *string

	// The endpoint type of the service.
	EndpointType ServiceEndpointType

	// The unique identifier of the environment.
	EnvironmentId *string

	// Any error associated with the service resource.
	Error *ErrorResponse

	// A summary of the configuration for the Lambda endpoint type.
	LambdaEndpoint *LambdaEndpointSummary

	// A timestamp that indicates when the service was last updated.
	LastUpdatedTime *time.Time

	// The name of the service.
	Name *string

	// The Amazon Web Services account ID of the service owner.
	OwnerAccountId *string

	// The unique identifier of the service.
	ServiceId *string

	// The current state of the service.
	State ServiceState

	// The tags assigned to the service.
	Tags map[string]string

	// The summary of the configuration for the URL endpoint type.
	UrlEndpoint *UrlEndpointSummary

	// The ID of the virtual private cloud (VPC).
	VpcId *string

	noSmithyDocumentSerde
}

// The configuration for the URI path route type.
type UriPathRouteInput struct {

	// If set to ACTIVE , traffic is forwarded to this route’s service after the route
	// is created.
	//
	// This member is required.
	ActivationState RouteActivationState

	// The path to use to match traffic. Paths must start with / and are relative to
	// the base of the application.
	//
	// This member is required.
	SourcePath *string

	// Indicates whether to match all subpaths of the given source path. If this value
	// is false , requests must match the source path exactly before they are forwarded
	// to this route's service.
	IncludeChildPaths *bool

	// A list of HTTP methods to match. An empty list matches all values. If a method
	// is present, only HTTP requests using that method are forwarded to this route’s
	// service.
	Methods []HttpMethod

	noSmithyDocumentSerde
}

// The configuration for the URL endpoint type.
type UrlEndpointConfig struct {

	// The health check URL of the URL endpoint type.
	HealthUrl *string

	// The HTTP URL endpoint.
	Url *string

	noSmithyDocumentSerde
}

// The configuration for the URL endpoint type.
type UrlEndpointInput struct {

	// The URL to route traffic to. The URL must be an rfc3986-formatted URL (https://datatracker.ietf.org/doc/html/rfc3986)
	// . If the host is a domain name, the name must be resolvable over the public
	// internet. If the scheme is https , the top level domain of the host must be
	// listed in the IANA root zone database (https://www.iana.org/domains/root/db) .
	//
	// This member is required.
	Url *string

	// The health check URL of the URL endpoint type. If the URL is a public endpoint,
	// the HealthUrl must also be a public endpoint. If the URL is a private endpoint
	// inside a virtual private cloud (VPC), the health URL must also be a private
	// endpoint, and the host must be the same as the URL.
	HealthUrl *string

	noSmithyDocumentSerde
}

// The summary of the configuration for the URL endpoint type.
type UrlEndpointSummary struct {

	// The health check URL of the URL endpoint type. If the URL is a public endpoint,
	// the HealthUrl must also be a public endpoint. If the URL is a private endpoint
	// inside a virtual private cloud (VPC), the health URL must also be a private
	// endpoint, and the host must be the same as the URL.
	HealthUrl *string

	// The URL to route traffic to. The URL must be an rfc3986-formatted URL (https://datatracker.ietf.org/doc/html/rfc3986)
	// . If the host is a domain name, the name must be resolvable over the public
	// internet. If the scheme is https , the top level domain of the host must be
	// listed in the IANA root zone database (https://www.iana.org/domains/root/db) .
	Url *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
