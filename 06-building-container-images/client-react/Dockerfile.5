# syntax=docker/dockerfile:1.5

FROM node:19.4-bullseye AS build

# Specify working directory other than /
WORKDIR /usr/src/app

# Copy only files required to install
# dependencies (better layer caching)
COPY package*.json ./

# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm install

COPY . .

RUN npm run build

# Use separate stage for deployable image
FROM nginxinc/nginx-unprivileged:1.23-alpine-perl

# Use COPY --link to avoid breaking cache if we change the second stage base image
COPY --link nginx.conf /etc/nginx/conf.d/default.conf

COPY --link --from=build usr/src/app/dist/ /usr/share/nginx/html

EXPOSE 8080