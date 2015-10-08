package system

import (
	"errors"
	"strings"

	"github.com/aelsabbahy/goss/util"
)

type RpmPackage struct {
	name      string
	versions  []string
	loaded    bool
	installed bool
}

func NewRpmPackage(name string, system *System) Package {
	return &RpmPackage{name: name}
}

func (p *RpmPackage) setup() {
	if p.loaded {
		return
	}
	p.loaded = true
	cmd := util.NewCommand("rpm", "-q", "--nosignature", "--nohdrchk", "--nodigest", "--qf", "%{VERSION}\n", p.name)
	if err := cmd.Run(); err != nil {
		return
	}
	p.installed = true
	p.versions = strings.Split(strings.TrimSpace(cmd.Stdout.String()), "\n")
}

func (p *RpmPackage) Name() string {
	return p.name
}

func (p *RpmPackage) Exists() (interface{}, error) { return p.Installed() }

func (p *RpmPackage) Installed() (interface{}, error) {
	p.setup()

	return p.installed, nil
}

func (p *RpmPackage) Versions() ([]string, error) {
	p.setup()
	if len(p.versions) == 0 {
		return p.versions, errors.New("Package version not found")
	}
	return p.versions, nil
}
