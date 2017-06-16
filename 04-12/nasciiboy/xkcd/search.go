package xkcd

import (
  "log"
  "fmt"
	"os"
  "path"
  "io"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "strconv"
  "strings"
)

const (
  pwd      = iota
  pJson
  pImg
  pDBInfo
  pDB
  n
)

var paths   [n]string
var files = [n]string {
   pJson  : "json",
   pImg   : "img",
   pDBInfo: "dbInfo.json",
   pDB    : "db.Json",
}

var hasPaths, hasInfo, hasDB, hasData bool

var lastComic *Comic
var info      *dbInfo
var db        *Comics

func LastComicData() (*Comic, error) {
  if lastComic != nil {
    return lastComic, nil
  }

	resp, err := http.Get(baseName + jsonRequest)
	if err != nil { return nil, err }
  defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("LastComic: search query failed: %s", resp.Status)
	}

  var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}

  lastComic = &comic
  return lastComic, nil
}

func GetComicData( n int ) (*Comic, error) {
  if lastComic == nil {
    if _, err := LastComicData(); err != nil {
      return nil, fmt.Errorf("Get: %v", err)
    }
  }

  if n > lastComic.Num {
    return nil, fmt.Errorf("GetComicData %d: lastComit is %d", n, lastComic.Num )
  }

	resp, err := http.Get(baseName + "/" + strconv.Itoa( n ) + jsonRequest)
	if err != nil { return nil, err }
  defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GetComicData %d: search query failed: %s", n, resp.Status)
	}

  var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}

  return &comic, nil
}

func MaxComic() int {
  if lastComic == nil {
    if _, err := LastComicData(); err != nil {
      return -1
    }
  }

  return lastComic.Num
}

func GetAllXKCD(){
  fmt.Println( "GetAllXKCD: init requests to https://xkcd.com" )

  if err := os.Mkdir( paths[ pJson ], 0755 ); os.IsNotExist( err ) {
    log.Fatalf( "GetAllXKCD: %v\n", err )
  }

  if err := os.Mkdir( paths[ pImg ], 0755 ); os.IsNotExist( err ) {
    log.Fatalf( "GetAllXKCD: %v\n", err )
  }

  info = new( dbInfo )
  *info = dbInfo{ Init: 1, Last: MaxComic() }
  for i := info.Init; i <=  info.Last; i++ {
    comicData, err := GetComicData( i )
    if err != nil {
      fmt.Fprintf( os.Stderr, "GetAllXKCD: %v\n", err )
      info.Fail = append( info.Fail, &FailRequest{ Num: i, Json: false, Img: false } )
      continue
    }

    dJson, err := json.MarshalIndent( comicData, "", "  " )
    if err != nil {
      fmt.Fprintf( os.Stderr, "GetAllXKCD: marshaling %d: %v\n", i, err )
      info.Fail = append( info.Fail, &FailRequest{ Num: i, Json: false, Img: false } )
      continue
    }

    jsonFileName := path.Join( paths[ pJson ], fmt.Sprintf( "%04d.json", i ) )
    err = ioutil.WriteFile( jsonFileName, dJson, 0644 )
    if err != nil {
      fmt.Fprintf( os.Stderr, "GetAllXKCD: write json %d: %v\n", i, err )
      info.Fail = append( info.Fail, &FailRequest{ Num: i, Json: false, Img: false } )
      continue
    }

    resp, err := http.Get( comicData.ImageURL )
    if err != nil {
      fmt.Fprintf( os.Stderr, "GetAllXKCD: get image %d: %v\n", i, err )
      info.Fail = append( info.Fail, &FailRequest{ Num: i, Json: true, Img: false } )
      continue
    }

    imgExt       := path.Ext( comicData.ImageURL )
    imgFileName  := path.Join( paths[ pImg ], fmt.Sprintf( "%04d%s", i, imgExt ) )
    imgFile, err := os.Create( imgFileName )
    if err != nil {
      fmt.Fprintf( os.Stderr, "GetAllXKCD: write img %d: %v\n", i, err )
      resp.Body.Close()
      info.Fail = append( info.Fail, &FailRequest{ Num: i, Json: true, Img: false } )
      continue
    }
    io.Copy( imgFile, resp.Body )
    imgFile.Close()
    resp.Body.Close()

    fmt.Printf( "GetAllXKCD: write %04d(.json|%s) ... DONE\n", i, imgExt )
    info.OK++
  }

  dJson, err := json.MarshalIndent( *info, "", "  " )
  if err != nil {
    log.Fatalf( "GetAllXKCD: write dbInfo: %v\n", err )
  }

  err = ioutil.WriteFile( paths[ pDBInfo ], dJson, 0644 )
  if err != nil {
    log.Fatalf( "GetAllXKCD: write dbInfo: %v\n", err )
  }

  hasInfo = true
}


