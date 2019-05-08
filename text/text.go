// Package text disponibiliza métodos para trabalhar com textos/strings.
package text

import (
	"math/rand"
	"regexp"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// encontra todos os caracteres não-numéricos.
var regexExceptNumbers = regexp.MustCompile("[^0-9]")

// OnlyNumbers retorna apenas os números de uma string.
//
//   str := OnlyNumbers("A123BC")  // "123"
//
func OnlyNumbers(str string) string {
	return regexExceptNumbers.ReplaceAllString(str, "")
}

// Normalize substitui caracteres especiais de uma string por caracteres
// da tabela ASCII.
//
//  str := Normalize("JI-PARANÁ")  // "JI-PARANA"
//
func Normalize(str string) (string, error) {
	transf := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	data, _, err := transform.String(transf, str)
	return data, err
}

// RandomCharacters gerar caracteres alertórios com base na quantidade informada
//
// str := RandomCharacters(4) // XwpT
//
func RandomCharacters(numberOfCharacters int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, numberOfCharacters)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}