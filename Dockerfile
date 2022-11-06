FROM golang:1.19 as BASE
WORKDIR /go/src/github.com/karlderkaefer/argocd-image-updater/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /bin/argocd-ecr-updater

FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=BASE /bin/argocd-ecr-updater .
CMD ["./argocd-ecr-updater"]
