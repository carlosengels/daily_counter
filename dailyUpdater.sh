#!/bin/bash

git pull

go run updater.go

commit_message=$(date +"%Y-%m-%d")

git add .

git commit -m "$commit_message"

git push origin main

date > last_run.txt
