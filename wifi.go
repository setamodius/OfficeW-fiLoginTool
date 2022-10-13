package main
import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "bufio"
    "log"
    "os"
)

func httpPostForm(user string, pass string) {
    transCfg := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
    }
    client := &http.Client{Transport: transCfg}
    parm := url.Values{}
    parm.Add("username", user)
    parm.Add("password", pass)
    req, err := http.NewRequest("POST", "https://192.168.7.1:888/", strings.NewReader(parm.Encode()))
    req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("%s", err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Println(string(body))
}
func main() {
    var user = ""
    var pass = ""
    file, err := os.Open("password.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var text []string
    for scanner.Scan() {
        text = append(text, scanner.Text())
    }
    user = text[0]
    pass = text[1]
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    httpPostForm(user, pass);
}