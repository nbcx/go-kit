package exec

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

func Command(name string, arg ...string) *Cmd {
	return &Cmd{
		Cmd: exec.Command(name, arg...),
	}
}

func New(command string) *Cmd {
	cmd := exec.Command("cmd", "/c", command)
	return &Cmd{
		Cmd: cmd,
	}
}

func (p *Cmd) Terminate(sig os.Signal) (err error) {
	if p.Process == nil {
		return nil
	}
	pid := p.Process.Pid
	c := exec.Command("taskkill", "/t", "/f", "/pid", strconv.Itoa(pid))
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

// SetUser not support on windws
func (k *Cmd) SetUser(name string) (err error) {
	log.Printf("Can not set user(%s) on windows", name)
	return nil
}
