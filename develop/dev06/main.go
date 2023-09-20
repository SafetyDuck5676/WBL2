package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Определение флагов
	fieldsPtr := flag.String("f", "1", "fields to select")
	delimiterPtr := flag.String("d", "\t", "delimiter")
	separatedPtr := flag.Bool("s", false, "only separated lines")

	// Парсинг флагов командной строки
	flag.Parse()

	// Чтение строк из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка строки на наличие разделителя
		if strings.Contains(line, *delimiterPtr) || !*separatedPtr {
			// Разделение строки на поля
			fields := strings.Split(line, *delimiterPtr)

			// Выбор запрошенных полей
			selectedFields := make([]string, 0)
			if *fieldsPtr == "" {
				// Если не указаны поля, выводим всю строку
				selectedFields = fields
			} else {
				// Иначе выбираем запрошенные поля
				fieldIndexes := strings.Split(*fieldsPtr, ",")
				for _, index := range fieldIndexes {
					selectedFields = append(selectedFields, fields[toInt(index)-1])
				}
			}

			// Вывод выбранных полей
			fmt.Println(strings.Join(selectedFields, *delimiterPtr))
		}
	}

	// Проверка ошибок сканнера
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Преобразование строки в целое число
func toInt(s string) int {
	var res int
	for _, c := range s {
		res = res*10 + int(c-'0')
	}
	return res
}
