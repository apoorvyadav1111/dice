package eval

const (
	BYTE = "BYTE"
	BIT  = "BIT"

	AND string = "AND"
	OR  string = "OR"
	XOR string = "XOR"
	NOT string = "NOT"

	Ex         string = "EX"
	Px         string = "PX"
	Pxat       string = "PXAT"
	Exat       string = "EXAT"
	XX         string = "XX"
	NX         string = "NX"
	GT         string = "GT"
	LT         string = "LT"
	KeepTTL    string = "KEEPTTL"
	Sync       string = "SYNC"
	Async      string = "ASYNC"
	Help       string = "HELP"
	Memory     string = "MEMORY"
	Count      string = "COUNT"
	GetKeys    string = "GETKEYS"
	List       string = "LIST"
	Info       string = "INFO"
	null       string = "null"
	WithValues string = "WITHVALUES"
	WithScores string = "WITHSCORES"
	REV        string = "REV"
	GET        string = "GET"
	SET        string = "SET"
	INCRBY     string = "INCRBY"
	OVERFLOW   string = "OVERFLOW"
	WRAP       string = "WRAP"
	SAT        string = "SAT"
	FAIL       string = "FAIL"
	SIGNED     string = "SIGNED"
	UNSIGNED   string = "UNSIGNED"
	FIELDS     string = "FIELDS"
	NOT_FOUND  int64  = -2
	PAST       int64  = -1
	EXPIRY_SET int64  = 1
)
