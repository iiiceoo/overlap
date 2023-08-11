package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/iiiceoo/iprange"
)

var version = "v0.0.1"

var (
	vmode   bool
	file    string
	verbose bool
)

func main() {
	flag.Usage = func() {
		fmt.Print("Usage: overlap [-v] [-f file] [IP ranges...]\n")
		fmt.Print("       overlap [-V]\n\n")
		fmt.Print("Options:\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&vmode, "V", false, "Display the version of overlap.")
	flag.BoolVar(&verbose, "v", false, "Be verbose, display details of overlaping IP ranges.")
	flag.StringVar(&file, "f", "", `The file path of the IP ranges list, which supports the following
formats of IP ranges:
    (IPv4)                  (IPv6)
    172.18.0.1              fd00::1
    172.18.0.0/24           fd00::/64
    172.18.0.1-10           fd00::1-a
    172.18.0.1-172.18.1.10  fd00::1-fd00::1:a`,
	)
	flag.Parse()

	if vmode {
		fmt.Println(version)
		return
	}

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	var rr []string
	rr = append(rr, flag.Args()...)
	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			r := strings.Trim(scanner.Text(), " ")
			if len(r) != 0 {
				rr = append(rr, r)
			}
		}
	}

	if len(rr) == 0 {
		fmt.Println("Empty IP range list")
		return
	}

	ranges, err := iprange.Parse(rr...)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !ranges.IsOverlap() {
		fmt.Println("No overlaping :)")
	} else {
		fmt.Println("Overlaping :(")
	}

	if !verbose {
		return
	}

	n := len(rr)
	var xrr []*iprange.IPRanges
	for i := 0; i < n; i++ {
		xr, _ := iprange.Parse(rr[i])
		xrr = append(xrr, xr)
	}

	fmt.Print("\nDetails:\n")
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			u := xrr[i].Intersect(xrr[j])
			if u.Size().Cmp(big.NewInt(0)) > 0 {
				fmt.Printf("%s and %s overlap at %s\n", rr[i], rr[j], u)
			}
		}
	}
}
