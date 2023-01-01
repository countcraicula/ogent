// Code generated by ent, DO NOT EDIT.

package ogent

import (
	"context"
	"net/http"

	"ariga.io/ogent/example/pets/ent"
	"ariga.io/ogent/example/pets/ent/category"
	"ariga.io/ogent/example/pets/ent/pet"
	"ariga.io/ogent/example/pets/ent/user"
	"github.com/go-faster/jx"
)

// OgentHandler implements the ogen generated Handler interface and uses Ent as data layer.
type OgentHandler struct {
	client *ent.Client
}

// NewOgentHandler returns a new OgentHandler.
func NewOgentHandler(c *ent.Client) *OgentHandler { return &OgentHandler{c} }

// rawError renders err as json string.
func rawError(err error) jx.Raw {
	var e jx.Encoder
	e.Str(err.Error())
	return e.Bytes()
}

// CreateCategory handles POST /categories requests.
func (h *OgentHandler) CreateCategory(ctx context.Context, req *CreateCategoryReq) (CreateCategoryRes, error) {
	b := h.client.Category.Create()
	// Add all fields.
	b.SetName(req.Name)
	// Add all edges.
	b.AddPetIDs(req.Pets...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Category.Query().Where(category.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewCategoryCreate(e), nil
}

// ReadCategory handles GET /categories/{id} requests.
func (h *OgentHandler) ReadCategory(ctx context.Context, params ReadCategoryParams) (ReadCategoryRes, error) {
	q := h.client.Category.Query().Where(category.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewCategoryRead(e), nil
}

// UpdateCategory handles PATCH /categories/{id} requests.
func (h *OgentHandler) UpdateCategory(ctx context.Context, req *UpdateCategoryReq, params UpdateCategoryParams) (UpdateCategoryRes, error) {
	b := h.client.Category.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	// Add all edges.
	b.ClearPets().AddPetIDs(req.Pets...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Category.Query().Where(category.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewCategoryUpdate(e), nil
}

// DeleteCategory handles DELETE /categories/{id} requests.
func (h *OgentHandler) DeleteCategory(ctx context.Context, params DeleteCategoryParams) (DeleteCategoryRes, error) {
	err := h.client.Category.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteCategoryNoContent), nil

}

// ListCategory handles GET /categories requests.
func (h *OgentHandler) ListCategory(ctx context.Context, params ListCategoryParams) (ListCategoryRes, error) {
	q := h.client.Category.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewCategoryLists(es)
	return (*ListCategoryOKApplicationJSON)(&r), nil
}

// ListCategoryPets handles GET /categories/{id}/pets requests.
func (h *OgentHandler) ListCategoryPets(ctx context.Context, params ListCategoryPetsParams) (ListCategoryPetsRes, error) {
	q := h.client.Category.Query().Where(category.IDEQ(params.ID)).QueryPets()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewCategoryPetsLists(es)
	return (*ListCategoryPetsOKApplicationJSON)(&r), nil
}

// CreatePet handles POST /pets requests.
func (h *OgentHandler) CreatePet(ctx context.Context, req *CreatePetReq) (CreatePetRes, error) {
	b := h.client.Pet.Create()
	// Add all fields.
	b.SetName(req.Name)
	if v, ok := req.Weight.Get(); ok {
		b.SetWeight(v)
	}
	if v, ok := req.Birthday.Get(); ok {
		b.SetBirthday(v)
	}
	// Add all edges.
	b.AddCategoryIDs(req.Categories...)
	b.SetOwnerID(req.Owner)
	b.AddFriendIDs(req.Friends...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Pet.Query().Where(pet.ID(e.ID))
	// Eager load edges that are required on create operation.
	q.WithCategories().WithOwner()
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewPetCreate(e), nil
}

// DeletePet handles DELETE /pets/{id} requests.
func (h *OgentHandler) DeletePet(ctx context.Context, params DeletePetParams) (DeletePetRes, error) {
	err := h.client.Pet.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeletePetNoContent), nil

}

// ListPet handles GET /pets requests.
func (h *OgentHandler) ListPet(ctx context.Context, params ListPetParams) (ListPetRes, error) {
	q := h.client.Pet.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewPetLists(es)
	return (*ListPetOKApplicationJSON)(&r), nil
}

// ReadPet handles GET /pets/{id} requests.
func (h *OgentHandler) ReadPet(ctx context.Context, params ReadPetParams) (ReadPetRes, error) {
	q := h.client.Pet.Query().Where(pet.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewPetRead(e), nil
}

// UpdatePet handles PATCH /pets/{id} requests.
func (h *OgentHandler) UpdatePet(ctx context.Context, req *UpdatePetReq, params UpdatePetParams) (UpdatePetRes, error) {
	b := h.client.Pet.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	if v, ok := req.Weight.Get(); ok {
		b.SetWeight(v)
	}
	if v, ok := req.Birthday.Get(); ok {
		b.SetBirthday(v)
	}
	// Add all edges.
	b.ClearCategories().AddCategoryIDs(req.Categories...)
	if v, ok := req.Owner.Get(); ok {
		b.SetOwnerID(v)
	}
	b.ClearFriends().AddFriendIDs(req.Friends...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.Pet.Query().Where(pet.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewPetUpdate(e), nil
}

// ListPetCategories handles GET /pets/{id}/categories requests.
func (h *OgentHandler) ListPetCategories(ctx context.Context, params ListPetCategoriesParams) (ListPetCategoriesRes, error) {
	q := h.client.Pet.Query().Where(pet.IDEQ(params.ID)).QueryCategories()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewPetCategoriesLists(es)
	return (*ListPetCategoriesOKApplicationJSON)(&r), nil
}

// ReadPetOwner handles GET /pets/{id}/owner requests.
func (h *OgentHandler) ReadPetOwner(ctx context.Context, params ReadPetOwnerParams) (ReadPetOwnerRes, error) {
	q := h.client.Pet.Query().Where(pet.IDEQ(params.ID)).QueryOwner()
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewPetOwnerRead(e), nil
}

// ListPetFriends handles GET /pets/{id}/friends requests.
func (h *OgentHandler) ListPetFriends(ctx context.Context, params ListPetFriendsParams) (ListPetFriendsRes, error) {
	q := h.client.Pet.Query().Where(pet.IDEQ(params.ID)).QueryFriends()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewPetFriendsLists(es)
	return (*ListPetFriendsOKApplicationJSON)(&r), nil
}

// CreateUser handles POST /users requests.
func (h *OgentHandler) CreateUser(ctx context.Context, req *CreateUserReq) (CreateUserRes, error) {
	b := h.client.User.Create()
	// Add all fields.
	b.SetName(req.Name)
	b.SetAge(req.Age)
	// Add all edges.
	b.AddPetIDs(req.Pets...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewUserCreate(e), nil
}

// ReadUser handles GET /users/{id} requests.
func (h *OgentHandler) ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID))
	e, err := q.Only(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return NewUserRead(e), nil
}

// UpdateUser handles PATCH /users/{id} requests.
func (h *OgentHandler) UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error) {
	b := h.client.User.UpdateOneID(params.ID)
	// Add all fields.
	if v, ok := req.Name.Get(); ok {
		b.SetName(v)
	}
	if v, ok := req.Age.Get(); ok {
		b.SetAge(v)
	}
	// Add all edges.
	b.ClearPets().AddPetIDs(req.Pets...)
	// Persist to storage.
	e, err := b.Save(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	// Reload the entity to attach all eager-loaded edges.
	q := h.client.User.Query().Where(user.ID(e.ID))
	e, err = q.Only(ctx)
	if err != nil {
		// This should never happen.
		return nil, err
	}
	return NewUserUpdate(e), nil
}

// DeleteUser handles DELETE /users/{id} requests.
func (h *OgentHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error) {
	err := h.client.User.DeleteOneID(params.ID).Exec(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsConstraintError(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	return new(DeleteUserNoContent), nil

}

// ListUser handles GET /users requests.
func (h *OgentHandler) ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error) {
	q := h.client.User.Query()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)

	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewUserLists(es)
	return (*ListUserOKApplicationJSON)(&r), nil
}

// ListUserPets handles GET /users/{id}/pets requests.
func (h *OgentHandler) ListUserPets(ctx context.Context, params ListUserPetsParams) (ListUserPetsRes, error) {
	q := h.client.User.Query().Where(user.IDEQ(params.ID)).QueryPets()
	page := 1
	if v, ok := params.Page.Get(); ok {
		page = v
	}
	itemsPerPage := 30
	if v, ok := params.ItemsPerPage.Get(); ok {
		itemsPerPage = v
	}
	q.Limit(itemsPerPage).Offset((page - 1) * itemsPerPage)
	es, err := q.All(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return &R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: rawError(err),
			}, nil
		case ent.IsNotSingular(err):
			return &R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: rawError(err),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}
	r := NewUserPetsLists(es)
	return (*ListUserPetsOKApplicationJSON)(&r), nil
}
