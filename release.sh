#!/usr/bin/env bash

# Todo: Add function to manage and automate the checksums for the binrary
# Version needs to be specified as first argument

version="$1"

if [[ -z "$version" ]]; then
	exit 1
fi

if ! command -v "git" >/dev/null 2>&1; then
	echo ":: Git package is not installed..."
	exit 1
fi

git add .
git commit -m "Added Version: $version"
git push origin main

git tag "$version"
git push origin "$version"
