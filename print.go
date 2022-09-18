package cmcm

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

func printPlain(w io.Writer, cmt ...*Comment) error {
	bw := bufio.NewWriter(w)
	for i, c := range cmt {
		var s string
		if i > 0 {
			s = "----------------------------\n"
		}
		s = s + fmt.Sprintln("ID:\t", c.ID)
		s = s + fmt.Sprintln("Author:\t", c.Author)
		s = s + fmt.Sprintln("URL:\t", c.HTMLURL)
		s = s + fmt.Sprintln("")
		s = s + fmt.Sprintln(c.Body)
		if _, err := bw.Write([]byte(s)); err != nil {
			return err
		}
	}

	return bw.Flush()
}

func printJSON(w io.Writer, cmt ...*Comment) error {
	var s []byte
	var err error
	if len(cmt) == 1 {
		s, err = json.Marshal(cmt[0])
	} else {
		s, err = json.Marshal(cmt)
	}
	if err != nil {
		return fmt.Errorf("failed to marshal: %w", err)
	}
	bw := bufio.NewWriter(w)
	if _, err := bw.Write(s); err != nil {
		return err
	}
	return bw.Flush()
}
