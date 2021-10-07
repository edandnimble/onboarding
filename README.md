# Onboarding

**Task description:**
https://proxit.atlassian.net/wiki/spaces/DEV/pages/2558919270/Onboarding+-+First+Issue+Task

## Build & Run

**build docker images**

Note a version is needed so Nomad won't try to pull latest image from docker hub

<pre><code>cd src/onboarding.com
docker build --tag api:1.0  -f api/Dockerfile .
docker build --tag number:1.0  -f number/Dockerfile .
docker build --tag tasks:1.0  -f tasks/Dockerfile .
docker build --tag guesser:1.0  -f guesser/Dockerfile .</pre></code>

**run Nomad & Consul**

install Nomad & Consul on your PC
https://learn.hashicorp.com/tutorials/nomad/get-started-install
https://learn.hashicorp.com/tutorials/consul/get-started-install

Add permissions to Consul for binding port 53 (on linux any port lower than 1024 needs binding permission)
<pre><code>sudo setcap CAP_NET_BIND_SERVICE=+eip /usr/bin/consul</pre></code>

Get your public IP (ifconfig/ip address) and change the nomad file build/onboarding.nomad so that Consul will and DNS servers passed to container will list/bind the public IP

Run Nomad agent (server & client)
<pre><code>sudo nomad agent -dev-connect -consul-address="{PUBLIC_IP}:8500"</pre></code>

* -dev-connect will configure nomad server & client with default configuration for development on local pc but unlike -dev it will bind the agent to the public IP
* -consul-address notify Nomad where the local Consul agent is listening
all options are described here:
https://www.nomadproject.io/docs/commands/agent#consul-address

Run Containers
<pre><code>nomad job run build/onboarding.nomad</pre></code>
Stop Containers
<pre><code>nomad stop onboarding</pre></code>

**Note:** to debug from inside container we need a bash shell, change docker file of the container to run with some image like ubuntu instead of alpine and add external DNS like 8.8.8.8 to the nomad file dns_servers config.

In order to run localy without containers run:
<pre><code>cd src/onboarding.com/run
go run ./main.go</pre></code>

