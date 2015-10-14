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
var Filename, Trainname string
var Savetrain bool
var PngBefore,PngAfter string

func main() {
    LoadParams()
    PngBefore = "Before.png"
    PngAfter = "After.png"
    
    //if Test {
        result := testing.Benchmark(Execute)
    //}
    nanoseconds := float64(result.T.Nanoseconds()) / float64(result.N)
    milliseconds := nanoseconds / 1000000.0
     
    fmt.Printf("%13.2f ns/op | %13.10f ms/op | %d Iterations\n", nanoseconds, milliseconds, result.N)

    var num float64
    
    for true {
        var a []float64
        a = append(a, 0,0,0)
        for i := 0; i < 3; i++ {
            fmt.Printf("%d >> ", i)
            n, err := fmt.Scanf("%g\n", &num)
            if err != nil {
                fmt.Println(n, err)
            }
            
            a[i] = num
        }
        somf.Koh.Test(a)
    }
    //if(Test){
       //Execute()
    //}
}

func Execute(b *testing.B){
    // faz a leitura dos dados de treinamento
    if Loadtype == 0 {
        somf.Koh = somf.LoadFile(Filename)
    }
    if Loadtype == 1 {
        somf.Koh = somf.LoadKDDCup("KDDCup")
    }
    if Loadtype == 2 {
        //Koh = LoadJson()
    }

    // Desenha o estado atual da grade antes do treino
    somf.Koh.Draw(PngBefore)
    // faz o treinamento da base de dados
    somf.Koh = somf.Koh.Train()
    // Desenha o estado atual da grade depois do treino
    somf.Koh.Draw(PngAfter)

    // verifica se deve salvar o treinamento
    if Savetrain{
        somf.SaveTrainJson(Trainname)
    }
}

func LoadParams(){
    fmt.Println("Params")
    flag.StringVar(&somf.Server,"server", "localhost", "Server name")
    flag.StringVar(&somf.Dbname,"base", "TCC", "Data base name")
    flag.IntVar(&somf.Gridsize,"grid", 10, "Grid Size")
    flag.IntVar(&somf.Dimensions,"dim", 3, "Dimensions Weigths")
    flag.IntVar(&somf.Interactions,"ite", 5000, "Iteractions")
    flag.Float64Var(&somf.TxVar,"var", 0.5, "Taxa Variation")

    flag.BoolVar(&Savetrain,"s", false, "Save Training?")
    flag.StringVar(&Trainname,"train", "GridTrain", "Training name")
    flag.IntVar(&Loadtype,"type", 0, "0-Load file,1-Load KddCup,2-Json File")
    flag.StringVar(&Filename,"f", "", "File name")
    flag.BoolVar(&Test,"t", false, "Test time")

    config:= flag.String("config", "", "Config file")

    flag.Parse()
    
    // passando parametro
    if *config=="" {
        fmt.Println("-type:", Loadtype)

        if Loadtype == 0{
            fmt.Println("-File:", Filename)
        }

        if Loadtype == 1{
            fmt.Println("-Server:", somf.Server)
            fmt.Println("-DataBase:", somf.Server)
        }

        fmt.Println("-Grid Size:", somf.Gridsize)
        fmt.Println("-Interactions:", somf.Interactions)
        fmt.Println("-Variation:", somf.TxVar)
        fmt.Println("-Save Training?:", Savetrain)

        if Savetrain{
            fmt.Println("-Training name:", Trainname)
        }
    }
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
    fmt.Println("Trainning...")
}