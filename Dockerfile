FROM golang:1.18 as builder

WORKDIR /usr/src

# Download quorum
RUN git clone https://github.com/ConsenSys/quorum.git

# Download redtmon
RUN git clone https://github.com/hesusruiz/redtmon.git

WORKDIR /usr/src/redtmon
RUN go mod tidy & go mod download && go mod verify

RUN go build -v -o redtmon .

# Now copy it into a minimal distroless base image
FROM gcr.io/distroless/base-debian11
WORKDIR /redtmon
COPY --from=builder /usr/src/redtmon/redtmon /redtmon/redtmon
COPY --from=builder /usr/src/redtmon/www /redtmon/www
COPY --from=builder /usr/src/redtmon/Caddyfile /redtmon/Caddyfile
ENTRYPOINT [ "redtmon", "run" ]
