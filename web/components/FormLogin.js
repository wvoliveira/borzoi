import Router from "next/router";
import React from "react";

import {AuthAPI} from "../lib/api/auth";
import useUser from "../lib/utils/useUser";
import {Box, TextField} from "@mui/material";

export default function FormLogin() {
    const {mutateUser} = useUser({
        redirectTo: "/profile",
        redirectIfFound: true,
    });

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
            const {data, status} = await AuthAPI.Login(email, password);
            if (status !== 200 && status !== 500) {
                setErrors(data.message);
                console.log(data.message);
            }

            if (status === 200) {
                mutateUser(data)
                Router.push("/");
            }
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <TextField
                id="email"
                label="Email"
                variant="outlined"
                name="email"
                type="email"
                placeholder="Email"
                value={email}
                onChange={handleEmailChange}
            />
            <br/>
            <TextField
                id="password"
                label="Password"
                variant="outlined"
                name="password"
                type="password"
                placeholder="Password"
                value={password}
                onChange={handlePasswordChange}
            />
            <br/>
            <button type="submit" disabled={isLoading}>
                Login
            </button>
            <br/>
            {errors ? errors : ""}
        </form>
    );
}
