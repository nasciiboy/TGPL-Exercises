The Go Programming Language Exercises

feel free to collaborate, only follow the structure

  chapter-exercise/username/solution

for example

  01-02/nasciiboy/main.go


======== Index

01.01 :: Modify the echo program to also print @c(os.Args[0])
01.02 :: echo program to print the index and value of each of its arguments, one per line.
01.03 :: str += VS strings.Join
01.04 :: dup2 to print the names of all files in which each duplicated line occurs
01.05 :: Change the Lissajous program’s color palette to green on black
01.06 :: Modify the Lissajous program to produce images in multiple colors
01.07 :: Use io.Copy(dst, src) it instead of ioutil.ReadAll to copy the response http.Body to os.Stdout
01.08 :: Modify fetch to add the prefix "http://" to each argument URL if it is missing.
01.09 :: Modify fetch to also print the HTTP status code, found in resp.Status.
01.10 :: Find a web site that produces a large amount of data. Investigate
         caching by running @c(fetchall) twice in succession to see whether the
         reported time changes much. Do you get the same content each time?
         Modify @c(fetchall) to print its output to a file so it can be examined.
01.11 :: Try fetchall) with longer argument lists, such as samples from the top
         million web sites available at http://www.alexa.com/.  How does the
         program behave if a web site just doesn’t respond? (Section 8.9
         describes mechanisms for coping in such cases.)
01.12 :: Modify the Lissajous server to read parameter values from the URL. For
         example, you might arrange it so that a URL like
         http://localhost:8000/?cycles=20 sets the number of cycles to 20
         instead of the default 5. Use the strconv.Atoi function to convert the
         string parameter into an integer. You can see its documentation with go
         doc strconv.Atoi.
02.01 :: Add types, constants, and functions to tempconv for processing
         temperatures in the Kelvin scale, where zero Kelvin is -273.15°C and a
         difference of 1K has the same magnitude as 1°C.
02.02 :: Write a general-purpose unit-conversion program analogous to cf that
         reads numbers from its command-line arguments or from the standard
         input if there are no arguments, and converts each number into units
         like temperature in Celsius and Fahrenheit, length in feet and meters,
         weight in pounds and kilograms, and the like.
02.03 :: Rewrite PopCount to use a loop instead of a single expression. Compare
         the performance of the two versions. (#Section 11.4 shows how to
         compare the performance of different implementations systematically.)
02.04 :: Write a version of PopCount that counts bits by shifting its argument
         through 64 bit positions, testing the rightmost bit each time. Compare
         its performance to the tablelookup version.
02.05 :: The expression x&(x-1) clears the rightmost non-zero bit of x. Write a
         version of PopCount that counts bits by using this fact, and assess its
         performance.
03.01 :: If the function f returns a non-finite float64 value, the SVG file
         will contain invalid <polygon> elements (although many SVG renderers
         handle this gracefully). Modify the program to skip invalid polygons.
03.02 :: Experiment with visualizations of other functions from the math
         package. Can you produce an egg box, moguls, or a saddle?
