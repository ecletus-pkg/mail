package mail

import (
	"github.com/aghape/db/common"
	"github.com/aghape/db/common/mixins"
)

type QorMail struct {
	common.Model
	mixins.CreationUpdationMixin
	Address string `gorm:"size:255"`
	Note    string `gorm:"size:255"`
}

func (p *QorMail) Stringify() (s string) {
	s += p.Address
	if p.Note != "" {
		s += " [" + p.Note + "]"
	}
	return
}
