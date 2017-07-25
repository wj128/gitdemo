package main

import (
	"userCenter/controllers"
	"net/http"
)

//	访问端口   http://192.168.33.10:9090/
func main(){
	http.HandleFunc("/", controllers.Error)
 	http.HandleFunc("/add", controllers.Post)      //增加资源，设置访问的路由
	http.HandleFunc("/delete",controllers.Delete)  //删除资源（通过id）
	http.HandleFunc("/update",controllers.Put)     //更新（修改）资源 
	http.HandleFunc("/query",controllers.Get)      //查找资源（通过id）
	http.ListenAndServe(":9090", nil)             //设置监听的端口
}


