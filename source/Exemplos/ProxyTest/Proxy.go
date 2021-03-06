package main 

import ( 
        "net" 
        "fmt" 
        "os"
        "encoding/hex"
        "bytes"
) 

func main() { 
      
        localAddr := ":8090"
        remoteAddr := "127.0.0.1:9090" 

        local, err := net.Listen("tcp", localAddr) 
        if local == nil { 
                fatal("cannot listen: %v", err) 
        } 

        for { 
                conn, err := local.Accept() 
                        
                if conn == nil { 
                        fatal("accept failed: %v", err) 
                }                 
                
                go forward(conn, remoteAddr) 
        } 
} 

func forward(local net.Conn, remoteAddr string) { 
        remote, err := net.Dial("tcp", remoteAddr) 
        if remote == nil { 
                fmt.Fprintf(os.Stderr, "remote dial failed: %v\n", err) 
                return 
        } 
        go Copy(local, remote,"local") 
        go Copy(remote, local,"remote") 
  
} 

func Copy(dst net.Conn, src net.Conn, id string) {
        
        buf := &bytes.Buffer{}
        for {
                buf.Reset()
                data := make([]byte, 256)
                n, err := dst.Read(data)
                if err != nil {
                        panic(err)
                }
                buf.Write(data[:n])
                fmt.Printf("(%s)received:\n%v\n", id, hex.Dump(data[:n]))
                
                src.Write(buf.Bytes())
        }
<<<<<<< HEAD
=======

        //io.Copy(dst, src)

>>>>>>> 9931db2260cc6fa8cfef43b655ada063bddcc324
}

func fatal(s string, a ... interface{}) { 
        fmt.Fprintf(os.Stderr, "netfwd: %s\n", fmt.Sprintf(s, a)) 
        os.Exit(2) 
} 
