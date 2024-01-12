package main

import (
"fmt"
"os"
)

type Furniture struct {
	Name string	
	Height int
	Width int
	Depth int
	Material string
}

type Appliances struct { 
	Name string
	Height int
	Width int
	Depth int
	Material string
}

func writeToFile(f Furniture) error {
	fileName := f.Name + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "Высота: %d см\nШирина: %d см\nГлубина: %d см\nМатериал: %s", f.Height, f.Width, f.Depth, f.Material)
	if err != nil {
		return err
	}
	return nil
}

func writeToFile_1(v Appliances) error {
	fileName := v.Name + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		
	}
	defer file.Close()
	
	_, err = fmt.Fprintf(file, "Высота: %d см\nШирина: %d см\nГлубина: %d см\nМатериал: %s", v.Height, v.Width, v.Depth, v.Material)
	if err != nil {
		
	}
	return nil
	}

func main() {

furniture := []Furniture{
	Furniture{Name: "Кровать", Height: 200, Width: 160, Depth: 220, Material: "Дерево"},
	Furniture{Name: "Шкаф", Height: 220, Width: 120, Depth: 50, Material: "ДСП"},
	Furniture{Name: "Стол", Height: 75, Width: 120, Depth: 70, Material: "Стекло"},
}

appliances:= []Appliances{
	Appliances{Name: "Тостер", Height: 200, Width: 160, Depth: 220, Material: "Алюминий"},
	Appliances{Name: "Телевизор", Height: 220, Width: 120, Depth: 50, Material: "Пластик"},
	Appliances{Name: "Холодильник", Height: 75, Width: 120, Depth: 70, Material: "Металл"},
}

for _, f := range furniture {
	err := writeToFile(f)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Файл: %s.txt не был записан!\n", f.Name)
	} else {
		fmt.Printf("Файл: %s.txt записан успешно!\n", f.Name)
	}
}

for _, v := range appliances {
	err := writeToFile(Furniture(v)) 
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Файл: %s.txt не был записан!\n", v.Name)
	} else {
	fmt.Printf("Файл: %s.txt записан успешно!\n", v.Name)
	}
}

}


