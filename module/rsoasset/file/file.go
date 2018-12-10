package file

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
	mb.Registry.MustAddMetricSet("rsoasset", "file", New)
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
	cfgwarn.Experimental("The rsoasset file metricset is experimental.")

  //config := struct{}{}
  config := struct{
    MaxConn    int      `config:"maxconn" validate:"min=1"`
    RsoSvr     string   `config:"rsosvr"`
    Hosts      []string `config:"hosts"`
    DataDir    string   `config:"datadir"`
  }{
    MaxConn: 10,
    RsoSvr: "",
    Hosts: make([]string, 0),
    DataDir: "data",
  }
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

  println("###################### MaxConn", config.MaxConn, " RsoSvr:", config.RsoSvr)
  println(" Hosts:", config.Hosts)
  for _, v := range config.Hosts {
    println("      - host:", v)
  }

  println("##################### DataDir", config.DataDir)
  if isInit == true {
    isInit = false
    initFileData(config.DataDir)
  }

	return &MetricSet{
		BaseMetricSet: base,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {

  list, err  := getFileAssets()
  if err != nil {
    println(err.Error())
    return
  }

  for _, itm := range list {
    report.Event(mb.Event{
      MetricSetFields: common.MapStr{
        "name": itm.Name,
        "size": itm.Size,
        "file_description": itm.FileDescription,
        "original_filename": itm.OriginalFilename,
        "file_version": itm.FileVersion,
        "product_name": itm.ProductName,
        "product_version": itm.ProductVersion,
        "company_name": itm.CompanyName,
        "legal_copyright": itm.LegalCopyright,
      },
    })
  }

  println("###################### filecount:",len(list))

}
