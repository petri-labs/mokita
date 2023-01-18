
#!/bin/bash
# microtick and bitcanna contributed significantly here.
# rocksdb doesn't work yet

# PRINT EVERY COMMAND
set -ux

# uncomment the three lines below to build mokita

# export GOPATH=~/go
# export PATH=$PATH:~/go/bin
# go install -./...


# MAKE HOME FOLDER AND GET GENESIS
mokitad init test --chain-id moki-test-4 
# wget -O ~/.mokitad/config/genesis.json https://cloudflare-ipfs.com/ipfs/QmXRvBT3hgoXwwPqbK6a2sXUuArGM8wPyo1ybskyyUwUxs

INTERVAL=1500

# GET TRUST HASH AND TRUST HEIGHT

LATEST_HEIGHT=$(curl -s https://rpc-test.osmosis.zone/block | jq -r .result.block.header.height);
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL))
TRUST_HASH=$(curl -s "https://rpc-test.osmosis.zone/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"


# export state sync vars
export MOKISISD_P2P_MAX_NUM_OUTBOUND_PEERS=200
export MOKISISD_STATESYNC_ENABLE=true
export MOKISISD_STATESYNC_RPC_SERVERS="https://rpc-test.osmosis.zone:443,https://rpc-test.osmosis.zone:443"
export MOKISISD_P2P_SEEDS="0f9a9c694c46bd28ad9ad6126e923993fc6c56b1@137.184.181.105:26656"
export MOKISISD_P2P_PERSISTENT_PEERS="4ab030b7fd75ed895c48bcc899b99c17a396736b@137.184.190.127:26656,3dbffa30baab16cc8597df02945dcee0aa0a4581@143.198.139.33:26656"
export MOKISISD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export MOKISISD_STATESYNC_TRUST_HASH=$TRUST_HASH



# THERE, NOW IT'S SYNCED AND YOU CAN PLAY
mokitad start 
