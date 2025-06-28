package initial

import "go_project_structure_be/configurations"

func NewInitialMock(confs *configurations.Configs) Initial {
	return initial{confs: confs}
}
