import Router from "next/router";
import React from "react";

import {AuthAPI} from "../lib/api/auth";
import useAuth from "../lib/utils/useAuth";
import {Button, TextField} from "@mui/material";

export default function FormLogin() {
    const {mutateAuth} = useAuth({
        redirectTo: "/",
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

        try {
            const {data, status} = await AuthAPI.Login(email, password);
            if (status !== 200 && status !== 500) {
                setErrors(data.message);
                console.log(data.message);
            }

            if (status === 200) {
                mutateAuth(true)
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
                type="email"
                placeholder="Email"
                value={email}
                onChange={handleEmailChange}
                size="small"
            />
            {" "}
            <TextField
                id="password"
                label="Password"
                type="password"
                placeholder="Password"
                value={password}
                onChange={handlePasswordChange}
                size="small"
            />
            {" "}
            <Button type="submit" variant="contained" disabled={isLoading}>
                Login
            </Button>
            <br/>
            {errors ? errors : ""}
        </form>
    );
}
