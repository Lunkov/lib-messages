package messages

import ( 
  "sync"
)

type ReqGetBalance struct {
  Address              string      `gorm:"column:address;type:string;primary_key"`
  Sign                 []byte      `gorm:"column:sign"`
  PublicKey            []byte      `gorm:"column:public_key"`
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
