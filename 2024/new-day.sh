#!/bin/bash
DAY=$(printf "%02d" $1)
mkdir -p src/$DAY
cp -r src/template/* src/$DAY/
