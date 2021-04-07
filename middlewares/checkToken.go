package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseWithRecorder struct{
	ResponseWriter  http.ResponseWriter
	statusCode     int
	body           bytes.Buffer
}

func AccessMiddle(f http.Handler)http.Handler{
	// 创建一个新的handler包装http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))

		//write
		r.Body = rdr

		wc := &ResponseWithRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           bytes.Buffer{},
		}
		// 调用下一个中间件或者最终的handler处理程序
		f.ServeHTTP(w, r)

		fmt.Println(wc)
	})
}