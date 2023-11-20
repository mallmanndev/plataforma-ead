package ports

type UUIDService interface {
	Generate() string
}
