package models

//ParamsRegister 注册需要参数
type ParamsRegister struct {
	Password  string `json:"password" binding:"required"`
	ReCAPTCHA string `form:"g_recaptcha_response" json:"g_recaptcha_response" binding:"required"`
}

//ParamLogout 退出登录
type ParamLogout struct {
	Token string `json:"token" binding:"required"`
}

//ParamLogin 用户名，密码登录需要的参数
type ParamLogin struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	ReCAPTCHA string `form:"g_recaptcha_response" json:"g_recaptcha_response" binding:"required"`
}

//ParamAddCompany 添加公司需要的参数
type ParamAddCompany struct {
	Name string `json:"name" binding:"required"`
}

//Page 分页参数
type Page struct {
	Offset int `json:"offset" form:"offset"`
	Count  int `json:"count" form:"count"`
}

//ParamGetCompanyList 获取公司需要的参数
type ParamGetCompanyList struct {
	Page
}
