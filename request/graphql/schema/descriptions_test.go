// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package schema

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sourcenetwork/defradb/client"
)

func TestSingleSimpleType(t *testing.T) {
	cases := []descriptionTestCase{
		{
			description: "Single simple type",
			sdl: `
			type user {
				name: String
				age: Int
				verified: Boolean
			}
			`,
			targetDescs: []client.CollectionDescription{
				{
					Name: "user",
					Schema: client.SchemaDescription{
						Name: "user",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "age",
								Kind: client.FieldKind_INT,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "verified",
								Kind: client.FieldKind_BOOL,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
			},
		},
		{
			description: "Multiple simple types",
			sdl: `
			type user {
				name: String
				age: Int
				verified: Boolean
			}

			type author {
				name: String
				publisher: String
				rating: Float
			}
			`,
			targetDescs: []client.CollectionDescription{
				{
					Name: "user",
					Schema: client.SchemaDescription{
						Name: "user",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "age",
								Kind: client.FieldKind_INT,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "verified",
								Kind: client.FieldKind_BOOL,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
				{
					Name: "author",
					Schema: client.SchemaDescription{
						Name: "author",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "publisher",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "rating",
								Kind: client.FieldKind_FLOAT,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
			},
		},
		{
			description: "Multiple types with relations (one-to-one)",
			sdl: `
			type book {
				name: String
				rating: Float
				author: author
			}

			type author {
				name: String
				age: Int
				published: book
			}
			`,
			targetDescs: []client.CollectionDescription{
				{
					Name: "book",
					Schema: client.SchemaDescription{
						Name: "book",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name:         "author",
								RelationName: "author_book",
								Kind:         client.FieldKind_FOREIGN_OBJECT,
								Typ:          client.NONE_CRDT,
								Schema:       "author",
								RelationType: client.Relation_Type_ONE | client.Relation_Type_ONEONE,
							},
							{
								Name:         "author_id",
								Kind:         client.FieldKind_DocKey,
								Typ:          client.LWW_REGISTER,
								RelationType: client.Relation_Type_INTERNAL_ID,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "rating",
								Kind: client.FieldKind_FLOAT,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
				{
					Name: "author",
					Schema: client.SchemaDescription{
						Name: "author",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "age",
								Kind: client.FieldKind_INT,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name:         "published",
								RelationName: "author_book",
								Kind:         client.FieldKind_FOREIGN_OBJECT,
								Typ:          client.NONE_CRDT,
								Schema:       "book",
								RelationType: client.Relation_Type_ONE | client.Relation_Type_ONEONE | client.Relation_Type_Primary,
							},
							{
								Name:         "published_id",
								Kind:         client.FieldKind_DocKey,
								Typ:          client.LWW_REGISTER,
								RelationType: client.Relation_Type_INTERNAL_ID,
							},
						},
					},
				},
			},
		},
		{
			description: "Multiple simple types",
			sdl: `
			type user {
				name: String
				age: Int
				verified: Boolean
			}

			type author {
				name: String
				publisher: String
				rating: Float
			}
			`,
			targetDescs: []client.CollectionDescription{
				{
					Name: "user",
					Schema: client.SchemaDescription{
						Name: "user",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "age",
								Kind: client.FieldKind_INT,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "verified",
								Kind: client.FieldKind_BOOL,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
				{
					Name: "author",
					Schema: client.SchemaDescription{
						Name: "author",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "publisher",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "rating",
								Kind: client.FieldKind_FLOAT,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
			},
		},
		{
			description: "Multiple types with relations (one-to-one)",
			sdl: `
			type book {
				name: String
				rating: Float
				author: author @relation(name:"book_authors")
			}

			type author {
				name: String
				age: Int
				published: book @relation(name:"book_authors")
			}
			`,
			targetDescs: []client.CollectionDescription{
				{
					Name: "book",
					Schema: client.SchemaDescription{
						Name: "book",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name:         "author",
								RelationName: "book_authors",
								Kind:         client.FieldKind_FOREIGN_OBJECT,
								Typ:          client.NONE_CRDT,
								Schema:       "author",
								RelationType: client.Relation_Type_ONE | client.Relation_Type_ONEONE,
							},
							{
								Name:         "author_id",
								Kind:         client.FieldKind_DocKey,
								Typ:          client.LWW_REGISTER,
								RelationType: client.Relation_Type_INTERNAL_ID,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "rating",
								Kind: client.FieldKind_FLOAT,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
				{
					Name: "author",
					Schema: client.SchemaDescription{
						Name: "author",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "age",
								Kind: client.FieldKind_INT,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name:         "published",
								RelationName: "book_authors",
								Kind:         client.FieldKind_FOREIGN_OBJECT,
								Typ:          client.NONE_CRDT,
								Schema:       "book",
								RelationType: client.Relation_Type_ONE | client.Relation_Type_ONEONE | client.Relation_Type_Primary,
							},
							{
								Name:         "published_id",
								Kind:         client.FieldKind_DocKey,
								Typ:          client.LWW_REGISTER,
								RelationType: client.Relation_Type_INTERNAL_ID,
							},
						},
					},
				},
			},
		},
		{
			description: "Multiple types with relations (one-to-one) with directive",
			sdl: `
			type book {
				name: String
				rating: Float
				author: author @primary
			}

			type author {
				name: String
				age: Int
				published: book
			}
			`,
			targetDescs: []client.CollectionDescription{
				{
					Name: "book",
					Schema: client.SchemaDescription{
						Name: "book",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name:         "author",
								RelationName: "author_book",
								Kind:         client.FieldKind_FOREIGN_OBJECT,
								Typ:          client.NONE_CRDT,
								Schema:       "author",
								RelationType: client.Relation_Type_ONE | client.Relation_Type_ONEONE | client.Relation_Type_Primary,
							},
							{
								Name:         "author_id",
								Kind:         client.FieldKind_DocKey,
								Typ:          client.LWW_REGISTER,
								RelationType: client.Relation_Type_INTERNAL_ID,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "rating",
								Kind: client.FieldKind_FLOAT,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
				{
					Name: "author",
					Schema: client.SchemaDescription{
						Name: "author",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "age",
								Kind: client.FieldKind_INT,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name:         "published",
								RelationName: "author_book",
								Kind:         client.FieldKind_FOREIGN_OBJECT,
								Typ:          client.NONE_CRDT,
								Schema:       "book",
								RelationType: client.Relation_Type_ONE | client.Relation_Type_ONEONE,
							},
							{
								Name:         "published_id",
								Kind:         client.FieldKind_DocKey,
								Typ:          client.LWW_REGISTER,
								RelationType: client.Relation_Type_INTERNAL_ID,
							},
						},
					},
				},
			},
		},
		{
			description: "Multiple types with relations (one-to-many)",
			sdl: `
			type book {
				name: String
				rating: Float
				author: author
			}

			type author {
				name: String
				age: Int
				published: [book]
			}
			`,
			targetDescs: []client.CollectionDescription{
				{
					Name: "book",
					Schema: client.SchemaDescription{
						Name: "book",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name:         "author",
								RelationName: "author_book",
								Kind:         client.FieldKind_FOREIGN_OBJECT,
								Typ:          client.NONE_CRDT,
								Schema:       "author",
								RelationType: client.Relation_Type_ONE | client.Relation_Type_ONEMANY | client.Relation_Type_Primary,
							},
							{
								Name:         "author_id",
								Kind:         client.FieldKind_DocKey,
								Typ:          client.LWW_REGISTER,
								RelationType: client.Relation_Type_INTERNAL_ID,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "rating",
								Kind: client.FieldKind_FLOAT,
								Typ:  client.LWW_REGISTER,
							},
						},
					},
				},
				{
					Name: "author",
					Schema: client.SchemaDescription{
						Name: "author",
						Fields: []client.FieldDescription{
							{
								Name: "_key",
								Kind: client.FieldKind_DocKey,
								Typ:  client.NONE_CRDT,
							},
							{
								Name: "age",
								Kind: client.FieldKind_INT,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name: "name",
								Kind: client.FieldKind_STRING,
								Typ:  client.LWW_REGISTER,
							},
							{
								Name:         "published",
								RelationName: "author_book",
								Kind:         client.FieldKind_FOREIGN_OBJECT_ARRAY,
								Typ:          client.NONE_CRDT,
								Schema:       "book",
								RelationType: client.Relation_Type_MANY | client.Relation_Type_ONEMANY,
							},
						},
					},
				},
			},
		},
	}

	for _, test := range cases {
		runCreateDescriptionTest(t, test)
	}
}

func runCreateDescriptionTest(t *testing.T, testcase descriptionTestCase) {
	ctx := context.Background()

	descs, err := FromString(ctx, testcase.sdl)
	assert.NoError(t, err, testcase.description)
	assert.Equal(t, len(descs), len(testcase.targetDescs), testcase.description)

	for i, d := range descs {
		assert.Equal(t, testcase.targetDescs[i], d, testcase.description)
	}
}

type descriptionTestCase struct {
	description string
	sdl         string
	targetDescs []client.CollectionDescription
}
