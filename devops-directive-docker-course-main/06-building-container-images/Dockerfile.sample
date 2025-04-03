# syntax=docker/dockerfile:1.5
# escape=\
# ^ OPTIONAL "directives" (must be at top if used)

# THIS IS A COMMENT

# ARG is the only instruction that can come before FROM 
ARG BASE_IMAGE_TAG=19.4
# ARGs can be overriden at build time 
# > docker build --build-arg BASE_VERSION=19.3 .

FROM node:${BASE_IMAGE_TAG}

LABEL org.opencontainers.image.authors="sid@devopsdirective.com"

RUN echo "Hey Team 👋 (shell form)"
RUN ["echo", "Hey Team 👋 (exec form)"]

# Heredocs allow for specifying multiple commands to 
# be run within a single step, across multiple lines
# without lots of && and \
RUN <<EOF
apt update
apt install iputils-ping -y
EOF

# --mount allows for mounting additional files
# into the build context
# RUN --mount=type=bind ...
# RUN --mount=type=cache ...
# RUN --mount=type=ssh ...
RUN --mount=type=secret,id=secret.txt,dst=/container-secret.txt \
  echo "Run the command that requires access to the secret here"

# Available only at build time
# (Still in image metadata though...)
ARG BUILD_ARG=foo

# Available at build and run time
ENV ENV_VAR=bar

# Set the default working directory 
# Use the convention of your language/framework 
WORKDIR path/to/the/working/directory

ENTRYPOINT [ "echo", "Hey Team 👋 (entrypoint)" ]
CMD [ "+ (cmd)" ]