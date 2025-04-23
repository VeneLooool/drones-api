package repository

// All вернет "*"
func All() string {
	return "*"
}

type ArbitrarySQLChangeable[T any] interface {
	SQL(string) T
}

func ReturningAll[T ArbitrarySQLChangeable[T]](sb T) T {
	return sb.SQL("RETURNING *")
}
