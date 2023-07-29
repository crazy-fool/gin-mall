package sid

import (
	"gin-mall/pkg/helper/convert"
	"gin-mall/pkg/log"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

type Sid struct {
	sf *sonyflake.Sonyflake
}

func init() {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	sid = &Sid{sf}
	log.GetLog().Info("=============初始化ssid================")
}

var sid *Sid

func GetSid() *Sid {
	return sid
}
func (s Sid) GenString() (string, error) {
	id, err := s.sf.NextID()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate sonyflake ID")
	}
	return convert.IntToBase62(int(id)), nil
}
func (s Sid) GenUint64() (uint64, error) {
	return s.sf.NextID()
}
