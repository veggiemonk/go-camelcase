#!/usr/bin/env bash

TAG="$(echo "$(date +"%Y-%m-%d")-$(git rev-parse --short HEAD)")"
echo "Tagging with $TAG"

MSG=${1:-"releasing $TAG"}

echo git tag -a "$TAG" -m "$MSG"
