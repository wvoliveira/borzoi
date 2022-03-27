import * as React from "react";
import Layout from "../components/layout";

export default function Register() {

    const handleSubmit = (event) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);
        console.log({
            email: data.get('email'),
            password: data.get('password'),
        });
    };

    return (
        <Layout>
            Register
            <br/><br/>

            Name: <input /><br/>
            Email: <input/><br/>
            Password: <input/><br/>
        </Layout>
    )
}
