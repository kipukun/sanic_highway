// Code generated by qtc from "admin_layout.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/admin_layout.qtpl:1
package templates

//line templates/admin_layout.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/admin_layout.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/admin_layout.qtpl:2
type AdminPage interface {
//line templates/admin_layout.qtpl:2
	Title() string
//line templates/admin_layout.qtpl:2
	StreamTitle(qw422016 *qt422016.Writer)
//line templates/admin_layout.qtpl:2
	WriteTitle(qq422016 qtio422016.Writer)
//line templates/admin_layout.qtpl:2
	Content() string
//line templates/admin_layout.qtpl:2
	StreamContent(qw422016 *qt422016.Writer)
//line templates/admin_layout.qtpl:2
	WriteContent(qq422016 qtio422016.Writer)
//line templates/admin_layout.qtpl:2
}

//line templates/admin_layout.qtpl:8
func StreamAdminPageTemplate(qw422016 *qt422016.Writer, p AdminPage) {
//line templates/admin_layout.qtpl:8
	qw422016.N().S(`
<!DOCTYPE html>
<html>
	<head>
		`)
//line templates/admin_layout.qtpl:12
	p.StreamTitle(qw422016)
//line templates/admin_layout.qtpl:12
	qw422016.N().S(`
		<link rel="stylesheet" href="/assets/style.css">
	</head>
	<body>
		<header>
			<a src="/" id="logo">
				Sanic Information Highway 
				<b>Maintenance & Administration</b>
			</a>
			<nav>
				<ul>
					<li><a href="/admin">Home</a></li>
					<li><a href="/admin/edit">Editing</a></li>
				</ul>
			</nav>
		</header>
	
		`)
//line templates/admin_layout.qtpl:29
	p.StreamContent(qw422016)
//line templates/admin_layout.qtpl:29
	qw422016.N().S(`
	</body>
</html>
`)
//line templates/admin_layout.qtpl:32
}

//line templates/admin_layout.qtpl:32
func WriteAdminPageTemplate(qq422016 qtio422016.Writer, p AdminPage) {
//line templates/admin_layout.qtpl:32
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/admin_layout.qtpl:32
	StreamAdminPageTemplate(qw422016, p)
//line templates/admin_layout.qtpl:32
	qt422016.ReleaseWriter(qw422016)
//line templates/admin_layout.qtpl:32
}

//line templates/admin_layout.qtpl:32
func AdminPageTemplate(p AdminPage) string {
//line templates/admin_layout.qtpl:32
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/admin_layout.qtpl:32
	WriteAdminPageTemplate(qb422016, p)
//line templates/admin_layout.qtpl:32
	qs422016 := string(qb422016.B)
//line templates/admin_layout.qtpl:32
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/admin_layout.qtpl:32
	return qs422016
//line templates/admin_layout.qtpl:32
}
