package aws

import (
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
)

// UseServiceDefaultRetries instructs the config to use the service's own
// default number of retries. This will be the default action if
// Config.MaxRetries is nil also.
const UseServiceDefaultRetries = -1

// RequestRetryer is an alias for a type that implements the request.Retryer
// interface.
type RequestRetryer interface***REMOVED******REMOVED***

// A Config provides service configuration for service clients. By default,
// all clients will use the defaults.DefaultConfig tructure.
//
//     // Create Session with MaxRetry configuration to be shared by multiple
//     // service clients.
//     sess := session.Must(session.NewSession(&aws.Config***REMOVED***
//         MaxRetries: aws.Int(3),
// ***REMOVED***))
//
//     // Create S3 service client with a specific Region.
//     svc := s3.New(sess, &aws.Config***REMOVED***
//         Region: aws.String("us-west-2"),
// ***REMOVED***)
type Config struct ***REMOVED***
	// Enables verbose error printing of all credential chain errors.
	// Should be used when wanting to see all errors while attempting to
	// retrieve credentials.
	CredentialsChainVerboseErrors *bool

	// The credentials object to use when signing requests. Defaults to a
	// chain of credential providers to search for credentials in environment
	// variables, shared credential file, and EC2 Instance Roles.
	Credentials *credentials.Credentials

	// An optional endpoint URL (hostname only or fully qualified URI)
	// that overrides the default generated endpoint for a client. Set this
	// to `""` to use the default generated endpoint.
	//
	// @note You must still provide a `Region` value when specifying an
	//   endpoint for a client.
	Endpoint *string

	// The resolver to use for looking up endpoints for AWS service clients
	// to use based on region.
	EndpointResolver endpoints.Resolver

	// EnforceShouldRetryCheck is used in the AfterRetryHandler to always call
	// ShouldRetry regardless of whether or not if request.Retryable is set.
	// This will utilize ShouldRetry method of custom retryers. If EnforceShouldRetryCheck
	// is not set, then ShouldRetry will only be called if request.Retryable is nil.
	// Proper handling of the request.Retryable field is important when setting this field.
	EnforceShouldRetryCheck *bool

	// The region to send requests to. This parameter is required and must
	// be configured globally or on a per-client basis unless otherwise
	// noted. A full list of regions is found in the "Regions and Endpoints"
	// document.
	//
	// @see http://docs.aws.amazon.com/general/latest/gr/rande.html
	//   AWS Regions and Endpoints
	Region *string

	// Set this to `true` to disable SSL when sending requests. Defaults
	// to `false`.
	DisableSSL *bool

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`.
	HTTPClient *http.Client

	// An integer value representing the logging level. The default log level
	// is zero (LogOff), which represents no logging. To enable logging set
	// to a LogLevel Value.
	LogLevel *LogLevelType

	// The logger writer interface to write logging messages to. Defaults to
	// standard out.
	Logger Logger

	// The maximum number of times that a request will be retried for failures.
	// Defaults to -1, which defers the max retry setting to the service
	// specific configuration.
	MaxRetries *int

	// Retryer guides how HTTP requests should be retried in case of
	// recoverable failures.
	//
	// When nil or the value does not implement the request.Retryer interface,
	// the client.DefaultRetryer will be used.
	//
	// When both Retryer and MaxRetries are non-nil, the former is used and
	// the latter ignored.
	//
	// To set the Retryer field in a type-safe manner and with chaining, use
	// the request.WithRetryer helper function:
	//
	//   cfg := request.WithRetryer(aws.NewConfig(), myRetryer)
	//
	Retryer RequestRetryer

	// Disables semantic parameter validation, which validates input for
	// missing required fields and/or other semantic request input errors.
	DisableParamValidation *bool

	// Disables the computation of request and response checksums, e.g.,
	// CRC32 checksums in Amazon DynamoDB.
	DisableComputeChecksums *bool

	// Set this to `true` to force the request to use path-style addressing,
	// i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default, the S3 client
	// will use virtual hosted bucket addressing when possible
	// (`http://BUCKET.s3.amazonaws.com/KEY`).
	//
	// @note This configuration option is specific to the Amazon S3 service.
	// @see http://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
	//   Amazon S3: Virtual Hosting of Buckets
	S3ForcePathStyle *bool

	// Set this to `true` to disable the SDK adding the `Expect: 100-Continue`
	// header to PUT requests over 2MB of content. 100-Continue instructs the
	// HTTP client not to send the body until the service responds with a
	// `continue` status. This is useful to prevent sending the request body
	// until after the request is authenticated, and validated.
	//
	// http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPUT.html
	//
	// 100-Continue is only enabled for Go 1.6 and above. See `http.Transport`'s
	// `ExpectContinueTimeout` for information on adjusting the continue wait
	// timeout. https://golang.org/pkg/net/http/#Transport
	//
	// You should use this flag to disble 100-Continue if you experience issues
	// with proxies or third party S3 compatible services.
	S3Disable100Continue *bool

	// Set this to `true` to enable S3 Accelerate feature. For all operations
	// compatible with S3 Accelerate will use the accelerate endpoint for
	// requests. Requests not compatible will fall back to normal S3 requests.
	//
	// The bucket must be enable for accelerate to be used with S3 client with
	// accelerate enabled. If the bucket is not enabled for accelerate an error
	// will be returned. The bucket name must be DNS compatible to also work
	// with accelerate.
	S3UseAccelerate *bool

	// Set this to `true` to disable the EC2Metadata client from overriding the
	// default http.Client's Timeout. This is helpful if you do not want the
	// EC2Metadata client to create a new http.Client. This options is only
	// meaningful if you're not already using a custom HTTP client with the
	// SDK. Enabled by default.
	//
	// Must be set and provided to the session.NewSession() in order to disable
	// the EC2Metadata overriding the timeout for default credentials chain.
	//
	// Example:
	//    sess := session.Must(session.NewSession(aws.NewConfig()
	//       .WithEC2MetadataDiableTimeoutOverride(true)))
	//
	//    svc := s3.New(sess)
	//
	EC2MetadataDisableTimeoutOverride *bool

	// Instructs the endpoint to be generated for a service client to
	// be the dual stack endpoint. The dual stack endpoint will support
	// both IPv4 and IPv6 addressing.
	//
	// Setting this for a service which does not support dual stack will fail
	// to make requets. It is not recommended to set this value on the session
	// as it will apply to all service clients created with the session. Even
	// services which don't support dual stack endpoints.
	//
	// If the Endpoint config value is also provided the UseDualStack flag
	// will be ignored.
	//
	// Only supported with.
	//
	//     sess := session.Must(session.NewSession())
	//
	//     svc := s3.New(sess, &aws.Config***REMOVED***
	//         UseDualStack: aws.Bool(true),
	// ***REMOVED***)
	UseDualStack *bool

	// SleepDelay is an override for the func the SDK will call when sleeping
	// during the lifecycle of a request. Specifically this will be used for
	// request delays. This value should only be used for testing. To adjust
	// the delay of a request see the aws/client.DefaultRetryer and
	// aws/request.Retryer.
	//
	// SleepDelay will prevent any Context from being used for canceling retry
	// delay of an API operation. It is recommended to not use SleepDelay at all
	// and specify a Retryer instead.
	SleepDelay func(time.Duration)

	// DisableRestProtocolURICleaning will not clean the URL path when making rest protocol requests.
	// Will default to false. This would only be used for empty directory names in s3 requests.
	//
	// Example:
	//    sess := session.Must(session.NewSession(&aws.Config***REMOVED***
	//         DisableRestProtocolURICleaning: aws.Bool(true),
	//***REMOVED***))
	//
	//    svc := s3.New(sess)
	//    out, err := svc.GetObject(&s3.GetObjectInput ***REMOVED***
	//    	Bucket: aws.String("bucketname"),
	//    	Key: aws.String("//foo//bar//moo"),
	//***REMOVED***)
	DisableRestProtocolURICleaning *bool
