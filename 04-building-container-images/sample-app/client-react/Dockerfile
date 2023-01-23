# DEV Dockerfile

FROM node:19.4-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .
EXPOSE 5173

CMD ["npm", "run", "dev"]