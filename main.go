package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	rpcPtr := flag.String("rpc", "https://testnet-rpc.monad.xyz", "Monad RPC endpoint URL")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	client, err := ethclient.Dial(*rpcPtr)
	if err != nil {
		fmt.Printf("Failed to connect to RPC: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	switch args[0] {
	case "tps":
		fs := flag.NewFlagSet("tps", flag.ExitOnError)
		blocks := fs.Int("blocks", 10, "Number of blocks to analyze")
		fs.Parse(args[1:])
		checkTPS(client, *blocks)

	case "blocktime":
		fs := flag.NewFlagSet("blocktime", flag.ExitOnError)
		blocks := fs.Int("blocks", 10, "Number of blocks to analyze")
		fs.Parse(args[1:])
		checkBlockTime(client, *blocks)

	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Monad testnet-2 Benchmark Tool")
	fmt.Println("Usage: monadbench [global flags] <command> [command flags]")
	fmt.Println("\nGlobal flags:")
	flag.PrintDefaults()
	fmt.Println("\nCommands:")
	fmt.Println("  tps       - Check transactions per second")
	fmt.Println("  blocktime - Check average block time")
	fmt.Println("\nExamples:")
	fmt.Println("  monadbench -rpc https://testnet-rpc.monad.xyz tps -blocks 20")
	fmt.Println("  monadbench -rpc https://testnet-rpc.monad.xyz blocktime -blocks 20")
}

func checkTPS(client *ethclient.Client, blocks int) {
	if blocks < 5 {
		fmt.Println("ERROR: Minimum 5 blocks required")
		os.Exit(1)
	}

	if blocks > 200 {
		fmt.Println("ERROR: Maximum 200 blocks required")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	latest, err := client.BlockNumber(ctx)
	if err != nil {
		fmt.Printf("Failed to get latest block: %v\n", err)
		os.Exit(1)
	}

	startBlock := latest - uint64(blocks-1)
	if startBlock < 0 {
		startBlock = 0
	}

	fmt.Printf("Analyzing %d blocks from %d to %d...\n", blocks, startBlock, latest)

	var (
		totalTxs    int
		firstTime   uint64
		lastTime    uint64
		blocksFound int
	)

	for i := startBlock; i <= latest; i++ {
		block, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			fmt.Printf("Warning: Failed to get block %d: %v\n", i, err)
			continue
		}

		txs := len(block.Transactions())
		totalTxs += txs
		blocksFound++

		if i == startBlock {
			firstTime = block.Time()
		}
		if i == latest {
			lastTime = block.Time()
		}
	}

	if blocksFound < 5 {
		fmt.Println("ERROR: Insufficient blocks retrieved for analysis")
		os.Exit(1)
	}

	duration := float64(lastTime - firstTime)
	tps := float64(totalTxs) / duration

	fmt.Println("\nTPS Analysis Results:")
	fmt.Printf("Blocks analyzed : %d - %d\n", startBlock, latest)
	fmt.Printf("Total blocks    : %d\n", blocksFound)
	fmt.Printf("Total txs       : %d\n", totalTxs)
	fmt.Printf("Time span       : %.2f seconds\n", duration)
	fmt.Printf("TPS             : %.2f transactions/second\n", tps)
}

func checkBlockTime(client *ethclient.Client, blocks int) {
	if blocks < 5 {
		fmt.Println("ERROR: Minimum 5 blocks required")
		os.Exit(1)
	}

	if blocks > 200 {
		fmt.Println("ERROR: Maximum 200 blocks required")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	latest, err := client.BlockNumber(ctx)
	if err != nil {
		fmt.Printf("Failed to get latest block: %v\n", err)
		os.Exit(1)
	}

	startBlock := latest - uint64(blocks-1)
	if startBlock < 0 {
		startBlock = 0
	}

	fmt.Printf("Analyzing %d blocks from %d to %d...\n", blocks, startBlock, latest)

	var (
		prevTime     uint64
		totalDiff    float64
		measurements int
	)

	for i := startBlock; i <= latest; i++ {
		block, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			fmt.Printf("Warning: Failed to get block %d: %v\n", i, err)
			continue
		}

		currentTime := block.Time()
		if i > startBlock {
			diff := float64(currentTime-prevTime) * 1000
			totalDiff += diff
			measurements++
		}
		prevTime = currentTime
	}

	if measurements < 4 {
		fmt.Println("ERROR: Insufficient blocks retrieved for analysis")
		os.Exit(1)
	}

	avgBlockTime := totalDiff / float64(measurements)

	fmt.Println("\nBlock Time Analysis Results:")
	fmt.Printf("Blocks analyzed : %d - %d\n", startBlock, latest)
	fmt.Printf("Total blocks    : %d\n", measurements+1)
	fmt.Printf("Intervals       : %d\n", measurements)
	fmt.Printf("Total time      : %.2f ms\n", totalDiff)
	fmt.Printf("Avg block time  : %.2f ms\n", avgBlockTime)
}
