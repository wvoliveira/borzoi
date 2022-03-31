import React, { useState } from "react";
import useUser from "../lib/iron/useUser";
import Layout from "../components/Layout";
import Form from "../components/Form";
import fetchJson, { FetchError } from "../lib/iron/fetchJson";
import FormLogin from "../components/FormLogin";

export default function Login() {
    // here we just check if user is already logged in and redirect to profile
    const { mutateUser } = useUser({
        redirectTo: "/profile",
        redirectIfFound: true,
    });

    const [errorMsg, setErrorMsg] = useState("");

    return (
        <Layout>
            <div className="login">
                <FormLogin />
            </div>
            <style jsx>{`
        .login {
          max-width: 21rem;
          margin: 0 auto;
          padding: 1rem;
          border: 1px solid #ccc;
          border-radius: 4px;
        }
      `}</style>
        </Layout>
    );
}
