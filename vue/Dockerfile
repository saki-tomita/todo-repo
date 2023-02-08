FROM node:18.1-alpine as vue-local
WORKDIR /vue/app
RUN apk update && \
        npm install -g npm && \
        npm install -g @vue/cli && \
        npm install axios

FROM nginx:1.15.12-alpine as vue-run
WORKDIR /vue/app
COPY --from=vue-local /app/dist /usr/share/nginx/html
COPY --from=vue-local /app/nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]