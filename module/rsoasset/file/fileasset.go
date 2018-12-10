package file

import (
  "fmt"
  "strings"
  "os"
  "path/filepath"
  "encoding/json"

  "github.com/ssokssok/metricbeat/module/rsoasset/utils"
  "bitbucket.org/realsighton/rso/servers/common/esmodels"
)

var (
  datadir string
  old []*esmodels.FileType 
)

func initFileData(p string) {
 
  pwd, err := os.Getwd()
  if err != nil {
    println(err)
    return
  }

  datadir = fmt.Sprintf("%s%c%s%c%s", pwd, filepath.Separator, p, filepath.Separator, "file.json")
  println("datadir :", datadir)

  buf := utils.GetJSONContents(datadir)

  if len(buf) <= 0 {
    return
  }

  old = make([]*esmodels.FileType, 0)
  
  err = json.Unmarshal(buf, &old)
  println("$$$$$$$$$$ initialize old length:", len(old))
  return
}

func writeFileData() {
  f, err := os.Create(datadir)
  if err != nil {
    println("file create error:", err)
    return 
  }

  defer f.Close()

  bctn, _ := json.Marshal(old)

  f.WriteString(string(bctn))
  f.Sync()
  println("****************** filedata write")
  return
}

func getFileAssets() ([]*esmodels.FileType, error) {
  cur, err := getEsModelFileType()
  if err != nil {
    println(err)
    return nil, err
  }

  ma := make([]*esmodels.FileType, 0) 


  for _, v := range cur {
    f := findExistList(v)
    if !f {
      ma = append(ma, v)
    }
  }

  for _, m := range ma {
    old = append(old, m)
  }
  //old = ma 

  writeFileData()

  return ma, nil
}


func findExistList(v *esmodels.FileType) bool {
  
  if old == nil {
    old = make([]*esmodels.FileType, 0)
    return false
  }

  if len(old) == 0 {
    return false
  }

  for _, ov := range old {
    if v.Name != nil && ov.Name != nil {
      if v.FileVersion != nil && ov.FileVersion != nil {
        if *v.Name == *ov.Name && *v.FileVersion == *ov.FileVersion {
          return true
        }    
      }
    }
  }
  return false
}


func getEsModelFileType() ([]*esmodels.FileType, error) {

  ma := make([]*esmodels.FileType, 0)

  prca, errprc := getProcAssets()
  if errprc != nil {
    return nil, errprc
  }

  println("Length of Process : ", len(prca))
  for _, prc := range prca {

    if prc.ExecutablePath == nil && prc.CommandLine == nil {
      continue
    }

    var path string 
    
    if prc.ExecutablePath != nil {
      path = *prc.ExecutablePath
    }
   
    if prc.ExecutablePath == nil {
      path = *prc.CommandLine
    }

    if excludeProcess(*prc.Name, path) {
      continue
    }

    m := new(esmodels.FileType)

    err := getVersionInfo(path, m ) 
    if err != nil {
      continue
    }

    ma = append(ma, m)
  }

  println("################### last file count", len(ma))
  return ma, nil  
}


func excludeProcess(nm string, path string) bool {
  
  lpath := strings.ToLower(path) 

  exclprefix := []string{"c:\\windows"}

  for _, pfx := range exclprefix {
    if strings.HasPrefix(lpath, pfx) {
      return true
    }
  }

  exclpath := []string {"microsoft","epson","google","kakao"}

  for _, pth := range exclpath {
    if strings.Contains(lpath, pth) {
      return true
    } 
  }

  exclptn := []string{ "conhost.exe",
  "powershell.exe",
  "dptf_helper.exe",
  "sihost.exe",
  "svchost.exe",
  "igfxEM.exe",
  "taskhostw.exe",
  "explorer.exe",
  "dllhost.exe",
  "searchui.exe",
  "msascuil.exe",
  "onedrive.exe",
  "slack.exe",
  "kakaotalk.exe",
  "jusched.exe",
  "chrome.exe",
  "gocode.exe",
  "sublime_text.exe", 
  "plugin_host.exe",
  "metricbeat.exe",
  }

  lnm := strings.ToLower(nm)
  
  for _, ptn := range exclptn {
    if lnm == ptn {
      return true
    }
  }

  return false
}