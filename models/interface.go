package models

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
)

type BkWebScanIntf struct {
	clt *BkTcpClient
}

func NewBkWebScanIntf(proto string, ipport string, keepAlive bool) *BkWebScanIntf {
	//proto maybe 'tcp4'
	clt := NewBkTcpClient(proto, ipport, keepAlive)
	return &BkWebScanIntf{clt: clt}
}

/*通信接口使用例子
mgrurl := beego.AppConfig.String("mgrurl")
wsIntf := models.NewBkWebScanIntf("tcp4", mgrurl, false)
selector := []string{"option_name", "option_value"}
request := models.SyncRequest{
	Table:       "table_engine_info", //漏洞表
	SyncVersion: 0,
	Limit:       100,
	Offset:      0,
	//Filter:      filter,
	Selector: selector,
}
req := models.SyncReq{
	TimeStamp: int32(time.Second),
	Request:   request,
}
rsp, SyncErr := wsIntf.SyncData(&req)
if SyncErr != nil {
	beego.Error("查询任务失败")
} else {
	fmt.Println("Request of query:", req)
	fmt.Println("Response of query:", rsp)
}
*/
func notSpaceIndex(str string) int {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			return i + 1
		}
	}
	return -1
}

/*模板，照着抄*/

//数据同步
func (this *BkWebScanIntf) SyncData(syncReq *SyncReq) (*SyncRsp, error) { //++
	data, err := json.Marshal(&syncReq) //++
	if err != nil {
		return nil, err
	}
	size := uint32(len(data))
	req := NewBkNotifyPacket(BKWS_MAGIC, NOTIFY_SYNC, size, data)
	rsp, err := this.clt.Send(req)
	if err != nil {
		beego.Error("send packet fail!")
		return nil, errors.New("send packet fail!")
	} else {
		rspData := rsp.GetData()
		rst := SyncRsp{} //++
		json.Unmarshal(rspData, &rst)
		return &rst, nil
	}
}

//任务添加
func (this *BkWebScanIntf) AddTask(addtaskReq *AddTaskReq) (*AddTaskRsp, error) {
	data, err := json.Marshal(&addtaskReq)
	if err != nil {
		return nil, err
	}

	size := uint32(len(data))
	req := NewBkNotifyPacket(BKWS_MAGIC, NOTIFY_ADD_TASK, size, data)
	rsp, err := this.clt.Send(req)

	if err != nil {
		beego.Error("send packet fail!")
		return nil, errors.New("send packet fail!")
	} else {
		rspData := rsp.GetData()
		rst := AddTaskRsp{}
		json.Unmarshal(rspData, &rst)
		return &rst, nil
	}
}

//任务查询
func (this *BkWebScanIntf) QueryTask(querytaskReq *QueryTaskReq) (*QueryTaskRsp, error) {
	data, err := json.Marshal(&querytaskReq)
	if err != nil {
		return nil, err
	}

	size := uint32(len(data))
	req := NewBkNotifyPacket(BKWS_MAGIC, NOTIFY_QUERY_TASK, size, data)
	rsp, err := this.clt.Send(req)

	if err != nil {
		beego.Error("send packet fail!")
		return nil, errors.New("send packet fail!")
	} else {
		rspData := rsp.GetData()
		rst := QueryTaskRsp{}
		json.Unmarshal(rspData, &rst)
		return &rst, nil
	}
}

//任务控制
func (this *BkWebScanIntf) ControlTask(controltaskReq *ControlTaskReq) (*ControlTaskRsp, error) {
	data, err := json.Marshal(&controltaskReq)
	if err != nil {
		return nil, err
	}

	size := uint32(len(data))
	req := NewBkNotifyPacket(BKWS_MAGIC, NOTIFY_CONTROL_TASK, size, data)
	rsp, err := this.clt.Send(req)

	if err != nil {
		beego.Error("send packet fail!")
		return nil, errors.New("send packet fail!")
	} else {
		rspData := rsp.GetData()
		rst := ControlTaskRsp{}
		json.Unmarshal(rspData, &rst)
		return &rst, nil
	}
}

//任务属性查询
func (this *BkWebScanIntf) QueryTaskAttr(taskattrReq *TaskAttrQueryReq) (*TaskAttrQueryRsp, error) {
	data, err := json.Marshal(&taskattrReq)
	if err != nil {
		return nil, err
	}

	size := uint32(len(data))
	req := NewBkNotifyPacket(BKWS_MAGIC, NOTIFY_QUERY_TASK_ATTR, size, data)
	rsp, err := this.clt.Send(req)

	if err != nil {
		beego.Error("send packet fail!")
		return nil, errors.New("send packet fail!")
	} else {
		rspData := rsp.GetData()
		rst := TaskAttrQueryRsp{}
		json.Unmarshal(rspData, &rst)
		return &rst, nil
	}
}

//任务属性设置
func (this *BkWebScanIntf) TaskAttrSet(taskattrsetReq *TaskAttrQueryReq) (*TaskAttrSetRsp, error) {
	data, err := json.Marshal(&taskattrsetReq)
	if err != nil {
		return nil, err
	}

	size := uint32(len(data))
	req := NewBkNotifyPacket(BKWS_MAGIC, NOTIFY_UPDATE_TASK_ATTR, size, data)
	rsp, err := this.clt.Send(req)

	if err != nil {
		beego.Error("send packet fail!")
		return nil, errors.New("send packet fail!")
	} else {
		rspData := rsp.GetData()
		rst := TaskAttrSetRsp{}
		json.Unmarshal(rspData, &rst)
		return &rst, nil
	}
}

//任务URL查询
func (this *BkWebScanIntf) TaskUrlQuery(taskurlqueryReq *TaskUrlQueryReq) (*TaskUrlQueryRsp, error) {
	data, err := json.Marshal(&taskurlqueryReq)
	if err != nil {
		return nil, err
	}

	size := uint32(len(data))
	req := NewBkNotifyPacket(BKWS_MAGIC, NOTIFY_QUERY_TASK_URL, size, data)
	rsp, err := this.clt.Send(req)

	if err != nil {
		beego.Error("send packet fail!")
		return nil, errors.New("send packet fail!")
	} else {
		rspData := rsp.GetData()
		rst := TaskUrlQueryRsp{}
		json.Unmarshal(rspData, &rst)
		return &rst, nil
	}
}
