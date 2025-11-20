FROM golang:1.25

# Install Air BEFORE entering your module
RUN go install github.com/air-verse/air@latest

ENV PATH="/go/bin:${PATH}"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["./air.sh"]
