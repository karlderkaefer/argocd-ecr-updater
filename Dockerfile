FROM golang:1.21 as BASE
WORKDIR /go/src/github.com/karlderkaefer/argocd-image-updater/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/argocd-ecr-updater

FROM gcr.io/distroless/static-debian11
EXPOSE 8080
WORKDIR /app
COPY --from=BASE /bin/argocd-ecr-updater .
CMD ["./argocd-ecr-updater"]
