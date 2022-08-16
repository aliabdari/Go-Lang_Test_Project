package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Rectangles struct {
	Main struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"main"`
	Input []struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"input"`
}

func check_rectangles(obj *Rectangles) bool {
	// fmt.Println("check_rectangles data:", reflect.TypeOf(obj.Main))
	for i := 0; i < len(obj.Input); i++ {
		if is_overlap(obj.Main, obj.Input[i]) {
			fmt.Println("is_overlap data:", obj.Input[i])
			write_to_file()
		}
	}
	return true
}

func is_overlap(rect1, rect2 Rect) bool {
	// fmt.Println("is_overlap data:", rect1)
	// fmt.Println("is_overlap data:", rect2)

	if (rect1.X >= rect2.X+rect2.Width) || (rect1.X+rect1.Width <= rect2.X) || (rect1.Y+rect1.Height <= rect2.Y) || (rect1.Y >= rect2.Y+rect2.Height) {
		return false
	}

	return true
}

func write_to_file() bool {
	return true
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data := &Rectangles{}
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("got data:", reflect.TypeOf(data))
		check_rectangles(data)
		w.WriteHeader(http.StatusCreated)
	})

	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}

}