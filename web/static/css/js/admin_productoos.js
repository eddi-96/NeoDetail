document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("formProducto");
  const msg = document.getElementById("msg");
  const tabla = document.getElementById("tablaProductos");

  async function cargarProductos() {
    const res = await fetch("/api/productos");
    const data = await res.json();

    if (!data.success) {
      tabla.innerHTML = `<tr><td colspan="6">Error cargando productos</td></tr>`;
      return;
    }

    tabla.innerHTML = data.data.map(p => `
      <tr>
        <td>${p.id}</td>
        <td>${p.nombre}</td>
        <td>${p.descripcion || ""}</td>
        <td>${p.precio}</td>
        <td>${p.stock}</td>
        <td>${p.imagen_url || ""}</td>
      </tr>
    `).join("");
  }

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const producto = {
      nombre: document.getElementById("nombre").value.trim(),
      descripcion: document.getElementById("descripcion").value.trim(),
      precio: parseFloat(document.getElementById("precio").value),
      stock: parseInt(document.getElementById("stock").value || "0", 10),
      imagen_url: document.getElementById("imagen_url").value.trim()
    };

    const res = await fetch("/api/productos", {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify(producto)
    });

    const data = await res.json();

    if (!data.success) {
      msg.textContent = "❌ " + (data.error || "No se pudo crear");
      return;
    }

    msg.textContent = "✅ Producto creado ID: " + data.id;
    form.reset();
    await cargarProductos();
  });

  cargarProductos();
});