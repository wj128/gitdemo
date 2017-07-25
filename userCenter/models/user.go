package models

import(
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
)

var db = &sql.DB{}

func init(){
	db,_ = sql.Open("mysql","root:123456@tcp(172.16.1.71:3306)/test?charset=utf8")
}

type User struct {
	User_id int
	User_name string
	User_password string 
}

/*
 *增加
 */
func UserAdd(user *User){
	stmt,err := db.Prepare("insert into user(name,password)value(?,?)")
	checkErr(err)
	rs,err := stmt.Exec(user.User_name,user.User_password)
	checkErr(err)
	id,err := rs.LastInsertId()
	fmt.Println("增加的数据的id：",id)
}

/*
 *删除
 */
func UserDelete(id int){
	stmt,err := db.Prepare("delete from user where id=?")
	checkErr(err)
	stmt.Exec(id)
	fmt.Println("删除的数据的id：",id)
}

/*
 *修改
 */
func UserUpdate(user *User){
	userId := user.User_id
	userName := user.User_name
	userPassword := user.User_password

	if(userName == "" && userPassword != ""){
		stmt,err := db.Prepare("update user set password=? where id=? ")
		checkErr(err)
		stmt.Exec(userPassword,userId)
		fmt.Println("修改了数据的密码，他的id是：",userId)
	}else if(userName != "" && userPassword == ""){
		stmt,err := db.Prepare("update user set name=? where id=? ")
		checkErr(err)
		stmt.Exec(userName,userId)
		fmt.Println("修改了数据的用户名，他的id是：",userId)
	}else if(userName != "" && userPassword != ""){
		stmt,err := db.Prepare("update user set name=?,password=? where id=? ")
		checkErr(err)
		stmt.Exec(userName,userPassword,userId)
		fmt.Println("修改了数据的用户名和密码，他的id是：",userId)
	}
}

 /*
 *查询
 */
func UserQuery(id int) (user User) {
/*
	stmt,err := db.Prepare("select * from user where id=?")
	checkErr(err)
	rs,err := stmt.Exec(id)
	checkErr(err)
	fmt.Println("-------------")
	fmt.Println("查询的用户的id是：",id)
	rows := rs.Query()
	fmt.Println("++++++++++++++++++++++")
	checkErr(err)
	fmt.Println("kkkkkkkk")
*/
	var u_id int
	var u_name string
	var u_password string
	rs := db.QueryRow("select * from user where id = ?", id)
	err := rs.Scan(&u_id,&u_name,&u_password)
	checkErr(err)
	fmt.Println(u_id)
	fmt.Println(u_name)
	fmt.Println(u_password)
	
	user1 := User{User_id:u_id,User_name:u_name,User_password:u_password}
	return user1

/*	for rows.Next(){
		var id int
		var name string
		var password string
		err := rows.Scan(&id,&name,&password)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(password)
	}
*/

}

 /*
 *异常处理
 */
func checkErr(err error){
	if err != nil {
		log.Println(err)
	}
}