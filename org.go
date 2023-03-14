package messages

import (
)

// Information about Organization
type OrgInfo struct {
  Name          string          `json:"name"`
  Logo          string          `json:"logo"`
}

func NewOrgInfo() *OrgInfo {
  return &OrgInfo{}
}
