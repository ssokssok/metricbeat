package file

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
  "github.com/elastic/beats/metricbeat/mb"

//  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("tlabasset", "file", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
  mb.BaseMetricSet
  Names *string   
  Size *int64   
  FileDescription *string   
  OriginalFilename *string  
  FileVersion *string   
  ProductName *string   
  ProductVersion *string 
  CompanyName *string
  LegalCopyright *string
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Experimental("The tlabasset file metricset is experimental.")

	config := struct{}{}
	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
    BaseMetricSet: base,
    Names: new(string),
    Size: new(int64),
    FileDescription: new(string),
    OriginalFilename: new(string),
    FileVersion: new(string),
    ProductName: new(string),
    ProductVersion: new(string),
    CompanyName: new(string),
    LegalCopyright: new(string),
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) {

  // path := `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`

  // itm := new(esmodels.FileType)

  // err := getVersionInfo(path, itm ) 
  // if err != nil {
  //   return
  // }

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
  
}
