FROM node:18.1.0-alpine
COPY . /root
WORKDIR /root

RUN apk update && \
    npm install -g npm && \
    npm install -g @vue/cli

EXPOSE 8080