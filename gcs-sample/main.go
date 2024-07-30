package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	// Define the bucket and object names
	// bucket := "your-bucket-name"
	// object := "your-object-name"
	bucket := "example-jh-test"
	// object := "addresses.csv"
	// object := "sample-1.tsv"
	object := "sample-5.tsv"

	// Use a buffer to capture the output
	var w bytes.Buffer

	// Call the downloadFileIntoMemory function
	data, err := downloadFileIntoMemory(&w, bucket, object)
	if err != nil {
		fmt.Printf("Failed to download file: %v\n", err)
		return
	}

	// Print the downloaded data
	fmt.Println("Downloaded data:", string(data))
	fmt.Println(w.String())
}

// downloadFileIntoMemory downloads an object.
func downloadFileIntoMemory(w io.Writer, bucket, object string) ([]byte, error) {
	// bucket := "bucket-name"
	// object := "object-name"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %w", object, err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}
	fmt.Fprintf(w, "Blob %v downloaded.\n", object)
	return data, nil
}
