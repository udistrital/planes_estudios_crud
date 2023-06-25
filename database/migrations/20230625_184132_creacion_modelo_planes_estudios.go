package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionModeloPlanesEstudios_20230625_184132 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionModeloPlanesEstudios_20230625_184132{}
	m.Created = "20230625_184132"

	migration.Register("CreacionModeloPlanesEstudios_20230625_184132", m)
}

// Run the migrations
func (m *CreacionModeloPlanesEstudios_20230625_184132) Up() {
	file, err := ioutil.ReadFile("../scripts/20230625_184132_creacion_modelo_planes_estudios_up.sql")

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
func (m *CreacionModeloPlanesEstudios_20230625_184132) Down() {
	file, err := ioutil.ReadFile("../scripts/20230625_184132_creacion_modelo_planes_estudios_down.sql")

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
