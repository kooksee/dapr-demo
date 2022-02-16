appPort=50051
serverName=server0
dapr run --app-id ${serverName} --app-port=${appPort} --config config.yaml -- go run server/main.go ${appPort} &

appPort=50052
serverName=server0
dapr run --app-id ${serverName} --app-port=${appPort} --config config.yaml -- go run server/main.go ${appPort} &

appPort=50053
serverName=server1
dapr run --app-id ${serverName} --app-port=${appPort} --config config.yaml -- go run server/main.go ${appPort} &

appPort=50054
serverName=server1
dapr run --app-id ${serverName} --app-port=${appPort} --config config.yaml -- go run server/main.go ${appPort} &