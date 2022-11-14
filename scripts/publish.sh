#!/bin/bash

# ---------------------------------------------
# -- Publish new version/tag to github
# ---------------------------------------------
#set -x # uncomment to debug script
set -e # exit on first error
set -o pipefail
set -u # fail on unset var

# ---------------------------------------------
# -- Constants
# ---------------------------------------------
readonly PARENT_DIR=$(readlink -f "$(dirname "${BASH_SOURCE[0]}")/..")

# ---------------------------------------------
# -- Script arguments
# ---------------------------------------------
readonly TAG=v0.0.4
readonly GITHUB_USER="wcarmon"
readonly PROJECT_NAME="otzap"

# ---------------------------------------------
# -- Publish
# ---------------------------------------------
cd "$PARENT_DIR/src" >/dev/null 2>&1

go mod tidy
go clean -modcache

git fetch --all
git tag $TAG

echo
echo "|-- Pushing"
git push origin $TAG || true

GOPROXY=proxy.golang.org go list -m github.com/${GITHUB_USER}/${PROJECT_NAME}@${TAG}

# ---------------------------------------------
# -- Report
# ---------------------------------------------
echo
echo "|-- Retrieve using "
echo "go get github.com/${GITHUB_USER}/${PROJECT_NAME}@${TAG}"
