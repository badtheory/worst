package worst

import (
	"github.com/badtheory/fuse"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
	"net/http"
)

type Security struct {

	// AllowedOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters
	// (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penality.
	// Only one wildcard can be used per origin.
	// Default value is ["*"]
	AllowedOrigins []string `default:"[\"*\"]"`

	// AllowOriginFunc is a custom function to validate the origin. It take the origin
	// as argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowedOrigins is ignored.
	AllowOriginFunc func(r *http.Request, origin string) bool

	// AllowedMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET and POST)
	AllowedMethods []string `default:"[\"GET\", \"POST\"]"`

	// AllowedHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	// If the special "*" value is present in the list, all headers will be allowed.
	// Default value is [] but "Origin" is always appended to the list.
	AllowedHeaders []string

	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposedHeaders []string

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached
	MaxAge int

	// OptionsPassthrough instructs preflight to let other potential next handlers to
	// process the OPTIONS method. Turn this on if your application handles OPTIONS.
	OptionsPassthrough bool

	// Debugging flag adds additional output to debug server side CORS issues
	Debug bool

	// If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
	BrowserXssFilter bool `default:"true"` // nolint: golint

	// If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
	ContentTypeNosniff bool

	// If ForceSTSHeader is set to true, the STS header will be added even when the connection is HTTP. Default is false.
	ForceSTSHeader bool

	// If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
	FrameDeny bool `default:"true"`

	// When developing, the AllowedHosts, SSL, and STS options can cause some unwanted effects. Usually testing happens on http, not https, and on localhost, not your production domain... so set this to true for dev environment.
	// If you would like your development environment to mimic production with complete Host blocking, SSL redirects, and STS headers, leave this as false. Default if false.
	IsDevelopment bool

	// If SSLRedirect is set to true, then only allow https requests. Default is false.
	SSLRedirect bool

	// If SSLForceHost is true and SSLHost is set, requests will be forced to use SSLHost even the ones that are already using SSL. Default is false.
	SSLForceHost bool

	// If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
	SSLTemporaryRedirect bool

	// If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
	STSIncludeSubdomains bool `default:"true"`

	// If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
	STSPreload bool `default:"true"`

	// ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
	ContentSecurityPolicy string `default:"script-src $NONCE"`

	// ContentSecurityPolicyReportOnly allows the Content-Security-Policy-Report-Only header value to be set with a custom value. Default is "".
	ContentSecurityPolicyReportOnly string

	// CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
	CustomBrowserXssValue string // nolint: golint

	// Passing a template string will replace `$NONCE` with a dynamic nonce value of 16 bytes for each request which can be later retrieved using the Nonce function.
	// Eg: script-src $NONCE -> script-src 'nonce-a2ZobGFoZg=='
	// CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
	CustomFrameOptionsValue string

	// PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
	PublicKey string

	// ReferrerPolicy allows sites to control when browsers will pass the Referer header to other sites. Default is "".
	ReferrerPolicy string

	// FeaturePolicy allows to selectively enable and disable use of various browser features and APIs. Default is "".
	FeaturePolicy string

	// SSLHost is the host name that is used to redirect http requests to https. Default is "", which indicates to use the same host.
	SSLHost string

	// AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
	AllowedHosts []string

	// AllowedHostsAreRegex determines, if the provided slice contains valid regular expressions. If this flag is set to true, every request's
	// host will be checked against these expressions. Default is false for backwards compatibility.
	AllowedHostsAreRegex bool

	// HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
	HostsProxyHeaders []string

	// SSLHostFunc is a function pointer, the return value of the function is the host name that has same functionality as `SSHost`. Default is nil.
	// If SSLHostFunc is nil, the `SSLHost` option will be used.
	SSLHostFunc *secure.SSLHostFunc

	// SSLProxyHeaders is set of header keys with associated values that would indicate a valid https request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
	SSLProxyHeaders map[string]string

	// STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
	STSSeconds int64

	// ExpectCTHeader allows the Expect-CT header value to be set with a custom value. Default is "".
	ExpectCTHeader string

}

func (s Security) fuse() (cors.Options, secure.Options) {
	co := fuse.Fuse(&s, &cors.Options{}).(cors.Options)
	so := fuse.Fuse(&s, &secure.Options{}).(secure.Options)
	return co, so
}
