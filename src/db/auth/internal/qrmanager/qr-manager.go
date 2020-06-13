package qrmanager

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func GenerateQr(content string) ([]byte, error) {
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		fmt.Println(err)
		return png, err
	}
	return png, nil
}
