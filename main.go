package main

// import (
// 	"os"

// 	"github.com/elastic/beats/libbeat/cmd/instance"
// 	"github.com/elastic/beats/metricbeat/beater"

// 	// Comment out the following line to exclude all official metricbeat modules and metricsets
// 	_ "github.com/elastic/beats/metricbeat/include"

// 	// Make sure all your modules and metricsets are linked in this file
// 	_ "github.com/ssokssok/metricbeat/include"
// )

// var settings = instance.Settings{
// 	Name: "metricbeat",
// }

// func main() {
// 	if err := instance.Run(settings, beater.DefaultCreator()); err != nil {
// 		os.Exit(1)
// 	}
// }
import (
  "fmt"
  "os"
  "strings"

  "github.com/elastic/beats/libbeat/beat"
  "github.com/elastic/beats/libbeat/cfgfile"
  "github.com/elastic/beats/libbeat/cmd"
  xpackcmd "github.com/elastic/beats/x-pack/libbeat/cmd"

  //  "github.com/elastic/beats/libbeat/cmd/instance"
  "github.com/elastic/beats/metricbeat/beater"

  // Comment out the following line to exclude all official metricbeat modules and metricsets
  _ "github.com/elastic/beats/metricbeat/include"

  // Make sure all your modules and metricsets are linked in this file
  _ "github.com/ssokssok/metricbeat/include"
)

var Name = "metricbeat"

var RootCmd = cmd.GenRootCmd(Name, "", beater.DefaultCreator())

func main() {

    RootCmd.AddCommand(cmd.GenModulesCmd(Name, "", buildModulesManager))
    
    //RootCmd.AddCommand(xpackcmd.genEnrollCmd(name, ""))
    xpackcmd.AddXPack(RootCmd, Name)

    if err := RootCmd.Execute(); err != nil {
        os.Exit(1)
    }
    /*
    if err := instance.Run(settings, beater.DefaultCreator()); err != nil {
        os.Exit(1)
    }
    */
}


func buildModulesManager(beat *beat.Beat) (cmd.ModulesManager, error) {
    config := beat.BeatConfig

    glob, err := config.String("config.modules.path", -1)
    if err != nil {
        return nil, fmt.Errorf("modules management requires 'metricbeat.config.modules.path' setting")
    }

    if !strings.HasSuffix(glob, "*.yml") {
        return nil, fmt.Errorf("wrong settings for config.modules.path, it is expected to end with *.yml. Got: %s", glob)
    }

    modulesManager, err := cfgfile.NewGlobManager(glob, ".yml", ".disabled")
    if err != nil {
        return nil, fmt.Errorf("initialization error: %v", err)
    }
    return modulesManager, nil
}
