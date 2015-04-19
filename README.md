# go-capnproto-gen
generator for Cap'n Proto schema files from go structure reflection

```
import "github.com/h4ck3rm1k3/captnproto-gen"

func main() {
	var prototype taginfo_model.Keys
	captnprotogen.Generate(prototype,"@0xbbfb7b1b11adeed6")
}

```
