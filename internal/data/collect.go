package data

// os depent functions

import (
	"github.com/ssleert/sfolib"
	"strings"
)

// +================= static info functions =================+

func getCpuModel(ch chan string, errch chan error) {
	var result strings.Builder
	result.Grow(100)

	s, err := sfolib.ReadLines("/proc/cpuinfo", 5)
	if err != nil {
		ch <- ""
		errch <- err
	}

	cpuName := strings.Fields(s[4])[3:]
	cpuVendor := strings.Fields(s[1])[2]
	switch cpuVendor {
	case "AuthenticAMD":
		for _, e := range cpuName {
			switch {
			case strings.Contains(e, "Ryzen"):
				result.WriteRune('R')
				result.WriteString(cpuName[2])
				result.WriteRune('-')
				result.WriteString(cpuName[3])
				break
			case strings.Contains(e, "Turion"):
				result.WriteString("Turion ")
				result.WriteString(cpuName[2])
				result.WriteRune(' ')
				result.WriteString(cpuName[3])
			}
		}
		if result.Len() == 0 {
			result.WriteString("AMD")
		}
	case "GenuineIntel":
		for i, e := range cpuName {
			switch {
			case strings.Contains(e, "i3"),
				strings.Contains(e, "i5"),
				strings.Contains(e, "i7"),
				strings.Contains(e, "i9"):
				result.WriteString(cpuName[i])
				break
			case strings.Contains(e, "Xeon"):
				result.WriteString(cpuName[3])
				break
			case strings.Contains(e, "Celeron"):
				result.WriteString("Celeron ")
				result.WriteString(cpuName[3])
				break
			case strings.Contains(e, "Duo"):
				result.WriteString("2d-")
				result.WriteString(cpuName[4])
				break
			case strings.Contains(e, "Pentium"):
				result.WriteString("Pentium ")
				result.WriteString(cpuName[2])
				break
			}
		}
		if result.Len() == 0 {
			result.WriteString("INTEL")
		}
	}

	ch <- result.String()
	errch <- err
}

// +=========================================================+