func MakeDB(){
  if !hasPaths { getPaths() }

  if lastComic == nil {
    if _, err := LastComicData(); err != nil {
      log.Fatalf( "GetAllXKCD: write dbInfo: %v\n", err )
    }
  }

  cDB := make(Comics, lastComic.Num + 1)
  dJson, err := ioutil.ReadDir( paths[ pJson ] )
  if err != nil {
    log.Fatalf( "MakeDB: %v\n", err )
  }

  for _, f := range dJson {
    file, err := os.Open( path.Join( paths[ pJson ], f.Name() ) )
    if err != nil {
      fmt.Fprintf( os.Stderr, "MakeDB: %v\n", err )
      continue
    }

    var comic Comic
    if err := json.NewDecoder(file).Decode(&comic); err != nil {
      fmt.Fprintf( os.Stderr, "MakeDB: %v\n", err )
      file.Close()
      continue
    }
    file.Close()

    cDB[comic.Num] = comic
  }

  dbSon, err := json.MarshalIndent( cDB, "", "  " )
  if err != nil {
    log.Fatalf( "MakeDB: write dbInfo: %v\n", err )
  }

  err = ioutil.WriteFile( paths[ pDB ], dbSon, 0644 )
  if err != nil {
    log.Fatalf( "MakeDB: write dbInfo: %v\n", err )
  }

  db = &cDB
  hasDB = true
}

func LoadDB(){
  if !hasPaths { getPaths() }
  if hasDB { return }

  file, err := os.Open( paths[ pDB ] )
  if err != nil {
    log.Fatalf( "MakeDB: %v\n", err )
  }
  defer file.Close()

  var cDB Comics
  if err := json.NewDecoder(file).Decode(&cDB); err != nil {
    log.Fatalf( "MakeDB: %v\n", err )
  }

  db = &cDB
  hasDB = true
}

func Search( str string ) (c Comics) {
  str = strings.ToLower( str )
  c = make( Comics, 0, 32 )
  for _, comic := range *db {
    if strings.Contains( strings.ToLower(comic.Title), str ) {
      c = append( c, comic )
    } else if strings.Contains( strings.ToLower(comic.Transcript), str ) {
      c = append( c, comic )
    }
  }

  return c
}

func ExistData() bool {
  if !hasPaths { getPaths() }

  f, err := os.Stat( paths[ pJson   ] )
  if err != nil || !f.IsDir() { return false }
  f, err = os.Stat( paths[ pImg    ] )
  if err != nil || !f.IsDir() { return false }
  f, err = os.Stat( paths[ pDBInfo ] )
  if err != nil ||  f.IsDir() { return false }
  f, err = os.Stat( paths[ pDB     ] )
  if err != nil ||  f.IsDir() { return false }

  hasData = true
  return true
}

func getPaths(){
  var err error
  paths[pwd], err = os.Getwd()
  if err != nil {
    log.Fatalf( "getPaths: %v\n", err )
  }

  paths[ pJson   ] = path.Join( paths[pwd], files[ pJson   ] )
  paths[ pImg    ] = path.Join( paths[pwd], files[ pImg    ] )
  paths[ pDBInfo ] = path.Join( paths[pwd], files[ pDBInfo ] )
  paths[ pDB     ] = path.Join( paths[pwd], files[ pDB     ] )

  hasPaths = true
}
