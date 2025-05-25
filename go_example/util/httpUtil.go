package util

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type httpUtil struct {
}

var httpUtilInstance httpUtil

func GetInstanceByHttpUtil() *httpUtil {
	return &httpUtilInstance
}

// GetRequestIp 获取客户端 ip，绕过代理，会去除端口号
func (*httpUtil) GetRequestIp(req *http.Request) string {
	ip := func(req *http.Request) string {
		// 优先使用 X-Forwarded-For
		fIp := req.Header.Get("X-Forwarded-For")
		if "" != fIp && !strings.Contains(fIp, "[") {
			// x.x.x.x,xx.xx.x.x,x.x.x.xx ...
			if strings.Contains(fIp, ",") {
				ips := strings.Split(fIp, ",")
				return ips[0]
			}
			return fIp
		}

		rIp := req.RemoteAddr
		// RemoteAddr=[::1] or 127.0.0.1
		if "" != rIp && !strings.Contains(rIp, "[") && !strings.HasPrefix(rIp, "127.") {
			return rIp
		}

		xIp := req.Header.Get("X-Real-IP")
		if "" != xIp && !strings.Contains(xIp, "[") {
			return xIp
		}

		remoteAddr := req.Header.Get("Remote_addr")
		if "" != remoteAddr && !strings.Contains(remoteAddr, "[") {
			return remoteAddr
		}

		return req.RemoteAddr
	}(req)
	if strings.Contains(ip, ":") {
		return strings.Split(ip, ":")[0]
	}
	return strings.Trim(ip, " ")
}

func (*httpUtil) Post(url string, data []byte, header map[string]string) ([]byte, error) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if nil != err {
		return nil, err
	}

	if nil == header || 0 == len(header) {
		req.Header.Set("Content-Type", "application/json")
	} else {
		if _, isOk := header["Content-Type"]; !isOk {
			req.Header.Set("Content-Type", "application/json")
		}
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(fmt.Sprintf("StatusCode=%d", resp.StatusCode))
	}

	return body, nil
}

func (*httpUtil) CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		h.ServeHTTP(w, r)
	})
}

func (*httpUtil) Get(u string, header map[string]string,
	timeOutSecond uint, transport *http.Transport) ([]byte, error) {
	if 0 == timeOutSecond {
		timeOutSecond = 30
	}

	client := &http.Client{
		Timeout: time.Duration(timeOutSecond) * time.Second,
	}

	if nil != transport {
		client.Transport = transport
	}

	req, err := http.NewRequest("GET", u, nil)
	if nil != err {
		return nil, err
	}

	if nil != header {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(fmt.Sprintf("StatusCode=%d Status=%s", resp.StatusCode, resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}

func (*httpUtil) PostClient(u string, header map[string]string, inBody []byte,
	timeOutSecond uint, transport *http.Transport) ([]byte, error) {
	if 0 == timeOutSecond {
		timeOutSecond = 30
	}

	client := &http.Client{
		Timeout: time.Duration(timeOutSecond) * time.Second,
	}

	if nil != transport {
		client.Transport = transport
	}

	req, err := http.NewRequest("POST", u, bytes.NewReader(inBody))
	if nil != err {
		return nil, errors.New(fmt.Sprintf("url=%s err=%s", u, err.Error()))
	}

	if nil == header || 0 == len(header) {
		req.Header.Set("Content-Type", "application/json")
	} else {
		if _, isOk := header["Content-Type"]; !isOk {
			req.Header.Set("Content-Type", "application/json")
		}
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New(fmt.Sprintf("StatusCode=%d Status=%s", resp.StatusCode, resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, err
	}

	return body, nil
}
