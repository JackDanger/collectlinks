package collectlinks

import (
  "fmt"
  "strings"
  "testing"
)

func TestAll(t *testing.T) {
  reader := strings.NewReader(` <p>
  <a href="http://motherfuckingwebsite.com">1</a>
  <a href='http://news.google.com'>2</a>
  <a style=\"\" href=http://imgur.com>3</a>
  http://dontcollectthis.com
</p>`)

  links := All(reader)

  fmt.Println(links)

  if len(links) != 3 {
    t.Error("Wrong number of links returned")
  }

  if links[0] != "http://motherfuckingwebsite.com" {
    t.Error("The first link is incorrect")
  }
  if links[1] != "http://news.google.com" {
    t.Error("The second link is incorrect")
  }
  if links[2] != "http://imgur.com" {
    t.Error("The third link is incorrect")
  }
}
