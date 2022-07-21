package main

func main() {
	chain := NewBlockChain()
	defer chain.db.Close()

	cli := CLI{chain}
	cli.Run()
}
