// Code generated by qtc from "content.xml.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line content.xml.qtpl:1
package ods

//line content.xml.qtpl:1
import "strings"

//line content.xml.qtpl:2
import "encoding/xml"

//line content.xml.qtpl:3
import "time"

//line content.xml.qtpl:4
import "fmt"

//line content.xml.qtpl:5
import "github.com/tgulacsi/go/spreadsheet"

//line content.xml.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line content.xml.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line content.xml.qtpl:8
func StreamXML(qw422016 *qt422016.Writer, s string) {
//line content.xml.qtpl:10
	var buf strings.Builder
	_ = xml.EscapeText(&buf, []byte(s))

//line content.xml.qtpl:13
	qw422016.N().S(buf.String())
//line content.xml.qtpl:14
}

//line content.xml.qtpl:14
func WriteXML(qq422016 qtio422016.Writer, s string) {
//line content.xml.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:14
	StreamXML(qw422016, s)
//line content.xml.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:14
}

//line content.xml.qtpl:14
func XML(s string) string {
//line content.xml.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:14
	WriteXML(qb422016, s)
//line content.xml.qtpl:14
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:14
	return qs422016
//line content.xml.qtpl:14
}

//line content.xml.qtpl:15
func streamgetValueString(qw422016 *qt422016.Writer, v interface{}) {
//line content.xml.qtpl:17
	var buf strings.Builder
	switch x := v.(type) {
	case time.Time:
		if !x.IsZero() {
			buf.WriteString(x.Format(time.RFC3339))
		}
	case int, int8, int16, int32, int64, uint, uint16, uint32, uint64:
		fmt.Fprintf(&buf, "%d", v)
	case float32, float64:
		fmt.Fprintf(&buf, "%f", v)
	case string:
		_ = xml.EscapeText(&buf, []byte(x))
	case fmt.Stringer:
		_ = xml.EscapeText(&buf, []byte(x.String()))
	default:
		_ = xml.EscapeText(&buf, []byte(fmt.Sprintf("%v", v)))
	}

//line content.xml.qtpl:34
	qw422016.N().S(buf.String())
//line content.xml.qtpl:35
}

//line content.xml.qtpl:35
func writegetValueString(qq422016 qtio422016.Writer, v interface{}) {
//line content.xml.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:35
	streamgetValueString(qw422016, v)
//line content.xml.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:35
}

//line content.xml.qtpl:35
func getValueString(v interface{}) string {
//line content.xml.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:35
	writegetValueString(qb422016, v)
//line content.xml.qtpl:35
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:35
	return qs422016
//line content.xml.qtpl:35
}

