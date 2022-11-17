#!/bin/bash

# -- Usage:
# 1. Update code/tests/docs
# 2. Auto-format
# 3. Bump version in VERSION file
# 4. Commit
# 5. Run this file

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
readonly TAG=$(cat $PARENT_DIR/VERSION)
readonly GITHUB_USER="wcarmon"
readonly PROJECT_NAME="otzap"

# ---------------------------------------------
# -- Validate
# ---------------------------------------------
#TODO: validate tag

# ---------------------------------------------
# -- Publish
# ---------------------------------------------
cd "$PARENT_DIR" >/dev/null 2>&1

go mod tidy
#go clean -modcache

git fetch --all --tags
git push origin HEAD

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
git tag $TAG || true
# Undo: git tag --delete v0.0.x

echo
echo "|-- Local tags"
git tag

echo
echo "|-- Pushing all tags ..."
git push origin --tags
# Undo: git push --delete origin v0.0.999
git fetch --all --tags

echo
echo "|-- Registering ${TAG} ..."
GOPROXY=proxy.golang.org go list -m github.com/${GITHUB_USER}/${PROJECT_NAME}@${TAG}

# ---------------------------------------------
# -- Report
# ---------------------------------------------
echo
echo "|-- Retrieve using ..."
echo "go get github.com/${GITHUB_USER}/${PROJECT_NAME}@${TAG}"

echo
echo "|-- See docs ..."
echo "https://pkg.go.dev/github.com/${GITHUB_USER}/${PROJECT_NAME}"
