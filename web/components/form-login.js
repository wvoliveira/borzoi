import * as React from "react";
import {GetUser, SetUser} from '../lib/user';
import {AuthPassword, SetAuth} from '../lib/auth';

export default function FormLogin() {
    const [loading, setLoading] = React.useState(false);
    const [error, setError] = React.useState("");

    const handleSubmit = async (event) => {
        setLoading(true);
        event.preventDefault();

        const inputs = new FormData(event.currentTarget);

        const payload = {
            "email": inputs.get('email'),
            "password": inputs.get('password'),
        }
        console.log("email: " + payload.email + " password: " + payload.password);

        AuthPassword(payload.email, payload.password)
            .then(response => {
                setLoading(false);
                if (!response.ok) {
                    if (response.status !== 500) {
                        let data = response.json();
                        setError(data.message);
                    } else {
                        setError("Unknown error.");
                    }
                } else {
                    SetAuth();
                }
            })
    };

    return (
        <form onSubmit={handleSubmit}>
            Email: <input name="email"/><br/>
            Password: <input name="password"/><br/>
            <button>Login</button><br/>
            {loading ? "Loading..." : ""}
            {error ? error : ""}
        </form>
    )
}
