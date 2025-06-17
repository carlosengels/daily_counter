### WTH
This is a simple project that increases the counter or "counter.txt" file and commits that change to Github.

### Why? Who is this for?
Por que because.

Maybe your favorite color is green or you vehemently scorn white tiles. 

### But how?

Either as a lambda function or a systemd service


## Lambda
Uses a Docker image, which is pushed to ECR.

# Notes
1. Timeout for lambda function was timing out at 3 seconds. I increased it to 1 minute and it works.

## Systemd service
# SETUP
1. Create local user/group on host
2. Install and setup Git
3. Create ssh key for Github
4. Clone repository
5. Adjust permissions
6. Set up systemd service and timer