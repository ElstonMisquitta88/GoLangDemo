package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

func (t *Todo) MarkDone() {
	t.Done = true
}

var sqldbconn *sql.DB

type Food struct {
	Id          int
	Title       string
	Description string
	Price       float64
}

func sqlconnect() *sql.DB {
	cfg, err := loadConfig("config.json")
	cfg.DB.Password = "phoenixuser1"
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	connString := fmt.Sprintf(
		"server=%s;port=%d;user id=%s;password=%s;database=%s;encrypt=disable",
		cfg.DB.Server,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Database,
	)

	// Open connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	// Ping DB to confirm connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Ping failed:", err.Error())
	}

	fmt.Println("Connected to SQL Server successfully!")
	return db
}

type DBConfig struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type AppConfig struct {
	DB DBConfig `json:"db"`
}

func loadConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
func getFoodList() ([]Food, error) {
	query := `SELECT Id, Title, [Description], Price FROM Food`

	rows, err := sqldbconn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	foods := []Food{} // initialize slice

	for rows.Next() {
		var f Food
		err := rows.Scan(&f.Id, &f.Title, &f.Description, &f.Price)
		if err != nil {
			return nil, err
		}
		foods = append(foods, f)
	}
	return foods, rows.Err()
}

func main() {

	sqldbconn = sqlconnect() // your working function
	defer sqldbconn.Close()

	allfood, err := getFoodList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(allfood)

	/*todo := Todo{ID: 1, Title: "Learn Go", Done: false}
	todo.MarkDone()
	fmt.Println(todo.Done) // true

	x := 42
	p := &x // p is a pointer to x
	fmt.Printf("The value of x is: %d\n", x)
	fmt.Println("Value of Pointer ", *p)
	*/

	/* filePath := "D:\\SampleFile\\test.txt" // Replace with your file path
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File content:\n%s\n", content) */

	/*var wg sync.WaitGroup
	wg.Go(func() {
		say("Processing task")
		exportfile()
		say("Processing Next task")
	})
	wg.Wait()
	fmt.Println("All tasks complete")
	*/

	/* 	fmt.Println("Hello, Go!")

	   	s1 := "Elston "
	   	s1 = "Welcome to " + s1
	   	fmt.Println(s1)

	   	say("Hello Go!")
	   	conditionalExample(10)
	   	demoarray()

	   	r := rect{width: 3, height: 4}
	   	measure(r)

	   	x := 10
	   	p := &x // p is a pointer to x

	   	fmt.Println("x =", x)   // 10
	   	fmt.Println("p =", p)   // memory address
	   	fmt.Println("*p =", *p) // 10 (value at that memory address)

	   	x, y := vals(3, 7)
	   	fmt.Println("X Value :", x)
	   	fmt.Println("Y Value :", y) */
}

func exportfile() {
	records := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "24", "London"},
		{"Charlie", "35", "Paris"},
	}

	// Create a new CSV file
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close() // Ensure the file is closed

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all buffered data is written to the file

	// Write all records to the CSV file
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalf("failed to write record: %v", err)
		}
	}

	log.Println(" Data successfully exported to output.csv")
}

func vals(a int, b int) (int, int) {
	return a + 1, b + 1
}

func demoarray() {
	//Arrays
	arr := [3]int{10, 20, 30}
	fmt.Println("Array elements are:", arr[0], arr[1], arr[2])

	//Slice
	arr1 := []int{40, 50, 60}
	fmt.Println("Array elements are:", arr1[0], arr1[1], arr1[2])
}

func conditionalExample(valnum int) {
	{
		if valnum > 5 {
			fmt.Println("Value is greater than 5")
		} else {
			fmt.Println("Value is less than or equal to 5")
		}
	}

}

func say(message string) {
	fmt.Println("You said: ", message)
}
