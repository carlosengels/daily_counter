# GitHub Streak Updater

A simple automation that updates a counter.txt file daily and pushes the change to GitHub — keeping activity streaks alive, or just for fun.


### Why?
Por que because...

Maybe your favorite color is green, or you vehemently scorn white tiles. 


### How It Works

The script increments a number in files/counter.txt and commits it to GitHub.
It can be run in two ways:

### AWS Lambda

Runs on a custom Docker image, pushed to ECR, and triggered by EventBridge Scheduler.

Tech highlights:

al2023 Lambda runtime (provided.al2023)

Git installed via dnf

Custom bootstrap script to handle event loop

IAM Role with lambda:InvokeFunction permissions
s
Schedule defined using aws scheduler create-schedule

⏱️ Note: Lambda timeout needed to be increased from 3 seconds to 1 minute to allow for git clone, commit, and push.


### systemd Service (Local)

Can also run locally as a recurring systemd timer.

Setup:

1. Create a dedicated user/group

2. Install & configure Git

3. Set up an SSH key for GitHub access

4. Clone the repo

5. Adjust file and directory permissions

6. Deploy *.service and *.timer files to systemd

