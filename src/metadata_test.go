package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/oasisprotocol/oasis-core/go/runtime/bundle"
	"github.com/stretchr/testify/require"
)

const metadataDirectory = "../metadata"

func TestMetadata(t *testing.T) {
	client := http.Client{
		Timeout: time.Minute,
	}

	dir, err := os.Open(metadataDirectory)
	require.NoError(t, err)
	defer dir.Close()

	files, err := dir.Readdirnames(0)
	require.NoError(t, err)

	for _, file := range files {
		fmt.Printf("Verifying metadata: %s\n", file)

		path := filepath.Join(metadataDirectory, file)
		bytes, err := os.ReadFile(path)
		require.NoError(t, err)

		url := strings.TrimSpace(string(bytes))
		fmt.Printf("Downloading bundle: %s\n", url)

		resp, err := client.Get(url)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode)

		tempFile, err := os.CreateTemp("", "bundle.orc")
		require.NoError(t, err)
		defer tempFile.Close()

		_, err = io.Copy(tempFile, resp.Body)
		require.NoError(t, err)

		bnd, err := bundle.Open(tempFile.Name())
		require.NoError(t, err)

		hash := bnd.Manifest.Hash().String()
		fmt.Printf("Manifest hash: %s\n", hash)
		require.Equal(t, hash, file)

		fmt.Println()
	}
}
