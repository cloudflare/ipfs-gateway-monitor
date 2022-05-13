#!/bin/sh

jwtFile="$PINATA_JWT_FILE"

if [ $jwtFile ]; then
    args="-pinata-jwt=$(cat $jwtFile)"
fi

sleepTime=30
echo "#Sleeping for $sleepTime seconds to wait for the IPFS nodes to be ready."
sleep $sleepTime

/usr/local/bin/ipfs-gw-monitor $@ $args
