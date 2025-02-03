package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type optiontest struct {
	Strings []string
	Option  Options
}

func TestTableParserSuccess(t *testing.T) {
	var tableTests = []struct {
		in  optiontest
		out []string
	}{
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{}}, []string{"I love music.", " ", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik."}},
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{C: true}}, []string{"3 I love music.", "1  ", "2 I love music of Kartik.",
			"1 Thanks.", "2 I love music of Kartik."}},
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{D: true}}, []string{"I love music.", "I love music of Kartik.", "I love music of Kartik."}},
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{U: true}}, []string{" ", "Thanks."}},
		{optiontest{Strings: []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", " ", "I love MuSIC of Kartik.", "I love music of kartik.",
			"Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."}, Option: Options{I: true}}, []string{"I LOVE MUSIC.", " ", "I love MuSIC of Kartik.", "Thanks.", "I love music of kartik."}},
		{optiontest{Strings: []string{"We love music.", "I love music.", "They love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
			"Thanks."}, Option: Options{F: 1}}, []string{"We love music.", " ", "I love music of Kartik.", "Thanks."}},
		{optiontest{Strings: []string{"I love music.", "A love music.", "C love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
			"Thanks."}, Option: Options{S: 1}}, []string{"I love music.", " ", "I love music of Kartik.", "We love music of Kartik.", "Thanks."}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{}}, []string{"Небо", "Облако", "Пляж", "Анапа"}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{C: true}}, []string{"2 Небо", "2 Облако", "1 Пляж", "2 Анапа"}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{D: true}}, []string{"Небо", "Облако", "Анапа"}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{U: true}}, []string{"Пляж"}},
		{optiontest{Strings: []string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1}}, []string{"Небо голуБое", "Пляж чистый", "Анапа солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1, C: true}}, []string{"4 Небо голуБое", "1 Пляж чистый", "3 Анапа солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1, D: true}}, []string{"Небо голуБое", "Анапа солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1, U: true}}, []string{"Пляж чистый"}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Q Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3}}, []string{"Небо и голуБое", "Пляж Я чистый", "Анапа т солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако А Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3, C: true}}, []string{"4 Небо и голуБое", "1 Пляж Я чистый", "3 Анапа т солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Д Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3, D: true}}, []string{"Небо и голуБое", "Анапа т солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Д Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3, U: true}}, []string{"Пляж Я чистый"}},
	}

	for _, item := range tableTests {
		result, err := Uniq(item.in.Strings, item.in.Option)
		if err != nil {
			t.Fatalf("%s", err)
		}
		require.Equal(t, item.out, result, "The two slices should be the same.")
	}

}

func TestTableParserFail(t *testing.T) {
	var tableTests = []struct {
		in  optiontest
		out []string
	}{
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{}}, []string{"I love music.", " ", "I love music of Kartik.",
			"Thanks."}},
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{C: true}}, []string{"1  ", "2 I love music of Kartik.",
			"1 Thanks.", "2 I love music of Kartik."}},
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{D: true}}, []string{"I love music of Kartik.", "I love music of Kartik."}},
		{optiontest{Strings: []string{"I love music.", "I love music.", "I love music.", " ", "I love music of Kartik.", "I love music of Kartik.",
			"Thanks.", "I love music of Kartik.", "I love music of Kartik."}, Option: Options{U: true}}, []string{"Thanks."}},
		{optiontest{Strings: []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", " ", "I love MuSIC of Kartik.", "I love music of kartik.",
			"Thanks.", "I love music of kartik.", "I love MuSIC of Kartik."}, Option: Options{I: true}}, []string{" ", "I love MuSIC of Kartik.", "Thanks.", "I love music of kartik."}},
		{optiontest{Strings: []string{"We love music.", "I love music.", "They love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
			"Thanks."}, Option: Options{F: 1}}, []string{"We love music.", "I love music of Kartik.", "Thanks."}},
		{optiontest{Strings: []string{"I love music.", "A love music.", "C love music.", " ", "I love music of Kartik.", "We love music of Kartik.",
			"Thanks."}, Option: Options{S: 1}}, []string{"I love music.", "I love music of Kartik.", "We love music of Kartik.", "Thanks."}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{}}, []string{"Небо", "Пляж", "Анапа"}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{C: true}}, []string{"2 Облако", "1 Пляж", "2 Анапа"}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{D: true}}, []string{"Облако", "Анапа"}},
		{optiontest{Strings: []string{"Небо", "Небо", "Облако", "Облако", "Пляж", "Анапа",
			"Анапа"}, Option: Options{U: true}}, []string{"Пляж", " "}},
		{optiontest{Strings: []string{"Небо голуБое", " ", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1}}, []string{"Небо голуБое", "Пляж чистый", "Анапа солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо голуБое", "Небо голубое", "Облако Голубое", " ", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1, C: true}}, []string{"4 Небо голуБое", " ", "1 Пляж чистый", "3 Анапа солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1, D: true}}, []string{"Небо голуБое", " ", "Анапа солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо голуБое", "Небо голубое", "Облако Голубое", "Облако ГолУбое", "Пляж чистый", "Анапа солнечнаЯ",
			"Анапа солНечная", "Сторона СолНечная"}, Option: Options{I: true, F: 1, U: true}}, []string{"Пляж чистый", " "}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Q Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3}}, []string{"Небо и голуБое", " ", "Пляж Я чистый", "Анапа т солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако А Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3, C: true}}, []string{"4 Небо и голуБое", " ", "1 Пляж Я чистый", "3 Анапа т солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Д Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3, D: true}}, []string{"Небо и голуБое", " ", "Анапа т солнечнаЯ"}},
		{optiontest{Strings: []string{"Небо и голуБое", "Небо я голубое", "Облако Б Голубое", "Облако Д Голубое", "Пляж Я чистый", "Анапа т солнечнаЯ",
			"Анапа г солНечная", "Сторона ф СолНечная"}, Option: Options{I: true, F: 1, S: 3, U: true}}, []string{"Пляж Я чистый", " "}},
	}

	for _, item := range tableTests {
		result, err := Uniq(item.in.Strings, item.in.Option)
		if err != nil {
			t.Fatalf("%s", err)
		}
		require.NotEqual(t, item.out, result, "The two slices should not be the same.")
	}

}
