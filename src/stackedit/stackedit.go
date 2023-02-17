package stackedit

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"go.goblog.app/app/pkgs/bufferpool"
	"go.goblog.app/app/pkgs/plugintypes"
)

type plugin struct{}

func (p *plugin) Render(rc plugintypes.RenderContext, rendered io.Reader, modified io.Writer) {
	doc, err := goquery.NewDocumentFromReader(rendered)
	if err != nil {
		io.Copy(modified, rendered)
		return
	}

	head := doc.Find("head")
	if head.Length() == 0 {
		io.Copy(modified, rendered)
		return
	}

	var buffer bytes.Buffer
buffer.WriteString("<script src=\"https://unpkg.com/stackedit-js@1.0.7/docs/lib/stackedit.min.js\"></script>")

	script := bufferpool.Get()
	script.WriteString(buffer.String())

	head.AppendHtml(script.String())

	bufferpool.Put(script)
_ = goquery.Render(modified, doc.Selection)


	// inject script stuff
	body := doc.Find("body")
	if body.Length() == 0 {
		io.Copy(modified, rendered)
		return
	}
	var scriptbuffer bytes.Buffer
	scriptbuffer.WriteString(`<script>
function makeEditButton(el) {
  const div = document.createElement('div');
  div.className = 'stackedit-button-wrapper';
  div.innerHTML = '<a href="javascript:void(0)">Open Editor</a>';
  el.parentNode.insertBefore(div, el.nextSibling);
  return div.getElementsByTagName('a')[0];
}

const textareaUpdate = document.querySelector('#editor-update');
const textareaCreate = document.querySelector('#editor-create');

makeEditButton(textareaUpdate)
  .addEventListener('click', function onClick() {
    const stackedit = new Stackedit();
    stackedit.on('fileChange', function onFileChange(file) {
      textareaUpdate.value = file.content.text;
    });
    stackedit.openFile({
      name: 'Update Post',
      content: {
        text: textareaUpdate.value
      }
    });
  });
makeEditButton(textareaCreate)
  .addEventListener('click', function onClick() {
    const stackedit = new Stackedit();
    stackedit.on('fileChange', function onFileChange(file) {
      textareaCreate.value = file.content.text;
    });
    stackedit.openFile({
      name: 'Create Post',
      content: {
        text: textareaCreate.value
      }
    });
  });
</script>`)
	
	scriptx := bufferpool.Get()
	scriptx.WriteString(scriptbuffer.String())

	body.AppendHtml(scriptx.String())
	bufferpool.Put(scriptx)
	_ = goquery.Render(modified, doc.Selection)


}

func GetPlugin() plugintypes.UI2 {
	//return &plugin{}
}

