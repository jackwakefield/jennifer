package jen

/*
	This file is generated by genjen - do not edit!
*/

// List inserts a comma seperated list
func List(code ...Code) *Statement {
	s := new(Statement)
	return s.List(code...)
}

// List inserts a comma seperated list
func (l *StatementList) List(code ...Code) *Statement {
	s := List(code...)
	*l = append(*l, s)
	return s
}

// List inserts a comma seperated list
func (s *Statement) List(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		seperator: ",",
	}
	*s = append(*s, b)
	return s
}

// Parens inserts parenthesis
func Parens(code ...Code) *Statement {
	s := new(Statement)
	return s.Parens(code...)
}

// Parens inserts parenthesis
func (l *StatementList) Parens(code ...Code) *Statement {
	s := Parens(code...)
	*l = append(*l, s)
	return s
}

// Parens inserts parenthesis
func (s *Statement) Parens(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "(",
		close:     ")",
	}
	*s = append(*s, b)
	return s
}

// Braces inserts curly braces
func Braces(code ...Code) *Statement {
	s := new(Statement)
	return s.Braces(code...)
}

// Braces inserts curly braces
func (l *StatementList) Braces(code ...Code) *Statement {
	s := Braces(code...)
	*l = append(*l, s)
	return s
}

// Braces inserts curly braces
func (s *Statement) Braces(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "{",
		close:     "}",
	}
	*s = append(*s, b)
	return s
}

// Values inserts curly braces containing a comma separated list
func Values(code ...Code) *Statement {
	s := new(Statement)
	return s.Values(code...)
}

// Values inserts curly braces containing a comma separated list
func (l *StatementList) Values(code ...Code) *Statement {
	s := Values(code...)
	*l = append(*l, s)
	return s
}

// Values inserts curly braces containing a comma separated list
func (s *Statement) Values(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "{",
		close:     "}",
		seperator: ",",
	}
	*s = append(*s, b)
	return s
}

// Index inserts square brackets containing a colon separated list
func Index(code ...Code) *Statement {
	s := new(Statement)
	return s.Index(code...)
}

// Index inserts square brackets containing a colon separated list
func (l *StatementList) Index(code ...Code) *Statement {
	s := Index(code...)
	*l = append(*l, s)
	return s
}

// Index inserts square brackets containing a colon separated list
func (s *Statement) Index(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "[",
		close:     "]",
		seperator: ":",
	}
	*s = append(*s, b)
	return s
}

// Block inserts curly braces containing a statements list
func Block(code ...Code) *Statement {
	s := new(Statement)
	return s.Block(code...)
}

// Block inserts curly braces containing a statements list
func (l *StatementList) Block(code ...Code) *Statement {
	s := Block(code...)
	*l = append(*l, s)
	return s
}

// Block inserts curly braces containing a statements list
func (s *Statement) Block(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "{",
		close:     "}",
		seperator: ";",
	}
	*s = append(*s, b)
	return s
}

// Call inserts parenthesis containing a comma separated list
func Call(code ...Code) *Statement {
	s := new(Statement)
	return s.Call(code...)
}

// Call inserts parenthesis containing a comma separated list
func (l *StatementList) Call(code ...Code) *Statement {
	s := Call(code...)
	*l = append(*l, s)
	return s
}

// Call inserts parenthesis containing a comma separated list
func (s *Statement) Call(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "(",
		close:     ")",
		seperator: ",",
	}
	*s = append(*s, b)
	return s
}

// Params inserts parenthesis containing a comma separated list
func Params(code ...Code) *Statement {
	s := new(Statement)
	return s.Params(code...)
}

// Params inserts parenthesis containing a comma separated list
func (l *StatementList) Params(code ...Code) *Statement {
	s := Params(code...)
	*l = append(*l, s)
	return s
}

// Params inserts parenthesis containing a comma separated list
func (s *Statement) Params(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "(",
		close:     ")",
		seperator: ",",
	}
	*s = append(*s, b)
	return s
}

