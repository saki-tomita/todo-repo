FROM node:18.1-alpine as vue-local
WORKDIR /vue/app
COPY ./app/package.json ./
#COPY ./app/package-lock.json ./
COPY ./app/yarn.lock ./
RUN apk update && \
    npm install -g npm && \
    #npm install -g @vue/cli && \
    #npm install axios && \
    yarn add -g @vue/cli && \
    yarn install && \
    #yarn install --prod --frozen-lockfile && \
    #yarn upgrade && \
    yarn run build
    #npm run build

FROM nginx:1.15.12-alpine as vue-run
COPY --from=vue-local /dist/ /usr/share/nginx/html/
COPY --from=vue-local /nginx/ /etc/nginx/conf.d/
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]