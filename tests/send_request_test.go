package tests

import (
	"fmt"
	"testing"

	"github.com/meloshub/meloshub/network"
)

func TestSendRequest(t *testing.T) {
	url := "https://baidu.com"
	resp, err := network.Get(url, nil)
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	if err := resp.IsSuccess(); err != nil {
		t.Fatalf("Request was not successful: %v", err)
	}
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Text: %s\n", resp.Text())
	fmt.Printf("Response Headers: %v\n", resp.Headers)
}
