package conv

func ToUint(v *int32) *uint {
	if v == nil {
		return nil
	}
	res := uint(*v)
	return &res
}

func ToInt32(v *uint) *int32 {
	if v == nil {
		return nil
	}
	res := int32(*v)
	return &res
}
