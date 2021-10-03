module onboarding.com/guesser

go 1.17

replace onboarding.com/api => ./../api/

replace onboarding.com/number => ./../number/

replace onboarding.com/utils => ./../utils/

replace onboarding.com/tasks => ./../tasks/

replace onboarding.com/guesser => ./

require (
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	onboarding.com/api v0.0.0-00010101000000-000000000000
	onboarding.com/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/golang/snappy v0.0.2 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.7.2 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
