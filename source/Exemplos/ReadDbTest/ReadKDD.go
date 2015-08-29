package main

import (
	"fmt"
	"os"
	"strings"
	"reflect"
	"strconv"
	"encoding/json"
)
// 

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
	num_access_files int
	num_outbound_cmds int
	is_host_login string
	is_guest_login string
	count int
	srv_count int
	serror_rate int
	srv_serror_rate int
	rerror_rate int
	srv_rerror_rate int
	same_srv_rate int
	diff_srv_rate int
	srv_diff_host_rate int
	dst_host_count int
	dst_host_srv_count int
	dst_host_same_srv_rate int
	dst_host_diff_srv_rate int
	dst_host_same_src_port_rate int
	dst_host_srv_diff_host_rate int
	dst_host_serror_rate int
	dst_host_srv_serror_rate int
	dst_host_rerror_rate int
	dst_host_srv_rerror_rate int
	attack string
}

type t struct {
        Duration int
    }

var DataBase []KDDCup
var totalreg int

func main() { 
	filename := "kddcup.data"
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

	b, err := json.Marshal(DataBase)
    if err != nil {
        fmt.Println(err)
        return
    }

    // open output file
    fo, err := os.Create("output.json")
    if err != nil {
        panic(err)
    }
    fo.WriteString(string(b))

	fmt.Printf("Total de registros: %d",totalreg)
}

func readline(f *os.File) string{
	ret :=""
	c := ""
	b := make([]byte, 1)

	// quan~to não for enter
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
            if f.CanSet() {
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
	}

	DataBase = append(DataBase, *newreg)
	totalreg = totalreg + 1
}

func readreg(l *string) string{
	i 	:= strings.Index(*l, ",")

	reg := (*l)[:i]
	*l 	= (*l)[i+1:]

	return reg
}

func checkerro(e error) {
    if e != nil {
        panic(e)
    }
}