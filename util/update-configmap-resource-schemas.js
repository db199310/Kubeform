#!/usr/bin/env node

const fs = require('fs');
const yaml = require('js-yaml');

const crds_path = process.argv[2];
const crd_names_list = fs.readdirSync(crds_path)
                         .filter(fn => fn.startsWith('modules.kubeform.com_'))
                         .filter(fn => fn.endsWith('.yaml'));

const config_map_file = process.argv[3];
var cm_obj = yaml.load(fs.readFileSync(config_map_file, { encoding: 'UTF-8' }));
var crd_schemas = JSON.parse(cm_obj['data']['schemas']);

crd_names_list.forEach((crd_file) => {
  console.log('Building ' + crd_file + '...');

  const crd_obj = yaml.load(
                  fs.readFileSync(crds_path + '/' + crd_file,
                                  { encoding: 'UTF-8' }));
  const crd_name = crd_obj['spec']['names']['singular'] + '.' +
                   crd_obj['spec']['group'];
  crd_schemas[crd_name] = {};

  // Build version-to-schema map
  crd_obj['spec']['versions'].forEach((version_obj) => {
    var version = version_obj['name'];

    var schema_root;
    if ('schema' in version_obj) {
      schema_root = version_obj['schema'];
    } else {
      schema_root = crd_obj['spec']['validation'];
    }

    var schema = schema_root['openAPIV3Schema'];

    crd_schemas[crd_name][version] = schema;
  });
});

cm_obj['data']['schemas'] = JSON.stringify(crd_schemas, null, 2);
fs.writeFileSync(config_map_file, yaml.dump(cm_obj));
