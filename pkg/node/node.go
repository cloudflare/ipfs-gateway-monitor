package node

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"time"

	api "github.com/ipfs/go-ipfs-api"
	files "github.com/ipfs/go-ipfs-files"
)

var (
	ErrKeyNotFound = errors.New("The key name is not found.")
)

var timeout = 15 * time.Minute

type Node string

func (node *Node) NewShell() *NodeShell {
	shell := api.NewShell(string(*node))
	shell.SetTimeout(timeout)
	return &NodeShell{
		node,
		shell,
	}
}

type NodeShell struct {
	node  *Node
	shell *api.Shell
}

// Cat the content at the given path.
func (nsh *NodeShell) Cat(path string) ([]byte, error) {
	sh := nsh.shell
	rc, err := sh.Cat(path)
	if err != nil {
		return nil, err
	}

	output, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	if err := rc.Close(); err != nil {
		return nil, err
	}
	return output, nil
}

// Find the key name in the key store.
func (nsh *NodeShell) KeyFind(name string) (*api.Key, error) {
	sh := nsh.shell
	keys, err := sh.KeyList(context.Background())
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		if key.Name == name {
			return key, nil
		}
	}
	return nil, ErrKeyNotFound
}

// Remove the key name in the key store.
func (nsh *NodeShell) KeyRm(name string) (*api.Key, error) {
	sh := nsh.shell
	keys, err := sh.KeyRm(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return keys[0], nil
}

// Publish a mutable name to point to a given value.
func (nsh *NodeShell) Publish(name, cid string) error {
	sh := nsh.shell
	_, err := sh.PublishWithDetails(cid, name, 0, 0, true)
	if err != nil {
		return err
	}
	return nil
}

// Import the key no matter whether there is already the key name or not.
func (nsh *NodeShell) KeyReimport(name, privateKey string) error {
	// Firstly, check if there is a key name in the key store already or not.
	key, err := nsh.KeyFind(name)
	if err != nil && err != ErrKeyNotFound {
		return err
	}
	if key != nil {
		// Remove the key
		_, err := nsh.KeyRm(name)
		if err != nil {
			return err
		}
	}

	sh := nsh.shell
	// Add the key. Since there is no key import method in go-ipfs-api, we
	// need to create a HTTP request to import a key to the node from scratch.
	decodedKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return err
	}
	// We use go-ipfs-files to send a key file in multipart/form-data format
	// similar to what it's done in https://github.com/ipfs/go-ipfs-api/blob/master/shell.go#L508-L515
	fr := files.NewReaderFile(bytes.NewReader(decodedKey))
	slf := files.NewSliceDirectory([]files.DirEntry{files.FileEntry("", fr)})
	fileReader := files.NewMultiFileReader(slf, true)

	resp, err := sh.Request("key/import", name).
		Body(fileReader).
		Send(context.Background())
	if err != nil {
		return err
	}
	defer resp.Close()
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

// Generate the key no matter whether there is already the key name or not.
func (nsh *NodeShell) KeyRegen(name string) (*api.Key, error) {
	// Firstly, check if there is a key name in the key store already or not.
	key, err := nsh.KeyFind(name)
	if err != nil && err != ErrKeyNotFound {
		return nil, err
	}
	if key != nil {
		// Remove the key
		_, err := nsh.KeyRm(name)
		if err != nil {
			return nil, err
		}
	}

	sh := nsh.shell
	key, err = sh.KeyGen(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return key, nil
}
