# Pin specific version
# Use slim for reduced image size
FROM node:19.6-bullseye-slim AS base

# Specify working directory
WORKDIR /usr/src/app

# Copy only required files
COPY package*.json ./

FROM base AS dev

# Install only production dependencies
# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm ci --only=production

COPY . . 

CMD ["npm", "run", "dev"]

from base AS production

# Set environment variables
ENV NODE_ENV production

# Install only production dependencies
# Use cache mount to speed up install of existing dependencies
RUN --mount=type=cache,target=/usr/src/app/.npm \
  npm set cache /usr/src/app/.npm && \
  npm ci --only=production

# Use non root user
USER node

#Copy remaining source code AFTER installing dependencies
COPY --chown=node:node ./src .

# Expose port
EXPOSE 3000

CMD ["node", "index.js"]