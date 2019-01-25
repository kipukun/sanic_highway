// This file is automatically generated by qtc from "index.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/index.qtpl:1
package templates

//line templates/index.qtpl:1
import "github.com/kipukun/sanic_highway/model"

//line templates/index.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/index.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/index.qtpl:4
type IndexPage struct {
	EroList             []model.Eroge
	Pagination          [5]int
	Current, Prev, Next int
}

//line templates/index.qtpl:11
func (p *IndexPage) StreamTitle(qw422016 *qt422016.Writer) {
	//line templates/index.qtpl:11
	qw422016.N().S(`
<title>sanic information highway</title>
`)
//line templates/index.qtpl:13
}

//line templates/index.qtpl:13
func (p *IndexPage) WriteTitle(qq422016 qtio422016.Writer) {
	//line templates/index.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/index.qtpl:13
	p.StreamTitle(qw422016)
	//line templates/index.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:13
}

//line templates/index.qtpl:13
func (p *IndexPage) Title() string {
	//line templates/index.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/index.qtpl:13
	p.WriteTitle(qb422016)
	//line templates/index.qtpl:13
	qs422016 := string(qb422016.B)
	//line templates/index.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/index.qtpl:13
	return qs422016
//line templates/index.qtpl:13
}

//line templates/index.qtpl:16
func (p *IndexPage) StreamContent(qw422016 *qt422016.Writer) {
	//line templates/index.qtpl:16
	qw422016.N().S(`
<table>
	<tr>
		<th>Circle</th>
		<th>Title</th>
		<th>DLsite</th>
		<th>VNDB</th>
		<th>Misc.</th>
		<th>Date</th>
		<th>XDCC</th>
		<th>HDD</th>
		<th>Torrent</th>
	</tr>
	`)
	//line templates/index.qtpl:29
	for _, ero := range p.EroList {
		//line templates/index.qtpl:29
		qw422016.N().S(`
	<tr>
		<td><a href="/circle/`)
		//line templates/index.qtpl:31
		qw422016.N().D(ero.Circle.ID)
		//line templates/index.qtpl:31
		qw422016.N().S(`">`)
		//line templates/index.qtpl:31
		qw422016.E().S(ero.CircleName)
		//line templates/index.qtpl:31
		qw422016.N().S(`</a></td>
		<td><a href="/ero/`)
		//line templates/index.qtpl:32
		qw422016.N().D(ero.ID)
		//line templates/index.qtpl:32
		qw422016.N().S(`">`)
		//line templates/index.qtpl:32
		qw422016.E().S(ero.Title)
		//line templates/index.qtpl:32
		qw422016.N().S(`</a></td>
		<td>
			`)
		//line templates/index.qtpl:34
		for _, id := range ero.DLsiteIDs {
			//line templates/index.qtpl:34
			qw422016.N().S(`
				<p>`)
			//line templates/index.qtpl:35
			qw422016.E().S(id)
			//line templates/index.qtpl:35
			qw422016.N().S(`</p>
			`)
			//line templates/index.qtpl:36
		}
		//line templates/index.qtpl:36
		qw422016.N().S(`
		</td>
		<td>
			`)
		//line templates/index.qtpl:39
		for _, id := range ero.VNDBIDs {
			//line templates/index.qtpl:39
			qw422016.N().S(`
				<p>`)
			//line templates/index.qtpl:40
			qw422016.E().S(id)
			//line templates/index.qtpl:40
			qw422016.N().S(`</p>
			`)
			//line templates/index.qtpl:41
		}
		//line templates/index.qtpl:41
		qw422016.N().S(`
		</td>
		<td>
			`)
		//line templates/index.qtpl:44
		for _, id := range ero.MiscIDs {
			//line templates/index.qtpl:44
			qw422016.N().S(`
				<p>`)
			//line templates/index.qtpl:45
			qw422016.E().S(id)
			//line templates/index.qtpl:45
			qw422016.N().S(`</p>
			`)
			//line templates/index.qtpl:46
		}
		//line templates/index.qtpl:46
		qw422016.N().S(`
		</td>

		<td>`)
		//line templates/index.qtpl:49
		qw422016.E().S(ero.Date)
		//line templates/index.qtpl:49
		qw422016.N().S(`</td>
		<td>`)
		//line templates/index.qtpl:50
		qw422016.E().V(ero.OnXDCC)
		//line templates/index.qtpl:50
		qw422016.N().S(`</td>
		<td>`)
		//line templates/index.qtpl:51
		qw422016.E().V(ero.OnHDD)
		//line templates/index.qtpl:51
		qw422016.N().S(`</td>
		<td>`)
		//line templates/index.qtpl:52
		qw422016.E().V(ero.InTorrent)
		//line templates/index.qtpl:52
		qw422016.N().S(`</td>
	</tr>
	`)
		//line templates/index.qtpl:54
	}
	//line templates/index.qtpl:54
	qw422016.N().S(`
</table>
<center>
	<div class="pg">
		<a href="/page/`)
	//line templates/index.qtpl:58
	qw422016.N().D(p.Prev)
	//line templates/index.qtpl:58
	qw422016.N().S(`">◀</a>
		`)
	//line templates/index.qtpl:59
	for _, pg := range p.Pagination {
		//line templates/index.qtpl:59
		qw422016.N().S(`
			<a href="/page/`)
		//line templates/index.qtpl:60
		qw422016.N().D(pg)
		//line templates/index.qtpl:60
		qw422016.N().S(`" `)
		//line templates/index.qtpl:60
		if pg == p.Current {
			//line templates/index.qtpl:60
			qw422016.N().S(` class="current" `)
			//line templates/index.qtpl:60
		}
		//line templates/index.qtpl:60
		qw422016.N().S(`>`)
		//line templates/index.qtpl:60
		qw422016.N().D(pg)
		//line templates/index.qtpl:60
		qw422016.N().S(`</a>
		`)
		//line templates/index.qtpl:61
	}
	//line templates/index.qtpl:61
	qw422016.N().S(`
		<a href="/page/`)
	//line templates/index.qtpl:62
	qw422016.N().D(p.Next)
	//line templates/index.qtpl:62
	qw422016.N().S(`">▶</a>
</center>
`)
//line templates/index.qtpl:64
}

//line templates/index.qtpl:64
func (p *IndexPage) WriteContent(qq422016 qtio422016.Writer) {
	//line templates/index.qtpl:64
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/index.qtpl:64
	p.StreamContent(qw422016)
	//line templates/index.qtpl:64
	qt422016.ReleaseWriter(qw422016)
//line templates/index.qtpl:64
}

//line templates/index.qtpl:64
func (p *IndexPage) Content() string {
	//line templates/index.qtpl:64
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/index.qtpl:64
	p.WriteContent(qb422016)
	//line templates/index.qtpl:64
	qs422016 := string(qb422016.B)
	//line templates/index.qtpl:64
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/index.qtpl:64
	return qs422016
//line templates/index.qtpl:64
}