// Decls inserts parenthesis containing a statement list
func Decls(code ...Code) *Statement {
	s := new(Statement)
	return s.Decls(code...)
}

// Decls inserts parenthesis containing a statement list
func (l *StatementList) Decls(code ...Code) *Statement {
	s := Decls(code...)
	*l = append(*l, s)
	return s
}

// Decls inserts parenthesis containing a statement list
func (s *Statement) Decls(code ...Code) *Statement {
	b := block{
		Statement: s,
		code:      code,
		open:      "(",
		close:     ")",
		seperator: ";",
	}
	*s = append(*s, b)
	return s
}

func Bool() *Statement {
	s := new(Statement)
	return s.Bool()
}

func (l *StatementList) Bool() *Statement {
	s := Bool()
	*l = append(*l, s)
	return s
}

func (s *Statement) Bool() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "bool",
	}
	*s = append(*s, t)
	return s
}

func Byte() *Statement {
	s := new(Statement)
	return s.Byte()
}

func (l *StatementList) Byte() *Statement {
	s := Byte()
	*l = append(*l, s)
	return s
}

func (s *Statement) Byte() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "byte",
	}
	*s = append(*s, t)
	return s
}

func Complex64() *Statement {
	s := new(Statement)
	return s.Complex64()
}

func (l *StatementList) Complex64() *Statement {
	s := Complex64()
	*l = append(*l, s)
	return s
}

func (s *Statement) Complex64() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "complex64",
	}
	*s = append(*s, t)
	return s
}

func Complex128() *Statement {
	s := new(Statement)
	return s.Complex128()
}

func (l *StatementList) Complex128() *Statement {
	s := Complex128()
	*l = append(*l, s)
	return s
}

func (s *Statement) Complex128() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "complex128",
	}
	*s = append(*s, t)
	return s
}

func Error() *Statement {
	s := new(Statement)
	return s.Error()
}

func (l *StatementList) Error() *Statement {
	s := Error()
	*l = append(*l, s)
	return s
}

func (s *Statement) Error() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "error",
	}
	*s = append(*s, t)
	return s
}

func Float32() *Statement {
	s := new(Statement)
	return s.Float32()
}

func (l *StatementList) Float32() *Statement {
	s := Float32()
	*l = append(*l, s)
	return s
}

func (s *Statement) Float32() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "float32",
	}
	*s = append(*s, t)
	return s
}

func Float64() *Statement {
	s := new(Statement)
	return s.Float64()
}

func (l *StatementList) Float64() *Statement {
	s := Float64()
	*l = append(*l, s)
	return s
}

func (s *Statement) Float64() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "float64",
	}
	*s = append(*s, t)
	return s
}

func Int() *Statement {
	s := new(Statement)
	return s.Int()
}

func (l *StatementList) Int() *Statement {
	s := Int()
	*l = append(*l, s)
	return s
}

func (s *Statement) Int() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "int",
	}
	*s = append(*s, t)
	return s
}

func Int8() *Statement {
	s := new(Statement)
	return s.Int8()
}

func (l *StatementList) Int8() *Statement {
	s := Int8()
	*l = append(*l, s)
	return s
}

func (s *Statement) Int8() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "int8",
	}
	*s = append(*s, t)
	return s
}

func Int16() *Statement {
	s := new(Statement)
	return s.Int16()
}

func (l *StatementList) Int16() *Statement {
	s := Int16()
	*l = append(*l, s)
	return s
}

func (s *Statement) Int16() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "int16",
	}
	*s = append(*s, t)
	return s
}

func Int32() *Statement {
	s := new(Statement)
	return s.Int32()
}

func (l *StatementList) Int32() *Statement {
	s := Int32()
	*l = append(*l, s)
	return s
}

func (s *Statement) Int32() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "int32",
	}
	*s = append(*s, t)
	return s
}

func Int64() *Statement {
	s := new(Statement)
	return s.Int64()
}

func (l *StatementList) Int64() *Statement {
	s := Int64()
	*l = append(*l, s)
	return s
}

