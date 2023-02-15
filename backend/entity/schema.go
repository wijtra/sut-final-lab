package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid: "required~ต้องไม่เป็นค่าว่าง"` // ต้องไม่เป็นค่าว่าง
	Email      string
	CustomerID string `"matches(L|M|H)~รหัสลูกค้าขึนต้นด้วย L หรือ M หรือ H, maxstringlength(7)~กรอกได้สูงสุด7ตัว"` // รหัสลูกค้าขึ8นต้นด้วย L หรือ M หรือ H แล้วตามด้วยตัวเลขจํานวน 7 ตัว
}
