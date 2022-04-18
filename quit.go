package main

import (
  "fmt"
  "flag"
  "os"
  "encoding/csv"
  //"github.com/shogo82148/go-shuffle"
  "math/rand"
  "time"
)

func main()  {
  var shufflee string
  csv_file := flag.String("f", "efj.csv", "read file path with -f") //bunch of flags
  flag.StringVar(&shufflee, "s", "", "shuffle the entries in your file, you have to pass double quotes")
  flag.Parse()


  file, err := os.Open(*csv_file) // opening the file
  if err != nil {
    exit(fmt.Sprintf("error opening file!: %s", *csv_file))
    return
  }

  read := csv.NewReader(file) //creating new reader and reading all the data from file
  lines, err := read.ReadAll()
  if err != nil{
    fmt.Println("failed to parse file!")
  }
  //shuffle.Slice(lines) shuffles for only 1 time
  if shufflee == "" {
  r := rand.New(rand.NewSource(time.Now().Unix())) // shuffle files everytime the program runs
  r.Shuffle((len(lines)), func(i, j int){
    lines[i], lines[j] = lines[j], lines[i]
  })

  problems := parseLines(lines)

  correct := 0 // correct answers, default 0
  for i, p := range problems{
    fmt.Printf("problem #%d : %s = \n", i+1, p.q) //p.q are the questions, i is the index
    var answer string
    fmt.Scanf("%s\n", &answer)
    if answer == p.a {
        correct++
      } // p.a are the answers
  }
  fmt.Printf("Your score: %d out of %d \n", correct, len(problems))
}
}
func exit(message string){
  fmt.Println(message)

  os.Exit(1)
}

func parseLines(lines [][]string) []problems { //parsing lines
  ret := make([]problems, len(lines))
  for i, lines := range lines{
    ret[i] = problems{
      q: lines[0],
      a: lines[1],
    }
    }
    return ret
}

type problems struct{
  q string
  a string
}
