/**
Homepage: https://github.com/ucloud/ucloud-sdk-go
Examples: https://github.com/ucloud/ucloud-sdk-go/tree/master/examples
*/

package main

import (
	"encoding/base64"
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

	pod := `apiVersion: v1beta1
kind: Pod
spec:
  containers:
  - name: cube01
    image: 'uhub.service.ucloud.cn/hello321/python:3.7.7-alpine'
    resources:
      limits:
        memory: 1024Mi
        cpu: 1000m
    volumeMounts: []
  volumes: []
  imagePullSecrets:
  - password: xxxx
    registryserver: uhub.service.ucloud.cn
    username: ethan.shen@ucloud.cn
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
