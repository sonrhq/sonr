package transfer

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

// ^ Helper: Chunks string based on B64ChunkSize ^ //
func ChunkBase64(s string) []string {
	chunkSize := B64ChunkSize
	ss := make([]string, 0, len(s)/chunkSize+1)
	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:chunkSize]), s[chunkSize:]
	}
	return ss
}

// ^ Helper: Encodes to Jpeg Image ^ //
func EncodeJpegBuffer(buf *bytes.Buffer, path string) error {
	// Open File at Meta Path
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Encode as Jpeg into buffer
	err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 100})
	if err != nil {
		return err
	}
	return nil
}

// ^ Helper: Encodes to PNG Image ^ //
func EncodePngBuffer(buf *bytes.Buffer, path string) error {
	// Open File at Meta Path
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert to Image Object
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Encode as Jpeg into buffer
	err = png.Encode(buf, img)
	if err != nil {
		return err
	}
	return nil
}

// ^ Helper: Opens File for Buffer and Writes ^ //
func WriteBuffer(buf *bytes.Buffer, path string) error {
	// Open File at Meta Path\
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Write Bytes to buffer
	_, err = buf.Write(dat)
	if err != nil {
		return err
	}
	return nil
}
