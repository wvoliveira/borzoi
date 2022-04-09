import React from "react";

import { AuthAPI } from "../lib/api/auth";
import useAuth from "../lib/utils/useAuth";
import { TextField, Typography } from "@mui/material";
import { LoadingButton } from "@mui/lab";
import Link from "next/link";

export default function FormLogin() {
    const { mutateAuth } = useAuth({
        redirectTo: "/",
        redirectIfFound: true,
    });

    const [loading, setLoading] = React.useState(false);
    const [error, setError] = React.useState("");
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
            const { data, status } = await AuthAPI.Login(email, password);
            if (status !== 200 && status !== 500) {
                setError(data.message);
                console.log(data.message);
            }

            if (status === 200) {
                mutateAuth(true);
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
                id="email"
                label="Email"
                type="email"
                placeholder="Email"
                value={email}
                onChange={handleEmailChange}
                size="small"
                required
            />{" "}
            <TextField
                id="password"
                label="Password"
                type="password"
                placeholder="Password"
                value={password}
                onChange={handlePasswordChange}
                size="small"
                required
            />{" "}
            <LoadingButton
                type="submit"
                onClick={handleSubmit}
                variant="contained"
                loading={loading}
            >
                Login
            </LoadingButton>{" "}
            {error ? error : ""}
            <br />
            <Typography>
                <Link href="/forgot">
                    <a style={{ textDecoration: "none", color: "#000" }}>
                        Forgot my password.
                    </a>
                </Link>
            </Typography>
        </form>
    );
}
