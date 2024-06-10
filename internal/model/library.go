package model

type Library string

const (
	LibraryGolangciLint Library = "golangci-lint"
	LibraryMockery      Library = "mockery"
)

var SupportedLibraries = map[Library]struct{}{
	LibraryGolangciLint: {},
	LibraryMockery:      {},
}

func (l Library) String() string {
	return string(l)
}
