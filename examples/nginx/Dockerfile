#FROM nginx:1.16-alpine
FROM nginx:1.17
#RUN yum -y update && yum install -y curl
#RUN apt -y update && apt install -y curl
RUN apt-get update && apt-get install -y curl && apt-get clean
WORKDIR /usr/share/nginx/html/
COPY index.html .
EXPOSE 80
CMD ["curl", "localhost"]
#CMD ["cat", "index.html"]
