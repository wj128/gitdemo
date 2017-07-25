package controllers

import (
	//_"userCenter/models"
	"userCenter/models"
	"net/http"
	"fmt"
	"strconv"//string与int的转换
	"encoding/json"
	"regexp"
)

/*
 *  Post方法主要是增加信息，这里调用了models包 下面的User结构体和UserAdd方法（首字母必须大写），这里将获取到了URL参数封装到了
 *  User结构体中，通过在UserAdd方法里面传入一个指针，通过”指针.属性“获取该值
 *
 */

func Post(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
 	U_name := r.Form.Get("user_name")
	U_psw := r.Form.Get("user_password")
	bool1,_ := regexp.MatchString("^[0-9]|[a-z]|[A-Z]$", U_name)
	bool2,_ := regexp.MatchString("^[0-9]|[a-z]|[A-Z][.][_]$", U_psw)
	if U_name != "" && U_psw != "" && bool1 == true && bool2 == true {
		user := &models.User{User_name:U_name,User_password:U_psw}
		models.UserAdd(user)
	}else if U_name == "" || U_psw == "" {
		fmt.Fprintf(w, "用户名或密码为空或参数名不正确!")
	}else if bool1 != true || bool2 != true {
		fmt.Fprintf(w, "用户名或密码为非法字符!")
	}else{
		fmt.Fprintf(w, "请重新检查URL是否正确!")
	}
}



/*
 *  删除操作用了数据类型转换：将 r.Form.Get("user_id")得到的string类型转换为
 *  int类型的数据，这里使用了strconv包下面的Atoi()方法，该方法传入一个string类型，
 *  返回转换后的int类型和一个error类型
 */
func Delete(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	U_id,_ := strconv.Atoi(r.Form.Get("user_id"))//返回值是int和error
	models.UserDelete(U_id)
}

 
/*
 *  更新资源
 *
 */
func Put(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	U_id,_ := strconv.Atoi(r.Form.Get("user_id"))
	U_name := r.Form.Get("user_name")
	U_psw := r.Form.Get("user_password")

	user := &models.User{User_id:U_id,User_name:U_name,User_password:U_psw}
	models.UserUpdate(user)
}

/*
 *  查找资源
 *
 */
func Get(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
/*	//获取完整的URL
	scheme := "http://"
    if r.TLS != nil {
        scheme = "https://"
    }
	//http://192.168.33.10:9090/query?user_id=4
	fmt.Println(scheme+r.Host+r.RequestURI)

	var v []string
	var i int
	for _,v = range r.Form {
		L := len(v)
		for i=0;i<L;i++ {
			//fmt.Println(v[i])
		}
	}
	fmt.Println(v[i])
*/

	U_id,_ := strconv.Atoi(r.Form.Get("user_id"))
	fmt.Println(U_id,"pppppppp")
	//json格式输出数据到页面
	user1 := models.UserQuery(U_id)
	json.NewEncoder(w).Encode(user1)
}



func Error(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "请输入正确的URL!")
}
