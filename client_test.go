package bsvrates

import (
	"net/http"
	"strings"
	"testing"
	"time"
)

// TestNewClient test new client
func TestNewClient(t *testing.T) {
	t.Parallel()

	client := NewClient(nil, nil, DefaultProviders)

	if client == nil {
		t.Fatal("failed to load client")
	}

	// Test default providers
	if client.Providers != DefaultProviders {
		t.Fatalf("expected default bits (%d) to be set, got: %d", DefaultProviders, client.Providers)
	}
}

// TestNewClient_CustomHttpClient test new client with custom HTTP client
func TestNewClient_CustomHttpClient(t *testing.T) {
	t.Parallel()

	client := NewClient(nil, http.DefaultClient, ProviderPreev)

	if client == nil {
		t.Fatal("failed to load client")
	}

	// Test providers

	if client.Providers&ProviderPreev == 0 {
		t.Fatalf("expected bit %d to be set, got: %d", ProviderPreev, client.Providers)
	}
}

// BenchmarkNewClient benchmarks the NewClient method
func BenchmarkNewClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewClient(nil, nil, DefaultProviders)
	}
}

// TestDefaultClientOptions tests setting DefaultClientOptions()
func TestDefaultClientOptions(t *testing.T) {
	t.Parallel()

	options := DefaultClientOptions()

	if options.UserAgent != defaultUserAgent {
		t.Fatalf("expected value: %s got: %s", defaultUserAgent, options.UserAgent)
	}

	if options.BackOffExponentFactor != 2.0 {
		t.Fatalf("expected value: %f got: %f", 2.0, options.BackOffExponentFactor)
	}

	if options.BackOffInitialTimeout != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffInitialTimeout)
	}

	if options.BackOffMaximumJitterInterval != 2*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 2*time.Millisecond, options.BackOffMaximumJitterInterval)
	}

	if options.BackOffMaxTimeout != 10*time.Millisecond {
		t.Fatalf("expected value: %v got: %v", 10*time.Millisecond, options.BackOffMaxTimeout)
	}

	if options.DialerKeepAlive != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.DialerKeepAlive)
	}

	if options.DialerTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.DialerTimeout)
	}

	if options.RequestRetryCount != 2 {
		t.Fatalf("expected value: %v got: %v", 2, options.RequestRetryCount)
	}

	if options.RequestTimeout != 10*time.Second {
		t.Fatalf("expected value: %v got: %v", 10*time.Second, options.RequestTimeout)
	}

	if options.TransportExpectContinueTimeout != 3*time.Second {
		t.Fatalf("expected value: %v got: %v", 3*time.Second, options.TransportExpectContinueTimeout)
	}

	if options.TransportIdleTimeout != 20*time.Second {
		t.Fatalf("expected value: %v got: %v", 20*time.Second, options.TransportIdleTimeout)
	}

	if options.TransportMaxIdleConnections != 10 {
		t.Fatalf("expected value: %v got: %v", 10, options.TransportMaxIdleConnections)
	}

	if options.TransportTLSHandshakeTimeout != 5*time.Second {
		t.Fatalf("expected value: %v got: %v", 5*time.Second, options.TransportTLSHandshakeTimeout)
	}
}

// TestDefaultClientOptions_NoRetry will set 0 retry counts
func TestDefaultClientOptions_NoRetry(t *testing.T) {
	options := DefaultClientOptions()
	options.RequestRetryCount = 0
	client := NewClient(options, nil, DefaultProviders)

	if client == nil {
		t.Fatal("failed to load client")
	}

	// todo: add additional checks on the Client (methods)
}

