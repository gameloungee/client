const url = "localhost:8080/api/v1/account/create"

function CreateAccount() {
    let data = {
        email: document.getElementById("email").value,
        password: document.getElementById("password").value
    }

    fetch(url, {
        method: "POST",
        body: JSON.stringify(data),
        headers:
        {
            "Content-Type": "application/json"
        }
    }
    ).then(resp => response.json())
    .then(console.log(resp))
}