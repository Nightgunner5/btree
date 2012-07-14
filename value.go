package btree

type value_t interface {
	Value() interface{}
	Equals(other value_t) bool
	LessThan(other value_t) bool
}

type value_impl_Sortable struct {
	val Sortable
}
func (v *value_impl_Sortable) Value() interface{} {
	return v.val
}
func (v *value_impl_Sortable) Equals(other value_t) bool {
	o, ok := other.(*value_impl_Sortable)
	return ok && v.val.Equals(o.val)
}
func (v *value_impl_Sortable) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_Sortable)
	return ok && v.val.LessThan(o.val)
}


type value_impl_int struct {
	val int
}
func (v *value_impl_int) Value() interface{} {
	return v.val
}
func (v *value_impl_int) Equals(other value_t) bool {
	o, ok := other.(*value_impl_int)
	return ok && v.val == o.val
}
func (v *value_impl_int) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_int)
	return ok && v.val < o.val
}


type value_impl_uint struct {
	val uint
}
func (v *value_impl_uint) Value() interface{} {
	return v.val
}
func (v *value_impl_uint) Equals(other value_t) bool {
	o, ok := other.(*value_impl_uint)
	return ok && v.val == o.val
}
func (v *value_impl_uint) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_uint)
	return ok && v.val < o.val
}


type value_impl_int8 struct {
	val int8
}
func (v *value_impl_int8) Value() interface{} {
	return v.val
}
func (v *value_impl_int8) Equals(other value_t) bool {
	o, ok := other.(*value_impl_int8)
	return ok && v.val == o.val
}
func (v *value_impl_int8) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_int8)
	return ok && v.val < o.val
}


type value_impl_uint8 struct {
	val uint8
}
func (v *value_impl_uint8) Value() interface{} {
	return v.val
}
func (v *value_impl_uint8) Equals(other value_t) bool {
	o, ok := other.(*value_impl_uint8)
	return ok && v.val == o.val
}
func (v *value_impl_uint8) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_uint8)
	return ok && v.val < o.val
}


type value_impl_int16 struct {
	val int16
}
func (v *value_impl_int16) Value() interface{} {
	return v.val
}
func (v *value_impl_int16) Equals(other value_t) bool {
	o, ok := other.(*value_impl_int16)
	return ok && v.val == o.val
}
func (v *value_impl_int16) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_int16)
	return ok && v.val < o.val
}


type value_impl_uint16 struct {
	val uint16
}
func (v *value_impl_uint16) Value() interface{} {
	return v.val
}
func (v *value_impl_uint16) Equals(other value_t) bool {
	o, ok := other.(*value_impl_uint16)
	return ok && v.val == o.val
}
func (v *value_impl_uint16) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_uint16)
	return ok && v.val < o.val
}


type value_impl_int32 struct {
	val int32
}
func (v *value_impl_int32) Value() interface{} {
	return v.val
}
func (v *value_impl_int32) Equals(other value_t) bool {
	o, ok := other.(*value_impl_int32)
	return ok && v.val == o.val
}
func (v *value_impl_int32) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_int32)
	return ok && v.val < o.val
}


type value_impl_uint32 struct {
	val uint32
}
func (v *value_impl_uint32) Value() interface{} {
	return v.val
}
func (v *value_impl_uint32) Equals(other value_t) bool {
	o, ok := other.(*value_impl_uint32)
	return ok && v.val == o.val
}
func (v *value_impl_uint32) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_uint32)
	return ok && v.val < o.val
}


type value_impl_int64 struct {
	val int64
}
func (v *value_impl_int64) Value() interface{} {
	return v.val
}
func (v *value_impl_int64) Equals(other value_t) bool {
	o, ok := other.(*value_impl_int64)
	return ok && v.val == o.val
}
func (v *value_impl_int64) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_int64)
	return ok && v.val < o.val
}


type value_impl_uint64 struct {
	val uint64
}
func (v *value_impl_uint64) Value() interface{} {
	return v.val
}
func (v *value_impl_uint64) Equals(other value_t) bool {
	o, ok := other.(*value_impl_uint64)
	return ok && v.val == o.val
}
func (v *value_impl_uint64) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_uint64)
	return ok && v.val < o.val
}


type value_impl_string struct {
	val string
}
func (v *value_impl_string) Value() interface{} {
	return v.val
}
func (v *value_impl_string) Equals(other value_t) bool {
	o, ok := other.(*value_impl_string)
	return ok && v.val == o.val
}
func (v *value_impl_string) LessThan(other value_t) bool {
	o, ok := other.(*value_impl_string)
	return ok && v.val < o.val
}