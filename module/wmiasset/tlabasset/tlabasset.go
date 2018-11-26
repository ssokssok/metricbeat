package tlabasset

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
	mb.Registry.MustAddMetricSet("wmiasset", "tlabasset", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
//  counter int
  //Win32_DiskDrive []*esmodels.DiskType
  Device  *esmodels.DeviceAssetType
  Printer *esmodels.PrinterAssetType
  Os      *esmodels.OsAssetType
  Sw      *esmodels.SwAssetType
  Patch   *esmodels.PatchType
  File    *esmodels.FileType
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The wmiasset tlabasset metricset is experimental.")

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet: base,
 //   counter:       1,
    Device:      new(esmodels.DeviceAssetType),
    Printer:     new(esmodels.PrinterAssetType),
    Os:          new(esmodels.OsAssetType),
    Sw:          new(esmodels.SwAssetType),
    Patch:       new(esmodels.PatchType),
    File:        new(esmodels.FileType),
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {

  dvc := getDevice()

	report.Event(mb.Event{
		MetricSetFields: common.MapStr{
 //     "counter": m.counter,
      "device": dvc,
		},
  })
//	m.counter++

  osa := getOsAsset()

  report.Event(mb.Event{
    MetricSetFields: common.MapStr{
      "os": osa,
    },
  })

  sws := getSws()
  if sws == nil {
    return
  }
  println("#### sws count:", len(sws))
  for _, sw := range sws {
    report.Event(mb.Event{
      MetricSetFields: common.MapStr{
        "sw": sw,
      },
    })
  }

  prts := getPrinters()
  if prts == nil {
    return
  }

  println("#### prts count:", len(prts))
  for _, prt := range prts {
    report.Event(mb.Event{
      MetricSetFields: common.MapStr{
        "printer": prt,
      },
    })
  }  

  ptchlst := getPatchs()
  if ptchlst == nil {
    return
  }

  println("#### ptchlst count:", len(ptchlst))
  for _, ptch := range ptchlst {
    report.Event(mb.Event{
      MetricSetFields: common.MapStr{
        "patch": ptch,
      },
    })
  }    
}
