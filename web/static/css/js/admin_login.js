document.addEventListener("DOMContentLoaded", () => {
  const btn = document.getElementById("loginBtn");

  btn.addEventListener("click", async () => {
    const usuario = document.getElementById("adminUser").value;
    const password = document.getElementById("adminPass").value;

    const body = new URLSearchParams();
    body.append("usuario", usuario);
    body.append("password", password);

    const res = await fetch("/admin/login", {
      method: "POST",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: body.toString(),
      redirect: "follow",
    });

    // Si el backend responde con redirect, el fetch NO cambia de página solo.
    // Por eso hacemos redirect manual:
    if (res.redirected) {
      window.location.href = res.url;
      return;
    }

    if (res.ok) {
      // por si devuelves 200
      window.location.href = "/admin/productos";
    } else {
      alert("Credenciales inválidas");
    }
  });
});