package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/wijtra/sut-final-lab/entity"
)
// ตรวจสอบทะเบียนรถแล้วต้องเจอ Error
func TestRegistrationNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	cutomer := Customer{
		Name "", // ต้องไม่เป็นค่าว่าง
		Email "wsgniy@gmail.com",
		CustomerID "M1234567",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(vehicle)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ต้องไม่เป็นค่าว่าง"))
}
