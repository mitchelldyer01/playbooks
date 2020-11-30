package applications

import {
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
}

type ImageConfig struct {
	tag string
	hostIP string
	hostPort string
	containerPort string
	protocol string
}

func Deploy(config ImageConfig) (string, error) {
	client, err = client.NewEnvClient()
	if err != nil {
		fmt.Println("Failed creating a Docker client")
		panic(err)
	}

	hostBinding := nat.PortBinding {
		HostIP: config.hostIP,
		HostPort: config.hostPort,
	}

	containerPort, err := nat.NewPort(config.protocol, config.containerPort)
	if err != nil {
		fmt.Println("Failed grabbing the port")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	
	container, err := client.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: config.tag,
		},
		&container.HostConfig{
			PortBindings: portBinding,
		}, 
		nil, 
		""
	)
	if err != nil {
		panic(err)
	}

	cli.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	fmt.Printf("Container % is started", container.ID)
	return container.ID, nil
}