使用Beego开发
====
#一些有用的小技巧 
##beego orm 的自动映射
```SQL
    Create table user(user_id varchar(32),user_name varchar(32));
```

```Go
import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/astaxie/beego/session/mysql"
)
  
func init() {
    orm.RegisterDriver("mysql",orm.DR_MySQL)
    orm.RegisterDataBase("default","mysql","root:@(tcp192.168.1.10:3306)/orm_test?charset=utf8")
}
func main() {
    type User struct{
	    UserName string
	    UserId string
    }
    o := orm.NewOrm()
    o.Using("default")//默认是default，要用其他的可用using替代
    var users []User
    //table中的user_id自动映射到UserId,user_name自动映射到UserName
    num , _ := o.Raw("select * from user").QueryRows(&users)
    

}


```
##json in Go
　Go语言中用反引号创建的字符串为原生字符串`raw string`，可有多行组成。
  ```Go
  type User struct{
    UserId string `json:"user_id"`
    UserName string `json:"user_name"`  
  }
  ```
  
  json本质是一定格式的字符串`{"name":"tom","list":[{1}{2}]}`
  ```Go
    user := User{"111","tom"}
    jsonStr := json.Marshal(user)//string(jsonStr) = `{"user_id":"111","user_name":"tom"}`

  ```