func (s *Statement) Int64() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "int64",
	}
	*s = append(*s, t)
	return s
}

func Rune() *Statement {
	s := new(Statement)
	return s.Rune()
}

func (l *StatementList) Rune() *Statement {
	s := Rune()
	*l = append(*l, s)
	return s
}

func (s *Statement) Rune() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "rune",
	}
	*s = append(*s, t)
	return s
}

func String() *Statement {
	s := new(Statement)
	return s.String()
}

func (l *StatementList) String() *Statement {
	s := String()
	*l = append(*l, s)
	return s
}

func (s *Statement) String() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "string",
	}
	*s = append(*s, t)
	return s
}

func Uint() *Statement {
	s := new(Statement)
	return s.Uint()
}

func (l *StatementList) Uint() *Statement {
	s := Uint()
	*l = append(*l, s)
	return s
}

func (s *Statement) Uint() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "uint",
	}
	*s = append(*s, t)
	return s
}

func Uint8() *Statement {
	s := new(Statement)
	return s.Uint8()
}

func (l *StatementList) Uint8() *Statement {
	s := Uint8()
	*l = append(*l, s)
	return s
}

func (s *Statement) Uint8() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "uint8",
	}
	*s = append(*s, t)
	return s
}

func Uint16() *Statement {
	s := new(Statement)
	return s.Uint16()
}

func (l *StatementList) Uint16() *Statement {
	s := Uint16()
	*l = append(*l, s)
	return s
}

func (s *Statement) Uint16() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "uint16",
	}
	*s = append(*s, t)
	return s
}

func Uint32() *Statement {
	s := new(Statement)
	return s.Uint32()
}

func (l *StatementList) Uint32() *Statement {
	s := Uint32()
	*l = append(*l, s)
	return s
}

func (s *Statement) Uint32() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "uint32",
	}
	*s = append(*s, t)
	return s
}

func Uint64() *Statement {
	s := new(Statement)
	return s.Uint64()
}

func (l *StatementList) Uint64() *Statement {
	s := Uint64()
	*l = append(*l, s)
	return s
}

func (s *Statement) Uint64() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "uint64",
	}
	*s = append(*s, t)
	return s
}

func Uintptr() *Statement {
	s := new(Statement)
	return s.Uintptr()
}

func (l *StatementList) Uintptr() *Statement {
	s := Uintptr()
	*l = append(*l, s)
	return s
}

func (s *Statement) Uintptr() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "uintptr",
	}
	*s = append(*s, t)
	return s
}

func True() *Statement {
	s := new(Statement)
	return s.True()
}

func (l *StatementList) True() *Statement {
	s := True()
	*l = append(*l, s)
	return s
}

func (s *Statement) True() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "true",
	}
	*s = append(*s, t)
	return s
}

func False() *Statement {
	s := new(Statement)
	return s.False()
}

func (l *StatementList) False() *Statement {
	s := False()
	*l = append(*l, s)
	return s
}

func (s *Statement) False() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "false",
	}
	*s = append(*s, t)
	return s
}

func Iota() *Statement {
	s := new(Statement)
	return s.Iota()
}

func (l *StatementList) Iota() *Statement {
	s := Iota()
	*l = append(*l, s)
	return s
}

func (s *Statement) Iota() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "iota",
	}
	*s = append(*s, t)
	return s
}

func Nil() *Statement {
	s := new(Statement)
	return s.Nil()
}

func (l *StatementList) Nil() *Statement {
	s := Nil()
	*l = append(*l, s)
	return s
}

func (s *Statement) Nil() *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "nil",
	}
	*s = append(*s, t)
	return s
}

func Append(c ...Code) *Statement {
	s := new(Statement)
	return s.Append(c...)
}

