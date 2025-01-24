This is a simple project that increases the counter or "counter.txt" file and commits that change to Github.

I've deployed it on an EC2 instance and automated it with systemd.

---SETUP---

Stuff that you will need to do if you want to run dailyUpdater.sh on a AWS Linux instance
1. Create local user/group on host
2. Install and setup Git
3. Create ssh key for Github
4. Clone repository
5. Adjust permissions
6. Set up systemd service and timer

---Roadmap---
1. Place counter.txt file in an S3 bucket and interact with it there
2. - Option A - containarize
   - Option B - make serverless w/ AWS Lambda