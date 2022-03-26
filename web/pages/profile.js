import * as React from "react";
import Layout from "../components/layout";
import {Fetcher} from "../lib/utils/fetcher";
import useSWR from "swr";

export default function Profile() {
    const { data, error } = useSWR('/api/v1/users/me', Fetcher)
    if (error) {
        return <>error: {error}</>
    }

    if (!data) {
        return <>Loading...</>
    }

    let user = data.data
    return (
        <Layout>
            <>
                ID: {user.id}<br/>
                Name: {user.name}
            </>
        </Layout>
    )
}
