#!/bin/sh

repo="$IPFS_PATH"
port="$IPFS_PORT"

# We need to initialize the repo first.
if [ -e "$repo/config" ]; then
  echo "Found IPFS fs-repo at $repo"
else
  /usr/local/bin/ipfs init
fi

# We need to set the multiaddr to listen on and advertise.
/usr/local/bin/ipfs config Addresses.Swarm --json "[\"/ip4/0.0.0.0/tcp/$port\"]"

/usr/local/bin/ipfs-daemon $@