03.03 :: Color each polygon based on its height, so that the peaks are colored
         red (#ff0000) and the valleys blue (#0000ff).
03.04 :: Following the approach of the Lissajous example in Section 1.7,
         construct a web server that computes surfaces and writes SVG data to
         the client. The server must set the Content-Type header like this:
03.05 :: Implement a full-color Mandelbrot set using the function image.NewRGBA
         and the type color.RGBA or color.YCbCr.
03.06 :: Supersampling is a technique to reduce the effect of pixelation by
         computing the color value at several points within each pixel and
         taking the average. The simplest method is to divide each pixel into
         four subpixels. Implement it.
03.07 :: Another simple fractal uses Newton’s method to find complex solutions
         to a function such as z⁴-1 = 0. Shade each starting point by the number
         of iterations required to get close to one of the four roots. Color
         each point by the root it approaches.
03.08 :: Rendering fractals at high zoom levels demands great arithmetic
         precision. Implement the same fractal using four different
         representations of numbers: complex64, complex128, big.Float, and
         big.Rat. (The latter two types are found in the math/big package.
         Float uses arbitrary but bounded-precision floating-point; Rat uses
         unbounded-precision rational numbers.) How do they compare in
         performance and memory usage? At what zoom levels do rendering
         artifacts become visible?
03.09 :: Write a web server that renders fractals and writes the image data to
         the client.  Allow the client to specify the x, y, and zoom values as
         parameters to the HTTP request.
03.10 :: Write a non-recursive version of comma, using bytes.Buffer instead of
          string concatenation.
03.11 :: Enhance comma so that it deals correctly with floating-point numbers
         and an optional sign.
03.12 :: Write a function that reports whether two strings are anagrams of each
         other, that is, they contain the same letters in a different order.
03.13 :: Write const declarations for KB, MB, up through YB as compactly as you
         can.
04.01 :: Write a function that counts the number of bits that are different in
         two SHA256 hashes. (See PopCount from Section 2.6.2.)
04.02 :: Write a program that prints the SHA256 hash of its standard input by
         default but supports a command-line flag to print the SHA384 or SHA512
         hash instead.
04.03 :: Rewrite reverse to use an array pointer instead of a slice.
04.04 :: Write a version of rotate that operates in a single pass.
04.05 :: Write an in-place function to eliminate adjacent duplicates in a
          []string slice.
04.06 :: Write an in-place function that squashes each run of adjacent Unicode
         spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a
         single ASCII space.
04.07 :: Modify reverse to reverse the characters of a []byte slice that
         represents a UTF-8-encoded string, in place. Can you do it without
         allocating new memory?
04.08 :: Modify charcount to count letters, digits, and so on in their Unicode
         categories, using functions like unicode.IsLetter.
04.09 :: Write a program wordfreq to report the frequency of each word in an
         input text file. Call input.Split(bufio.ScanWords) before the first
         call to Scan to break the input into words instead of lines.
04.10 :: Modify issues to report the results in age categories, say less than a
         month old, less than a year old, and more than a year old.
04.11 :: Build a tool that lets users create, read, update, and delete GitHub
         issues from the command line, invoking their preferred text editor when
         substantial text input is required.
04.12 :: The popular web comic xkcd has a JSON interface. For example, a request
         to https://xkcd.com/571/info.0.json produces a detailed description of
         comic 571, one of many favorites. Download each URL (once!)  and build
         an offline index. Write a tool xkcd that, using this index, prints the
         URL and transcript of each comic that matches a search term provided on
         the command line.
04.13 :: The JSON-based web service of the Open Movie Database lets you search
         https://omdbapi.com/ for a movie by name and download its poster
         image. Write a tool poster that downloads the poster image for the
         movie named on the command line.
04.14 :: Create a web server that queries GitHub once and then allows navigation
         of the list of bug reports, milestones, and users.
05.01 :: Change the findlinks program to traverse the n.FirstChild linked list
         using recursive calls to visit instead of a loop.
05.02 :: Write a function to populate a mapping from element names—p, div, span,
         and so on—to the number of elements with that name in an HTML document
         tree.
05.03 :: Write a function to print the contents of all text nodes in an HTML
         document tree. Do not descend into <script> or <style> elements, since
         their contents are not visible in a web browser.
05.04 :: Extend the visit function so that it extracts other kinds of links from
         the document, such as images, scripts, and style sheets.
05.05 :: Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)
05.06 :: Modify the corner function in gopl.io/ch3/surface (Section 3.2) to use
         named results and a bare return statement.
