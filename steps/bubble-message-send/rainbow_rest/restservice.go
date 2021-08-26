package rainbow_rest

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Rest struct {
	host   string
	token  string
	client *resty.Client
}

func userAgent() string {
	return fmt.Sprintf("Golang Rainbow SDK/%s", "0.0.0")
}

func NewRest(host string) *Rest {
	r := Rest{host: host, client: resty.New()}
	r.client.SetDebug(true)
	r.client.SetHostURL(fmt.Sprintf("https://%s", host))
	r.client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   userAgent(),
	})
	return &r
}

func (r *Rest) SetToken(token string) {
	r.token = token
	headers := r.client.Header
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	for n, v := range r.client.Header {
		fmt.Println(n, ":", v)
	}
}

func (r *Rest) get(url string, pathparams map[string]string, params map[string]string, typ interface{}) (*resty.Response, error) {
	if typ == nil {
		return r.client.R().SetPathParams(pathparams).SetQueryParams(params).Get(url)
	} else {
		return r.client.R().SetPathParams(pathparams).SetQueryParams(params).SetResult(typ).Get(url)
	}
}

func (r *Rest) put(url string, pathParams map[string]string, urlParams map[string]string, body interface{}, typ interface{}) (*resty.Response, error) {
	if typ == nil {
		return r.client.R().SetPathParams(pathParams).SetQueryParams(urlParams).SetBody(body).Put(url)
	} else {
		return r.client.R().SetPathParams(pathParams).SetQueryParams(urlParams).SetBody(body).SetResult(typ).Put(url)
	}
}

func (r *Rest) post(url string, pathParams map[string]string, urlParams map[string]string, body interface{}, typ interface{}) (*resty.Response, error) {
	if typ == nil {
		return r.client.R().SetPathParams(pathParams).SetQueryParams(urlParams).SetBody(body).Post(url)
	} else {
		return r.client.R().SetPathParams(pathParams).SetQueryParams(urlParams).SetBody(body).SetResult(typ).Post(url)
	}
}

func (r *Rest) delete(url string, pathParams map[string]string, urlParams map[string]string) (*resty.Response, error) {
	return r.client.R().SetPathParams(pathParams).SetQueryParams(urlParams).Delete(url)

}
