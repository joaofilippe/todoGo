FROM golang as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENV ENV=PROD

EXPOSE 3000
EXPOSE 5432

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./cmd/server ./cmd/main.go

FROM scratch
COPY --from=builder /app/cmd/server /app/cmd/server
COPY --from=builder /app/config /app/config

WORKDIR /app/cmd
ENTRYPOINT [ "./server" ]