package converter

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
	"os"
	"text-converter/internal/cfg"
)

type Converter struct {

}

func (converter *Converter) ConvertPdf2Text(binaryContent []byte) (string, error) {
	config := cfg.GetConfig()
	// Create temporary file to read from
	tmpfilePath := fmt.Sprintf("/tmp/text-converter-%s.pdf", config.RequestId)
	tmpfile, err := os.Create(tmpfilePath)
	if err != nil {
		return "", err
	}
	tmpfile.Write(binaryContent)
	tmpfile.Close()
	defer os.Remove(tmpfilePath)

	// Read and convert temporary file
	f, r, err := pdf.Open(tmpfilePath)
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	buf.ReadFrom(b)
	text := buf.String()
	return text, nil
}