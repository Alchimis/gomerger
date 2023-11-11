package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type строка struct {
	string
}

type логическоеЗначене struct {
	bool
}

func новаяОшибка(стр строка) error {
	return errors.New(стр.string)
}

func новыйФайл(стр строка) (*os.File, error) {
	return os.Open(стр.string)
}

func функция() {

}

func вывестиСтроку(стр строка) {
	fmt.Println(стр.string)
}

/*
Задача: есть много файлов с похожим названием типо кот-1.png, кот-2.png, кот-3.png,
	Программа должна поместить эти файла в папку с названием кот.
Возможное решение: Проходим по названиям всех фалов в папке, выделяем название до цыферок
и складываем в папку
*/

func получитьИнформациюОФайле(путьКФайлу строка) (fs.FileInfo, error) {
	return os.Stat(путьКФайлу.string)
}

func неРавно(значениеОдин interface{}, значениДва interface{}) логическоеЗначене {
	return логическоеЗначене{значениеОдин != значениДва}
}

func файлЯвляетьсяПапкой(путьКПапке строка) логическоеЗначене {
	информацияОФале, ошибка := получитьИнформациюОФайле(путьКПапке)
	if неРавно(ошибка, nil).bool {
		return логическоеЗначене{false}
	}
	return логическоеЗначене{информацияОФале.IsDir()}
}

func открытьПапку(путьКПапке строка) (файл *os.File, ошибка error) {
	файл = nil
	ошибка = nil
	if !файлЯвляетьсяПапкой(путьКПапке).bool {
		ошибка = новаяОшибка(строка{"Файл не являеться папкой"})
		return
	}
	файл, ошибка = новыйФайл(путьКПапке)
	return
}

func вывестиОшибку(ошибка error) {
	fmt.Println(ошибка)
}

func вывестиЧтоУгодно(чтоУгодно interface{}) {
	fmt.Println(чтоУгодно)
}

type ТаблицаИмёнФайлов struct {
	Таблица map[строка][]строка
}

func (таблицаИмёнФалов *ТаблицаИмёнФайлов) Укамплектовать(ключь, значение строка) {
	массивСтрок, нашёл := таблицаИмёнФалов.Таблица[ключь]
	if !нашёл { // если записи нет то добавляем запись
		таблицаИмёнФалов.Таблица[ключь] = []строка{
			значение,
		}
	} else { // обновляем записи добавлением нового хначения
		таблицаИмёнФалов.Таблица[ключь] = append(массивСтрок, значение)
	}
}

func примерМапы() {

	/*
		в map можно деалть записи по типу

		карта := make(map[string]string)
		карта["кот"] = "васька"

		мы добавил  в мап запись кот - васька
		в дальнейшем если мы обратимся к записи "кот" то мы получим "васька"

	*/
	карта := make(map[string]string)
	// добавили запись кот - васька
	карта["кот"] = "васька"
	// достали запись по ключу кот. в итоге получили ваську
	вывестиСтроку(строка{карта["кот"]})
}

func получитьКлючевоеНазваниеФайла(названиеФайла строка) строка {
	// отрезает суфикс у файла. то есть расширение. например "кот.png" превратиться в "кот"
	расширениеФайла := filepath.Ext(названиеФайла.string)
	названиеФалаБезРасшиения, _ := strings.CutSuffix(названиеФайла.string, расширениеФайла) // функция дополнительно оповещает, нашла ли она суффикс, но мне это не надо

	// ищём есть ли в строке цифры
	индексПервогоЧислаВСтроке := strings.IndexFunc(названиеФалаБезРасшиения, func(руна rune) bool {
		return руна < '9' && руна > '0'
	})

	// если не нашёл ни одного числа то функция вернёт посто название файла без расширения
	if индексПервогоЧислаВСтроке == -1 {
		return строка{названиеФалаБезРасшиения}
	}
	// если первое число найдено то отрезает всё до числа
	return строка{названиеФалаБезРасшиения[:индексПервогоЧислаВСтроке]}
}

func main() {
	путьКПапке := строка{"путь к папке"}
	открытаяПапка, ошибка := открытьПапку(путьКПапке)
	if ошибка != nil {
		вывестиОшибку(ошибка)
		return
	}
	вывестиЧтоУгодно(открытаяПапка)
	вывестиСтроку(строка{"вот папка"})
	назвнияФайлов, ошибка := открытаяПапка.Readdir(0) //
	if ошибка != nil {
		вывестиОшибку(ошибка)
		return
	}
	//вывестиЧтоУгодно(назвнияФайлов)

	таблицаИмёнФайлов := ТаблицаИмёнФайлов{
		make(map[строка][]строка),
	}
	fmt.Println(назвнияФайлов)
	for _, информацияОФайле := range назвнияФайлов {
		обёрткаНазванияФайла := строка{информацияОФайле.Name()}
		ключевоеназваниеФайла := получитьКлючевоеНазваниеФайла(обёрткаНазванияФайла)
		таблицаИмёнФайлов.Укамплектовать(ключевоеназваниеФайла, обёрткаНазванияФайла)
	}
	вывестиЧтоУгодно(таблицаИмёнФайлов.Таблица)
	вывестиЧтоУгодно(открытаяПапка.Name())

	fmt.Println(filepath.Dir(открытаяПапка.Name()))
	//os.Rename()
	for клбчь := range таблицаИмёнФайлов.Таблица {
		if ошибка = os.Mkdir(открытаяПапка.Name()+"\\"+клбчь.string, 0755); ошибка != nil {
			вывестиЧтоУгодно(ошибка)
		} else {
			for _, названиеФайла := range таблицаИмёнФайлов.Таблица[клбчь] {
				os.Create(клбчь.string + "\\" + названиеФайла.string)
				os.Rename(открытаяПапка.Name()+"\\"+названиеФайла.string, открытаяПапка.Name()+"\\"+клбчь.string+"\\"+названиеФайла.string)
			}
		}
	}
}
