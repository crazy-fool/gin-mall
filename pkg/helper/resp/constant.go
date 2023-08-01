package resp

const SuccessCode = 0

const CommonError = 1               //	失败
const CommonErrorMessage = "failed" //	失败
const ParamError = 2                //	请求参数错误
const OpFailed = 3                  //	操作失败
const ResourceNotFound = 4          //	数据不存在

const CustomerNotFound = 101 //	用户不存在
const LoginFailed = 102      //	登录失败
const CustomerNotLogin = 103 //	未登录

var responseMessage = map[int]string{

	SuccessCode:      "success",
	ParamError:       "请求参数有误",
	OpFailed:         "操作失败",
	ResourceNotFound: "数据不存在",
	CustomerNotFound: "用户不存在",
	LoginFailed:      "登录失败",
	CustomerNotLogin: "用户未登录",
}
