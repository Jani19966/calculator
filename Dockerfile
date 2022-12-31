from golang:latest

WORKDIR /go/src/github.com/jani19966/calculator

COPY backend backend
COPY frontend frontend
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build -o /bin/calculator

CMD ["/bin/calculator"]