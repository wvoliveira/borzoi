import Layout from "../components/Layout";
import {Link} from "@mui/material";
import Typography from "@mui/material/Typography";
import * as React from "react";

export default function Home() {
    return (
        <Layout>
            <Typography noWrap component="div">
                A client management system.
            </Typography>
        </Layout>
    );
}
