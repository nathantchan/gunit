package generate

import "strings"

const GeneratedFilename = "generated_by_gunit_test.go"

//////////////////////////////////////////////////////////////////////////////

var rawTestFunction = strings.TrimSpace(`
{{range .TestCases}}
func Test_{{$.StructName}}__{{.Name}}(t *testing.T) { {{if .Skipped}}
	t.SkipNow()
	{{else if .LongRunning}}if testing.Short() {
		t.SkipNow()
	}
	{{end}}
	t.Parallel()
	fixture := gunit.NewFixture(t, testing.Verbose())
	defer fixture.Finalize()
	test := &{{$.StructName}}{Fixture: fixture}{{if $.TestTeardownName}}
	defer test.{{$.TestTeardownName}}(){{end}}{{if $.TestSetupName}}
	test.{{$.TestSetupName}}(){{end}}
	test.{{.Name}}()
}
{{else}}
func Test_{{$.StructName}}(t *testing.T) {
	t.Skip("Fixture '{{$.StructName}}' has no test cases.")
}
{{end}}
`)

//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

const header = `//////////////////////////////////////////////////////////////////////////////
// Generated Code ////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////

package %s

import (
	"testing"

	"github.com/smartystreets/gunit"
)
`

const footer = `

func init() {
	gunit.Validate("%s")
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////// Generated Code //
///////////////////////////////////////////////////////////////////////////////
`

//////////////////////////////////////////////////////////////////////////////
