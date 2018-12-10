package os

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
  old  *esmodels.OsAssetType
  cur  *esmodels.OsAssetType
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("rsoasset", "os", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
  *esmodels.OsAssetType
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The rsoasset os metricset is experimental.")

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
    initOsData(config.DataDir)
  }

	return &MetricSet{
		BaseMetricSet: base,
    OsAssetType: new(esmodels.OsAssetType),
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {

  var err error

  m.Os, err = getOs()
  if err != nil {
    return
  }

  m.Timezone, err = getTZ()
  if err != nil {
    return
  }

  m.Shares, err = getShares()
  if err != nil {
    return
  }

  m.UserAccounts, err = getUserAccounts()
  if err != nil {
    return
  }
  
  cur.Os = m.Os
  cur.Timezone = m.Timezone 
  cur.Shares = m.Shares 
  cur.UserAccounts = m.UserAccounts 

  isEq := checkEqualOS()
  
  if isEq {
    println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& no changed")
    writeOsData()
    return 
  }

  old = cur

  writeOsData()

	report.Event(mb.Event{
		MetricSetFields: common.MapStr{
      "os": m.Os,
      "timezone": m.Timezone,
      "shares": m.Shares,
      "useraccounts": m.UserAccounts,
		},
	})
}


func checkEqualOS() bool {
  var isEq bool

  if old.Os == nil {
    println("hear 1")
    return false
  }

  isEq = cur.Os.Equals(old.Os)
  if !isEq {
    println("hear 2")
    return false
  }

  if old.Timezone == nil {
    println("hear 3")
    return false
  }

  isEq = cur.Timezone.Equals(old.Timezone)
  if !isEq {
    println("hear 4")
    return false
  }

  if cur.Shares != nil && old.Shares == nil {
    println("here 5")
    return false
  }

  if len(cur.Shares) != len(old.Shares) {
    println("here 6")
    return false
  }

  sort.Slice(cur.Shares, func(i, j int) bool {return *cur.Shares[i].Name < *cur.Shares[j].Name })
  sort.Slice(old.Shares, func(i, j int) bool {return *old.Shares[i].Name < *old.Shares[j].Name })

  for i:=0; i<len(cur.Shares); i++ {
    iseq := cur.Shares[i].Equals(old.Shares[i])
    if !iseq {
      println("here 7")
      return false
    }
  }

  if cur.UserAccounts != nil && old.UserAccounts == nil {
    println("here 8")
    return false
  }

  if len(cur.UserAccounts) != len(old.UserAccounts) {
    println("here 9")
    return false
  }

  sort.Slice(cur.UserAccounts, func(i, j int) bool {return *cur.UserAccounts[i].Name < *cur.UserAccounts[j].Name })
  sort.Slice(old.UserAccounts, func(i, j int) bool {return *old.UserAccounts[i].Name < *old.UserAccounts[j].Name })

  for i:=0; i<len(cur.UserAccounts); i++ {
    iseq := cur.UserAccounts[i].Equals(old.UserAccounts[i])
    if !iseq {
      println("here 10")
      return false
    }
  }

  return true
}


func initOsData(p string) {

  old = new(esmodels.OsAssetType)
  cur = new(esmodels.OsAssetType)

  pwd, err := os.Getwd()
  if err != nil {
    println(err)
    return
  }

  datadir = fmt.Sprintf("%s%c%s%c%s", pwd, filepath.Separator, p, filepath.Separator, "os.json")
  println("datadir :", datadir)

  buf := utils.GetJSONContents(datadir)

  if len(buf) <= 0 {
    return
  }
 
  err = json.Unmarshal(buf, old)
  println("$$$$$$$$$$ initialize old share data:", len(old.Shares))
  return
}

func writeOsData() {
  f, err := os.Create(datadir)
  if err != nil {
    println("os create error:", err)
    return 
  }

  defer f.Close()

  bctn, _ := json.Marshal(old)

  f.WriteString(string(bctn))
  f.Sync()
  println("****************** os data write")
  return
}