// TestClientOptions_ToPreevOptions tests setting ToPreevOptions()
func TestClientOptions_ToPreevOptions(t *testing.T) {
	t.Parallel()

	options := DefaultClientOptions()

	preevOptions := options.ToPreevOptions()

	if !strings.Contains(preevOptions.UserAgent, "go-preev") {
		t.Fatalf("expected value: %s got: %s", "go-preev", preevOptions.UserAgent)
	}

	if preevOptions.BackOffExponentFactor != options.BackOffExponentFactor {
		t.Fatalf("expected value: %f got: %f", options.BackOffExponentFactor, preevOptions.BackOffExponentFactor)
	}

	if preevOptions.BackOffInitialTimeout != options.BackOffInitialTimeout {
		t.Fatalf("expected value: %v got: %v", options.BackOffInitialTimeout, preevOptions.BackOffInitialTimeout)
	}

	if preevOptions.BackOffMaximumJitterInterval != options.BackOffMaximumJitterInterval {
		t.Fatalf("expected value: %v got: %v", options.BackOffMaximumJitterInterval, preevOptions.BackOffMaximumJitterInterval)
	}

	if preevOptions.BackOffMaxTimeout != options.BackOffMaxTimeout {
		t.Fatalf("expected value: %v got: %v", options.BackOffMaxTimeout, preevOptions.BackOffMaxTimeout)
	}

	if preevOptions.DialerKeepAlive != options.DialerKeepAlive {
		t.Fatalf("expected value: %v got: %v", options.DialerKeepAlive, preevOptions.DialerKeepAlive)
	}

	if preevOptions.DialerTimeout != options.DialerTimeout {
		t.Fatalf("expected value: %v got: %v", options.DialerTimeout, preevOptions.DialerTimeout)
	}

	if preevOptions.RequestRetryCount != options.RequestRetryCount {
		t.Fatalf("expected value: %v got: %v", options.RequestRetryCount, preevOptions.RequestRetryCount)
	}

	if preevOptions.RequestTimeout != options.RequestTimeout {
		t.Fatalf("expected value: %v got: %v", options.RequestTimeout, preevOptions.RequestTimeout)
	}

	if preevOptions.TransportExpectContinueTimeout != options.TransportExpectContinueTimeout {
		t.Fatalf("expected value: %v got: %v", options.TransportExpectContinueTimeout, preevOptions.TransportExpectContinueTimeout)
	}

	if preevOptions.TransportIdleTimeout != options.TransportIdleTimeout {
		t.Fatalf("expected value: %v got: %v", options.TransportIdleTimeout, preevOptions.TransportIdleTimeout)
	}

	if preevOptions.TransportMaxIdleConnections != options.TransportMaxIdleConnections {
		t.Fatalf("expected value: %v got: %v", options.TransportMaxIdleConnections, preevOptions.TransportMaxIdleConnections)
	}

	if preevOptions.TransportTLSHandshakeTimeout != options.TransportTLSHandshakeTimeout {
		t.Fatalf("expected value: %v got: %v", options.TransportTLSHandshakeTimeout, preevOptions.TransportTLSHandshakeTimeout)
	}
}

// TestClientOptions_ToWhatsOnChainOptions tests setting ToWhatsOnChainOptions()
func TestClientOptions_ToWhatsOnChainOptions(t *testing.T) {
	t.Parallel()

	options := DefaultClientOptions()

	whatsOnChainOptions := options.ToWhatsOnChainOptions()

	if !strings.Contains(whatsOnChainOptions.UserAgent, "go-whatsonchain") {
		t.Fatalf("expected value: %s got: %s", "go-whatsonchain", whatsOnChainOptions.UserAgent)
	}

	if whatsOnChainOptions.BackOffExponentFactor != options.BackOffExponentFactor {
		t.Fatalf("expected value: %f got: %f", options.BackOffExponentFactor, whatsOnChainOptions.BackOffExponentFactor)
	}

	if whatsOnChainOptions.BackOffInitialTimeout != options.BackOffInitialTimeout {
		t.Fatalf("expected value: %v got: %v", options.BackOffInitialTimeout, whatsOnChainOptions.BackOffInitialTimeout)
	}

	if whatsOnChainOptions.BackOffMaximumJitterInterval != options.BackOffMaximumJitterInterval {
		t.Fatalf("expected value: %v got: %v", options.BackOffMaximumJitterInterval, whatsOnChainOptions.BackOffMaximumJitterInterval)
	}

	if whatsOnChainOptions.BackOffMaxTimeout != options.BackOffMaxTimeout {
		t.Fatalf("expected value: %v got: %v", options.BackOffMaxTimeout, whatsOnChainOptions.BackOffMaxTimeout)
	}

	if whatsOnChainOptions.DialerKeepAlive != options.DialerKeepAlive {
		t.Fatalf("expected value: %v got: %v", options.DialerKeepAlive, whatsOnChainOptions.DialerKeepAlive)
	}

	if whatsOnChainOptions.DialerTimeout != options.DialerTimeout {
		t.Fatalf("expected value: %v got: %v", options.DialerTimeout, whatsOnChainOptions.DialerTimeout)
	}

	if whatsOnChainOptions.RequestRetryCount != options.RequestRetryCount {
		t.Fatalf("expected value: %v got: %v", options.RequestRetryCount, whatsOnChainOptions.RequestRetryCount)
	}

	if whatsOnChainOptions.RequestTimeout != options.RequestTimeout {
		t.Fatalf("expected value: %v got: %v", options.RequestTimeout, whatsOnChainOptions.RequestTimeout)
	}

	if whatsOnChainOptions.TransportExpectContinueTimeout != options.TransportExpectContinueTimeout {
		t.Fatalf("expected value: %v got: %v", options.TransportExpectContinueTimeout, whatsOnChainOptions.TransportExpectContinueTimeout)
	}

	if whatsOnChainOptions.TransportIdleTimeout != options.TransportIdleTimeout {
		t.Fatalf("expected value: %v got: %v", options.TransportIdleTimeout, whatsOnChainOptions.TransportIdleTimeout)
	}

	if whatsOnChainOptions.TransportMaxIdleConnections != options.TransportMaxIdleConnections {
		t.Fatalf("expected value: %v got: %v", options.TransportMaxIdleConnections, whatsOnChainOptions.TransportMaxIdleConnections)
	}

	if whatsOnChainOptions.TransportTLSHandshakeTimeout != options.TransportTLSHandshakeTimeout {
		t.Fatalf("expected value: %v got: %v", options.TransportTLSHandshakeTimeout, whatsOnChainOptions.TransportTLSHandshakeTimeout)
	}
}
