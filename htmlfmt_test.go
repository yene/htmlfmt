package main

import (
	"testing"

	"github.com/yosssi/gohtml"
)

func TestHTMLDocumentHTML(t *testing.T) {
	s := `<!DOCTYPE html><html><head><title>This is a title.</title></head><body><p>Line1<br>Line2</p><br/></body></html><!-- aaa -->`
	actual := gohtml.Format(s)
	expected := `<!DOCTYPE html>
<html>
  <head>
    <title>
      This is a title.
    </title>
  </head>
  <body>
    <p>
      Line1
      <br>
      Line2
    </p>
    <br/>
  </body>
</html>
<!-- aaa -->`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestMultilineComments(t *testing.T) {
	s := `
	<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<!--[if lt IE 9]>
  <script src="https://cdn.jsdelivr.net/npm/html5shiv@3.7.3/dist/html5shiv.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/respond.js@1.4.2/dest/respond.min.js"></script>
<![endif]-->
	</head>
	`
	actual := gohtml.Format(s)
	expected := `<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <!--[if lt IE 9]>
    <script src="https://cdn.jsdelivr.net/npm/html5shiv@3.7.3/dist/html5shiv.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/respond.js@1.4.2/dest/respond.min.js"></script>
  <![endif]-->
</head>`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}

func TestHTMLTextWithNewline(t *testing.T) {
	s := `
<!DOCTYPE html><html><head></head><body>
<div>
  <span>
    I am content.
  </span>
</div>
</body></html>
	`
	actual := gohtml.Format(s)
	expected := `<!DOCTYPE html>
<html>
  <head>
  </head>
  <body>
    <div>
      <span>
        I am content.
      </span>
    </div>
  </body>
</html>`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}
