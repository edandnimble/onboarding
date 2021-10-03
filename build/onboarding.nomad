job "onboarding" {
  datacenters = ["dc1"]

  group "consul" {
    count = 1
    task "consul" {
      driver = "raw_exec"

      config {
        command = "consul"
        args = ["agent", "-dev"]
      }

      artifact {
        source = "https://releases.hashicorp.com/consul/1.10.3/consul_1.10.3_linux_amd64.zip"
      }
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
      port "http" {}
    }

    task "api" {
      driver = "docker"

      env {
        NUMBER_GRPC_PORT = "${NOMAD_PORT_number_grpc}"
        GUESSER_GRPC_PORT = "${NOMAD_PORT_guesser_grpc}"
        API_GRPC_PORT = "${NOMAD_PORT_grpc}"
        API_HTTP_PORT = "${NOMAD_PORT_http}"
      }

      config {
        image = "api:1.0"
        ports = ["http", "grpc"]
      }

      service {
        name = "api"
        port = "http"
        tags = ["http"]
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
        API_GRPC_PORT = "${NOMAD_PORT_api_grpc}"
      }

      config {
        image = "guesser:1.0"
        ports = ["grpc"]
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

      env {
        NUMBER_PORT = "${NOMAD_PORT_grpc}"
      }

      config {
        image = "tasks:1.0"
        ports = ["grpc"]
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
      port "redis" {}
    }

    task "redis" {
      driver = "docker"

      env {
        REDIS_PORT = "${NOMAD_PORT_redis}"
      }

      config {
        image = "redis:6.2"
        ports = ["redis"]
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
      port "mongo" {}
    }

    task "mongo" {
      driver = "docker"

      env {
        MONGO_PORT = "${NOMAD_PORT_mongo}"
      }

      config {
        image = "mongo:5.0.3"
        ports = ["mongo"]
      }

      service {
        name = "mongo"
        port = "mongo"
      }
    }
  }
}

