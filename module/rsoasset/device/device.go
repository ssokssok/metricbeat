package device

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
  "github.com/elastic/beats/metricbeat/mb"
  
  "bitbucket.org/realsighton/rso/servers/common/esmodels"  
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

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
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
		},
  })
}
