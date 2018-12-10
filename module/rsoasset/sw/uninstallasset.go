package sw

import (
  "os/exec"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/realsighton/rso/servers/common/esmodels"
)

var (
  oldun []*esmodels.SwAssetType 
)


// UninstallSw is ...
type UninstallSw struct {
  DisplayName *string `json:"DisplayName,omitempty"`
  DispalyVersion *string `json:"DisplayVersion,omitempty"`
  Publisher *string `json:"Publisher,omitempty"`
  UninstallString *string `json:"UninstallString,omitempty"`
  InstallDate *json.Number `json:"InstallDate,Number,omitempty"`
  EstimatedSize *int64 `json:"EstimatedSize,omitempty"`
}

func getSwUninstallAssets() ([]*esmodels.SwAssetType, error) {
  cur := make([]*esmodels.SwAssetType, 0)

  var err error
  err = getEsModelSwUninstallHKCU(&cur)
  if err != nil {
    println(err)
    return nil, err
  }

  err = getEsModelSwUninstallHKLM1(&cur)
  if err != nil {
    println(err)
    return nil, err
  }
  
  err = getEsModelSwUninstallHKLM2(&cur)
  if err != nil {
    println(err)
    return nil, err
  }

  ma := make([]*esmodels.SwAssetType, 0)

  for _, v := range cur {
    f := findExistListUninstall(v)
    if !f {
      ma = append(ma, v)
    }
  }

  for _, m := range ma {
    oldun = append(oldun, m)
  }

  return ma, nil
}


func findExistListUninstall(v *esmodels.SwAssetType) bool {
  
  if oldun == nil {
    oldun = make([]*esmodels.SwAssetType, 0)
    return false
  }

  if len(oldun) == 0 {
    return false
  }

  for _, ov := range oldun {
    if ov.Equals(v) {
      return true
    }
  }

  return false
}

func getEsModelSwUninstallHKCU(ma *[]*esmodels.SwAssetType) error {

  qry := `Powershell /command "Get-ItemProperty HKCU:\Software\Microsoft\Windows\CurrentVersion\Uninstall\* | Select-Object DisplayName,DisplayVersion,Publisher,UninstallString,InstallDate,EstimatedSize | Convertto-Json | out-file UninstallHKCU.json -encoding UTF8"`
  fn := `UninstallHKCU.json`
  
  execcmd := "PowerShell"
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return err
  }

  buf := utils.GetContents(fn) 
  println(string(buf))

  uma := make([]*UninstallSw, 0)
  um := new(UninstallSw)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &uma)
  } else {
    err = json.Unmarshal(buf, um)
    uma = append(uma, um)
  }

  if err != nil {
    println(err.Error())
    return err
  }

  for _, ui := range uma {
    m := new(esmodels.SwAssetType)
    m.Name = ui.DisplayName 
    m.Version = ui.DispalyVersion
    m.Vendor = ui.Publisher
    if ui.InstallDate != nil {
      tmp := string(*ui.InstallDate)
      m.InstallDate = &tmp
    }
    m.Size = ui.EstimatedSize
    *ma = append(*ma, m)
  }

  return nil  
}


func getEsModelSwUninstallHKLM1(ma *[]*esmodels.SwAssetType) error {

  qry := `Powershell /command "Get-ItemProperty HKLM:\Software\Wow6432Node\Microsoft\Windows\CurrentVersion\Uninstall\* | Select-Object DisplayName, DisplayVersion, Publisher, UninstallString,InstallDate,EstimatedSize | ConvertTo-Json | out-file UninstallHKLM1.json -encoding UTF8"`
  fn := `UninstallHKLM1.json`
  
  execcmd := "PowerShell"
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return err
  }

  buf := utils.GetContents(fn) 
  println(string(buf))

  uma := make([]*UninstallSw, 0)
  um := new(UninstallSw)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &uma)
  } else {
    err = json.Unmarshal(buf, um)
    uma = append(uma, um)
  }

  if err != nil {
    println(err.Error())
    return err
  }

  for _, ui := range uma {
    m := new(esmodels.SwAssetType)
    m.Name = ui.DisplayName 
    m.Version = ui.DispalyVersion
    m.Vendor = ui.Publisher
    if ui.InstallDate != nil {
      tmp := string(*ui.InstallDate)
      m.InstallDate = &tmp
    }
    m.Size = ui.EstimatedSize
    *ma = append(*ma, m)
  }

  return nil  
}

func getEsModelSwUninstallHKLM2(ma *[]*esmodels.SwAssetType) error {

  qry := `Powershell /command "Get-ItemProperty HKLM:\Software\Microsoft\Windows\CurrentVersion\Uninstall\* | Select-Object DisplayName,DisplayVersion,Publisher,UninstallString,InstallDate,EstimatedSize | Convertto-Json | out-file UninstallHKLM2.json -encoding UTF8"`
  fn := `UninstallHKLM2.json`
  
  execcmd := "PowerShell"
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return err
  }

  buf := utils.GetContents(fn) 
  println(string(buf))

  uma := make([]*UninstallSw, 0)
  um := new(UninstallSw)

  if buf[0] == '[' {
    err = json.Unmarshal(buf, &uma)
  } else {
    err = json.Unmarshal(buf, um)
    uma = append(uma, um)
  }

  if err != nil {
    println(err.Error())
    return err
  }

  for _, ui := range uma {
    m := new(esmodels.SwAssetType)
    m.Name = ui.DisplayName 
    m.Version = ui.DispalyVersion
    m.Vendor = ui.Publisher
    if ui.InstallDate != nil {
      tmp := string(*ui.InstallDate)
      m.InstallDate = &tmp
    }
    m.Size = ui.EstimatedSize
    *ma = append(*ma, m)
  }

  return nil  
}