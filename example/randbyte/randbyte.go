package randbyte

import (
	"encoding/binary"
	"io"
	"math/rand"
)

type generator struct {
	rnd rand.Source // Генератор случайных чисел. Вообще rand.Rand уже реализует интерфейс io.Reader, но для примера мы реализуем его самостоятельно.
}

// New — обратите внимание, что мы возвращаем generator, присвоенный интерфейсу io.Reader, сама структура generator неэкспортируемая.
// Мы скрыли внутри пакета все детали.
func New(seed int64) io.Reader {
	return &generator{
		rnd: rand.NewSource(seed),
	}
}

// Read — реализация io.Reader
func (g *generator) Read(bytes []byte) (n int, err error) { // error — это тип ошибки, подробнее мы рассмотрим его в следующем разделе.
	// for min := range len(bytes) / 8 {
	// min *= 8
	for min := 0; min+8 < len(bytes); min += 8 {
		binary.LittleEndian.PutUint64(bytes[min:min+8], uint64(g.rnd.Int63()))
	}
	return len(bytes), nil
}
