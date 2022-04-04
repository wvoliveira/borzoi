import Layout from "../components/Layout";
import {Link} from "@mui/material";
import Typography from "@mui/material/Typography";
import * as React from "react";
import FormCreateClient from "../components/clients/FormCreateClient";
import TableClients from "../components/clients/TableClients";
import useAuth from "../lib/utils/useAuth";
import {useRouter} from "next/router";

export default function Clients() {
    const router = useRouter();
    const {auth} = useAuth({redirectTo: "/"});

    return (
        <Layout>
            <FormCreateClient />
            <br/>
            <TableClients />
        </Layout>
    );
}
