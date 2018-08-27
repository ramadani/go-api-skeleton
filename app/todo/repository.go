package todo

// Repository interface for concrete repository
type Repository interface {
	Create()
	Find(id uint)
	Update(id uint)
	Delete(id uint)
}
