package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:EstadoAprobacionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planes_estudios_crud/controllers:PlanEstudioProyectoAcademicoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
