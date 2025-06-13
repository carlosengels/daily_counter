# Start with Amazon Linux 2023 custom runtime base image
FROM public.ecr.aws/lambda/provided:al2023

# Install Git
RUN yum update -y && \
    yum install -y git && \
    yum clean all

# Copy your Bash script
COPY bootstrap /var/task/bootstrap

# Ensure the bootstrap file is executable
RUN chmod +x /var/task/bootstrap

# Lambda will run /var/task/bootstrap as the entry point
CMD ["/var/task/bootstrap"]
