import Router from "next/router";
import React from "react";

import {AuthAPI} from "../lib/api/auth";
import {Button, TextField} from "@mui/material";
import {LoadingButton} from "@mui/lab";

export default function FormRegister() {
    const [loading, setLoading] = React.useState(false);
    const [errors, setErrors] = React.useState([]);
    const [name, setName] = React.useState("");
    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");

    const handleName = React.useCallback((e) => setName(e.target.value), []);
    const handleEmail = React.useCallback((e) => setEmail(e.target.value), []);
    const handlePassword = React.useCallback((e) => setPassword(e.target.value), []);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        console.log("name: " + name);
        console.log("email: " + email);
        console.log("password: " + password);

        try {
            const {data, status} = await AuthAPI.Register(name, email, password);
            if (status !== 200 && status !== 500) {
                setErrors(data.message);
                console.log(data.message);
            }

            if (status === 200) {
                await Router.push("/login");
            }
        } catch (error) {
            console.error(error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <form>
            <TextField
                name="name"
                label="Name"
                type="text"
                placeholder="Name"
                value={name}
                onChange={handleName}
                size="small"
                required
            />
            {" "}
            <TextField
                name="email"
                label="Email"
                type="email"
                placeholder="Email"
                value={email}
                onChange={handleEmail}
                size="small"
                required
            />
            {" "}
            <TextField
                name="password"
                label="Password"
                type="password"
                placeholder="Password"
                value={password}
                onChange={handlePassword}
                size="small"
                required
            />
            {" "}
            <LoadingButton type="submit" onClick={handleSubmit} variant="contained" loading={loading}>
                Register
            </LoadingButton>
            {" "}
            {errors ? errors : ""}
        </form>
    );
}
