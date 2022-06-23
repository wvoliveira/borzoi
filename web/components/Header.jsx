import * as React from "react";
import { useRouter } from "next/router";
import Link from "next/link";

import useAuth from "../lib/utils/useAuth";
import { AuthAPI } from "../lib/api/auth";

export default function Header() {
    const { auth, mutateAuth } = useAuth();
    const router = useRouter();

    const handleLogout = (e) => {
        e.preventDefault();
        AuthAPI.Logout().then(mutateAuth(null) && router.push("/"));
    };

    return (
        <header>
            <div className="logo">
                <Link href="/">
                    <a>BORZOI</a>
                </Link>
            </div>

            {!auth && (
                <div className="menu button">
                    <Link href="/login">
                        <a>Login</a>
                    </Link>
                    {" | "}
                    <Link href="/register">
                        <a>Register</a>
                    </Link>
                </div>
            )}

            {auth && (
                <div className="menu button">
                    <Link href="/jobs">
                        <a>Jobs</a>
                    </Link>
                    {" | "}
                    <Link href="/clients">
                        <a>Clients</a>
                    </Link>
                    {" | "}
                    <Link href="/profile">
                        <a>Profile</a>
                    </Link>
                    {" | "}
                    <Link href="/logout" onClick={handleLogout}>
                        <a>Logout</a>
                    </Link>
                </div>
            )}
        </header>
    );
}
