package utils

// 接口返回码（2018-10-25 之后）
const (
	//系统类
	SYS_EXCEPTION   = -1  //系统异常
	SYS_MAINTENANCE = 999 //系统维护中

	//请求类
	REQ_SUCCESS            = 0    //请求成功
	REQ_LACK_PARAMS        = 1000 //缺少参数
	REQ_PARAM_VERIF_FAILED = 1001 //参数验证失败

	//响应类
	RES_NOT_FOUND   = 1002 //没有相关数据
	RES_SAVE_FAILED = 1003 //保存失败
	RES_DEL_FAILED  = 1004 //删除失败
	RES_GET_FAILED  = 1005 //获取失败

	//权限类
	AUTH_PERMISSION_DENIED = 1006 //没有权限
)

/*
历史返回码 2018-10-25 之前

-1 系统异常
0 请求成功
999 系统维护中
1001 参数验证失败
1002 没有相关数据
1003 保存失败
1004 删除失败
1005 获取失败
1006 没有权限
1007 重复操作
1008 无效操作
*/
