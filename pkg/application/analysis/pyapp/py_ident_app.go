package pyapp

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/python"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_python"
)

func streamToParser(is antlr.CharStream) *parser.PythonParser {
	lexer := parser.NewPythonLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return parser.NewPythonParser(tokens)
}

func ProcessTsString(code string) *parser.PythonParser {
	is := antlr.NewInputStream(code)
	return streamToParser(is)
}

type PythonIdentApp struct {

}

func (p *PythonIdentApp) AnalysisPackageManager(path string) core_domain.CodePackageManagerInfo  {
	return core_domain.CodePackageManagerInfo{}
}

func (p *PythonIdentApp) Analysis(code string, fileName string) core_domain.CodeFile {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Root()

	listener := ast_python.NewPythonIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetCodeFileInfo()
}

func (p *PythonIdentApp) SetExtensions(extension interface{})  {

}

func (p *PythonIdentApp) IdentAnalysis(s string, file string) []core_domain.CodeMember {
	return nil
}
