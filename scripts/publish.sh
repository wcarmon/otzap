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
readonly TAG=v0.0.9
readonly GITHUB_USER="wcarmon"
readonly PROJECT_NAME="otzap"

# ---------------------------------------------
# -- Publish
# ---------------------------------------------
cd "$PARENT_DIR" >/dev/null 2>&1

go mod tidy
#go clean -modcache

git fetch --all

CLEAN_WORKSPACE_INDICATOR=$(
  git diff-index --quiet HEAD
  echo $?
)
if [[ "0" -ne "$CLEAN_WORKSPACE_INDICATOR" ]]; then
  echo "ERROR: Working directory is not clean, commit first"
  exit 1
fi

echo
echo "|-- Tagging ..."
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
