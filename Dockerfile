FROM golang:1.27 as BASE
WORKDIR /go/src/github.com/karlderkaefer/argocd-image-updater/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/argocd-ecr-updater

# Hardened, minimal base for the statically linked binary (CGO_ENABLED=0).
# chainguard/static ships a CA-certificates bundle and runs as the non-root
# user "nonroot" (uid/gid 65532) by default. It has no shell and no package
# manager, which removes the OS package layer entirely (no Debian tzdata, so
# no recurring tzdata advisories).
FROM cgr.dev/chainguard/static:latest
EXPOSE 8080
WORKDIR /app
COPY --from=BASE /bin/argocd-ecr-updater .
CMD ["./argocd-ecr-updater"]
