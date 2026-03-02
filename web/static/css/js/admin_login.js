document.getElementById("loginBtn").addEventListener("click", ()=>{
  const u = document.getElementById("adminUser").value.trim();
  const p = document.getElementById("adminPass").value.trim();

  if(!u || !p){ alert("Completa credenciales"); return; }

  // guardamos en localStorage para usar en panel admin
  localStorage.setItem("admin_user", u);
  localStorage.setItem("admin_pass", p);

  window.location.href = "/admin/productos";
});