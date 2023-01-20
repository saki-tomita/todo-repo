FROM node:18.1.0-alpine
COPY ./app /app
WORKDIR /app

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli

EXPOSE 8080