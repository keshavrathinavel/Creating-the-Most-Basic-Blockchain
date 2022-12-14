package main
import (
	"crypto/sha256"
	"fmt"
	"time"
)

// block struct with block features

type Block struct{
	timestamp time.Time
	transactions []string
	prevHash []byte
	Hash []byte
}

// constructor that returns pointer to create new blocks

func NewBlock(transactions []string, prevHash []byte) *Block{

	currentTime := time.Now()
	return &Block {
		timestamp : currentTime,
		transactions : transactions,
		prevHash : prevHash,
		Hash : NewHash( currentTime, transactions, prevHash),
		}
}

// create hashing function

func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {

	input := append(prevHash, time.String()...)
	for transactions := range transactions {
		input = append(input, string(rune(transactions))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

// printing the contents of every block 

func printBlockInformation(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions: ")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}

// creating transactions, hashes and finally, blocks to implement the above
// first block is GenesisBlock

func main() {
	genesisTransactions := []string{"Keshav needs to find himself", "Alone."}
	genesisBlock := NewBlock(genesisTransactions, []byte{})
	fmt.Println("--- First Block ---")
	printBlockInformation(genesisBlock)

	block2Transactions := []string{"Keshav has it now."}
	block2 := NewBlock(block2Transactions, genesisBlock.Hash)
	fmt.Println("--- Second Block ---")
	printBlockInformation(block2)

	block3Transactions := []string{"Why does he feel empty", "Because he is."}
	block3 := NewBlock(block3Transactions, block2.Hash)
	fmt.Println("--- Third Block ---")
	printBlockInformation(block3)
}