#!/usr/bin/env node

const fs = require('fs');
const yaml = require('js-yaml');

const crds_path = 'api/crds';
const crd_names_list = [
  'modules.kubeform.com_thomasstorageaccounts.yaml'
];

const config_map_file = process.argv[2];
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
    var schema = version_obj['schema']['openAPIV3Schema'];

    crd_schemas[crd_name][version] = schema;
  });
});

cm_obj['data']['schemas'] = JSON.stringify(crd_schemas, null, 2);
fs.writeFileSync(config_map_file, yaml.dump(cm_obj));