func (l *StatementList) Append(c ...Code) *Statement {
	s := Append(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Append(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "append",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Cap(c ...Code) *Statement {
	s := new(Statement)
	return s.Cap(c...)
}

func (l *StatementList) Cap(c ...Code) *Statement {
	s := Cap(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Cap(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "cap",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Close(c ...Code) *Statement {
	s := new(Statement)
	return s.Close(c...)
}

func (l *StatementList) Close(c ...Code) *Statement {
	s := Close(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Close(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "close",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Complex(c ...Code) *Statement {
	s := new(Statement)
	return s.Complex(c...)
}

func (l *StatementList) Complex(c ...Code) *Statement {
	s := Complex(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Complex(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "complex",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Copy(c ...Code) *Statement {
	s := new(Statement)
	return s.Copy(c...)
}

func (l *StatementList) Copy(c ...Code) *Statement {
	s := Copy(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Copy(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "copy",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Delete(c ...Code) *Statement {
	s := new(Statement)
	return s.Delete(c...)
}

func (l *StatementList) Delete(c ...Code) *Statement {
	s := Delete(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Delete(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "delete",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Imag(c ...Code) *Statement {
	s := new(Statement)
	return s.Imag(c...)
}

func (l *StatementList) Imag(c ...Code) *Statement {
	s := Imag(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Imag(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "imag",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Len(c ...Code) *Statement {
	s := new(Statement)
	return s.Len(c...)
}

func (l *StatementList) Len(c ...Code) *Statement {
	s := Len(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Len(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "len",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Make(c ...Code) *Statement {
	s := new(Statement)
	return s.Make(c...)
}

func (l *StatementList) Make(c ...Code) *Statement {
	s := Make(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Make(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "make",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func New(c ...Code) *Statement {
	s := new(Statement)
	return s.New(c...)
}

func (l *StatementList) New(c ...Code) *Statement {
	s := New(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) New(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "new",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Panic(c ...Code) *Statement {
	s := new(Statement)
	return s.Panic(c...)
}

func (l *StatementList) Panic(c ...Code) *Statement {
	s := Panic(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Panic(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "panic",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Print(c ...Code) *Statement {
	s := new(Statement)
	return s.Print(c...)
}

func (l *StatementList) Print(c ...Code) *Statement {
	s := Print(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Print(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "print",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Println(c ...Code) *Statement {
	s := new(Statement)
	return s.Println(c...)
}

func (l *StatementList) Println(c ...Code) *Statement {
	s := Println(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Println(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "println",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Real(c ...Code) *Statement {
	s := new(Statement)
	return s.Real(c...)
}

func (l *StatementList) Real(c ...Code) *Statement {
	s := Real(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Real(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "real",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Recover(c ...Code) *Statement {
	s := new(Statement)
	return s.Recover(c...)
}

func (l *StatementList) Recover(c ...Code) *Statement {
	s := Recover(c...)
	*l = append(*l, s)
	return s
}

func (s *Statement) Recover(c ...Code) *Statement {
	t := Token{
		Statement: s,
		typ:       identifierToken,
		content:   "recover",
	}
	ca := Call(c...)
	*s = append(*s, t, ca)
	return s
}

func Break() *Statement {
	s := new(Statement)
	return s.Break()
}

func (l *StatementList) Break() *Statement {
	s := Break()
	*l = append(*l, s)
	return s
}

func (s *Statement) Break() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "break",
	}
	*s = append(*s, t)
	return s
}

func Default() *Statement {
	s := new(Statement)
	return s.Default()
}

func (l *StatementList) Default() *Statement {
	s := Default()
	*l = append(*l, s)
	return s
}

func (s *Statement) Default() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "default",
	}
	*s = append(*s, t)
	return s
}

func Func() *Statement {
	s := new(Statement)
	return s.Func()
}

func (l *StatementList) Func() *Statement {
	s := Func()
	*l = append(*l, s)
	return s
}

func (s *Statement) Func() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "func",
	}
	*s = append(*s, t)
	return s
}

func Interface() *Statement {
	s := new(Statement)
	return s.Interface()
}

func (l *StatementList) Interface() *Statement {
	s := Interface()
	*l = append(*l, s)
	return s
}

func (s *Statement) Interface() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "interface",
	}
	*s = append(*s, t)
	return s
}

func Select() *Statement {
	s := new(Statement)
	return s.Select()
}

func (l *StatementList) Select() *Statement {
	s := Select()
	*l = append(*l, s)
	return s
}

func (s *Statement) Select() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "select",
	}
	*s = append(*s, t)
	return s
}

func Case() *Statement {
	s := new(Statement)
	return s.Case()
}

func (l *StatementList) Case() *Statement {
	s := Case()
	*l = append(*l, s)
	return s
}

func (s *Statement) Case() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "case",
	}
	*s = append(*s, t)
	return s
}

func Defer() *Statement {
	s := new(Statement)
	return s.Defer()
}

func (l *StatementList) Defer() *Statement {
	s := Defer()
	*l = append(*l, s)
	return s
}

func (s *Statement) Defer() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "defer",
	}
	*s = append(*s, t)
	return s
}

func Go() *Statement {
	s := new(Statement)
	return s.Go()
}

func (l *StatementList) Go() *Statement {
	s := Go()
	*l = append(*l, s)
	return s
}

func (s *Statement) Go() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "go",
	}
	*s = append(*s, t)
	return s
}

func Map() *Statement {
	s := new(Statement)
	return s.Map()
}

func (l *StatementList) Map() *Statement {
	s := Map()
	*l = append(*l, s)
	return s
}

func (s *Statement) Map() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "map",
	}
	*s = append(*s, t)
	return s
}

func Struct() *Statement {
	s := new(Statement)
	return s.Struct()
}

func (l *StatementList) Struct() *Statement {
	s := Struct()
	*l = append(*l, s)
	return s
}

func (s *Statement) Struct() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "struct",
	}
	*s = append(*s, t)
	return s
}

func Chan() *Statement {
	s := new(Statement)
	return s.Chan()
}

func (l *StatementList) Chan() *Statement {
	s := Chan()
	*l = append(*l, s)
	return s
}

func (s *Statement) Chan() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "chan",
	}
	*s = append(*s, t)
	return s
}

func Else() *Statement {
	s := new(Statement)
	return s.Else()
}

func (l *StatementList) Else() *Statement {
	s := Else()
	*l = append(*l, s)
	return s
}

func (s *Statement) Else() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "else",
	}
	*s = append(*s, t)
	return s
}

func Goto() *Statement {
	s := new(Statement)
	return s.Goto()
}

func (l *StatementList) Goto() *Statement {
	s := Goto()
	*l = append(*l, s)
	return s
}

func (s *Statement) Goto() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "goto",
	}
	*s = append(*s, t)
	return s
}

func Package() *Statement {
	s := new(Statement)
	return s.Package()
}

func (l *StatementList) Package() *Statement {
	s := Package()
	*l = append(*l, s)
	return s
}

func (s *Statement) Package() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "package",
	}
	*s = append(*s, t)
	return s
}

func Switch() *Statement {
	s := new(Statement)
	return s.Switch()
}

func (l *StatementList) Switch() *Statement {
	s := Switch()
	*l = append(*l, s)
	return s
}

func (s *Statement) Switch() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "switch",
	}
	*s = append(*s, t)
	return s
}

func Const() *Statement {
	s := new(Statement)
	return s.Const()
}

func (l *StatementList) Const() *Statement {
	s := Const()
	*l = append(*l, s)
	return s
}

func (s *Statement) Const() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "const",
	}
	*s = append(*s, t)
	return s
}

func Fallthrough() *Statement {
	s := new(Statement)
	return s.Fallthrough()
}

func (l *StatementList) Fallthrough() *Statement {
	s := Fallthrough()
	*l = append(*l, s)
	return s
}

func (s *Statement) Fallthrough() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "fallthrough",
	}
	*s = append(*s, t)
	return s
}

func If() *Statement {
	s := new(Statement)
	return s.If()
}

func (l *StatementList) If() *Statement {
	s := If()
	*l = append(*l, s)
	return s
}

func (s *Statement) If() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "if",
	}
	*s = append(*s, t)
	return s
}

func Range() *Statement {
	s := new(Statement)
	return s.Range()
}

func (l *StatementList) Range() *Statement {
	s := Range()
	*l = append(*l, s)
	return s
}

func (s *Statement) Range() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "range",
	}
	*s = append(*s, t)
	return s
}

func Type() *Statement {
	s := new(Statement)
	return s.Type()
}

func (l *StatementList) Type() *Statement {
	s := Type()
	*l = append(*l, s)
	return s
}

func (s *Statement) Type() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "type",
	}
	*s = append(*s, t)
	return s
}

func Continue() *Statement {
	s := new(Statement)
	return s.Continue()
}

func (l *StatementList) Continue() *Statement {
	s := Continue()
	*l = append(*l, s)
	return s
}

func (s *Statement) Continue() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "continue",
	}
	*s = append(*s, t)
	return s
}

