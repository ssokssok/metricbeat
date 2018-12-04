package file

import (
  "fmt"
  "os/exec"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

// VersionInfo is ...
type VersionInfo struct {
  CompanyName *string `json:"CompanyName,omitempty"`
  FileDescription *string `json:"FileDescription,omitempty"`
  OriginalFilename *string `json:"OriginalFilename,omitempty"`
  ProductName *string `json:"ProductName,omitempty"`
  ProductVersion *string `json:"ProductVersion,omitempty"`
  LegalCopyright *string `json:"LegalCopyright,omitempty"`
  FileVersion *string `json:"FileVersion,omitempty"`
}

// RSOFileInfoType is ...
type RSOFileInfoType struct {
  Name *string `json:"Name,omitempty"`
  FullName *string `json:"FullName,omitempty"`
  Size  *int64 `json:"Size,omitempty"`
  CreationTime *string `json:"CreationTime,omitempty"`
  Version  *VersionInfo `json:"VersionInfo,omitempty"`
}

func getVersionInfo(path string, ft *esmodels.FileType) error {

  qry := `Get-ItemProperty -Path '%s' | select-object name, fullname,@{Label='Size';Expression={$_.length}}, versioninfo, @{Label='CreationTime';Expression={Get-Date  $_.creationtime -Format 'yyyy-MM-ddThh:mm:sszzz' } } | convertto-json | out-file rsoFileInfo.json -encoding UTF8`
  //qry := `Powershell.exe "Get-ItemProperty -Path '%s' | select-object name, fullname, length, versioninfo, @{Label='creationtime';Expression={Get-Date  $_.creationtime -Format 'yyyy-MM-ddThh:mm:sszzz' } } | convertto-json | out-file rsoFileInfo.json -encoding UTF8"`
  //qry := `Get-ItemProperty -Path '%s' | select-object name, fullname, length, versioninfo, @{Label='CreationTime';Expression={Get-Date  $_.creationtime -Format 'yyyy-MM-ddThh:mm:sszzz' } } | convertto-json | out-file tlabFileInfo.json -encoding UTF8"`
  qry = fmt.Sprintf(qry, path)
  fn := `rsoFileInfo.json`

  execcmd := "Powershell.exe"  // PowerShell
  println("################# :", qry)
  cmd := exec.Command(execcmd, qry)
  _, err := cmd.CombinedOutput()
  if err != nil {
    return err
  }

  buf := utils.GetContents(fn) 
  println(string(buf))

  m := new(RSOFileInfoType)

  err = json.Unmarshal(buf, m)

  if err != nil {
    println("@@@@@@@@@@@@@@@@ error :", err.Error())
    return err
  }

  ft.Name = m.Name 
  ft.Size = m.Size 
  ft.FileDescription = m.Version.FileDescription
  ft.OriginalFilename = m.Version.OriginalFilename
  ft.FileVersion = m.Version.FileVersion
  ft.ProductName = m.Version.ProductName
  ft.ProductVersion = m.Version.ProductVersion
  ft.CompanyName = m.Version.CompanyName
  ft.LegalCopyright = m.Version.LegalCopyright
  return nil  
}

