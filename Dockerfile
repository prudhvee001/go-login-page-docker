FROM golang:1.21 AS builder
WORKDIR /mpr-app
COPY . .
RUN go build -o myloginapp

FROM alpine:latest
WORKDIR /mpr-app
COPY --from=builder /mpr-app/myloginapp .
EXPOSE 8080
CMD ["./myloginapp"]
