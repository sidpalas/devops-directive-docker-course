# Pin base image
FROM node:19.4-bullseye

WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

COPY . .

RUN npm install

CMD ["npm", "run", "dev"]