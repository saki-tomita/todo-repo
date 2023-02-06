FROM node:18.1-alpine

WORKDIR /app

RUN apk update && \
        npm install -g npm && \
        npm install -g @vue/cli && \
        npm install axios

EXPOSE 8080