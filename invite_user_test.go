package messages

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestInviteUser(t *testing.T) {
	
  dbui := NewDBUserInvite()
  dbui.Init("token", "user", "login@domain.org")
  
  dbui.Hash()
     
  uio := NewReqUserInviteFromOrg()
  uio.Init("ANO", dbui)
  assert.Equal(t, 64, len(uio.Hash()))
  
  msg := uio.Pack()
  assert.Equal(t, 244, len(msg))
  
  uio2 := NewReqUserInviteFromOrg()
  uio2.Unpack(msg)
  
  assert.Equal(t, *uio, *uio2)
}

func BenchmarkMutexMap(b *testing.B) {

  b.ResetTimer()
}
