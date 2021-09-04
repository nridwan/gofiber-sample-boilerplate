#!/bin/bash
# Script for building all build types in a finger snap
# args
# $1 - first argument, set JDK path
if [ "$1" "!=" "" ]; then
	VERSION=$(date +%s)
	touch "migration/${VERSION}_${1}.up.sql"
	touch "migration/${VERSION}_${1}.down.sql"
else
	exit 1
fi
