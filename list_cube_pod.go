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
		"Action":    "ListCubePod",
		"Region":    "cn-bj2",
		"Zone":      "cn-bj2-02",
		"ProjectId": "org-znwjjr",
		"VPCId":     "uvnet-q1anvtk2",
		"SubnetId":  "subnet-r0oxidkt",
		"Limit":     "20",
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


