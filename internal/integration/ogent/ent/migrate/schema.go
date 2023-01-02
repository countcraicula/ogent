// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AllTypesColumns holds the columns for the "all_types" table.
	AllTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "int", Type: field.TypeInt},
		{Name: "int8", Type: field.TypeInt8},
		{Name: "int16", Type: field.TypeInt16},
		{Name: "int32", Type: field.TypeInt32},
		{Name: "int64", Type: field.TypeInt64},
		{Name: "uint", Type: field.TypeUint},
		{Name: "uint8", Type: field.TypeUint8},
		{Name: "uint16", Type: field.TypeUint16},
		{Name: "uint32", Type: field.TypeUint32},
		{Name: "uint64", Type: field.TypeUint64},
		{Name: "float32", Type: field.TypeFloat32},
		{Name: "float64", Type: field.TypeFloat64},
		{Name: "string_type", Type: field.TypeString},
		{Name: "bool", Type: field.TypeBool},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "time", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"on", "off"}},
		{Name: "bytes", Type: field.TypeBytes},
		{Name: "nilable", Type: field.TypeString, Size: 2147483647},
	}
	// AllTypesTable holds the schema information for the "all_types" table.
	AllTypesTable = &schema.Table{
		Name:       "all_types",
		Columns:    AllTypesColumns,
		PrimaryKey: []*schema.Column{AllTypesColumns[0]},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "readonly", Type: field.TypeString, Nullable: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "weight", Type: field.TypeInt, Nullable: true},
		{Name: "birthday", Type: field.TypeTime, Nullable: true},
		{Name: "tag_id", Type: field.TypeBytes, Nullable: true},
		{Name: "height", Type: field.TypeInt, Nullable: true},
		{Name: "user_pets", Type: field.TypeInt},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_users_pets",
				Columns:    []*schema.Column{PetsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeUint},
		{Name: "height", Type: field.TypeUint, Nullable: true},
		{Name: "favorite_cat_breed", Type: field.TypeEnum, Enums: []string{"siamese", "bengal", "lion", "tiger", "leopard", "other"}},
		{Name: "favorite_dog_breed", Type: field.TypeEnum, Nullable: true, Enums: []string{"Kuro"}},
		{Name: "favorite_fish_breed", Type: field.TypeEnum, Nullable: true, Enums: []string{"gold", "koi", "shark"}},
		{Name: "user_best_friend", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_best_friend",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CategoryPetsColumns holds the columns for the "category_pets" table.
	CategoryPetsColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeInt},
		{Name: "pet_id", Type: field.TypeInt},
	}
	// CategoryPetsTable holds the schema information for the "category_pets" table.
	CategoryPetsTable = &schema.Table{
		Name:       "category_pets",
		Columns:    CategoryPetsColumns,
		PrimaryKey: []*schema.Column{CategoryPetsColumns[0], CategoryPetsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_pets_category_id",
				Columns:    []*schema.Column{CategoryPetsColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "category_pets_pet_id",
				Columns:    []*schema.Column{CategoryPetsColumns[1]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PetFriendsColumns holds the columns for the "pet_friends" table.
	PetFriendsColumns = []*schema.Column{
		{Name: "pet_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// PetFriendsTable holds the schema information for the "pet_friends" table.
	PetFriendsTable = &schema.Table{
		Name:       "pet_friends",
		Columns:    PetFriendsColumns,
		PrimaryKey: []*schema.Column{PetFriendsColumns[0], PetFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pet_friends_pet_id",
				Columns:    []*schema.Column{PetFriendsColumns[0]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "pet_friends_friend_id",
				Columns:    []*schema.Column{PetFriendsColumns[1]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AllTypesTable,
		CategoriesTable,
		PetsTable,
		UsersTable,
		CategoryPetsTable,
		PetFriendsTable,
	}
)

func init() {
	PetsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	CategoryPetsTable.ForeignKeys[0].RefTable = CategoriesTable
	CategoryPetsTable.ForeignKeys[1].RefTable = PetsTable
	PetFriendsTable.ForeignKeys[0].RefTable = PetsTable
	PetFriendsTable.ForeignKeys[1].RefTable = PetsTable
}
