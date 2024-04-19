package alipay

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Config struct {
	Gateway         string
	AppID           string
	AppKey          string
	AlipayPublicKey string
}

func LoadPrivateKey(privateKeyPem string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}
	return key, nil
}

func LoadPublicKey(publicKeyPem string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPem))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}
	return publicKey, nil
}

type Client struct {
	client *http.Client
	config *Config
}

type Response = map[string]interface{}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		client: http.DefaultClient,
	}
}

func (c *Client) createBaseParams() (base url.Values) {
	base = url.Values{}
	base.Add("app_id", c.config.AppID)
	base.Add("charset", "utf8")
	base.Add("format", "JSON")
	base.Add("sign_type", "RSA2")
	base.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	base.Add("version", "1.0")
	return base
}

// create signature
// @docs https://opendocs.alipay.com/common/057k53?pathHash=7b14a0af
func (c *Client) createSignature(params url.Values) (string, error) {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var strs []string
	for _, k := range keys {
		if k == "sign" {
			continue
		}
		value := params.Get(k)
		if value == "" {
			continue
		}
		strs = append(strs, fmt.Sprintf("%s=%s", k, value))
	}
	signContent := strings.Join(strs, "&")
	hash := sha256.New()
	hash.Write([]byte(signContent))
	hashed := hash.Sum(nil)
	paivateKey, err := LoadPrivateKey(c.config.AppKey)
	if err != nil {
		return "", fmt.Errorf("failed to load private key: %v", err)
	}
	signature, err := rsa.SignPKCS1v15(nil, paivateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}
	sign := base64.StdEncoding.EncodeToString(signature)
	return sign, nil
}

func sortJSON(jsonObj map[string]interface{}) (string, error) {
	var sortedKeys []string
	for key := range jsonObj {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	var sortedStrs []string
	for _, key := range sortedKeys {
		value, err := json.Marshal(jsonObj[key])
		if err != nil {
			return "", fmt.Errorf("failed to marshal JSON value: %v", err)
		}
		sortedStrs = append(sortedStrs, fmt.Sprintf("\"%s\":%s", key, string(value)))
	}
	str := "{" + strings.Join(sortedStrs, ",") + "}"
	str = strings.ReplaceAll(str, "/", "\\/")
	return str, nil
}

// Verify signature of the response
// @docs https://opendocs.alipay.com/common/02mse7?pathHash=096e611e
func (c *Client) verifySignature(response map[string]interface{}, sign string) (bool, error) {
	signContent, err := sortJSON(response)
	if err != nil {
		return false, fmt.Errorf("failed to sort JSON: %v", err)
	}
	hash := sha256.New()
	hash.Write([]byte(signContent))
	hashed := hash.Sum(nil)
	// 对签名进行 Base64 解码
	decodedSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, fmt.Errorf("failed to decode sign: %v", err)
	}
	// 加载支付宝公钥
	alipayPublicKey, err := LoadPublicKey(c.config.AlipayPublicKey)
	if err != nil {
		return false, fmt.Errorf("failed to load Alipay public key: %v", err)
	}
	// 验证签名
	err = rsa.VerifyPKCS1v15(alipayPublicKey, crypto.SHA256, hashed, decodedSign)
	if err != nil {
		return false, fmt.Errorf("failed to verify sign: %v", err)
	}
	return true, nil
}

func (c *Client) Execute(method string, params interface{}) (response Response, err error) {
	qs := c.createBaseParams()
	bizContent, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %v", err)
	}
	qs.Add("method", method)
	qs.Add("biz_content", string(bizContent))

	sign, err := c.createSignature(qs)
	if err != nil {
		return nil, fmt.Errorf("failed to create signature: %v", err)
	}
	qs.Del("biz_content")
	qs.Add("sign", sign)
	body := url.Values{}
	body.Add("biz_content", string(bizContent))
	req, err := http.NewRequest(http.MethodPost, c.config.Gateway, bytes.NewBufferString(body.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.URL.RawQuery = qs.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	sign, ok := result["sign"].(string)
	if !ok {
		return nil, errors.New("no sign found in the response")
	}
	responseKey := strings.Replace(method, ".", "_", -1) + "_response"
	response = result[responseKey].(Response)
	if response["code"].(string) != "10000" {
		return nil, fmt.Errorf("failed to execute request: %v", response["msg"])
	}
	valid, err := c.verifySignature(response, sign)
	if !valid || err != nil {
		return nil, fmt.Errorf("failed to verify signature: %v", err)
	}
	return response, nil
}
