FROM golang:1.24-alpine AS build
# Allow go to retrieve the dependencies for the build step
WORKDIR /go-modules
RUN apk update && apk upgrade && apk add --no-cache git && apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY . ./
# Compile the binary, we don't want to run the cgo resolver

RUN echo "building for GOOS: $TARGETOS, GOARCH: $TARGETARCH"
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -tags timetzdata -mod=vendor -a -o lead_raptor_api

# final stage
FROM scratch AS final
WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go-modules/lead_raptor_api .
COPY ./configurations /configurations
COPY ./scripts /scripts




ENV TZ=Asia/Bangkok

VOLUME /configurations
VOLUME /scripts
VOLUME /certs
EXPOSE 8179
ENTRYPOINT ["./lead_raptor_api"]
