package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type mouse struct {
	id      string
	cur_val int
	new_val int
}

func good_exit() {
	fmt.Println("You can now close this window.\n\t\tPress 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
}

func print_err_help() {
	fmt.Println("If you are seeing this message, you probably ran into an issue with the program.")
	fmt.Println("\nHere are some things you can try to fix the issue:")
	fmt.Println("1:\tMake sure you are running the program as an administrator.")
	fmt.Println("2:\tYour mouse is not supported by the program.")
	fmt.Println("3:\tYou are not running the program on Windows.")
	fmt.Println("If none of these help, please open an issue on GitHub: https://github.com/carter4299/windows-scroll-direction")
	good_exit()
}

func get_val(s string) (int, error) {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return int(r - '0'), nil
		}
	}
	return -1, fmt.Errorf("No value found in string, mouse ID: %s", s)
}

func run_powershell_script(script string) (string, error) {
	cmd := exec.Command("powershell", "-Command", script)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func get_mouse_id() (string, error) {
	script := `Get-WmiObject Win32_PointingDevice | Select-Object DeviceID`
	device_id, err := run_powershell_script(script)
	if err != nil {
		return "", err
	}
	device_id = strings.TrimSpace(device_id)
	return strings.Split(device_id, "\n")[2], nil
}

func get_mouse_vals(id string) (int, int, error) {
	script := fmt.Sprintf(`Get-ItemProperty -Path "HKLM:\\SYSTEM\\CurrentControlSet\\Enum\\%s\\Device Parameters" -Name "FlipFlopWheel"`, id)
	ffw_value, err := run_powershell_script(script)
	if err != nil {
		return -1, -1, err
	}
	ffw_value = strings.TrimSpace(ffw_value)
	ffw_value = strings.Split(ffw_value, "\n")[0]
	cur_val, err := get_val(ffw_value)
	if err != nil || cur_val < 0 || cur_val > 1 {
		return -1, -1, err
	}
	return cur_val, 1 - cur_val, nil
}

func main() {
	fmt.Println("Welcome to the FlipFlopWheel value changer!\n\tStarting...")

	fairwell := [2]string{"Reverse Scroll", "Natural Scroll"}
	var user mouse
	var err error

	user.id, err = get_mouse_id()
	if err != nil {
		fmt.Println("Error getting device ID:", err)
		print_err_help()
	}
	fmt.Println("Your Mouse ID is:", user.id)

	user.cur_val, user.new_val, err = get_mouse_vals(user.id)
	if err != nil {
		fmt.Println("Error getting FlipFlopWheel value:", err)
		print_err_help()
	}
	fmt.Printf("Your current FlipFlopWheel value is %s(%d)...\tChanging to %s(%d)\n", fairwell[user.cur_val], user.cur_val, fairwell[user.new_val], user.new_val)

	script := fmt.Sprintf(`Set-ItemProperty -Path "HKLM:\\SYSTEM\\CurrentControlSet\\Enum\\%s\\Device Parameters" -Name "FlipFlopWheel" -Value %d`, user.id, user.new_val)
	_, err = run_powershell_script(script)
	if err != nil {
		fmt.Println("Error setting new FlipFlopWheel value:", err)
		print_err_help()
	}

	fmt.Printf("FlipFlopWheel value was successfully changed to %s(%d)\n", fairwell[user.new_val], user.new_val)
	fmt.Println("\nDon't forget to restart your computer to apply changes.")
	good_exit()
}
