
package easytier

import (
	"os/exec"
	"syscall"
)

// configureHiddenProcess hides Windows console windows.
// Stdout and Stderr redirection remain unaffected.
func configureHiddenProcess(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000, // CREATE_NO_WINDOW
	}
}
