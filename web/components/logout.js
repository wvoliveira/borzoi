import Link from "next/link";
import * as React from "react";

export default function Logout() {
    const handleLogout = async (e) => {
        e.preventDefault();
        console.log("logout clicked!")
    }

    return (
        <Link href={"/logout"} >
            <a onClick={(e) => handleLogout(e, "/logout")}>Logout</a>
        </Link>
    )
}
