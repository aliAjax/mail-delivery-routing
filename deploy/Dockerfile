FROM golang:1.22 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /maildelivery ./cmd/server
FROM gcr.io/distroless/static-debian12
COPY --from=build /maildelivery /maildelivery
COPY --from=build /src/configs /configs
EXPOSE 18089 19089 2525
ENTRYPOINT ["/maildelivery"]