***REMOVED***

// NewConfig returns a new Config pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
//
//     // Create Session with MaxRetry configuration to be shared by multiple
//     // service clients.
//     sess := session.Must(session.NewSession(aws.NewConfig().
//         WithMaxRetries(3),
//     ))
//
//     // Create S3 service client with a specific Region.
//     svc := s3.New(sess, aws.NewConfig().
//         WithRegion("us-west-2"),
//     )
func NewConfig() *Config ***REMOVED***
	return &Config***REMOVED******REMOVED***
***REMOVED***

// WithCredentialsChainVerboseErrors sets a config verbose errors boolean and returning
// a Config pointer.
func (c *Config) WithCredentialsChainVerboseErrors(verboseErrs bool) *Config ***REMOVED***
	c.CredentialsChainVerboseErrors = &verboseErrs
	return c
***REMOVED***

// WithCredentials sets a config Credentials value returning a Config pointer
// for chaining.
func (c *Config) WithCredentials(creds *credentials.Credentials) *Config ***REMOVED***
	c.Credentials = creds
	return c
***REMOVED***

// WithEndpoint sets a config Endpoint value returning a Config pointer for
// chaining.
func (c *Config) WithEndpoint(endpoint string) *Config ***REMOVED***
	c.Endpoint = &endpoint
	return c
