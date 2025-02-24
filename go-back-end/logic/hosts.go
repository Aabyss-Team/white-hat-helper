package logic

import (
	"web_app/dao/redis"
	"web_app/models"

	"go.uber.org/zap"
)

//GetHostList 获取主机列表
func GetHostList(params *models.ParamGetHostList) ([]*models.HostListItem, error) {
	//1.查询主机列表
	hostList, err := redis.GetHostsByCompanyID(params)
	if err != nil {
		zap.L().Error("GetHostList redis.GetHostsByCompanyID error", zap.Error(err))
		return nil, err
	}
	//2.查询主机列表的操作系统
	zap.L().Debug("GetHostList redis.GetHostsByCompanyID success", zap.Any("hostList", hostList))
	if err := redis.GetOS(hostList); err != nil {
		zap.L().Error("GetHostList redis.GetOS error", zap.Error(err))
		return nil, err
	}
	//3.查询主机列表的端口信息
	if err := redis.GetPort(hostList); err != nil {
		zap.L().Error("GetHostList redis.GetPort error", zap.Error(err))
		return nil, err
	}
	//3.查询主机列表的web服务
	if err := redis.GetWeb(hostList); err != nil {
		zap.L().Error("GetHostList redis.GetWeb error", zap.Error(err))
		return nil, err
	}
	return hostList, nil
}

func GetHostDetail(ip string) (host models.HostDetail, err error) {
	host.IP = ip
	//1.查询操作系统类型
	os, err := redis.GetOSByIP(ip)
	if err != nil {
		zap.L().Error("GetOSByIP failed ", zap.Error(err))
		return host, nil
	}
	host.OS = os
	//2.查询端口信息
	portDetailList, err := redis.GetPortDetailByIP(ip)
	if err != nil {
		zap.L().Error("GetHostList redis.GetPort error", zap.Error(err))
		return host, err
	}
	host.PortList = portDetailList
	//3.查询web服务信息
	webServiceList, err := redis.GetWebServiceByIP(ip)
	if err != nil {
		zap.L().Error("GetHostList redis.GetWebServiceByIP error", zap.Error(err))
		return host, err
	}
	host.WebList = webServiceList
	return host, nil
}
