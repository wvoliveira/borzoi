import * as React from 'react';
import Typography from '@mui/material/Typography';
import useAuth from "../lib/utils/useAuth";
import {useRouter} from "next/router";
import {AuthAPI} from "../lib/api/auth";
import {LinkStyle} from "./Utils";

export default function Header() {
    const {auth, mutateAuth} = useAuth();
    const router = useRouter();

    const handleLogout = (e) => {
        e.preventDefault();
        AuthAPI.Logout().then(mutateAuth(null) && router.push("/"));
    };

    return (
        <header>
            <Typography variant="h6" noWrap component="div">
                <LinkStyle href="/" text="BORZOI"/>
            </Typography>

            {!auth && (
                <>
                    <Typography>
                        <LinkStyle href="/login" text="Login"/>
                        {" | "}
                        <LinkStyle href="/register" text="Register"/>
                    </Typography>
                </>
            )}

            {auth && (
                <>
                    <Typography>
                        <LinkStyle href="/profile" text="Profile"/>
                        {" | "}
                        <LinkStyle href="/logout" text="Logout" onClick={handleLogout}/>
                    </Typography>
                </>
            )}
        </header>
    );
}
