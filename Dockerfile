FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY *.go go.* views /src/
ADD views /src/views
RUN CGO_ENABLED=0 go build -o /tmp/Bon

FROM scratch
COPY --from=build /tmp/Bon  /opt/Bon
COPY --from=build /src/views /opt/views
WORKDIR /opt
ENTRYPOINT ["/opt/Bon"]
