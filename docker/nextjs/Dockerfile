FROM node

ENV NODE_ENV=development

WORKDIR /usr/src/app

COPY ./src/front/package*.json ./
RUN npm install --legacy-peer-deps

CMD ["npm", "run", "dev"]