***REMOVED***

// WithEndpointResolver sets a config EndpointResolver value returning a
// Config pointer for chaining.
func (c *Config) WithEndpointResolver(resolver endpoints.Resolver) *Config ***REMOVED***
	c.EndpointResolver = resolver
	return c
***REMOVED***

// WithRegion sets a config Region value returning a Config pointer for
// chaining.
func (c *Config) WithRegion(region string) *Config ***REMOVED***
	c.Region = &region
	return c
***REMOVED***

// WithDisableSSL sets a config DisableSSL value returning a Config pointer
// for chaining.
func (c *Config) WithDisableSSL(disable bool) *Config ***REMOVED***
	c.DisableSSL = &disable
	return c
***REMOVED***

// WithHTTPClient sets a config HTTPClient value returning a Config pointer
// for chaining.
func (c *Config) WithHTTPClient(client *http.Client) *Config ***REMOVED***
	c.HTTPClient = client
	return c
***REMOVED***

// WithMaxRetries sets a config MaxRetries value returning a Config pointer
// for chaining.
func (c *Config) WithMaxRetries(max int) *Config ***REMOVED***
	c.MaxRetries = &max
	return c
***REMOVED***

// WithDisableParamValidation sets a config DisableParamValidation value
// returning a Config pointer for chaining.
func (c *Config) WithDisableParamValidation(disable bool) *Config ***REMOVED***
	c.DisableParamValidation = &disable
	return c
***REMOVED***

// WithDisableComputeChecksums sets a config DisableComputeChecksums value
// returning a Config pointer for chaining.
func (c *Config) WithDisableComputeChecksums(disable bool) *Config ***REMOVED***
	c.DisableComputeChecksums = &disable
	return c
***REMOVED***

// WithLogLevel sets a config LogLevel value returning a Config pointer for
// chaining.
func (c *Config) WithLogLevel(level LogLevelType) *Config ***REMOVED***
	c.LogLevel = &level
	return c
***REMOVED***

// WithLogger sets a config Logger value returning a Config pointer for
// chaining.
func (c *Config) WithLogger(logger Logger) *Config ***REMOVED***
	c.Logger = logger
	return c
***REMOVED***

// WithS3ForcePathStyle sets a config S3ForcePathStyle value returning a Config
// pointer for chaining.
func (c *Config) WithS3ForcePathStyle(force bool) *Config ***REMOVED***
	c.S3ForcePathStyle = &force
	return c
***REMOVED***

// WithS3Disable100Continue sets a config S3Disable100Continue value returning
// a Config pointer for chaining.
func (c *Config) WithS3Disable100Continue(disable bool) *Config ***REMOVED***
	c.S3Disable100Continue = &disable
	return c
***REMOVED***

// WithS3UseAccelerate sets a config S3UseAccelerate value returning a Config
// pointer for chaining.
func (c *Config) WithS3UseAccelerate(enable bool) *Config ***REMOVED***
	c.S3UseAccelerate = &enable
	return c
***REMOVED***

// WithUseDualStack sets a config UseDualStack value returning a Config
// pointer for chaining.
func (c *Config) WithUseDualStack(enable bool) *Config ***REMOVED***
	c.UseDualStack = &enable
	return c
***REMOVED***

// WithEC2MetadataDisableTimeoutOverride sets a config EC2MetadataDisableTimeoutOverride value
// returning a Config pointer for chaining.
func (c *Config) WithEC2MetadataDisableTimeoutOverride(enable bool) *Config ***REMOVED***
	c.EC2MetadataDisableTimeoutOverride = &enable
	return c
