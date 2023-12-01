ARG PYIPAD_IMAGE
ARG PYIPADMINER_IMAGE

FROM ${PYIPAD_IMAGE} as pyipad
FROM ${PYIPADMINER_IMAGE} as pyrinminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/Pyrinpyi/pyipad

WORKDIR /go/src/github.com/Pyrinpyi/pyipad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=pyipad /app/ /app/
COPY --from=pyrinminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/Pyrinpyi/pyipad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
