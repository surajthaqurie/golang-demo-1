package main

import (
	"context"
	"fmt"
	"log"

	immudb "github.com/codenotary/immudb/pkg/client"
	"google.golang.org/grpc/metadata"
)

func Run() error {
	// Step 1: Configure immudb client options with server address and port
	opts := immudb.DefaultOptions().
		WithAddress("localhost"). // Server address
		WithPort(3322)            // Server port

	// Step 2: Create a new immudb client with the specified options
	client := immudb.NewClient().WithOptions(opts)

	// Step 3: Open a session with the immudb server using provided credentials
	err := client.OpenSession(
		context.TODO(),   // Context for session
		[]byte(`immudb`), // Username
		[]byte(`immudb`), // Password
		"defaultdb",      // Database name
	)

	// Error handling for failed session opening
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Step 4: Create a background context for operations
	ctx := context.Background()

	// Step 5: Ensure the client session is closed at the end
	defer client.CloseSession(ctx)

	// Step 6: Perform a VerifiedSet operation to store data securely with proof
	verifyTransaction, err := client.VerifiedSet(ctx, []byte(`welcome`), []byte(`suraj`))
	if err != nil {
		// Error handling for failed VerifiedSet operation
		log.Fatal(err)
		return err
	}

	// Step 7: Print the transaction ID of the VerifiedSet operation
	fmt.Printf("Successfully set key with transaction ID: %d\n", verifyTransaction.Id)

	// Step 8: Perform a VerifiedGet operation to retrieve the stored data securely
	verifyEntry, err := client.VerifiedGet(ctx, []byte(`welcome`))

	// Error handling for failed VerifiedGet operation
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Step 9: Print the retrieved and verified entry
	fmt.Printf("Successfully retrieved and verified entry %v\n", verifyEntry)

	// Step 10: Return nil to indicate the function completed successfully
	return nil
}

func main() {
	fmt.Println("ImmueDb")

	if err := Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func RunOldConnection() error {
	client, err := immudb.NewImmuClient(immudb.DefaultOptions())
	if err != nil {
		log.Fatal(err)
		return err

	}

	ctx := context.Background()

	lr, err := client.Login(ctx, []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Fatal(err)
		return err
	}

	md := metadata.Pairs("authorization", lr.Token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	vtx, err := client.VerifiedSet(ctx, []byte(`welcome`), []byte(`gophers`))
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Successfully set key with transaction ID: %d\n", vtx.Id)

	ventry, err := client.VerifiedGet(ctx, []byte(`welcome`))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Successfully retrieved and verified entry %v\n", ventry)

	return nil

}
