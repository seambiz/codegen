package codegen

// TType template
func TEnd(bb *GenBuffer) {

	bb.NewLine()
	bb.S("// ^^ END OF GENERATED BY CODEGEN. DO NOT EDIT. ^^")
	bb.NewLine()
}