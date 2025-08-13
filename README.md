# 📟 Monad Benchmark Tools (Linux)
Monad CLI Tools for Benchmark & Analysis. A simple command-line tool written in Go to check TPS (Transactions Per Second) and average block time on the Monad testnet-2 network.

## 🚀 Features
### Check TPS (Transactions Per Second)
- Calculates transactions per second within a specified block range
- Displays total transactions, time span, and TPS

### Check Block Time:
- Calculates average block creation time within a specified block range
- Displays total time, number of intervals, and average block time.

## 🛠️ Installation
- Clone this repository
- Go to binary directory

### Commands:
```bash
  git clone https://github.com/rawakinode/benchmark-monad.git
  cd benchmark-monad/bin
  sudo chmod+x monadbench
```


## 📦 Usage
### Global Flags

| Flag | Description | Default |
|---------|---------|---------|
| -rpc | RPC endpoint URL	| https://testnet-rpc.monad.xyz |

### Commands


Check TPS
```bash
./monadbench [-rpc RPC_URL] tps [-blocks BLOCK_COUNT]
```
Check Block Time
```bash
./monadbench [-rpc RPC_URL] blocktime [-blocks BLOCK_COUNT]
```

Example:
```bash
./monadbench -rpc https://testnet-rpc.monad.xyz tps -blocks 20
./monadbench -rpc https://testnet-rpc.monad.xyz blocktime -blocks 20
```

Result:
```bash
Analyzing 20 blocks from 4123400 to 4123419...

TPS Analysis Results:
Blocks analyzed : 4123400 - 4123419
Total blocks    : 20
Total txs       : 184
Time span       : 240000.00 ms
TPS             : 0.77 transactions/second

Analyzing 15 blocks from 4123410 to 4123424...

Block Time Analysis Results:
Blocks analyzed : 4123410 - 4123424
Total blocks    : 15
Intervals       : 14
Total time      : 168000.00 ms
Avg block time  : 12000.00 ms
```

## 🤝 Contributing
Got an idea or found a bug? Jump in! Fork the repo, make your changes, and send us a PR. We love seeing what the community builds.

## 🙏 Credits
Created and maintained by **rawakinode**.  
Thanks to all contributors and Monad Community.  
