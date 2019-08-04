package stats

import (
	"errors"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// CPUData simple data regarding the CPU used for the machine this application
// is running on.
type CPUData struct {
	Vendor     string  `json:"vendor,omitempty"`
	PhysicalID string  `json:"physicalId,omitempty"`
	Family     string  `json:"family,omitempty"`
	Cores      int32   `json:"cores,omitempty"`
	Model      string  `json:"model,omitempty"`
	Speed      float64 `json:"speedMhz,omitempty"`
}

// Host simple data regarding the machines host name for the machine this
// application is running on.
type Host struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Uptime string `json:"uptime,omitempty"`
}

// Disk snapshot of the disk usage for this application
type Disk struct {
	Total        uint64  `json:"totalBytes,omitempty"`
	Free         uint64  `json:"freeBytes,omitempty"`
	Used         uint64  `json:"usedBytes,omitempty"`
	UsedPerc     float64 `json:"usedPercentage,omitempty"`
	MaxPercLimit float64 `json:"maxPercentage,omitempty"`
	Status       string  `json:"status,omitempty"`
}

// Mem memory snapshot of the disk usage for this application
type Mem struct {
	Total        uint64  `json:"totalBytes,omitempty"`
	Free         uint64  `json:"freeBytes,omitempty"`
	Used         uint64  `json:"usedBytes,omitempty"`
	UsedPerc     float64 `json:"usedPercentage,omitempty"`
	MaxPercLimit float64 `json:"maxPercentage,omitempty"`
	Status       string  `json:"status,omitempty"`
}

// CPUs aggregated CPU data for the machine this application is running on
type CPUs struct {
	Cnt int       `json:"count,omitempty"`
	CPU []CPUData `json:"cpu,omitempty"`
}

// SystemDetails details about the machine this application is running on.
type SystemDetails struct {
	OS   string `json:"os,omitempty"`
	CPUs *CPUs  `json:"cpu,omitempty"`
	Disk *Disk  `json:"disk,omitempty"`
	Host *Host  `json:"host,omitempty"`
	Mem  *Mem   `json:"mem,omitempty"`
}

// Options initialization data/configurations for Stats
type Options struct {
	//StartTime the time that this application was started (e.g., time.Now())
	StartTime time.Time `json:"startTime,omitempty"`
}

// Stats aggregated snapshot of the state of the machine this application is
// running on
type Stats struct {
	//SysDetails details about the system/machine the application is running on
	SysDetails *SystemDetails `json:"sysDetails,omitempty"`
	//AppDetails details about the application
	AppDetails map[string]interface{} `json:"appDetails,omitempty"`
}

// pullMem gather memory usage/availablity info
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func (s *Stats) pullMem() (err error) {
	if s == nil {
		err = errors.New("Stats should be initialized before use. (pullMem)")
		return
	}
	vmStat, err := mem.VirtualMemory()

	if err != nil {
		return
	}
	if s.SysDetails.Mem == nil {
		s.SysDetails.Mem = &Mem{}
	}

	s.SysDetails.Mem.Total = vmStat.Total
	s.SysDetails.Mem.Free = vmStat.Free
	s.SysDetails.Mem.Used = vmStat.Used
	s.SysDetails.Mem.UsedPerc = vmStat.UsedPercent

	return
}

// pullCPU gather CPU info
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func (s *Stats) pullCPU() (err error) {
	if s == nil {
		err = errors.New("Stats should be initialized before use. (pullCPU)")
		return
	}
	cpuStat, err := cpu.Info()

	if err != nil {
		return
	}
	if s.SysDetails.CPUs == nil {
		s.SysDetails.CPUs = &CPUs{}
	}
	s.SysDetails.CPUs.Cnt = len(cpuStat)
	for _, cpu := range cpuStat {
		var st = CPUData{}
		st.Cores = cpu.Cores
		st.Family = cpu.Family
		st.Model = cpu.Model
		st.Vendor = cpu.VendorID
		st.Speed = cpu.Mhz
		st.PhysicalID = cpu.PhysicalID
		s.SysDetails.CPUs.CPU = append(s.SysDetails.CPUs.CPU, st)
	}

	return
}

// pullDisk gather disk usage/availablity info
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func (s *Stats) pullDisk() (err error) {
	if s == nil {
		err = errors.New("Stats should be initialized before use. (pullDisk)")
		return
	}
	diskStat, err := disk.Usage("/")
	if err != nil {
		return
	}
	if s.SysDetails.Disk == nil {
		s.SysDetails.Disk = &Disk{}
	}
	s.SysDetails.Disk.Total = diskStat.Total
	s.SysDetails.Disk.Free = diskStat.Free
	s.SysDetails.Disk.Used = diskStat.Used
	s.SysDetails.Disk.UsedPerc = diskStat.UsedPercent

	return
}

// pullHost gather host info
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func (s *Stats) pullHost() (err error) {
	if s == nil {
		err = errors.New("Stats should be initialized before use. (pullHost)")
		return
	}

	hostStat, err := host.Info()
	if err != nil {
		return err
	}
	if s.SysDetails.Host == nil {
		s.SysDetails.Host = &Host{}
	}
	s.SysDetails.Host.ID = hostStat.HostID
	s.SysDetails.Host.Name = hostStat.Hostname
	s.SysDetails.Host.Uptime = (time.Duration(hostStat.Uptime) * time.Second).String()

	return
}

// pullUptime if startTime is provided, then calcuate the amount of time this
// application has been running. Otherwise sets uptime as UNKNOWN
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func (s *Stats) pullUptime() (err error) {
	if s == nil {
		err = errors.New("Stats should be initialized before use. (pullUptime)")
		return
	}
	s.AppDetails["uptime"] = "UNKNOWN"

	if st, ok := s.AppDetails["startTime"]; ok {
		start := st.(time.Time)
		uptime := time.Since(start)
		s.AppDetails["uptime"] = uptime.String()
	}
	return
}

// PullStats gather primarily system stats. Also calculates application uptime
// if possible
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func (s *Stats) PullStats() (err error) {
	if s == nil {
		err = errors.New("Stats should be initialized before use. (PullStats)")
		return
	}

	if s.SysDetails == nil {
		s.SysDetails = &SystemDetails{}
	}
	// s.SysDetails.Mem = &Mem{}
	// s.SysDetails.Disk = &Mem{}
	err = s.pullMem()
	err = s.pullCPU()
	err = s.pullDisk()
	err = s.pullHost()
	err = s.pullUptime()

	if runtime.GOOS != s.SysDetails.OS {
		s.SysDetails.OS = runtime.GOOS
	}
	return
}

// NewStats create a new stat data structure
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func NewStats(opts Options) Stats {

	s := Stats{}
	if !opts.StartTime.IsZero() {
		s.AppDetails = make(map[string]interface{})
		s.AppDetails["startTime"] = opts.StartTime
	}

	return s
}
