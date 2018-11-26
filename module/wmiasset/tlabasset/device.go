package tlabasset

import (
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getDevice() *esmodels.DeviceAssetType {

  dvc := new(esmodels.DeviceAssetType)

  // SystemType
  sys, errsys := getSystem()
  if errsys == nil {
    dvc.SystemType = sys
  }

  // PcBiosType
  bios, errbios := getBios()
  if errbios == nil {
    dvc.Bios = bios
  }

  // ProcessorType
  prca, errprc := getProcessors()
  if errprc == nil {
    dvc.Processors = prca
  }

  // DiskType
  diska, errdisk := getDisks()

  if errdisk == nil {
    dvc.Disks = diska
  }

  // DriveType
  drva, errdrv := getDrives()
  if errdrv == nil {
    dvc.Drives = drva
  }

  // NicType
  nica, errnic := getNics()
  if errnic == nil {
    dvc.Nics = nica
  }

  // NwConfigType
  nwcfga, errnwcfg := getNicConfigs()
  if errnwcfg == nil {
    dvc.NicConfigs = nwcfga
  }

  return dvc
}

