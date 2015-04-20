package capnprotogen_transfer
/*
generate a transfer function to get and set the object
*/
import "fmt"
import "reflect"

func Generate(obj interface{}){
	fmt.Printf("package transfer\n")
	obj_type_reflect := reflect.TypeOf(obj)

	switch obj_type_reflect.Kind() {
	case reflect.Struct:
		fmt.Printf("func Set(src %v, dest cap.%v){\n",obj_type_reflect.Name(),obj_type_reflect.Name())
		for i := 0; i < obj_type_reflect.NumField(); i += 1 {
			f := obj_type_reflect.Field(i).Name
			fmt.Printf("\tdest.Set%v(src.%v)\n",f,f)
		}
		fmt.Printf("}\n")
		
		fmt.Printf("func Get(src cap.%v, dest %v){\n",obj_type_reflect.Name(),obj_type_reflect.Name())
		for i := 0; i < obj_type_reflect.NumField(); i += 1 {
			f := obj_type_reflect.Field(i).Name
			fmt.Printf("\tdest.%v=src.%v()\n",f,f)

		}
		fmt.Printf("}\n")
	}
}
