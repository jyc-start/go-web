package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi，this is home page")
}

// r.Body只能读一次，意味着你读了别人就不能读了
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	fmt.Fprintf(w, "read body first: %v", body)
	// 尝试再次读取，不会报错，但是什么也读不到
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}
	fmt.Fprintf(w, "read body second: %v", body)
}

// r.GetBody原则上可以多次读取，但是在原生的http.Request里面，这个是nil
// 需要别人设置好！！！
func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprint(w, "getBody Is Nil \n")
	} else {
		fmt.Fprint(w, "getBody not is Nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

// 有的key为空，调用前需要判断
func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

// r.Header会自动把首字母改成大写
func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "header is %v \n", r.Header)
}

// ParseForm()只返回err
func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v \n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "err=%v \n", err)
	}
	fmt.Fprintf(w, "after parse form %v \n", r.Form)
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/read/once", readBodyOnce)
	http.HandleFunc("/getBody", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/form", form)
	http.HandleFunc("/wholeUrl", wholeUrl)
	if err := http.ListenAndServe("localhost:8081", nil); err != nil {
		fmt.Println(err)
	}
}
