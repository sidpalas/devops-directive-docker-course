# Pin specific version for stability
# Use slim for reduced image size
FROM node:19.6-bullseye-slim AS base

# Specify working directory other than /
WORKDIR /usr/src/app

# Copy only files required to install
# dependencies (better layer caching)
COPY package*.json ./

FROM base as dev

RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm install

COPY . .

CMD ["npm", "run", "dev"]

FROM base as production

# Set NODE_ENV
ENV NODE_ENV production

# Install only production dependencies
# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm ci --only=production

# Use non-root user
# Use --chown on COPY commands to set file permissions
USER node

# Copy the healthcheck script
COPY --chown=node:node ./healthcheck/ .

# Copy remaining source code AFTER installing dependencies. 
# Again, copy only the necessary files
COPY --chown=node:node ./src/ .

# Indicate expected port
EXPOSE 3000

CMD [ "node", "index.js" ]
