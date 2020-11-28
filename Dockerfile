FROM golang:1.14-alpine AS build

WORKDIR /src

# Layer caching shit.
COPY go.mod .
COPY go.sum .

# Build a static binary.
COPY . .
RUN go mod vendor
RUN CGO=0 go install \
    -mod=vendor \
    -ldflags '-s' -tags 'netgo osusergo' \
    ./...

# Verify if the binary is truly static.
RUN ldd /go/bin/plantai 2>&1 | grep -q 'Not a valid dynamic program'

# Runeet.
FROM alpine
COPY --from=build /go/bin/plantai /usr/local/bin
EXPOSE 8080
ENTRYPOINT ["plantai", "serve"]
