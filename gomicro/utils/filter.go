package utils

import "net/http"

//自用，带error 这里的res req单纯只是接收参数
type FilterHandle func(res http.ResponseWriter,req *http.Request)error


type Filter struct {
	urlMap map[string]FilterHandle
}

func NewFilter ()*Filter{
	return &Filter{
		urlMap:make(map[string]FilterHandle),
	}
}
//符合http参数类型 不带error
type WebHandle func(rw http.ResponseWriter, req *http.Request)

func (f *Filter) RegisterUrl(url string,handleFunc FilterHandle){
	f.urlMap[url] = handleFunc
}

func (f *Filter) Handle(webHandle WebHandle)(func(res http.ResponseWriter,req *http.Request)){
	return func(res http.ResponseWriter, req *http.Request) {
		//执行满足条件的验证函数
		handle,ok := f.urlMap[req.RequestURI]
		if ok {
			if err := handle(res,req);err != nil {
				res.Write([]byte(err.Error()))
				return
			}
		}

		//实际http的HandleFunc放上面的匿名函数，且被执行，这里属于执行里面的代码,无关闭包和中间件
		webHandle(res,req)
	}
}