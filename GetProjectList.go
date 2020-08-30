/**
Homepage: https://github.com/ucloud/ucloud-sdk-go
Examples: https://github.com/ucloud/ucloud-sdk-go/tree/master/examples
*/

package main

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go"
	"github.com/ucloud/ucloud-sdk-go/services/uaccount"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func main() {
	cfg := ucloud.NewConfig()
	cfg.Region = "cn-bj2"
	cfg.BaseUrl = "https://api.ucloud.cn"

	cred := auth.NewCredential()
	cred.PublicKey = "helloPublicKey"
	cred.PrivateKey = "helloPrivateKey"

	uaccountClient := uaccount.NewClient(&cfg, &cred)

	req := uaccountClient.NewGetProjectListRequest()
	req.IsFinance = ucloud.String("No")

	resp, err := uaccountClient.GetProjectList(req)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println("[RESPONSE]", resp)
}
