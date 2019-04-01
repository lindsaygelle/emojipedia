package content_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gellel/emojipedia/cache"
	"github.com/gellel/emojipedia/content"
)

func TestCreate(t *testing.T) {
	body := `
	<!doctype html>
	<html>
		<head>
			<title></title>
		</head>
		<body>
			<table>
				<tbody>
				<tr>
					<th colspan="5" class="bighead">
						<a href="#smileys_&amp;_emotion" name="smileys_&amp;_emotion">
							Smileys &amp; Emotion
						</a>
					</th>
				</tr>
				<tr>
					<th colspan="5" class="mediumhead">
						<a href="#face-smiling" name="face-smiling">
							face-smiling
						</a>
					</th>
				</tr>
				<tr>
					<td class="rchars">494</td>
					<td class="code">
						<a href="#1f417" name="1f417">U+1F417</a>
					</td>
					<td class="andr">
						<a href="full-emoji-list.html#1f417" target="full">
							<img alt="🐗" title="U+1F417 🐗 boar" class="imga" src="">
						</a>
					</td>
					<td class="name">
						boar
					</td>
					<td class="name">
						boar | pig
					</td>
				</tr>
				<tbody>
			</table>
		</body>
	</html>`
	response := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	err := cache.WriteHTML("test.html", response)
	if err != nil {
		message := fmt.Sprintf("content.go err != nil: reason %s", err)
		t.Errorf(message)
	}
	document, err := cache.GetHTML("test.html")
	if err != nil {
		message := fmt.Sprintf("content.go err != nil: reason %s", err)
		t.Errorf(message)
	}
	if len(content.Create(document)) != 1 {
		t.Errorf("content.go 0 != 1")
	}
}