05.07 :: Develop startElement and endElement into a general HTML pretty-printer.
         Print comment nodes, text nodes, and the attributes of each element (<a
         href='...'>). Use short forms like <img/> instead of <img></img> when
         an element has no children. Write a test to ensure that the output can
         be parsed successfully. (See #Chapter 11.)
05.08 :: Modify forEachNode so that the pre and post functions return a boolean
         result indicating whether to continue the traversal. Use it to write a
         function ElementByID with the following signature that finds the first
         HTML element with the specified id attribute. The function should stop
         the traversal as soon as a match is found.
05.09 :: Write a function expand(s string, f func(string) string) string that
         replaces each substring $foo within s by the text returned by f("foo").
05.10 :: Rewrite topoSort to use maps instead of slices and eliminate the
         initial sort.  Verify that the results, though nondeterministic, are
         valid topological orderings.
05.11 :: The instructor of the linear algebra course decides that calculus is
         now a prerequisite. Extend the topoSort function to report cycles.
05.12 :: The startElement and endElement functions in gopl.io/ch5/outline2
         (Section 5.5) share a global variable, depth. Turn them into anonymous
         functions that share a variable local to the outline function.
05.13 :: Modify crawl to make local copies of the pages it finds, creating
         directories as necessary. Don’t make copies of pages that come from a
         different domain. For example, if the original page comes from
         https://golang.org, save all files from there, but exclude ones from
         https://vimeo.com.
05.14 :: Use the breadthFirst function to explore a different structure. For
         example, you could use the course dependencies from the topoSort
         example (a directed graph), the file system hierarchy on your computer
         (a tree), or a list of bus or subway routes downloaded from your city
         government’s web site (an undirected graph).
05.15 :: Write variadic functions max and min, analogous to sum. What should
         these functions do when called with no arguments? Write variants that
         require at least one argument.
05.16 :: Write a variadic version of strings.Join.
05.17 :: Write a variadic function ElementsByTagName that, given an HTML node
         tree and zero or more names, returns all the elements that match one of
         those names. Here are two example calls:
05.18 :: Without changing its behavior, rewrite the fetch function to use defer
         to close the writable file.
05.19 :: Use panic and recover to write a function that contains no return
         statement yet returns a non-zero value.
06.01 :: Implement these additional methods:
06.02 :: Define a variadic (*IntSet).AddAll(...int) method that allows a list of
         values to be added, such as s.AddAll(1, 2, 3).
06.03 :: (*IntSet).UnionWith computes the union of two sets using |, the
         word-parallel bitwise OR operator. Implement methods for IntersectWith,
         DifferenceWith, and SymmetricDifference for the corresponding set
         operations. (The symmetric difference of two sets contains the elements
         present in one set or the other but not both.)
06.04 :: Add a method Elems that returns a slice containing the elements of the
         set, suitable for iterating over with a range loop.
06.05 :: The type of each word used by IntSet is uint64, but 64-bit arithmetic
         may be inefficient on a 32-bit platform. Modify the program to use the
         uint type, which is the most efficient unsigned integer type for the
         platform. Instead of dividing by 64, define a constant holding the
         effective size of uint in bits, 32 or 64. You can use the perhaps
         too-clever expression 32 << (^uint(0) >> 63) for this purpose.
07.01 :: Using the ideas from ByteCounter, implement counters for words and for
         lines.  You will find bufio.ScanWords useful.
07.02 :: Write a function CountingWriter with the signature below that, given an
         io.Writer, returns a new Writer that wraps the original, and a pointer
         to an int64 variable that at any moment contains the number of bytes
         written to the new Writer.
07.03 :: Write a String method for the *tree type in gopl.io/ch4/treesort
         (Section 4.4) that reveals the sequence of values in the tree.
07.04 :: The strings.NewReader function returns a value that satisfies the
         io.Reader interface (and others) by reading from its argument, a
         string. Implement a simple version of NewReader yourself, and use it to
         make the HTML parser (Section 5.2) take input from a string.
07.05 :: The LimitReader function in the io package accepts an io.Reader r and a
         number of bytes n, and returns another Reader that reads from r but
         reports an end-of-file condition after n bytes. Implement it.
07.06 :: Add support for Kelvin temperatures to tempflag.
07.07 :: Explain why the help message contains °C when the default value of 20.0
         does not.
07.08 :: Many GUIs provide a table widget with a stateful multi-tier sort: the
         primary sort key is the most recently clicked column head, the
         secondary sort key is the second-most recently clicked column head, and
         so on. Define an implementation of sort.Interface for use by such a
         table. Compare that approach with repeated sorting using sort.Stable.
07.09 :: Use the html/template package (Section 4.6) to replace printTracks with
         a function that displays the tracks as an HTML table. Use the solution
         to the previous exercise to arrange that each click on a column head
         makes an HTTP request to sort the table.
07.10 :: The sort.Interface type can be adapted to other uses. Write a function
         IsPalindrome(s sort.Interface) bool that reports whether the sequence s
         is a palindrome, in other words, reversing the sequence would not
         change it. Assume that the elements at indices i and j are equal if
         !s.Less(i, j) && !s.Less(j, i).
07.11 :: Add additional handlers so that clients can create, read, update, and
         delete database entries. For example, a request of the form
         /update?item=socks&price=6 will update the price of an item in the
         inventory and report an error if the item does not exist or if the
         price is invalid. (Warning: this change introduces concurrent variable
         updates.)
07.12 :: Change the handler for /list to print its output as an HTML table, not
         text.  You may find the html/template package (Section 4.6) useful.
07.13 :: Add a String method to Expr to pretty-print the syntax tree. Check that
         the results, when parsed again, yield an equivalent tree.
07.14 :: Define a new concrete type that satisfies the Expr interface and
         provides a new operation such as computing the minimum value of its
         operands. Since the Parse function does not create instances of this
         new type, to use it you will need to construct a syntax tree directly
         (or extend the parser).
07.15 :: Write a program that reads a single expression from the standard input,
         prompts the user to provide values for any variables, then evaluates
         the expression in the resulting environment. Handle all errors
         gracefully.
07.16 :: Write a web-based calculator program.
07.17 :: Extend xmlselect so that elements may be selected not just by name, but
         by their attributes too, in the manner of CSS, so that, for instance,
         an element like <div id="page" class="wide"> could be selected by a
         matching id or class as well as its name.
07.18 :: Using the token-based decoder API, write a program that will read an
         arbitrary XML document and construct a tree of generic nodes that
         represents it. Nodes are of two kinds: CharData nodes represent text
         strings, and Element nodes represent named elements and their
         attributes. Each element node has a slice of child nodes.
08.01 :: Modify clock2 to accept a port number, and write a program, clockwall,
         that acts as a client of several clock servers at once, reading the
         times from each one and displaying the results in a table, akin to the
         wall of clocks seen in some business offices. If you have access to
         geographically distributed computers, run instances remotely ;
         otherwise run local instances on different ports with fake time zones.
08.02 :: Implement a concurrent File Transfer Protocol (FTP) server. The server
         should interpret commands from each client such as cd to change
         directory, ls to list a directory, get to send the contents of a file,
         and close to close the connection. You can use the standard ftp command
         as the client, or write your own.
08.03 :: In netcat3, the interface value conn has the concrete type *net.TCPConn,
         which represents a TCP connection. A TCP connection consists of two
         halves that may be closed independently using its CloseRead and
         CloseWrite methods. Modify the main goroutine of netcat3 to close only
         the write half of the connection so that the program will continue to
         print the final echoes from the reverb1 server even after the standard
         input has been closed.  (Doing this for the reverb2 server is harder;
         see Exercise 8.4.)
08.04 :: Modify the reverb2 server to use a sync.WaitGroup per connection to
         count the number of active echo goroutines. When it falls to zero,
         close the write half of the TCP connection as described in Exercise
         8.3. Verify that your modified netcat3 client from that exercise waits
         for the final echoes of multiple concurrent shouts, even after the
         standard input has been closed.
08.05 :: Take an existing CPU-bound sequential program, such as the Mandelbrot
         program of Section 3.3 or the 3-D surface computation of Section 3.2,
         and execute its main loop in parallel using channels for
         communication. How much faster does it run on a multiprocessor machine?
         What is the optimal number of goroutines to use?
08.06 :: Add depth-limiting to the concurrent crawler. That is, if the user sets
         -depth=3, then only URLs reachable by at most three links will be
         fetched.
08.07 :: Write a concurrent program that creates a local mirror of a web site,
         fetching each reachable page and writing it to a directory on the local
         disk. Only pages within the original domain (for instance,
         https://golang.org) should be fetched. URLs within mirrored pages
         should be altered as needed so that they refer to the mirrored page,
         not the original.
08.08 :: Using a select statement, add a timeout to the echo server from Section
         8.3 so that it disconnects any client that shouts nothing within 10
         seconds.
08.09 :: Write a version of du that computes and periodically displays separate
         totals for each of the root directories.
08.10 :: HTTP requests may be cancelled by closing the optional Cancel channel
         in the http.Request struct. Modify the web crawler of Section 8.6 to
         support cancellation.
08.11 :: Following the approach of mirroredQuery in Section 8.4.4, implement a
         variant of fetch that requests several URLs concurrently. As soon as
         the first response arrives, cancel the other requests.
08.12 :: Make the broadcaster announce the current set of clients to each new
         arrival.  This requires that the clients set and the entering and
         leaving channels record the client name too.
08.13 :: Make the chat server disconnect idle clients, such as those that have
         sent no messages in the last five minutes. Hint: calling conn.Close()
         in another goroutine unblocks active Read calls such as the one done by
         input.Scan().
08.14 :: Change the chat server’s network protocol so that each client provides
         its name on entering. Use that name instead of the network address when
         prefixing each message with its sender’s identity.
08.15 :: Failure of any client program to read data in a timely manner
         ultimately causes all clients to get stuck. Modify the broadcaster to
         skip a message rather than wait if a client writer is not ready to
         accept it. Alternatively, add buffering to each client’s outgoing
         message channel so that most messages are not dropped; the broadcaster
         should use a non-blocking send to this channel.
09.01 :: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1
         program. The result should indicate whether the transaction succeeded
         or failed due to insufficient funds. The message sent to the monitor
         goroutine must contain both the amount to withdraw and a new channel
         over which the monitor goroutine can send the boolean result back to
         Withdraw.
09.02 :: Rewrite the PopCount example from Section 2.6.2 so that it initializes
         the lookup table using sync.Once the first time it is
         needed. (Realistically, the cost of synchronization would be
         prohibitive for a small and highly optimized function like PopCount.)
09.03 :: Extend the Func type and the (*Memo).Get method so that callers may
         provide an optional done channel through which they can cancel the
         operation (Section 8.9). The results of a cancelled Func call should
         not be cached.
09.04 :: Construct a pipeline that connects an arbitrary number of goroutines
         with channels. What is the maximum number of pipeline stages you can
         create without running out of memory? How long does a value take to
         transit the entire pipeline?
09.05 :: Write a program with two goroutines that send messages back and forth
         over two unbuffered channels in ping-pong fashion. How many
         communications per second can the program sustain?
09.06 :: Measure how the performance of a compute-bound parallel program (see
         Exercise 8.5) varies with GOMAXPROCS. What is the optimal value on your
         computer? How many CPUs does your computer have?
10.01 :: Extend the jpeg program so that it converts any supported input format
         to any output format, using image.Decode to detect the input format and
         a flag to select the output format.
10.02 :: Define a generic archive file-reading function capable of reading ZIP
         files (archive/zip) and POSIX tar files (archive/tar).  Use a
         registration mechanism similar to the one described above so that
         support for each file format can be plugged in using blank imports.
10.03 :: Using fetch http://gopl.io/ch1/helloworld?go-get=1, find out which
          service hosts the code samples for this book. (HTTP requests from go
          get include the go-get parameter so that servers can distinguish them
          from ordinary browser requests.)
10.04 :: Construct a tool that reports the set of all packages in the workspace
         that transitively depend on the packages specified by the
         arguments. Hint: you will need to run go list twice, once for the
         initial packages and once for all packages. You may want to parse its
         JSON output using the encoding/json package (Section 4.5).
11.01 :: Write tests for the charcount program in Section 4.3.
11.02 :: Write a set of tests for IntSet (Section 6.5) that checks that its
         behavior after each operation is equivalent to a set based on built-in
         maps. Save your implementation for benchmarking in Exercise 11.7.
11.03 :: TestRandomPalindromes only tests palindromes. Write a randomized test
         that generates and verifies non-palindromes.
11.04 :: Modify randomPalindrome to exercise IsPalindrome’s handling of
         punctuation and spaces.
11.05 :: Extend TestSplit to use a table of inputs and expected outputs.
11.06 :: Write benchmarks to compare the PopCount implementation in Section
         2.6.2 with your solutions to Exercise 2.4 and Exercise 2.5. At what
         point does the table-based approach break even?
11.07 :: Write benchmarks for Add, UnionWith, and other methods of *IntSet
         (Section 6.5) using large pseudo-random inputs. How fast can you make
         these methods run? How does the choice of word size affect performance?
         How fast is IntSet compared to a set implementation based on the
         built-in map type?
12.01 :: Extend Display so that it can display maps whose keys are structs or
         arrays.
12.02 :: Make display safe to use on cyclic data structures by bounding the
         number of steps it takes before abandoning the recursion. (In Section
         13.3, we’ll see another way to detect cycles.)
12.03 :: Implement the missing cases of the encode function. Encode booleans as
         t and nil, floating-point numbers using Go’s notation, and complex
         numbers like 1+2i as #C(1.0 2.0). Interfaces can be encoded as a pair
         of a type name and a value, for instance ("[]int" (1 2 3)), but beware
         that this notation is ambiguous: the reflect.Type.String method may
         return the same string for different types.
12.04 :: Modify encode to pretty-print the S-expression in the style shown
         above.
12.05 :: Adapt encode to emit JSON instead of S-expressions. Test your encoder
         using the standard decoder, json.Unmarshal.
12.06 :: Adapt encode so that, as an optimization, it does not encode a field
         whose value is the zero value of its type.
12.07 :: Create a streaming API for the S-expression decoder, following the
         style of json.Decoder (Section 4.5).
12.08 :: The sexpr.Unmarshal function, like json.Marshal, requires the complete
         input in a byte slice before it can begin decoding. Define a
         sexpr.Decoder type that, like json.Decoder, allows a sequence of values
         to be decoded from an io.Reader. Change sexpr.Unmarshal to use this new
         type.
12.09 :: Write a token-based API for decoding S-expressions, following the style
         of xml.Decoder (Section 7.14). You will need five types of tokens:
         Symbol, String, Int, StartList, and EndList.
12.10 :: Extend sexpr.Unmarshal to handle the booleans, floating-point numbers,
         and interfaces encoded by your solution to Exercise 12.3. (Hint: to
         decode interfaces, you will need a mapping from the name of each
         supported type to its reflect.Type.)
12.11 :: Write the corresponding Pack function. Given a struct value, Pack
         should return a URL incorporating the parameter values from the struct.
12.12 :: Extend the field tag notation to express parameter validity
         requirements. For example, a string might need to be a valid email
         address or credit-card number, and an integer might need to be a valid
         US ZIP code. Modify Unpack to check these requirements.
12.13 :: Modify the S-expression encoder (Section 12.4) and decoder (Section
         12.6) so that they honor the sexpr:"..." field tag in a similar manner
         to encoding/json (Section 4.5).
13.01 :: Define a deep comparison function that considers numbers (of any type)
         equal if they differ by less than one part in a billion.
13.02 :: Write a function that reports whether its argument is a cyclic data
         structure.
13.03 :: Use sync.Mutex to make bzip2.writer safe for concurrent use by multiple
         goroutines.
13.04 :: Depending on C libraries has its drawbacks. Provide an alternative
         pure-Go implementation of bzip.NewWriter that uses the os/exec package
         to run /bin/bzip2 as a subprocess.
