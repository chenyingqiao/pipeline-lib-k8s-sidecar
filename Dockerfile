FROM golang:1.18 as build
WORKDIR /app
ENV GOPROXY=https://goproxy.cn
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o pipeline-common-k8s-sidecar --ldflags "-extldflags -static" main.go

FROM busybox
WORKDIR /app
COPY --from=build /app/pipeline-common-k8s-sidecar /app/
RUN chmod +x /app/pipeline-common-k8s-sidecar
EXPOSE 9999
CMD /app/pipeline-common-k8s-sidecar