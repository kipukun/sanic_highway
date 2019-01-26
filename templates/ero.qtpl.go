// This file is automatically generated by qtc from "ero.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/ero.qtpl:1
package templates

//line templates/ero.qtpl:1
import "github.com/kipukun/sanic_highway/model"

//line templates/ero.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/ero.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/ero.qtpl:4
type ErogePage struct {
	Ero  model.Eroge
	Tags []string
}

//line templates/ero.qtpl:11
func (p *ErogePage) StreamTitle(qw422016 *qt422016.Writer) {
	//line templates/ero.qtpl:11
	qw422016.N().S(`
<title>`)
	//line templates/ero.qtpl:12
	qw422016.E().S(p.Ero.Title)
	//line templates/ero.qtpl:12
	qw422016.N().S(`</title>
`)
//line templates/ero.qtpl:13
}

//line templates/ero.qtpl:13
func (p *ErogePage) WriteTitle(qq422016 qtio422016.Writer) {
	//line templates/ero.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/ero.qtpl:13
	p.StreamTitle(qw422016)
	//line templates/ero.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line templates/ero.qtpl:13
}

//line templates/ero.qtpl:13
func (p *ErogePage) Title() string {
	//line templates/ero.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/ero.qtpl:13
	p.WriteTitle(qb422016)
	//line templates/ero.qtpl:13
	qs422016 := string(qb422016.B)
	//line templates/ero.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/ero.qtpl:13
	return qs422016
//line templates/ero.qtpl:13
}

//line templates/ero.qtpl:16
func (p *ErogePage) StreamContent(qw422016 *qt422016.Writer) {
	//line templates/ero.qtpl:16
	qw422016.N().S(`

<article>
	<h1>`)
	//line templates/ero.qtpl:19
	qw422016.E().S(p.Ero.Title)
	//line templates/ero.qtpl:19
	qw422016.N().S(`</h1>
	<p>Made by `)
	//line templates/ero.qtpl:20
	qw422016.E().S(p.Ero.CircleName)
	//line templates/ero.qtpl:20
	qw422016.N().S(`</p>
	`)
	//line templates/ero.qtpl:21
	for _, img := range p.Ero.Images {
		//line templates/ero.qtpl:21
		qw422016.N().S(`
		<img src="`)
		//line templates/ero.qtpl:22
		qw422016.E().S(img)
		//line templates/ero.qtpl:22
		qw422016.N().S(`"/>
	`)
		//line templates/ero.qtpl:23
	}
	//line templates/ero.qtpl:23
	qw422016.N().S(`
	<h3>Tags</h3>
	`)
	//line templates/ero.qtpl:25
	for _, tag := range p.Tags {
		//line templates/ero.qtpl:25
		qw422016.N().S(`
		<span>`)
		//line templates/ero.qtpl:26
		qw422016.E().S(tag)
		//line templates/ero.qtpl:26
		qw422016.N().S(`</span>
	`)
		//line templates/ero.qtpl:27
	}
	//line templates/ero.qtpl:27
	qw422016.N().S(`
</article>
`)
//line templates/ero.qtpl:29
}

//line templates/ero.qtpl:29
func (p *ErogePage) WriteContent(qq422016 qtio422016.Writer) {
	//line templates/ero.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/ero.qtpl:29
	p.StreamContent(qw422016)
	//line templates/ero.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line templates/ero.qtpl:29
}

//line templates/ero.qtpl:29
func (p *ErogePage) Content() string {
	//line templates/ero.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/ero.qtpl:29
	p.WriteContent(qb422016)
	//line templates/ero.qtpl:29
	qs422016 := string(qb422016.B)
	//line templates/ero.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/ero.qtpl:29
	return qs422016
//line templates/ero.qtpl:29
}
