package draw

import (
	"fmt"
	"github.com/ssleert/sterm"
	"strconv"
)

// draw static info for bat block
func (s *Info) diskStatic() {
	var (
		rootText = iconed{s.colorList[0], "", "root"}
		homeText = iconed{s.colorList[1], "", "/home"}
		userText = iconed{s.colorList[2], "", "/usr"}
	)

	s.writeIconed(38, s.y+19, &rootText)
	s.writeIconed(38, s.y+20, &homeText)
	s.writeIconed(38, s.y+21, &userText)

	s.tui.WriteString(s.colorMid)
	s.putStr(48, s.y+19, "%")
	s.putStr(48, s.y+20, "%")
	s.putStr(48, s.y+21, "%")

	s.putStr(60, s.y+18, " GiB ")
	s.putStr(62, s.y+19, "GiB")
	s.putStr(62, s.y+20, "GiB")
	s.putStr(62, s.y+21, "GiB")

	s.reset()
	s.putStr(59, s.y+18, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataStat.DiskTotal)))
}

// draw dynamic info for bat block
func (s *Info) diskDynamic() {
	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Disk.RootUsedPerc))
	s.putStr(47, s.y+19, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Disk.RootUsedPerc)))
	s.reset()
	s.putStr(60, s.y+19, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Disk.RootUsed)))

	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Disk.HomeUsedPerc))
	s.putStr(47, s.y+20, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Disk.HomeUsedPerc)))
	s.reset()
	s.putStr(60, s.y+20, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Disk.HomeUsed)))

	s.tui.WriteString(colorForPercent(&s.colorTempr, s.DataDyn.Disk.UsrUsedPerc))
	s.putStr(47, s.y+21, sterm.RevPrint(" "+strconv.Itoa(s.DataDyn.Disk.UsrUsedPerc)))
	s.reset()
	s.putStr(60, s.y+21, sterm.RevPrint(fmt.Sprintf(" %.2f", s.DataDyn.Disk.UsrUsed)))
}
