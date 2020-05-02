package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thayen/sample-rest/domain/model"
	"github.com/thayen/sample-rest/usecase"
	"log"
	"net/http"
	"strconv"
)

type HttpContactServer struct {
	*mux.Router
	contactUsecase usecase.ContactUsecase
}

func NewHttpContact(contactUsecase usecase.ContactUsecase) HttpContactServer {

	router := mux.NewRouter()
	contact := HttpContactServer{contactUsecase: contactUsecase}
	contact.Router = router

	router.HandleFunc("/contacts", contact.findAll).Methods("GET")
	router.HandleFunc("/contacts", contact.createContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", contact.getContact).Methods("GET")
	router.HandleFunc("/contacts/{id}", contact.updateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", contact.deleteContact).Methods("DELETE")
	return contact
}

func (h HttpContactServer) updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact ContactUpdate
	params := mux.Vars(r)
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		w.WriteHeader(400)
		_, _ = w.Write([]byte("invalid JSON " + err.Error()))
		return
	}
	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(400)
		_, _ = w.Write([]byte("invalid id"))
		return
	}
	contact.Id = int(id)
	if err := h.contactUsecase.Update(contact); err != nil {
		setError(w, err)
		return
	}
	json.NewEncoder(w).Encode(&contact)
}

func (h HttpContactServer) getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}
	contact, err := h.contactUsecase.Get(int(id))
	if err != nil {
		setError(w, err)
		return
	}
	json.NewEncoder(w).Encode(toContactResponse(contact))
}

func (h HttpContactServer) createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact ContactCreation
	_ = json.NewDecoder(r.Body).Decode(&contact)
	id, err := h.contactUsecase.Create(contact)
	if err != nil {
		setError(w, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("contacts/%d", id))
	w.WriteHeader(201)
}

func (h HttpContactServer) findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact ContactCreation
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts, err := h.contactUsecase.FindAll()
	if err != nil {
		setError(w, err)
		return
	}
	all := make([]ContactResponse, 0, len(contacts))
	for _, c := range contacts {
		all = append(all, toContactResponse(c))
	}
	_ = json.NewEncoder(w).Encode(all)

}

func (h HttpContactServer) deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("invalid id"))
		return
	}
	if err := h.contactUsecase.Delete(int(id)); err != nil {
		setError(w, err)
		return
	}
}

func setError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case model.EmailAlreadyExists:
		w.WriteHeader(400)
		w.Write([]byte("Email already exists"))
	case model.NotFound:
		w.WriteHeader(404)
	default:
		w.WriteHeader(500)
	}
	log.Println(err)
}
