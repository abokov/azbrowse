#!/usr/bin/env bash
set -e
cd `dirname $0`

if [ -z "$BUILD_NUMBER" ]; then
    echo "Envvar 'BUILD_NUMBER' must be set for this script to work correctly. When building locally for debugging/testing this script is not needed use 'go build' instead."
    exit 1
fi 

git tag -f v1.0.$BUILD_NUMBER
export GOVERSION=$(go version)

echo "->Installing tools"
bash -f ./install_dev_tools.sh
cd ../


# echo "->Checking swagger-generated code is up-to-date"
# make swagger-codegen
# if [[ $(git diff --stat) != '' ]]; then
#   echo "Ditry GIT: Failing as swagger-generated caused changes, please run `make swagger-generate` and commit changes for build to pass"
#   exit 1
# else
#   echo '`swagger-gen` ran and no changes detected in code: Success'
# fi


echo "->Running dep"
make dep
echo "->Installing tests"
make test

