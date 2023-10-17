FROM golang:1.21-alpine as builder

WORKDIR /app
COPY * ./
RUN go mod download
RUN go build -o /ecoforest-exporter

FROM scratch
EXPOSE 2112
COPY --from=builder /ecoforest-exporter /ecoforest-exporter

CMD [ "/ecoforest-exporter" ]
