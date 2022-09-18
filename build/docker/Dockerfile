FROM centos:centos7

COPY ./blog-backend /home/blog/backend/bin
WORKDIR /home/blog/backend/

RUN wget https://golang.google.cn/dl/go1.17.13.linux-amd64.tar.gz \
    && rm -rf /usr/local/go \
    && tar -zxf go1.17.13.linux-amd64.tar.gz -C /usr/local \
    && rm -f go1.17.13.linux-amd64.tar.gz \
    && export GOROOT=/usr/local/go \
    && mkdir -p /home/GoProject \
    && export GO111MODULE=on \
    && export GOPROXY=https://goproxy.cn,direct \
    && mkdir conf

EXPOSE 3000

ENTRYPOINT ['./blog-backend']