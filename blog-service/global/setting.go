package global

import (
	"go-programming-tour/blog-service/pkg/logger"
	"go-programming-tour/blog-service/pkg/setting"
)

var(
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
)
