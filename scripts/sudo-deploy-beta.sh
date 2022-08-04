# Build sonr
sudo ignite chain build -t linux:amd64 -o release --release

#  Unzip it
sudo tar -xzvf release/sonr_linux_amd64.tar.gz -C release

nodes=( v1-beta )

# Stop the existing binaries
for i in "${nodes[@]}"; do ssh root@$(dig "$i".sonr.ws +short) 'kill $(pidof sonrd)'; done

# Copy the binary over
# scp sonrd root@143.198.29.209:~/sonrd
for i in "${nodes[@]}"; do scp release/sonrd root@$(dig "$i".sonr.ws +short):~/sonrd; done

# Copy over the setup_chain_dev.sh script
for i in "${nodes[@]}"; do scp scripts/setup_chain_dev.sh root@$(dig "$i".sonr.ws +short):~/setup_chain_dev.sh; done

# Move the binary to the correct location
# scp sonrd root@143.198.29.209:~/sonrd
for i in "${nodes[@]}"; do ssh root@$(dig "$i".sonr.ws +short) 'sudo mv sonrd /usr/bin/sonrd'; done

# start the new binary
# for i in "${nodes[@]}"; do ssh root@$(dig "$i".sonr.ws +short) 'tmux new -A sonrd start --rpc.laddr tcp://0.0.0.0:26657'; done