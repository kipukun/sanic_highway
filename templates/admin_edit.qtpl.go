// Code generated by qtc from "admin_edit.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/admin_edit.qtpl:1
package templates

//line templates/admin_edit.qtpl:1
import "fmt"

//line templates/admin_edit.qtpl:2
import "strings"

//line templates/admin_edit.qtpl:3
import "github.com/kipukun/sanic_highway/model"

//line templates/admin_edit.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/admin_edit.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/admin_edit.qtpl:6
type AdminEditPage struct {
	User                *model.User
	EroList             []model.Eroge
	Pagination          []int
	Current, Prev, Next int
}

//line templates/admin_edit.qtpl:14
func (p *AdminEditPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/admin_edit.qtpl:14
	qw422016.N().S(`
<title>highway maintenance</title>
`)
//line templates/admin_edit.qtpl:16
}

//line templates/admin_edit.qtpl:16
func (p *AdminEditPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/admin_edit.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/admin_edit.qtpl:16
	p.StreamTitle(qw422016)
//line templates/admin_edit.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line templates/admin_edit.qtpl:16
}

//line templates/admin_edit.qtpl:16
func (p *AdminEditPage) Title() string {
//line templates/admin_edit.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/admin_edit.qtpl:16
	p.WriteTitle(qb422016)
//line templates/admin_edit.qtpl:16
	qs422016 := string(qb422016.B)
//line templates/admin_edit.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/admin_edit.qtpl:16
	return qs422016
//line templates/admin_edit.qtpl:16
}

//line templates/admin_edit.qtpl:19
func (p *AdminEditPage) StreamContent(qw422016 *qt422016.Writer) {
//line templates/admin_edit.qtpl:19
	qw422016.N().S(`
<marquee><h1>editing</h1></marquee>
<table>
<tr>
	<th>Title</th>
	<th>
		MetaIDs 
		<button onclick=both("vndb")>Fill VNDB</button>
		<button onclick=both("dlsite")>Fill DLsite</button>
	</th>
</tr>
`)
//line templates/admin_edit.qtpl:30
	for _, ero := range p.EroList {
//line templates/admin_edit.qtpl:30
		qw422016.N().S(`
<tr>
	<td>
		`)
//line templates/admin_edit.qtpl:33
		qw422016.E().S(ero.Filename)
//line templates/admin_edit.qtpl:33
		qw422016.N().S(`
		<input type="button" value="Copy" />
	</td>
	<td>
		`)
//line templates/admin_edit.qtpl:37
		for t, d := range ero.Meta {
//line templates/admin_edit.qtpl:37
			qw422016.N().S(`
			<span><b>`)
//line templates/admin_edit.qtpl:38
			qw422016.E().S(t)
//line templates/admin_edit.qtpl:38
			qw422016.N().S(`:</b> `)
//line templates/admin_edit.qtpl:38
			qw422016.E().S(fmt.Sprintf(strings.Join(d, ", ")))
//line templates/admin_edit.qtpl:38
			qw422016.N().S(`</span>
		`)
//line templates/admin_edit.qtpl:39
		}
//line templates/admin_edit.qtpl:39
		qw422016.N().S(`
		<form action="/api/edit/`)
//line templates/admin_edit.qtpl:40
		qw422016.N().D(ero.ID)
//line templates/admin_edit.qtpl:40
		qw422016.N().S(`" method="POST">
			<input type="hidden" name="page" value="`)
//line templates/admin_edit.qtpl:41
		qw422016.N().D(p.Current)
//line templates/admin_edit.qtpl:41
		qw422016.N().S(`">
			<select name="group">
				`)
//line templates/admin_edit.qtpl:43
		for t, _ := range ero.Meta {
//line templates/admin_edit.qtpl:43
			qw422016.N().S(`
				<option value="`)
//line templates/admin_edit.qtpl:44
			qw422016.E().S(t)
//line templates/admin_edit.qtpl:44
			qw422016.N().S(`">`)
//line templates/admin_edit.qtpl:44
			qw422016.E().S(t)
//line templates/admin_edit.qtpl:44
			qw422016.N().S(`</option>
				`)
//line templates/admin_edit.qtpl:45
		}
//line templates/admin_edit.qtpl:45
		qw422016.N().S(`
				<option value="new">create new</option>
			</select>
			<input name="new-group" type="text" placeholder="new id group">
			<input name="id" type="text" placeholder="id" required>
			<input name="op" type="submit" value="+">
			<input name="op" type="submit" value="-">
		</form>
	</td>
</tr>
`)
//line templates/admin_edit.qtpl:55
	}
//line templates/admin_edit.qtpl:55
	qw422016.N().S(`
</table>
<center>
	<div class="pg">
		<a href="/admin/edit/page/`)
//line templates/admin_edit.qtpl:59
	qw422016.N().D(p.Prev)
//line templates/admin_edit.qtpl:59
	qw422016.N().S(`">&ltrif;</a>
		`)
//line templates/admin_edit.qtpl:60
	for _, pg := range p.Pagination {
//line templates/admin_edit.qtpl:60
		qw422016.N().S(`
			<a href="/admin/edit/page/`)
//line templates/admin_edit.qtpl:61
		qw422016.N().D(pg)
//line templates/admin_edit.qtpl:61
		qw422016.N().S(`" `)
//line templates/admin_edit.qtpl:61
		if pg == p.Current {
//line templates/admin_edit.qtpl:61
			qw422016.N().S(` class="current" `)
//line templates/admin_edit.qtpl:61
		}
//line templates/admin_edit.qtpl:61
		qw422016.N().S(`>`)
//line templates/admin_edit.qtpl:61
		qw422016.N().D(pg)
//line templates/admin_edit.qtpl:61
		qw422016.N().S(`</a>
		`)
//line templates/admin_edit.qtpl:62
	}
//line templates/admin_edit.qtpl:62
	qw422016.N().S(`
		<a href="/admin/edit/page/`)
//line templates/admin_edit.qtpl:63
	qw422016.N().D(p.Next)
//line templates/admin_edit.qtpl:63
	qw422016.N().S(`">&rtrif;</a>
</center>

<script type="text/javascript">
document.addEventListener('DOMContentLoaded', function() {
	if (typeof(Storage) !== "undefined") {
		fill()	
	} else {
		console.log("no localStorage support, for some reason.")
	}
}, false);
function fill() {
	f = localStorage.getItem("fill")
	if (f !== null) {
		groups = document.querySelectorAll('[name="new-group"]')
		for (let el of groups) {el.value = f} 
	}
}
function both(f) {
	localStorage.setItem("fill", f)
	fill()
}
</script>
`)
//line templates/admin_edit.qtpl:86
}

//line templates/admin_edit.qtpl:86
func (p *AdminEditPage) WriteContent(qq422016 qtio422016.Writer) {
//line templates/admin_edit.qtpl:86
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/admin_edit.qtpl:86
	p.StreamContent(qw422016)
//line templates/admin_edit.qtpl:86
	qt422016.ReleaseWriter(qw422016)
//line templates/admin_edit.qtpl:86
}

//line templates/admin_edit.qtpl:86
func (p *AdminEditPage) Content() string {
//line templates/admin_edit.qtpl:86
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/admin_edit.qtpl:86
	p.WriteContent(qb422016)
//line templates/admin_edit.qtpl:86
	qs422016 := string(qb422016.B)
//line templates/admin_edit.qtpl:86
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/admin_edit.qtpl:86
	return qs422016
//line templates/admin_edit.qtpl:86
}
