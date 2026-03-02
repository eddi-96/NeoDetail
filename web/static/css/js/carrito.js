const USER_ID = 1;

async function cargarCarrito(){
  const res = await fetch(`/api/carrito?usuario_id=${USER_ID}`);
  const json = await res.json();
  const data = json.data;

  const list = document.getElementById("carritoList");
  if(!data || !data.items || data.items.length === 0){
    list.innerHTML = "<p class='hint'>Tu carrito está vacío.</p>";
    document.getElementById("total").textContent = "0.00";
    return;
  }

  let total = 0;
  list.innerHTML = "";

  data.items.forEach(it=>{
    const sub = Number(it.cantidad) * Number(it.precio_unitario);
    total += sub;

    list.innerHTML += `
      <div class="card">
        <div class="name">${it.nombre || ("Producto #" + it.producto_id)}</div>
        <div class="row">
          <div>Cant: <b>${it.cantidad}</b></div>
          <div>$${Number(it.precio_unitario).toFixed(2)}</div>
        </div>
        <div class="row">
          <div class="hint">Subtotal</div>
          <div class="price">$${sub.toFixed(2)}</div>
        </div>
      </div>
    `;
  });

  document.getElementById("total").textContent = total.toFixed(2);
}

document.getElementById("btnFinalizar")?.addEventListener("click", async () => {
  const usuario_id = 1; // por ahora fijo para tarea (luego lo tomas del login)
  const res = await fetch("/api/checkout", {
    method: "POST",
    headers: {"Content-Type":"application/json"},
    body: JSON.stringify({ usuario_id })
  });
  const data = await res.json();

  const msg = document.getElementById("msgCompra");
  if (data.success) {
    msg.innerText = "✅ Gracias por tu compra. Pedido #" + data.pedido_id;
  } else {
    msg.innerText = "❌ " + (data.error || "No se pudo finalizar");
  }
});

cargarCarrito();
