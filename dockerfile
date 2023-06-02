FROM golang:1.20 as builder

WORKDIR /gin-api-rest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY templates/ templates/
COPY assets/ assets/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /gin-api-rest/main .
COPY --from=builder /gin-api-rest/templates/ templates/
COPY --from=builder /gin-api-rest/assets/ assets/

CMD ["./main"]
