package patch

import (
	"github.com/elastic/beats/libbeat/common"
  "github.com/elastic/beats/libbeat/common/cfgwarn"
  "github.com/elastic/beats/libbeat/logp"
  "github.com/elastic/beats/libbeat/paths"  
	"github.com/elastic/beats/metricbeat/mb"
  
  //"bitbucket.org/truslab/pcon/servers/common/esmodels"
)

var (
  isInit = true
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("rsoasset", "patch", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
  //*esmodels.PatchType
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The rsoasset patch metricset is experimental.")

	config := struct{
    DataDir    string   `config:"datadir"`
  }{
    DataDir: "data",
  }
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

  logp.Info("##################### DataDir: %s, path: %s", paths.Paths.Data, config.DataDir)
  if isInit == true {
    isInit = false
    initPatchData(paths.Paths.Data, config.DataDir)
  }

	return &MetricSet{
		BaseMetricSet: base,
    // PatchType: new(esmodels.PatchType),
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {
  
  list, err := getPatchAssets()
  if err != nil {
    return
  }

  if len(list) == 0 {
    println("$$$$$$$$$$$$$$$$$$$$$$$$$$$ no change")
    return
  }

  for _, itm := range list {
    report.Event(mb.Event{
      MetricSetFields: common.MapStr{
        "hotfixid": itm.HotFixID,
        "description": itm.Description,
        "caption": itm.Caption,
        "installedon": itm.InstalledOn,
      },
    })
  }
}
