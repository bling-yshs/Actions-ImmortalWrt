package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

const basePath = "C:\\Users\\Administrator\\Desktop\\nodejs挂载测试\\"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        filePath := filepath.Join(basePath, r.URL.Path)

        data, err := ioutil.ReadFile(filePath)
        if err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/plain")
        w.WriteHeader(http.StatusOK)
        w.Write(data)
    })

    go func() {
        fmt.Println("Server running at http://localhost:8080/")
        http.ListenAndServe(":8080", nil)
    }()

    for {
        var input string
        fmt.Scanln(&input)
        if strings.TrimSpace(input) == "stop" {
            fmt.Println("Stopping server...")
            os.Exit(0)
        }
    }
}
