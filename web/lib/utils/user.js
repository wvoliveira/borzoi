let User = {
    Name: "",
}

export function GetUser() {
    let data = localStorage.getItem("user")
    if (!data === undefined) {
        return JSON.parse(data)
    }
}

export function SetUser(user) {
    let data = JSON.stringify(user)
    localStorage.setItem("user", data)
}
