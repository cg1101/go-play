# Featured Video

## Topic 1 Interface

In Go, there is no need to declare that a type *implements* an interface, just define all the methods it requires.

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, %s\n", new(World))
}

type World struct{}

func (w *World) String() string {
	return "世界"
}
```

```go
package main

import (
	"fmt"
)

type Office int

const (
	Boston Office = iota
	NewYork
)

var officePlace = []string{"Cambridge, MA", "New York, NY"}

// this is a method on integer not a pointer of struct
func (o Office) String() string {

	return "Google, " + officePlace[o]
}

func main() {
	fmt.Printf("Hello, %s\n", Boston)
}
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()
	fmt.Printf("Hello, %s (%d)\n", day, day)
}
```

```go
package main

import (
	"fmt"
	"time"
)

func fetch(url string) {
	y, x := 1, 2
	for (i )
}

func main() {
	start := time.Now()
	fetch("http://www.google.com/")
	fmt.Println(time.Since(start))
}
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "Hello, 世界\n")
}
```

```go
package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	fmt.Fprintf(h, "Hello, 世界\n")
	fmt.Printf("hash=%#x\n", h.Sum32())
}
```

```go
package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	h := crc32.NewIEEE()
	w := io.MultiWriter(h, os.Stdout)
	fmt.Fprintf(w, "Hello, 世界\n")
	fmt.Printf("hash=%#x\n", h.Sum32())
}
```

```go
package main

import (
	"fmt"
)

type Lang struct {
	Name string
	Year int
	URL string
}

func main() {
	lang := Lang{"Go", 2009, "http://golang.org/"}
	fmt.Printf("%v\n", lang)	// or %+v
	// fmt.Printf("%+v\n", lang)
	// fmt.Printf("%#v\n", lang)
}
```

## Topic 2 Reflection

```go
package main

import (
	"os"
	"reflect"
	"strconv"
)

func main() {
	myPrint("Hello ", 42, "\n")
}

func myPrint(args ...interface{}) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		}
	}
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Lang struct {
	Name string
	Year int
	URL string
}

func main() {
	lang := Lang{"Go", 2009, "http://golang.org/"}
	data, err := json.Marshal(lang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
```

```go
package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Lang struct {
	Name string
	Year int
	URL string
}

func main() {
	lang := Lang{"Go", 2009, "http://golang.org/"}
	data, err := xml.MarshalIndent(lang, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
```

```go
package main

import (
	"io"
	"log"
	"os"
	
)

func main() {
	input, err := os.Open("/Users/rsc/lang.json")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, input)
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	input, err := os.Open("./lang.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("%v\n", lang)
	}
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func do(f func(Lang)) {
	input, err := os.Open("./lang.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}

func main() {
	do(func(lang Lang) {
		fmt.Printf("%v\n", lang)
	})
}
```

```go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func do(f func(Lang)) {
	input, err := os.Open("./lang.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}

func main() {
	do(func(lang Lang) {
		data, err := xml.MarshalIndent(lang, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", data)
	})
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func do(f func(Lang)) {
	input, err := os.Open("./lang.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}

func count(name, url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	fmt.Printf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}

func main() {
	start := time.Now()
	do(func(lang Lang) {
		count(lang.Name, lang.URL)
	})
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
```

## Topic 3 Concurrency

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func do(f func(Lang)) {
	input, err := os.Open("./lang.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}

func count(name, url string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	fmt.Printf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}

func main() {
	do(func(lang Lang) {
		go count(lang.Name, lang.URL)
	})
	time.Sleep(10 * time.Second)
}
```

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func do(f func(Lang)) {
	input, err := os.Open("./lang.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}

func count(name, url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	dt := time.Since(start).Seconds()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, dt)
}

func main() {
	start := time.Now()
	c := make(chan string)
	n := 0
	do(func(lang Lang) {
		n++
		go count(lang.Name, lang.URL, c)
	})
	for i := 0; i < n; i++ {
		fmt.Print(<-c)
	}
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
```

```go
func main() {
	start := time.Now()
	c := make(chan string)
	n := 0
	do(func(lang Lang) {
		n++
		go count(lang.Name, lang.URL, c)
	})

	timeout := time.After(10 * time.Second)
	for i := 0; i < n; i++ {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-timeout:
			fmt.Print("Timed out\n")
			return
		}
	}
	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}
```
