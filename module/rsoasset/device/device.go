package device

import (
  "fmt"
  "os"
  "sort"
  "path/filepath"
  "encoding/json"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
  "github.com/elastic/beats/metricbeat/mb"
  
  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/realsighton/rso/servers/common/esmodels"  
)

var (
  isInit = true
  datadir string
  old  *esmodels.DeviceAssetType
  cur  *esmodels.DeviceAssetType
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("rsoasset", "device", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
  *esmodels.DeviceAssetType
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The rsoasset device metricset is experimental.")

	config := struct{
    DataDir    string   `config:"datadir"`
  }{
    DataDir: "data",
  }
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

  println("##################### DataDir", config.DataDir)
  if isInit == true {
    isInit = false
    initDeviceData(config.DataDir)
  }

	return &MetricSet{
		BaseMetricSet: base,
    DeviceAssetType: new(esmodels.DeviceAssetType),
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {
  var err error 

  // System
  m.System, err = getSystem()
  if err != nil {
    return
  }

  // PcBiosType
  m.Bios, err = getBios()
  if err != nil {
    return
  }

  // ProcessorType
  m.Processors, err = getProcessors()
  if err != nil {
    return
  }

  // DiskType
  m.Disks, err = getDisks()
  if err != nil {
    return
  }

  // DriveType
  m.Drives, err = getDrives()
  if err != nil {
    return
  }

  // NicType
  m.Nics, err = getNics()
  if err != nil {
    return
  }

  // NwConfigType
  m.NicConfigs, err = getNicConfigs()
  if err != nil {
    return
  }

  // VideoControllerType
  m.Videos, err = getVideoController()
  if err != nil {
    return
  }

  m.Memories, err = getMemories()
  if err != nil {
    return
  }

  cur.System     = m.System 
  cur.Bios       = m.Bios
  cur.Processors = m.Processors
  cur.Disks      = m.Disks
  cur.Drives     = m.Drives
  cur.Nics       = m.Nics 
  cur.NicConfigs = m.NicConfigs 
  cur.Videos     = m.Videos
  cur.Memories   = m.Memories 

  isEq := checkEqualDevice()
  
  if isEq {
    println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& no changed")
    writeDeviceData()
    return 
  }

  old = cur
  
  writeDeviceData()

	report.Event(mb.Event{
		MetricSetFields: common.MapStr{
      "system": m.System,
      "bios": m.Bios,
      "processors": m.Processors,
      "disks": m.Disks,
      "drives": m.Drives,
      "nics": m.Nics,
      "nicconfigs": m.NicConfigs,
      "videos": m.Videos,
      "memories": m.Memories,
		},
  })
}

func checkEqualDevice() bool {
  var isEq bool

  if old.System == nil {
    println("hear 1")
    return false
  }

  isEq = cur.System.Equals(old.System)
  if !isEq {
    println("hear 2")
    return false
  }

  if old.Bios == nil {
    println("hear 3")
    return false
  }

  isEq = cur.Bios.Equals(old.Bios)
  if !isEq {
    println("hear 4")
    return false
  }

  if cur.Processors != nil && old.Processors == nil {
    println("here 5")
    return false
  }

  if len(cur.Processors) != len(old.Processors) {
    println("here 6")
    return false
  }

  sort.Slice(cur.Processors, func(i, j int) bool {return *cur.Processors[i].Name < *cur.Processors[j].Name })
  sort.Slice(old.Processors, func(i, j int) bool {return *old.Processors[i].Name < *old.Processors[j].Name })

  for i:=0; i<len(cur.Processors); i++ {
    iseq := cur.Processors[i].Equals(old.Processors[i])
    if !iseq {
      println("here 7")
      return false
    }
  }

  if old.Disks == nil {
    println("here 8")
    return false
  }

  if len(cur.Disks) != len(old.Disks) {
    println("here 9")
    return false
  }

  sort.Slice(cur.Disks, func(i, j int) bool {return *cur.Disks[i].Name < *cur.Disks[j].Name })
  sort.Slice(old.Disks, func(i, j int) bool {return *old.Disks[i].Name < *old.Disks[j].Name })

  for i:=0; i<len(cur.Disks); i++ {
    iseq := cur.Disks[i].Equals(old.Disks[i])
    if !iseq {
      println("here 10")
      return false
    }
  }

  if old.Drives == nil {
    println("here 11")
    return false
  }

  if len(cur.Drives) != len(old.Drives) {
    println("here 12")
    return false
  }

  sort.Slice(cur.Drives, func(i, j int) bool {return *cur.Drives[i].Name < *cur.Drives[j].Name })
  sort.Slice(old.Drives, func(i, j int) bool {return *old.Drives[i].Name < *old.Drives[j].Name })

  for i:=0; i<len(cur.Drives); i++ {
    iseq := cur.Drives[i].Equals(old.Drives[i])
    if !iseq {
      println("here 13")
      return false
    }
  }

  if old.Nics == nil {
    println("here 14")
    return false
  }

  if len(cur.Nics) != len(old.Nics) {
    println("here 15")
    return false
  }

  sort.Slice(cur.Nics, func(i, j int) bool {return *cur.Nics[i].Name < *cur.Nics[j].Name })
  sort.Slice(old.Nics, func(i, j int) bool {return *old.Nics[i].Name < *old.Nics[j].Name })

  for i:=0; i<len(cur.Nics); i++ {
    iseq := cur.Nics[i].Equals(old.Nics[i])
    if !iseq {
      println("here 16")
      return false
    }
  }

  if old.NicConfigs == nil {
    println("here 17")
    return false
  }

  if len(cur.NicConfigs) != len(old.NicConfigs) {
    println("here 18")
    return false
  }

  sort.Slice(cur.NicConfigs, func(i, j int) bool {return *cur.NicConfigs[i].Description < *cur.NicConfigs[j].Description })
  sort.Slice(old.NicConfigs, func(i, j int) bool {return *old.NicConfigs[i].Description < *old.NicConfigs[j].Description })

  for i:=0; i<len(cur.NicConfigs); i++ {
    iseq := cur.NicConfigs[i].Equals(old.NicConfigs[i])
    if !iseq {
      println("here 19")
      return false
    }
  }

  if old.Videos == nil {
    return false
  }

  if len(cur.Videos) != len(old.Videos) {
    return false
  }

  sort.Slice(cur.Videos, func(i, j int) bool {return *cur.Videos[i].Name < *cur.Videos[j].Name })
  sort.Slice(old.Videos, func(i, j int) bool {return *old.Videos[i].Name < *old.Videos[j].Name })

  for i:=0; i<len(cur.Videos); i++ {
    iseq := cur.Videos[i].Equals(old.Videos[i])
    if !iseq {
      return false
    }
  }
  

  if old.Memories == nil {
    println("here 14")
    return false
  }

  if len(cur.Memories) != len(old.Memories) {
    println("here 15")
    return false
  }

  sort.Slice(cur.Memories, func(i, j int) bool {return *cur.Memories[i].BankLabel < *cur.Memories[j].BankLabel })
  sort.Slice(old.Memories, func(i, j int) bool {return *old.Memories[i].BankLabel < *old.Memories[j].BankLabel })

  for i:=0; i<len(cur.Memories); i++ {
    iseq := cur.Memories[i].Equals(old.Memories[i])
    if !iseq {
      println("here 16")
      return false
    }
  }

  return true
}



func initDeviceData(p string) {

  old = new(esmodels.DeviceAssetType)
  cur = new(esmodels.DeviceAssetType)

  pwd, err := os.Getwd()
  if err != nil {
    println(err)
    return
  }

  datadir = fmt.Sprintf("%s%c%s%c%s", pwd, filepath.Separator, p, filepath.Separator, "device.json")
  println("datadir :", datadir)

  buf := utils.GetJSONContents(datadir)

  if len(buf) <= 0 {
    return
  }
 
  err = json.Unmarshal(buf, old)
  println("$$$$$$$$$$ initialize old nic data:", len(old.Nics))
  return
}

func writeDeviceData() {
  f, err := os.Create(datadir)
  if err != nil {
    println("device create error:", err)
    return 
  }

  defer f.Close()

  bctn, _ := json.Marshal(old)

  f.WriteString(string(bctn))
  f.Sync()
  println("****************** device data write")
  return
}
