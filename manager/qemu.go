package manager

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mainflux/mainflux/logger"
)

const script = "cmd/manager/script/launch-qemu.sh"

type Config struct {
	HDAFile          string `env:"HDA_FILE" envDefault:"cmd/manager/img/focal-server-cloudimg-amd64.qcow2"`
	GuestSizeInMB    int    `env:"GUEST_SIZE_IN_MB" envDefault:"4096"`
	SevGuest         bool   `env:"SEV_GUEST" envDefault:"1"`
	SmpNCPUs         int    `env:"SMP_NCPUS" envDefault:"4"`
	Console          string `env:"CONSOLE" envDefault:"serial"`
	VNCPort          string `env:"VNC_PORT"`
	UseVirtio        bool   `env:"USE_VIRTIO" envDefault:"1"`
	UEFIBiosCode     string `env:"UEFI_BIOS_CODE" envDefault:"/usr/share/OVMF/OVMF_CODE.fd"`
	UEFIBiosVarsOrig string `env:"UEFI_BIOS_VARS_ORIG" envDefault:"/usr/share/OVMF/OVMF_VARS.fd"`
	UEFIBiosVarsCopy string `env:"UEFI_BIOS_VARS_COPY" envDefault:"cmd/manager/img/OVMF_VARS.fd"`
	CBitPos          int    `env:"CBITPOS" envDefault:"51"`
	HostHTTPPort     int    `env:"HOST_HTTP_PORT" envDefault:"9301"`
	GuestHTTPPort    int    `env:"GUEST_HTTP_PORT" envDefault:"9031"`
	HostGRPCPort     int    `env:"HOST_GRPC_PORT" envDefault:"7020"`
	GuestGRPCPort    int    `env:"GUEST_GRPC_PORT" envDefault:"7002"`
	EnableFileLog    bool   `env:"ENABLE_FILE_LOG" envDefault:"0"`
	ExecQemuCmdLine  bool   `env:"EXEC_QEMU_CMDLINE" envDefault:"0"`
	Sudo             bool   `env:"SUDO" envDefault:"0"`
}

// RunQemuVM runs a QEMU virtual machine: constructs the QEMU command line arguments by executing the launch-qemu.sh,
// extracts the QEMU command and its arguments, and starts the QEMU process
func RunQemuVM(qemuConfig Config, logger logger.Logger) (*exec.Cmd, error) {
	args := constructQemuCmd(qemuConfig)

	output, err := ExeShCmdStdout(script, args...)
	if err != nil {
		return nil, err
	}

	command, args := ExtractCmdAndArgs(output, qemuConfig.Sudo)
	cmd, err := RunCmdStart(command, args...)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// constructQemuCmd constructs the command line arguments for executing the launch-qemu.sh script,
// which in turn provides the QEMU command to run
func constructQemuCmd(config Config) []string {
	args := []string{
		"-hda", config.HDAFile,
		"-mem", strconv.Itoa(config.GuestSizeInMB),
	}

	if !config.SevGuest {
		args = append(args, "-nosev")
	}

	if config.VNCPort != "" {
		args = append(args, "-vnc", config.VNCPort)
	}

	if config.UseVirtio {
		args = append(args, "-virtio")
	}

	if config.EnableFileLog {
		args = append(args, "-filelog")
	}

	if config.ExecQemuCmdLine {
		args = append(args, "-exec")
	}

	if config.HostHTTPPort != 0 {
		args = append(args, "-hosthttp", strconv.Itoa(config.HostHTTPPort))
	}

	if config.GuestHTTPPort != 0 {
		args = append(args, "-guesthttp", strconv.Itoa(config.GuestHTTPPort))
	}

	if config.HostGRPCPort != 0 {
		args = append(args, "-hostgrpc", strconv.Itoa(config.HostGRPCPort))
	}

	if config.GuestGRPCPort != 0 {
		args = append(args, "-guestgrpc", strconv.Itoa(config.GuestGRPCPort))
	}

	if config.UEFIBiosCode != "" {
		args = append(args, "-bios", config.UEFIBiosCode)
	}

	if config.UEFIBiosVarsOrig != "" {
		args = append(args, "-origuefivars", config.UEFIBiosVarsOrig)
	}

	if config.UEFIBiosVarsCopy != "" {
		args = append(args, "-copyuefivars", config.UEFIBiosVarsCopy)
	}

	if config.Console != "" {
		args = append(args, "-console", config.Console)
	}

	if config.CBitPos != 0 {
		args = append(args, "-cbitpos", strconv.Itoa(config.CBitPos))
	}

	if config.SmpNCPUs != 0 {
		args = append(args, "-smp", strconv.Itoa(config.SmpNCPUs))
	}

	return args
}

// ExeShCmdStdout executes a shell command capturing the standard output
func ExeShCmdStdout(command string, args ...string) (string, error) {
	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command(command, args...)

	// Capture stdout and stderr using buffers
	cmd.Stdout = io.MultiWriter(&stdoutBuf, os.Stdout)
	cmd.Stderr = io.MultiWriter(&stderrBuf, os.Stderr)

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing command '%s': %s", cmd.String(), err)
	}

	return stdoutBuf.String(), nil
}

// ExtractCmdAndArgs extracts the command and its arguments from the output string
func ExtractCmdAndArgs(cmdLine string, sudo bool) (string, []string) {
	lines := strings.Split(cmdLine, "\n")
	if len(lines) == 0 {
		return "", nil
	}

	parts := strings.Fields(lines[0])
	if len(parts) == 0 {
		return "", nil
	}

	if sudo {
		parts = append([]string{"sudo"}, parts...)
	}

	cmd := parts[0]
	args := parts[1:]

	return cmd, args
}

// RunCmdOutput runs the specified command and returns its standard output as a string
func RunCmdOutput(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error executing command '%s': %s", cmd.String(), err)
	}

	return string(output), nil
}

// RunCmdStart starts the specified command and returns the *exec.Cmd for the running process
func RunCmdStart(command string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(command, args...)

	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("error starting command '%s': %s", cmd.String(), err)
	}

	return cmd, nil
}
