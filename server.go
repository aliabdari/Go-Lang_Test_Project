package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// the struct of rectangle with time, used for output information
type Rect_time struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Time   string `json:"time"`
}

// the struct of a Rectangle,used for Input information
type Rect struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// the struct of received data
type Rectangles struct {
	Main  Rect   `json:"main"`
	Input []Rect `json:"input"`
}

// investigate all of the received information
func check_rectangles(obj *Rectangles) {

	var rect_time_list = []Rect_time{}
	json.Unmarshal([]byte(read_from_file()), &rect_time_list)

	for i := 0; i < len(obj.Input); i++ {
		if is_overlap(obj.Main, obj.Input[i]) {
			// fmt.Println("is_overlap data:", obj.Input[i])
			currentTime := time.Now()

			rect_time_obj := Rect_time{}
			rect_time_obj.X = obj.Input[i].X
			rect_time_obj.Y = obj.Input[i].Y
			rect_time_obj.Width = obj.Input[i].Width
			rect_time_obj.Height = obj.Input[i].Height
			rect_time_obj.Time = currentTime.Format("2006-01-02 3:4:5 PM")

			rect_time_list = append(rect_time_list, rect_time_obj)
		}
	}
	write_to_file(rect_time_list)
}

// check if two rectangles have mutual area or not
func is_overlap(rect1, rect2 Rect) bool {
	if (rect1.X >= rect2.X+rect2.Width) || (rect1.X+rect1.Width <= rect2.X) || (rect1.Y+rect1.Height <= rect2.Y) || (rect1.Y >= rect2.Y+rect2.Height) {
		return false
	}

	return true
}

// write acceptable information in a txt file
func write_to_file(rect_time_list []Rect_time) {

	file, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	data, _ := json.Marshal(rect_time_list)
	// fmt.Println(string(data))

	if _, err := file.WriteString(string(data)); err != nil {
		log.Fatal(err)
	}
}

// read data file in order to return saved information
func read_from_file() string {
	file, err := os.OpenFile("data.txt", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var string_json string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		string_json = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return string_json
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			data := &Rectangles{}
			err := json.NewDecoder(r.Body).Decode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			check_rectangles(data)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(string("Data processing get started")))

		} else if r.Method == http.MethodGet {
			output := read_from_file()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(output))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	})

	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}

}
