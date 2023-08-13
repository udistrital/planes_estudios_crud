package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type EliminarProyectoAcademicoPlanEstudioProyectoAcademico_20230813_091005 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarProyectoAcademicoPlanEstudioProyectoAcademico_20230813_091005{}
	m.Created = "20230813_091005"

	migration.Register("EliminarProyectoAcademicoPlanEstudioProyectoAcademico_20230813_091005", m)
}

// Run the migrations
func (m *EliminarProyectoAcademicoPlanEstudioProyectoAcademico_20230813_091005) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230813_091005_eliminar_proyecto_academico_plan_estudio_proyecto_academico_up.sql")
	if err != nil {
		//handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}

// Reverse the migrations
func (m *EliminarProyectoAcademicoPlanEstudioProyectoAcademico_20230813_091005) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230813_091005_eliminar_proyecto_academico_plan_estudio_proyecto_academico_down.sql")
	if err != nil {
		//handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