func For() *Statement {
	s := new(Statement)
	return s.For()
}

func (l *StatementList) For() *Statement {
	s := For()
	*l = append(*l, s)
	return s
}

func (s *Statement) For() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "for",
	}
	*s = append(*s, t)
	return s
}

func Import() *Statement {
	s := new(Statement)
	return s.Import()
}

func (l *StatementList) Import() *Statement {
	s := Import()
	*l = append(*l, s)
	return s
}

func (s *Statement) Import() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "import",
	}
	*s = append(*s, t)
	return s
}

func Return() *Statement {
	s := new(Statement)
	return s.Return()
}

func (l *StatementList) Return() *Statement {
	s := Return()
	*l = append(*l, s)
	return s
}

func (s *Statement) Return() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "return",
	}
	*s = append(*s, t)
	return s
}

func Var() *Statement {
	s := new(Statement)
	return s.Var()
}

func (l *StatementList) Var() *Statement {
	s := Var()
	*l = append(*l, s)
	return s
}

func (s *Statement) Var() *Statement {
	t := Token{
		Statement: s,
		typ:       keywordToken,
		content:   "var",
	}
	*s = append(*s, t)
	return s
}

// Sum inserts the sum operator (+)
func Sum() *Statement {
	s := new(Statement)
	return s.Sum()
}

