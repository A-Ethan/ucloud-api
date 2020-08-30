/**
Homepage: https://github.com/ucloud/ucloud-sdk-go
Examples: https://github.com/ucloud/ucloud-sdk-go/tree/master/examples
*/

package main

import (
	"encoding/base64"
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func main() {
	cfg := ucloud.NewConfig()
	cfg.Region = "cn-bj2"
	cfg.BaseUrl = "https://api.ucloud.cn"

	cred := auth.NewCredential()
	cred.PublicKey = "helloPublicKey"
	cred.PrivateKey = "helloPrivateKey"

	pod := `apiVersion: v1beta1
kind: Pod
spec:
  containers:
  - name: cube01
    image: 'uhub.service.ucloud.cn/hello123/nginx:1.17.10-alpine'
    resources:
      limits:
        memory: 1024Mi
        cpu: 1000m
    volumeMounts: []
  volumes: []
  restartPolicy: Always`
	podBase := base64.StdEncoding.EncodeToString([]byte(pod))

	// fmt.Println("[Pod]", podBase)

	// New Client
	Client := ucloud.NewClient(&cfg, &cred)

	req := Client.NewGenericRequest()
	err := req.SetPayload(map[string]interface{}{
		"Action":    "CreateCubePod",
		"Region":    "cn-bj2",
		"Zone":      "cn-bj2-02",
		"ProjectId": "org-znwjjr",
		"VPCId":     "uvnet-q1anvtk2",
		"SubnetId":  "subnet-r0oxidkt",
		"Pod":       podBase,
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
