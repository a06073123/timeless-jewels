FROM nginx
COPY timeless.conf /etc/nginx/conf.d/default.conf
COPY frontend/build/ /usr/share/nginx/html