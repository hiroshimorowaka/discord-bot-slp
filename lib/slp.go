package lib

import (
	"fmt"
	"os"

	"github.com/sch8ill/mclib"
	"github.com/sch8ill/mclib/slp"
)

func GetServerInfo() (*slp.Response, error) {
	host := os.Getenv("HOST")

	if host == "" {
		return nil, fmt.Errorf("server ip not set")
	}

	client, _ := mclib.NewClient(host)
	return client.StatusPing()

}
