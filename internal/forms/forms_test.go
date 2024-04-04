package forms

import (
	"net/http/httptest"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	//como é um reciver é preciso criar um Form antes de testar
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid form, should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	//como é um reciver é preciso criar um Form antes de testar
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form show valid when required field is empty")
	}
	//TODO: terminar o test aula 79
}
