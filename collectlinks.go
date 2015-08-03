// This package does the extraordinarily simple operation of parsing a given piece of html
// and providing you with all the hyperlinks hrefs it finds.
package collectlinks

import (
  "golang.org/x/net/html"
  "io"
)

/* `All` takes a reader object (like the one returned from http.Get())
   It returns a slice of strings representing the "href" attributes from
   anchor links found in the provided html.

   It does not close the reader passed to it.
*/
func All(httpBody io.Reader) []string {
  links := make([]string, 0)
  page := html.NewTokenizer(httpBody)
  for {
    tokenType := page.Next()
    if tokenType == html.ErrorToken {
      return links
    }
    token := page.Token()
    if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
      for _, attr := range token.Attr {
        if attr.Key == "href" {
          links = append(links, attr.Val)
        }
      }
    }
  }
}
