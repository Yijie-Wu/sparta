package service

import (
	"context"
	"fmt"
	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/spf13/viper"
	"sparta/service/dto"
)

var hostService *HostService

type HostService struct {
	BaseService
}

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (u *HostService) Shutdown(iShutdownHostDTO dto.ShutdownHostDTO) error {
	var errResult error

	hostIP := iShutdownHostDTO.HostIP
	fmt.Println(hostIP)

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.username"),
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  fmt.Sprintf("%s,", hostIP),
		ModuleName: "command",
		Args:       viper.GetString("ansible.shutdownArgs"),
		ExtraVars: map[string]interface{}{
			"ansible_password": viper.GetString("ansible.password"),
		},
	}

	iAdhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "oneline",
	}

	errResult = iAdhoc.Run(context.TODO())

	return errResult
}
