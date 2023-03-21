#!/usr/bin/env bash

message=${1:-"nothing"}
version=${2:-"skip"}

git add .
git commit -m "${version} - ${message}"
if [ "${version}" != "skip" ]; then
  git tag "v${version}"
fi
git push origin main --tags