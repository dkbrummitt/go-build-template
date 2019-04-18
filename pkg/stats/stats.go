package stats

import (
	"fmt"
	"net"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

const (
	defaultMax = 80.0
)

type StatsOptions struct {
	MaxDisk float64
	MaxHost float64
	MaxMem  float64
	MaxNet  float64
}

type CPUData struct {
	Vendor     string  `json:"vendor,omitempty"`
	PhysicalID string  `json:"physicalId,omitempty"`
	Family     string  `json:"family,omitempty"`
	Cores      int32   `json:"cores,omitempty"`
	Model      string  `json:"model,omitempty"`
	Speed      float64 `json:"speedMhz,omitempty"`
}
type Host struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Uptime string `json:"uptime,omitempty"`
}
type Disk struct {
	Total        uint64  `json:"totalBytes,omitempty"`
	Free         uint64  `json:"freeBytes,omitempty"`
	Used         uint64  `json:"usedBytes,omitempty"`
	UsedPerc     float64 `json:"usedPercentage,omitempty"`
	MaxPercLimit float64 `json:"maxPercentage,omitempty"`
	Status       string  `json:"status,omitempty"`
}
type NetData struct {
	Name    string `json:"name,omitempty"`
	MacAddr string `json:"macAddress,omitempty"`
}
type Mem struct {
	Total        uint64  `json:"totalBytes,omitempty"`
	Free         uint64  `json:"freeBytes,omitempty"`
	Used         uint64  `json:"usedBytes,omitempty"`
	UsedPerc     float64 `json:"usedPercentage,omitempty"`
	MaxPercLimit float64 `json:"maxPercentage,omitempty"`
	Status       string  `json:"status,omitempty"`
}
type Nets struct {
	Cnt int       `json:"count,omitempty"`
	Net []NetData `json:"cpu,omitempty"`
}
type CPUs struct {
	Cnt int       `json:"count,omitempty"`
	CPU []CPUData `json:"cpu,omitempty"`
}
type Stats struct {
	OS   string `json:"os,omitempty"`
	CPUs *CPUs  `json:"cpu,omitempty"`
	Disk *Disk  `json:"disk,omitempty"`
	Host *Host  `json:"host,omitempty"`
	Mem  *Mem   `json:"mem,omitempty"`
	Nets *Nets  `json:"net,omitempty"`
}

func NewStats(opts StatsOptions) Stats {
	s := Stats{}
	validateOpts(&opts) //validate and set defaults
	s.Disk = &Disk{}
	s.Disk.MaxPercLimit = opts.MaxDisk
	s.Mem = &Mem{}
	s.Mem.MaxPercLimit = opts.MaxMem

	return s
}

func validateOpts(opts *StatsOptions) {

	if opts.MaxDisk == 0.0 {
		opts.MaxDisk = defaultMax
	}
	if opts.MaxHost == 0.0 {
		opts.MaxHost = defaultMax
	}
	if opts.MaxMem == 0.0 {
		opts.MaxMem = defaultMax
	}
	if opts.MaxNet == 0.0 {
		opts.MaxNet = defaultMax
	}
}

func (s *Stats) UpdateOpts(opts StatsOptions) {
	validateOpts(&opts) //validate and set defaults

	s.Disk.MaxPercLimit = opts.MaxDisk
	s.Mem.MaxPercLimit = opts.MaxMem
}

func (s *Stats) pullMem() error {
	vmStat, err := mem.VirtualMemory()

	if err != nil {
		fmt.Println(err)
		return err
	}
	s.Mem.Total = vmStat.Total
	s.Mem.Free = vmStat.Free
	s.Mem.Used = vmStat.Used
	s.Mem.UsedPerc = vmStat.UsedPercent
	if s.Mem.UsedPerc <= s.Mem.MaxPercLimit {
		s.Mem.Status = "OK"
	}
	if s.Mem.UsedPerc > s.Mem.MaxPercLimit {
		s.Mem.Status = "NOT OK"
	}
	return err
}

func (s *Stats) pullCPU() error {
	cpuStat, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if s.CPUs == nil {
		s.CPUs = &CPUs{}
	}
	s.CPUs.Cnt = len(cpuStat)
	for _, cpu := range cpuStat {
		var st = CPUData{}
		st.Cores = cpu.Cores
		st.Family = cpu.Family
		st.Model = cpu.Model
		st.Vendor = cpu.VendorID
		st.Speed = cpu.Mhz
		st.PhysicalID = cpu.PhysicalID
		s.CPUs.CPU = append(s.CPUs.CPU, st)
	}
	return err
}

func (s *Stats) pullDisk() error {
	diskStat, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	s.Disk.Total = diskStat.Total
	s.Disk.Free = diskStat.Free
	s.Disk.Used = diskStat.Used
	s.Disk.UsedPerc = diskStat.UsedPercent
	if s.Disk.UsedPerc <= s.Disk.MaxPercLimit {
		s.Disk.Status = "OK"
	}
	if s.Disk.UsedPerc > s.Disk.MaxPercLimit {
		s.Disk.Status = "NOT OK"
	}
	return err
}

func (s *Stats) pullHost() error {
	var err error

	hostStat, err := host.Info()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if s.Host == nil {
		s.Host = &Host{}
	}
	s.Host.ID = hostStat.HostID
	s.Host.Name = hostStat.Hostname
	// var up time.Duration = time.Duration(hostStat.Uptime) * time.Nanosecond)
	s.Host.Uptime = (time.Duration(hostStat.Uptime) * time.Second).String()
	return err
}

func (s *Stats) pullNet() error {
	var err error
	netStats, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return err
	}
	if s.Nets == nil {
		s.Nets = &Nets{}
	}
	s.Nets.Cnt = len(netStats)
	for _, netStat := range netStats {
		net := NetData{}
		net.Name = netStat.Name
		net.MacAddr = netStat.HardwareAddr.String()
		s.Nets.Net = append(s.Nets.Net, net)
		// TODO Support networking flags. // netStat.Flags.String()
		// TODO Support IP Address(es) for loop required. //netStat.Addrs()

	}
	return err
}

func (s *Stats) PullStats() {
	s.pullMem()
	s.pullCPU()
	s.pullDisk()
	s.pullHost()
	s.pullNet()
	fmt.Println(runtime.GOOS)

	if runtime.GOOS != s.OS {
		s.OS = runtime.GOOS
	}
}
