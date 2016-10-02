package container

import (
	"errors"
	"os"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
	"github.com/docker/engine-api/types/strslice"
	"github.com/docker/go-connections/nat"
	"github.com/maliceio/malice/config"
	"github.com/maliceio/malice/malice/docker/client"
	er "github.com/maliceio/malice/malice/errors"
)

// Start starts a malice docker container
func Start(
	docker *client.Docker,
	cmd strslice.StrSlice,
	name string,
	image string,
	logs bool,
	binds []string,
	portBindings nat.PortMap,
	links []string,
	env []string,
) (types.ContainerJSONBase, error) {

	if docker.Ping() {
		// Check that all requirements for the container to run are ready
		checkContainerRequirements(docker, name, image)

		createContConf := &container.Config{
			Image: image,
			Cmd:   cmd,
			Env:   env,
		}
		hostConfig := &container.HostConfig{
			Binds:        binds,
			NetworkMode:  "malice",
			PortBindings: portBindings,
			Links:        links,
			Privileged:   false,
		}
		networkingConfig := &network.NetworkingConfig{}

		contResponse, err := docker.Client.ContainerCreate(context.Background(), createContConf, hostConfig, networkingConfig, name)
		if err != nil {
			log.WithFields(log.Fields{"env": config.Conf.Environment.Run}).Errorf("CreateContainer error = %s\n", err)
		}

		err = docker.Client.ContainerStart(context.Background(), contResponse.ID, types.ContainerStartOptions{})
		if err != nil {
			log.WithFields(log.Fields{"env": config.Conf.Environment.Run}).Errorf("StartContainer error = %s\n", err)
		}

		if logs {
			LogContainer(docker, contResponse.ID)
		}

		contJSON, err := Inspect(docker, contResponse.ID)
		return *contJSON.ContainerJSONBase, err
	}
	return types.ContainerJSONBase{}, errors.New("Cannot connect to the Docker daemon. Is the docker daemon running on this host?")
}

// LogContainer tails container logs to terminal
func LogContainer(docker *client.Docker, contID string) {

	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		// Since       string
		// Timestamps  bool
		Follow: true,
		// Tail        string
	}

	logs, err := docker.Client.ContainerLogs(context.Background(), contID, options)
	defer logs.Close()
	er.CheckError(err)

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, logs)
	er.CheckError(err)
}
