package main

import (
	"fmt"
	"net/http"
	"yxr/blog/pkg/setting"
	"yxr/blog/routers"
)

/*
*  @author liqiqiorz
*  @data 2020/5/15 19:10
 */
func main(){
    router:=routers.InitRouter()
	s:=&http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}