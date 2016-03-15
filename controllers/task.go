package controllers

import (
	"WebScan/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TaskController struct {
	beego.Controller
}

func (this *TaskController) Prepare() {
	//检测用户是否已经登陆
	isLogin := this.GetSession("IsLogin").(bool)
	if isLogin != true {
		beego.Info("用户未登录,跳转到登陆界面")
		this.Redirect("/login", 302)
	}
}
func (this *TaskController) QueryTask() {

	userId := this.GetSession("user_id").(int)

	type TaskInfo struct {
		TaskId     int    `json:"task_id"`
		TaskName   string `json:"task_name"`
		TaskUrl    string `json:"task_url"`
		TaskNote   string `json:"task_note"`
		TaskPeriod string `json:"task_period"`
		TaskTime   string `json:"task_time"`
		TaskDay    string `json:"task_day"`
		//报文通信获取以下字段
		TaskProgress  int    `json:"task_progress"`
		CreateTime    string `json:"create_time"`
		StartTime     string `json:"start_time"`
		CompleteTime  string `json:"complete_time"`
		LastScanTime  string `json:"last_scan_time"`
		LastCrawlTime string `json:"last_crawl_time"`
	}
	type TaskGroupInfo struct {
		TaskGroupId   int    `json:"task_group_id"`
		TaskGroupName string `json:"task_group_name"`
	}
	type RspInfo struct {
		TaskInfoList  []TaskInfo      `json:"task_info_list"`
		TaskGroupList []TaskGroupInfo `json:"task_group_list"`
	}
	var rspInfo RspInfo
	//查本地数据库，获取任务分组信息
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("select * from table_task_group where user_id=?", userId).QueryRows(&(rspInfo.TaskGroupList))

	//和C++后台通信，获取任务信息
	mgrurl := beego.AppConfig.String("mgrurl")
	wsIntf := models.NewBkWebScanIntf("tcp4", mgrurl, false)
	//假设那些字段会自动映射，例如数据库task_id的值映射到结构体TaskId中
	num2, _ := o.Raw("select * from table_task_info where user_id=?", userId).QueryRows(&(rspInfo.TaskInfoList))
	for i := 0; i < int(num2); i++ {

		//strNum := strA[notEmpty(strA):]
		//a1, err := strconv.Atoi(strNum)
		taskIdNum := rspInfo.TaskInfoList[i].TaskId
		taskIdStr := fmt.Sprintf("%36d", taskIdNum)

		request := models.QueryTaskRequest{
			TaskId:      taskIdStr,
			SyncVersion: 0,
		}
		req := models.QueryTaskReq{
			TimeStamp: int32(time.Now().Unix()),
			Request:   request,
		}
		rsp, queryTaskErr := wsIntf.QueryTask(&req)
		if queryTaskErr != nil {
			beego.Error("没有查询到Task：", taskIdNum)
			continue
		}
		rspInfo.TaskInfoList[i].CreateTime = rsp.Response.TaskInfo.CreateTime
		rspInfo.TaskInfoList[i].StartTime = rsp.Response.TaskInfo.StartTime
		rspInfo.TaskInfoList[i].CompleteTime = rsp.Response.TaskInfo.CompleteTime
		rspInfo.TaskInfoList[i].LastCrawlTime = rsp.Response.TaskInfo.LastCrawlTime
		rspInfo.TaskInfoList[i].LastScanTime = rsp.Response.TaskInfo.LastScanTime
		rspInfo.TaskInfoList[i].TaskProgress = rsp.Response.TaskInfo.TaskState
	}
	this.Data["json"] = &rspInfo
	this.ServeJson()
}
func (this *TaskController) Get() {
	this.TplNames = "task.html"
}
