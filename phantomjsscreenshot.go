package gnomescreenshot

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"os"
	"os/exec"
)

var dir = "/tmp/"

// EncodingField tells the browser what we've encoded and how to decode it
const EncodingField = "data:image/png;base64,"

// getUUID returns a UUID string to be used for naming
func getUUID() string {
	uid, _ := uuid.NewRandom()
	return uid.String()
}

// GrabScreenshot returns a link to a screenshot for retrieval on localhost
// and and errors encountered along the way with the calling application
// returns fileloc, command output, error
func GrabScreenshot(link string) (string, error) {

	//get new filename
	filename := dir + getUUID() + ".png"

	command := "gnome-web-photo"
	thumbnail := "--mode=thumbnail"
	size := "-s"
	sizeVal := "256"
	hyperlink := link

	args := []string{thumbnail, size, sizeVal, hyperlink, filename}
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", errors.Wrap(err, "command run")
	}

	b64, err := b64png(filename)

	err = os.Remove(filename)
	if err != nil {
		return b64, errors.Wrap(err, "error removing file, b64 may not be nil")
	}

	return b64, nil
}

func b64png(filename string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// create a new buffer base on file size
	fi, _ := f.Stat()

	size := fi.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(f)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	b64 := base64.StdEncoding.EncodeToString(buf)

	// add data encoding information for html
	b64 = EncodingField + b64

	// Embed into an html without PNG file
	return b64, nil
}
