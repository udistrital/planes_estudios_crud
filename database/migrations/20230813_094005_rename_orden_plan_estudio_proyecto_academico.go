package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type RenameOrdenPlanEstudioProyectoAcademico_20230813_094005 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RenameOrdenPlanEstudioProyectoAcademico_20230813_094005{}
	m.Created = "20230813_094005"

	migration.Register("RenameOrdenPlanEstudioProyectoAcademico_20230813_094005", m)
}

// Run the migrations
func (m *RenameOrdenPlanEstudioProyectoAcademico_20230813_094005) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230813_094005_rename_orden_plan_estudio_proyecto_academico_up.sql")

	if err != nil {
		// handle error
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
func (m *RenameOrdenPlanEstudioProyectoAcademico_20230813_094005) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230813_094005_rename_orden_plan_estudio_proyecto_academico_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
