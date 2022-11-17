#!/bin/bash

# ---------------------------------------------
# -- Publish new version/tag to github
# ---------------------------------------------
set -x # uncomment to debug script
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
readonly TAG=v0.0.8
readonly GITHUB_USER="wcarmon"
readonly PROJECT_NAME="otzap"

# ---------------------------------------------
# -- Publish
# ---------------------------------------------
cd "$PARENT_DIR" >/dev/null 2>&1

go mod tidy
#go clean -modcache

git fetch --all

if [[ "0" -ne "$(git diff-index --quiet HEAD;echo $?)" ]]; then
  echo "ERROR: Working directory is not clean"
  exit 0

else
  echo "ready to tag!"
  exit 0
fi

git tag $TAG

echo
echo "|-- Pushing all tags ..."
git push origin --tags
# Undo: git push --delete origin v0.0.999

echo
echo "|-- Registering ${TAG} ..."
GOPROXY=proxy.golang.org go list -m github.com/${GITHUB_USER}/${PROJECT_NAME}@${TAG}

# ---------------------------------------------
# -- Report
# ---------------------------------------------
echo
echo "|-- Retrieve using "
echo "go get github.com/${GITHUB_USER}/${PROJECT_NAME}@${TAG}"
