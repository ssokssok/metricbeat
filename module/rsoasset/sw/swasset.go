package sw

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "encoding/json"

  "github.com/elastic/beats/libbeat/logp"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/realsighton/rso/servers/common/esmodels"
)

var (
  datadir string
  old []*esmodels.SwAssetType 
)


func initSWData(pwd, p string) {
 
  // pwd, err := os.Getwd()
  // if err != nil {
  //   println(err)
  //   return
  // }

  // datadir = fmt.Sprintf("%s%c%s%c%s", pwd, filepath.Separator, p, filepath.Separator, "sw.json")
  datadir = fmt.Sprintf("%s%c%s", pwd, filepath.Separator, "sw.json")
  logp.Info("datadir : %s", datadir)

  buf := utils.GetJSONContents(datadir)

  if len(buf) <= 0 {
    return
  }

  old = make([]*esmodels.SwAssetType, 0)
  
  err := json.Unmarshal(buf, &old)
  if err != nil {
    logp.Warn("initialize get info err: %v", err)
    return 
  }
  logp.Info("$$$$$$$$$$ initialize old length: %d", len(old))
  return
}

func writeSWData() {
  f, err := os.Create(datadir)
  if err != nil {
    println("sw create error:", err)
    return 
  }

  defer f.Close()

  bctn, _ := json.Marshal(old)

  f.WriteString(string(bctn))
  f.Sync()
  println("****************** sw data write")
  return
}


func getSwAssets() ([]*esmodels.SwAssetType, error) {
  m, err := getEsModelSwAssetType()
  if err != nil {
    println(err)
    return nil, err
  }

  um, uerr := getSwUninstallAssets()
  if uerr != nil {
    return nil, uerr
  }

  if m != nil {
    for _, itm := range um {
      m = append(m, itm)
    }
    
    for _, mi := range m {
      old = append(old, mi)
    }
    writeSWData()
    return m, nil
  }

  for _, mi := range um {
    old = append(old, mi)
  }
  writeSWData()
  return um, nil
}

func getSwAssetType() (string, error) {

  m := new(esmodels.SwAssetType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return "", err
  }

  buf := utils.GetContents(fn) 

  ma := make([]*esmodels.SwAssetType, 0)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &ma)
  } else {
    err = json.Unmarshal(buf, m)
    ma = append(ma, m)
  }

  if err != nil {
    return "", err
  }

  s := ""

  if b, err := json.Marshal(ma); err == nil {
    s = string(b)
  } 

  return s, nil
}

func getEsModelSwAssetType() ([]*esmodels.SwAssetType, error) {

  m := new(esmodels.SwAssetType)

  qry, fn := m.GetPsQuery()
  
  execcmd := "PowerShell"
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return nil, err
  }

  buf := utils.GetContents(fn) 
  //println(string(buf))
  cur := make([]*esmodels.SwAssetType, 0)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &cur)
  } else {
    err = json.Unmarshal(buf, m)
    cur = append(cur, m)
  }

  if err != nil {
    println(err.Error())
    return nil, err
  }

  ma := make([]*esmodels.SwAssetType, 0)

  for _, v := range cur {
     f := findExistList(v)
     if !f {
       ma = append(ma, v)
     }
  }


  return ma, nil  
}

func findExistList(v *esmodels.SwAssetType) bool {
  
  if old == nil {
    old = make([]*esmodels.SwAssetType, 0)
    return false
  }

  if len(old) == 0 {
    return false
  }

  for _, ov := range old {
    if ov.Equals(v) {
      return true
    }
  }

  return false
}