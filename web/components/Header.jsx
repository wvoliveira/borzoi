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
                        <a>login</a>
                    </Link>
                    <Link href="/register">
                        <a>register</a>
                    </Link>
                </div>
            )}

            {auth && (
                <div className="menu button">
                    <Link href="/jobs">
                        <a>jobs</a>
                    </Link>
                    <Link href="/clients">
                        <a>clients</a>
                    </Link>
                    <Link href="/profile">
                        <a>profile</a>
                    </Link>
                    <Link href="/logout" onClick={handleLogout}>
                        <a>logout</a>
                    </Link>
                </div>
            )}
        </header>
    );
}
