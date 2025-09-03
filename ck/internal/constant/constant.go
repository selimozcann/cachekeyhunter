package constants

// Base
const (
	UserAgentHeader = "User-Agent"
	UserAgentValue  = "CacheKeyHunter/1.0"

	DefaultTimeoutSeconds = 15

	HeaderXCache        = "X-Cache"
	HeaderCFCacheStatus = "CF-Cache-Status"
	HeaderAge           = "Age"
	HeaderVary          = "Vary"
	Hit                 = "HIT"
)

// Header keys
const (
	HeaderXForwardedHost  = "X-Forwarded-Host"
	HeaderXForwardedProto = "X-Forwarded-Proto"
	HeaderForwarded       = "Forwarded"
)

// Default values
const (
	DefaultExampleDomain   = "example.com"
	DefaultProto           = "http"
	DefaultForwardedPrefix = "host="
)

const (
	SeverityHIGH   = "HIGH"
	SeverityMedium = "MEDIUM"
	SeverityLow    = "LOW"
)
