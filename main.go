package main

import (
	"net/url"

	_ "github.com/udistrital/planes_estudios_crud/routers"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/auditoria"
	"github.com/udistrital/utils_oas/security"
	"github.com/udistrital/utils_oas/xray"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	"github.com/udistrital/utils_oas/customerrorv2"
	"github.com/udistrital/utils_oas/database"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		url.QueryEscape(beego.AppConfig.String("PGpass"))+"@"+
		beego.AppConfig.String("PGhost")+":"+
		beego.AppConfig.String("PGport")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschema")+"")
}

func main() {
	// orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	conn, err := database.BuildPostgresConnectionString()
	if err != nil {
		logs.Error("error consultando la cadena de conexión: %v", err)
		return
	}

	err = orm.RegisterDataBase("default", "postgres", conn)
	if err != nil {
		logs.Error("error al conectarse a la base de datos: %v", err)
		return
	}
	////////////////////

	allowedOrigins := []string{"*.udistrital.edu.co"}
	if beego.BConfig.RunMode == beego.DEV {
		allowedOrigins = []string{"*"}
		orm.Debug = true // Solo para APIs CRUD
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"DELETE", "GET", "OPTIONS", "PATCH", "POST", "PUT"}, // ajustar según los métodos usados en el api
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"User-Agent",
			"X-Amzn-Trace-Id"},
		ExposeHeaders:    []string{"Content-Length"}, // agregar otros headers según sea el caso
		AllowCredentials: true,
	}))

	err = xray.InitXRay()
	if err != nil {
		logs.Error("error configurando AWS XRay: %v", err)
	}
	apistatus.Init()
	auditoria.InitMiddleware()
	beego.ErrorController(&customerrorv2.CustomErrorController{})
	security.SetSecurityHeaders()
	beego.Run()
}
