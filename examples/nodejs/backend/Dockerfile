FROM node:10.15.3-alpine

RUN apk add curl
WORKDIR /app
EXPOSE 3000
CMD ["npm", "run", "dev"]

COPY package* ./
RUN npm ci
COPY . .