// Sum inserts the sum operator (+)
func (l *StatementList) Sum() *Statement {
	s := Sum()
	*l = append(*l, s)
	return s
}

// Sum inserts the sum operator (+)
func (s *Statement) Sum() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "+",
	}
	*s = append(*s, t)
	return s
}

// Diff inserts the difference operator (-)
func Diff() *Statement {
	s := new(Statement)
	return s.Diff()
}

// Diff inserts the difference operator (-)
func (l *StatementList) Diff() *Statement {
	s := Diff()
	*l = append(*l, s)
	return s
}

// Diff inserts the difference operator (-)
func (s *Statement) Diff() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "-",
	}
	*s = append(*s, t)
	return s
}

// Product inserts the product operator (*)
func Product() *Statement {
	s := new(Statement)
	return s.Product()
}

// Product inserts the product operator (*)
func (l *StatementList) Product() *Statement {
	s := Product()
	*l = append(*l, s)
	return s
}

// Product inserts the product operator (*)
func (s *Statement) Product() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "*",
	}
	*s = append(*s, t)
	return s
}

// Quotient inserts the quotient operator (/)
func Quotient() *Statement {
	s := new(Statement)
	return s.Quotient()
}

// Quotient inserts the quotient operator (/)
func (l *StatementList) Quotient() *Statement {
	s := Quotient()
	*l = append(*l, s)
	return s
}

// Quotient inserts the quotient operator (/)
func (s *Statement) Quotient() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "/",
	}
	*s = append(*s, t)
	return s
}

// Remainder inserts the remainder operator (%)
func Remainder() *Statement {
	s := new(Statement)
	return s.Remainder()
}

