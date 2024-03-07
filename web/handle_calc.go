package web

import (
	"net/http"

	"github.com/pushmarket/render"
)

type calcResponse struct {
	Result int `json:"result" validate:"required"`
}

// handleCalcAdd adds top two numbers from the stack of calculators memory
//
//	@ID			calcAdd
//	@Tags		calc
//	@Summary	Adds top two numbers from the stack of calculators memory
//	@Accept		json,x-www-form-urlencoded
//	@Produce	json
//	@Success	200		{object}	calcResponse
//	@Failure	400		{object}	errorResponse	"Bad request."
//	@Failure	default	{object}	errorResponse
//	@Router		/calc/add [post]
func (a *API) handleCalcAdd(w http.ResponseWriter, r *http.Request) {
	res, err := a.Config.Calculator.Add()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, calcResponse{Result: res})
}

// handleCalcDiv divides the top two numbers from the stack of calculators memory
//
//	@ID			calcDiv
//	@Tags		calc
//	@Summary	Divides top two numbers from the stack of calculators memory
//	@Accept		json,x-www-form-urlencoded
//	@Produce	json
//	@Success	200		{object}	calcResponse
//	@Failure	400		{object}	errorResponse	"Bad request."
//	@Failure	default	{object}	errorResponse
//	@Router		/calc/div [post]
func (a *API) handleCalcDiv(w http.ResponseWriter, r *http.Request) {
	res, err := a.Config.Calculator.Div()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, calcResponse{Result: res})
}

// handleCalcMultip multiplies the top two numbers from the stack of calculators memory
//
//	@ID			calcMultip
//	@Tags		calc
//	@Summary	Multiplies top two numbers from the stack of calculators memory
//	@Accept		json,x-www-form-urlencoded
//	@Produce	json
//	@Success	200		{object}	calcResponse
//	@Failure	400		{object}	errorResponse	"Bad request."
//	@Failure	default	{object}	errorResponse
//	@Router		/calc/multip [post]
func (a *API) handleCalcMultip(w http.ResponseWriter, r *http.Request) {
	res, err := a.Config.Calculator.Mul()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, calcResponse{Result: res})
}

// handleCalcSubs multiplies the top two numbers from the stack of calculators memory
//
//	@ID			calcSubs
//	@Tags		calc
//	@Summary	Subtracts top two numbers from the stack of calculators memory
//	@Accept		json,x-www-form-urlencoded
//	@Produce	json
//	@Success	200		{object}	calcResponse
//	@Failure	400		{object}	errorResponse	"Bad request."
//	@Failure	default	{object}	errorResponse
//	@Router		/calc/subs [post]
func (a *API) handleCalcSubs(w http.ResponseWriter, r *http.Request) {
	res, err := a.Config.Calculator.Sub()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, calcResponse{Result: res})
}
