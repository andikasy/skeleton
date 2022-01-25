package skeleton

import "os"

type directories struct {
	names []string
}

func New() *directories {
	names := []string{
		"handler",
		"usecase",
		"repository",
		"domain",
		"driver",
	}

	return &directories{
		names: names,
	}
}

func (d *directories) Create() error {
	for _, directory := range d.names {
		_, err := os.Stat(directory)

		if os.IsNotExist(err) {
			err := os.MkdirAll(directory, 0755)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
