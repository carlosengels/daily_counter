#!/bin/bash
python3 updater.py

commit_message=$(date +"%Y-%m-%d")

git add .

git commit -m "$commit_message"

git push origin main
