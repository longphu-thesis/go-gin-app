package example

import (
	"reflect"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Foo struct {
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f *Foo) reflect() {
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		log.WithFields(log.Fields{
			"Field Name": typeField.Name,
			"Field Value": valueField.Interface(),
			"Tag Value": tag.Get("tag_name"),
		}).Debug(fmt.Sprintf("Reflect"))
	}
}
