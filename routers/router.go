// @APIVersion 1.0.0
// @Title Plan de estudios
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/planes_estudios_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_aprobacion",
			beego.NSInclude(
				&controllers.EstadoAprobacionController{},
			),
		),

		beego.NSNamespace("/plan_estudio",
			beego.NSInclude(
				&controllers.PlanEstudioController{},
			),
		),

		beego.NSNamespace("/plan_estudio_proyecto_academico",
			beego.NSInclude(
				&controllers.PlanEstudioProyectoAcademicoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
