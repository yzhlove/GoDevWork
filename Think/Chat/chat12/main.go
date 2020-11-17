package main

import "fmt"

func main() {

	hc := create()
	pg := hc.Get(1)
	pg.Points = 12138
	pg.Name = "create page"
	hc.show()
}

type Hero struct {
	Nick  string
	Pages []*Page
}

type Page struct {
	Name   string
	Points uint32
}

func (h *Hero) Get(i int) *Page {
	return h.Pages[i]
}

func create() (h *Hero) {
	h = &Hero{}
	h.Nick = "nick"
	h.Pages = append(h.Pages, &Page{Name: "one", Points: 12})
	h.Pages = append(h.Pages, &Page{Name: "two", Points: 23})
	h.Pages = append(h.Pages, &Page{Name: "three", Points: 34})
	return
}

func (h *Hero) show() {
	fmt.Println("==", h.Nick)
	for _, p := range h.Pages {
		fmt.Println("====", p.Name, p.Points)
	}
}
