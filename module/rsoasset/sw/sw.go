package sw

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/metricbeat/mb"
)

var (
  isInit = true
)
// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("rsoasset", "sw", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The rsoasset sw metricset is experimental.")

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
    initSWData(config.DataDir)
  }


	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {
  list, err := getSwAssets()
  if err != nil {
    return
  }

  for _, itm := range list {
    report.Event(mb.Event{
      MetricSetFields: common.MapStr{
        "name": itm.Name,
        "version": itm.Version,
        "productid": itm.ProductID,
        "vendor": itm.Vendor,
        "language": itm.Language,
        "packagecode": itm.PackageCode,
        "skunumber": itm.SKUNumber,
        "size": itm.Size,
        "identifyingnumber": itm.IdentifyingNumber,
        "installdate": itm.InstallDate,
      },
    })
  }

  println("#################### last sw count:", len(list))
}
