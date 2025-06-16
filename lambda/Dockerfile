# Start with Amazon Linux 2023 custom runtime base image
FROM public.ecr.aws/lambda/provided:al2023

# Install Git
RUN dnf update -y && \
    dnf install -y git && \
    dnf clean all

# Copy bootstrap script to the correct location
COPY bootstrap /var/runtime/bootstrap

# Ensure the bootstrap file is executable
RUN chmod +x /var/runtime/bootstrap
