FROM node:18.1-alpine as vue-local
WORKDIR /app
COPY ./app/package.json ./
COPY ./app/yarn.lock ./
RUN apk update && \
    npm install -g npm && \
    yarn install --frozen-lockfile && \
    yarn add @vue/cli && \
    yarn build

FROM nginx:1.15.12-alpine as vue-run
COPY --from=vue-local /app/dist/ /usr/share/nginx/html/
COPY --from=vue-local /app/nginx/ /etc/nginx/conf.d/
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]