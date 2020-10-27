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
	// 替换nfs挂载点

	pod := `apiVersion: v1beta1
	kind: Pod
	metadata:
	spec:
	  containers:
		- image: 'uhub.service.ucloud.cn/hello123/nginx:1.17.10-alpine'
		  name: nginx
		  resources:
			limits:
			  cpu: '1'
			  memory: 1Gi
		  volumeMounts:
			- mountPath: /etc/nginx/conf.d/default.conf
			  name: defaultconf
			  subPath: default.conf
			- mountPath: /usr/share/nginx/html/
			  name: wpnfs
		- args:
			- '-R'
		  command:
			- php-fpm
		  image: 'uhub.service.ucloud.cn/hello123/wordpress:php7.4-fpm'
		  name: php
		  resources:
			limits:
			  cpu: '1'
			  memory: 1Gi
		  volumeMounts:
			- mountPath: /var/www/html/
			  name: wpnfs
	  initContainers:
		- args:
			- '-c'
			- >-
			  wget http://cube.cn-bj.ufileos.com/wordpress-5.4.2.tar && tar -zxvf
			  wordpress-5.4.2.tar && mv /wordpress /data/wordpress
		  command:
			- sh
		  image: 'uhub.service.ucloud.cn/hello123/busybox:1.28'
		  name: initcube01
		  volumeMounts:
			- mountPath: /data
			  name: wpnfs
	  restartPolicy: Always
	  volumes:
		- name: defaultconf
		  secret:
			default.conf: >-
			  c2VydmVyIHsKICAgIGxpc3RlbiAgICAgICA4MDsKICAgIHNlcnZlcl9uYW1lICBsb2NhbGhvc3Q7CgogICAgbG9jYXRpb24gLyB7CiAgICAgICAgcm9vdCAgIC91c3Ivc2hhcmUvbmdpbngvaHRtbC93b3JkcHJlc3M7CiAgICAgICAgaW5kZXggIGluZGV4Lmh0bWwgaW5kZXgucGhwOwogICAgfQogICAgZXJyb3JfcGFnZSAgIDUwMCA1MDIgNTAzIDUwNCAgLzUweC5odG1sOwogICAgbG9jYXRpb24gPSAvNTB4Lmh0bWwgewogICAgICAgIHJvb3QgICAvdXNyL3NoYXJlL25naW54L2h0bWw7CiAgICB9CiAgICBsb2NhdGlvbiB+IFwucGhwJCB7CiAgICAgICAgZmFzdGNnaV9wYXNzIGxvY2FsaG9zdDo5MDAwOwogICAgICAgIGZhc3RjZ2lfaW5kZXggaW5kZXgucGhwOwogICAgICAgIGZhc3RjZ2lfYnVmZmVycyAxNiAxNms7CiAgICAgICAgZmFzdGNnaV9idWZmZXJfc2l6ZSAzMms7CiAgICAgICAgZmFzdGNnaV9wYXJhbSBTQ1JJUFRfRklMRU5BTUUgL3Zhci93d3cvaHRtbC93b3JkcHJlc3MvJGZhc3RjZ2lfc2NyaXB0X25hbWU7CiAgICAgICAgI2ZpeGVzIHRpbWVvdXRzCiAgICAgICAgZmFzdGNnaV9yZWFkX3RpbWVvdXQgNjAwOwogICAgICAgIGluY2x1ZGUgZmFzdGNnaV9wYXJhbXM7CiAgICB9Cn0=
		- name: wpnfs
		  nfs:
			address: '10.9.17.238:/'
			mountOption: vers=4.0`
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
