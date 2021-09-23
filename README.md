APM template

# commonly used command
* `go run . ` or `go run main.go` to run the app 
* `go build <file_you_want_to_build>.go` to build the file, then `./<file_you_built>` to tun the go program directly

here is a list of other commands
https://www.ubuntupit.com/go-command-examples-for-aspiring-golang-developers/

# Deployment to Azure
## pre-requisite: 
* have docker installed
* created an account on Docker Hub
* have access to `MFC-Global-Canada-PRODLNENG-AFF_HD-NonProduction-S1` subscription so you have access to this `cdnii-sandbox-cac-hd-rg` resource group to deploy the app

https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker

steps to push to docker hub:
https://docs.docker.com/docker-hub/repos/#:~:text=To%20push%20an%20image%20to,docs%2Fbase%3Atesting%20).


`docker build -t <your_docker_username>/trivia_gophers -f Dockerfile.production .`
`docker push <your_docker_username>/trivia_gophers`
