package httpclient

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

//Config : HTTP Client config
type Config struct {
	Header     map[string]string
	Method     string
	Data       map[string]interface{}
	Query      map[string]interface{}
	Timeout    time.Duration
	RequestURI string
}

//Get default get config
func Get() *Config {
	return &Config{
		Method: "GET",
		Header: map[string]string{},
		Data:   map[string]interface{}{},
		Query:  map[string]interface{}{},
	}
}

//Post default POST config
func Post() *Config {
	return &Config{
		Method: "POST",
		Header: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
		Data:  map[string]interface{}{},
		Query: map[string]interface{}{},
	}
}

//SetHeader for adding/updating Header
func (c *Config) SetHeader(name string, value string) *Config {
	c.Header[name] = value
	return c
}

//UnsetHeader for removing Header
func (c *Config) UnsetHeader(name string) *Config {
	delete(c.Header, name)
	return c
}

//SetData for set data
func (c *Config) SetData(name string, value interface{}) *Config {
	c.Data[name] = value
	return c
}

//SetQuery for set data
func (c *Config) SetQuery(name string, value string) *Config {
	c.Query[name] = value
	return c
}

//SetRequestURI for set data
func (c *Config) SetRequestURI(path string) *Config {
	c.RequestURI = path
	return c
}

//Do Request
func (c *Config) Do() (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.SetRequestURI(c.RequestURI)
	for key, val := range c.Query {
		req.URI().QueryArgs().Set(key, fmt.Sprintf("%v", val))
	}
	if val, ok := c.Header["Content-Type"]; ok {
		req.Header.SetContentType(val)
	}
	for key, val := range c.Header {
		if key == "Content-Type" {
			req.Header.SetContentType(val)
			if method := strings.ToLower(c.Method); method == "post" || method == "put" {
				switch contenttype := strings.ToLower(val); contenttype {
				case "application/json":
					v, _ := json.Marshal(c.Data)
					req.SetBody(v)
				default:
					args := req.PostArgs()
					args.Reset()
					for key2, val2 := range c.Data {
						args.Set(key2, fmt.Sprintf("%v", val2))
					}
				}
			}
		} else {
			req.Header.Set(key, val)
		}
	}

	// define web client request Method
	req.Header.SetMethod(c.Method)

	var timeOut = 30 * time.Second
	if c.Timeout != 0 {
		timeOut = c.Timeout
	}
	// DO GET request
	var err = fasthttp.DoTimeout(req, resp, timeOut)

	if err != nil {
		return nil, err
	}

	// add your logic code here , to handle response

	var out = fasthttp.AcquireResponse()
	resp.CopyTo(out)

	return out, nil
}
