package controllers

import (
	"web_app/logic"
	"web_app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//AddCompanyHandler 添加公司
func AddCompanyHandler(c *gin.Context) {
	//1.接收传参
	var param models.ParamAddCompany
	if msg, err := ValidateJSONParam(c, &param); err != nil {
		zap.L().Error("add company handler param error", zap.Error(err))
		RespErrMsg(c, CodeInvalidParam, msg)
		return
	}
	//2.查重
	exist, err := logic.IsCompanyExist(param.Name)
	if err != nil {
		zap.L().Error("add company handler check company exist error", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	if exist {
		RespErrMsg(c, CodeAlreadyExist, "公司已存在")
		return
	}
	//2.业务逻辑层
	company, err := logic.AddCompany(&param)
	if err != nil {
		zap.L().Error("add company handler error", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, company)
}

//GetCompanyHandler 获取公司列表
func GetCompanyListHandler(c *gin.Context) {
	//1.接收传参
	var param models.ParamGetCompanyList
	if msg, err := ValidateQueryParam(c, &param); err != nil {
		zap.L().Error("get company list handler param error", zap.Error(err))
		RespErrMsg(c, CodeInvalidParam, msg)
		return
	}
	//2.业务逻辑层
	companies, err := logic.GetCompanyList(&param)
	if err != nil {
		zap.L().Error("get company list handler error", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, companies)
}
