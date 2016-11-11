// +build linux

package lxc

import (
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/axsh/openvdc/hypervisor"
	lxc "gopkg.in/lxc/go-lxc.v2"
)

func init() {
	hypervisor.RegisterProvider("lxc", &LXCHypervisorProvider{})
}

type LXCHypervisorProvider struct {
}

func (p *LXCHypervisorProvider) Name() string {
	return "lxc"
}

func (p *LXCHypervisorProvider) CreateDriver() (hypervisor.HypervisorDriver, error) {
	return &LXCHypervisorDriver{
		log: log.WithField("hypervisor", "lxc"),
		lxcpath: lxc.DefaultConfigPath(),
		name: "lxc-test",
		// Set pre-defined template option from gopkg.in/lxc/go-lxc.v2/options.go
		template: lxc.DownloadTemplateOption,
	}, nil
}

type LXCHypervisorDriver struct {
	log        *log.Entry
	imageName  string
	hostName   string
	lxcpath    string
	template   lxc.TemplateOptions
	name       string
}

func (d *LXCHypervisorDriver) CreateInstance() error {

	c, err := lxc.NewContainer(d.name, d.lxcpath)
	if err != nil {
		d.log.Errorln(err)
		return err
	}

	d.log.Infoln("Creating lxc-container...")

	if err := c.Create(d.template); err != nil {
		d.log.Errorln(err)
		return err
	}
	return nil
}

func (d *LXCHypervisorDriver) DestroyInstance() error {
	c, err := lxc.NewContainer(d.name, d.lxcpath)
	if err != nil {
		d.log.Errorln(err)
		return err
	}

	d.log.Infoln("Destroying lxc-container..")
	if err := c.Destroy(); err != nil {
		d.log.Errorln(err)
		return err
	}
	return nil
}

func (d *LXCHypervisorDriver) StartInstance() error {

	c, err := lxc.NewContainer(d.name, d.lxcpath)
	if err != nil {
		d.log.Errorln(err)
		return err
	}

	d.log.Infoln("Starting lxc-container...")
	if err := c.Start(); err != nil {
		d.log.Errorln(err)
		return err
	}

	d.log.Infoln("Waiting for lxc-container to start networking")
	if _, err := c.WaitIPAddresses(5 * time.Second); err != nil {
		d.log.Errorln(err)
		return err
	}
	return nil
}

func (d *LXCHypervisorDriver) StopInstance() error {

	c, err := lxc.NewContainer(d.name, d.lxcpath)
	if err != nil {
		d.log.Errorln(err)
		return err
	}

	d.log.Infoln("Stopping lxc-container..")
	if err := c.Stop(); err != nil {
		d.log.Errorln(err)
		return err
	}
	return nil
}