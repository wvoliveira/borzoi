export function GetAuth() {
    return localStorage.getItem("auth");
}

export function SetAuth() {
    localStorage.setItem("auth", "true");
}

export async function IsAuthenticated() {
    let auth = getFromStorage("auth");
    return auth !== undefined;
}

export async function AuthPassword(email, password) {
    let data = {
        "email": email,
        "password": password,
    }
    return await fetch("/api/v1/auth/password/login", {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
        .then(response => {
            if (response.ok) {
                console.log("successful: " + response.json());
                return response;
            } else {
                console.log('Network response was not ok. Status: ' + response.status);
                response.status !== 500 ? console.log(response.json()) : null;
               return response;
            }
        })
        .catch(response => {
            console.log(response);
            return response;
        })
}
