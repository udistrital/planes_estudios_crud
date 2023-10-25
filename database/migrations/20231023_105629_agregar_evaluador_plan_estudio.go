package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type AgregarEvaluadorPlanEstudio_20231023_105629 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarEvaluadorPlanEstudio_20231023_105629{}
	m.Created = "20231023_105629"

	migration.Register("AgregarEvaluadorPlanEstudio_20231023_105629", m)
}

// Run the migrations
func (m *AgregarEvaluadorPlanEstudio_20231023_105629) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20231023_105629_agregar_evaluador_plan_estudio_up.sql")

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
func (m *AgregarEvaluadorPlanEstudio_20231023_105629) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20231023_105629_agregar_evaluador_plan_estudio_down.sql")

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
