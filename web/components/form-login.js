import Router from "next/router";
import React from "react";
import { mutate } from "swr";

import { AuthAPI } from "../lib/api/auth";
import { UserAPI } from "../lib/api/user";

export default function FormLogin() {
    const [isLoading, setLoading] = React.useState(false);
    const [errors, setErrors] = React.useState([]);
    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");

    const handleEmailChange = React.useCallback(
        (e) => setEmail(e.target.value),
        []
    );
    const handlePasswordChange = React.useCallback(
        (e) => setPassword(e.target.value),
        []
    );

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        console.log("email: " + email);
        console.log("password: " + password);

        try {
            const { data, status } = await AuthAPI.Login(email, password);
            if (status !== 200) {
                setErrors(data.errors);
            }
            
            if (status === 200) {
                const { data, status } = await UserAPI.Me();
                if (status !== 200) {
                    setErrors(data.errors);
                }
                if (data) {
                    console.log(data);
                    window.localStorage.setItem("user", JSON.stringify(data.data));
                    mutate("user", data?.data);
                    Router.push("/");
                }
            }
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            Email:{" "}
            <input
                name="email"
                type="email"
                placeholder="Email"
                value={email}
                onChange={handleEmailChange}
            />
            <br />
            Password:{" "}
            <input
                name="password"
                type="password"
                placeholder="Password"
                value={password}
                onChange={handlePasswordChange}
            />
            <br />
            <button type="submit" disabled={isLoading}>
                Login
            </button>
            <br />
            {errors ? errors : ""}
        </form>
    );
}
