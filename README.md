# GitHub Streak Updater

A simple automation that updates a counter.txt file daily and pushes the change to GitHub — keeping activity streaks alive.

Runs on a custom Docker image on Lambda, pushed to ECR, and triggered by EventBridge Scheduler.

Tech highlights:

- al2023 Lambda runtime (provided.al2023)

- Git installed via dnf

- Custom bootstrap script to handle event loop

- IAM Role with lambda:InvokeFunction permissions

- Schedule defined using aws scheduler create-schedule

Note: Lambda timeout needed to be increased from 3 seconds to 1 minute to allow for git clone, commit, and push.

### Why?
Por que because...

Maybe your favorite color is green, or you vehemently scorn white tiles. 


### How It Works

The script increments a number in files/counter.txt and commits it to GitHub.
It can be run in two ways:

### AWS Lambda Set up

## 1. Clone repository and cd into lambda dir

## 2.  Building the docker image
# Authenticate Docker to your ECR registry (expires after 12 hours)
`aws ecr get-login-password --region <region> \
  | docker login --username AWS --password-stdin <account-ID>.dkr.ecr.<region>.amazonaws.com`

# Build the Docker image for Lambda (specify platform for compatibility)
docker buildx build --platform linux/amd64 --load -t lambda-git-update .

# Tag the image for your ECR repository
docker tag lambda-git-update <account-ID>.dkr.ecr.us-west-2.amazonaws.com/lambda-git-update:latest

# Push the image to your ECR repo
docker push <account-ID>.dkr.ecr.us-west-2.amazonaws.com/lambda-git-update:latest

## 3. Create a Github token and add it to Secrets Manager
`aws secretsmanager create-secret \
  --name GitHubToken \
  --secret-string '{"token":"ghp_your_actual_token_here"}'
`
## 4. Deploy Cloudformation template

## 5. Test
`aws lambda invoke \
  --function-name your-lambda-function-name \
  --invocation-type Event \
  --payload '{}' \
  response.json
`

### Alertnative setup: systemd Service

Can also run locally as a recurring systemd timer.

Setup:

1. Create a dedicated user/group

2. Install & configure Git

3. Set up an SSH key for GitHub access

4. Clone the repo

5. Adjust file and directory permissions

6. Deploy *.service and *.timer files to systemd

