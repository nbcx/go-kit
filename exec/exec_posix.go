//go:build !windows
// +build !windows

package exec

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

func setupCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Setsid = true
}

func Command(name string, arg ...string) *Cmd {
	cmd := exec.Command(name, arg...)
	setupCmd(cmd)
	return &Cmd{
		Cmd: cmd,
	}
}

func New(command string) *Cmd {
	cmd := exec.Command("/bin/bash", "-c", command)
	setupCmd(cmd)
	return &Cmd{Cmd: cmd}
}

func (p *Cmd) Terminate(sig os.Signal) (err error) {
	if p.Process == nil {
		return
	}
	group, err := os.FindProcess(-1 * p.Process.Pid)
	if err == nil {
		err = group.Signal(sig)
	}
	return err
}

func (k *Cmd) SetUser(name string) (err error) {
	u, err := user.Lookup(name)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return err
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return err
	}
	if k.SysProcAttr == nil {
		k.SysProcAttr = &syscall.SysProcAttr{}
	}
	k.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
	return nil
}
