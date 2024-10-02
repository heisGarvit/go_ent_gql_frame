package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
	"log/slog"
	"os"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./src/models/ent/ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.TemplateDir("./src/utils/templates"),
	}
	if err := entc.Generate("./src/models/ent/schema", &gen.Config{
		Target:  "./src/models/ent",
		Package: "project/src/models/ent",
		Features: []gen.Feature{
			gen.FeatureModifier,
		},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	g, e := entc.LoadGraph("./src/models/schema", &gen.Config{})
	if e != nil {
		log.Fatalf("loading ent graph: %v", e)
	}
	var gql string
	var fieldsEnum string
	var mutations string
	for _, node := range g.Nodes {
		fields := ""
		for _, field := range node.Fields {
			fields += gen.Funcs["camel"].(func(string) string)(field.Name) + "\n"
		}
		fieldsEnum += `
		enum ` + node.Name + `Field {
		  ` + fields + `
		}
		`
		gql += `
  	` + node.QueryName() + `(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """
    Ordering options for Members returned from the connection.
    """
    orderBy: [` + node.Name + `Order!]

    """Filtering options for Users returned from the connection."""
    where: ` + node.Name + `WhereInput
  ): ` + node.Name + `Connection!

` + node.Name + `Export(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """
    Ordering options for Members returned from the connection.
    """
    orderBy: [` + node.Name + `Order!]

    """Filtering options for Users returned from the connection."""
    where: ` + node.Name + `WhereInput

	fileType: String!,
	excelUrl: String,
  ): String
`
		mutations += `
	` + node.CreateName() + `(input: Create` + node.Name + `Input!): ` + node.Name + `!
	` + node.CreateBulkName() + `(inputs: [Create` + node.Name + `Input]!): [` + node.Name + `!]
	` + node.UpdateName() + `(id: ID!, input: Update` + node.Name + `Input!): ` + node.Name + `!
	` + node.DeleteName() + `(id: ID!): ` + node.Name + `!
`
		slog.Info(node.Name)
	}

	slog.Info("Writing to file /src/models/ent/resolvers/r.graphql")
	scalar := `
			scalar DateOnly
			scalar TimeOnly
			scalar NullString
			scalar JSON
	`
	err = os.WriteFile("./src/models/ent/resolvers/r.graphql", []byte(scalar+`extend type Query {`+gql+`}`+` type Mutation {`+mutations+`}`+fieldsEnum), 0644)
	if err != nil {
		slog.Error("Error writing to file /src/models/ent/resolvers/r.graphql", "error", err)
		os.Exit(1)
	}

}
