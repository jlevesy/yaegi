// Code generated by 'goexports strconv'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"go/constant"
	"go/token"
	"reflect"
	"strconv"
)

func init() {
	Symbols["strconv"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AppendBool":               reflect.ValueOf(strconv.AppendBool),
		"AppendFloat":              reflect.ValueOf(strconv.AppendFloat),
		"AppendInt":                reflect.ValueOf(strconv.AppendInt),
		"AppendQuote":              reflect.ValueOf(strconv.AppendQuote),
		"AppendQuoteRune":          reflect.ValueOf(strconv.AppendQuoteRune),
		"AppendQuoteRuneToASCII":   reflect.ValueOf(strconv.AppendQuoteRuneToASCII),
		"AppendQuoteRuneToGraphic": reflect.ValueOf(strconv.AppendQuoteRuneToGraphic),
		"AppendQuoteToASCII":       reflect.ValueOf(strconv.AppendQuoteToASCII),
		"AppendQuoteToGraphic":     reflect.ValueOf(strconv.AppendQuoteToGraphic),
		"AppendUint":               reflect.ValueOf(strconv.AppendUint),
		"Atoi":                     reflect.ValueOf(strconv.Atoi),
		"CanBackquote":             reflect.ValueOf(strconv.CanBackquote),
		"ErrRange":                 reflect.ValueOf(&strconv.ErrRange).Elem(),
		"ErrSyntax":                reflect.ValueOf(&strconv.ErrSyntax).Elem(),
		"FormatBool":               reflect.ValueOf(strconv.FormatBool),
		"FormatFloat":              reflect.ValueOf(strconv.FormatFloat),
		"FormatInt":                reflect.ValueOf(strconv.FormatInt),
		"FormatUint":               reflect.ValueOf(strconv.FormatUint),
		"IntSize":                  reflect.ValueOf(constant.MakeFromLiteral("64", token.INT, 0)),
		"IsGraphic":                reflect.ValueOf(strconv.IsGraphic),
		"IsPrint":                  reflect.ValueOf(strconv.IsPrint),
		"Itoa":                     reflect.ValueOf(strconv.Itoa),
		"ParseBool":                reflect.ValueOf(strconv.ParseBool),
		"ParseFloat":               reflect.ValueOf(strconv.ParseFloat),
		"ParseInt":                 reflect.ValueOf(strconv.ParseInt),
		"ParseUint":                reflect.ValueOf(strconv.ParseUint),
		"Quote":                    reflect.ValueOf(strconv.Quote),
		"QuoteRune":                reflect.ValueOf(strconv.QuoteRune),
		"QuoteRuneToASCII":         reflect.ValueOf(strconv.QuoteRuneToASCII),
		"QuoteRuneToGraphic":       reflect.ValueOf(strconv.QuoteRuneToGraphic),
		"QuoteToASCII":             reflect.ValueOf(strconv.QuoteToASCII),
		"QuoteToGraphic":           reflect.ValueOf(strconv.QuoteToGraphic),
		"Unquote":                  reflect.ValueOf(strconv.Unquote),
		"UnquoteChar":              reflect.ValueOf(strconv.UnquoteChar),

		// type definitions
		"NumError": reflect.ValueOf((*strconv.NumError)(nil)),
	}
}
