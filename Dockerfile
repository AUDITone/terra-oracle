from golang:alpine as build

workdir /app

copy . .

run go install ./cmd/terra-oracle

# user your terraproject/core image
from terra as build-env

FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates rsync jq curl fish

# Copy over binaries from the build-env
COPY --from=build-env /usr/bin/terrad /usr/bin/terrad
COPY --from=build-env /usr/bin/terracli /usr/bin/terracli
COPY --from=build /go/bin/terra-oracle /usr/bin/terra-oracle

