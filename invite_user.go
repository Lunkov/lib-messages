package messages

import (
  "time"
  "bytes"
  "crypto/sha512"
  "encoding/gob"
  "encoding/base64"
)

// Invite User to Organization
type DBUserInvite struct {
  InviteToken   string          `json:"token"`
  Login         string          `json:"login"`
  Role          string          `json:"role"`
  EMail         string          `json:"email" gorm:"index:idx_email,unique"`
  UpdatedAt     time.Time
  OrgSign     []byte
}

type ReqUserInviteFromOrg struct {
  NodeUrl       string          `json:"url"`
  InviteToken   string          `json:"token"`
  Login         string          `json:"login"`
  EMail         string          `json:"email"`
  UpdatedAt     time.Time       `json:"updated_at"`
  Sign          []byte          `json:"sign"`
}

type ReqUserRegisterToOrg struct {
  NodeUrl       string          `json:"url"`
  InviteToken   string          `json:"token"`
  Login         string          `json:"login"`
  EMail         string          `json:"email"`

  DisplayName   string          `json:"displayName"`
  FirstName     string          `json:"first_name"`
  MiddleName    string          `json:"middle_name"`
  LastName      string          `json:"last_name"`

  Country       string          `json:"country"`
  Locality      string          `json:"locality"`
  OrgUnit       string          `json:"unit"`

  PubKey        []byte          `json:"pubkey"`
  Sign          []byte          `json:"sign"`
}

/*
 * DBUserInvite
 * 
 * 
 */
 
func NewDBUserInvite() *DBUserInvite {
  return &DBUserInvite{}
}

func (i *DBUserInvite) Init(token string, login string, email string) {
  i.InviteToken = token
  i.Login = login
  i.EMail = email
  i.UpdatedAt = time.Now()
}

func (i *DBUserInvite) Hash() []byte {
  sha_512 := sha512.New()
  sha_512.Write([]byte(i.InviteToken + i.Login + i.EMail + i.UpdatedAt.String()))
  return sha_512.Sum(nil)
}

/*
 * ReqUserInviteFromOrg
 * 
 * 
 */

func NewReqUserInviteFromOrg() (*ReqUserInviteFromOrg) {
  return &ReqUserInviteFromOrg{}
}

func (i *ReqUserInviteFromOrg) Init(nodeUrl string, dbui *DBUserInvite) {
  i.NodeUrl = nodeUrl
  i.InviteToken = dbui.InviteToken
  i.Login = dbui.Login 
  i.EMail = dbui.EMail
  i.UpdatedAt = dbui.UpdatedAt
}

func (i *ReqUserInviteFromOrg) Hash() []byte {
  sha_512 := sha512.New()
  sha_512.Write([]byte(i.InviteToken + i.NodeUrl + i.Login + i.EMail + i.UpdatedAt.String()))
  return sha_512.Sum(nil)
}

func (i *ReqUserInviteFromOrg) Pack() string {
  var buff bytes.Buffer
  encoder := gob.NewEncoder(&buff)
  encoder.Encode(i)
  return base64.StdEncoding.EncodeToString(buff.Bytes())
}

func (i *ReqUserInviteFromOrg) Unpack(msg string) bool {
  input, err := base64.StdEncoding.DecodeString(msg)
  if err != nil {
    return false
  }
  buf := bytes.NewBuffer(input)
  decoder := gob.NewDecoder(buf)
  err = decoder.Decode(i)
  if err != nil {
    return false
  }
  return true
}

/*
 * ReqUserRegisterToOrg
 * 
 * 
 */

func NewReqUserRegisterToOrg() (*ReqUserRegisterToOrg) {
  return &ReqUserRegisterToOrg{}
}

func (i *ReqUserRegisterToOrg) Hash() []byte {
  sha_512 := sha512.New()
  sha_512.Write([]byte(i.InviteToken + i.NodeUrl + i.Login + i.EMail + i.DisplayName + 
                       i.FirstName + i.MiddleName + i.LastName + 
                       i.Country + i.Locality +
                       string(i.PubKey)))
  return sha_512.Sum(nil)
}
