package models

const (
	NOTIFY_SYNC = 6001 //数据同步请求

	NOTIFY_ADD_TASK = 6003 //任务添加请求

	NOTIFY_DELETE_TASK = 6005 //任务删除请求

	NOTIFY_QUERY_TASK = 6007 //任务查询请求

	NOTIFY_CONTROL_TASK = 6009 //任务控制请求

	NOTIFY_UPDATE_TASK_ATTR = 6011 //任务属性更新请求

	NOTIFY_QUERY_TASK_ATTR = 6013 //任务属性查询请求

	NOTIFY_QUERY_TASK_URL = 6015 //任务URL结果查询请求

	NOTIFY_QUERY_TASK_RESULT = 6017 //任务扫描结果查询请求

	NOTIFY_GET_CRAWL_TASK = 1001 //爬行任务获取请求

	NOTIFY_GET_CRAWL_TASK_ATTR = 1003 //爬行任务属性查询请求

	NOTIFY_UPLOAD_CRAWL_RESULT = 1005 //爬行任务结果上传请求

	NOTIFY_GET_SCAN_TASK = 2001 //扫描任务获取请求

	NOTIFY_GET_SCAN_TASK_ATTR = 2003 //扫描任务属性查询请求

	NOTIFY_UPLOAD_SCAN_RESULT = 2005 //扫描任务结果上报请求

)

var MapErrMsg map[int32]string

const (
	//error message code
	EMC_OK                     int32 = 0
	EMC_MSG_TYPE_INVALD        int32 = 300
	EMC_MSG_FMT_INVALID        int32 = 301
	EMC_SYS_ERROR              int32 = 302
	EMC_CONNECT_ERROR          int32 = 303
	EMC_CONFIG_INVALID         int32 = 304
	EMC_PROCESS_ALREADY_EXISTS int32 = 305
	EMC_PROCESS_KILL_ERROR     int32 = 306
)

const (
	//任务状态  0- 新建 1- 扫描中 2- 停止 3- 暂停 5- 已完成
	TASK_STATUS_INIT    int32 = 0
	TASK_STATUS_RUNNING int32 = 1
	TASK_STATUS_STOP    int32 = 2
	TASK_STATUS_PAUSE   int32 = 3
	TASK_STATUS_DONE    int32 = 4
)

func init() {
	MapErrMsg = make(map[int32]string, 10)
	MapErrMsg[EMC_OK] = "ok"
	MapErrMsg[EMC_MSG_TYPE_INVALD] = "invalid message type!"
	MapErrMsg[EMC_MSG_FMT_INVALID] = "invalid message json format!"
	MapErrMsg[EMC_SYS_ERROR] = "system inner error!"
	MapErrMsg[EMC_CONNECT_ERROR] = "connect refused!"
	MapErrMsg[EMC_CONFIG_INVALID] = "config invalid"
	MapErrMsg[EMC_PROCESS_ALREADY_EXISTS] = "process already exists"
	MapErrMsg[EMC_PROCESS_KILL_ERROR] = "process kill error"
}
