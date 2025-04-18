package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAPIAnimales(t *testing.T) {
	server := http.HandlerFunc(handleAnimales)

	// 1. GET inicial debe devolver lista vacía
	t.Log("1. GET inicial - Debe devolver lista vacía")
	req := httptest.NewRequest(http.MethodGet, "/api/animales", nil)
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Esperado 200, obtuve %d", w.Code)
	}
	var lista []Animal
	json.Unmarshal(w.Body.Bytes(), &lista)
	if len(lista) != 0 {
		t.Errorf("Esperado lista vacía, obtuve %v", lista)
	} else {
		t.Log("GET inicial exitoso, lista vacía confirmada")
	}

	// Datos base para las siguientes pruebas
	baseAnimal := Animal{
		Nombre:    "Max",
		Especie:   "Perro",
		Edad:      "3",
		Dueno:     "Carlos",
		Telefono:  "3210001234",
		Direccion: "Cra 10 #20-30",
		Barrio:    "Centro",
	}

	// 2. POST - Registrar un animal
	t.Logf("2. POST - Registrando animal: %+v", baseAnimal)
	body, _ := json.Marshal(baseAnimal)
	req = httptest.NewRequest(http.MethodPost, "/api/animales", bytes.NewReader(body))
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("POST falló con código %d", w.Code)
	}
	var creado Animal
	json.Unmarshal(w.Body.Bytes(), &creado)
	if creado.ID != 1 || creado.Nombre != "Max" {
		t.Errorf("Animal creado no válido: %v", creado)
	} else {
		t.Logf("Animal creado con éxito: %+v", creado)
	}

	// 3. GET - Debe contener al animal creado
	t.Log("3. GET - Verificando existencia del animal creado")
	req = httptest.NewRequest(http.MethodGet, "/api/animales", nil)
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &lista)
	if len(lista) != 1 || lista[0].Nombre != "Max" {
		t.Errorf("Lista incorrecta: %v", lista)
	} else {
		t.Logf("GET exitoso, lista contiene: %+v", lista[0])
	}

	// 4-10. POST adicionales
	t.Log("4-10. POST - Agregando más animales")
	for i := 2; i <= 8; i++ {
		baseAnimal.Nombre = "Animal" + strconv.Itoa(i)
		body, _ := json.Marshal(baseAnimal)
		req = httptest.NewRequest(http.MethodPost, "/api/animales", bytes.NewReader(body))
		w = httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("POST #%d falló", i)
		} else {
			t.Logf("POST #%d exitoso: %s", i, baseAnimal.Nombre)
		}
	}

	// 11. PUT - Actualizar animal con ID 1
	t.Log("11. PUT - Actualizando animal con ID 1")
	actualizado := creado
	actualizado.Nombre = "Maximus"
	body, _ = json.Marshal(actualizado)
	req = httptest.NewRequest(http.MethodPut, "/api/animales", bytes.NewReader(body))
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	var modificado Animal
	json.Unmarshal(w.Body.Bytes(), &modificado)
	if modificado.Nombre != "Maximus" {
		t.Errorf("PUT falló: nombre esperado 'Maximus', obtuve %s", modificado.Nombre)
	} else {
		t.Logf("PUT exitoso: %+v", modificado)
	}

	// 12. GET - Verificar que actualización fue exitosa
	t.Log("12. GET - Verificando actualización")
	req = httptest.NewRequest(http.MethodGet, "/api/animales", nil)
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &lista)
	if lista[0].Nombre != "Maximus" {
		t.Errorf("Actualización no persistida: %v", lista[0])
	} else {
		t.Log("Actualización confirmada: Nombre cambiado a Maximus")
	}

	// 13. DELETE - Eliminar animal con ID 1
	t.Log("13. DELETE - Eliminando animal con ID 1")
	req = httptest.NewRequest(http.MethodDelete, "/api/animales?id=1", nil)
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("DELETE falló con código %d", w.Code)
	}
	respBody, _ := ioutil.ReadAll(w.Body)
	if !bytes.Contains(respBody, []byte("deleted")) {
		t.Errorf("DELETE no retornó estado esperado")
	} else {
		t.Log("Animal con ID 1 eliminado correctamente")
	}

	// 14. GET - Verificar que ya no existe el animal con ID 1
	t.Log("14. GET - Verificando que animal con ID 1 no existe")
	req = httptest.NewRequest(http.MethodGet, "/api/animales", nil)
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	json.Unmarshal(w.Body.Bytes(), &lista)
	encontrado := false
	for _, a := range lista {
		if a.ID == 1 {
			encontrado = true
			break
		}
	}
	if encontrado {
		t.Errorf("Animal con ID 1 no fue eliminado")
	} else {
		t.Log("Confirmado: Animal con ID 1 ya no está en la lista")
	}

	// 15. DELETE - Intentar eliminar un ID inexistente
	t.Log("15. DELETE - Intentando eliminar ID inexistente (999)")
	req = httptest.NewRequest(http.MethodDelete, "/api/animales?id=999", nil)
	w = httptest.NewRecorder()
	server.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("DELETE con ID inexistente falló")
	} else {
		t.Log("DELETE con ID inexistente manejado correctamente")
	}
}
