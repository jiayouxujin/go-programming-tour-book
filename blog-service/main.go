package main

import (
	"github.com/gin-gonic/gin"
	"go-programming-tour/blog-service/global"
	"go-programming-tour/blog-service/internal/model"
	"go-programming-tour/blog-service/internal/routers"
	"go-programming-tour/blog-service/pkg/logger"
	setting2 "go-programming-tour/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init()  {
	err:=setupSetting()
	if err!=nil{
		log.Fatalf("init.setupSetting err: %v",err)
	}
	err=setupDBEngine()
	if err!=nil{
		log.Fatalf("init.setupDBEngine err: %v",err)
	}
	err=setupLogger()
	if err!=nil{
		log.Fatalf("init.setupLogger err: %v",err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go编程之旅：一起用Go做项目
func main()  {
	gin.SetMode(global.ServerSetting.RunMode)
	router:=routers.NewRouter()
	s:=&http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting,err:=setting2.NewSetting()
	if err!=nil{
		return err
	}
	err=setting.ReadSection("App",&global.AppSetting)
	if err!=nil{
		return err
	}
	err=setting.ReadSection("Server",&global.ServerSetting)
	if err!=nil{
		return err
	}
	err=setting.ReadSection("Database",&global.DatabaseSetting)
	if err!=nil{
		return err
	}
	global.ServerSetting.ReadTimeout*=time.Second
	global.ServerSetting.WriteTimeout*=time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine,err=model.NewDBEngine(global.DatabaseSetting)
	if err!=nil{
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger=logger.NewLogger(&lumberjack.Logger{
		Filename:global.AppSetting.LogSavePath+"/"+global.AppSetting.LogFileName+global.AppSetting.LogFileExt,
		MaxSize:600,
		MaxAge:10,
		LocalTime:true,
	},"",log.LstdFlags).WithCaller(2)
	return nil
}