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
  echo "bumping appVersion of chart to $version"
  CHART_VERSION_OLD=$(cat ${CHART_DIR}/Chart.yaml | yq '. | .version')
  CHART_VERSION_PATCH_INCREASE=$(echo "$CHART_VERSION_OLD" | awk -F. '{$NF = $NF + 1;} 1' | sed 's/ /./g')
  APP_VERSION_OLD=$(cat ${CHART_DIR}/Chart.yaml | yq '. | .appVersion')
  if [[ "${APP_VERSION_OLD}" != "${version}" ]]; then
    echo "bumping helm chart version from $CHART_VERSION_OLD to $CHART_VERSION_PATCH_INCREASE"
    yq e -i ".version = \"${CHART_VERSION_PATCH_INCREASE}\"" ${CHART_DIR}/Chart.yaml
  fi
  yq e -i ".appVersion = \"${version}\"" ${CHART_DIR}/Chart.yaml
fi
