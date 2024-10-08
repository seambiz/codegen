package templates

import (
	"github.com/seambiz/codegen"
	"github.com/seambiz/gen"
)

type DBTypes struct{}

func (DBTypes) Generate(conf codegen.Config) string {
	g := gen.NewGenerator()

	g.Package(conf.RootPackage)
	g.Import("math/big")

	g.Comment("GENERATED BY CODEGEN.")
	g.NewLine()

	/**********************************************************************
	* entity types
	**********************************************************************/
	for _, schema := range conf.Schemas {
		for _, t := range schema.Tables {
			if t.Generate {
				/**********************************************************************
				 * basic struct
				**********************************************************************/
				g.Comment(t.Title, " represents a row from ", t.Name, ".")
				g.Struct(t.Title, func() {
					for _, f := range t.Fields {
						g.Field(f.Title).Type(gen.PtrIf(f.IsNullable, f.GoType)).Tag("json", f.Name).Tag("db", f.Name).Go()
					}
					g.NewLine()

					for _, fk := range t.ForeignKeys {
						g.Field(fk.CustomName).Slice(!fk.IsUnique).PointerType(fk.RefTableTitle).Tag("json", fk.CustomName+",omitempty").Go()
					}
				})

				/**********************************************************************
				 * audit
				**********************************************************************/
				if t.Audit {
					g.Method("Diff").Receiver(t.Initials, t.Title).Param("other", t.Title).Returns("[]Change").Body(func() {
						g.Go("changes := []Change{}")

						for _, f := range t.Fields {
							if !f.IsPrimaryKey {
								g.If(g.Id(t.Initials, f.Title), "!=", g.Id("other", f.Title)).Body(func() {
									g.Append("changes", gen.NewStruct("Change", &gen.KV{
										"Field":    f.Name,
										"OldValue": gen.Sprintf("%v", g.Id("other", f.Title)),
										"NewValue": gen.Sprintf("%v", g.Id(t.Initials, f.Title)),
									}))
								})
							}
						}

						g.Return("changes")
					})

					g.Method("Apply").Receiver(t.Initials, t.Title).Param("other", "*"+t.Title).Body(func() {
						g.Panic("not implemented")
					})

					g.Method("Key").Receiver(t.Initials, t.Title).Returns("int64").Body(func() {
						g.Return(g.Id(t.Initials, "UUID"))
					})
				}
			}
		}
	}

	/**********************************************************************
	* partial entity types
	**********************************************************************/
	for _, schema := range conf.Schemas {
		for _, t := range schema.Tables {
			if t.Generate {
				/**********************************************************************
				 * basic struct
				**********************************************************************/
				g.Comment(t.Title, "Partial is used for updating specific columns from ", t.Name, ".")
				g.Struct(g.S(t.Title, "Partial"), func() {
					g.Field(t.Title).Go()
					g.Field("Touched").Type("big.Int").Go()
				})

				/**********************************************************************
				 * setter for applying bit
				**********************************************************************/
				for _, f := range t.Fields {
					g.Method("Set", f.Title).Receiver("p", gen.Ptr(t.Title, "Partial")).Param("param", gen.PtrIf(f.IsNullable, f.GoType)).Body(func() {
						g.Go("p.", f.Title, " = param")
						g.Go("p.Touched.SetBit(&p.Touched, ", t.Title, "_", f.Title, ", 1)")
					})
				}

				/**********************************************************************
				 * audit
				**********************************************************************/
				if t.Audit {
					g.Method("Diff").Receiver(t.Initials, g.S(t.Title, "Partial")).Param("other", t.Title).Returns("[]Change").Body(func() {
						g.Go("changes := []Change{}")

						for _, f := range t.Fields {
							if !f.IsPrimaryKey {
								g.If(g.S(g.Id(t.Initials, "Touched"), ".Bit(", t.Title, "_", f.Title), "==", "1").Body(func() {
									g.If(g.Id(t.Initials, f.Title), "!=", g.Id("other", f.Title)).Body(func() {
										g.Append("changes", gen.NewStruct("Change", &gen.KV{
											"Field":    f.Name,
											"OldValue": gen.Sprintf("%v", g.Id("other", f.Title)),
											"NewValue": gen.Sprintf("%v", g.Id(t.Initials, f.Title)),
										}))
									})
								})
							}
						}

						g.Return("changes")
					})

					g.Method("Apply").Receiver(t.Initials, t.Title).Param("other", gen.Ptr(t.Title)).Body(func() {
						for _, f := range t.Fields {
							if !(f.IsPrimaryKey || f.NoAudit) {
								g.If(g.S(g.Id(t.Initials, "Touched"), ".Bit(", t.Title, "_", f.Title), "==", "1").Body(func() {
									g.Go(g.Id("other", f.Title), " = ", g.Id(t.Initials, f.Title))
								})
							}
						}
					})

					g.Method("Key").Receiver(t.Initials, t.Title).Returns("int64").Body(func() {
						g.Return(g.Id(t.Initials, "UUID"))
					})
				}
			}
		}
	}

	g.NewLine()
	g.Comment("^^ END OF GENERATED BY CODEGEN. ^^")

	return g.String()
}
