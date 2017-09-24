package main
import (
      "fmt"
      "io"
      "strings"
      "os/exec"
      "net/http"
)

func main() {
      http.HandleFunc("/", repoUpdate)
      http.ListenAndServe(":8080", nil)
}

func repoUpdate(w http.ResponseWriter, r *http.Request) {
      var socket = strings.Split(r.RemoteAddr, ":")
      var repo string = r.URL.Query().Get("repo")
      var branch string = r.URL.Query().Get("branch")
      var dir string = "/var/www/html/" + repo

      cronLog("Connection from " + socket[0])

      cmd := exec.Command("git", "checkout", branch)
      cmd.Dir = dir
      out, err := cmd.Output()
      if err != nil {
            cronLog(err.Error())
            io.WriteString(w, "Branch checkout went wrong. Is " + branch + " a valid branch?\n")
            io.WriteString(w, string(out))

      } else {
            cronLog(string(out))
            io.WriteString(w, "Branch checkout success\n")
            w.Write(out)
            io.WriteString(w, "https://gv-capstone.duckdns.org")
            io.WriteString(w, "\n\n")
      }

      cmd = exec.Command("git", "pull")
      cmd.Dir = dir
      if err != nil {
            cronLog(err.Error())
            io.WriteString(w, "Branch pull went wrong\n")
            io.WriteString(w, string(out))

      } else {
            cronLog(string(out))
            io.WriteString(w, "Grand pull success\n")
            w.Write(out)
            io.WriteString(w, "https://gv-capstone.duckdns.org")
      }

}

func cronLog(message string) {
      err := exec.Command("logger", message).Run()
      if err != nil {
            fmt.Println("Failed to write to syslog")
            fmt.Println(message)
      } else {
            fmt.Println(message)
      }
}
