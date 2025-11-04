package initial

import "go_ci/configurations"

func NewInitialMock(confs *configurations.Configs) Initial {
	return initial{confs: confs}
}