***REMOVED***

// WithSleepDelay overrides the function used to sleep while waiting for the
// next retry. Defaults to time.Sleep.
func (c *Config) WithSleepDelay(fn func(time.Duration)) *Config ***REMOVED***
	c.SleepDelay = fn
	return c
***REMOVED***

// MergeIn merges the passed in configs into the existing config object.
func (c *Config) MergeIn(cfgs ...*Config) ***REMOVED***
	for _, other := range cfgs ***REMOVED***
		mergeInConfig(c, other)
	***REMOVED***
***REMOVED***

func mergeInConfig(dst *Config, other *Config) ***REMOVED***
	if other == nil ***REMOVED***
		return
	***REMOVED***

	if other.CredentialsChainVerboseErrors != nil ***REMOVED***
		dst.CredentialsChainVerboseErrors = other.CredentialsChainVerboseErrors
	***REMOVED***

	if other.Credentials != nil ***REMOVED***
		dst.Credentials = other.Credentials
	***REMOVED***

	if other.Endpoint != nil ***REMOVED***
		dst.Endpoint = other.Endpoint
	***REMOVED***

	if other.EndpointResolver != nil ***REMOVED***
		dst.EndpointResolver = other.EndpointResolver
	***REMOVED***

	if other.Region != nil ***REMOVED***
		dst.Region = other.Region
	***REMOVED***

	if other.DisableSSL != nil ***REMOVED***
		dst.DisableSSL = other.DisableSSL
	***REMOVED***

	if other.HTTPClient != nil ***REMOVED***
		dst.HTTPClient = other.HTTPClient
	***REMOVED***

	if other.LogLevel != nil ***REMOVED***
		dst.LogLevel = other.LogLevel
	***REMOVED***

	if other.Logger != nil ***REMOVED***
		dst.Logger = other.Logger
	***REMOVED***

	if other.MaxRetries != nil ***REMOVED***
		dst.MaxRetries = other.MaxRetries
	***REMOVED***

	if other.Retryer != nil ***REMOVED***
		dst.Retryer = other.Retryer
	***REMOVED***

	if other.DisableParamValidation != nil ***REMOVED***
		dst.DisableParamValidation = other.DisableParamValidation
	***REMOVED***

	if other.DisableComputeChecksums != nil ***REMOVED***
		dst.DisableComputeChecksums = other.DisableComputeChecksums
	***REMOVED***

	if other.S3ForcePathStyle != nil ***REMOVED***
		dst.S3ForcePathStyle = other.S3ForcePathStyle
	***REMOVED***

	if other.S3Disable100Continue != nil ***REMOVED***
		dst.S3Disable100Continue = other.S3Disable100Continue
	***REMOVED***

	if other.S3UseAccelerate != nil ***REMOVED***
		dst.S3UseAccelerate = other.S3UseAccelerate
	***REMOVED***

	if other.UseDualStack != nil ***REMOVED***
		dst.UseDualStack = other.UseDualStack
	***REMOVED***

	if other.EC2MetadataDisableTimeoutOverride != nil ***REMOVED***
		dst.EC2MetadataDisableTimeoutOverride = other.EC2MetadataDisableTimeoutOverride
	***REMOVED***

	if other.SleepDelay != nil ***REMOVED***
		dst.SleepDelay = other.SleepDelay
	***REMOVED***

	if other.DisableRestProtocolURICleaning != nil ***REMOVED***
		dst.DisableRestProtocolURICleaning = other.DisableRestProtocolURICleaning
	***REMOVED***

	if other.EnforceShouldRetryCheck != nil ***REMOVED***
		dst.EnforceShouldRetryCheck = other.EnforceShouldRetryCheck
	***REMOVED***
***REMOVED***

// Copy will return a shallow copy of the Config object. If any additional
// configurations are provided they will be merged into the new config returned.
func (c *Config) Copy(cfgs ...*Config) *Config ***REMOVED***
	dst := &Config***REMOVED******REMOVED***
	dst.MergeIn(c)

	for _, cfg := range cfgs ***REMOVED***
		dst.MergeIn(cfg)
	***REMOVED***

	return dst
***REMOVED***
