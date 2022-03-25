import * as React from "react";
import Layout from '../components/layout';
import FormLogin from '../components/form-login';
import {IsAuthenticated} from "../lib/auth";
import { useRouter } from 'next/router'

export default function Login() {
    const router = useRouter()

    if (IsAuthenticated()) {
        router.push("/").then(r => {});
    }

    return (
        <Layout>
            Login
            <br/>
            <br/>
            <FormLogin/>
        </Layout>
    )
}
