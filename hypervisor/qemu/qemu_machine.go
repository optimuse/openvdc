package qemu

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type State int

const (
	STOPPED State = iota // 0
	STARTING
	RUNNING
	STOPPING
	REBOOTING
	SHUTTINGDOWN
	TERMINATING
	FAILED
)

var stateValues = map[State]string{
	STOPPED:      "STOPPED",
	STARTING:     "STARTING",
	RUNNING:      "RUNNING",
	STOPPING:     "STOPPING",
	REBOOTING:    "REBOOTING",
	SHUTTINGDOWN: "SHUTTINGDOWN",
	FAILED:       "FAILED",
}

var MachineState = map[string]State{
	"STOPPED":      STOPPED,
	"STARTING":     STARTING,
	"RUNNING":      RUNNING,
	"STOPPING":     STOPPING,
	"REBOOTING":    REBOOTING,
	"SHUTTINGDOWN": SHUTTINGDOWN,
	"TERMINATING":  TERMINATING,
	"FAILED":       FAILED,
}

type Machine struct {
	State             State
	Cores             int
	Memory            uint64
	Name              string
	Display           string
	Vnc               string
	MonitorSocketPath string
	SerialSocketPath  string
	AgentSocketPath   string
	Devices           []*Device
	Pidfile           string
	Nics              []Nic
	Drives            map[string]Drive
	Process           *os.Process
	Kvm               bool
}

type Nic struct {
	IfName       string
	Index        string
	Ipv4Addr     string
	MacAddr      string
	Bridge       string
	BridgeHelper string
	Type         string
}

func (m *Machine) promptPattern() string {
	return fmt.Sprintf("openvdc@%s", m.Name)
}

func (m *Machine) HavePrompt() bool {
	c, err := net.Dial("unix", m.SerialSocketPath)
	if err != nil {
		return false
	}

	buf := bufio.NewReader(c)
	matchprompt := make(chan bool, 1)
	defer close(matchprompt)

	go func() {
		defer c.Close()
		c.SetReadDeadline(time.Now().Add(time.Second))
		tries := 0
		for {
			if tries > 10 {
				matchprompt <- false
				return
			}
			line, _, _ := buf.ReadLine()
			if strings.Contains(string(line), m.promptPattern()) {
				matchprompt <- true
				return
			}
			tries = tries + 1
		}
	}()

	// send new line in order to trigger the prompt
	fmt.Fprintf(c, "\n")
	return <-matchprompt
}

func (m *Machine) WaitForPrompt() bool {
	c, err := net.Dial("unix", m.SerialSocketPath)
	if err != nil {
		return false
	}
	defer c.Close()
	buf := bufio.NewReader(c)

	if err := c.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return false
	}
	b, _ := buf.ReadBytes('\n')
	return (strings.Contains(string(b), m.promptPattern()))
}

// since machine struct does not get saved in memory for each instance there may not be any points
// in scheduling states as they are not stored anywhere
func (m *Machine) ScheduleState(nextState State, timeout time.Duration, evaluation func() bool) error {
	errorc := make(chan error)

	go func() {
		defer close(errorc)

		timeoutc := time.After(timeout)
		for {
			select {
			case <-timeoutc:
				errorc <- errors.Errorf("Timed out scheduling state %s", stateValues[nextState])
				return
			default:
				if evaluation() {
					errorc <- nil
					return
				}
			}
		}
	}()

	if err := <-errorc; err != nil {
		return err
	}
	m.State = nextState
	return nil
}

func NewMachine(cores int, mem uint64) *Machine {
	return &Machine{
		Cores:   cores,
		Memory:  mem,
		Drives:  make(map[string]Drive),
		Devices: make([]*Device, 0),
		Display: "none",
	}
}

func (m *Machine) AddNICs(nics []Nic) []*Device {
	var netDevs []*Device
	for _, nic := range nics {
		hostDev := NewDevice(NetType)
		hostDev.AddDriver("tap")
		hostDev.AddDriverOption("ifname", nic.IfName)

		guestDev := NewDevice(DevType)
		guestDev.AddDriver("virtio-net-pci")

		if len(nic.MacAddr) > 0 {
			guestDev.AddDriverOption("mac", nic.MacAddr)
		}

		hostDev.LinkToGuestDevice(nic.IfName, guestDev)
		netDevs = append(netDevs, hostDev)
		netDevs = append(netDevs, guestDev)
		m.Nics = append(m.Nics, nic)
	}
	return netDevs
}

func (m *Machine) AddDevice(device *Device) {
	m.Devices = append(m.Devices, device)
}

func (m *Machine) Start(startCmd string) error {
	cmdLine := &cmdLine{args: make([]string, 0)}

	cmd := exec.Command(startCmd, cmdLine.QemuBootCmd(m)...)
	if err := cmd.Run(); err != nil {
		return errors.Errorf("Failed to execute cmd: %s", cmd.Args)
	}

	m.Process = cmd.Process
	// TODO: add some error handling

	return m.ScheduleState(STARTING, (10 * time.Minute), func() bool {
		err := m.MonitorCommand("info name")
		return (err != nil)
	})
}

func (m *Machine) MonitorCommand(cmd string) error {
	c, err := net.Dial("unix", m.MonitorSocketPath)
	if err != nil {
		return errors.Errorf("Failed to connect to monitor socket %s:", m.MonitorSocketPath)
	}
	defer c.Close()

	fmt.Fprintf(c, "%s\n", cmd)
	return nil
}

func (m *Machine) Stop() error {
	if err := m.MonitorCommand("quit"); err != nil {
		return err
	}
	if err := os.Remove(m.MonitorSocketPath); err != nil {
		return errors.Errorf("Unable remove monitor socket path: %s", m.MonitorSocketPath)
	}
	if err := os.Remove(m.SerialSocketPath); err != nil {
		return errors.Errorf("Unable remove serial socket path: %s", m.SerialSocketPath)
	}
	return nil
}

func (m *Machine) Reboot() error {
	m.MonitorCommand("system_reset")
	return m.ScheduleState(REBOOTING, (10 * time.Minute), func() bool {
		return (m.HavePrompt() == false)
	})
}
