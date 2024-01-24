# Pin base image
FROM node:19.4-bullseye

WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm install

COPY . .

CMD ["npm", "run", "dev"]