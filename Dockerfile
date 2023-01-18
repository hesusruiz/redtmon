FROM golang:1.19-alpine as builder

# Install the buil tools required
RUN apk add --no-cache make gcc musl-dev linux-headers git ca-certificates

# We will build in this directory
WORKDIR /usr/src

# Clone Quorum into the 'quorum' subdirectory.
# The monitor uses Quorum/Ethereum client to connect to the blockchain node
# via JSON-RPC with websockets transport.
# Because Quorum is a fork of go-ethereum we need it locally to be able to compile.
RUN git clone https://github.com/ConsenSys/quorum.git

# The monitor will be compiled in a sibling directory
WORKDIR /usr/src/redtmon

# Copy the go.mod for pre-downloading dependencies and only redownloading them
# in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy redtmon sources
COPY . .

# Build the executable
RUN go build -v -o caddy .

#################################################################
# Create a minimal image by copying files to the caddy base image
#################################################################

FROM caddy:2.6.2

# Overwrite the standard 'caddy' binary with our own
COPY --from=builder /usr/src/redtmon/caddy /usr/bin/caddy

# Use our Caddyfile
# If the RedT node websocket RPC is available locally at "ws://127.0.0.1:22001",
# there is no need to modify it.
# Otherwise, modify it to your liking before building the image.
# You can use anything a standard version of Caddy can do.
COPY --from=builder /usr/src/redtmon/Caddyfile /etc/caddy/Caddyfile

# Copy the HTML files to the /srv/www directory for the Caddy file server
# If you modify the Caddyfile, make sure you also make the corresponding modification here.
COPY --from=builder /usr/src/redtmon/www /srv/www
