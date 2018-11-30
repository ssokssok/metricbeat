package os

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
  "github.com/elastic/beats/metricbeat/mb"

  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("tlabasset", "os", New)
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
	cfgwarn.Experimental("The tlabasset os metricset is experimental.")

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
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
  
	report.Event(mb.Event{
		MetricSetFields: common.MapStr{
      "os": m.Os,
      "timezone": m.Timezone,
      "shares": m.Shares,
      "useraccounts": m.UserAccounts,
		},
	})
}
