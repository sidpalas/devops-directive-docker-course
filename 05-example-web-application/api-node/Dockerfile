# Pin specific version
# Use slim for reduced image size
FROM node:19.6-bullseye-slim

# Specify working directory
WORKDIR /usr/src/app

# Copy only required files
COPY package*.json ./

RUN npm install

# Use non root user
USER node

#Copy remaining source code AFTER installing dependencies
COPY --chown=node:node ./src .

CMD ["node", "index.js"]