package captnprotogen

import "fmt"
import "reflect"
import "strings"


func GenerateField(i int, obj_type reflect.StructField ){
	type_map := map[string]string{
		"string": "Text",
		"int64":   "Int64",
	}
	var intypename string = obj_type.Type.Name()
	var typename string = type_map[intypename]
	var name= strings.ToLower(obj_type.Name)
	fmt.Printf("\t%v @%d : %s;\n",name,i,typename)
}

func Generate(obj interface{}, id string){
	fmt.Printf("%s;  # unique file ID, generated by `capnp id`\n",id)
	fmt.Printf("using Go = import \"go.capnp\";\n")
	fmt.Printf("$Go.package(\"captnproto\");")

	obj_reflect := reflect.ValueOf(obj)
	obj_type_reflect := reflect.TypeOf(obj)


	switch obj_reflect.Kind() {
	case reflect.Struct:
		fmt.Printf("struct %v{\n",obj_type_reflect.Name())
		for i := 0; i < obj_reflect.NumField(); i += 1 {
			GenerateField(i,obj_type_reflect.Field(i))
		}
		fmt.Printf("}\n")	
	}
}
