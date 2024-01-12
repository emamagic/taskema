package mysql

type Scanner interface {
	Scan(dest ...any) error
}