FROM node:18.1-alpine

WORKDIR /app

RUN apk update && \
        npm install -g npm && \
        npm install -g @vue/cli && \
        npm install axios

ENV HOST 0.0.0.0
EXPOSE 8080

CMD ["npm", "start"]