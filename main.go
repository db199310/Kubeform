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

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"kubeform.dev/kubeform/util"

	"github.com/gobuffalo/flect"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm"
)

var licenseHeaderFile string

func init() {
	flag.StringVar(&licenseHeaderFile, "license-header-file", "", "Path to the license file.")
}

func GenerateModuleApis() (err error) {
	// Collect versions from the defined modules
	var versions []string
	modulesFileName := "module-sources.json"

	byt, err := ioutil.ReadFile(modulesFileName)
	if err != nil {
		return
	}

	var moduleObjs map[string]interface{}
	err = json.Unmarshal(byt, &moduleObjs)
	if err != nil {
		return
	}

	versionsSet := make(map[string]struct{})
	for _, moduleObj := range moduleObjs {
		for version := range moduleObj.(map[string]interface{}) {
			versionsSet[version] = struct{}{}
		}
	}

	for version := range versionsSet {
		versions = append(versions, version)
	}
	sort.Strings(versions)
	log.Printf("Module Versions: %v", versions)

	// Generate apis for each version
	for _, version := range versions {
		err = util.GenerateProviderAPIS("modules", version, nil, nil)
		if err != nil {
			return
		}
	}

	return
}

func main() {
	flag.Parse()

	if licenseHeaderFile != "" {
		byt, err := ioutil.ReadFile(licenseHeaderFile)
		if err != nil {
			log.Println(err.Error())
		}

		util.License = string(byt)
	}

	version := "v1alpha1"

	providersMap := map[string]terraform.ResourceProvider{
		"azurerm": azurerm.Provider(),
	}

	for key, provider := range providersMap {
		p, ok := provider.(*schema.Provider)
		if ok {
			var structNames []string
			var schemas []map[string]*schema.Schema
			providerPrefix := key
			for key, values := range p.ResourcesMap {
				key = strings.TrimPrefix(key, providerPrefix+"_")
				structNames = append(structNames, flect.Capitalize(flect.Camelize(key)))
				schemas = append(schemas, values.Schema)
			}

			err := util.GenerateProviderAPIS(key, version, schemas, structNames)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}

	err := GenerateModuleApis()
	if err != nil {
		log.Println(err.Error())
	}
}
