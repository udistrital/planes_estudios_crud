package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type PlanEstudio struct {
	Id                           int               `orm:"column(id);pk"`
	Nombre                       string            `orm:"column(nombre)"`
	CodigoAbreaviacion           string            `orm:"column(codigo_abreaviacion);null"`
	Codigo                       string            `orm:"column(codigo)"`
	ProyectoAcademicoId          int               `orm:"column(proyecto_academico_id)"`
	TotalCreditos                float64           `orm:"column(total_creditos)"`
	NumeroSemestres              float64           `orm:"column(numero_semestres)"`
	NumeroResolucion             float64           `orm:"column(numero_resolucion)"`
	AnoResolucion                float64           `orm:"column(ano_resolucion)"`
	EspaciosSemestreDistribucion string            `orm:"column(espacios_semestre_distribucion);type(json);null"`
	ResumenPlanEstudios          string            `orm:"column(resumen_plan_estudios);type(json);null"`
	SoporteDocumental            string            `orm:"column(soporte_documental);type(json);null"`
	Observacion                  string            `orm:"column(observacion);null"`
	Activo                       bool              `orm:"column(activo)"`
	FechaCreacion                time.Time         `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion            time.Time         `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	PlanEstudioPadreId           *PlanEstudio      `orm:"column(plan_estudio_padre_id);rel(fk)"`
	EstadoAprobacionId           *EstadoAprobacion `orm:"column(estado_aprobacion_id);rel(fk)"`
}

func (t *PlanEstudio) TableName() string {
	return "plan_estudio"
}

func init() {
	orm.RegisterModel(new(PlanEstudio))
}

// AddPlanEstudio insert a new PlanEstudio into database and returns
// last inserted Id on success.
func AddPlanEstudio(m *PlanEstudio) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPlanEstudioById retrieves PlanEstudio by Id. Returns error if
// Id doesn't exist
func GetPlanEstudioById(id int) (v *PlanEstudio, err error) {
	o := orm.NewOrm()
	v = &PlanEstudio{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPlanEstudio retrieves all PlanEstudio matches certain condition. Returns empty list if
// no records exist
func GetAllPlanEstudio(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PlanEstudio))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []PlanEstudio
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePlanEstudio updates PlanEstudio by Id and returns error if
// the record to be updated doesn't exist
func UpdatePlanEstudioById(m *PlanEstudio) (err error) {
	o := orm.NewOrm()
	v := PlanEstudio{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePlanEstudio deletes PlanEstudio by Id and returns error if
// the record to be deleted doesn't exist
func DeletePlanEstudio(id int) (err error) {
	o := orm.NewOrm()
	v := PlanEstudio{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PlanEstudio{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
