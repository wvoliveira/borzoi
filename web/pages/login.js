import React from "react";
import Layout from "../components/Layout";
import FormLogin from "../components/FormLogin";

export default function Login() {
    return (
        <Layout>
            <div className="login">
                <FormLogin/>
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
