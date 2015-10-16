# ngrok - Introspected tunnels to localhost ([homepage](https://ngrok.com))

## 基于官方ngrok修改

* 修改依赖包问题
* 修改编译,去掉静态资源文件编译到二进制文件
* 修改编译产出物为静态二进制文件,基于Dockerfile打包镜像更小
* 新增支持跨平台编译
* 修改ngrok客户端连接ngrokd服务端失败的bug

注意:    
因为bug,去掉了源代码中的验证功能,后续将完善.

## 编译

GOOS=linux GOARCH=amd64 ./build.sh

## 部署
部署依赖assets中的资源文件

部署服务端
cd ./bin && ./ngrokd -domain="example.com"

客户端
cd ./bin && ./ngrok -config=client.yaml 30001





### "I want to securely expose a web server to the internet and capture all traffic for detailed inspection and replay"
![](https://ngrok.com/static/img/overview.png)

## What is ngrok?
ngrok is a reverse proxy that creates a secure tunnel from a public endpoint to a locally running web service.
ngrok captures and analyzes all traffic over the tunnel for later inspection and replay.

## ngrok 2.0
**NOTE** This repository contains the code for ngrok 1.0. The code for ngrok 2.0 is not yet open source.

## What can I do with ngrok?
- Expose any http service behind a NAT or firewall to the internet on a subdomain of ngrok.com
- Expose any tcp service behind a NAT or firewall to the internet on a random port of ngrok.com
- Inspect all http requests/responses that are transmitted over the tunnel
- Replay any request that was transmitted over the tunnel


## What is ngrok useful for?
- Temporarily sharing a website that is only running on your development machine
- Demoing an app at a hackathon without deploying
- Developing any services which consume webhooks (HTTP callbacks) by allowing you to replay those requests
- Debugging and understanding any web service by inspecting the HTTP traffic
- Running networked services on machines that are firewalled off from the internet


## Downloading and installing ngrok
ngrok has _no_ runtime dependencies. Just download a single binary for your platform and run it. Some premium features
are only available by creating an account on ngrok.com. If you need them, [create an account on ngrok.com](https://ngrok.com/signup).

#### [Download ngrok for your platform](https://ngrok.com/download)

## Developing on ngrok
[ngrok developer's guide](docs/DEVELOPMENT.md)