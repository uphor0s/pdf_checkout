package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func generate_invoice(invoice_data Invoice) string {
	file_path := fmt.Sprintf("pdfs/N%v %s %s.pdf", loadCounter(), invoice_data.Buyer_name, time.Now().Format("02.01.2006 15-04-05"))
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.AddUTF8Font("dejavu", consts.Normal, "fonts/dejavu/DejaVuSans.ttf")
	m.AddUTF8Font("dejavu", consts.Bold, "fonts/dejavu/DejaVuSans-Bold.ttf")

	m.SetDefaultFontFamily("dejavu")

	m.SetPageMargins(20, 10, 20)
	buildFruitList(invoice_data, m)
	err := m.OutputFileAndClose(file_path)
	if err != nil {
		log.Panic(err)
	}
	increment()
	output, _ := filepath.Abs(file_path)
	return output
}

func buildFruitList(data Invoice, m pdf.Maroto) {
	// company_data := map[string]string{
	// 	"ИНН":         "6164147042",
	// 	"КПП":         "616401001",
	// 	"Номер счёта": "40702810626340000613",
	// 	"Банк":        "ФИЛИАЛ \"РОСТОВСКИЙ\" АО \"АЛЬФА-БАНК\"",
	// 	"БИК":         "046015207",
	// 	"Корреспондентский счёт": "30101810500000000207",
	// }
	var sum float64
	tableHeadings := []string{"№", "Наименование", "Стоимость", "Количество", "Сумма"}
	contents := [][]string{}
	for i, item := range data.Contents {
		contents = append(contents, []string{strconv.Itoa((i + 1)), item.Name, fmt.Sprintf("%.2f", item.Price), fmt.Sprintf("%d", item.Count), fmt.Sprintf("%.2f", item.Price*float64(item.Count))})
		sum += item.Price * float64(item.Count)

		// fmt.Println(item.Price, item.Count, item.Price*float64(item.Count))
		// fmt.Printf("%f\n", sum)
	}
	contents = append(contents, []string{"", "", "", "Итого", fmt.Sprintf("%.2f", sum)})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("ООО АйСофтПро", props.Text{
				Top:  2,
				Size: 15,

				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(10, func() {
			m.Text("344002, РОССИЯ, Ростовская область, Ростов-на-Дону, ул. Социалистическая, д. 74, офис 500, помещ. 11", props.Text{
				Top:             2,
				Size:            12,
				VerticalPadding: 2,
				Style:           consts.Bold,
				Align:           consts.Left,
			})
		})
	})

	m.Row(10, func() {})

	for _, p := range company_data {
		m.Row(10, func() {
			m.Col(5, func() {
				m.Text(p.key, props.Text{
					Top:  2,
					Size: 12,

					Style: consts.Normal,
					Align: consts.Left,
				})
			})
			m.Col(10, func() {
				m.Text(p.value, props.Text{
					Top:  2,
					Size: 12,

					Style: consts.Normal,
					Align: consts.Left,
				})
			})
		})
	}
	m.Row(10, func() {})
	m.Row(10, func() {
		m.Col(5, func() {
			m.Text("Плательщик:", props.Text{
				Top:  2,
				Size: 12,

				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(10, func() {
			m.Text(data.Buyer_name, props.Text{
				Top:  2,
				Size: 12,

				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})

	m.Row(10, func() {})

	m.SetBackgroundColor(color.NewWhite())
	m.Row(10, func() {})
	m.Row(10, func() {
		m.Col(11, func() {
			m.Text(fmt.Sprint("Счет №", loadCounter(), " от ", time.Now().Format("02.01.2006")), props.Text{
				Top:  2,
				Size: 15,

				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.Row(10, func() {})
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      12,
			GridSizes: []uint{1, 4, 2, 3, 2},
		},
		ContentProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{1, 4, 2, 3, 2},
		},
		Align:                  consts.Center,
		VerticalContentPadding: 6,
		HeaderContentSpace:     3,
		Line:                   true,
	})
	m.Row(10, func() {

		m.Col(6, func() {})

		m.Col(6, func() {
			m.Text(fmt.Sprintf("К оплате %.2f руб.", sum), props.Text{
				Top:  2,
				Size: 14,

				Style: consts.Bold,
				Align: consts.Right,
			})
		})
	})

	m.Row(10, func() {

		m.Col(12, func() {
			m.Text(cases.Title(language.Russian).String(numToStr(sum)), props.Text{
				Top:  2,
				Size: 12,

				Style: consts.Bold,
				Align: consts.Right,
			})
		})

	})

	// m.Row(50, func() {
	// 	m.Col(20, func() {
	// 		_ = m.FileImage("seal.png", props.Rect{
	// 			Left:    30,
	// 			Center:  true,
	// 			Percent: 100,
	// 		})
	// 	})

	// })
}
