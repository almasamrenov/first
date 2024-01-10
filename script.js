document.addEventListener("DOMContentLoaded", function () {
  const registrationForm = document.getElementById("registrationForm");

  registrationForm.addEventListener("submit", function (event) {
    event.preventDefault();

    const name = document.getElementById("input-name").value;
    const username = document.getElementById("input-username").value;
    const email = document.getElementById("input-email").value;
    const password = document.getElementById("input-password").value;

    const requestData = {
      person: {
        name,
        username,
        email,
        password,
      },
    };

    fetch("http://localhost:8080/person-endpoint", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestData),
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => {
        displayResponse(data);
      })
      .catch((error) => {
        console.error("Error:", error);
        displayError(error.message);
      });
  });

  function displayResponse(responseData) {
    const responseContainer = document.getElementById("responseContainer");
    responseContainer.innerHTML = `<div class="alert alert-${
      responseData.status === "success" ? "success" : "danger"
    }">Success</div>`;
  }

  function displayError(errorMessage) {
    const responseContainer = document.getElementById("responseContainer");
    responseContainer.innerHTML = `<div class="alert alert-danger">${errorMessage}</div>`;
  }
});
