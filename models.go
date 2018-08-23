package mail

import (
	"github.com/moisespsena-go/aorm"
)

type Mail struct {
	aorm.AuditedModel
	Address string `gorm:"size:255"`
	Note    string `gorm:"size:255"`
}

func (p *Mail) Stringify() (s string) {
	s += p.Address
	if p.Note != "" {
		s += " [" + p.Note + "]"
	}
	return
}
