package updater

import (
	"encoding/hex"
	"hash"
	"io"
	"os"
)

func getFileHash(filename string, hasher hash.Hash) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(hasher, file)
	if err != nil {
		return "", err
	}
	res := hex.EncodeToString(hasher.Sum(nil))

	return res, nil
}
