package main

import (
	"github.com/spf13/pflag"
	"go_project/api/router"
	"go_project/internal/repository"
	"go_project/internal/service"
	"go_project/pkg/global"
	"go_project/pkg/propertie"
	"go_project/pkg/storage/orm"
	"go_project/pkg/tools"
)

var (
	resourceDir = pflag.StringP("config dir", "c", "resource", "config path.")
	env         = pflag.StringP("env name", "e", "", "env var name.")
)

func init(){
	reloadResource()
	tools.Logger = tools.NewLogger()
}

func main() {
	db := orm.NewMySQL(orm.DbConf)
	//model.InitDB()
	service.Svc = service.New(repository.New(db, nil))
	// 初始化路由
	apiRouter := router.APIRouter()
	// 监听端口，默认在8080
	port := global.Conf.PprofPort
	err := apiRouter.Run(port)
	if err != nil {
		return
	}
}

func reloadResource() {
	pflag.Parse()
	// init properties
	c := propertie.New(*resourceDir, propertie.WithEnv(*env))
	var cfg global.Config
	if err := c.Load("global", &cfg); err != nil {
		panic(err)
	}
	// set global
	global.Conf = &cfg
	var dbCfg orm.Config
	if err := c.Load("database", &dbCfg); err != nil {
		panic(err)
	}
	orm.DbConf = &dbCfg
}