//line content.xml.qtpl:38
func StreamBeginSpreadsheet(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:38
	qw422016.N().S(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-content xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0" xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0" xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0" xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0" xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0" xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0" xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0" xmlns:config="urn:oasis:names:tc:opendocument:xmlns:config:1.0" xmlns:math="http://www.w3.org/1998/Math/MathML" xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0" xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0" xmlns:ooo="http://openoffice.org/2004/office" xmlns:ooow="http://openoffice.org/2004/writer" xmlns:oooc="http://openoffice.org/2004/calc" xmlns:tableooo="http://openoffice.org/2009/table" xmlns:of="urn:oasis:names:tc:opendocument:xmlns:of:1.2" xmlns:dom="http://www.w3.org/2001/xml-events" xmlns:xforms="http://www.w3.org/2002/xforms" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:gnm="http://www.gnumeric.org/odf-extension/1.0" xmlns:css3t="http://www.w3.org/TR/css3-text/" xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0" xmlns:calcext="urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0" office:version="1.2">
  <office:scripts/>
  <office:font-face-decls/>
  <office:automatic-styles/>
  <office:body>
    <office:spreadsheet>
      <table:calculation-settings table:null-year="1930" table:automatic-find-labels="false" table:case-sensitive="false" table:precision-as-shown="false" table:search-criteria-must-apply-to-whole-cell="true" table:use-regular-expressions="false" table:use-wildcards="false">
        <table:null-date table:date-value="1899-12-30" table:value-type="date"/>
        <table:iteration table:maximum-difference="0.001" table:status="enable" table:steps="100"/>
      </table:calculation-settings>
`)
//line content.xml.qtpl:49
}

//line content.xml.qtpl:49
func WriteBeginSpreadsheet(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:49
	StreamBeginSpreadsheet(qw422016)
//line content.xml.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:49
}

//line content.xml.qtpl:49
func BeginSpreadsheet() string {
//line content.xml.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:49
	WriteBeginSpreadsheet(qb422016)
//line content.xml.qtpl:49
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:49
	return qs422016
//line content.xml.qtpl:49
}

//line content.xml.qtpl:51
func (ow *ODSWriter) StreamBeginSheet(qw422016 *qt422016.Writer, name string, cols []spreadsheet.Column) {
//line content.xml.qtpl:51
	qw422016.N().S(`<table:table table:name="`)
//line content.xml.qtpl:51
	StreamXML(qw422016, name)
//line content.xml.qtpl:51
	qw422016.N().S(`" table:print="true">`)
//line content.xml.qtpl:52
	var hasHeader bool

//line content.xml.qtpl:53
	for _, c := range cols {
//line content.xml.qtpl:53
		qw422016.N().S(`<table:table-column table:style-name="`)
//line content.xml.qtpl:53
		qw422016.E().S(ow.getStyleName(c.Column))
//line content.xml.qtpl:53
		qw422016.N().S(`" />`)
//line content.xml.qtpl:53
		if c.Name != "" {
			hasHeader = true
		}

//line content.xml.qtpl:54
	}
//line content.xml.qtpl:55
	if hasHeader {
//line content.xml.qtpl:55
		qw422016.N().S(`<table:table-row>`)
//line content.xml.qtpl:56
		for _, c := range cols {
//line content.xml.qtpl:56
			qw422016.N().S(`<table:table-cell office:value-type="string" table:style-name="`)
//line content.xml.qtpl:56
			qw422016.N().S(ow.getStyleName(c.Header))
//line content.xml.qtpl:56
			qw422016.N().S(`"><text:p>`)
//line content.xml.qtpl:56
			StreamXML(qw422016, c.Name)
//line content.xml.qtpl:56
			qw422016.N().S(`</text:p></table:table-cell>`)
//line content.xml.qtpl:57
		}
//line content.xml.qtpl:57
		qw422016.N().S(`</table:table-row>`)
//line content.xml.qtpl:58
	}
//line content.xml.qtpl:58
	qw422016.N().S(`
`)
//line content.xml.qtpl:59
}

//line content.xml.qtpl:59
func (ow *ODSWriter) WriteBeginSheet(qq422016 qtio422016.Writer, name string, cols []spreadsheet.Column) {
//line content.xml.qtpl:59
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:59
	ow.StreamBeginSheet(qw422016, name, cols)
//line content.xml.qtpl:59
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:59
}

//line content.xml.qtpl:59
func (ow *ODSWriter) BeginSheet(name string, cols []spreadsheet.Column) string {
//line content.xml.qtpl:59
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:59
	ow.WriteBeginSheet(qb422016, name, cols)
//line content.xml.qtpl:59
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:59
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:59
	return qs422016
//line content.xml.qtpl:59
}

//line content.xml.qtpl:61
func (ow *ODSWriter) StreamEndSheet(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:61
	qw422016.N().S(`
      </table:table>
`)
//line content.xml.qtpl:63
}

//line content.xml.qtpl:63
func (ow *ODSWriter) WriteEndSheet(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:63
	ow.StreamEndSheet(qw422016)
//line content.xml.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:63
}

//line content.xml.qtpl:63
func (ow *ODSWriter) EndSheet() string {
//line content.xml.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:63
	ow.WriteEndSheet(qb422016)
//line content.xml.qtpl:63
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:63
	return qs422016
//line content.xml.qtpl:63
}

//line content.xml.qtpl:66
func StreamRow(qw422016 *qt422016.Writer, values ...interface{}) {
//line content.xml.qtpl:66
	qw422016.N().S(`<table:table-row>`)
//line content.xml.qtpl:67
	for _, v := range values {
//line content.xml.qtpl:67
		typ := getValueType(v)

//line content.xml.qtpl:67
		qw422016.N().S(`
	<table:table-cell office:value-type="`)
//line content.xml.qtpl:68
		qw422016.N().S(typ.String())
//line content.xml.qtpl:68
		qw422016.N().S(`"`)
//line content.xml.qtpl:69
		if typ == FloatType {
//line content.xml.qtpl:69
			qw422016.N().S(` office:value="`)
//line content.xml.qtpl:69
			qw422016.N().S(fmt.Sprintf("%v", v))
//line content.xml.qtpl:69
			qw422016.N().S(`"`)
//line content.xml.qtpl:70
		} else if typ == DateType {
//line content.xml.qtpl:70
			qw422016.N().S(` office:date-value="getValueString(v)"`)
//line content.xml.qtpl:71
		}
//line content.xml.qtpl:71
		qw422016.N().S(`><text:p>`)
//line content.xml.qtpl:71
		streamgetValueString(qw422016, v)
//line content.xml.qtpl:71
		qw422016.N().S(`</text:p></table:table-cell>`)
//line content.xml.qtpl:72
	}
//line content.xml.qtpl:72
	qw422016.N().S(`</table:table-row>
`)
//line content.xml.qtpl:73
}

//line content.xml.qtpl:73
func WriteRow(qq422016 qtio422016.Writer, values ...interface{}) {
//line content.xml.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:73
	StreamRow(qw422016, values...)
//line content.xml.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:73
}

//line content.xml.qtpl:73
func Row(values ...interface{}) string {
//line content.xml.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:73
	WriteRow(qb422016, values...)
//line content.xml.qtpl:73
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:73
	return qs422016
//line content.xml.qtpl:73
}

//line content.xml.qtpl:75
func StreamEndSpreadsheet(qw422016 *qt422016.Writer) {
//line content.xml.qtpl:75
	qw422016.N().S(`
    </office:spreadsheet>
  </office:body>
</office:document-content>
`)
//line content.xml.qtpl:79
}

//line content.xml.qtpl:79
func WriteEndSpreadsheet(qq422016 qtio422016.Writer) {
//line content.xml.qtpl:79
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:79
	StreamEndSpreadsheet(qw422016)
//line content.xml.qtpl:79
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:79
}

//line content.xml.qtpl:79
func EndSpreadsheet() string {
//line content.xml.qtpl:79
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:79
	WriteEndSpreadsheet(qb422016)
//line content.xml.qtpl:79
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:79
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:79
	return qs422016
//line content.xml.qtpl:79
}

//line content.xml.qtpl:81
func StreamStyles(qw422016 *qt422016.Writer, styles map[string]string) {
//line content.xml.qtpl:81
	qw422016.N().S(`<?xml version="1.0" encoding="UTF-8"?>
<office:document-styles xmlns:office="urn:oasis:names:tc:opendocument:xmlns:office:1.0" xmlns:style="urn:oasis:names:tc:opendocument:xmlns:style:1.0" xmlns:text="urn:oasis:names:tc:opendocument:xmlns:text:1.0" xmlns:table="urn:oasis:names:tc:opendocument:xmlns:table:1.0" xmlns:draw="urn:oasis:names:tc:opendocument:xmlns:drawing:1.0" xmlns:fo="urn:oasis:names:tc:opendocument:xmlns:xsl-fo-compatible:1.0" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:meta="urn:oasis:names:tc:opendocument:xmlns:meta:1.0" xmlns:number="urn:oasis:names:tc:opendocument:xmlns:datastyle:1.0" xmlns:svg="urn:oasis:names:tc:opendocument:xmlns:svg-compatible:1.0" xmlns:chart="urn:oasis:names:tc:opendocument:xmlns:chart:1.0" xmlns:dr3d="urn:oasis:names:tc:opendocument:xmlns:dr3d:1.0" xmlns:config="urn:oasis:names:tc:opendocument:xmlns:config:1.0" xmlns:math="http://www.w3.org/1998/Math/MathML" xmlns:form="urn:oasis:names:tc:opendocument:xmlns:form:1.0" xmlns:script="urn:oasis:names:tc:opendocument:xmlns:script:1.0" xmlns:ooo="http://openoffice.org/2004/office" xmlns:ooow="http://openoffice.org/2004/writer" xmlns:oooc="http://openoffice.org/2004/calc" xmlns:tableooo="http://openoffice.org/2009/table" xmlns:of="urn:oasis:names:tc:opendocument:xmlns:of:1.2" xmlns:dom="http://www.w3.org/2001/xml-events" xmlns:xforms="http://www.w3.org/2002/xforms" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:gnm="http://www.gnumeric.org/odf-extension/1.0" xmlns:css3t="http://www.w3.org/TR/css3-text/" xmlns:loext="urn:org:documentfoundation:names:experimental:office:xmlns:loext:1.0" xmlns:calcext="urn:org:documentfoundation:names:experimental:calc:xmlns:calcext:1.0" office:version="1.2">
  <office:styles>
    <style:default-style style:family="table-column">
      <style:table-column-properties style:use-optimal-column-width="true"/>
    </style:default-style>
    <style:default-style style:family="table-row">
      <style:table-row-properties style:use-optimal-row-height="true"/>
    </style:default-style>
  </office:styles>
  <office:automatic-styles>
	`)
//line content.xml.qtpl:92
	for _, s := range styles {
//line content.xml.qtpl:92
		qw422016.N().S(s)
//line content.xml.qtpl:93
	}
//line content.xml.qtpl:93
	qw422016.N().S(`
  </office:automatic-styles>
</office:document-styles>
`)
//line content.xml.qtpl:96
}

//line content.xml.qtpl:96
func WriteStyles(qq422016 qtio422016.Writer, styles map[string]string) {
//line content.xml.qtpl:96
	qw422016 := qt422016.AcquireWriter(qq422016)
//line content.xml.qtpl:96
	StreamStyles(qw422016, styles)
//line content.xml.qtpl:96
	qt422016.ReleaseWriter(qw422016)
//line content.xml.qtpl:96
}

//line content.xml.qtpl:96
func Styles(styles map[string]string) string {
//line content.xml.qtpl:96
	qb422016 := qt422016.AcquireByteBuffer()
//line content.xml.qtpl:96
	WriteStyles(qb422016, styles)
//line content.xml.qtpl:96
	qs422016 := string(qb422016.B)
//line content.xml.qtpl:96
	qt422016.ReleaseByteBuffer(qb422016)
//line content.xml.qtpl:96
	return qs422016
//line content.xml.qtpl:96
}
