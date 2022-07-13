import * as React from "react";

import JobsTable from "../components/JobsTable";

import useAuth from "../lib/hooks/useAuth";


export default function Jobs() {
    useAuth({
        redirectTo: "/login",
    });

    return (
        <>
            <h2>Jobs</h2>
            <JobsTable />
        </>
    );
}
