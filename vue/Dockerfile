FROM node:18.1-alpine as vue-local
WORKDIR /vue/app
COPY package.json ./
COPY yarn.lock ./
RUN apk update && \
        npm install -g npm && \
        npm install -g @vue/cli && \
        npm install axios && \
       yarn install \
COPY . .
RUN yarn build

FROM nginx:1.15.12-alpine as vue-run
WORKDIR /vue/app
COPY --from=vue-local /dist/ /usr/share/nginx/html/
COPY --from=vue-local /nginx/ /etc/nginx/conf.d/
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]