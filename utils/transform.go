package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TransformToName трансформация имени
func TransformToName(str string) string {
	str = strings.TrimSpace(str)
	// Если строка пуста, просто возвращаем её
	if len(str) == 0 {
		return str
	}

	// Преобразуем первое слово в строке
	firstWord := strings.SplitN(str, " ", 2)[0]

	caser := cases.Title(language.BrazilianPortuguese)
	capitalizedFirst := caser.String(strings.ToLower(firstWord))

	// Преобразуем оставшуюся часть строки в нижний регистр
	rest := strings.ToLower(str[len(firstWord):])

	// Объединяем и возвращаем результат
	return capitalizedFirst + rest
}

// TransformToNamePtr трансформация имени
func TransformToNamePtr(strPtr *string) *string {
	if strPtr == nil {
		return nil
	}

	str := TransformToName(*strPtr)

	if str == "" {
		return nil
	}
	return &str
}
