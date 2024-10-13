package main

import (
	"context"
	cctprelayer "dydx-cctp-relayer/cctp_relayer"

	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	godotenv.Load()
}

func main() {
	hexPrivateKey := os.Getenv("PRIVATE_KEY")
	nobleMintRecipient := os.Getenv("NOBLE_MINT_RECIPIENT")

	ctx := context.Background()

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	cctpClient, err := cctprelayer.NewCCTPClient(
		ctx,
		log,
		hexPrivateKey,
		cctprelayer.Arbitrum,
		nobleMintRecipient,
		11, // in USDC
	)
	if err != nil {
		log.Fatalf("failed to create CCTP client: %s", err)
	}

	err = cctpClient.RequestCCTPTransfer()
	if err != nil {
		log.Fatalf("failed to request CCTP transfer: %s", err)
	}
}
