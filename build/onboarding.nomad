job "onboarding" {
  datacenters = ["dc1"]

  group "consul" {
    count = 1
    task "consul" {
      driver = "raw_exec"

      config {
        command = "consul"
        args = ["agent", "-dev", "-bind", "10.254.67.12", "-advertise", "10.254.67.12", "-client", "10.254.67.12", "-dns-port", "53"]
      }

      # cannot use artifact with DNS because of bind 53 port restictions on linux
      # install consul and run:
      # $ sudo setcap CAP_NET_BIND_SERVICE=+eip /usr/bin/consul
      # artifact {
      #   source = "https://releases.hashicorp.com/consul/1.10.3/consul_1.10.3_linux_amd64.zip"
      # }
    }
  }

  group "number" {
    count = 1

    network {
      port "grpc" {}
    }

    task "number" {
      driver = "docker"

      env {
        NUMBER_GRPC_PORT = "${NOMAD_PORT_grpc}"
      }

      config {
        image = "number:1.0"
        ports = ["grpc"]
        dns_servers = ["10.254.67.12"]
      }

      service {
        name = "number"
        port = "grpc"
      }
    }
  }

  group "api" {
    count = 1

    network {
      port "grpc" {}
      port "http" {
        static = 8080
      }
    }

    task "api" {
      driver = "docker"

      env {
        API_GRPC_PORT = "${NOMAD_PORT_grpc}"
        API_HTTP_PORT = "${NOMAD_PORT_http}"
      }

      config {
        image = "api:1.0"
        ports = ["http", "grpc"]
        dns_servers = ["10.254.67.12"]
        # dns_servers = ["10.254.67.12", "8.8.8.8"] debug with internet access
      }

      service {
        name = "api"
        port = "http"
        tags = ["http"]
      }
      service {
        # dig grpc.api.service.consul works but cannot get DNS with tag from golang (
        # name = "api"
        name = "grpcapi"
        port = "grpc"
        tags = ["grpc"]
      }
    }
  }

  group "guesser" {
    count = 1

    network {
      port "grpc" {}
    }

    task "guesser" {
      driver = "docker"

      env {
        GUESSER_GRPC_PORT = "${NOMAD_PORT_grpc}"
      }

      config {
        image = "guesser:1.0"
        ports = ["grpc"]
        dns_servers = ["10.254.67.12"]
      }

      service {
        name = "guesser"
        port = "grpc"
      }
    }
  }

  group "tasks" {
    count = 1

    network {
      port "grpc" {}
    }

    task "tasks" {
      driver = "docker"

      config {
        image = "tasks:1.0"
        ports = ["grpc"]
        dns_servers = ["10.254.67.12"]
      }

      service {
        name = "tasks"
        port = "grpc"
      }
    }
  }

  group "redis" {
    count = 1

    network {
        port "redis" {
            to = 6379
        }
    }

    task "redis" {
      driver = "docker"

      env {
        REDIS_PORT = "${NOMAD_PORT_redis}"
      }

      config {
        image = "redis:6.2"
        ports = ["redis"]
        dns_servers = ["10.254.67.12"]
      }

      service {
        name = "redis"
        port = "redis"
      }
    }
  }

  group "mongo" {
    count = 1

    network {
        port "mongo" {
            to = 27017
        }
    }

    task "mongo" {
      driver = "docker"

      config {
        image = "mongo:5.0.3"
        ports = ["mongo"]
        dns_servers = ["10.254.67.12"]
      }

      service {
        name = "mongo"
        port = "mongo"
      }
    }
  }
}

