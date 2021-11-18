package lfs

// src: https://github.com/git-lfs/git-lfs/blob/1a4431a2e9a8cbe5cf0949c527b4554d03169634/lfs/pointer.go

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

const (
	// blobSizeCutoff is used to determine which files to scan for Git LFS
	// pointers.  Any file with a size below this cutoff will be scanned.
	blobSizeCutoff = 1024
)

var (
	v1Aliases = []string{
		"http://git-media.io/v/2",            // alpha
		"https://hawser.github.com/spec/v1",  // pre-release
		"https://git-lfs.github.com/spec/v1", // public launch
	}
	latest      = "https://git-lfs.github.com/spec/v1"
	oidType     = "sha256"
	oidRE       = regexp.MustCompile(`\A[0-9a-f]{64}\z`)
	matcherRE   = regexp.MustCompile("git-media|hawser|git-lfs")
	extRE       = regexp.MustCompile(`\Aext-\d{1}-\w+`)
	pointerKeys = []string{"version", "oid", "size"}

	emptyObjectSHA256 = hex.EncodeToString(sha256.New().Sum(nil))
)

type Pointer struct {
	Version   string
	Oid       string
	Size      int64
	OidType   string
	Canonical bool
}

// A PointerExtension is parsed from the Git LFS Pointer file.
func NewPointer(oid string, size int64) *Pointer {
	return &Pointer{latest, oid, size, oidType, true}
}

func EmptyPointer() *Pointer {
	return NewPointer(emptyObjectSHA256, 0)
}

func (p *Pointer) Encoded() string {
	if p.Size == 0 {
		return ""
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("version %s\n", latest))
	buffer.WriteString(fmt.Sprintf("oid %s:%s\n", p.OidType, p.Oid))
	buffer.WriteString(fmt.Sprintf("size %d\n", p.Size))
	return buffer.String()
}

// DecodeFrom decodes an *lfs.Pointer from the given io.Reader, "reader".
// If the pointer encoded in the reader could successfully be read and decoded,
// it will be returned with a nil error.
//
// If the pointer could not be decoded, an io.Reader containing the entire
// blob's data will be returned, along with a parse error.
func DecodeFrom(meta *filesystem.Metadata, reader io.Reader) (*Pointer, io.Reader, error) {
	if meta.Size < blobSizeCutoff {
		return nil, reader, fmt.Errorf("file not big enough")
	}

	buf := make([]byte, blobSizeCutoff)
	n, err := reader.Read(buf)
	buf = buf[:n]

	var contents io.Reader = bytes.NewReader(buf)
	if err != io.EOF {
		contents = io.MultiReader(contents, reader)
	}

	if err != nil && err != io.EOF {
		return nil, contents, err
	}

	if len(buf) == 0 {
		return EmptyPointer(), contents, nil
	}

	p, err := decodeKV(bytes.TrimSpace(buf))
	if err == nil && p != nil {
		p.Canonical = p.Encoded() == string(buf)
	}
	return p, contents, err
}

func verifyVersion(version string) error {
	if len(version) == 0 {
		return errors.New("missing version")
	}

	for _, v := range v1Aliases {
		if v == version {
			return nil
		}
	}

	return errors.New("Invalid version: " + version)
}

func decodeKV(data []byte) (*Pointer, error) {
	kvps, err := decodeKVData(data)
	if err != nil {
		return nil, err
	}

	if err := verifyVersion(kvps["version"]); err != nil {
		return nil, err
	}

	value, ok := kvps["oid"]
	if !ok {
		return nil, errors.New("invalid Oid")
	}

	oid, err := parseOid(value)
	if err != nil {
		return nil, err
	}

	value = kvps["size"]
	size, err := strconv.ParseInt(value, 10, 64)
	if err != nil || size < 0 {
		return nil, fmt.Errorf("invalid size: %q", value)
	}

	return NewPointer(oid, size), nil
}

func parseOid(value string) (string, error) {
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return "", errors.New("Invalid Oid value: " + value)
	}
	if parts[0] != oidType {
		return "", errors.New("Invalid Oid type: " + parts[0])
	}
	oid := parts[1]
	if !oidRE.Match([]byte(oid)) {
		return "", errors.New("Invalid Oid: " + oid)
	}
	return oid, nil
}

func decodeKVData(data []byte) (kvps map[string]string, err error) {
	kvps = make(map[string]string)

	if !matcherRE.Match(data) {
		err = errors.New("invalid header")
		return
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	line := 0
	numKeys := len(pointerKeys)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		parts := strings.SplitN(text, " ", 2)
		if len(parts) < 2 {
			err = fmt.Errorf("error reading line %d: %s", line, text)
			return
		}

		key := parts[0]
		value := parts[1]

		if numKeys <= line {
			err = fmt.Errorf("extra line: %s", text)
			return
		}

		if expected := pointerKeys[line]; key != expected {
			if !extRE.Match([]byte(key)) {
				err = fmt.Errorf("expected %s, got %s", expected, key)
				return
			}
			continue
		}

		line += 1
		kvps[key] = value
	}

	err = scanner.Err()
	return
}
