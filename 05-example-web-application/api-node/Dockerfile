# Pin specific version
# Use slim for reduced image size
FROM node:19.6-bullseye-slim

# Specify working directory
WORKDIR /usr/src/app

COPY . .

RUN npm install

CMD ["node", "index.js"]