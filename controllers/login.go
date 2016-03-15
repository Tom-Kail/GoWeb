package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

/*
func (this *LoginController) Prepare() {
	userId := this.GetSession("user_id").(string)
	if userId == nil {
		this.Redirect("/login", 302)
	}
}
*/
func (this *LoginController) Logout() {
	this.Data["IsLogin"] = false
	this.DelSession("user_name")
	this.DelSession("IsLogin")
	this.DelSession("user_id")
	logoutInfo := struct {
		Status bool
		Msg    string
	}{false, ""}
	this.Data["json"] = &logoutInfo
	this.ServeJson()
	this.Redirect("/", 302)
}
func (this *LoginController) Post() {
	type LoginInfo struct {
		LoginName  string `json:"login_name"`
		Password   string `json:"password"`
		VerifyCode string `json:"verify_code"`
	}
	var loginInfo LoginInfo
	loginJsonInfo := this.GetString("data")
	json.Unmarshal([]byte(loginJsonInfo), loginInfo)

	rsp := struct {
		Status bool
		Msg    string
	}{false, ""}

	// 验证用户名及密码
	if status, msg, uname, uid := IsLoginRight(loginInfo.LoginName, loginInfo.Password, loginInfo.VerifyCode); status == true {

		//this.SetSession("login_name", loginName)
		this.SetSession("user_name", uname)
		this.SetSession("user_id", uid)
		this.SetSession("IsLogin", true)
		rsp.Status = false //登陆成功
		rsp.Msg = ""
		this.Data["json"] = &rsp
	} else {
		rsp.Status = true //登陆失败
		rsp.Msg = msg     //失败原因
		this.Data["json"] = &rsp
		this.SetSession("IsLogin", false)
	}
	this.ServeJson()
}
func IsLoginRight(loginName string, pwd string, VerifyCode string) (bool, string, string, string) {
	o := orm.NewOrm()
	o.Using("default")
	type User struct {
		UserId    string `field:"user_id"`
		UserName  string `field:"user_name"`
		LoginName string `field:"login_name"`
		Password  string `field:"password"`
	}
	var user User
	err := o.Raw("select * from table_user_info where login_name=? and password=?", loginName, pwd).QueryRow(&user)
	if err == nil && user.UserId != "" {
		return true, "", user.UserName, user.UserId
	} else {
		return false, "用户不存在或密码错误", "", ""
	}
}
