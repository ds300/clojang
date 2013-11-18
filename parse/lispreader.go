package parse

import "os"
import "bufio"
import "unicode"
import "log"
import "bytes"
import "strings"
import "regexp"
import "strconv"

import "clojang/data/types"

func whitespace (c rune) bool {
  return unicode.IsSpace(c) || int(c) == 44 //comma
}

var intRegex = regexp.MustCompile(`^-?(0|[1-9]\d*)$`)
var floatRegex = regexp.MustCompile(`^-?(0|[1-9]\d*)(\.\d+)?((e|E)(\+|-)?\d+)?$`)
var fracRegex = regexp.MustCompile(`^-?(0|[1-9]\d*)/([1-9]\d*)$`)

func Read (rdr *bufio.Reader) types.IObj {
  c, _, _ := rdr.ReadRune()

  for whitespace(c) {
    c, _, _  = rdr.ReadRune()
  }

  switch {
    case '(' == c:
      return readList(rdr)

    case ')' == c:
      panic ("unexpected )")

    case '"' == c:
      return readString(rdr)

    default:
      rdr.UnreadRune()

      tkn := readToken(rdr)

      if strings.HasPrefix(tkn, ":") {
        log.Println("making a keyword:", tkn)
        ns, name := splitNs(tkn[1:])
        return &types.Keyword{ns, name}

      } else if intRegex.MatchString(tkn) {
        num, _ := strconv.ParseInt(tkn, 10, 64)
        i := types.Long(num)
        return &i

      } else if floatRegex.MatchString(tkn) {
        num, _ := strconv.ParseFloat(tkn, 64)
        i := types.Double(num)
        return &i

      } else if fracRegex.MatchString(tkn) {
        log.Println("me gots a frac:")
        parts := strings.Split(tkn, "/")
        nom, _ := strconv.ParseInt(parts[0], 10, 64)
        denom, _ := strconv.ParseInt(parts[1], 10, 64)
        i := types.Fraction{nom, denom}
        return &i

      } else {
        ns, name := splitNs(tkn)
        return &types.Symbol{ns, name}
      } 
  }
}

func splitNs (s string) (string, string) {
  parts := strings.Split(s, "/")
  if len(parts) == 1 {
    return "", s
  } else if len(parts) == 2 {
    return parts[0], parts[1]
  } else {
    panic("invalid symbol: " + s)
  }
}

func readString(rdr *bufio.Reader) *types.String {
  var buffer bytes.Buffer

  c, _, err := rdr.ReadRune()

  for c != '"' && err == nil {
    if c == '\\' {
      c, _, err = rdr.ReadRune()
      if c == 'n' {
        buffer.WriteRune('\n')

      } else if c == '"' {
        buffer.WriteRune('"')

      } else if c == '\\' {
        buffer.WriteRune('\\')

      } else if c == 'b' {
        buffer.WriteRune('\b')

      } else if c == 'f' {
        buffer.WriteRune('\f')

      } else if c == 'r' {
        buffer.WriteRune('\r')

      } else if c == 't' {
        buffer.WriteRune('\t')

      } else if c == 'u' {
        uffff := 0

        for i := 0; i<4; i++ {
          c, _, err = rdr.ReadRune()
          if '0' <= c && c <= '9' {
            uffff = uffff * 16 + (int(c) - 48)

          } else if 'a' <= c && c <= 'f' {
            uffff = uffff * 16 + (int(c) - 87)

          } else if 'A' <= c && c <= 'F' {
            uffff = uffff * 16 + (int(c) - 55)

          } else {
            panic(string(c) + " is not a hexadecimal digit")
          }
        }

        buffer.WriteRune(rune(uffff))
      }

    } else {
      buffer.WriteRune(c)
    }
    c, _, err = rdr.ReadRune()
  }

  result := types.String(buffer.String())

  return &result
}


func readList (rdr *bufio.Reader) *types.List {
  var head types.List
  tail := &head

  c, _, _ := rdr.ReadRune()

  for c != ')' {
    rdr.UnreadRune()
    tail.Val = Read(rdr)
    c, _, _ = rdr.ReadRune()
    tail.Next = &types.List{}
    tail = tail.Next
  }

  return &head
}

func readToken (rdr *bufio.Reader) string {
  c, _, err := rdr.ReadRune()

  for whitespace(c) {
    c, _, err  = rdr.ReadRune()
  }

  var buffer bytes.Buffer

  for !whitespace(c) && c != ')' && err == nil {
    buffer.WriteRune(c)
    c, _, err = rdr.ReadRune()
  }

  rdr.UnreadRune()

  log.Println(buffer.String() + " and '"+string(c)+"'")

  return buffer.String()
}

func ReadFile (path string) types.IObj {
  log.Println("ReadFile yo")
  file, err := os.Open(path)
  if err != nil { panic(err) }
  rdr := bufio.NewReader(file)
  result := Read (rdr)
  file.Close()
  log.Println("done with that jazz")
  return result
}