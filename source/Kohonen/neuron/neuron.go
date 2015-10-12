package kohonen

import (
    "math"
    //"fmt"
    //"bufio"
    //"os"
)

type Neuron struct {
    Weights []float64
    ExitWeights []float64
    RGB []int
    X,Y,Length int
    nf float64
    maxint float64
    maxvar float64
}

func (r Neuron) Create(x int, y int, l int,maxite int) Neuron{
    r.X = x
    r.Y = y
    r.Length = l
    r.maxint = float64(maxite)
    r.maxvar = 0.05

    dl:=float64(l)

    log:=math.Log(dl)
    r.nf = float64(r.maxint) / log
    
    return r
}

func (r Neuron) Gauss(it int,le float64,dist float64) float64 {
    dit:=float64(it)

    return math.Exp(-0.1*math.Pow(dist, 2) / (2 * le * dit))
}

func (r Neuron) Strength(it int) float64 {
    dit:=float64(it)
    dl:=float64(r.Length / 2)

    result:= math.Exp(-dit / r.nf) * dl

    return result
}

func (r Neuron) LearningRate(it int) float64 {
    dit:=float64(it)

    return math.Exp(-1.0*dit / r.nf) * r.maxvar
}

func (r Neuron) UpdateWeigths(pattern,Exitpattern []float64,winner Neuron,it int) (Neuron) {
    
    le:=r.Strength(it)
    dx:=float64(winner.X - r.X)
    dy:=float64(winner.Y - r.Y)
    //fmt.Printf("wx:%d wy:%d x:%d y:%d\n",winner.X,winner.Y,r.X,r.Y)

    dist:=math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
    //fmt.Printf("dist:%g le:%g\n",dist,le)
    if(dist < le){
        Gau:=r.Gauss(it,le,dist)
        Lea:=r.LearningRate(it)
        //fmt.Printf("Lea:%g Str:%g Gau:%g\n",Lea,le,Gau)
        //fmt.Print("Press 'Enter' to continue...")
        //bufio.NewReader(os.Stdin).ReadBytes('\n') 

        for i := 0; i < len(r.Weights); i++ {
            delta:= Lea * Gau * (pattern[i] - r.Weights[i])

            r.Weights[i]+=delta

            r.RGB[i] = int((r.Weights[i] * 255))

            if r.RGB[i] > 255 {
                r.RGB[i] = 255
            }
        }

        for i := 0; i < len(r.ExitWeights); i++ {
            delta:= Lea * Gau * (Exitpattern[i] - r.ExitWeights[i])

            r.ExitWeights[i]+=delta
        }
    }
    return r
}