// Remainder inserts the remainder operator (%)
func (l *StatementList) Remainder() *Statement {
	s := Remainder()
	*l = append(*l, s)
	return s
}

// Remainder inserts the remainder operator (%)
func (s *Statement) Remainder() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "%",
	}
	*s = append(*s, t)
	return s
}

// As inserts the assignment operator (=)
func As() *Statement {
	s := new(Statement)
	return s.As()
}

// As inserts the assignment operator (=)
func (l *StatementList) As() *Statement {
	s := As()
	*l = append(*l, s)
	return s
}

// As inserts the assignment operator (=)
func (s *Statement) As() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "=",
	}
	*s = append(*s, t)
	return s
}

// Sas inserts the short assignment operator (:=)
func Sas() *Statement {
	s := new(Statement)
	return s.Sas()
}

// Sas inserts the short assignment operator (:=)
func (l *StatementList) Sas() *Statement {
	s := Sas()
	*l = append(*l, s)
	return s
}

// Sas inserts the short assignment operator (:=)
func (s *Statement) Sas() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   ":=",
	}
	*s = append(*s, t)
	return s
}

// Sel inserts the selector operator (.)
func Sel() *Statement {
	s := new(Statement)
	return s.Sel()
}

// Sel inserts the selector operator (.)
func (l *StatementList) Sel() *Statement {
	s := Sel()
	*l = append(*l, s)
	return s
}

// Sel inserts the selector operator (.)
func (s *Statement) Sel() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   ".",
	}
	*s = append(*s, t)
	return s
}

// Eq inserts the equal operator (==)
func Eq() *Statement {
	s := new(Statement)
	return s.Eq()
}

// Eq inserts the equal operator (==)
func (l *StatementList) Eq() *Statement {
	s := Eq()
	*l = append(*l, s)
	return s
}

// Eq inserts the equal operator (==)
func (s *Statement) Eq() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "==",
	}
	*s = append(*s, t)
	return s
}

// Neq inserts the not equal operator (!=)
func Neq() *Statement {
	s := new(Statement)
	return s.Neq()
}

// Neq inserts the not equal operator (!=)
func (l *StatementList) Neq() *Statement {
	s := Neq()
	*l = append(*l, s)
	return s
}

// Neq inserts the not equal operator (!=)
func (s *Statement) Neq() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "!=",
	}
	*s = append(*s, t)
	return s
}

// Lt inserts the less than operator (<)
func Lt() *Statement {
	s := new(Statement)
	return s.Lt()
}

// Lt inserts the less than operator (<)
func (l *StatementList) Lt() *Statement {
	s := Lt()
	*l = append(*l, s)
	return s
}

// Lt inserts the less than operator (<)
func (s *Statement) Lt() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "<",
	}
	*s = append(*s, t)
	return s
}

// Lte inserts the less than or equal operator (<=)
func Lte() *Statement {
	s := new(Statement)
	return s.Lte()
}

// Lte inserts the less than or equal operator (<=)
func (l *StatementList) Lte() *Statement {
	s := Lte()
	*l = append(*l, s)
	return s
}

// Lte inserts the less than or equal operator (<=)
func (s *Statement) Lte() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "<=",
	}
	*s = append(*s, t)
	return s
}

// Gt inserts the greater than operator (>)
func Gt() *Statement {
	s := new(Statement)
	return s.Gt()
}

// Gt inserts the greater than operator (>)
func (l *StatementList) Gt() *Statement {
	s := Gt()
	*l = append(*l, s)
	return s
}

// Gt inserts the greater than operator (>)
func (s *Statement) Gt() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   ">",
	}
	*s = append(*s, t)
	return s
}

// Gte inserts the greater than or equal operator (>=)
func Gte() *Statement {
	s := new(Statement)
	return s.Gte()
}

// Gte inserts the greater than or equal operator (>=)
func (l *StatementList) Gte() *Statement {
	s := Gte()
	*l = append(*l, s)
	return s
}

