package models

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
	"time"
	"github.com/satori/go.uuid"
	"strconv"
	"net/http"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

// Client sts client
type Client struct {
	AccessKeyId     string
	AccessKeySecret string
	RoleArn         string
	SessionName     string
}

// ServiceError sts service error
type ServiceError struct {
	Code       string
	Message    string
	RequestId  string
	HostId     string
	RawMessage string
	StatusCode int
}

// Credentials the credentials obtained by AssumedRole,
// used for the peration of Alibaba Cloud service.
type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	Expiration      time.Time
	SecurityToken   string
}

// AssumedRoleUser the user to AssumedRole
type AssumedRoleUser struct {
	Arn           string
	AssumedRoleId string
}

// Response the response of AssumeRole
type Response struct {
	Credentials     Credentials
	AssumedRoleUser AssumedRoleUser
	RequestId       string
}

// Error implement interface error
func (e *ServiceError) Error() string {
	return fmt.Sprintf("oss: service returned error: StatusCode=%d, ErrorCode=%s, ErrorMessage=%s, RequestId=%s",
		e.StatusCode, e.Code, e.Message, e.RequestId)
}

// NewClient New STS Client
func NewClient(accessKeyId, accessKeySecret, roleArn, sessionName string) *Client {
	return &Client{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		RoleArn:         roleArn,
		SessionName:     sessionName,
	}
}

const (
	// StsSignVersion sts sign version
	StsSignVersion = "1.0"
	// StsAPIVersion sts api version
	StsAPIVersion = "2015-04-01"
	// StsHost sts host
	StsHost = "https://sts.aliyuncs.com/"
	// TimeFormat time fomrat
	TimeFormat = "2006-01-02T15:04:05Z"
	// RespBodyFormat  respone body format
	RespBodyFormat = "JSON"
	// PercentEncode '/'
	PercentEncode = "%2F"
	// HTTPGet http get method
	HTTPGet = "GET"
)

// AssumeRole assume role
func (c *Client) AssumeRole(expiredTime uint) (*Response, error) {
	url, err := c.generateSignedURL(expiredTime)
	if err != nil {
		return nil, err
	}

	body, status, err := c.sendRequest(url)
	if err != nil {
		return nil, err
	}

	return c.handleResponse(body, status)
}

// Private function
func (c *Client) generateSignedURL(expiredTime uint) (string, error) {
	queryStr := "SignatureVersion=" + StsSignVersion
	queryStr += "&Format=" + RespBodyFormat
	queryStr += "&Timestamp=" + url.QueryEscape(time.Now().UTC().Format(TimeFormat))
	queryStr += "&RoleArn=" + url.QueryEscape(c.RoleArn)
	queryStr += "&RoleSessionName=" + c.SessionName
	queryStr += "&AccessKeyId=" + c.AccessKeyId
	queryStr += "&SignatureMethod=HMAC-SHA1"
	queryStr += "&Version=" + StsAPIVersion
	queryStr += "&Action=AssumeRole"
	uuidv4, _ := uuid.NewV4()
	queryStr += "&SignatureNonce=" + uuidv4.String()
	queryStr += "&DurationSeconds=" + strconv.FormatUint((uint64)(expiredTime), 10)

	// Sort query string
	queryParams, err := url.ParseQuery(queryStr)
	if err != nil {
		return "", err
	}
	result := queryParams.Encode()

	strToSign := HTTPGet + "&" + PercentEncode + "&" + url.QueryEscape(result)

	// Generate signature
	hashSign := hmac.New(sha1.New, []byte(c.AccessKeySecret+"&"))
	hashSign.Write([]byte(strToSign))
	signature := base64.StdEncoding.EncodeToString(hashSign.Sum(nil))

	// Build url
	assumeURL := StsHost + "?" + queryStr + "&Signature=" + url.QueryEscape(signature)

	return assumeURL, nil
}

func (c *Client) sendRequest(url string) ([]byte, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, err
}

func (c *Client) handleResponse(responseBody []byte, statusCode int) (*Response, error) {
	if statusCode != http.StatusOK {
		se := ServiceError{StatusCode: statusCode, RawMessage: string(responseBody)}
		err := json.Unmarshal(responseBody, &se)
		if err != nil {
			return nil, err
		}
		return nil, &se
	}

	resp := Response{}
	err := json.Unmarshal(responseBody, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
