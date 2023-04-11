package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type Data struct {
	DataSample string `json:"dataSample"`
	TribeOwner string `json:"tribeOwner"`
	IsPII      bool   `json:"isPII"`
	TypePII    string `json:"typePII"`
}

// input column data here
type Column struct {
	ID          Data `json:"id"           dataType:"int"        tribeOwner:"3" isPII:"true"  typePII:"C3"`
	Channel     Data `json:"channel"      dataType:"string"     tribeOwner:"3" isPII:"false" typePII:""`
	Name        Data `json:"name"         dataType:"string"     tribeOwner:"3" isPII:"false" typePII:""`
	CreatedTime Data `json:"created_time" dataType:"timestamp"  tribeOwner:"3" isPII:"false" typePII:""`
}

// input table data here
type Table struct {
	EventLog Column `json:"event_log"`
}

func main() {
	table := Table{}
	column := Column{}

	colType := reflect.TypeOf(column)
	colVal := reflect.ValueOf(column)

	for i := 0; i < colVal.NumField(); i++ {
		field := colType.Field(i)
		isPII := false

		if field.Tag.Get("isPII") == "true" {
			isPII = true
		}

		var dataSample string
		// adjust to cover any data types
		switch field.Tag.Get("dataType") {
		case "int":
			dataSample = "123"
		case "string":
			dataSample = "sample"
		case "timestamp":
			now := time.Now().In(time.FixedZone("Asia/Jakarta", 7*3600))
			dataSample = now.String()
		}

		a := Data{
			DataSample: dataSample,
			TribeOwner: field.Tag.Get("tribeOwner"),
			IsPII:      isPII,
			TypePII:    field.Tag.Get("typePII"),
		}
		SetField(&column, field.Tag.Get("json"), a)
	}

	TableType := reflect.TypeOf(table)
	field := TableType.Field(0)
	SetField(&table, field.Tag.Get("json"), column)

	outputPath, _ := filepath.Abs("files/output.json")

	// Convert to json
	output, err := json.MarshalIndent(table, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write to output
	_, err2 := f.WriteString(string(output))
	if err2 != nil {
		log.Fatal(err2)
	}
}

// ref: https://gist.github.com/lelandbatey/a5c957b537bed39d1d6fb202c3b8de06
// set each field in a struct
func SetField(item interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(item).Elem()
	if !v.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}
	// It's possible we can cache this, which is why precompute all these ahead of time.
	findJsonName := func(t reflect.StructTag) (string, error) {
		if jt, ok := t.Lookup("json"); ok {
			return strings.Split(jt, ",")[0], nil
		}
		return "", fmt.Errorf("tag provided does not define a json tag", fieldName)
	}
	fieldNames := map[string]int{}
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		jname, _ := findJsonName(tag)
		fieldNames[jname] = i
	}

	fieldNum, ok := fieldNames[fieldName]
	if !ok {
		return fmt.Errorf("field %s does not exist within the provided item", fieldName)
	}
	fieldVal := v.Field(fieldNum)
	fieldVal.Set(reflect.ValueOf(value))
	return nil
}
