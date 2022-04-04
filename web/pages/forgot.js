import React from "react";
import Layout from "../components/Layout";
import FormLogin from "../components/FormLogin";
import Box from '@mui/material/Box';
import Link from "next/link";
import FormForgotPassword from "../components/FormForgotPassword";

export default function Login() {
    return (
        <Layout>
            <FormForgotPassword />
        </Layout>
    );
}
