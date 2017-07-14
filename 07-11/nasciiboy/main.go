// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
  "fmt"
  "log"
  "net/http"
  "strconv"
  "sync"
)

var mu sync.Mutex

func main() {
  db := database{"shoes": 50, "socks": 5}
  http.HandleFunc("/list", db.list)
  http.HandleFunc("/price", db.price)
  http.HandleFunc("/update", db.update)
  http.HandleFunc("/remove", db.remove)
  http.HandleFunc("/add", db.add)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
  for item, price := range db {
    fmt.Fprintf(w, "%s: %s\n", item, price)
  }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
  item := req.URL.Query().Get("item")
  if price, ok := db[item]; ok {
    fmt.Fprintf(w, "%s\n", price)
  } else {
    w.WriteHeader(http.StatusNotFound) // 404
    fmt.Fprintf(w, "no such item: %q\n", item)
  }
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
  mu.Lock()
  defer mu.Unlock()

  item := req.URL.Query().Get("item")
  if _, ok := db[item]; ok {
    price := req.URL.Query().Get("price")
    v, err := strconv.ParseFloat( price, 32 )
    if err != nil {
      fmt.Fprintf(w, "invalid price: %v\n", err)
    } else {
      db[item] = dollars( v )
      fmt.Fprintf(w, "now %q = %s\n", item, db[item])
    }
  } else {
    w.WriteHeader(http.StatusNotFound) // 404
    fmt.Fprintf(w, "no such item: %q\n", item)
  }
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
  mu.Lock()
  defer mu.Unlock()

  item := req.URL.Query().Get("item")
  if _, ok := db[item]; ok {
    delete( db, item )
    fmt.Fprintf(w, "remove %q\n", item)
  } else {
    w.WriteHeader(http.StatusNotFound) // 404
    fmt.Fprintf(w, "no such item: %q\n", item)
  }
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
  mu.Lock()
  defer mu.Unlock()

  item := req.URL.Query().Get("item")
  if _, ok := db[item]; !ok {
    price := req.URL.Query().Get("price")
    if len( price ) != 0 {
      v, err := strconv.ParseFloat( price, 32 )
      if err != nil {
        fmt.Fprintf(w, "invalid price: %v\n", err)
      } else {
        db[item] = dollars( v )
        fmt.Fprintf(w, "add new item %q = %s\n", item, db[item])
      }
    } else {
      db[item] = 0
      fmt.Fprintf(w, "add new item %q = %s\n", item, db[item])
    }
  } else {
    w.WriteHeader(http.StatusNotFound) // 404
    fmt.Fprintf(w, "item %q already exist\n", item)
  }
}
