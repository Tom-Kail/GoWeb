package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Post() {
	//解析请求
	jsonStr := this.GetString("data")
	type RegisterInfo struct {
		UserName       string `json:"user_name"`
		LoginName      string `json:"login_name"`
		Password       string `json:"user_password"`
		Mail           string `json:"mail"`
		Phone          string `json:"phone"`
		ActivationCode string `json:"activation_code"`
	}
	regiInfo := RegisterInfo{}
	json.Unmarshal([]byte(jsonStr), &regiInfo)
	//检测用户名何登陆名是否已经存在
	status, msg := checkUserExists(regiInfo.UserName, regiInfo.LoginName)
	if status {
		//在本地数据库中注册用户
		o := orm.NewOrm()
		o.Using("default")
		rsp := struct {
			Status bool
			Msg    string
		}{true, ""}
		//role 1:admin 3:normal user
		_, err := o.Raw("insert into table_user_info(user_name,login_name,password,mail,phone,activation_code,role) values(?,?,?,?,?,?,?)", regiInfo.UserName, regiInfo.LoginName, regiInfo.Password, regiInfo.Mail, regiInfo.Phone, regiInfo.ActivationCode, 3).Exec()
		if err == nil {
			rsp.Status = false
			rsp.Msg = ""
		} else {
			rsp.Status = true
			rsp.Msg = "插入数据库不成功"
			beego.Error(err)
		}

		this.Data["json"] = &rsp
		this.ServeJson()

	} else {
		rsp := struct {
			Status bool
			Msg    string
		}{false, msg}
		this.Data["json"] = &rsp
		this.ServeJson()
	}

}

func checkUserExists(userName string, loginName string) (bool, string) {
	o := orm.NewOrm()
	o.Using("default")
	type User struct {
		UserId    string `orm:"user_id"`
		UserName  string `orm:"user_name"`
		LoginName string `orm:"login_name"`
	}
	var user User
	err := o.Raw("select login_name from table_user_info where login_name=?", user.LoginName).QueryRow(&user)
	if err != nil {
		beego.Error(err)
	} else if user.LoginName != "" {
		return false, "登陆名已被注册"
	}
	err = o.Raw("select user_name from table_user_info where user_name=?", user.UserName).QueryRow(&user)
	if err != nil {
		beego.Error(err)
	} else if user.UserName != "" {
		return false, "用户名已被注册"
	}
	return true, ""

}
