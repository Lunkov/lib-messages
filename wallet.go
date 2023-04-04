package messages

import ( 
  "sync"

  "bytes"
  
  "encoding/gob"
  "encoding/base64"
)

type ReqGetBalance struct {
  Address              string      `json:"address"      gorm:"column:address;type:string;primary_key"`
  Sign                 []byte      `json:"sign"         gorm:"column:sign"`
  PublicKey            []byte      `json:"public_key"   gorm:"column:public_key"`
}

func NewReqGetBalance() *ReqGetBalance {
  return &ReqGetBalance{}
}

func (t *ReqGetBalance) Pack() string {
  var buff bytes.Buffer
  encoder := gob.NewEncoder(&buff)
  encoder.Encode(t)
  return base64.StdEncoding.EncodeToString(buff.Bytes())
}

func (t *ReqGetBalance) Unpack(msg string) bool {
  input, err := base64.StdEncoding.DecodeString(msg)
  if err != nil {
    return false
  }
  buf := bytes.NewBuffer(input)
  decoder := gob.NewDecoder(buf)
  err = decoder.Decode(t)
  if err != nil {
    return false
  }
  return true
}

type Balance struct {
  Address              string      `gorm:"column:address;type:string;primary_key"`
  Coin                 string      `gorm:"column:coin;type:string"`
  Balance              uint64      `gorm:"column:balance"`
  UnconfirmedBalance   uint64      `gorm:"column:unconfirmed_balance"`
  TotalReceived        uint64      `gorm:"column:total_received"`
  TotalSent            uint64      `gorm:"column:total_sent"`
  
  //LastTransaction      string      `gorm:"column:last_transaction"`
  //UpdatedAt            time.Time   `gorm:"column:updated_at;type:timestamp with time zone"`
  //Hash                 []byte      `gorm:"column:hash"`
}

func NewBalance() *Balance {
  return &Balance{}
}

func (b *Balance) Pack() string {
  var buff bytes.Buffer
  encoder := gob.NewEncoder(&buff)
  encoder.Encode(b)
  return base64.StdEncoding.EncodeToString(buff.Bytes())
}

func (b *Balance) Unpack(msg string) bool {
  input, err := base64.StdEncoding.DecodeString(msg)
  if err != nil {
    return false
  }
  buf := bytes.NewBuffer(input)
  decoder := gob.NewDecoder(buf)
  err = decoder.Decode(b)
  if err != nil {
    return false
  }
  return true
}

type Balances struct {
  a     []*Balance
  mu    sync.RWMutex
}

func NewBalances() *Balances {
  return &Balances{a: make([]*Balance, 0)}
}

func (b *Balances) Add(i *Balance) {
  b.mu.Lock()
  b.a = append(b.a, i)
  b.mu.Unlock()
}

func (b *Balances) GetBalanses() ([]*Balance) {
  return b.a
}
