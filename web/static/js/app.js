(function () {
  "use strict";

  var el = document.getElementById("status-value");

  fetch("/api/health")
    .then(function (res) {
      if (!res.ok) {
        throw new Error("HTTP " + res.status);
      }
      return res.json();
    })
    .then(function (data) {
      el.textContent = data.status;
      el.className = "value ok";
    })
    .catch(function (err) {
      el.textContent = "error: " + err.message;
      el.className = "value error";
    });
})();
