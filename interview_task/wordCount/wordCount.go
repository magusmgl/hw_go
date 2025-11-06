package wordcount

type WordCounter struct {
	Order  []string
	Counts map[string]int
	Limit  int
}

func NewWordCounter(limit int) *WordCounter {
	return &WordCounter{
		Counts: make(map[string]int),
		Limit:  limit,
	}
}

func (wc *WordCounter) CountWord(word string) {
	if _, ok := wc.Counts[word]; !ok {
		wc.Order = append(wc.Order, word)
	}

	wc.Counts[word]++

	if len(wc.Counts) > wc.Limit {
		delete(wc.Counts, wc.Order[0])
		wc.Order = wc.Order[1:] //Удаление
	}
}
