package main

import (
	"fmt"
	"os"
	"strings"
	"reflect"
	"strconv"
	"log"
    "gopkg.in/mgo.v2"
)

/* ataques
back,
buffer_overflow,
ftp_write,
guess_passwd,
imap,ipsweep,
land,
loadmodule,
multihop,
neptune,
nmap,
normal,
perl,
phf,
pod,
portsweep,
rootkit,
satan,
smurf,
spy,
teardrop,
warezclient,
warez
*/
type KDDCup struct{
	Duration int
	Protocol_type string
	Service string
	Flag string
	Src_bytes int
	Dst_bytes int
	Land string
	Wrong_fragment int
	Urgent int
	Hot int
	Num_failed_logins int
	Logged_in string
	Num_compromised int
	Root_shell int
	Su_attempted int 
	Num_root int
	Num_file_creations int
	Num_shells int
	Num_access_files int
	Num_outbound_cmds int
	Is_host_login string
	Is_guest_login string
	Count int
	Srv_count int
	Serror_rate int
	Srv_serror_rate int
	Rerror_rate int
	Srv_rerror_rate int
	Same_srv_rate int
	Diff_srv_rate int
	Srv_diff_host_rate int
	Dst_host_count int
	Dst_host_srv_count int
	Dst_host_same_srv_rate int
	Dst_host_diff_srv_rate int
	Dst_host_same_src_port_rate int
	Dst_host_srv_diff_host_rate int
	Dst_host_serror_rate int
	Dst_host_srv_serror_rate int
	Dst_host_rerror_rate int
	Dst_host_srv_rerror_rate int
	Attack string
}

type t struct {
        Duration int
    }

var DataBase []KDDCup
var totalreg int

func main() { 
	server:="localhost"
	filename := "kddcup.data_10_percent_corrected"
	// faz a conexao com a base de dados
	session, err := mgo.Dial(server)
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    Colletion := session.DB("TCC").C("KDDCup")
	
	// faz a leitura do arquivo
	file,err := os.Open(filename)
	checkerro(err)
	totalreg = 0

	// faz a leitura da linha
	line:=readline(file)
	for line!="" {
		scanline(line)
		// faz a leitura da proxima linha
		line=readline(file)
	}
	for _,item := range DataBase {
  		err =Colletion.Insert(item)
  		if err != nil {
			log.Fatal(err)
	    }
	}

	fmt.Printf("Total de registros: %i",totalreg)
}

func readline(f *os.File) string{
	ret :=""
	c := ""
	b := make([]byte, 1)

	// quanto não for enter
	for c != "\n" {
		f.Read(b)

		// se não ler nada, arquivo acabou, então sai do loop
		if b[0] == 0{
			break
		}

		c = string(b)
		ret += c
	}
	return ret
}

func scanline(l string){
	newreg := new(KDDCup)

	val := reflect.ValueOf(newreg).Elem()
	
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)

        f := val.FieldByName(typeField.Name)        

        if f.IsValid() {
            if f.Kind() == reflect.Int {
                x,_ := strconv.Atoi(readreg(&l))
                f.SetInt(int64(x))
            }

            if f.Kind() == reflect.String {
                x := readreg(&l)
                 f.SetString(x)
            }
        }
	}
	
	DataBase = append(DataBase, *newreg)
	totalreg = totalreg + 1
}

func readreg(l *string) string{
	i 	:= strings.Index(*l, ",")

	// valida final de arquivo
	if i < 0 {
		i=len(*l)-2
	}

	reg := (*l)[:i]
	*l 	= (*l)[i+1:]

	return reg
}

func checkerro(e error) {
    if e != nil {
        panic(e)
    }
}