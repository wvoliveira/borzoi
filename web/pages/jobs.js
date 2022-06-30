import * as React from "react";
import { useRouter } from "next/router";
import Router from "next/router";

import FormCreateClient from "../components/clients/FormCreateClient";
import TableClients from "../components/clients/TableClients";

import useAuth from "../lib/utils/useAuth";

export default function Jobs() {
    const router = useRouter();
    const { auth } = useAuth({ redirectTo: "/" });

    return (
        <h3>Jobs</h3>
    );
}
