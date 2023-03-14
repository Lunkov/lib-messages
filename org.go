package messages

import (
  "time"
  "bytes"
  "crypto/sha512"
  "encoding/gob"
  "encoding/base64"
)

// Information about Organization
type OrgInfo struct {
  Name          string          `json:"name"`
  Logo          string          `json:"logo"`
}

func NewOrgInfo() *OrgInfo {
  return &OrgInfo{}
}
