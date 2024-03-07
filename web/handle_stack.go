package web

import (
	"net/http"

	"github.com/pushmarket/render"
)

type stackFetchResponse struct {
	Stack []int `json:"stack"`
}

type stackPushRequest struct {
	Number int `json:"number" form:"number"`
}

type stackPopResponse struct {
	PoppedNumber int `json:"popped_number"`
}

// handleStackFetch prints the stack of numbers
//
//	@ID			stackFetch
//	@Tags		stack
//	@Summary	Prints all numbers from stack
//	@Accept		json,x-www-form-urlencoded
//	@Produce	json
//	@Success	200		"Succesfully printed stack"
//	@Failure	400		{object}	errorResponse	"Bad request."
//	@Failure	default	{object}	errorResponse
//	@Router		/stack [get]
func (a *API) handleStackFetch(w http.ResponseWriter, r *http.Request) {
	stack := a.Config.Stack.Fetch()
	render.Status(r, http.StatusOK)
	render.JSON(w, r, stackFetchResponse{Stack: stack})
}

// handleStackAdd adds a number to the stack
//
//	@ID			stackAdd
//	@Tags		stack
//	@Summary	Adds a number to the stack
//	@Param		body	body	stackPushRequest	true	"Number to add to the stack"
//	@Accept		json,x-www-form-urlencoded
//	@Produce	json
//	@Success	200		"Succesfully added number to stack"
//	@Failure	default	"Internal server error"
//	@Router		/stack [post]
func (a *API) handleStackAdd(w http.ResponseWriter, r *http.Request) {
	req := stackPushRequest{}
	if err := render.Decode(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.Config.Stack.Push(req.Number)
	render.Status(r, http.StatusOK)
	render.NoContent(w, r)
}

// handleStackPop pops a number from the stack
//
//	@ID			stackPop
//	@Tags		stack
//	@Summary	Pops a number from the stack
//	@Accept		json,x-www-form-urlencoded
//	@Produce	json
//	@Success	200		{object}	stackPopResponse
//	@Failure	400		{object}	errorResponse	"Bad request."
//	@Failure	default	{object}	errorResponse
//	@Router		/stack [delete]
func (a *API) handleStackPop(w http.ResponseWriter, r *http.Request) {
	n := a.Config.Stack.Pop()
	render.Status(r, http.StatusOK)
	render.JSON(w, r, stackPopResponse{PoppedNumber: n})
}
