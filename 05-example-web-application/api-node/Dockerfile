FROM ubuntu

RUN apt-get update && apt-get install nodejs npm -y

COPY . .

RUN npm install

CMD ["node", "index.js"]