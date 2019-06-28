package main

import (
	"net/http"

	"encoding/json"
	"log"
)

func userlogin(writer http.ResponseWriter, request *http.Request) {
	//数据库操作
	//逻辑处理
	//restapi json
	//获得参数
	//解析参数


	username := request.PostForm.Get("username")
	password := request.PostForm.Get("passwd")
	loginok := false
	if (username == "jilin7105" && password =="123qwe") {
		loginok =true
	}

	if(loginok){
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] ="test"
		Resp(writer,0,data,"")
	}else{
		Resp(writer,-1,nil,"密码不正确")
	}

}

type H struct {
	Code int
	Msg string
	Data interface{}
}

func Resp(w http.ResponseWriter ,code int ,data interface{} ,msg string){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	h := H{
		Code:code,
		Data:data,
		Msg:msg,
	}

	ret,err :=json.Marshal(h)
	if err!=nil{
		log.Panicln(err.Error())
	}
	w.Write([]byte(ret))
}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login",userlogin)

	//启动web服务器
	http.ListenAndServe(":9099",nil)
}
