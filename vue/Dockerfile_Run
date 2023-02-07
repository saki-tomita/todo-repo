FROM node:18.1-alpine

COPY ./app/dist ./dist
WORKDIR /app

RUN apk update && \
        npm install -g npm && \
        npm install -g @vue/cli && \
        npm install axios

ENV HOST 0.0.0.0
EXPOSE 8080

