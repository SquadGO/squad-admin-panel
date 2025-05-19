FROM golang:1.24-alpine AS build

WORKDIR /backend

COPY go.mod go.sum ./

RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
  --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
  go mod download

COPY . .

RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
  --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux go build -o /server cmd/server/main.go

FROM alpine:3.18

WORKDIR /
COPY --from=build /server /server

CMD ["/server"]