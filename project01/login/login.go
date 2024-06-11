package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
var users = make(map[string]User) //存储用户信息的map

func main(){
	
	http.HandleFunc("/login", loginHandler) 
	http.HandleFunc("/register", registerHandler)
	fmt.Println("Listening on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err!= nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		var user User
		//解析请求体中的json数据
		err := json.NewDecoder(r.Body).Decode(&user)
		if err!= nil {
			//解析失败
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _,ok := users[user.Username]; ok{
			//用户名已存在
			http.Error(w, "用户已经存在", http.StatusBadRequest)
			return
		}
		//将用户信息存储到map中
		users[user.Username] = user

		response := map[string]string{"message": "注册成功"}
		//返回json格式的响应
		json.NewEncoder(w).Encode(response)

	}else{
		http.Error(w, "请求方法错误", http.StatusMethodNotAllowed)
	}

}
func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST"{
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		storedUser,ok := users[user.Username]
		if !ok || storedUser.Password != user.Password{
			http.Error(w, "用户名或密码错误", http.StatusUnauthorized)
			return
		}
		response := map[string]string{"message": "登录成功"}
		json.NewEncoder(w).Encode(response)
	}else{

		http.Error(w, "请求方法错误", http.StatusMethodNotAllowed)
	}

}