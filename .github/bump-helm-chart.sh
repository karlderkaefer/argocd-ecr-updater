#!/bin/bash

CHART_DIR=${CHART_DIR:-"charts/argocd-ecr-updater"}

while getopts v:b: flag
do
    case "${flag}" in
        v) version=${OPTARG};;
        b) bump=${OPTARG};;
    esac
done

# create version file
echo "creating version file with $version"
echo "$version" > .version

if [ "$bump" = "true" ]; then
  echo "bumping chart to $version"
  yq e -i ".appVersion = \"${version}\"" ${CHART_DIR}/Chart.yaml
fi
