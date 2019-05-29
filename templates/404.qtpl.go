// Code generated by qtc from "404.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line templates/404.qtpl:1
package templates

//line templates/404.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/404.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/404.qtpl:2
type StopPage struct{}

//line templates/404.qtpl:5
func (p *StopPage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/404.qtpl:5
	qw422016.N().S(`
<title>turn back</title>
`)
//line templates/404.qtpl:7
}

//line templates/404.qtpl:7
func (p *StopPage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/404.qtpl:7
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/404.qtpl:7
	p.StreamTitle(qw422016)
//line templates/404.qtpl:7
	qt422016.ReleaseWriter(qw422016)
//line templates/404.qtpl:7
}

//line templates/404.qtpl:7
func (p *StopPage) Title() string {
//line templates/404.qtpl:7
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/404.qtpl:7
	p.WriteTitle(qb422016)
//line templates/404.qtpl:7
	qs422016 := string(qb422016.B)
//line templates/404.qtpl:7
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/404.qtpl:7
	return qs422016
//line templates/404.qtpl:7
}

//line templates/404.qtpl:9
func (p *StopPage) StreamContent(qw422016 *qt422016.Writer) {
//line templates/404.qtpl:9
	qw422016.N().S(`
<div class="stop">
	<img src="/assets/image/stop.png" >
	<h2>You must turn back.</h2>
</div>
`)
//line templates/404.qtpl:14
}

//line templates/404.qtpl:14
func (p *StopPage) WriteContent(qq422016 qtio422016.Writer) {
//line templates/404.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/404.qtpl:14
	p.StreamContent(qw422016)
//line templates/404.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line templates/404.qtpl:14
}

//line templates/404.qtpl:14
func (p *StopPage) Content() string {
//line templates/404.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/404.qtpl:14
	p.WriteContent(qb422016)
//line templates/404.qtpl:14
	qs422016 := string(qb422016.B)
//line templates/404.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/404.qtpl:14
	return qs422016
//line templates/404.qtpl:14
}
