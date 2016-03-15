package models

//数据同步
type FilterStruct struct {
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
}

type SyncRequest struct {
	Table       string `json:"table"`
	SyncVersion int32  `json:"sync_version"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
	//Filter      FilterStruct `json:"filter"`
	Selector []string `json:"selector"`
}
type SyncReq struct {
	TimeStamp int32       `json:"timestamp"`
	Request   SyncRequest `json:"request"`
}

type RecordSetStruct struct {
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
	SyncVersion string `json:"sync_version"`
}
type SyncResponse struct {
	MaxSyncVersion int32             `json:"max_sync_version"`
	MaxCount       int32             `json:"max_count"`
	Table          string            `json:"table"`
	RecordSet      []RecordSetStruct `json:"recordset"`
}
type SyncRsp struct {
	Err       int32        `json:"err"`
	Msg       string       `json:"msg"`
	TimeStamp int32        `json:"timestamp"`
	Response  SyncResponse `json:"response"`
}

//任务添加
/**/
type AddTaskAttrStruct struct {
	Options string `json:"options"`
}
type AddTaskInfoStruct struct {
	TaskName string `json:"task_name"`
	TaskUrl  string `json:"task_url"`
}
type AddTaskRequest struct {
	TaskInfo AddTaskInfoStruct `json:"task_info"`
	TaskAttr AddTaskAttrStruct `json:"task_attr"`
}
type AddTaskReq struct {
	TimeStamp int32          `json:"timestamp"`
	Request   AddTaskRequest `json:"addTaskReq"`
}
type AddTaskResponse struct {
	TaskId string `json:task_id`
}
type AddTaskRsp struct {
	Err       int32           `json:"err"`
	Msg       string          `json:"msg"`
	TimeStamp int32           `json:"timestamp"`
	Response  AddTaskResponse `json:"response"`
}

//任务查询
type QueryTaskRequest struct {
	TaskId      string `json:"task_id"`
	SyncVersion int32  `json:"sync_version"`
}
type QueryTaskReq struct {
	TimeStamp int32            `json:"timestamp"`
	Request   QueryTaskRequest `json:"request"`
}
type QueryTaskInfo struct {
	TaskId        string `json:"task_id"`
	TaskName      string `json:"task_name"`
	TaskUrl       string `json:"task_url"`
	TaskState     int    `json:"task_state"`
	CreateTime    string `json:"create_time"`
	StartTime     string `json:"start_time"`
	CompleteTime  string `json:"complete_time"`
	LastScanTime  string `json:"last_scan_time"`
	LastCrawlTime string `json:"last_crawl_time"`
	SyncVersion   int32  `json:"sync_version"`
}
type QueryTaskResponse struct {
	TaskInfo QueryTaskInfo `json:"task_info"`
}
type QueryTaskRsp struct {
	Err       int32             `json:"err"`
	Msg       string            `json:"msg"`
	TimeStamp int32             `json:"timestamp"`
	Response  QueryTaskResponse `json:"response"`
}

//任务控制
type ControlTaskRequest struct {
	TaskId      string `json:"task_id"`
	ControlCode int    `json:"control_code"`
}
type ControlTaskReq struct {
	TimeStamp int32              `json:"timestamp"`
	Request   ControlTaskRequest `json:"request"`
}
type ControlTaskResponse struct {
	OldState int `json:"old_state"`
	NewState int `json:"new_state"`
}
type ControlTaskRsp struct {
	Err       int32               `json:"err"`
	Msg       string              `json:"msg"`
	TimeStamp int32               `json:"timestamp"`
	Response  ControlTaskResponse `json:"response"`
}

//任务属性查询
type TaskAttrQueryRequest struct {
	TaskId      string `json:"task_id"`
	SyncVersion int    `json:"sync_version"`
}
type TaskAttrQueryReq struct {
	TimeStamp int32                `json:"timestamp"`
	Request   TaskAttrQueryRequest `json:"request"`
}
type TaskAttrStruct struct {
	AttrName    string `json:"attr_name"`
	AttrValue   string `json:"attr_value"`
	SyncVersion int    `json:"sync_version"`
}
type TaskAttrQueryResponse struct {
	TaskAttr []TaskAttrStruct `json:"task_attr"`
	TaskId   string           `json:"task_id"`
}
type TaskAttrQueryRsp struct {
	Err       int                   `json:"err"`
	Response  TaskAttrQueryResponse `json:"response"`
	TimeStamp int32                 `json:"timestamp"`
}

//任务属性设置

type SetTaskAttrStruct struct {
	Options string `json:"options"`
	Test    string `json:"test"`
}
type TaskAttrSetRequest struct {
	TaskId   string            `json:"task_id"`
	TaskAttr SetTaskAttrStruct `json:"task_attr"`
}
type TaskAttrSetReq struct {
	TimeStamp int32              `json:"timestamp"`
	Request   TaskAttrSetRequest `json:"request"`
}

type TaskAttrSetRsp struct {
	Err       int32  `json:"err"`
	Msg       string `json:"msg"`
	TimeStamp int32  `json:timestamp`
	Response  string `json:response`
}

//任务url查询
type TaskUrlQueryRequest struct {
	TaskID      string `json:"task_id"`
	SyncVersion int    `json:"sync_version"`
}

type TaskUrlQueryReq struct {
	TimeStamp int32               `json:"timestamp"`
	Request   TaskUrlQueryRequest `json:"request"`
}

type TaskUrlQueryRecordSet struct {
	UrlId       int    `json:"url_id"`
	UrlValue    string `json:"url_value"`
	TaskId      string `json:"task_id"`
	CreateTime  string `json:"create_time"`
	SyncVersion int    `json:"sync_version"`
}

type TaskUrlQueryResponse struct {
	RecordSet []TaskUrlQueryRecordSet `json:"recordset"`
}

type TaskUrlQueryRsp struct {
	Err       int32                `json:"err"`
	Msg       string               `json:"msg"`
	TimeStamp int32                `json:"timestamp"`
	Response  TaskUrlQueryResponse `json:"response"`
}

//任务结果查询
type TaskResultQueryRequest struct {
	TaskId      string `json:"task_id"`
	SyncVersion int    `json:"sync_version"`
}

type TaskResultQueryReq struct {
	TimeStamp int32                  `json:"timestamp"`
	Request   TaskResultQueryRequest `json:"request"`
}
