FROM golang:1.25-alpine AS build
WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache git ca-certificates
RUN update-ca-certificates

COPY . ./

RUN echo "building for GOOS: $TARGETOS, GOARCH: $TARGETARCH"
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -tags timetzdata -o go_ci

# Final stage
FROM scratch AS final
WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/go_ci ./
COPY ./configurations /configurations

ENV TZ=Asia/Bangkok
VOLUME /certs
EXPOSE 8179
ENTRYPOINT ["./go_ci"]
