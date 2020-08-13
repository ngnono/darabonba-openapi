// This file is auto-generated, don't edit it. Thanks.
/**
 * This is for OpenApi SDK
 */
package client

import (
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	credential "github.com/aliyun/credentials-go/credentials"
)

/**
 * Model for initing client
 */
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// region id
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// read timeout
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// credential
	Credential credential.Credential `json:"credential,omitempty" xml:"credential,omitempty"`
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// network for endpoint
	Network *string `json:"network,omitempty" xml:"network,omitempty"`
	// user agent
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// suffix for endpoint
	Suffix *string `json:"suffix,omitempty" xml:"suffix,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// endpoint type
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// OpenPlatform endpoint
	OpenPlatformEndpoint *string `json:"openPlatformEndpoint,omitempty" xml:"openPlatformEndpoint,omitempty"`
	// Deprecated
	// credential type
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetRegionId(v string) *Config {
	s.RegionId = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetCredential(v credential.Credential) *Config {
	s.Credential = v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetNetwork(v string) *Config {
	s.Network = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSuffix(v string) *Config {
	s.Suffix = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetEndpointType(v string) *Config {
	s.EndpointType = &v
	return s
}

func (s *Config) SetOpenPlatformEndpoint(v string) *Config {
	s.OpenPlatformEndpoint = &v
	return s
}

func (s *Config) SetType(v string) *Config {
	s.Type = &v
	return s
}

type OpenApiRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   map[string]*string     `json:"query,omitempty" xml:"query,omitempty"`
	Body    map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenApiRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenApiRequest) GoString() string {
	return s.String()
}

func (s *OpenApiRequest) SetHeaders(v map[string]*string) *OpenApiRequest {
	s.Headers = v
	return s
}

func (s *OpenApiRequest) SetQuery(v map[string]*string) *OpenApiRequest {
	s.Query = v
	return s
}

func (s *OpenApiRequest) SetBody(v map[string]interface{}) *OpenApiRequest {
	s.Body = v
	return s
}

type Client struct {
	Endpoint             *string
	RegionId             *string
	Protocol             *string
	UserAgent            *string
	EndpointRule         *string
	EndpointMap          map[string]*string
	Suffix               *string
	ReadTimeout          *int
	ConnectTimeout       *int
	HttpProxy            *string
	HttpsProxy           *string
	Socks5Proxy          *string
	Socks5NetWork        *string
	NoProxy              *string
	Network              *string
	ProductId            *string
	MaxIdleConns         *int
	EndpointType         *string
	OpenPlatformEndpoint *string
	Credential           credential.Credential
}

/**
 * Init client with Config
 * @param config config contains the necessary information to create a client
 */
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(tea.ToMap(config))) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	if !tea.BoolValue(util.Empty(config.AccessKeyId)) && !tea.BoolValue(util.Empty(config.AccessKeySecret)) {
		if !tea.BoolValue(util.Empty(config.SecurityToken)) {
			config.Type = tea.String("sts")
		} else {
			config.Type = tea.String("access_key")
		}

		credentialConfig := &credential.Config{
			AccessKeyId:     config.AccessKeyId,
			Type:            config.Type,
			AccessKeySecret: config.AccessKeySecret,
			SecurityToken:   config.SecurityToken,
		}
		client.Credential, _err = credential.NewCredential(credentialConfig)
		if _err != nil {
			return _err
		}

	} else if !tea.BoolValue(util.IsUnset(config.Credential)) {
		client.Credential = config.Credential
	}

	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.RegionId = config.RegionId
	client.UserAgent = config.UserAgent
	client.ReadTimeout = config.ReadTimeout
	client.ConnectTimeout = config.ConnectTimeout
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = config.MaxIdleConns
	return nil
}

/**
 * Encapsulate the request and invoke the network
 * @param action api name
 * @param version product version
 * @param protocol http or https
 * @param method e.g. GET
 * @param authType authorization type e.g. AK
 * @param bodyType response body type e.g. String
 * @param request object of OpenApiRequest
 * @param runtime which controls some details of call api, such as retry times
 * @return the response
 */
func (client *Client) DoRPCRequest(action *string, version *string, protocol *string, method *string, authType *string, bodyType *string, request *OpenApiRequest, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":      "retry",
		"readTimeout":    tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout": tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":      tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":     tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":        tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":   tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = tea.String("/")
			request_.Query = tea.Merge(map[string]*string{
				"Action":         action,
				"Format":         tea.String("json"),
				"Version":        version,
				"Timestamp":      openapiutil.GetTimestamp(),
				"SignatureNonce": util.GetNonce(),
			}, request.Query)
			// endpoint is setted in product client
			request_.Headers = map[string]*string{
				"host":          client.Endpoint,
				"x-acs-version": version,
				"x-acs-action":  action,
				"user-agent":    client.GetUserAgent(),
			}
			if !tea.BoolValue(util.IsUnset(request.Body)) {
				tmp := util.AnyifyMapValue(openapiutil.Query(request.Body))
				request_.Body = tea.ToReader(util.ToFormString(tmp))
				request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			}

			if !tea.BoolValue(util.EqualString(authType, tea.String("Anonymous"))) {
				accessKeyId, _err := client.GetAccessKeyId()
				if _err != nil {
					return _result, _err
				}

				accessKeySecret, _err := client.GetAccessKeySecret()
				if _err != nil {
					return _result, _err
				}

				securityToken, _err := client.GetSecurityToken()
				if _err != nil {
					return _result, _err
				}

				if !tea.BoolValue(util.Empty(securityToken)) {
					request_.Query["SecurityToken"] = securityToken
				}

				request_.Query["SignatureMethod"] = tea.String("HMAC-SHA1")
				request_.Query["SignatureVersion"] = tea.String("1.0")
				request_.Query["AccessKeyId"] = accessKeyId
				signedParam := tea.Merge(request_.Query,
					openapiutil.Query(request.Body))
				request_.Query["Signature"] = openapiutil.GetRPCSignature(signedParam, request_.Method, accessKeySecret)
			}

			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			if tea.BoolValue(util.Is4xx(response_.StatusCode)) || tea.BoolValue(util.Is5xx(response_.StatusCode)) {
				_res, _err := util.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				err := util.AssertAsMap(_res)
				_err = tea.NewSDKError(map[string]interface{}{
					"message": err["Message"],
					"data":    err,
					"code":    err["Code"],
				})
				return _result, _err
			}

			if tea.BoolValue(util.EqualString(bodyType, tea.String("binary"))) {
				resp := map[string]interface{}{
					"body":    response_.Body,
					"headers": response_.Headers,
				}
				_result = resp
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("byte"))) {
				byt, _err := util.ReadAsBytes(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    byt,
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("string"))) {
				str, _err := util.ReadAsString(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    tea.StringValue(str),
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("json"))) {
				obj, _err := util.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				res := util.AssertAsMap(obj)
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    res,
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else {
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]map[string]*string{
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			}

		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

/**
 * Encapsulate the request and invoke the network
 * @param action api name
 * @param version product version
 * @param protocol http or https
 * @param method e.g. GET
 * @param authType authorization type e.g. AK
 * @param pathname pathname of every api
 * @param bodyType response body type e.g. String
 * @param request object of OpenApiRequest
 * @param runtime which controls some details of call api, such as retry times
 * @return the response
 */
func (client *Client) DoROARequest(action *string, version *string, protocol *string, method *string, authType *string, pathname *string, bodyType *string, request *OpenApiRequest, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":      "retry",
		"readTimeout":    tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout": tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":      tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":     tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":        tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":   tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Headers = tea.Merge(map[string]*string{
				"date":                    util.GetDateUTCString(),
				"host":                    client.Endpoint,
				"accept":                  tea.String("application/json"),
				"x-acs-signature-nonce":   util.GetNonce(),
				"x-acs-signature-method":  tea.String("HMAC-SHA1"),
				"x-acs-signature-version": tea.String("1.0"),
				"x-acs-version":           version,
				"x-acs-action":            action,
				"user-agent":              util.GetUserAgent(client.UserAgent),
			}, request.Headers)
			if !tea.BoolValue(util.IsUnset(request.Body)) {
				request_.Body = tea.ToReader(util.ToJSONString(request.Body))
				request_.Headers["content-type"] = tea.String("application/json; charset=UTF-8;")
			}

			if !tea.BoolValue(util.IsUnset(request.Query)) {
				request_.Query = request.Query
			}

			if !tea.BoolValue(util.EqualString(authType, tea.String("Anonymous"))) {
				accessKeyId, _err := client.Credential.GetAccessKeyId()
				if _err != nil {
					return _result, _err
				}

				accessKeySecret, _err := client.Credential.GetAccessKeySecret()
				if _err != nil {
					return _result, _err
				}

				securityToken, _err := client.Credential.GetSecurityToken()
				if _err != nil {
					return _result, _err
				}

				if !tea.BoolValue(util.Empty(securityToken)) {
					request_.Headers["x-acs-accesskey-id"] = accessKeyId
					request_.Headers["x-acs-security-token"] = securityToken
				}

				stringToSign := openapiutil.GetStringToSign(request_)
				request_.Headers["authorization"] = tea.String("acs " + tea.StringValue(accessKeyId) + ":" + tea.StringValue(openapiutil.GetROASignature(stringToSign, accessKeySecret)))
			}

			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			if tea.BoolValue(util.EqualNumber(response_.StatusCode, tea.Int(204))) {
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]map[string]*string{
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			}

			if tea.BoolValue(util.Is4xx(response_.StatusCode)) || tea.BoolValue(util.Is5xx(response_.StatusCode)) {
				_res, _err := util.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				err := util.AssertAsMap(_res)
				_err = tea.NewSDKError(map[string]interface{}{
					"message": err["Message"],
					"data":    err,
					"code":    err["Code"],
				})
				return _result, _err
			}

			if tea.BoolValue(util.EqualString(bodyType, tea.String("binary"))) {
				resp := map[string]interface{}{
					"body":    response_.Body,
					"headers": response_.Headers,
				}
				_result = resp
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("byte"))) {
				byt, _err := util.ReadAsBytes(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    byt,
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("string"))) {
				str, _err := util.ReadAsString(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    tea.StringValue(str),
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("json"))) {
				obj, _err := util.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				res := util.AssertAsMap(obj)
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    res,
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else {
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]map[string]*string{
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			}

		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

/**
 * Encapsulate the request and invoke the network with form body
 * @param action api name
 * @param version product version
 * @param protocol http or https
 * @param method e.g. GET
 * @param authType authorization type e.g. AK
 * @param pathname pathname of every api
 * @param bodyType response body type e.g. String
 * @param request object of OpenApiRequest
 * @param runtime which controls some details of call api, such as retry times
 * @return the response
 */
func (client *Client) DoROARequestWithForm(action *string, version *string, protocol *string, method *string, authType *string, pathname *string, bodyType *string, request *OpenApiRequest, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":      "retry",
		"readTimeout":    tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout": tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":      tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":     tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":        tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":   tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Headers = tea.Merge(map[string]*string{
				"date":                    util.GetDateUTCString(),
				"host":                    client.Endpoint,
				"accept":                  tea.String("application/json"),
				"x-acs-signature-nonce":   util.GetNonce(),
				"x-acs-signature-method":  tea.String("HMAC-SHA1"),
				"x-acs-signature-version": tea.String("1.0"),
				"x-acs-version":           version,
				"x-acs-action":            action,
				"user-agent":              util.GetUserAgent(client.UserAgent),
			}, request.Headers)
			if !tea.BoolValue(util.IsUnset(request.Body)) {
				request_.Body = tea.ToReader(openapiutil.ToForm(request.Body))
				request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			}

			if !tea.BoolValue(util.IsUnset(request.Query)) {
				request_.Query = request.Query
			}

			if !tea.BoolValue(util.EqualString(authType, tea.String("Anonymous"))) {
				accessKeyId, _err := client.Credential.GetAccessKeyId()
				if _err != nil {
					return _result, _err
				}

				accessKeySecret, _err := client.Credential.GetAccessKeySecret()
				if _err != nil {
					return _result, _err
				}

				securityToken, _err := client.Credential.GetSecurityToken()
				if _err != nil {
					return _result, _err
				}

				if !tea.BoolValue(util.Empty(securityToken)) {
					request_.Headers["x-acs-accesskey-id"] = accessKeyId
					request_.Headers["x-acs-security-token"] = securityToken
				}

				stringToSign := openapiutil.GetStringToSign(request_)
				request_.Headers["authorization"] = tea.String("acs " + tea.StringValue(accessKeyId) + ":" + tea.StringValue(openapiutil.GetROASignature(stringToSign, accessKeySecret)))
			}

			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			if tea.BoolValue(util.EqualNumber(response_.StatusCode, tea.Int(204))) {
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]map[string]*string{
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			}

			if tea.BoolValue(util.Is4xx(response_.StatusCode)) || tea.BoolValue(util.Is5xx(response_.StatusCode)) {
				_res, _err := util.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				err := util.AssertAsMap(_res)
				_err = tea.NewSDKError(map[string]interface{}{
					"message": err["Message"],
					"data":    err,
					"code":    err["Code"],
				})
				return _result, _err
			}

			if tea.BoolValue(util.EqualString(bodyType, tea.String("binary"))) {
				resp := map[string]interface{}{
					"body":    response_.Body,
					"headers": response_.Headers,
				}
				_result = resp
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("byte"))) {
				byt, _err := util.ReadAsBytes(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    byt,
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("string"))) {
				str, _err := util.ReadAsString(response_.Body)
				if _err != nil {
					return _result, _err
				}

				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    tea.StringValue(str),
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else if tea.BoolValue(util.EqualString(bodyType, tea.String("json"))) {
				obj, _err := util.ReadAsJSON(response_.Body)
				if _err != nil {
					return _result, _err
				}

				res := util.AssertAsMap(obj)
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]interface{}{
					"body":    res,
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			} else {
				_result = make(map[string]interface{})
				_err = tea.Convert(map[string]map[string]*string{
					"headers": response_.Headers,
				}, &_result)
				return _result, _err
			}

		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

/**
 * Get user agent
 * @return user agent
 */
func (client *Client) GetUserAgent() (_result *string) {
	userAgent := util.GetUserAgent(client.UserAgent)
	_result = userAgent
	return _result
}

/**
 * Get accesskey id by using credential
 * @return accesskey id
 */
func (client *Client) GetAccessKeyId() (_result *string, _err error) {
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		return _result, _err
	}

	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	_result = accessKeyId
	return _result, _err
}

/**
 * Get accesskey secret by using credential
 * @return accesskey secret
 */
func (client *Client) GetAccessKeySecret() (_result *string, _err error) {
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		return _result, _err
	}

	secret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	_result = secret
	return _result, _err
}

/**
 * Get security token by using credential
 * @return security token
 */
func (client *Client) GetSecurityToken() (_result *string, _err error) {
	if tea.BoolValue(util.IsUnset(client.Credential)) {
		return _result, _err
	}

	token, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	_result = token
	return _result, _err
}

/**
 * If inputValue is not null, return it or return defaultValue
 * @param inputValue  users input value
 * @param defaultValue default value
 * @return the final result
 */
func DefaultAny(inputValue interface{}, defaultValue interface{}) (_result interface{}) {
	if tea.BoolValue(util.IsUnset(inputValue)) {
		_result = defaultValue
		return _result
	}

	_result = inputValue
	return _result
}

/**
 * If the endpointRule and config.endpoint are empty, throw error
 * @param config config contains the necessary information to create a client
 */
func (client *Client) CheckConfig(config *Config) (_err error) {
	if tea.BoolValue(util.Empty(client.EndpointRule)) && tea.BoolValue(util.Empty(config.Endpoint)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config.endpoint' can not be empty",
		})
		return _err
	}

	return _err
}