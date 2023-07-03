package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://appx.laflorida.cl/licencia_conducir/controllers/get_dias_cupo.php?mes=07&anno=2023&id_tramite=5"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  //Headers
  req.Header.Add("Accept", "*/*")
  req.Header.Add("Accept-Language", "es-419,es;q=0.9,en-CA;q=0.8,en;q=0.7,es-CL;q=0.6,es-US;q=0.5")
  req.Header.Add("Connection", "keep-alive")
  req.Header.Add("Cookie", "_ga=GA1.2.1362950678.1686834006; _gid=GA1.2.713970021.1686834006; PHPSESSID=fq1m7sum116kgm96sc1t1aanl5; _ga_WT2Z1T3Y8H=GS1.1.1686834004.1.1.1686834023.41.0.0")
  req.Header.Add("Referer", "https://appx.laflorida.cl/licencia_conducir/calendario.php?id_tramite=5")
  req.Header.Add("Sec-Fetch-Dest", "empty")
  req.Header.Add("Sec-Fetch-Mode", "cors")
  req.Header.Add("Sec-Fetch-Site", "same-origin")
  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
  req.Header.Add("X-Requested-With", "XMLHttpRequest")
  req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Google Chrome\";v=\"114\"")
  req.Header.Add("sec-ch-ua-mobile", "?0")
  req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}