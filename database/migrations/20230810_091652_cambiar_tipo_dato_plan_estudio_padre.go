package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type CambiarTipoDatoPlanEstudioPadre_20230810_091652 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambiarTipoDatoPlanEstudioPadre_20230810_091652{}
	m.Created = "20230810_091652"

	migration.Register("CambiarTipoDatoPlanEstudioPadre_20230810_091652", m)
}

// Run the migrations
func (m *CambiarTipoDatoPlanEstudioPadre_20230810_091652) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20230810_091652_cambiar_tipo_dato_plan_estudio_padre_up.sql")

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
func (m *CambiarTipoDatoPlanEstudioPadre_20230810_091652) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20230810_091652_cambiar_tipo_dato_plan_estudio_padre_down.sql")

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
