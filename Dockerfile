from golang:alpine as build

workdir /app

copy . .

run go install ./cmd/terra-oracle

FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates rsync jq curl fish

# Copy over binaries from the build
COPY --from=build /go/bin/terra-oracle /usr/bin/terra-oracle

cmd ["terra-oracle"]
