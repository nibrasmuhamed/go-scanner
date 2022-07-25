/*
Copyright © 2022 Nibras Muhammed nibrasmn027@gmail.com

*/
package cmd

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/j-keck/arping"
	"github.com/nibrasmuhamed/go-scanner/macvendorfinder"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "subcommand to scan network",
	Long:  `scan subcommand used to scan network followed by router ip address.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("please provide router ip to scan.")
		} else {
			hosts, err := Cdirgetter(args[0])
			if err != nil {
				print("Please Input a valid CIDR in this format ")
				os.Exit(0)
			}
			fmt.Print("\u001b[37;1m|HOST SCANNED|\t|STATE|\t|Mac Addresses|\n\u001b[0m")
			for i := range hosts {
				mac, host := Arpscan_lan(hosts[i])
				if isnotempty(host) {
					printer(mac, host, MAXPORT+1, "")
					hosts_online := append(hosts_online, host)
					fmt.Println(hosts_online)
					// time.Sleep(time.Second)
					fmt.Println(macvendorfinder.Macfind(mac))
				}
			}
			fmt.Println(cmd.Aliases)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const (
	MAXPORT = 65535
)

var hosts_online []string

func Arpscan_lan(ips string) (string, string) {
	ip := net.ParseIP(ips)
	arping.SetTimeout(500 * time.Millisecond)
	HwAddr, _, err := arping.Ping(ip)
	mac := HwAddr.String()
	if err == arping.ErrTimeout {
		return mac, ""
	} else if err != nil {
		if strings.Contains(err.Error(), "operation not") {
			print("Please run as root\n")
			os.Exit(1)
		} else if strings.Contains(err.Error(), "ip+net") {
			return mac, "Fail in net resources occurred Running again" + "\n"
			// Arpscan_lan(ips)

		} else if strings.Contains(err.Error(), "no usable interface found") {
			print("You put CIDR of another net OR Try Run same as root\n")
			os.Exit(0)
			return mac, "Probably you put a CIDR outside ur net" + "\n"

		} else {

			print("Running again: Unknown Error succedeed try run program in root\n")
			os.Exit(1)
			return mac, "Running again: Unknown Error succedeed try run program in root\n"

		}
	} else {
		return mac, ips
	}
	return mac, "Error"
}

func Cdirgetter(cidr string) ([]string, error) {
	var hosts []string
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		print("Please Input a valid CIDR in this format (192.168.1.1/24, 10.0.0.0/8)")
		os.Exit(0)
	}
	mascara := binary.BigEndian.Uint32(subnet.Mask)
	fAddr := binary.BigEndian.Uint32(subnet.IP)
	lAddr := (fAddr & mascara) | (mascara ^ 0xffffffff)
	for i := fAddr; i <= lAddr; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ips := ip.String()
		hosts = append(hosts, ips)
	}
	return hosts, err
}

func isnotempty(s string) bool {
	return len(s) > 0
}

func printer(mac string, host string, port int, service string) {
	if !isnotempty(service) {
		service = "Not Found"
	}
	if port == 0 {
		print("")
	} else if port != MAXPORT+1 {

		print(fmt.Sprintf("%d\tOpen\t"+service+"\n", port))
	} else {
		if strings.Contains(host, "Running") {
			print(host + "\n")
			fmt.Print("|HOST SCANNED|\t|STATE|\t|Mac Addresses|\n")
		} else {
			print(fmt.Sprintf("%s\t\u001b[32;1mOnline\u001b[0m\t"+mac+"\t", host))
		}
	}
}
