package model

type Library string

const (
	LibraryGolangciLint Library = "golangci-lint"
	LibraryMockery      Library = "mockery"
)

func (l Library) String() string {
	return string(l)
}
