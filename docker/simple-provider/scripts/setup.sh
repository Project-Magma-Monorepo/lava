#!/bin/sh
set -e

# Check if the wallet already exists
if ! lavap keys show ${PROVIDER_WALLET:-servicer1} --keyring-backend ${KEYRING_BACKEND:-test} > /dev/null 2>&1; then
  echo "Creating new wallet: ${PROVIDER_WALLET:-servicer1}"
  lavap keys add ${PROVIDER_WALLET:-servicer1} --keyring-backend ${KEYRING_BACKEND:-test}
  echo "⚠️ IMPORTANT: Save the mnemonic phrase above! ⚠️"
  echo "Wallet created. Please fund this address with LAVA tokens."
else
  echo "Wallet ${PROVIDER_WALLET:-servicer1} already exists."
fi

# Start the provider service
exec lavap rpcprovider /lava/config/provider.yml \
  --from "${PROVIDER_WALLET:-servicer1}" \
  --geolocation "${GEOLOCATION:-2}" \
  --chain-id "${CHAIN_ID:-lava-testnet-2}" \
  --keyring-backend "${KEYRING_BACKEND:-test}" \
  --node "${LAVA_RPC_NODE:-https://public-rpc.lavanet.xyz:443/rpc/}" \
  --log_level "debug"