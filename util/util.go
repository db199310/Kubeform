/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
//nolint:goconst
package util

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"kubeform.dev/kubeform/data"

	"github.com/dave/jennifer/jen"
	"github.com/gobuffalo/flect"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type TypeData struct {
	Name    string
	Spec    string
	License string
}

type ApisData struct {
	ProviderName string
	Version      string
	StructName   []string
	License      string
}

var License string

var execeptionList = map[string]string{
	"ConfigConfigurationRecorderStatus": "ConfigConfigurationRecorderStatus_",
	"EventhubNamespace":                 "EventhubNamespace_",
	"NotificationHubNamespace":          "NotificationHubNamespace_",
	"ProjectServices":                   "ProjectServiceBatch",
}

func GenerateProviderAPIS(providerName, version string, schmeas []map[string]*schema.Schema, structNames []string) error {
	templatePath := "templates"
	goPath := os.Getenv("GOPATH")
	basePath := filepath.Join(goPath,
		filepath.Join("src",
			filepath.Join("kubeform.dev",
				filepath.Join("kubeform"))))
	providerPath := filepath.Join(basePath,
		filepath.Join("apis", providerName))
	versionPath := filepath.Join(providerPath, version)
	dataPath := filepath.Join(basePath, "data")

	err := os.MkdirAll(versionPath, 0777)
	if err != nil {
		return err
	}
	if providerName == "modules" {
		var modulePaths []string
		modulePaths, structNames, err = GenerateModuleData(filepath.Join(basePath, "module-sources.json"), filepath.Join(templatePath, "cfg_data.tmpl"), dataPath)
		if err != nil {
			return err
		}
		for i, modulePath := range modulePaths {
			name := structNames[i]
			out := GenerateModuleCRD(modulePath, name)
			typeData := TypeData{
				Name:    name,
				Spec:    out,
				License: License,
			}
			modulePaths[i] = name
			templateToGoFile(filepath.Join(templatePath, "module_types.tmpl"), filepath.Join(versionPath, flect.Underscore(name)+"_types.go"), typeData)
		}
	} else {
		for i, structName := range structNames {
			var out string
			if val, ok := execeptionList[structName]; ok {
				structName = val
				structNames[i] = val
			}
			genSecret := false
			TerraformSchemaToStruct(schmeas[i], structName+"Spec", providerName, true, &genSecret, &out)
			typeData := TypeData{
				Name:    structName,
				Spec:    out,
				License: License,
			}

			templateToGoFile(filepath.Join(templatePath, "types.tmpl"), filepath.Join(versionPath, flect.Underscore(structName)+"_types.go"), typeData)
		}
	}

	sort.Strings(structNames)

	apiData := ApisData{
		ProviderName: providerName,
		Version:      version,
		StructName:   structNames,
		License:      License,
	}

	templateToGoFile(filepath.Join(templatePath, "doc.tmpl"), filepath.Join(versionPath, "doc.go"), apiData)
	templateToGoFile(filepath.Join(templatePath, "register.tmpl"), filepath.Join(versionPath, "register.go"), apiData)
	templateToGoFile(filepath.Join(templatePath, "provider_register.tmpl"), filepath.Join(providerPath, "register.go"), apiData)

	return nil
}

