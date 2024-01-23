# Pin specific version
# Use slim for reduced image size
FROM node:19.6-bullseye-slim

# Set environment variables
ENV NODE_ENV production

# Specify working directory
WORKDIR /usr/src/app

# Copy only required files
COPY package*.json ./

# Install only production dependencies
RUN npm ci --only=production

# Use non root user
USER node

#Copy remaining source code AFTER installing dependencies
COPY --chown=node:node ./src .

CMD ["node", "index.js"]