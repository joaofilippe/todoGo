FROM golang as builder

WORKDIR /app

COPY ./config/.env /app/config/.env

COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/main.go

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT [ "/server" ]