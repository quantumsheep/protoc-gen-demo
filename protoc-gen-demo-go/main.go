package main

import (
	"net/http"
	"path"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func main() {
	options := protogen.Options{}
	options.Run(run)
}

func run(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		generateFile(plugin, file)
	}

	return nil
}

func generateFile(plugin *protogen.Plugin, file *protogen.File) {
	directory, name := path.Split(file.Desc.Path())
	name = strings.TrimSuffix(name, ".proto") + ".go"

	g := plugin.NewGeneratedFile(directory+name, "")
	packagePath := file.Desc.Package()
	packageName := packagePath.Parent().Name() + packagePath.Name()

	g.P("package ", packageName)
	g.P()
	g.P("import (")
	g.P(`  "bytes"`)
	g.P(`  "encoding/json"`)
	g.P(`  "io"`)
	g.P(`  "net/http"`)
	g.P(")")
	g.P()

	for _, message := range file.Messages {
		name := message.Desc.Name()
		g.P("type ", name, " struct {")
		for _, field := range message.Fields {
			g.P(field.Desc.Name(), " ", field.Desc.Kind())
		}
		g.P("}")
		g.P()
	}

	for _, service := range file.Services {
		for _, method := range service.Methods {
			name := method.Desc.Name()
			requestType := method.Input.Desc.Name()
			responseType := method.Output.Desc.Name()

			httpRule := proto.GetExtension(method.Desc.Options(), annotations.E_Http).(*annotations.HttpRule)

			httpMethod := ""
			httpPath := ""
			switch httpRule.GetPattern().(type) {
			case *annotations.HttpRule_Get:
				httpMethod, httpPath = http.MethodGet, httpRule.GetGet()
			case *annotations.HttpRule_Put:
				httpMethod, httpPath = http.MethodPut, httpRule.GetPut()
			case *annotations.HttpRule_Post:
				httpMethod, httpPath = http.MethodPost, httpRule.GetPost()
			case *annotations.HttpRule_Delete:
				httpMethod, httpPath = http.MethodDelete, httpRule.GetDelete()
			case *annotations.HttpRule_Patch:
				httpMethod, httpPath = http.MethodPatch, httpRule.GetPatch()
			case *annotations.HttpRule_Custom:
				httpMethod, httpPath = httpRule.GetCustom().GetKind(), httpRule.GetCustom().GetPath()
			}

			g.P("func ", name, "(request *", requestType, ") (*", responseType, ", error) {")
			g.P("  var buf bytes.Buffer")
			g.P("  err := json.NewEncoder(&buf).Encode(request)")
			g.P("  if err != nil {")
			g.P("    return nil, err")
			g.P("  }")
			g.P()
			if httpMethod == http.MethodGet {
				g.P("  res, err := http.", httpMethod[:1]+strings.ToLower(httpMethod[1:]), `("`, httpPath, `")`)
			} else {
				g.P("  res, err := http.", httpMethod[:1]+strings.ToLower(httpMethod[1:]), `("`, httpPath, `", "application/json", &buf)`)
			}
			g.P("  if err != nil {")
			g.P("    return nil, err")
			g.P("  }")
			g.P("  defer res.Body.Close()")
			g.P()
			g.P("  body, err := io.ReadAll(res.Body)")
			g.P("  if err != nil {")
			g.P("    return nil, err")
			g.P("  }")
			g.P()
			g.P("  value := &", responseType, "{}")
			g.P("  err = json.Unmarshal(body, value)")
			g.P("  if err != nil {")
			g.P("    return nil, err")
			g.P("  }")
			g.P()
			g.P("  return value, nil")
			g.P("}")
			g.P()
		}
	}
}
