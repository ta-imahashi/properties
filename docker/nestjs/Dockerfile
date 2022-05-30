FROM node

ENV NODE_ENV=development

WORKDIR /usr/src/app

RUN npm i -g @nestjs/cli

COPY ./src/bff/package*.json ./
RUN npm install --legacy-peer-deps

CMD ["npm", "run", "start:dev"]