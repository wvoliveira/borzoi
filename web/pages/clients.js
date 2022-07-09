import * as React from "react";
import ClientTable from "../components/ClientTable";
import ClientSection from "../components/ClientSection";

import useAuth from "../lib/hooks/useAuth";


export default function Clients() {
    useAuth({
        redirectTo: "/login",
    });

    return (<>
        <ClientSection />
    </>);
}
