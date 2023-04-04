FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY /front-app/build .
COPY /data /data