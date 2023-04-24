#!/bin/bash -e

if [[ $# != 1 ]]; then
  echo "must send the required tag as a single parameter"
  exit 1
fi

TAG=$1
git tag "${TAG}"
VERSION="${TAG}" make build-all
