package main

import (
	"fmt"
	"go_example/util"
	"io"
	"net/http"
)

const platformPubKeyG = "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERXRLWWdBWDVXS2h6R2ljTVNSdVl2VlB2QgpNeGpvN1JVbzROWk9CU2ZIRkt6ZnppUVFKcUtRK3dUd2p0UmdhTndyQzd5bEtGcGNxamlqTTBMU2VscHpoWmwrCnFxTmVacE4yaDJkMW5wQ0wzbVBPYjJZYjJyUHc4T01oZmZ0WTc1dWxXYVY1cXBWTW1WSitSZ3VKOVlDN2tIaW8KQU9tUVVkUWhqOStPTmdCZFV3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="

func handler(w http.ResponseWriter, r *http.Request) {
	platformPub, err := util.DecodeBase64(platformPubKeyG)
	if nil != err {
		panic(err)
	}

	sign := r.Header.Get("sign")
	println(sign)
	if "" == sign {
		http.Error(w, "Signature error", http.StatusForbidden)
		return
	}
	fmt.Println("sign: ", sign)
	signData, err := util.DecodeBase64(sign)
	if nil != err {
		println(err)
		http.Error(w, "Signature error", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodPost {
		// 读取 POST 请求数据
		body, err := io.ReadAll(r.Body)
		println(string(body))
		if err != nil {
			println(err)
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		// 验证签名
		isOk, err := util.NewRSAPemPKCS8().RsaVerySignWithSha256(body, signData, platformPub)
		if nil != err {
			println(err)
			http.Error(w, "Signature error", http.StatusForbidden)
			return
		}

		if !isOk {
			http.Error(w, "Signature inaccuracy", http.StatusForbidden)
			return
		}
		fmt.Println(string(body))
		fmt.Println("成功...!")
		w.WriteHeader(http.StatusOK)
		return
	}
	// 非 POST 请求
	//fmt.Fprintf(w, "This endpoint accepts POST requests only.")
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	// 注册路由
	http.HandleFunc("/", handler)

	fmt.Println(platformPubKeyG)

	// 启动 HTTP 服务
	err := http.ListenAndServe(":19900", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}
