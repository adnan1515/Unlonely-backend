package models

type Identity struct {
	id uint
}
func (identity *Identity) SetId (_id uint) {
	identity.id = _id
}
func (identity * Identity) GetID() uint {
	return identity.id
}