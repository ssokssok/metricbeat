package tlabasset

import (
  "bitbucket.org/truslab/pcon/servers/common/esmodels"
)

func getPrinters() []*esmodels.PrinterAssetType {

  // PrinterAssetType
  list, err := getPrinterAssets()
  if err != nil {
    return nil
  }
  println("printer length: ", len(list))
  return list
}