func TerraformSchemaToStruct(s map[string]*schema.Schema, structName, providerName string, genProviderRef bool, genSecret *bool, out *string) {
	var statements jen.Statement
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := s[key]
		id := flect.Capitalize(flect.Camelize(key))
		if value.Description != "" {
			statements = append(statements, jen.Comment("// "+value.Description))
		}
		if value.Removed != "" {
			continue
		}

		jk := flect.Camelize(key) // json key
		tk := key                 // terraform key
		if value.Optional || value.Computed {
			statements = append(statements, jen.Comment("// +optional"))
			jk = jk + ",omitempty"
			tk = tk + ",omitempty"
		}

		if value.MaxItems != 0 {
			statements = append(statements, jen.Comment("// +kubebuilder:validation:MaxItems="+strconv.Itoa(value.MaxItems)))
		}

		if value.MinItems != 0 {
			statements = append(statements, jen.Comment("// +kubebuilder:validation:MinItems="+strconv.Itoa(value.MinItems)))
		}

		if value.Sensitive {
			if value.Type == schema.TypeString {
				statements = append(statements, jen.Id(id).String().Tag(map[string]string{"json": "-", "tf": tk, "sensitive": "true"}))
			} else if value.Type == schema.TypeMap {
				statements = append(statements, jen.Id(id).Map(jen.String()).String().Tag(map[string]string{"json": "-", "tf": tk, "sensitive": "true"}))
			}
			*genSecret = true
			continue
		}

		if value.Deprecated != "" {
			statements = append(statements, jen.Comment("// Deprecated"))
		}

		switch value.Type {
		case schema.TypeString:
			statements = append(statements, jen.Id(id).String().Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeInt:
			statements = append(statements, jen.Id(id).Int64().Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeBool:
			statements = append(statements, jen.Id(id).Bool().Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeFloat:
			statements = append(statements, jen.Id(id).Float64().Tag(map[string]string{"json": jk, "tf": tk}))
		case schema.TypeMap:
			switch value.Elem.(type) {
			case *schema.Schema:
				switch value.Elem.(*schema.Schema).Type {
				case schema.TypeInt:
					statements = append(statements, jen.Id(id).Map(jen.String()).Int64().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeFloat:
					statements = append(statements, jen.Id(id).Map(jen.String()).Float64().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeBool:
					statements = append(statements, jen.Id(id).Map(jen.String()).Bool().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeString:
					statements = append(statements, jen.Id(id).Map(jen.String()).String().Tag(map[string]string{"json": jk, "tf": tk}))
				}
			case *schema.Resource:
				statements = append(statements, jen.Id(id).Map(jen.String()).Id(structName+id).Tag(map[string]string{"json": jk, "tf": tk}))
				TerraformSchemaToStruct(value.Elem.(*schema.Resource).Schema, structName+id, providerName, false, genSecret, out)
			default:
				statements = append(statements, jen.Id(id).Map(jen.String()).String().Tag(map[string]string{"json": jk, "tf": tk}))
			}
		default:
			switch value.Elem.(type) {
			case *schema.Schema:
				switch value.Elem.(*schema.Schema).Type {
				case schema.TypeInt:
					statements = append(statements, jen.Id(id).Index().Int64().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeFloat:
					statements = append(statements, jen.Id(id).Index().Float64().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeBool:
					statements = append(statements, jen.Id(id).Index().Bool().Tag(map[string]string{"json": jk, "tf": tk}))
				case schema.TypeString:
					statements = append(statements, jen.Id(id).Index().String().Tag(map[string]string{"json": jk, "tf": tk}))
				}
			case *schema.Resource:
				statements = append(statements, jen.Id(id).Index().Id(structName+id).Tag(map[string]string{"json": jk, "tf": tk}))
				TerraformSchemaToStruct(value.Elem.(*schema.Resource).Schema, structName+id, providerName, false, genSecret, out)
			default:
				log.Fatalf("Provider %s has resource %s type %s.%s with unknown schema type %s", providerName, structName, structName, id, value.Elem)
			}

		}
	}

	if genProviderRef {
		if *genSecret {
			statements = append(jen.Statement{jen.Id("SecretRef").Id("*core.LocalObjectReference").Tag(map[string]string{"json": "secretRef,omitempty", "tf": "-"}).Line()}, statements...)
		}
		statements = append(jen.Statement{jen.Id("RemoteBackend").Id("*base.Backend").Tag(map[string]string{"json": "remoteBackend,omitempty", "tf": "-"}).Line()}, statements...)
		statements = append(jen.Statement{jen.Id("ID").Id("string").Tag(map[string]string{"json": "id,omitempty", "tf": "id,omitempty"}).Line()}, statements...)
		statements = append(jen.Statement{jen.Id("ProviderRef").Id("core.LocalObjectReference").Tag(map[string]string{"json": "providerRef", "tf": "-"}).Line()}, statements...)
	}

	if val, ok := execeptionList[structName]; ok {
		structName = val
	}

	c := jen.Type().Id(structName).Struct(statements...)
	*out = *out + fmt.Sprintf("%#v\n\n", c)
}

func templateToGoFile(templateFile, generatedFilePath string, templateData interface{}) {
	tmpl := template.Must(template.ParseFiles(templateFile))
	var buffer bytes.Buffer
	err := tmpl.Execute(&buffer, templateData)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ioutil.WriteFile(generatedFilePath, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func parseObjectVariables(objectVar *tfconfig.Variable) map[string]*tfconfig.Variable {
	object := strings.TrimFunc(objectVar.Type, func(r rune) bool {
		return r != '{' && r != '}'
	})
	object = strings.Trim(object, "{}")
	object = strings.ReplaceAll(object, " ", "")
	object = strings.TrimSpace(object)
	varsArray := strings.Split(object, "\n")
	objectVars := make(map[string]*tfconfig.Variable)
	for _, e := range varsArray {
		parts := strings.Split(e, "=")
		objectVars[parts[0]] = &tfconfig.Variable{
			Name:     parts[0],
			Type:     parts[1],
			Required: objectVar.Required,
		}
	}
	return objectVars
}

func terraformVariablesToStatement(name string, variables map[string]*tfconfig.Variable) (jen.Statement, map[string]map[string]*tfconfig.Variable) {
	var specStatements jen.Statement

	var keys []string
	for k := range variables {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	objects := make(map[string]map[string]*tfconfig.Variable)

	for _, key := range keys {
		variable := variables[key]
		if variable.Type == "" {
			variable.Type = "any"
		}
		tk := key
		jk := flect.Camelize(tk)
		id := flect.Capitalize(jk)
		if !variable.Required {
			specStatements = append(specStatements, jen.Comment("// +optional"))
			tk = tk + ",omitempty"
			jk = jk + ",omitempty"
		}
		if variable.Description != "" {
			specStatements = append(specStatements, jen.Comment(variable.Description))
		}

		switch {
		case variable.Type == "number":
			specStatements = append(specStatements, jen.Id(id).Qual("encoding/json", "Number").Tag(map[string]string{"json": jk, "tf": tk}))
		case variable.Type == "string":
			specStatements = append(specStatements, jen.Id(id).String().Tag(map[string]string{"json": jk, "tf": tk}))
		case variable.Type == "bool":
			//Making bool a pointer to preserve optional false values during json marshalling
			specStatements = append(specStatements, jen.Id(id).Op("*").Bool().Tag(map[string]string{"json": jk, "tf": tk}))
		case variable.Type == "any":
			specStatements = append(specStatements, jen.Id(id).Qual("encoding/json", "RawMessage").Tag(map[string]string{"json": jk, "tf": tk}))
		case strings.HasPrefix(variable.Type, "object"):
			specStatements = append(specStatements, jen.Id(id).Id(name+id).Tag(map[string]string{"json": jk, "tf": tk}))
			objectVars := parseObjectVariables(variable)
			objects[name+id] = objectVars
		case strings.HasPrefix(variable.Type, "list") || strings.HasPrefix(variable.Type, "set"):
			//try to extract the list type, fallback to default switch case if type's not set
			listType := strings.TrimFunc(variable.Type, func(r rune) bool {
				return r != '(' && r != ')'
			})
			listType = strings.Trim(listType, "()")

			switch {
			case listType == "bool":
				specStatements = append(specStatements, jen.Id(id).Index().Op("*").Bool().Tag(map[string]string{"json": jk, "tf": tk}))
			case listType == "number":
				specStatements = append(specStatements, jen.Id(id).Index().Qual("encoding/json", "Number").Tag(map[string]string{"json": jk, "tf": tk}))
			case listType == "string":
				specStatements = append(specStatements, jen.Id(id).Index().String().Tag(map[string]string{"json": jk, "tf": tk}))
			case strings.HasPrefix(listType, "object"):
				specStatements = append(specStatements, jen.Id(id).Index().Id(name+id).Tag(map[string]string{"json": jk, "tf": tk}))
				objectVars := parseObjectVariables(variable)
				objects[name+id] = objectVars
			//type not set or "any"
			default:
				specStatements = append(specStatements, jen.Id(id).Index().Qual("encoding/json", "RawMessage").Tag(map[string]string{"json": jk, "tf": tk}))
			}
		case strings.HasPrefix(variable.Type, "map"):
			mapType := strings.TrimFunc(variable.Type, func(r rune) bool {
				return r != '(' && r != ')'
			})
			mapType = strings.Trim(mapType, "()")

			switch {
			case mapType == "bool":
				specStatements = append(specStatements, jen.Id(id).Map(jen.String()).Op("*").Bool().Tag(map[string]string{"json": jk, "tf": tk}))
			case mapType == "number":
				specStatements = append(specStatements, jen.Id(id).Map(jen.String()).Qual("encoding/json", "Number").Tag(map[string]string{"json": jk, "tf": tk}))
			case mapType == "string":
				specStatements = append(specStatements, jen.Id(id).Map(jen.String()).String().Tag(map[string]string{"json": jk, "tf": tk}))
			case strings.HasPrefix(mapType, "object"):
				specStatements = append(specStatements, jen.Id(id).Map(jen.String()).Id(name+id).Tag(map[string]string{"json": jk, "tf": tk}))
				objectVars := parseObjectVariables(variable)
				objects[name+id] = objectVars
			default:
				specStatements = append(specStatements, jen.Id(id).Map(jen.String()).Qual("encoding/json", "RawMessage").Tag(map[string]string{"json": jk, "tf": tk}))
			}
		}
	}

	return specStatements, objects
}

func GenerateModuleCRD(path, name string) string {
	if tfconfig.IsModuleDir(path) {
		m, _ := tfconfig.LoadModule(path)
		variables := m.Variables
		outputs := m.Outputs

		specStatements, objects := terraformVariablesToStatement(name, variables)

		specStatements = append(jen.Statement{jen.Id("Source").String().Tag(map[string]string{"json": "source", "tf": "source"}).Line()}, specStatements...)
		specStatements = append(jen.Statement{jen.Comment("// +optional")}, specStatements...)
		specStatements = append(jen.Statement{jen.Id("RemoteBackend").Id("*base.Backend").Tag(map[string]string{"json": "remoteBackend,omitempty", "tf": "-"}).Line()}, statements...)
		specStatements = append(jen.Statement{jen.Id("ProviderRef").Id("core.LocalObjectReference").Tag(map[string]string{"json": "providerRef", "tf": "-"})}, specStatements...)
		specStatements = append(jen.Statement{jen.Id("SecretRef").Id("*core.LocalObjectReference").Tag(map[string]string{"json": "secretRef,omitempty", "tf": "-"})}, specStatements...)
		specStatements = append(jen.Statement{jen.Comment("// +optional")}, specStatements...)

		out := fmt.Sprintf("%#v\n\n", jen.Type().Id(name+"Spec").Struct(specStatements...))

		var objectNames []string
		for k := range objects {
			objectNames = append(objectNames, k)
		}

		sort.Strings(objectNames)

		for _, name := range objectNames {
			objectStatements, _ := terraformVariablesToStatement(name, objects[name])
			out = out + fmt.Sprintf("%#v\n\n", jen.Type().Id(name).Struct(objectStatements...))
		}

		var outputStatements jen.Statement
		keys := []string{}
		for k := range outputs {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {
			varName := key
			out := outputs[key]
			tk := varName
			jk := flect.Camelize(varName)
			id := flect.Capitalize(jk)
			outputStatements = append(outputStatements, jen.Comment(out.Description))
			outputStatements = append(outputStatements, jen.Comment("// +optional"))
			outputStatements = append(outputStatements, jen.Id(id).String().Tag(map[string]string{"json": jk, "tf": tk}))
		}
		out = out + fmt.Sprintf("%#v\n", jen.Type().Id(name+"Output").Struct(outputStatements...))

		return out
	}

	return ""
}

func DownloadRepository(name, link, outputPath string) (string, error) {
	tempDirPath, err := ioutil.TempDir(os.TempDir(), "")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempDirPath)

	zipPath := filepath.Join(tempDirPath, name+".zip")

	client := &http.Client{}

	req, err := http.NewRequest("GET", link, nil)
	if strings.Contains(link, "sede-ds-adp") {
		req.SetBasicAuth("test", os.Getenv("GIT_READ_TOKEN"))
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s is not a valid github url", link)
	}

	out, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	extractedRepoPath, err := Unzip(zipPath, filepath.Join(outputPath, name))
	if err != nil {
		return "", err
	}
	if strings.Contains(link, "sede-ds-adp") {
		return filepath.Join(outputPath, name), nil
	}
	return extractedRepoPath[0], nil
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		fpath := filepath.Join(dest, f.Name)

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, os.ModePerm)
			if err != nil {
				return filenames, err
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func GenerateModuleData(fileName, templateFile, generatedFilePath string) (modulePaths []string, structNames []string, err error) {
	byt, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	var config []data.Config

	err = json.Unmarshal(byt, &config)
	if err != nil {
		return
	}

	var cfgStr string
	for _, cfg := range config {
		cfgStr = cfgStr + `{
"` + cfg.Name + `","` + cfg.Source + `","` + cfg.Provider + `","","",
},
`
		path, err := DownloadRepository(cfg.Name, cfg.URL, "/tmp")
		if err != nil {
			fmt.Println(err)
		}

		modulePaths = append(modulePaths, filepath.Join(path, cfg.SubDirectory))
		structNames = append(structNames, cfg.Name)
	}

	tmpl := template.Must(template.ParseFiles(templateFile))
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, struct {
		Data    string
		License string
	}{
		Data:    cfgStr,
		License: License,
	})
	if err != nil {
		return
	}
	configFile := filepath.Join(generatedFilePath, "config_data.go")
	_, err = os.Create(configFile)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(configFile, buffer.Bytes(), 0644)
	return
}
