package modules

import (
	"fmt"

	"github.com/RadhiFadlillah/go-sastrawi"
)

func Chat() string {
	// Kalimat asal
	sentence := "Rakyat memenuhi halaman gedung untuk menyuarakan isi hatinya. Baca berita selengkapnya di http://www.kompas.com."
	// Ubah kata berimbuhan menjadi kata dasar
	dictionary := sastrawi.DefaultDictionary()
	stemmer := sastrawi.NewStemmer(dictionary)
	for _, word := range sastrawi.Tokenize(sentence) {
		fmt.Printf("%s => %s\n", word, stemmer.Stem(word))
	}
	return sentence
}
