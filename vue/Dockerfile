

FROM nginx:1.15.12-alpine
WORKDIR /vue
COPY ./app/dist /usr/share/nginx/html
COPY ./app/nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]