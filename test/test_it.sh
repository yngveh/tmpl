#!/bin/sh

export MYVAR=myvar

echo "Start testing"
PATH=$PATH:./target/

echo "=== Testing template from StdIn ==="
echo '{{ Env "HOME" }}' | tmpl -tmpl -
echo

echo "=== Testing template with data from yaml file ==="
tmpl -tmpl test/template-test.tpl -data test/test-data.yaml
echo

echo "=== Testing template with data from json file ==="
tmpl -tmpl test/template-test.tpl -data test/test-data.json
echo

