/**
Homepage: https://github.com/ucloud/ucloud-sdk-go
Examples: https://github.com/ucloud/ucloud-sdk-go/tree/master/examples
*/

package main

import (
	"fmt"
	"os"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func main() {
	cfg := ucloud.NewConfig()
	cfg.Region = "cn-bj2"
	cfg.BaseUrl = "https://api.ucloud.cn"

	cred := auth.NewCredential()
	cred.PublicKey = os.Getenv("UCLOUD_API_PUBLICKEY")
	cred.PrivateKey = os.Getenv("UCLOUD_API_PRIVATEKEY")

	// New Client
	Client := ucloud.NewClient(&cfg, &cred)

	req := Client.NewGenericRequest()
	err := req.SetPayload(map[string]interface{}{
		"Action":    "GetCubePod",
		"Region":    "cn-bj2",
		"Zone":      "cn-bj2-02",
		"ProjectId": "org-znwjjr",
		"CubeId":    "cube-4o3w3kzw",
	})

	if err != nil {
		panic(err)
	}

	// send request
	resp, err := Client.GenericInvoke(req)

	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[RESPONSE]", resp)
}

