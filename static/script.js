document.addEventListener("DOMContentLoaded", cargarAnimales);
const form = document.getElementById("animalForm");

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const animal = {
    nombre: form.nombre.value,
    especie: form.especie.value,
    edad: form.edad.value,
    dueno: form.dueno.value,
    telefono: form.telefono.value,
    direccion: form.direccion.value,
    barrio: form.barrio.value
  };
  const res = await fetch("/api/animales", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(animal)
  });
  const nuevo = await res.json();
  agregarAnimalTabla(nuevo);
  form.reset();
});

async function cargarAnimales() {
  const res = await fetch("/api/animales");
  const data = await res.json();
  data.forEach(agregarAnimalTabla);
}

function agregarAnimalTabla(animal) {
  const tbody = document.querySelector("#tablaAnimales tbody");
  const row = document.createElement("tr");
  row.setAttribute("id", `fila-${animal.id}`);
  row.innerHTML = `
    <td>${animal.nombre}</td>
    <td>${animal.especie}</td>
    <td>${animal.edad}</td>
    <td>${animal.dueno}</td>
    <td>${animal.telefono}</td>
    <td>${animal.barrio}</td>
    <td>
      <button class="btn btn-warning btn-sm" onclick="editar(${animal.id})">Editar</button>
      <button class="btn btn-danger btn-sm" onclick="eliminar(${animal.id})">Eliminar</button>
    </td>
  `;
  tbody.appendChild(row);
}

async function eliminar(id) {
  await fetch(`/api/animales?id=${id}`, { method: "DELETE" });
  document.getElementById(`fila-${id}`).remove();
}

async function editar(id) {
  const res = await fetch("/api/animales");
  const animales = await res.json();
  const a = animales.find(animal => animal.id === id);
  document.getElementById("editModalId").value = a.id;
  document.getElementById("editNombre").value = a.nombre;
  document.getElementById("editEspecie").value = a.especie;
  document.getElementById("editEdad").value = a.edad;
  document.getElementById("editDueno").value = a.dueno;
  document.getElementById("editTelefono").value = a.telefono;
  document.getElementById("editDireccion").value = a.direccion;
  document.getElementById("editBarrio").value = a.barrio;
  new bootstrap.Modal(document.getElementById("editarModal")).show();
}

async function guardarCambios() {
  const animal = {
    id: parseInt(document.getElementById("editModalId").value),
    nombre: document.getElementById("editNombre").value,
    especie: document.getElementById("editEspecie").value,
    edad: document.getElementById("editEdad").value,
    dueno: document.getElementById("editDueno").value,
    telefono: document.getElementById("editTelefono").value,
    direccion: document.getElementById("editDireccion").value,
    barrio: document.getElementById("editBarrio").value
  };
  const res = await fetch("/api/animales", {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(animal)
  });
  const actualizado = await res.json();
  const fila = document.getElementById(`fila-${actualizado.id}`);
  fila.innerHTML = `
    <td>${actualizado.nombre}</td>
    <td>${actualizado.especie}</td>
    <td>${actualizado.edad}</td>
    <td>${actualizado.dueno}</td>
    <td>${actualizado.telefono}</td>
    <td>${actualizado.barrio}</td>
    <td>
      <button class="btn btn-warning btn-sm" onclick="editar(${actualizado.id})">Editar</button>
      <button class="btn btn-danger btn-sm" onclick="eliminar(${actualizado.id})">Eliminar</button>
    </td>
  `;
  bootstrap.Modal.getInstance(document.getElementById("editarModal")).hide();
}
