// +build !windows

package console

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

const CommandExample = `
% openvdc console i-000000001
root@i-00000001#

% echo ls | openvdc console i-000000001
% cat test.sh | openvdc console i-000000001
`

func (s *SshConsole) bindFDs(session *ssh.Session) (func() error, error) {
	closeFunc := func() error { return nil }
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			s.log.WithError(err).Warn("Failed to get console size. Set to 80x40")
			w = 80
			h = 40
		}
		modes := ssh.TerminalModes{}
		term, ok := os.LookupEnv("TERM")
		if !ok {
			term = defaultTermInfo
		}
		if err := session.RequestPty(term, h, w, modes); err != nil {
			return nil, errors.Wrap(err, "Failed session.RequestPty")
		}

		origstate, err := terminal.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			return nil, errors.Wrap(err, "Failed terminal.MakeRaw")
		}

		closeFunc = func() error {
			if err := terminal.Restore(int(os.Stdin.Fd()), origstate); err != nil {
				if errno, ok := err.(syscall.Errno); (ok && errno != 0) || !ok {
					return errors.Wrap(err, "Failed terminal.Restore")
				}
			}
			return nil
		}
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	return closeFunc, nil
}

func (s *SshConsole) signal(c chan<- os.Signal) error {
	signal.Notify(c, os.Interrupt, syscall.SIGWINCH)
	return nil
}

func (s *SshConsole) signalHandle(sig os.Signal, session *ssh.Session) error {
	fd := int(os.Stdin.Fd())
	switch sig {
	case syscall.SIGWINCH:
		w, h, err := terminal.GetSize(fd)
		if err != nil {
			return errors.Wrap(err, "Failed terminal.GetSize")
		}
		winchMsg := struct {
			Columns uint32
			Rows    uint32
			Width   uint32
			Height  uint32
		}{
			Columns: uint32(w),
			Rows:    uint32(h),
			Width:   uint32(w * 8),
			Height:  uint32(h * 8),
		}
		if _, err := session.SendRequest("window-change", false, ssh.Marshal(&winchMsg)); err != nil {
			return errors.Wrap(err, "Failed session.SendRequest(window-change)")
		}
	case os.Interrupt:
		sshSig := ssh.SIGINT
		if err := session.Signal(sshSig); err != nil {
			return errors.Wrapf(err, "Failed session.Signal: %s", sig)
		}
	default:
		return errors.Errorf("Unhandled signal: %s", sig)
	}
	return nil
}
