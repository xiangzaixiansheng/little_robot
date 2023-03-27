FROM registry.cn-hangzhou.aliyuncs.com/hanxiang/golang_basic

WORKDIR /app

COPY . ./
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /little_robot

EXPOSE 3000

CMD [ "/little_robot" ]