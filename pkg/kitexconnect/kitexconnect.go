package kitexconnect

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

const (
	// Default RPC timeout value, equal to kitex default value.
	defaultRPCTimeout = time.Duration(0)
	// Default connection timeout value (50 milliseconds), equal to kitex default value.
	defaultConnectTimeout = time.Millisecond * 50
	// Default read/write timeout value (5 seconds), equal to kitex default value.
	defaultReadWriteTimeout = time.Second * 5

	// Default maximum retry times for the RPC calls, equal to kitex default value.
	defaultMaxRetryTimes = 2
)

// connectOptions holds configurations for establishing connections, including timeouts and retry policies.
type connectOptions struct {
	connectTimeout time.Duration // Timeout for establishing a connection
	rpcTimeout     time.Duration // Timeout for the RPC call
	rwTimeout      time.Duration // Timeout for read/write operations

	retryTimes *int // Pointer to the maximum retry times

	// Optional retry functions
	errorRetry func(ctx context.Context, err error, ri rpcinfo.RPCInfo) bool        // Function to determine if an error should trigger a retry
	respRetry  func(ctx context.Context, resp interface{}, ri rpcinfo.RPCInfo) bool // Function to determine if a response should trigger a retry
}

// NewConnectOptions creates and returns a new connectOptions instance with default values.
func NewConnectOptions() *connectOptions {
	return &connectOptions{
		connectTimeout: defaultConnectTimeout,
		rpcTimeout:     defaultRPCTimeout,
		rwTimeout:      defaultReadWriteTimeout,
	}
}

// SetConnectTimeout sets the connection timeout and returns the updated connectOptions for chaining.
func (co *connectOptions) SetConnectTimeout(timeout time.Duration) *connectOptions {
	co.connectTimeout = timeout
	return co
}

// SetRPCTimeout sets the RPC timeout and returns the updated connectOptions for chaining.
func (co *connectOptions) SetRPCTimeout(timeout time.Duration) *connectOptions {
	co.rpcTimeout = timeout
	return co
}

// SetReadWriteTimeout sets the read/write timeout and returns the updated connectOptions for chaining.
func (co *connectOptions) SetReadWriteTimeout(timeout time.Duration) *connectOptions {
	co.rwTimeout = timeout
	return co
}

// SetMaxRetryTimes sets the maximum retry times and returns the updated connectOptions for chaining.
func (co *connectOptions) SetMaxRetryTimes(retryTimes int) *connectOptions {
	co.retryTimes = &retryTimes
	return co
}

// SetErrorRetry sets the error retry function and returns the updated connectOptions for chaining.
func (co *connectOptions) SetErrorRetry(errorRetry func(ctx context.Context, err error, ri rpcinfo.RPCInfo) bool) *connectOptions {
	co.errorRetry = errorRetry
	return co
}

// SetResponseRetry sets the response retry function and returns the updated connectOptions for chaining.
func (co *connectOptions) SetResponseRetry(respRetry func(ctx context.Context, resp interface{}, ri rpcinfo.RPCInfo) bool) *connectOptions {
	co.respRetry = respRetry
	return co
}

// RPCTimeout returns the RPC timeout value.
func (co *connectOptions) RPCTimeout() time.Duration {
	return co.rpcTimeout
}

// ConnectTimeout returns the connection timeout value.
func (co *connectOptions) ConnectTimeout() time.Duration {
	return co.connectTimeout
}

// ReadWriteTimeout returns the read/write timeout value.
func (co *connectOptions) ReadWriteTimeout() time.Duration {
	return co.rwTimeout
}

// connectPolicy manages connection options for RPC methods, ensuring safe access with mutex.
type connectPolicy struct {
	mu       sync.Mutex                 // Mutex for thread safety
	registry map[string]*connectOptions // Map to store connection options for each method
}

// NewConnectPolicy creates and returns a new connectPolicy instance with an initialized registry.
func NewConnectPolicy() *connectPolicy {
	return &connectPolicy{
		registry: make(map[string]*connectOptions, 0),
	}
}

// M retrieves the connectOptions for a given method, creating a new one if it does not exist.
func (cp *connectPolicy) M(method string) *connectOptions {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	// Create new options if none found.
	co, ok := cp.registry[method]
	if !ok {
		co = NewConnectOptions()
		cp.registry[method] = co
	}

	return co
}

// Timeouts retrieves the timeout settings for a given RPCInfo.
func (cp *connectPolicy) Timeouts(ri rpcinfo.RPCInfo) rpcinfo.Timeouts {
	timeouts, ok := cp.registry[ri.To().Method()]
	if ok {
		return timeouts
	}

	return NewConnectOptions()
}

// Build constructs client options based on the defined connectPolicy.
func (cp *connectPolicy) Build() []client.Option {
	kitexClientOptions := make([]client.Option, 0)

	policies := make(map[string]retry.Policy)
	for method, co := range cp.registry {
		// Create a new failure policy with default values.
		fp := retry.NewFailurePolicy()
		// Set maximum retry times if setted.
		if co.retryTimes != nil {
			fp.WithMaxRetryTimes(*co.retryTimes)
		}

		fp.ShouldResultRetry = new(retry.ShouldResultRetry)

		if co.errorRetry != nil {
			fp.ShouldResultRetry.ErrorRetryWithCtx = co.errorRetry
		}

		if co.respRetry != nil {
			fp.ShouldResultRetry.RespRetryWithCtx = co.respRetry
		}
		// Build and add the retry policy for the method
		policies[method] = retry.BuildFailurePolicy(fp)
	}

	// Add retry policies to client options.
	kitexClientOptions = append(kitexClientOptions, client.WithRetryMethodPolicies(policies))
	// Register timeout provider.
	kitexClientOptions = append(kitexClientOptions, client.WithTimeoutProvider(cp))

	return kitexClientOptions
}
