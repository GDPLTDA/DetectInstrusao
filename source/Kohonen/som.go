package main

import (
    //"./somstructs"
    somf "./somfunctions"
    "fmt"
    "flag"
    "testing"
    "os"
    "bufio"
)
var Test bool
var Loadtype int
var Filename, Savename string
var Saved, Train bool
var PngBefore,PngAfter string

func main() {
    LoadParams()
    PngBefore = "Before.png"
    PngAfter = "After.png"
    
    if Test {
        result := testing.Benchmark(ExecuteTest)
    
        seconds := float64(result.T.Seconds()) / float64(result.N)
        fmt.Printf("%13.2f s\n\n", seconds)
    }

    if !Test  {
        Execute()
    }
    ShowParams()

    var num float64
    
    for !Saved {
        var a []float64
        a = append(a, 0,0,0)
        for i := 0; i < somf.Dimensions; i++ {
            fmt.Printf("%d >> ", i)
            n, err := fmt.Scanf("%g\n", &num)
            if err != nil {
                fmt.Println(n, err)
            }
            
            a[i] = num
        }
        somf.Koh.Test(a)
    }
}
func ExecuteTest(b *testing.B) {
    Execute()
}

func Execute(){
    // faz a leitura dos dados de treinamento
    if Loadtype == 0 {
        somf.Koh = somf.LoadFile(Filename)
    }
    if Loadtype == 1 {
        somf.Koh = somf.LoadKDDCup("KDDCup")
    }
    if Loadtype == 2 {
        somf.Koh = somf.LoadJson(Filename)
    }

    if somf.Koh.Empty() {
        return
    }
    // Desenha o estado atual da grade antes do treino
    somf.Koh.Draw(PngBefore)

    if Train {
        // faz o treinamento da base de dados
        somf.Koh = somf.Koh.Train()
    }
    // Desenha o estado atual da grade depois do treino
    somf.Koh.Draw(PngAfter)

    // verifica se deve salvar o treinamento
    if Saved{
        somf.SaveJson(Savename)
    }

}

func LoadParams(){
    flag.StringVar(&somf.Server,"server", "localhost", "Server name")
    flag.StringVar(&somf.Dbname,"base", "TCC", "Data base name")
    flag.IntVar(&somf.Gridsize,"grid", 10, "Grid Size")
    flag.IntVar(&somf.Dimensions,"dim", 3, "Dimensions Weigths")
    flag.IntVar(&somf.Interactions,"ite", 5000, "Iteractions")
    flag.Float64Var(&somf.TxVar,"var", 0.5, "Taxa Variation")

    flag.BoolVar(&Saved,"s", false, "Save?")
    flag.BoolVar(&Train,"t", false, "Train?")
    flag.StringVar(&Savename,"sname", "", "Save file name")
    flag.IntVar(&Loadtype,"type", 0, "0-Load file, 1-Load KddCup, 2-Json File")
    flag.StringVar(&Filename,"f", "", "File name")
    flag.BoolVar(&Test,"test", false, "Test time")

    config:= flag.String("config", "", "Config file")

    flag.Parse()
    
    
    // usando arquivo de configuracao
    if *config!="" {
        fmt.Println("-Config:", *config)
        file,err := os.Open(*config)
        somf.Checkerro(err)

        reader := bufio.NewReader(file)
        scanner := bufio.NewScanner(reader)
        for scanner.Scan() {
            line:=scanner.Text()
            fmt.Println("--"+line)
        }
    }

    if  Train {
        fmt.Println("Trainning...")
    } else {
        fmt.Println("Loading...")
    }
    
}

func ShowParams() {
    fmt.Println("Params")
    fmt.Println("-Type:", Loadtype)

    if Loadtype == 0{
        fmt.Println("-File:", Filename)
    }

    if Loadtype == 1{
        fmt.Println("-Server:", somf.Server)
        fmt.Println("-DataBase:", somf.Server)
    }

    if Loadtype == 2{
        fmt.Println("-Json:", Filename)
    }

    fmt.Println("-Grid Size:", somf.Koh.Gridsize)
    fmt.Println("-Interactions:", somf.Koh.Interactions)
    fmt.Println("-Variation:", somf.Koh.TxVar)
    fmt.Println("-Save?:", Saved)

    if Saved{
        fmt.Println("  -File name:", Savename)
    }

}