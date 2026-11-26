# identity. Generated from testenv/topology/topology.yaml by testenv/topology/generate.js. Do not edit by hand.
# Built with the repo root as the context.
FROM golang:1.23 AS build
WORKDIR /src
COPY . .
ENV GOWORK=off CGO_ENABLED=0
RUN go build -o /out/app .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/app /app
USER 65532
EXPOSE 8080
ENTRYPOINT ["/app"]
