const USER_ID = 1; // por ahora fijo para demo (luego login real)

let productosCache = [];

function setCartCount(n){ document.getElementById("cartCount").textContent = String(n); }

async function cargarProductos(){
  const res = await fetch("/api/productos");
  const json = await res.json();
  productosCache = json.data || [];
  render(productosCache);
  await cargarContadorCarrito();
}

function render(lista){
  const grid = document.getElementById("grid");
  grid.innerHTML = "";
  lista.forEach(p=>{
    grid.innerHTML += `
      <div class="card">
        <div class="name">${p.nombre}</div>
        <div class="desc">${p.descripcion || ""}</div>
        <div class="row">
          <div class="price">$${Number(p.precio).toFixed(2)}</div>
          <button class="btn" onclick="agregar(${p.id})">Agregar</button>
        </div>
        <div class="hint">Stock: ${p.stock}</div>
      </div>
    `;
  });
}

async function agregar(productoId){
  const res = await fetch("/api/carrito/items", {
    method:"POST",
    headers:{ "Content-Type":"application/json" },
    body: JSON.stringify({ usuario_id: USER_ID, producto_id: productoId, cantidad: 1 })
  });

  const json = await res.json().catch(()=>({}));
  if(!res.ok){
    alert(json.error || "No se pudo agregar");
    return;
  }
  await cargarContadorCarrito();
}

async function cargarContadorCarrito(){
  const res = await fetch(`/api/carrito?usuario_id=${USER_ID}`);
  const json = await res.json();
  const items = (json.data && json.data.items) ? json.data.items : [];
  const count = items.reduce((acc,it)=>acc+(it.cantidad||0),0);
  setCartCount(count);
}

document.getElementById("buscador").addEventListener("input", (e)=>{
  const q = e.target.value.toLowerCase().trim();
  const filtrados = productosCache.filter(p =>
    (p.nombre||"").toLowerCase().includes(q) ||
    (p.descripcion||"").toLowerCase().includes(q)
  );
  render(filtrados);
});

cargarProductos();