// Gte inserts the greater than or equal operator (>=)
func (s *Statement) Gte() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   ">=",
	}
	*s = append(*s, t)
	return s
}

// And inserts the conditional and operator (&&)
func And() *Statement {
	s := new(Statement)
	return s.And()
}

// And inserts the conditional and operator (&&)
func (l *StatementList) And() *Statement {
	s := And()
	*l = append(*l, s)
	return s
}

// And inserts the conditional and operator (&&)
func (s *Statement) And() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "&&",
	}
	*s = append(*s, t)
	return s
}

// Or inserts the conditional or operator (||)
func Or() *Statement {
	s := new(Statement)
	return s.Or()
}

// Or inserts the conditional or operator (||)
func (l *StatementList) Or() *Statement {
	s := Or()
	*l = append(*l, s)
	return s
}

// Or inserts the conditional or operator (||)
func (s *Statement) Or() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "||",
	}
	*s = append(*s, t)
	return s
}

// Not inserts the conditional not operator (!)
func Not() *Statement {
	s := new(Statement)
	return s.Not()
}

// Not inserts the conditional not operator (!)
func (l *StatementList) Not() *Statement {
	s := Not()
	*l = append(*l, s)
	return s
}

// Not inserts the conditional not operator (!)
func (s *Statement) Not() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "!",
	}
	*s = append(*s, t)
	return s
}

// Rcv inserts the channel receive operator (<-)
func Rcv() *Statement {
	s := new(Statement)
	return s.Rcv()
}

// Rcv inserts the channel receive operator (<-)
func (l *StatementList) Rcv() *Statement {
	s := Rcv()
	*l = append(*l, s)
	return s
}

// Rcv inserts the channel receive operator (<-)
func (s *Statement) Rcv() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "<-",
	}
	*s = append(*s, t)
	return s
}

// Vari inserts the variadic operator (...)
func Vari() *Statement {
	s := new(Statement)
	return s.Vari()
}

// Vari inserts the variadic operator (...)
func (l *StatementList) Vari() *Statement {
	s := Vari()
	*l = append(*l, s)
	return s
}

// Vari inserts the variadic operator (...)
func (s *Statement) Vari() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "...",
	}
	*s = append(*s, t)
	return s
}

// Inc inserts the increment operator (++)
func Inc() *Statement {
	s := new(Statement)
	return s.Inc()
}

// Inc inserts the increment operator (++)
func (l *StatementList) Inc() *Statement {
	s := Inc()
	*l = append(*l, s)
	return s
}

// Inc inserts the increment operator (++)
func (s *Statement) Inc() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "++",
	}
	*s = append(*s, t)
	return s
}

// Dec inserts the decrement operator (--)
func Dec() *Statement {
	s := new(Statement)
	return s.Dec()
}

// Dec inserts the decrement operator (--)
func (l *StatementList) Dec() *Statement {
	s := Dec()
	*l = append(*l, s)
	return s
}

// Dec inserts the decrement operator (--)
func (s *Statement) Dec() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "--",
	}
	*s = append(*s, t)
	return s
}

// Ptr inserts the pointer operator (*)
func Ptr() *Statement {
	s := new(Statement)
	return s.Ptr()
}

// Ptr inserts the pointer operator (*)
func (l *StatementList) Ptr() *Statement {
	s := Ptr()
	*l = append(*l, s)
	return s
}

// Ptr inserts the pointer operator (*)
func (s *Statement) Ptr() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "*",
	}
	*s = append(*s, t)
	return s
}

// Adr inserts the address operator (&)
func Adr() *Statement {
	s := new(Statement)
	return s.Adr()
}

// Adr inserts the address operator (&)
func (l *StatementList) Adr() *Statement {
	s := Adr()
	*l = append(*l, s)
	return s
}

// Adr inserts the address operator (&)
func (s *Statement) Adr() *Statement {
	t := Token{
		Statement: s,
		typ:       operatorToken,
		content:   "&",
	}
	*s = append(*s, t)
	return s
}
