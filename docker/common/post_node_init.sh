#!/bin/sh
set -e
set -o pipefail

operator_address() {
    lavad q staking validators -o json | jq -r '.validators[0].operator_address'
}

check_votes_status() {
  lavad q gov proposals --output json | jq -r '.proposals[] | select(.status == "PROPOSAL_STATUS_VOTING_PERIOD")' 
}

# Function to get the current account sequence
get_account_sequence() {
  local account=$1
  lavad query account $(lavad keys show $account -a) -o json | jq -r '.sequence'
}

# get_base_specs() {
#     local priority_specs=(
#         "specs/mainnet-1/specs/ibc.json"
#         "specs/mainnet-1/specs/cosmoswasm.json"
#         "specs/mainnet-1/specs/tendermint.json"
#         "specs/mainnet-1/specs/cosmossdk.json"
#         "specs/testnet-2/specs/cosmossdkv45.json"
#         "specs/testnet-2/specs/cosmossdk_full.json"
#         "specs/mainnet-1/specs/cosmossdkv50.json"
#         "specs/mainnet-1/specs/ethermint.json"
#         "specs/mainnet-1/specs/ethereum.json"
#         "specs/mainnet-1/specs/solana.json"
#         "specs/mainnet-1/specs/aptos.json"
#         "specs/mainnet-1/specs/btc.json"
#     )

#     (IFS=,; echo "${priority_specs[*]}")
# }

vote_yes_on_all_pending_proposals() {
  echo "Waiting for at least one proposal to be active"
  while true; do
    sleep 1
    res=$(check_votes_status)
    if [ -n $res ]; then
      echo "No active proposals yet"
    else
      echo "Found active proposal!"
      lavad q gov proposals --output json | jq -r '.proposals[] | select(.status == "PROPOSAL_STATUS_VOTING_PERIOD") | .id' | while read -r proposal_id; do
        echo "$FROM voting yes on proposal id: $proposal_id"
        lavad tx gov vote $proposal_id yes -y --from $FROM --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE > /dev/null &
      done
      lavad q gov proposals --output json | jq -r '.proposals[] | select(.status == "PROPOSAL_STATUS_VOTING_PERIOD") | .id' | while read -r proposal_id; do
        echo "Waiting for proposal id: $proposal_id to pass..."
        until [ "$(lavad q gov proposal $proposal_id --output json | jq -r .status)" == "PROPOSAL_STATUS_PASSED" ]; do
          echo "Proposal id: $proposal_id didn't pass yet..."
          sleep 1
        done
        echo "Proposal id: $proposal_id finally passed!"
      done
      break
    fi
  done
}

echo "### Starting post node init ###"
wget -O jq https://github.com/jqlang/jq/releases/download/jq-1.7.1/jq-linux-amd64
chmod +x ./jq
mv jq /usr/bin

GASPRICE="${GASPRICE:-0.000000001ulava}"
FROM="${FROM:-servicer1}"
NODE="${NODE:-tcp://lava-node:26657}"

lavad config node $NODE
(
cd /lava/specs/mainnet-1/specs/
# specs=$(get_base_specs)
lavad tx gov submit-legacy-proposal spec-add ./ibc.json,./tendermint.json --lava-dev-test -y --from $FROM --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE
)
sleep 8
vote_yes_on_all_pending_proposals


cd /lava/specs/mainnet-1/specs/
lavad tx gov submit-legacy-proposal spec-add ./cosmossdk.json --lava-dev-test -y --from $FROM --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE
sleep 5
vote_yes_on_all_pending_proposals

cd /lava/specs/mainnet-1/specs/
lavad tx gov submit-legacy-proposal spec-add ./lava.json,./sui.json --lava-dev-test -y --from $FROM --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE
sleep 5
vote_yes_on_all_pending_proposals

echo "Adding plan: DefaultPlan"
lavad tx gov submit-legacy-proposal plans-add /lava/cookbook/plans/test_plans/default.json -y --from $FROM --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE
sleep 5
vote_yes_on_all_pending_proposals

echo "Buying plan: DefaultPlan for $FROM"
lavad tx subscription buy DefaultPlan $(lavad keys show $FROM -a) --enable-auto-renewal -y --from $FROM --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE 2> /dev/null

sleep 20
echo "Staking provider"
PROVIDERSTAKE="500000000000ulava"

# Stake SUI JSON-RPC provider on port 80
echo "Staking SUI JSON-RPC provider1"
SUI_PROVIDER_ADDRESS="nginx:80"
echo "Current sequence for servicer1: $sequence"
lavad tx pairing stake-provider SUIJSONRPC $PROVIDERSTAKE "$SUI_PROVIDER_ADDRESS,2" 2 $(operator_address) -y --delegate-commission 50 --from servicer1 --provider-moniker "servicer1-sui" --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE

sleep 20
# Stake LAV1 provider on port 81 (this will cover all LAV1 API interfaces)
echo "Staking LAV1 provider"
LAV_PROVIDER_ADDRESS="nginx:81"
echo "Current sequence for servicer2: $sequence"
lavad tx pairing stake-provider LAV1 $PROVIDERSTAKE "$LAV_PROVIDER_ADDRESS,1" 1 $(operator_address) -y --delegate-commission 50 --from servicer2 --provider-moniker "servicer2-lav" --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE

sleep 20
echo "Staking SUI JSON-RPC provider2"
SUI_PROVIDER_ADDRESS="nginx:82"
echo "Current sequence for servicer2: $sequence"
lavad tx pairing stake-provider SUIJSONRPC $PROVIDERSTAKE "$SUI_PROVIDER_ADDRESS,1" 1 $(operator_address) -y --delegate-commission 50 --from servicer2 --provider-moniker "servicer2-sui" --gas-adjustment "1.5" --gas "auto" --gas-prices $GASPRICE


echo "### Post node init finished successfully ###"

# Exit with success
exit 0