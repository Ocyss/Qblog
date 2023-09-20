package conv

func ToUint(v *int64) *uint {
	if v == nil {
		return nil
	}
	res := uint(*v)
	return &res
}

func ToInt64(v *uint) *int64 {
	if v == nil {
		return nil
	}
	res := int64(*v)
	return &res
}
