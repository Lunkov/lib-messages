package messages

import ( 
  "sync"
)

type Balance struct {
  Address              string
  Coin                 string
  Balance              uint64
  UnconfirmedBalance   uint64
  TotalReceived        uint64
  TotalSent            uint64
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
