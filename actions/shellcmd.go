// +build !windows
// +build !plan9

package actions

func getShellCmd() (cmd, args string) {
	return "sh", "-c"
}
